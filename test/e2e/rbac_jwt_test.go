package e2e_test

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"net/http"
	"sync/atomic"

	envoyutil "github.com/envoyproxy/go-control-plane/pkg/util"
	jwt2 "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/jwt"
	rbac2 "github.com/solo-io/solo-projects/projects/gloo/pkg/plugins/rbac"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/gogo/protobuf/types"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"

	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
	"github.com/solo-io/solo-projects/test/services"

	jwtplugin "github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/jwt"
	"github.com/solo-io/solo-projects/projects/gloo/pkg/api/v1/plugins/rbac"

	"github.com/fgrosse/zaptest"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	gloov1static "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/plugins/static"
	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients/memory"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	"github.com/solo-io/solo-projects/test/v1helpers"
	jose "gopkg.in/square/go-jose.v2"

	"github.com/dgrijalva/jwt-go"
)

var (
	baseJwksPort = uint32(28000)
)

const (
	issuer   = "issuer"
	audience = "thats-us"

	admin = "admin"
	user  = "user"
)

func jwks(ctx context.Context) (uint32, *rsa.PrivateKey) {
	priv, err := rsa.GenerateKey(rand.Reader, 512)
	Expect(err).NotTo(HaveOccurred())
	key := jose.JSONWebKey{
		Key:       priv.Public(),
		Algorithm: "RS256",
		Use:       "sig",
	}

	keySet := jose.JSONWebKeySet{
		Keys: []jose.JSONWebKey{key},
	}

	jwksBytes, err := json.Marshal(keySet)
	Expect(err).NotTo(HaveOccurred())

	jwksPort := atomic.AddUint32(&baseJwksPort, 1) + uint32(config.GinkgoConfig.ParallelNode*1000)
	jwtHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("content-type", "application/json")
		w.Write(jwksBytes)
	}
	h2s := &http2.Server{}

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", jwksPort),
		Handler: h2c.NewHandler(http.HandlerFunc(jwtHandler), h2s),
	}

	go s.ListenAndServe()
	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	// serialize json and show
	return jwksPort, priv
}

func getToken(claims jwt.StandardClaims, key *rsa.PrivateKey) string {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	s, err := token.SignedString(key)
	Expect(err).NotTo(HaveOccurred())
	return s
}

var _ = Describe("JWT + RBAC", func() {

	var (
		ctx            context.Context
		cancel         context.CancelFunc
		testClients    services.TestClients
		jwksPort       uint32
		privateKey     *rsa.PrivateKey
		jwtksServerRef core.ResourceRef
	)

	BeforeEach(func() {

		logger := zaptest.LoggerWriter(GinkgoWriter)
		contextutils.SetFallbackLogger(logger.Sugar())

		ctx, cancel = context.WithCancel(context.Background())
		cache := memory.NewInMemoryResourceCache()

		testClients = services.GetTestClients(cache)
		testClients.GlooPort = int(services.AllocateGlooPort())

		jwksPort, privateKey = jwks(ctx)

		jwtksServer := &gloov1.Upstream{
			Metadata: core.Metadata{
				Name:      "jwks-server",
				Namespace: "default",
			},
			UpstreamSpec: &gloov1.UpstreamSpec{
				UpstreamType: &gloov1.UpstreamSpec_Static{
					Static: &gloov1static.UpstreamSpec{
						Hosts: []*gloov1static.Host{{
							Addr: "localhost",
							Port: jwksPort,
						}},
						UseHttp2: true,
					},
				},
			},
		}

		testClients.UpstreamClient.Write(jwtksServer, clients.WriteOpts{})
		jwtksServerRef = jwtksServer.Metadata.Ref()
		rbacSettings := &rbac.Settings{
			RequireRbac: true,
		}
		settingsStruct, err := envoyutil.MessageToStruct(rbacSettings)
		Expect(err).NotTo(HaveOccurred())

		extensions := &gloov1.Extensions{
			Configs: map[string]*types.Struct{
				rbac2.ExtensionName: settingsStruct,
			},
		}

		what := services.What{
			DisableGateway: true,
			DisableUds:     true,
			DisableFds:     true,
		}

		services.RunGlooGatewayUdsFdsOnPort(ctx, cache, int32(testClients.GlooPort), what, defaults.GlooSystem, nil, extensions)
	})

	AfterEach(func() {
		cancel()
	})

	Context("With envoy", func() {

		var (
			envoyInstance *services.EnvoyInstance
			testUpstream  *v1helpers.TestUpstream
			envoyPort     = uint32(8080)
		)

		BeforeEach(func() {
			var err error
			envoyInstance, err = envoyFactory.NewEnvoyInstance()
			Expect(err).NotTo(HaveOccurred())

			err = envoyInstance.Run(testClients.GlooPort)
			Expect(err).NotTo(HaveOccurred())

			testUpstream = v1helpers.NewTestHttpUpstream(ctx, envoyInstance.LocalAddr())

			var opts clients.WriteOpts
			up := testUpstream.Upstream
			_, err = testClients.UpstreamClient.Write(up, opts)
			Expect(err).NotTo(HaveOccurred())
		})

		AfterEach(func() {
			if envoyInstance != nil {
				envoyInstance.Clean()
			}
		})

		addToken := func(req *http.Request, sub string) {
			claims := jwt.StandardClaims{
				Issuer:   issuer,
				Audience: audience,
				Subject:  sub,
			}
			token := getToken(claims, privateKey)
			req.Header.Add("Authorization", "Bearer "+token)
		}

		Context("user access tests", func() {
			BeforeEach(func() {

				// drain channel as we dont care about it
				go func() {
					for range testUpstream.C {
					}
				}()

				envoyPort += 1

				proxy := getProxyJwtRbac(envoyPort, jwtksServerRef, testUpstream.Upstream.Metadata.Ref())

				_, err := testClients.ProxyClient.Write(proxy, clients.WriteOpts{})
				Expect(err).NotTo(HaveOccurred())

				Eventually(func() (core.Status, error) {
					proxy, err := testClients.ProxyClient.Read(proxy.Metadata.Namespace, proxy.Metadata.Name, clients.ReadOpts{})
					if err != nil {
						return core.Status{}, err
					}

					return proxy.Status, nil
				}, "5s", "0.1s").Should(MatchFields(IgnoreExtras, Fields{
					"Reason": BeEmpty(),
					"State":  Equal(core.Status_Accepted),
				}))

				// wait for key service to start
				Eventually(func() error {
					_, err := http.Get(fmt.Sprintf("http://%s:%d/", "localhost", jwksPort))
					return err
				}, "5s", "0.5s").ShouldNot(HaveOccurred())

			})

			Context("non admin user", func() {

				It("should deny to /bar", func() {
					Eventually(func() (int, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/bar", "localhost", envoyPort), nil)
						addToken(req, "user")
						if err != nil {
							return 0, err
						}

						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusForbidden))
				})

				It("should deny to /foo with POST and allow /foo with GET", func() {
					Eventually(func() (int, error) {
						req, err := http.NewRequest("POST", fmt.Sprintf("http://%s:%d/foo", "localhost", envoyPort), nil)
						addToken(req, "user")
						if err != nil {
							return 0, err
						}

						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusForbidden))

					// All is initialized so this should work consistently
					Consistently(func() (int, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/foo", "localhost", envoyPort), nil)
						addToken(req, "user")
						if err != nil {
							return 0, err
						}

						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "1s", "0.1s").Should(Equal(http.StatusOK))
				})
			})

			Context("admin user", func() {
				It("should allow everything", func() {
					Eventually(func() (int, error) {
						req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/bar", "localhost", envoyPort), nil)
						addToken(req, "admin")
						if err != nil {
							return 0, err
						}

						resp, err := http.DefaultClient.Do(req)
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusOK))

					// No need to do eventually here as all is initialized.
					foo := fmt.Sprintf("http://%s:%d/foo", "localhost", envoyPort)

					req, err := http.NewRequest("POST", foo, nil)
					addToken(req, "admin")
					Expect(err).NotTo(HaveOccurred())
					resp, err := http.DefaultClient.Do(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(http.StatusOK))

					req, err = http.NewRequest("GET", foo, nil)
					addToken(req, "admin")
					Expect(err).NotTo(HaveOccurred())
					resp, err = http.DefaultClient.Do(req)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
				})
			})

			Context("anonymous user", func() {
				It("should not allow anything other than the public route", func() {
					Eventually(func() (int, error) {
						resp, err := http.Get(fmt.Sprintf("http://%s:%d/bar", "localhost", envoyPort))
						if err != nil {
							return 0, err
						}
						return resp.StatusCode, nil
					}, "5s", "0.5s").Should(Equal(http.StatusUnauthorized))

					// No need to do eventually here as all is initialized.
					foo := fmt.Sprintf("http://%s:%d/foo", "localhost", envoyPort)

					resp, err := http.Post(foo, "application/json", nil)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))

					resp, err = http.Get(foo)
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(http.StatusUnauthorized))

					resp, err = http.Get(fmt.Sprintf("http://%s:%d/public_route", "localhost", envoyPort))
					Expect(err).NotTo(HaveOccurred())
					Expect(resp.StatusCode).To(Equal(http.StatusOK))
				})
			})

		})
	})
})

func getProxyJwtRbac(envoyPort uint32, jwtksServerRef, upstream core.ResourceRef) *gloov1.Proxy {

	jwtCfg := &jwtplugin.VhostExtension{
		JwksUrl:         "http://test/keys",
		JwksUpstreamRef: &jwtksServerRef,
		Issuer:          issuer,
		Audiences:       []string{audience},
	}

	rbacCfg := &rbac.VhostExtension{
		Config: &rbac.Config{
			Policies: map[string]*rbac.Policy{
				"user": {
					Principals: []*rbac.Principal{{
						JwtPrincipal: &rbac.JWTPrincipal{
							Claims: map[string]string{
								"iss": issuer,
								"sub": user,
							},
						},
					}},
					Permissions: &rbac.Permissions{
						PathPrefix: "/foo",
						Methods:    []string{"GET"},
					},
				},
				"admin": {
					Principals: []*rbac.Principal{{
						JwtPrincipal: &rbac.JWTPrincipal{
							Claims: map[string]string{
								"iss": issuer,
								"sub": admin,
							},
						},
					}},
					Permissions: &rbac.Permissions{},
				},
			},
		},
	}

	return getProxyJwtRbacWithExtensions(envoyPort, jwtksServerRef, upstream, jwtCfg, rbacCfg)
}

func getProxyJwtRbacWithExtensions(envoyPort uint32, jwtksServerRef, upstream core.ResourceRef, jwtCfg *jwtplugin.VhostExtension, rbacCfg *rbac.VhostExtension) *gloov1.Proxy {
	var extensions *gloov1.Extensions

	jwtStruct, err := envoyutil.MessageToStruct(jwtCfg)
	Expect(err).NotTo(HaveOccurred())
	rbacStruct, err := envoyutil.MessageToStruct(rbacCfg)
	Expect(err).NotTo(HaveOccurred())
	protos := map[string]*types.Struct{
		jwt2.ExtensionName:  jwtStruct,
		rbac2.ExtensionName: rbacStruct,
	}
	extensions = &gloov1.Extensions{
		Configs: protos,
	}

	var vhosts []*gloov1.VirtualHost

	vhost := &gloov1.VirtualHost{
		Name:               "virt1",
		Domains:            []string{"*"},
		VirtualHostPlugins: &gloov1.VirtualHostPlugins{},
		Routes: []*gloov1.Route{
			{
				RoutePlugins: &gloov1.RoutePlugins{
					Extensions: getDisabled(),
				},
				Matcher: &gloov1.Matcher{
					PathSpecifier: &gloov1.Matcher_Prefix{
						Prefix: "/public_route",
					},
				},
				Action: &gloov1.Route_RouteAction{
					RouteAction: &gloov1.RouteAction{
						Destination: &gloov1.RouteAction_Single{
							Single: &gloov1.Destination{
								Upstream: upstream,
							},
						},
					},
				},
			}, {
				Matcher: &gloov1.Matcher{
					PathSpecifier: &gloov1.Matcher_Prefix{
						Prefix: "/",
					},
				},
				Action: &gloov1.Route_RouteAction{
					RouteAction: &gloov1.RouteAction{
						Destination: &gloov1.RouteAction_Single{
							Single: &gloov1.Destination{
								Upstream: upstream,
							},
						},
					},
				},
			}},
	}

	vhost.VirtualHostPlugins.Extensions = extensions

	vhosts = append(vhosts, vhost)

	p := &gloov1.Proxy{
		Metadata: core.Metadata{
			Name:      "proxy",
			Namespace: "default",
		},
		Listeners: []*gloov1.Listener{{
			Name:        "listener",
			BindAddress: "127.0.0.1",
			BindPort:    envoyPort,
			ListenerType: &gloov1.Listener_HttpListener{
				HttpListener: &gloov1.HttpListener{
					VirtualHosts: vhosts,
				},
			},
		}},
	}

	return p
}

func getDisabled() *gloov1.Extensions {

	jwtCfg := &jwtplugin.RouteExtension{
		Disable: true,
	}
	rbacCfg := &rbac.RouteExtension{
		Route: &rbac.RouteExtension_Disable{
			Disable: true,
		},
	}
	jwtStruct, err := envoyutil.MessageToStruct(jwtCfg)
	Expect(err).NotTo(HaveOccurred())
	rbacStruct, err := envoyutil.MessageToStruct(rbacCfg)
	Expect(err).NotTo(HaveOccurred())
	protos := map[string]*types.Struct{
		jwt2.ExtensionName:  jwtStruct,
		rbac2.ExtensionName: rbacStruct,
	}
	return &gloov1.Extensions{
		Configs: protos,
	}
}

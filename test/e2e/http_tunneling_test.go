package e2e_test

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/onsi/gomega/types"

	"github.com/solo-io/gloo/test/v1helpers"

	"github.com/golang/protobuf/ptypes/wrappers"

	static_plugin_gloo "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options/static"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	gatewayv1 "github.com/solo-io/gloo/projects/gateway/pkg/api/v1"
	gatewaydefaults "github.com/solo-io/gloo/projects/gateway/pkg/defaults"
	gloohelpers "github.com/solo-io/gloo/test/helpers"

	"github.com/solo-io/gloo/projects/gloo/pkg/defaults"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	gloov1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	"github.com/solo-io/gloo/test/services"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

var _ = Describe("tunneling", func() {

	var (
		ctx            context.Context
		cancel         context.CancelFunc
		testClients    services.TestClients
		envoyInstance  *services.EnvoyInstance
		up             *gloov1.Upstream
		tuPort         uint32
		tlsUpstream    bool
		writeNamespace = defaults.GlooSystem
	)

	checkProxy := func() {
		// ensure the proxy is created
		Eventually(func() (*gloov1.Proxy, error) {
			return testClients.ProxyClient.Read(writeNamespace, gatewaydefaults.GatewayProxyName, clients.ReadOpts{})
		}, "5s", "0.1s").ShouldNot(BeNil())
	}

	checkVirtualService := func(testVs *gatewayv1.VirtualService) {
		Eventually(func() (*gatewayv1.VirtualService, error) {
			return testClients.VirtualServiceClient.Read(testVs.Metadata.GetNamespace(), testVs.Metadata.GetName(), clients.ReadOpts{})
		}, "5s", "0.1s").ShouldNot(BeNil())
	}

	BeforeEach(func() {
		tlsUpstream = false
		var err error
		ctx, cancel = context.WithCancel(context.Background())
		defaults.HttpPort = services.NextBindPort()

		// run gloo
		ro := &services.RunOptions{
			NsToWrite: writeNamespace,
			NsToWatch: []string{"default", writeNamespace},
			WhatToRun: services.What{
				DisableFds: true,
				DisableUds: true,
			},
		}
		testClients = services.RunGlooGatewayUdsFds(ctx, ro)

		// write gateways and wait for them to be created
		err = gloohelpers.WriteDefaultGateways(writeNamespace, testClients.GatewayClient)
		Expect(err).NotTo(HaveOccurred(), "Should be able to write default gateways")
		Eventually(func() (gatewayv1.GatewayList, error) {
			return testClients.GatewayClient.List(writeNamespace, clients.ListOpts{})
		}, "10s", "0.1s").Should(HaveLen(2), "Gateways should be present")

		// run envoy
		envoyInstance, err = envoyFactory.NewEnvoyInstance()
		Expect(err).NotTo(HaveOccurred())
		err = envoyInstance.RunWithRoleAndRestXds(writeNamespace+"~"+gatewaydefaults.GatewayProxyName, testClients.GlooPort, testClients.RestXdsPort)
		Expect(err).NotTo(HaveOccurred())
	})

	JustBeforeEach(func() {
		// start http proxy and setup upstream that points to it
		port := startHttpProxy(ctx)

		tu := v1helpers.NewTestHttpUpstreamWithTls(ctx, envoyInstance.LocalAddr(), tlsUpstream)
		tuPort = tu.Upstream.UpstreamType.(*gloov1.Upstream_Static).Static.Hosts[0].Port

		up = &gloov1.Upstream{
			Metadata: &core.Metadata{
				Name:      "local-1",
				Namespace: "default",
			},
			UpstreamType: &gloov1.Upstream_Static{
				Static: &static_plugin_gloo.UpstreamSpec{
					Hosts: []*static_plugin_gloo.Host{
						{
							Addr: envoyInstance.LocalAddr(),
							Port: uint32(port),
						},
					},
				},
			},
			HttpProxyHostname: &wrappers.StringValue{Value: fmt.Sprintf("%s:%d", envoyInstance.LocalAddr(), tuPort)}, // enable HTTP tunneling,
		}

		// write a virtual service so we have a proxy to our test upstream
		testVs := getTrivialVirtualServiceForUpstream(writeNamespace, up.Metadata.Ref())
		_, err := testClients.VirtualServiceClient.Write(testVs, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true})
		Expect(err).NotTo(HaveOccurred())
		checkVirtualService(testVs)
	})

	AfterEach(func() {
		envoyInstance.Clean()
		cancel()
	})

	expectResponseBodyOnRequest := func(requestJsonBody string, expectedResponseStatusCode int, expectedResponseBodyMatcher types.GomegaMatcher) {
		EventuallyWithOffset(1, func() (string, error) {
			var client http.Client
			scheme := "http"
			var json = []byte(requestJsonBody)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s://%s:%d/test", scheme, "localhost", defaults.HttpPort), bytes.NewBuffer(json))
			if err != nil {
				return "", err
			}
			res, err := client.Do(req)
			if err != nil {
				return "", err
			}
			if res.StatusCode != expectedResponseStatusCode {
				return "", fmt.Errorf("not ok")
			}
			p := new(bytes.Buffer)
			if _, err := io.Copy(p, res.Body); err != nil {
				return "", err
			}
			defer res.Body.Close()
			return p.String(), nil
		}, "10s", "0.5s").Should(expectedResponseBodyMatcher)
	}

	Context("plaintext", func() {

		JustBeforeEach(func() {
			_, err := testClients.UpstreamClient.Write(up, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true})
			Expect(err).NotTo(HaveOccurred())

			checkProxy()
		})

		It("should proxy http", func() {
			// the request path here is envoy -> local HTTP proxy (HTTP CONNECT) -> test upstream
			// and back. The HTTP proxy is sending unencrypted HTTP bytes over
			// TCP to the test upstream (an echo server)
			jsonStr := `{"value":"Hello, world!"}`
			expectResponseBodyOnRequest(jsonStr, http.StatusOK, ContainSubstring(jsonStr))
		})
	})

	Context("with TLS", func() {

		JustBeforeEach(func() {

			secret := &gloov1.Secret{
				Metadata: &core.Metadata{
					Name:      "secret",
					Namespace: "default",
				},
				Kind: &gloov1.Secret_Tls{
					Tls: &gloov1.TlsSecret{
						CertChain:  gloohelpers.Certificate(),
						PrivateKey: gloohelpers.PrivateKey(),
						RootCa:     gloohelpers.Certificate(),
					},
				},
			}

			_, err := testClients.SecretClient.Write(secret, clients.WriteOpts{OverwriteExisting: true})
			Expect(err).NotTo(HaveOccurred())

			sslCfg := &gloov1.UpstreamSslConfig{
				SslSecrets: &gloov1.UpstreamSslConfig_SecretRef{
					SecretRef: &core.ResourceRef{Name: "secret", Namespace: "default"},
				},
			}

			if tlsUpstream {
				up.SslConfig = sslCfg
			}
			up.HttpProxyHostname = &wrappers.StringValue{Value: fmt.Sprintf("%s:%d", envoyInstance.LocalAddr(), tuPort)} // enable HTTP tunneling,
			_, err = testClients.UpstreamClient.Write(up, clients.WriteOpts{Ctx: ctx, OverwriteExisting: true})
			Expect(err).NotTo(HaveOccurred())

			checkProxy()
		})

		Context("with back TLS", func() {

			BeforeEach(func() {
				tlsUpstream = true
			})

			It("should proxy encrypted bytes over plaintext HTTP Connect", func() {
				// the request path here is [envoy] -- plaintext --> [local HTTP Connect proxy] -- encrypted --> TLS upstream
				jsonStr := `{"value":"Hello, world!"}`
				expectResponseBodyOnRequest(jsonStr, http.StatusOK, ContainSubstring(jsonStr))
			})
		})
	})
})

func startHttpProxy(ctx context.Context) int {
	listener, err := net.Listen("tcp", ":0")
	Expect(err).ToNot(HaveOccurred())

	addr := listener.Addr().String()
	_, portStr, err := net.SplitHostPort(addr)
	Expect(err).ToNot(HaveOccurred())

	port, err := strconv.Atoi(portStr)
	Expect(err).ToNot(HaveOccurred())

	fmt.Fprintln(GinkgoWriter, "go proxy addr", addr)

	go func() {
		defer GinkgoRecover()
		server := &http.Server{Addr: addr, Handler: http.HandlerFunc(connectProxy)}
		server.Serve(listener)
		<-ctx.Done()
		server.Close()
	}()

	return port
}

func isEof(r *bufio.Reader) bool {
	_, err := r.Peek(1)
	if errors.Is(err, io.EOF) {
		return true
	}
	return false
}

func connectProxy(w http.ResponseWriter, r *http.Request) {
	if r.Method != "CONNECT" {
		http.Error(w, "not connect", 400)
		return
	}

	if r.TLS != nil {
		fmt.Fprintf(GinkgoWriter, "handshake complete %v\n", r.TLS.HandshakeComplete)
		fmt.Fprintf(GinkgoWriter, "tls version %v\n", r.TLS.Version)
		fmt.Fprintf(GinkgoWriter, "cipher suite %v\n", r.TLS.CipherSuite)
		fmt.Fprintf(GinkgoWriter, "negotiated protocol %v\n", r.TLS.NegotiatedProtocol)
	}

	if proxyAuth := r.Header.Get("Proxy-Authorization"); proxyAuth != "" {
		fmt.Fprintf(GinkgoWriter, "proxy authorization: %s\n", proxyAuth)
		if username, password := parseBasicAuth(proxyAuth); username != "test" || password != "secret" {
			w.WriteHeader(http.StatusProxyAuthRequired)
			return
		}
	}

	hij, ok := w.(http.Hijacker)
	if !ok {
		Fail("no hijacker")
	}
	host := r.URL.Host

	targetConn, err := net.Dial("tcp", host)
	if err != nil {
		http.Error(w, "can't connect", 500)
		return
	}
	defer targetConn.Close()

	conn, buf, err := hij.Hijack()
	if err != nil {
		Expect(err).ToNot(HaveOccurred())
	}
	defer conn.Close()

	fmt.Fprintf(GinkgoWriter, "Accepting CONNECT to %s\n", host)
	// note to devs! will only work with HTTP 1.1 request from envoy!
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))

	// now just copy:
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer GinkgoRecover()
		for {
			// read bytes from buf.Reader until EOF
			bts := []byte{1}
			_, err := targetConn.Read(bts)
			if errors.Is(err, io.EOF) {
				break
			}
			Expect(err).NotTo(HaveOccurred())
			_, err = conn.Write(bts)
			if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, syscall.EPIPE) {
				fmt.Fprintf(GinkgoWriter, "error writing from upstream to envoy %v\n", err)
				Fail("error writing from upstream to envoy")
			}
		}
		err = buf.Flush()
		Expect(err).NotTo(HaveOccurred())
		wg.Done()
	}()
	go func() {
		defer GinkgoRecover()
		for !isEof(buf.Reader) {
			// read bytes from buf.Reader until EOF
			bts := []byte{1}
			_, err := buf.Read(bts)
			Expect(err).NotTo(HaveOccurred())
			_, err = targetConn.Write(bts)
			if err != nil && !errors.Is(err, io.EOF) && !errors.Is(err, syscall.EPIPE) {
				fmt.Fprintf(GinkgoWriter, "error writing from envoy to upstream %v\n", err)
				Fail("error writing from envoy to upstream")
			}
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Fprintf(GinkgoWriter, "done proxying\n")
}

func parseBasicAuth(auth string) (string, string) {
	const basicPrefix = "Basic "
	if !strings.HasPrefix(auth, basicPrefix) {
		return "", ""
	}
	decodedAuth, err := base64.StdEncoding.DecodeString(auth[len(basicPrefix):])
	if err != nil {
		return "", ""
	}
	decodedAuthString := string(decodedAuth)
	username, password, ok := strCut(decodedAuthString, ":")
	if !ok {
		return "", ""
	}
	return username, password
}

// COPIED FROM strings.Cut ADDED IN GO 1.18
// Cut slices s around the first instance of sep,
// returning the text before and after sep.
// The found result reports whether sep appears in s.
// If sep does not appear in s, cut returns s, "", false.
func strCut(s, sep string) (before, after string, found bool) {
	if i := strings.Index(s, sep); i >= 0 {
		return s[:i], s[i+len(sep):], true
	}
	return s, "", false
}

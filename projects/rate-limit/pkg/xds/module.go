package xds

import (
	"context"
	"os"

	solo_api_rl "github.com/solo-io/solo-apis/pkg/api/ratelimit.solo.io/v1alpha1"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_struct "github.com/golang/protobuf/ptypes/struct"
	"github.com/rotisserie/eris"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise"
	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/rate-limiter/pkg/modules"
	ratelimit "github.com/solo-io/rate-limiter/pkg/service"
	"github.com/solo-io/solo-projects/projects/rate-limit/pkg/shims"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const ModuleName = "xDS"

var ConfigGenErr = func(err error, domain string) error {
	return eris.Wrapf(err, "failed to generate configuration for domain [%s]", domain)
}

// This `RunnableModule` implementation uses xDS to get rate limit configuration from the GlooE control plane.
type configSource struct {
	glooAddress     string
	domainGenerator shims.RateLimitDomainGenerator
}

func NewConfigSource(settings Settings, domainGenerator shims.RateLimitDomainGenerator) modules.RunnableModule {
	return &configSource{
		glooAddress:     settings.GlooAddress,
		domainGenerator: domainGenerator,
	}
}

func (*configSource) Name() string {
	return ModuleName
}

func (x *configSource) Run(ctx context.Context, service ratelimit.RateLimitServiceServer) error {
	var nodeinfo core.Node
	var err error
	nodeinfo.Id, err = os.Hostname()
	// TODO(yuval-k): unhardcode this
	if err != nil {
		nodeinfo.Id = "ratelimit-unknown"
	}
	nodeinfo.Cluster = "ratelimit"
	role := "ratelimit"
	nodeinfo.Metadata = &_struct.Struct{
		Fields: map[string]*_struct.Value{
			"role": {
				Kind: &_struct.Value_StringValue{
					StringValue: role,
				},
			},
		},
	}

	xDSClientLoopFunc := func(ctx context.Context) error {
		logger := contextutils.LoggerFrom(ctx)

		client := v1.NewRateLimitConfigClient(&nodeinfo, func(version string, resources []*v1.RateLimitConfig) error {
			logger.Debugw("received new rate limit config", zap.Any("config", resources))

			for _, cfg := range resources {
				domain, err := x.domainGenerator.NewRateLimitDomain(ctx, cfg.Domain,
					&solo_api_rl.RateLimitConfigSpec_Raw{Descriptors: cfg.Descriptors, SetDescriptors: cfg.SetDescriptors})
				if err != nil {
					logger.Errorw("failed to generate configuration for domain", zap.String("domain", cfg.Domain), zap.Error(err))
					return ConfigGenErr(err, cfg.Domain)
				}
				service.SetDomain(domain)
			}
			return nil
		})

		// We are using non secure gRPC to gloo with the assumption that it will be
		// secured by envoy. if this assumption is not correct this needs to change.
		conn, err := grpc.DialContext(ctx, x.glooAddress, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			logger.Errorw("failed to establish connection to gloo xDS server", zap.String("glooAddress", x.glooAddress), zap.Error(err))
			return err
		}

		// TODO(yuval-k): a stat that indicates we are connected, with the reverse one deferred.
		// TODO(yuval-k): write a warning log
		return client.Start(ctx, conn)
	}

	err = contextutils.NewExponentioalBackoff(contextutils.ExponentioalBackoff{}).Backoff(ctx, xDSClientLoopFunc)
	if err == context.Canceled {
		return nil
	}
	return err
}

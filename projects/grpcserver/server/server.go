package server

import (
	"context"
	"net"

	"github.com/solo-io/solo-projects/projects/grpcserver/server/internal/client"

	"github.com/solo-io/go-utils/contextutils"
	v1 "github.com/solo-io/solo-projects/projects/grpcserver/api/v1"
	"google.golang.org/grpc"
)

type GlooGrpcService struct {
	server   *grpc.Server
	listener net.Listener
}

func NewGlooGrpcService(
	listener net.Listener,
	upstreamService v1.UpstreamApiServer,
	artifactService v1.ArtifactApiServer,
	configService v1.ConfigApiServer,
	secretService v1.SecretApiServer,
	virtualService v1.VirtualServiceApiServer,
	gatewayService v1.GatewayApiServer,
	proxyService v1.ProxyApiServer,
	envoyService v1.EnvoyApiServer,
	clientUpdater client.Updater) (*GlooGrpcService, error) {
	server := &GlooGrpcService{
		server:   grpc.NewServer(),
		listener: listener,
	}

	v1.RegisterUpstreamApiServer(server.server, upstreamService)
	v1.RegisterArtifactApiServer(server.server, artifactService)
	v1.RegisterConfigApiServer(server.server, configService)
	v1.RegisterSecretApiServer(server.server, secretService)
	v1.RegisterVirtualServiceApiServer(server.server, virtualService)
	v1.RegisterGatewayApiServer(server.server, gatewayService)
	v1.RegisterProxyApiServer(server.server, proxyService)
	v1.RegisterEnvoyApiServer(server.server, envoyService)

	// just responsible for kicking off the settings watch loop that rebuilds all the clients
	// (the client updater has to be passed somewhere, otherwise wire complains about an unused provider)
	err := clientUpdater.StartWatch()
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (s *GlooGrpcService) Run(ctx context.Context) error {
	contextutils.LoggerFrom(ctx).Infow("Starting gloo grpc server")
	return s.server.Serve(s.listener)
}

func (s *GlooGrpcService) Stop() {
	s.server.Stop()
}

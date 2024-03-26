package xds

import (
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
)

type DiscoveryInputs struct {
	Upstreams v1.UpstreamList
	Endpoints v1.EndpointList
}

type SecretInputs struct {
	Secrets v1.SecretList
}

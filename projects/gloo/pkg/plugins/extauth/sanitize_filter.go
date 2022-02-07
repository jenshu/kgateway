package extauth

import (
	envoy_config_route_v3 "github.com/envoyproxy/go-control-plane/envoy/config/route/v3"
	. "github.com/solo-io/gloo/projects/gloo/pkg/api/external/envoy/extensions/extauth"
	extauthapi "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/enterprise/options/extauth/v1"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins"
	"github.com/solo-io/gloo/projects/gloo/pkg/plugins/pluginutils"
)

const SanitizeFilterName = "io.solo.filters.http.sanitize"

func buildSanitizeFilter(userIdHeader string, includeCustomAuthServiceName bool) (plugins.StagedHttpFilter, error) {
	sanitizeConf := &Sanitize{}

	if userIdHeader != "" {
		sanitizeConf.HeadersToRemove = []string{userIdHeader}
	}

	if includeCustomAuthServiceName {
		sanitizeConf.CustomAuthServerName = DefaultAuthServiceName
	}

	return plugins.NewStagedFilterWithConfig(SanitizeFilterName, sanitizeConf, sanitizeFilterStage)
}

func setVirtualHostCustomAuth(out *envoy_config_route_v3.VirtualHost, extAuthExtension *extauthapi.ExtAuthExtension, availableCustomAuth map[string]*extauthapi.Settings) error {
	customAuthConfig := buildSanitizePerRouteConfig(extAuthExtension, availableCustomAuth)
	if customAuthConfig == nil {
		return nil
	}
	return pluginutils.SetVhostPerFilterConfig(out, SanitizeFilterName, customAuthConfig)
}

func setWeightedClusterCustomAuth(out *envoy_config_route_v3.WeightedCluster_ClusterWeight, extAuthExtension *extauthapi.ExtAuthExtension, availableCustomAuth map[string]*extauthapi.Settings) error {
	customAuthConfig := buildSanitizePerRouteConfig(extAuthExtension, availableCustomAuth)
	if customAuthConfig == nil {
		return nil
	}
	return pluginutils.SetWeightedClusterPerFilterConfig(out, SanitizeFilterName, customAuthConfig)
}

func setRouteCustomAuth(out *envoy_config_route_v3.Route, extAuthExtension *extauthapi.ExtAuthExtension, availableCustomAuth map[string]*extauthapi.Settings) error {
	customAuthConfig := buildSanitizePerRouteConfig(extAuthExtension, availableCustomAuth)
	if customAuthConfig == nil {
		return nil
	}
	return pluginutils.SetRoutePerFilterConfig(out, SanitizeFilterName, customAuthConfig)
}

func buildSanitizePerRouteConfig(extAuthExtension *extauthapi.ExtAuthExtension, availableCustomAuth map[string]*extauthapi.Settings) *SanitizePerRoute {
	customAuth := extAuthExtension.GetCustomAuth()

	// we have been passed an AuthConfig
	if extAuthExtension.GetConfigRef() != nil {
		return &SanitizePerRoute{CustomAuthServerName: DefaultAuthServiceName}
	}
	// we don't have a route-level customAuth.  Default to higher-order config
	if customAuth == nil || availableCustomAuth == nil {
		return nil
	}

	// we have a route-level customAuth...
	customAuthServiceName := customAuth.GetName()
	if customAuthServiceName == "" {
		// ...which is unnamed.  So we name it our default value.  This is for backwards compatibility, as, originally, customAuth did not require to be named
		customAuthServiceName = DefaultAuthServiceName
	} else if _, ok := availableCustomAuth[customAuthServiceName]; !ok {
		// ...which is not found.  Default to higher-order config
		return nil
	}
	return &SanitizePerRoute{
		CustomAuthServerName: customAuthServiceName,
	}
}

// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway/api/v1/matchable_tcp_gateway.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3"

	github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl "github.com/solo-io/gloo/projects/controller/pkg/api/v1/ssl"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = clone.Cloner(nil)
	_ = proto.Message(nil)
)

// Clone function
func (m *MatchableTcpGateway) Clone() proto.Message {
	var target *MatchableTcpGateway
	if m == nil {
		return target
	}
	target = &MatchableTcpGateway{}

	if h, ok := interface{}(m.GetNamespacedStatuses()).(clone.Cloner); ok {
		target.NamespacedStatuses = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	} else {
		target.NamespacedStatuses = proto.Clone(m.GetNamespacedStatuses()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.NamespacedStatuses)
	}

	if h, ok := interface{}(m.GetMetadata()).(clone.Cloner); ok {
		target.Metadata = h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	} else {
		target.Metadata = proto.Clone(m.GetMetadata()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.Metadata)
	}

	if h, ok := interface{}(m.GetMatcher()).(clone.Cloner); ok {
		target.Matcher = h.Clone().(*MatchableTcpGateway_Matcher)
	} else {
		target.Matcher = proto.Clone(m.GetMatcher()).(*MatchableTcpGateway_Matcher)
	}

	if h, ok := interface{}(m.GetTcpGateway()).(clone.Cloner); ok {
		target.TcpGateway = h.Clone().(*TcpGateway)
	} else {
		target.TcpGateway = proto.Clone(m.GetTcpGateway()).(*TcpGateway)
	}

	return target
}

// Clone function
func (m *MatchableTcpGateway_Matcher) Clone() proto.Message {
	var target *MatchableTcpGateway_Matcher
	if m == nil {
		return target
	}
	target = &MatchableTcpGateway_Matcher{}

	if m.GetSourcePrefixRanges() != nil {
		target.SourcePrefixRanges = make([]*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange, len(m.GetSourcePrefixRanges()))
		for idx, v := range m.GetSourcePrefixRanges() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SourcePrefixRanges[idx] = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			} else {
				target.SourcePrefixRanges[idx] = proto.Clone(v).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_config_core_v3.CidrRange)
			}

		}
	}

	if h, ok := interface{}(m.GetSslConfig()).(clone.Cloner); ok {
		target.SslConfig = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.SslConfig)
	} else {
		target.SslConfig = proto.Clone(m.GetSslConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_v1_ssl.SslConfig)
	}

	if m.GetPassthroughCipherSuites() != nil {
		target.PassthroughCipherSuites = make([]string, len(m.GetPassthroughCipherSuites()))
		for idx, v := range m.GetPassthroughCipherSuites() {

			target.PassthroughCipherSuites[idx] = v

		}
	}

	return target
}

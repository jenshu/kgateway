// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/trace/v3/zipkin.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_solo_kit_pkg_api_v1_resources_core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *ZipkinConfig) Clone() proto.Message {
	var target *ZipkinConfig
	if m == nil {
		return target
	}
	target = &ZipkinConfig{}

	target.CollectorEndpoint = m.GetCollectorEndpoint()

	if h, ok := interface{}(m.GetTraceId_128Bit()).(clone.Cloner); ok {
		target.TraceId_128Bit = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.TraceId_128Bit = proto.Clone(m.GetTraceId_128Bit()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if h, ok := interface{}(m.GetSharedSpanContext()).(clone.Cloner); ok {
		target.SharedSpanContext = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.SharedSpanContext = proto.Clone(m.GetSharedSpanContext()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	target.CollectorEndpointVersion = m.GetCollectorEndpointVersion()

	switch m.CollectorCluster.(type) {

	case *ZipkinConfig_CollectorUpstreamRef:

		if h, ok := interface{}(m.GetCollectorUpstreamRef()).(clone.Cloner); ok {
			target.CollectorCluster = &ZipkinConfig_CollectorUpstreamRef{
				CollectorUpstreamRef: h.Clone().(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		} else {
			target.CollectorCluster = &ZipkinConfig_CollectorUpstreamRef{
				CollectorUpstreamRef: proto.Clone(m.GetCollectorUpstreamRef()).(*github_com_solo_io_solo_kit_pkg_api_v1_resources_core.ResourceRef),
			}
		}

	case *ZipkinConfig_ClusterName:

		target.CollectorCluster = &ZipkinConfig_ClusterName{
			ClusterName: m.GetClusterName(),
		}

	}

	return target
}

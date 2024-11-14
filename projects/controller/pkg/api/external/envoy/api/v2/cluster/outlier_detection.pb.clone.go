// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/api/v2/cluster/outlier_detection.proto

package cluster

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_durationpb "google.golang.org/protobuf/types/known/durationpb"

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
func (m *OutlierDetection) Clone() proto.Message {
	var target *OutlierDetection
	if m == nil {
		return target
	}
	target = &OutlierDetection{}

	if h, ok := interface{}(m.GetConsecutive_5Xx()).(clone.Cloner); ok {
		target.Consecutive_5Xx = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.Consecutive_5Xx = proto.Clone(m.GetConsecutive_5Xx()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetInterval()).(clone.Cloner); ok {
		target.Interval = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.Interval = proto.Clone(m.GetInterval()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetBaseEjectionTime()).(clone.Cloner); ok {
		target.BaseEjectionTime = h.Clone().(*google_golang_org_protobuf_types_known_durationpb.Duration)
	} else {
		target.BaseEjectionTime = proto.Clone(m.GetBaseEjectionTime()).(*google_golang_org_protobuf_types_known_durationpb.Duration)
	}

	if h, ok := interface{}(m.GetMaxEjectionPercent()).(clone.Cloner); ok {
		target.MaxEjectionPercent = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.MaxEjectionPercent = proto.Clone(m.GetMaxEjectionPercent()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetEnforcingConsecutive_5Xx()).(clone.Cloner); ok {
		target.EnforcingConsecutive_5Xx = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.EnforcingConsecutive_5Xx = proto.Clone(m.GetEnforcingConsecutive_5Xx()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetEnforcingSuccessRate()).(clone.Cloner); ok {
		target.EnforcingSuccessRate = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.EnforcingSuccessRate = proto.Clone(m.GetEnforcingSuccessRate()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetSuccessRateMinimumHosts()).(clone.Cloner); ok {
		target.SuccessRateMinimumHosts = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.SuccessRateMinimumHosts = proto.Clone(m.GetSuccessRateMinimumHosts()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetSuccessRateRequestVolume()).(clone.Cloner); ok {
		target.SuccessRateRequestVolume = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.SuccessRateRequestVolume = proto.Clone(m.GetSuccessRateRequestVolume()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetSuccessRateStdevFactor()).(clone.Cloner); ok {
		target.SuccessRateStdevFactor = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.SuccessRateStdevFactor = proto.Clone(m.GetSuccessRateStdevFactor()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetConsecutiveGatewayFailure()).(clone.Cloner); ok {
		target.ConsecutiveGatewayFailure = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.ConsecutiveGatewayFailure = proto.Clone(m.GetConsecutiveGatewayFailure()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetEnforcingConsecutiveGatewayFailure()).(clone.Cloner); ok {
		target.EnforcingConsecutiveGatewayFailure = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.EnforcingConsecutiveGatewayFailure = proto.Clone(m.GetEnforcingConsecutiveGatewayFailure()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	target.SplitExternalLocalOriginErrors = m.GetSplitExternalLocalOriginErrors()

	if h, ok := interface{}(m.GetConsecutiveLocalOriginFailure()).(clone.Cloner); ok {
		target.ConsecutiveLocalOriginFailure = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.ConsecutiveLocalOriginFailure = proto.Clone(m.GetConsecutiveLocalOriginFailure()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetEnforcingConsecutiveLocalOriginFailure()).(clone.Cloner); ok {
		target.EnforcingConsecutiveLocalOriginFailure = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.EnforcingConsecutiveLocalOriginFailure = proto.Clone(m.GetEnforcingConsecutiveLocalOriginFailure()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetEnforcingLocalOriginSuccessRate()).(clone.Cloner); ok {
		target.EnforcingLocalOriginSuccessRate = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.EnforcingLocalOriginSuccessRate = proto.Clone(m.GetEnforcingLocalOriginSuccessRate()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upstream.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import core_solo_io "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
import core_solo_io1 "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
// @solo-kit:resource.short_name=us
// @solo-kit:resource.plural_name=upstreams
// @solo-kit:resource.resource_groups=api.gloo.solo.io
//
// Upstreams represent destination for routing HTTP requests. Upstreams can be compared to
// [clusters](https://www.envoyproxy.io/docs/envoy/latest/api-v1/cluster_manager/cluster.html?highlight=cluster) in Envoy terminology.
// Each upstream in Gloo has a type. Supported types include `static`, `kubernetes`, `aws`, `consul`, and more.
// Each upstream type is handled by a corresponding Gloo plugin.
type Upstream struct {
	// Type-specific configuration. Examples include static, kubernetes, and aws.
	// The type-specific config for the upstream is called a spec.
	UpstreamSpec *UpstreamSpec `protobuf:"bytes,2,opt,name=upstream_spec,json=upstreamSpec" json:"upstream_spec,omitempty"`
	// Status indicates the validation status of the resource. Status is read-only by clients, and set by gloo during validation
	Status core_solo_io1.Status `protobuf:"bytes,6,opt,name=status" json:"status" testdiff:"ignore"`
	// Metadata contains the object metadata for this resource
	Metadata core_solo_io.Metadata `protobuf:"bytes,7,opt,name=metadata" json:"metadata"`
	// Upstreams and their configuration can be automatically by Gloo Discovery
	// if this upstream is created or modified by Discovery, metadata about the operation will be placed here.
	DiscoveryMetadata *DiscoveryMetadata `protobuf:"bytes,8,opt,name=discovery_metadata,json=discoveryMetadata" json:"discovery_metadata,omitempty"`
}

func (m *Upstream) Reset()                    { *m = Upstream{} }
func (m *Upstream) String() string            { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()               {}
func (*Upstream) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{0} }

func (m *Upstream) GetUpstreamSpec() *UpstreamSpec {
	if m != nil {
		return m.UpstreamSpec
	}
	return nil
}

func (m *Upstream) GetStatus() core_solo_io1.Status {
	if m != nil {
		return m.Status
	}
	return core_solo_io1.Status{}
}

func (m *Upstream) GetMetadata() core_solo_io.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core_solo_io.Metadata{}
}

func (m *Upstream) GetDiscoveryMetadata() *DiscoveryMetadata {
	if m != nil {
		return m.DiscoveryMetadata
	}
	return nil
}

// created by discovery services
type DiscoveryMetadata struct {
}

func (m *DiscoveryMetadata) Reset()                    { *m = DiscoveryMetadata{} }
func (m *DiscoveryMetadata) String() string            { return proto.CompactTextString(m) }
func (*DiscoveryMetadata) ProtoMessage()               {}
func (*DiscoveryMetadata) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{1} }

func init() {
	proto.RegisterType((*Upstream)(nil), "gloo.solo.io.Upstream")
	proto.RegisterType((*DiscoveryMetadata)(nil), "gloo.solo.io.DiscoveryMetadata")
}
func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.UpstreamSpec.Equal(that1.UpstreamSpec) {
		return false
	}
	if !this.Status.Equal(&that1.Status) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !this.DiscoveryMetadata.Equal(that1.DiscoveryMetadata) {
		return false
	}
	return true
}
func (this *DiscoveryMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DiscoveryMetadata)
	if !ok {
		that2, ok := that.(DiscoveryMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}

func init() { proto.RegisterFile("upstream.proto", fileDescriptorUpstream) }

var fileDescriptorUpstream = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4a, 0x03, 0x31,
	0x14, 0xc7, 0x6d, 0x91, 0x5a, 0x62, 0x2b, 0x34, 0x16, 0xa9, 0x5d, 0x58, 0x99, 0x95, 0x1b, 0x13,
	0xaa, 0x2e, 0xa4, 0x1b, 0xa1, 0x08, 0xae, 0x74, 0x31, 0xc5, 0x8d, 0x9b, 0x92, 0x66, 0xd2, 0x18,
	0xfb, 0xf1, 0x42, 0x92, 0x29, 0x78, 0x19, 0xd7, 0x1e, 0xc5, 0x53, 0x74, 0xe1, 0x11, 0x3c, 0x81,
	0x4c, 0x26, 0x33, 0xb4, 0x0a, 0xe2, 0x2a, 0x79, 0x79, 0xff, 0xff, 0xef, 0x7d, 0x04, 0x1d, 0xa4,
	0xda, 0x3a, 0x23, 0xd8, 0x82, 0x68, 0x03, 0x0e, 0x70, 0x43, 0xce, 0x01, 0x88, 0x85, 0x39, 0x10,
	0x05, 0xdd, 0xb6, 0x04, 0x09, 0x3e, 0x41, 0xb3, 0x5b, 0xae, 0xe9, 0xf6, 0xa5, 0x72, 0xcf, 0xe9,
	0x84, 0x70, 0x58, 0xd0, 0x4c, 0x79, 0xae, 0x20, 0x3f, 0x67, 0xca, 0x51, 0xa6, 0x15, 0x5d, 0xf5,
	0xe9, 0x42, 0x38, 0x96, 0x30, 0xc7, 0x82, 0x85, 0xfe, 0xc3, 0x62, 0x1d, 0x73, 0xa9, 0x0d, 0x86,
	0xa6, 0x9e, 0xa7, 0x52, 0x2d, 0x43, 0x18, 0xbd, 0x55, 0x51, 0xfd, 0x31, 0x74, 0x8a, 0x6f, 0x50,
	0xb3, 0xe8, 0x7a, 0x6c, 0xb5, 0xe0, 0x9d, 0xea, 0x69, 0xe5, 0x6c, 0xff, 0xa2, 0x4b, 0x36, 0x7b,
	0x27, 0x85, 0x7c, 0xa4, 0x05, 0x8f, 0x1b, 0xe9, 0x46, 0x84, 0xef, 0x50, 0x2d, 0x2f, 0xd6, 0xa9,
	0x79, 0x67, 0x9b, 0x70, 0x30, 0xa2, 0x74, 0x8e, 0x7c, 0x6e, 0x78, 0xfc, 0xb1, 0xee, 0xed, 0x7c,
	0xad, 0x7b, 0x2d, 0x27, 0xac, 0x4b, 0xd4, 0x74, 0x3a, 0x88, 0x94, 0x5c, 0x82, 0x11, 0x51, 0x1c,
	0xec, 0xf8, 0x1a, 0xd5, 0x8b, 0x41, 0x3b, 0x7b, 0x1e, 0x75, 0xb4, 0x8d, 0xba, 0x0f, 0xd9, 0xe1,
	0x6e, 0x06, 0x8b, 0x4b, 0x35, 0x7e, 0x40, 0x38, 0x51, 0x96, 0xc3, 0x4a, 0x98, 0xd7, 0x71, 0xc9,
	0xa8, 0x7b, 0x46, 0x6f, 0x7b, 0x90, 0xdb, 0x42, 0x57, 0xc0, 0xe2, 0x56, 0xf2, 0xf3, 0x29, 0x3a,
	0x44, 0xad, 0x5f, 0xba, 0xe1, 0xe0, 0xfd, 0xf3, 0xa4, 0xf2, 0x74, 0xf5, 0xd7, 0xee, 0xb5, 0x81,
	0x17, 0xc1, 0x9d, 0xa5, 0x59, 0x49, 0xaa, 0x67, 0x32, 0xfc, 0xc6, 0xa4, 0xe6, 0x17, 0x7f, 0xf9,
	0x1d, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x27, 0xfb, 0x7a, 0x21, 0x02, 0x00, 0x00,
}

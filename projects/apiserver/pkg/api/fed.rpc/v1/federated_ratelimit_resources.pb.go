// Code generated by skv2. DO NOT EDIT.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-projects/projects/apiserver/api/fed.rpc/v1/federated_ratelimit_resources.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	v11 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1"
	types "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.ratelimit.solo.io/v1alpha1/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FederatedRateLimitConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *v1.ObjectMeta                        `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Spec     *types.FederatedRateLimitConfigSpec   `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Status   *types.FederatedRateLimitConfigStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *FederatedRateLimitConfig) Reset() {
	*x = FederatedRateLimitConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FederatedRateLimitConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FederatedRateLimitConfig) ProtoMessage() {}

func (x *FederatedRateLimitConfig) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FederatedRateLimitConfig.ProtoReflect.Descriptor instead.
func (*FederatedRateLimitConfig) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescGZIP(), []int{0}
}

func (x *FederatedRateLimitConfig) GetMetadata() *v1.ObjectMeta {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *FederatedRateLimitConfig) GetSpec() *types.FederatedRateLimitConfigSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *FederatedRateLimitConfig) GetStatus() *types.FederatedRateLimitConfigStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

type ListFederatedRateLimitConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListFederatedRateLimitConfigsRequest) Reset() {
	*x = ListFederatedRateLimitConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederatedRateLimitConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederatedRateLimitConfigsRequest) ProtoMessage() {}

func (x *ListFederatedRateLimitConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederatedRateLimitConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListFederatedRateLimitConfigsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescGZIP(), []int{1}
}

type ListFederatedRateLimitConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FederatedRateLimitConfigs []*FederatedRateLimitConfig `protobuf:"bytes,1,rep,name=federated_rate_limit_configs,json=federatedRateLimitConfigs,proto3" json:"federated_rate_limit_configs,omitempty"`
}

func (x *ListFederatedRateLimitConfigsResponse) Reset() {
	*x = ListFederatedRateLimitConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederatedRateLimitConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederatedRateLimitConfigsResponse) ProtoMessage() {}

func (x *ListFederatedRateLimitConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederatedRateLimitConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListFederatedRateLimitConfigsResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescGZIP(), []int{2}
}

func (x *ListFederatedRateLimitConfigsResponse) GetFederatedRateLimitConfigs() []*FederatedRateLimitConfig {
	if x != nil {
		return x.FederatedRateLimitConfigs
	}
	return nil
}

type GetFederatedRateLimitConfigYamlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FederatedRateLimitConfigRef *v11.ObjectRef `protobuf:"bytes,1,opt,name=federated_rate_limit_config_ref,json=federatedRateLimitConfigRef,proto3" json:"federated_rate_limit_config_ref,omitempty"`
}

func (x *GetFederatedRateLimitConfigYamlRequest) Reset() {
	*x = GetFederatedRateLimitConfigYamlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFederatedRateLimitConfigYamlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFederatedRateLimitConfigYamlRequest) ProtoMessage() {}

func (x *GetFederatedRateLimitConfigYamlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFederatedRateLimitConfigYamlRequest.ProtoReflect.Descriptor instead.
func (*GetFederatedRateLimitConfigYamlRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescGZIP(), []int{3}
}

func (x *GetFederatedRateLimitConfigYamlRequest) GetFederatedRateLimitConfigRef() *v11.ObjectRef {
	if x != nil {
		return x.FederatedRateLimitConfigRef
	}
	return nil
}

type GetFederatedRateLimitConfigYamlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YamlData *v1.ResourceYaml `protobuf:"bytes,1,opt,name=yaml_data,json=yamlData,proto3" json:"yaml_data,omitempty"`
}

func (x *GetFederatedRateLimitConfigYamlResponse) Reset() {
	*x = GetFederatedRateLimitConfigYamlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFederatedRateLimitConfigYamlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFederatedRateLimitConfigYamlResponse) ProtoMessage() {}

func (x *GetFederatedRateLimitConfigYamlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFederatedRateLimitConfigYamlResponse.ProtoReflect.Descriptor instead.
func (*GetFederatedRateLimitConfigYamlResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescGZIP(), []int{4}
}

func (x *GetFederatedRateLimitConfigYamlResponse) GetYamlData() *v1.ResourceYaml {
	if x != nil {
		return x.YamlData
	}
	return nil
}

var File_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDesc = []byte{
	0x0a, 0x66, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x65, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x66, 0x65,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x72, 0x61, 0x74,
	0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6b, 0x76, 0x32, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70,
	0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x01, 0x0a, 0x18,
	0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x61, 0x74, 0x65,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73, 0x70, 0x65, 0x63,
	0x12, 0x4d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x26, 0x0a, 0x24, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x93, 0x01, 0x0a, 0x25, 0x4c, 0x69, 0x73, 0x74,
	0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x6a, 0x0a, 0x1c, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72,
	0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x19, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x8c, 0x01,
	0x0a, 0x26, 0x47, 0x65, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61,
	0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x62, 0x0a, 0x1f, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x73, 0x6b, 0x76, 0x32, 0x2e, 0x73, 0x6f,
	0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52,
	0x1b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x66, 0x22, 0x6b, 0x0a, 0x27,
	0x47, 0x65, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x79, 0x61, 0x6d, 0x6c, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x59, 0x61, 0x6d, 0x6c, 0x52,
	0x08, 0x79, 0x61, 0x6d, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x32, 0xcb, 0x02, 0x0a, 0x1d, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x70, 0x69, 0x12, 0x90, 0x01, 0x0a, 0x1d,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x35, 0x2e,
	0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74,
	0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x96,
	0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61,
	0x6d, 0x6c, 0x12, 0x37, 0x2e, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x66, 0x65,
	0x64, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x61, 0x74, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x59, 0x61, 0x6d, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x50, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x65, 0x64, 0x2e, 0x72, 0x70, 0x63, 0x2f, 0x76,
	0x31, 0xc0, 0xf5, 0x04, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescData = file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDesc
)

func file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDescData
}

var file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_goTypes = []interface{}{
	(*FederatedRateLimitConfig)(nil),                // 0: fed.rpc.solo.io.FederatedRateLimitConfig
	(*ListFederatedRateLimitConfigsRequest)(nil),    // 1: fed.rpc.solo.io.ListFederatedRateLimitConfigsRequest
	(*ListFederatedRateLimitConfigsResponse)(nil),   // 2: fed.rpc.solo.io.ListFederatedRateLimitConfigsResponse
	(*GetFederatedRateLimitConfigYamlRequest)(nil),  // 3: fed.rpc.solo.io.GetFederatedRateLimitConfigYamlRequest
	(*GetFederatedRateLimitConfigYamlResponse)(nil), // 4: fed.rpc.solo.io.GetFederatedRateLimitConfigYamlResponse
	(*v1.ObjectMeta)(nil),                           // 5: rpc.edge.gloo.solo.io.ObjectMeta
	(*types.FederatedRateLimitConfigSpec)(nil),      // 6: fed.ratelimit.solo.io.FederatedRateLimitConfigSpec
	(*types.FederatedRateLimitConfigStatus)(nil),    // 7: fed.ratelimit.solo.io.FederatedRateLimitConfigStatus
	(*v11.ObjectRef)(nil),                           // 8: core.skv2.solo.io.ObjectRef
	(*v1.ResourceYaml)(nil),                         // 9: rpc.edge.gloo.solo.io.ResourceYaml
}
var file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_depIdxs = []int32{
	5, // 0: fed.rpc.solo.io.FederatedRateLimitConfig.metadata:type_name -> rpc.edge.gloo.solo.io.ObjectMeta
	6, // 1: fed.rpc.solo.io.FederatedRateLimitConfig.spec:type_name -> fed.ratelimit.solo.io.FederatedRateLimitConfigSpec
	7, // 2: fed.rpc.solo.io.FederatedRateLimitConfig.status:type_name -> fed.ratelimit.solo.io.FederatedRateLimitConfigStatus
	0, // 3: fed.rpc.solo.io.ListFederatedRateLimitConfigsResponse.federated_rate_limit_configs:type_name -> fed.rpc.solo.io.FederatedRateLimitConfig
	8, // 4: fed.rpc.solo.io.GetFederatedRateLimitConfigYamlRequest.federated_rate_limit_config_ref:type_name -> core.skv2.solo.io.ObjectRef
	9, // 5: fed.rpc.solo.io.GetFederatedRateLimitConfigYamlResponse.yaml_data:type_name -> rpc.edge.gloo.solo.io.ResourceYaml
	1, // 6: fed.rpc.solo.io.FederatedRatelimitResourceApi.ListFederatedRateLimitConfigs:input_type -> fed.rpc.solo.io.ListFederatedRateLimitConfigsRequest
	3, // 7: fed.rpc.solo.io.FederatedRatelimitResourceApi.GetFederatedRateLimitConfigYaml:input_type -> fed.rpc.solo.io.GetFederatedRateLimitConfigYamlRequest
	2, // 8: fed.rpc.solo.io.FederatedRatelimitResourceApi.ListFederatedRateLimitConfigs:output_type -> fed.rpc.solo.io.ListFederatedRateLimitConfigsResponse
	4, // 9: fed.rpc.solo.io.FederatedRatelimitResourceApi.GetFederatedRateLimitConfigYaml:output_type -> fed.rpc.solo.io.GetFederatedRateLimitConfigYamlResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_init()
}
func file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_init() {
	if File_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FederatedRateLimitConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFederatedRateLimitConfigsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFederatedRateLimitConfigsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFederatedRateLimitConfigYamlRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFederatedRateLimitConfigYamlResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto = out.File
	file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_rawDesc = nil
	file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_goTypes = nil
	file_github_com_solo_io_solo_projects_projects_apiserver_api_fed_rpc_v1_federated_ratelimit_resources_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FederatedRatelimitResourceApiClient is the client API for FederatedRatelimitResourceApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FederatedRatelimitResourceApiClient interface {
	ListFederatedRateLimitConfigs(ctx context.Context, in *ListFederatedRateLimitConfigsRequest, opts ...grpc.CallOption) (*ListFederatedRateLimitConfigsResponse, error)
	GetFederatedRateLimitConfigYaml(ctx context.Context, in *GetFederatedRateLimitConfigYamlRequest, opts ...grpc.CallOption) (*GetFederatedRateLimitConfigYamlResponse, error)
}

type federatedRatelimitResourceApiClient struct {
	cc grpc.ClientConnInterface
}

func NewFederatedRatelimitResourceApiClient(cc grpc.ClientConnInterface) FederatedRatelimitResourceApiClient {
	return &federatedRatelimitResourceApiClient{cc}
}

func (c *federatedRatelimitResourceApiClient) ListFederatedRateLimitConfigs(ctx context.Context, in *ListFederatedRateLimitConfigsRequest, opts ...grpc.CallOption) (*ListFederatedRateLimitConfigsResponse, error) {
	out := new(ListFederatedRateLimitConfigsResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.FederatedRatelimitResourceApi/ListFederatedRateLimitConfigs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedRatelimitResourceApiClient) GetFederatedRateLimitConfigYaml(ctx context.Context, in *GetFederatedRateLimitConfigYamlRequest, opts ...grpc.CallOption) (*GetFederatedRateLimitConfigYamlResponse, error) {
	out := new(GetFederatedRateLimitConfigYamlResponse)
	err := c.cc.Invoke(ctx, "/fed.rpc.solo.io.FederatedRatelimitResourceApi/GetFederatedRateLimitConfigYaml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederatedRatelimitResourceApiServer is the server API for FederatedRatelimitResourceApi service.
type FederatedRatelimitResourceApiServer interface {
	ListFederatedRateLimitConfigs(context.Context, *ListFederatedRateLimitConfigsRequest) (*ListFederatedRateLimitConfigsResponse, error)
	GetFederatedRateLimitConfigYaml(context.Context, *GetFederatedRateLimitConfigYamlRequest) (*GetFederatedRateLimitConfigYamlResponse, error)
}

// UnimplementedFederatedRatelimitResourceApiServer can be embedded to have forward compatible implementations.
type UnimplementedFederatedRatelimitResourceApiServer struct {
}

func (*UnimplementedFederatedRatelimitResourceApiServer) ListFederatedRateLimitConfigs(context.Context, *ListFederatedRateLimitConfigsRequest) (*ListFederatedRateLimitConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFederatedRateLimitConfigs not implemented")
}
func (*UnimplementedFederatedRatelimitResourceApiServer) GetFederatedRateLimitConfigYaml(context.Context, *GetFederatedRateLimitConfigYamlRequest) (*GetFederatedRateLimitConfigYamlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFederatedRateLimitConfigYaml not implemented")
}

func RegisterFederatedRatelimitResourceApiServer(s *grpc.Server, srv FederatedRatelimitResourceApiServer) {
	s.RegisterService(&_FederatedRatelimitResourceApi_serviceDesc, srv)
}

func _FederatedRatelimitResourceApi_ListFederatedRateLimitConfigs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederatedRateLimitConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedRatelimitResourceApiServer).ListFederatedRateLimitConfigs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.FederatedRatelimitResourceApi/ListFederatedRateLimitConfigs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedRatelimitResourceApiServer).ListFederatedRateLimitConfigs(ctx, req.(*ListFederatedRateLimitConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedRatelimitResourceApi_GetFederatedRateLimitConfigYaml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederatedRateLimitConfigYamlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedRatelimitResourceApiServer).GetFederatedRateLimitConfigYaml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fed.rpc.solo.io.FederatedRatelimitResourceApi/GetFederatedRateLimitConfigYaml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedRatelimitResourceApiServer).GetFederatedRateLimitConfigYaml(ctx, req.(*GetFederatedRateLimitConfigYamlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FederatedRatelimitResourceApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fed.rpc.solo.io.FederatedRatelimitResourceApi",
	HandlerType: (*FederatedRatelimitResourceApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListFederatedRateLimitConfigs",
			Handler:    _FederatedRatelimitResourceApi_ListFederatedRateLimitConfigs_Handler,
		},
		{
			MethodName: "GetFederatedRateLimitConfigYaml",
			Handler:    _FederatedRatelimitResourceApi_GetFederatedRateLimitConfigYaml_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-projects/projects/apiserver/api/fed.rpc/v1/federated_ratelimit_resources.proto",
}

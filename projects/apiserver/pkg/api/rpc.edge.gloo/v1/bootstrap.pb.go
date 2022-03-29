// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/bootstrap.proto

package v1

import (
	context "context"
	reflect "reflect"
	sync "sync"

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

type GlooFedCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GlooFedCheckRequest) Reset() {
	*x = GlooFedCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlooFedCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlooFedCheckRequest) ProtoMessage() {}

func (x *GlooFedCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlooFedCheckRequest.ProtoReflect.Descriptor instead.
func (*GlooFedCheckRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescGZIP(), []int{0}
}

type GlooFedCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// returns true if Gloo Fed is enabled in the current Gloo Edge installation
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *GlooFedCheckResponse) Reset() {
	*x = GlooFedCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GlooFedCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GlooFedCheckResponse) ProtoMessage() {}

func (x *GlooFedCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GlooFedCheckResponse.ProtoReflect.Descriptor instead.
func (*GlooFedCheckResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescGZIP(), []int{1}
}

func (x *GlooFedCheckResponse) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

type GraphqlCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GraphqlCheckRequest) Reset() {
	*x = GraphqlCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphqlCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphqlCheckRequest) ProtoMessage() {}

func (x *GraphqlCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphqlCheckRequest.ProtoReflect.Descriptor instead.
func (*GraphqlCheckRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescGZIP(), []int{2}
}

type GraphqlCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// returns true if GraphQL is enabled in the provided license
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *GraphqlCheckResponse) Reset() {
	*x = GraphqlCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GraphqlCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GraphqlCheckResponse) ProtoMessage() {}

func (x *GraphqlCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GraphqlCheckResponse.ProtoReflect.Descriptor instead.
func (*GraphqlCheckResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescGZIP(), []int{3}
}

func (x *GraphqlCheckResponse) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

// Options for the current Gloo Edge UI.
type ConsoleOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Returns true if the UI is read-only, false if UI allows creates/updates/deletes.
	ReadOnly bool `protobuf:"varint,1,opt,name=read_only,json=readOnly,proto3" json:"read_only,omitempty"`
	// Returns true if the GraphQL explorer is enabled.
	ApiExplorerEnabled bool `protobuf:"varint,2,opt,name=api_explorer_enabled,json=apiExplorerEnabled,proto3" json:"api_explorer_enabled,omitempty"`
}

func (x *ConsoleOptions) Reset() {
	*x = ConsoleOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsoleOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsoleOptions) ProtoMessage() {}

func (x *ConsoleOptions) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsoleOptions.ProtoReflect.Descriptor instead.
func (*ConsoleOptions) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescGZIP(), []int{4}
}

func (x *ConsoleOptions) GetReadOnly() bool {
	if x != nil {
		return x.ReadOnly
	}
	return false
}

func (x *ConsoleOptions) GetApiExplorerEnabled() bool {
	if x != nil {
		return x.ApiExplorerEnabled
	}
	return false
}

type GetConsoleOptionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConsoleOptionsRequest) Reset() {
	*x = GetConsoleOptionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsoleOptionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsoleOptionsRequest) ProtoMessage() {}

func (x *GetConsoleOptionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsoleOptionsRequest.ProtoReflect.Descriptor instead.
func (*GetConsoleOptionsRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescGZIP(), []int{5}
}

type GetConsoleOptionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *ConsoleOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *GetConsoleOptionsResponse) Reset() {
	*x = GetConsoleOptionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConsoleOptionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConsoleOptionsResponse) ProtoMessage() {}

func (x *GetConsoleOptionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConsoleOptionsResponse.ProtoReflect.Descriptor instead.
func (*GetConsoleOptionsResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescGZIP(), []int{6}
}

func (x *GetConsoleOptionsResponse) GetOptions() *ConsoleOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

var File_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto protoreflect.FileDescriptor

var file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDesc = []byte{
	0x0a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64,
	0x67, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x74, 0x73,
	0x74, 0x72, 0x61, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x72, 0x70, 0x63, 0x2e,
	0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x6c, 0x6f, 0x6f, 0x46, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x14, 0x47, 0x6c, 0x6f, 0x6f,
	0x46, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x72,
	0x61, 0x70, 0x68, 0x71, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x30, 0x0a, 0x14, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x22, 0x5f, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6f, 0x6e,
	0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x61, 0x64, 0x4f, 0x6e,
	0x6c, 0x79, 0x12, 0x30, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x5f, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72,
	0x65, 0x72, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x61, 0x70, 0x69, 0x45, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f,
	0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x5c, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0xe6,
	0x02, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x41, 0x70, 0x69, 0x12,
	0x6d, 0x0a, 0x10, 0x49, 0x73, 0x47, 0x6c, 0x6f, 0x6f, 0x46, 0x65, 0x64, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x2a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x6c, 0x6f, 0x6f,
	0x46, 0x65, 0x64, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2b, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x6c, 0x6f, 0x6f, 0x46, 0x65, 0x64, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d,
	0x0a, 0x10, 0x49, 0x73, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x2a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68,
	0x71, 0x6c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x72, 0x61, 0x70, 0x68, 0x71, 0x6c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x78, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x2f, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x4e, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f,
	0x6c, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x64, 0x67, 0x65, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescOnce sync.Once
	file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescData = file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDesc
)

func file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescData)
	})
	return file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDescData
}

var file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_goTypes = []interface{}{
	(*GlooFedCheckRequest)(nil),       // 0: rpc.edge.gloo.solo.io.GlooFedCheckRequest
	(*GlooFedCheckResponse)(nil),      // 1: rpc.edge.gloo.solo.io.GlooFedCheckResponse
	(*GraphqlCheckRequest)(nil),       // 2: rpc.edge.gloo.solo.io.GraphqlCheckRequest
	(*GraphqlCheckResponse)(nil),      // 3: rpc.edge.gloo.solo.io.GraphqlCheckResponse
	(*ConsoleOptions)(nil),            // 4: rpc.edge.gloo.solo.io.ConsoleOptions
	(*GetConsoleOptionsRequest)(nil),  // 5: rpc.edge.gloo.solo.io.GetConsoleOptionsRequest
	(*GetConsoleOptionsResponse)(nil), // 6: rpc.edge.gloo.solo.io.GetConsoleOptionsResponse
}
var file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_depIdxs = []int32{
	4, // 0: rpc.edge.gloo.solo.io.GetConsoleOptionsResponse.options:type_name -> rpc.edge.gloo.solo.io.ConsoleOptions
	0, // 1: rpc.edge.gloo.solo.io.BootstrapApi.IsGlooFedEnabled:input_type -> rpc.edge.gloo.solo.io.GlooFedCheckRequest
	2, // 2: rpc.edge.gloo.solo.io.BootstrapApi.IsGraphqlEnabled:input_type -> rpc.edge.gloo.solo.io.GraphqlCheckRequest
	5, // 3: rpc.edge.gloo.solo.io.BootstrapApi.GetConsoleOptions:input_type -> rpc.edge.gloo.solo.io.GetConsoleOptionsRequest
	1, // 4: rpc.edge.gloo.solo.io.BootstrapApi.IsGlooFedEnabled:output_type -> rpc.edge.gloo.solo.io.GlooFedCheckResponse
	3, // 5: rpc.edge.gloo.solo.io.BootstrapApi.IsGraphqlEnabled:output_type -> rpc.edge.gloo.solo.io.GraphqlCheckResponse
	6, // 6: rpc.edge.gloo.solo.io.BootstrapApi.GetConsoleOptions:output_type -> rpc.edge.gloo.solo.io.GetConsoleOptionsResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_init()
}
func file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_init() {
	if File_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlooFedCheckRequest); i {
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
		file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GlooFedCheckResponse); i {
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
		file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphqlCheckRequest); i {
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
		file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GraphqlCheckResponse); i {
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
		file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsoleOptions); i {
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
		file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsoleOptionsRequest); i {
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
		file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConsoleOptionsResponse); i {
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
			RawDescriptor: file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto = out.File
	file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_rawDesc = nil
	file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_goTypes = nil
	file_github_com_solo_io_solo_projects_projects_apiserver_api_rpc_edge_gloo_v1_bootstrap_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BootstrapApiClient is the client API for BootstrapApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BootstrapApiClient interface {
	IsGlooFedEnabled(ctx context.Context, in *GlooFedCheckRequest, opts ...grpc.CallOption) (*GlooFedCheckResponse, error)
	IsGraphqlEnabled(ctx context.Context, in *GraphqlCheckRequest, opts ...grpc.CallOption) (*GraphqlCheckResponse, error)
	GetConsoleOptions(ctx context.Context, in *GetConsoleOptionsRequest, opts ...grpc.CallOption) (*GetConsoleOptionsResponse, error)
}

type bootstrapApiClient struct {
	cc grpc.ClientConnInterface
}

func NewBootstrapApiClient(cc grpc.ClientConnInterface) BootstrapApiClient {
	return &bootstrapApiClient{cc}
}

func (c *bootstrapApiClient) IsGlooFedEnabled(ctx context.Context, in *GlooFedCheckRequest, opts ...grpc.CallOption) (*GlooFedCheckResponse, error) {
	out := new(GlooFedCheckResponse)
	err := c.cc.Invoke(ctx, "/rpc.edge.gloo.solo.io.BootstrapApi/IsGlooFedEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bootstrapApiClient) IsGraphqlEnabled(ctx context.Context, in *GraphqlCheckRequest, opts ...grpc.CallOption) (*GraphqlCheckResponse, error) {
	out := new(GraphqlCheckResponse)
	err := c.cc.Invoke(ctx, "/rpc.edge.gloo.solo.io.BootstrapApi/IsGraphqlEnabled", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bootstrapApiClient) GetConsoleOptions(ctx context.Context, in *GetConsoleOptionsRequest, opts ...grpc.CallOption) (*GetConsoleOptionsResponse, error) {
	out := new(GetConsoleOptionsResponse)
	err := c.cc.Invoke(ctx, "/rpc.edge.gloo.solo.io.BootstrapApi/GetConsoleOptions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BootstrapApiServer is the server API for BootstrapApi service.
type BootstrapApiServer interface {
	IsGlooFedEnabled(context.Context, *GlooFedCheckRequest) (*GlooFedCheckResponse, error)
	IsGraphqlEnabled(context.Context, *GraphqlCheckRequest) (*GraphqlCheckResponse, error)
	GetConsoleOptions(context.Context, *GetConsoleOptionsRequest) (*GetConsoleOptionsResponse, error)
}

// UnimplementedBootstrapApiServer can be embedded to have forward compatible implementations.
type UnimplementedBootstrapApiServer struct {
}

func (*UnimplementedBootstrapApiServer) IsGlooFedEnabled(context.Context, *GlooFedCheckRequest) (*GlooFedCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsGlooFedEnabled not implemented")
}
func (*UnimplementedBootstrapApiServer) IsGraphqlEnabled(context.Context, *GraphqlCheckRequest) (*GraphqlCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsGraphqlEnabled not implemented")
}
func (*UnimplementedBootstrapApiServer) GetConsoleOptions(context.Context, *GetConsoleOptionsRequest) (*GetConsoleOptionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsoleOptions not implemented")
}

func RegisterBootstrapApiServer(s *grpc.Server, srv BootstrapApiServer) {
	s.RegisterService(&_BootstrapApi_serviceDesc, srv)
}

func _BootstrapApi_IsGlooFedEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GlooFedCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapApiServer).IsGlooFedEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.edge.gloo.solo.io.BootstrapApi/IsGlooFedEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapApiServer).IsGlooFedEnabled(ctx, req.(*GlooFedCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BootstrapApi_IsGraphqlEnabled_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GraphqlCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapApiServer).IsGraphqlEnabled(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.edge.gloo.solo.io.BootstrapApi/IsGraphqlEnabled",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapApiServer).IsGraphqlEnabled(ctx, req.(*GraphqlCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BootstrapApi_GetConsoleOptions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsoleOptionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BootstrapApiServer).GetConsoleOptions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.edge.gloo.solo.io.BootstrapApi/GetConsoleOptions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BootstrapApiServer).GetConsoleOptions(ctx, req.(*GetConsoleOptionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BootstrapApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.edge.gloo.solo.io.BootstrapApi",
	HandlerType: (*BootstrapApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsGlooFedEnabled",
			Handler:    _BootstrapApi_IsGlooFedEnabled_Handler,
		},
		{
			MethodName: "IsGraphqlEnabled",
			Handler:    _BootstrapApi_IsGraphqlEnabled_Handler,
		},
		{
			MethodName: "GetConsoleOptions",
			Handler:    _BootstrapApi_GetConsoleOptions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/bootstrap.proto",
}

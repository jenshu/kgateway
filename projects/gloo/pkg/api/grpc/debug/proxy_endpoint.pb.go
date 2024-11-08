// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/grpc/debug/proxy_endpoint.proto

package debug

import (
	context "context"
	reflect "reflect"
	sync "sync"

	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
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

type ProxyEndpointRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. The namespace to look for proxies. If this is omitted, all namespaces will be considered.
	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// Optional. The name of the proxy to look up. If this is provided, a namespace must be included as well.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Equality-based selector to use to filter returned proxies. This will be ignored if a name is provided.
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#equality-based-requirement
	// If both `selector` and `expressionSelector` are defined, then `expressionSelector` is used.
	Selector map[string]string `protobuf:"bytes,3,rep,name=selector,proto3" json:"selector,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Optional. Set-based selector to use to filter returned proxies. This will be ignored if a name is provided.
	// See https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/#set-based-requirement
	// If both `selector` and `expressionSelector` are defined, then `expressionSelector` is used.
	ExpressionSelector string `protobuf:"bytes,5,opt,name=expression_selector,json=expressionSelector,proto3" json:"expression_selector,omitempty"`
}

func (x *ProxyEndpointRequest) Reset() {
	*x = ProxyEndpointRequest{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProxyEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyEndpointRequest) ProtoMessage() {}

func (x *ProxyEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyEndpointRequest.ProtoReflect.Descriptor instead.
func (*ProxyEndpointRequest) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescGZIP(), []int{0}
}

func (x *ProxyEndpointRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ProxyEndpointRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProxyEndpointRequest) GetSelector() map[string]string {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *ProxyEndpointRequest) GetExpressionSelector() string {
	if x != nil {
		return x.ExpressionSelector
	}
	return ""
}

type ProxyEndpointResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of proxies returned from the gloo pod.
	Proxies []*v1.Proxy `protobuf:"bytes,1,rep,name=proxies,proto3" json:"proxies,omitempty"`
}

func (x *ProxyEndpointResponse) Reset() {
	*x = ProxyEndpointResponse{}
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProxyEndpointResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProxyEndpointResponse) ProtoMessage() {}

func (x *ProxyEndpointResponse) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProxyEndpointResponse.ProtoReflect.Descriptor instead.
func (*ProxyEndpointResponse) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescGZIP(), []int{1}
}

func (x *ProxyEndpointResponse) GetProxies() []*v1.Proxy {
	if x != nil {
		return x.Proxies
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x5f, 0x65, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69,
	0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x73, 0x6f, 0x6c, 0x6f,
	0x2d, 0x6b, 0x69, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8a, 0x02, 0x0a, 0x14, 0x50,
	0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x4c, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x64, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x65, 0x78, 0x70, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6c, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x1a, 0x3b, 0x0a, 0x0d, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x22, 0x46, 0x0a, 0x15, 0x50, 0x72, 0x6f, 0x78, 0x79,
	0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2d, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x69, 0x65, 0x73, 0x32,
	0x6f, 0x0a, 0x14, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x78, 0x69, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c,
	0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x45, 0x6e,
	0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73,
	0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x64, 0x65, 0x62, 0x75, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_goTypes = []any{
	(*ProxyEndpointRequest)(nil),  // 0: gloo.solo.io.ProxyEndpointRequest
	(*ProxyEndpointResponse)(nil), // 1: gloo.solo.io.ProxyEndpointResponse
	nil,                           // 2: gloo.solo.io.ProxyEndpointRequest.SelectorEntry
	(*v1.Proxy)(nil),              // 3: gloo.solo.io.Proxy
}
var file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_depIdxs = []int32{
	2, // 0: gloo.solo.io.ProxyEndpointRequest.selector:type_name -> gloo.solo.io.ProxyEndpointRequest.SelectorEntry
	3, // 1: gloo.solo.io.ProxyEndpointResponse.proxies:type_name -> gloo.solo.io.Proxy
	0, // 2: gloo.solo.io.ProxyEndpointService.GetProxies:input_type -> gloo.solo.io.ProxyEndpointRequest
	1, // 3: gloo.solo.io.ProxyEndpointService.GetProxies:output_type -> gloo.solo.io.ProxyEndpointResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_init() }
func file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_grpc_debug_proxy_endpoint_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProxyEndpointServiceClient is the client API for ProxyEndpointService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProxyEndpointServiceClient interface {
	// Get a filtered list of proxies.
	GetProxies(ctx context.Context, in *ProxyEndpointRequest, opts ...grpc.CallOption) (*ProxyEndpointResponse, error)
}

type proxyEndpointServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProxyEndpointServiceClient(cc grpc.ClientConnInterface) ProxyEndpointServiceClient {
	return &proxyEndpointServiceClient{cc}
}

func (c *proxyEndpointServiceClient) GetProxies(ctx context.Context, in *ProxyEndpointRequest, opts ...grpc.CallOption) (*ProxyEndpointResponse, error) {
	out := new(ProxyEndpointResponse)
	err := c.cc.Invoke(ctx, "/gloo.solo.io.ProxyEndpointService/GetProxies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProxyEndpointServiceServer is the server API for ProxyEndpointService service.
type ProxyEndpointServiceServer interface {
	// Get a filtered list of proxies.
	GetProxies(context.Context, *ProxyEndpointRequest) (*ProxyEndpointResponse, error)
}

// UnimplementedProxyEndpointServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProxyEndpointServiceServer struct {
}

func (*UnimplementedProxyEndpointServiceServer) GetProxies(context.Context, *ProxyEndpointRequest) (*ProxyEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProxies not implemented")
}

func RegisterProxyEndpointServiceServer(s *grpc.Server, srv ProxyEndpointServiceServer) {
	s.RegisterService(&_ProxyEndpointService_serviceDesc, srv)
}

func _ProxyEndpointService_GetProxies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProxyEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProxyEndpointServiceServer).GetProxies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gloo.solo.io.ProxyEndpointService/GetProxies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProxyEndpointServiceServer).GetProxies(ctx, req.(*ProxyEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProxyEndpointService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gloo.solo.io.ProxyEndpointService",
	HandlerType: (*ProxyEndpointServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProxies",
			Handler:    _ProxyEndpointService_GetProxies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/gloo/projects/gloo/api/grpc/debug/proxy_endpoint.proto",
}

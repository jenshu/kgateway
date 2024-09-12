// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/extensions/extauth/sanitize.proto

package extauth

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sanitize struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HeadersToRemove []string `protobuf:"bytes,1,rep,name=headers_to_remove,json=headersToRemove,proto3" json:"headers_to_remove,omitempty"`
	// The name of the custom ext auth server
	// This is only set in the case where multiple auth servers exist.
	CustomAuthServerName string `protobuf:"bytes,2,opt,name=custom_auth_server_name,json=customAuthServerName,proto3" json:"custom_auth_server_name,omitempty"`
}

func (x *Sanitize) Reset() {
	*x = Sanitize{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sanitize) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sanitize) ProtoMessage() {}

func (x *Sanitize) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sanitize.ProtoReflect.Descriptor instead.
func (*Sanitize) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescGZIP(), []int{0}
}

func (x *Sanitize) GetHeadersToRemove() []string {
	if x != nil {
		return x.HeadersToRemove
	}
	return nil
}

func (x *Sanitize) GetCustomAuthServerName() string {
	if x != nil {
		return x.CustomAuthServerName
	}
	return ""
}

type SanitizePerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the custom ext auth server
	// This is only set in the case where multiple auth servers exist.
	CustomAuthServerName string `protobuf:"bytes,1,opt,name=custom_auth_server_name,json=customAuthServerName,proto3" json:"custom_auth_server_name,omitempty"`
}

func (x *SanitizePerRoute) Reset() {
	*x = SanitizePerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SanitizePerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SanitizePerRoute) ProtoMessage() {}

func (x *SanitizePerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SanitizePerRoute.ProtoReflect.Descriptor instead.
func (*SanitizePerRoute) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescGZIP(), []int{1}
}

func (x *SanitizePerRoute) GetCustomAuthServerName() string {
	if x != nil {
		return x.CustomAuthServerName
	}
	return ""
}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDesc = []byte{
	0x0a, 0x5a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x73, 0x61,
	0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x65, 0x6e,
	0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x73, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x2e,
	0x76, 0x32, 0x22, 0x6d, 0x0a, 0x08, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x12, 0x2a,
	0x0a, 0x11, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x54, 0x6f, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x49, 0x0a, 0x10, 0x53, 0x61, 0x6e, 0x69, 0x74, 0x69, 0x7a, 0x65, 0x50, 0x65, 0x72,
	0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x17, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x41, 0x75,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x51, 0x5a, 0x4f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d,
	0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x65, 0x78, 0x74, 0x61, 0x75, 0x74, 0x68, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_goTypes = []any{
	(*Sanitize)(nil),         // 0: envoy.config.filter.http.sanitize.v2.Sanitize
	(*SanitizePerRoute)(nil), // 1: envoy.config.filter.http.sanitize.v2.SanitizePerRoute
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Sanitize); i {
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
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SanitizePerRoute); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_extensions_extauth_sanitize_proto_depIdxs = nil
}

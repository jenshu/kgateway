// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gloo/api/external/envoy/config/core/v3/socket_option.proto

package v3

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/solo-io/gloo/projects/gloo/pkg/api/external/udpa/annotations"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type SocketOption_SocketState int32

const (
	// Socket options are applied after socket creation but before binding the socket to a port
	SocketOption_STATE_PREBIND SocketOption_SocketState = 0
	// Socket options are applied after binding the socket to a port but before calling listen()
	SocketOption_STATE_BOUND SocketOption_SocketState = 1
	// Socket options are applied after calling listen()
	SocketOption_STATE_LISTENING SocketOption_SocketState = 2
)

// Enum value maps for SocketOption_SocketState.
var (
	SocketOption_SocketState_name = map[int32]string{
		0: "STATE_PREBIND",
		1: "STATE_BOUND",
		2: "STATE_LISTENING",
	}
	SocketOption_SocketState_value = map[string]int32{
		"STATE_PREBIND":   0,
		"STATE_BOUND":     1,
		"STATE_LISTENING": 2,
	}
)

func (x SocketOption_SocketState) Enum() *SocketOption_SocketState {
	p := new(SocketOption_SocketState)
	*p = x
	return p
}

func (x SocketOption_SocketState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SocketOption_SocketState) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_enumTypes[0].Descriptor()
}

func (SocketOption_SocketState) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_enumTypes[0]
}

func (x SocketOption_SocketState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SocketOption_SocketState.Descriptor instead.
func (SocketOption_SocketState) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescGZIP(), []int{0, 0}
}

// Generic socket option message. This would be used to set socket options that
// might not exist in upstream kernels or precompiled Envoy binaries.
// [#next-free-field: 7]
type SocketOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An optional name to give this socket option for debugging, etc.
	// Uniqueness is not required and no special meaning is assumed.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// Corresponding to the level value passed to setsockopt, such as IPPROTO_TCP
	Level int64 `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	// The numeric name as passed to setsockopt
	Name int64 `protobuf:"varint,3,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to Value:
	//
	//	*SocketOption_IntValue
	//	*SocketOption_BufValue
	Value isSocketOption_Value `protobuf_oneof:"value"`
	// The state in which the option will be applied. When used in BindConfig
	// STATE_PREBIND is currently the only valid value.
	State SocketOption_SocketState `protobuf:"varint,6,opt,name=state,proto3,enum=solo.io.envoy.config.core.v3.SocketOption_SocketState" json:"state,omitempty"`
}

func (x *SocketOption) Reset() {
	*x = SocketOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SocketOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SocketOption) ProtoMessage() {}

func (x *SocketOption) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SocketOption.ProtoReflect.Descriptor instead.
func (*SocketOption) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescGZIP(), []int{0}
}

func (x *SocketOption) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SocketOption) GetLevel() int64 {
	if x != nil {
		return x.Level
	}
	return 0
}

func (x *SocketOption) GetName() int64 {
	if x != nil {
		return x.Name
	}
	return 0
}

func (m *SocketOption) GetValue() isSocketOption_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *SocketOption) GetIntValue() int64 {
	if x, ok := x.GetValue().(*SocketOption_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (x *SocketOption) GetBufValue() []byte {
	if x, ok := x.GetValue().(*SocketOption_BufValue); ok {
		return x.BufValue
	}
	return nil
}

func (x *SocketOption) GetState() SocketOption_SocketState {
	if x != nil {
		return x.State
	}
	return SocketOption_STATE_PREBIND
}

type isSocketOption_Value interface {
	isSocketOption_Value()
}

type SocketOption_IntValue struct {
	// Because many sockopts take an int value.
	IntValue int64 `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3,oneof"`
}

type SocketOption_BufValue struct {
	// Otherwise it's a byte buffer.
	BufValue []byte `protobuf:"bytes,5,opt,name=buf_value,json=bufValue,proto3,oneof"`
}

func (*SocketOption_IntValue) isSocketOption_Value() {}

func (*SocketOption_BufValue) isSocketOption_Value() {}

var File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDesc = []byte{
	0x0a, 0x5b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1d, 0x75, 0x64, 0x70,
	0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x02, 0x0a, 0x0c, 0x53,
	0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6c, 0x65,
	0x76, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x08, 0x69, 0x6e,
	0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x62, 0x75, 0x66, 0x5f, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x62, 0x75, 0x66,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x56, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x33, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x08, 0xfa, 0x42,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x46, 0x0a,
	0x0b, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x11, 0x0a, 0x0d,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x50, 0x52, 0x45, 0x42, 0x49, 0x4e, 0x44, 0x10, 0x00, 0x12,
	0x0f, 0x0a, 0x0b, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01,
	0x12, 0x13, 0x0a, 0x0f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x45, 0x4e,
	0x49, 0x4e, 0x47, 0x10, 0x02, 0x3a, 0x2f, 0x8a, 0xc8, 0xde, 0x8e, 0x04, 0x29, 0x0a, 0x27, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x32, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x03, 0xf8, 0x42, 0x01, 0x42, 0xa2, 0x01, 0xb8, 0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0xd0,
	0xf5, 0x04, 0x01, 0xe2, 0xb5, 0xdf, 0xcb, 0x07, 0x02, 0x10, 0x02, 0x0a, 0x2a, 0x69, 0x6f, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e,
	0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x42, 0x11, 0x53, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescData = file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_goTypes = []any{
	(SocketOption_SocketState)(0), // 0: solo.io.envoy.config.core.v3.SocketOption.SocketState
	(*SocketOption)(nil),          // 1: solo.io.envoy.config.core.v3.SocketOption
}
var file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_depIdxs = []int32{
	0, // 0: solo.io.envoy.config.core.v3.SocketOption.state:type_name -> solo.io.envoy.config.core.v3.SocketOption.SocketState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() {
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_init()
}
func file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_init() {
	if File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SocketOption); i {
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
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_msgTypes[0].OneofWrappers = []any{
		(*SocketOption_IntValue)(nil),
		(*SocketOption_BufValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto = out.File
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gloo_api_external_envoy_config_core_v3_socket_option_proto_depIdxs = nil
}

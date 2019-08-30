// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-projects/projects/grpcserver/api/v1/artifact.proto

package v1

import (
	bytes "bytes"
	context "context"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type GetArtifactRequest struct {
	Ref                  *core.ResourceRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *GetArtifactRequest) Reset()         { *m = GetArtifactRequest{} }
func (m *GetArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*GetArtifactRequest) ProtoMessage()    {}
func (*GetArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{0}
}
func (m *GetArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactRequest.Unmarshal(m, b)
}
func (m *GetArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactRequest.Marshal(b, m, deterministic)
}
func (m *GetArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactRequest.Merge(m, src)
}
func (m *GetArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_GetArtifactRequest.Size(m)
}
func (m *GetArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactRequest proto.InternalMessageInfo

func (m *GetArtifactRequest) GetRef() *core.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

type GetArtifactResponse struct {
	Artifact             *v1.Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetArtifactResponse) Reset()         { *m = GetArtifactResponse{} }
func (m *GetArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*GetArtifactResponse) ProtoMessage()    {}
func (*GetArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{1}
}
func (m *GetArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetArtifactResponse.Unmarshal(m, b)
}
func (m *GetArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetArtifactResponse.Marshal(b, m, deterministic)
}
func (m *GetArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetArtifactResponse.Merge(m, src)
}
func (m *GetArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_GetArtifactResponse.Size(m)
}
func (m *GetArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetArtifactResponse proto.InternalMessageInfo

func (m *GetArtifactResponse) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type ListArtifactsRequest struct {
	Namespaces           []string `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListArtifactsRequest) Reset()         { *m = ListArtifactsRequest{} }
func (m *ListArtifactsRequest) String() string { return proto.CompactTextString(m) }
func (*ListArtifactsRequest) ProtoMessage()    {}
func (*ListArtifactsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{2}
}
func (m *ListArtifactsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArtifactsRequest.Unmarshal(m, b)
}
func (m *ListArtifactsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArtifactsRequest.Marshal(b, m, deterministic)
}
func (m *ListArtifactsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArtifactsRequest.Merge(m, src)
}
func (m *ListArtifactsRequest) XXX_Size() int {
	return xxx_messageInfo_ListArtifactsRequest.Size(m)
}
func (m *ListArtifactsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArtifactsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListArtifactsRequest proto.InternalMessageInfo

func (m *ListArtifactsRequest) GetNamespaces() []string {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

type ListArtifactsResponse struct {
	Artifacts            []*v1.Artifact `protobuf:"bytes,1,rep,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListArtifactsResponse) Reset()         { *m = ListArtifactsResponse{} }
func (m *ListArtifactsResponse) String() string { return proto.CompactTextString(m) }
func (*ListArtifactsResponse) ProtoMessage()    {}
func (*ListArtifactsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{3}
}
func (m *ListArtifactsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListArtifactsResponse.Unmarshal(m, b)
}
func (m *ListArtifactsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListArtifactsResponse.Marshal(b, m, deterministic)
}
func (m *ListArtifactsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListArtifactsResponse.Merge(m, src)
}
func (m *ListArtifactsResponse) XXX_Size() int {
	return xxx_messageInfo_ListArtifactsResponse.Size(m)
}
func (m *ListArtifactsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListArtifactsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListArtifactsResponse proto.InternalMessageInfo

func (m *ListArtifactsResponse) GetArtifacts() []*v1.Artifact {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

type CreateArtifactRequest struct {
	// Deprecated
	Ref *core.ResourceRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"` // Deprecated: Do not use.
	// Deprecated
	// Artifact data is a map to be compatible with Kubernetes ConfigMaps.
	// The key refers to the ConfigMap key, the value is the associated ConfigMap value.
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Deprecated: Do not use.
	Artifact             *v1.Artifact      `protobuf:"bytes,3,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CreateArtifactRequest) Reset()         { *m = CreateArtifactRequest{} }
func (m *CreateArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*CreateArtifactRequest) ProtoMessage()    {}
func (*CreateArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{4}
}
func (m *CreateArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArtifactRequest.Unmarshal(m, b)
}
func (m *CreateArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArtifactRequest.Marshal(b, m, deterministic)
}
func (m *CreateArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArtifactRequest.Merge(m, src)
}
func (m *CreateArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_CreateArtifactRequest.Size(m)
}
func (m *CreateArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArtifactRequest proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *CreateArtifactRequest) GetRef() *core.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

// Deprecated: Do not use.
func (m *CreateArtifactRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateArtifactRequest) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type CreateArtifactResponse struct {
	Artifact             *v1.Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreateArtifactResponse) Reset()         { *m = CreateArtifactResponse{} }
func (m *CreateArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*CreateArtifactResponse) ProtoMessage()    {}
func (*CreateArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{5}
}
func (m *CreateArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateArtifactResponse.Unmarshal(m, b)
}
func (m *CreateArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateArtifactResponse.Marshal(b, m, deterministic)
}
func (m *CreateArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateArtifactResponse.Merge(m, src)
}
func (m *CreateArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_CreateArtifactResponse.Size(m)
}
func (m *CreateArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateArtifactResponse proto.InternalMessageInfo

func (m *CreateArtifactResponse) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type UpdateArtifactRequest struct {
	// Deprecated
	Ref *core.ResourceRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"` // Deprecated: Do not use.
	// Deprecated
	// Artifact data is a map to be compatible with Kubernetes ConfigMaps.
	// The key refers to the ConfigMap key, the value is the associated ConfigMap value.
	Data                 map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Deprecated: Do not use.
	Artifact             *v1.Artifact      `protobuf:"bytes,3,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *UpdateArtifactRequest) Reset()         { *m = UpdateArtifactRequest{} }
func (m *UpdateArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateArtifactRequest) ProtoMessage()    {}
func (*UpdateArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{6}
}
func (m *UpdateArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateArtifactRequest.Unmarshal(m, b)
}
func (m *UpdateArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateArtifactRequest.Marshal(b, m, deterministic)
}
func (m *UpdateArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateArtifactRequest.Merge(m, src)
}
func (m *UpdateArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateArtifactRequest.Size(m)
}
func (m *UpdateArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateArtifactRequest proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *UpdateArtifactRequest) GetRef() *core.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

// Deprecated: Do not use.
func (m *UpdateArtifactRequest) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *UpdateArtifactRequest) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type UpdateArtifactResponse struct {
	Artifact             *v1.Artifact `protobuf:"bytes,1,opt,name=artifact,proto3" json:"artifact,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdateArtifactResponse) Reset()         { *m = UpdateArtifactResponse{} }
func (m *UpdateArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateArtifactResponse) ProtoMessage()    {}
func (*UpdateArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{7}
}
func (m *UpdateArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateArtifactResponse.Unmarshal(m, b)
}
func (m *UpdateArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateArtifactResponse.Marshal(b, m, deterministic)
}
func (m *UpdateArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateArtifactResponse.Merge(m, src)
}
func (m *UpdateArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateArtifactResponse.Size(m)
}
func (m *UpdateArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateArtifactResponse proto.InternalMessageInfo

func (m *UpdateArtifactResponse) GetArtifact() *v1.Artifact {
	if m != nil {
		return m.Artifact
	}
	return nil
}

type DeleteArtifactRequest struct {
	Ref                  *core.ResourceRef `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeleteArtifactRequest) Reset()         { *m = DeleteArtifactRequest{} }
func (m *DeleteArtifactRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteArtifactRequest) ProtoMessage()    {}
func (*DeleteArtifactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{8}
}
func (m *DeleteArtifactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArtifactRequest.Unmarshal(m, b)
}
func (m *DeleteArtifactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArtifactRequest.Marshal(b, m, deterministic)
}
func (m *DeleteArtifactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArtifactRequest.Merge(m, src)
}
func (m *DeleteArtifactRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteArtifactRequest.Size(m)
}
func (m *DeleteArtifactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArtifactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArtifactRequest proto.InternalMessageInfo

func (m *DeleteArtifactRequest) GetRef() *core.ResourceRef {
	if m != nil {
		return m.Ref
	}
	return nil
}

type DeleteArtifactResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteArtifactResponse) Reset()         { *m = DeleteArtifactResponse{} }
func (m *DeleteArtifactResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteArtifactResponse) ProtoMessage()    {}
func (*DeleteArtifactResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_31231f9377f1df05, []int{9}
}
func (m *DeleteArtifactResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteArtifactResponse.Unmarshal(m, b)
}
func (m *DeleteArtifactResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteArtifactResponse.Marshal(b, m, deterministic)
}
func (m *DeleteArtifactResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteArtifactResponse.Merge(m, src)
}
func (m *DeleteArtifactResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteArtifactResponse.Size(m)
}
func (m *DeleteArtifactResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteArtifactResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteArtifactResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetArtifactRequest)(nil), "glooeeapi.solo.io.GetArtifactRequest")
	proto.RegisterType((*GetArtifactResponse)(nil), "glooeeapi.solo.io.GetArtifactResponse")
	proto.RegisterType((*ListArtifactsRequest)(nil), "glooeeapi.solo.io.ListArtifactsRequest")
	proto.RegisterType((*ListArtifactsResponse)(nil), "glooeeapi.solo.io.ListArtifactsResponse")
	proto.RegisterType((*CreateArtifactRequest)(nil), "glooeeapi.solo.io.CreateArtifactRequest")
	proto.RegisterMapType((map[string]string)(nil), "glooeeapi.solo.io.CreateArtifactRequest.DataEntry")
	proto.RegisterType((*CreateArtifactResponse)(nil), "glooeeapi.solo.io.CreateArtifactResponse")
	proto.RegisterType((*UpdateArtifactRequest)(nil), "glooeeapi.solo.io.UpdateArtifactRequest")
	proto.RegisterMapType((map[string]string)(nil), "glooeeapi.solo.io.UpdateArtifactRequest.DataEntry")
	proto.RegisterType((*UpdateArtifactResponse)(nil), "glooeeapi.solo.io.UpdateArtifactResponse")
	proto.RegisterType((*DeleteArtifactRequest)(nil), "glooeeapi.solo.io.DeleteArtifactRequest")
	proto.RegisterType((*DeleteArtifactResponse)(nil), "glooeeapi.solo.io.DeleteArtifactResponse")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-projects/projects/grpcserver/api/v1/artifact.proto", fileDescriptor_31231f9377f1df05)
}

var fileDescriptor_31231f9377f1df05 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x5d, 0xd2, 0x81, 0xe8, 0xad, 0x40, 0x60, 0xda, 0x2a, 0xe4, 0x61, 0x9a, 0x22, 0x01, 0x45,
	0x80, 0x23, 0x0a, 0x02, 0xc4, 0x9e, 0x3a, 0x8a, 0x10, 0xd3, 0x78, 0x89, 0xc4, 0x0b, 0x0f, 0x48,
	0x5e, 0x7a, 0x1b, 0xcc, 0xba, 0xda, 0xd8, 0x6e, 0xa5, 0x3d, 0xf2, 0x03, 0x7c, 0x07, 0xdf, 0xc5,
	0x97, 0xa0, 0x24, 0x4d, 0x68, 0x52, 0x4f, 0x0b, 0xaa, 0x84, 0xf6, 0x54, 0xc7, 0x3d, 0xf7, 0x9c,
	0xa3, 0xe3, 0x7b, 0x6d, 0x38, 0x4a, 0xb8, 0xf9, 0xba, 0x38, 0xa1, 0xb1, 0x38, 0x0b, 0xb5, 0x98,
	0x89, 0xa7, 0x5c, 0xe4, 0xbf, 0x52, 0x89, 0x6f, 0x18, 0x1b, 0x1d, 0x96, 0x8b, 0x44, 0xc9, 0x58,
	0xa3, 0x5a, 0xa2, 0x0a, 0x99, 0xe4, 0xe1, 0xf2, 0x59, 0xc8, 0x94, 0xe1, 0x53, 0x16, 0x1b, 0x2a,
	0x95, 0x30, 0x82, 0xdc, 0x49, 0x66, 0x42, 0x20, 0x32, 0xc9, 0x69, 0x4a, 0x41, 0xb9, 0xf0, 0xbb,
	0x89, 0x48, 0x44, 0xf6, 0x6f, 0x98, 0xae, 0x72, 0xa0, 0x7f, 0x60, 0x11, 0x4d, 0x6b, 0xd7, 0xb4,
	0xd2, 0x2f, 0xab, 0x8a, 0xff, 0xe4, 0x22, 0xc7, 0xa7, 0xdc, 0x14, 0x25, 0x0a, 0xa7, 0x39, 0x3a,
	0x18, 0x01, 0x79, 0x8f, 0x66, 0xb4, 0xa2, 0x88, 0xf0, 0xfb, 0x02, 0xb5, 0x21, 0x8f, 0xa1, 0xa5,
	0x70, 0xea, 0x39, 0xfb, 0xce, 0xa0, 0x33, 0xbc, 0x47, 0x63, 0xa1, 0xb0, 0xb0, 0x4c, 0x23, 0xd4,
	0x62, 0xa1, 0x62, 0x8c, 0x70, 0x1a, 0xa5, 0xa8, 0xe0, 0x03, 0xdc, 0xad, 0x50, 0x68, 0x29, 0xe6,
	0x1a, 0xc9, 0x10, 0x6e, 0x14, 0xce, 0x56, 0x44, 0x7d, 0x9a, 0xda, 0x2e, 0x89, 0xca, 0x8a, 0x12,
	0x17, 0xbc, 0x84, 0xee, 0x31, 0xd7, 0x25, 0x97, 0x2e, 0xfc, 0xec, 0x01, 0xcc, 0xd9, 0x19, 0x6a,
	0xc9, 0x62, 0xd4, 0x9e, 0xb3, 0xdf, 0x1a, 0xb4, 0xa3, 0xb5, 0x9d, 0xe0, 0x23, 0xf4, 0x6a, 0x75,
	0x2b, 0x13, 0x2f, 0xa0, 0x5d, 0x90, 0xe7, 0x75, 0x17, 0xbb, 0xf8, 0x0b, 0x0c, 0x7e, 0xb8, 0xd0,
	0x7b, 0xab, 0x90, 0x19, 0xac, 0x07, 0x13, 0x36, 0x0b, 0xe6, 0xd0, 0xf5, 0x9c, 0x2c, 0x1c, 0x72,
	0x04, 0xbb, 0x13, 0x66, 0x98, 0xe7, 0x66, 0xda, 0x43, 0xba, 0xd1, 0x02, 0xd4, 0x2a, 0x44, 0xc7,
	0xcc, 0xb0, 0x77, 0x73, 0xa3, 0xce, 0x33, 0xaa, 0x8c, 0xa3, 0x92, 0x68, 0xab, 0x59, 0xa2, 0xfe,
	0x2b, 0x68, 0x97, 0x54, 0xe4, 0x36, 0xb4, 0x4e, 0xf1, 0x3c, 0x73, 0xdf, 0x8e, 0xd2, 0x25, 0xe9,
	0xc2, 0xb5, 0x25, 0x9b, 0x2d, 0xd0, 0x73, 0xb3, 0xbd, 0xfc, 0xe3, 0x8d, 0xfb, 0xda, 0x09, 0x8e,
	0xa1, 0x5f, 0x77, 0xb6, 0xc5, 0xc1, 0xa6, 0x89, 0x7e, 0x92, 0x93, 0xff, 0x93, 0xa8, 0x55, 0xe8,
	0x4a, 0x24, 0x5a, 0x77, 0xb6, 0x45, 0xa2, 0x63, 0xe8, 0x8d, 0x71, 0x86, 0x9b, 0x81, 0xfe, 0xd3,
	0xec, 0x7a, 0xd0, 0xaf, 0xb3, 0xe4, 0x9e, 0x86, 0x3f, 0x77, 0xa1, 0x53, 0x6c, 0x8e, 0x24, 0x27,
	0x5f, 0xa0, 0xb3, 0x36, 0xe5, 0xe4, 0xbe, 0x25, 0xf7, 0xcd, 0x8b, 0xc4, 0x7f, 0x70, 0x19, 0x2c,
	0x57, 0x0b, 0x76, 0xc8, 0x04, 0x6e, 0x56, 0x46, 0x98, 0x3c, 0xb4, 0x94, 0xda, 0x2e, 0x07, 0x7f,
	0x70, 0x39, 0xb0, 0x54, 0x49, 0xe0, 0x56, 0xb5, 0xab, 0xc9, 0xa0, 0xe9, 0x48, 0xfa, 0x8f, 0x1a,
	0x20, 0xd7, 0x85, 0xaa, 0x87, 0x6d, 0x15, 0xb2, 0x76, 0xaa, 0x55, 0xc8, 0xde, 0x39, 0xb9, 0x50,
	0xf5, 0x04, 0xad, 0x42, 0xd6, 0x56, 0xb1, 0x0a, 0xd9, 0xdb, 0x21, 0xd8, 0x39, 0x1c, 0xfd, 0xfa,
	0xbd, 0xe7, 0x7c, 0x3e, 0xd8, 0xe2, 0x3d, 0x3c, 0xb9, 0x9e, 0xbd, 0x39, 0xcf, 0xff, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xaf, 0xc3, 0xbf, 0xb6, 0x55, 0x07, 0x00, 0x00,
}

func (this *GetArtifactRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetArtifactRequest)
	if !ok {
		that2, ok := that.(GetArtifactRequest)
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
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *GetArtifactResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetArtifactResponse)
	if !ok {
		that2, ok := that.(GetArtifactResponse)
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
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListArtifactsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListArtifactsRequest)
	if !ok {
		that2, ok := that.(ListArtifactsRequest)
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
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if this.Namespaces[i] != that1.Namespaces[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *ListArtifactsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ListArtifactsResponse)
	if !ok {
		that2, ok := that.(ListArtifactsResponse)
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
	if len(this.Artifacts) != len(that1.Artifacts) {
		return false
	}
	for i := range this.Artifacts {
		if !this.Artifacts[i].Equal(that1.Artifacts[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CreateArtifactRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateArtifactRequest)
	if !ok {
		that2, ok := that.(CreateArtifactRequest)
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
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if len(this.Data) != len(that1.Data) {
		return false
	}
	for i := range this.Data {
		if this.Data[i] != that1.Data[i] {
			return false
		}
	}
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *CreateArtifactResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*CreateArtifactResponse)
	if !ok {
		that2, ok := that.(CreateArtifactResponse)
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
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpdateArtifactRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateArtifactRequest)
	if !ok {
		that2, ok := that.(UpdateArtifactRequest)
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
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if len(this.Data) != len(that1.Data) {
		return false
	}
	for i := range this.Data {
		if this.Data[i] != that1.Data[i] {
			return false
		}
	}
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *UpdateArtifactResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpdateArtifactResponse)
	if !ok {
		that2, ok := that.(UpdateArtifactResponse)
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
	if !this.Artifact.Equal(that1.Artifact) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DeleteArtifactRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeleteArtifactRequest)
	if !ok {
		that2, ok := that.(DeleteArtifactRequest)
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
	if !this.Ref.Equal(that1.Ref) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DeleteArtifactResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DeleteArtifactResponse)
	if !ok {
		that2, ok := that.(DeleteArtifactResponse)
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
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ArtifactApiClient is the client API for ArtifactApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ArtifactApiClient interface {
	GetArtifact(ctx context.Context, in *GetArtifactRequest, opts ...grpc.CallOption) (*GetArtifactResponse, error)
	ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error)
	CreateArtifact(ctx context.Context, in *CreateArtifactRequest, opts ...grpc.CallOption) (*CreateArtifactResponse, error)
	UpdateArtifact(ctx context.Context, in *UpdateArtifactRequest, opts ...grpc.CallOption) (*UpdateArtifactResponse, error)
	DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error)
}

type artifactApiClient struct {
	cc *grpc.ClientConn
}

func NewArtifactApiClient(cc *grpc.ClientConn) ArtifactApiClient {
	return &artifactApiClient{cc}
}

func (c *artifactApiClient) GetArtifact(ctx context.Context, in *GetArtifactRequest, opts ...grpc.CallOption) (*GetArtifactResponse, error) {
	out := new(GetArtifactResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/GetArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactApiClient) ListArtifacts(ctx context.Context, in *ListArtifactsRequest, opts ...grpc.CallOption) (*ListArtifactsResponse, error) {
	out := new(ListArtifactsResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/ListArtifacts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactApiClient) CreateArtifact(ctx context.Context, in *CreateArtifactRequest, opts ...grpc.CallOption) (*CreateArtifactResponse, error) {
	out := new(CreateArtifactResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/CreateArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactApiClient) UpdateArtifact(ctx context.Context, in *UpdateArtifactRequest, opts ...grpc.CallOption) (*UpdateArtifactResponse, error) {
	out := new(UpdateArtifactResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/UpdateArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactApiClient) DeleteArtifact(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*DeleteArtifactResponse, error) {
	out := new(DeleteArtifactResponse)
	err := c.cc.Invoke(ctx, "/glooeeapi.solo.io.ArtifactApi/DeleteArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactApiServer is the server API for ArtifactApi service.
type ArtifactApiServer interface {
	GetArtifact(context.Context, *GetArtifactRequest) (*GetArtifactResponse, error)
	ListArtifacts(context.Context, *ListArtifactsRequest) (*ListArtifactsResponse, error)
	CreateArtifact(context.Context, *CreateArtifactRequest) (*CreateArtifactResponse, error)
	UpdateArtifact(context.Context, *UpdateArtifactRequest) (*UpdateArtifactResponse, error)
	DeleteArtifact(context.Context, *DeleteArtifactRequest) (*DeleteArtifactResponse, error)
}

// UnimplementedArtifactApiServer can be embedded to have forward compatible implementations.
type UnimplementedArtifactApiServer struct {
}

func (*UnimplementedArtifactApiServer) GetArtifact(ctx context.Context, req *GetArtifactRequest) (*GetArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetArtifact not implemented")
}
func (*UnimplementedArtifactApiServer) ListArtifacts(ctx context.Context, req *ListArtifactsRequest) (*ListArtifactsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArtifacts not implemented")
}
func (*UnimplementedArtifactApiServer) CreateArtifact(ctx context.Context, req *CreateArtifactRequest) (*CreateArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateArtifact not implemented")
}
func (*UnimplementedArtifactApiServer) UpdateArtifact(ctx context.Context, req *UpdateArtifactRequest) (*UpdateArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateArtifact not implemented")
}
func (*UnimplementedArtifactApiServer) DeleteArtifact(ctx context.Context, req *DeleteArtifactRequest) (*DeleteArtifactResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteArtifact not implemented")
}

func RegisterArtifactApiServer(s *grpc.Server, srv ArtifactApiServer) {
	s.RegisterService(&_ArtifactApi_serviceDesc, srv)
}

func _ArtifactApi_GetArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).GetArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/GetArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).GetArtifact(ctx, req.(*GetArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactApi_ListArtifacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtifactsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).ListArtifacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/ListArtifacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).ListArtifacts(ctx, req.(*ListArtifactsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactApi_CreateArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).CreateArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/CreateArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).CreateArtifact(ctx, req.(*CreateArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactApi_UpdateArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).UpdateArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/UpdateArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).UpdateArtifact(ctx, req.(*UpdateArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactApi_DeleteArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactApiServer).DeleteArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/glooeeapi.solo.io.ArtifactApi/DeleteArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactApiServer).DeleteArtifact(ctx, req.(*DeleteArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArtifactApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "glooeeapi.solo.io.ArtifactApi",
	HandlerType: (*ArtifactApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArtifact",
			Handler:    _ArtifactApi_GetArtifact_Handler,
		},
		{
			MethodName: "ListArtifacts",
			Handler:    _ArtifactApi_ListArtifacts_Handler,
		},
		{
			MethodName: "CreateArtifact",
			Handler:    _ArtifactApi_CreateArtifact_Handler,
		},
		{
			MethodName: "UpdateArtifact",
			Handler:    _ArtifactApi_UpdateArtifact_Handler,
		},
		{
			MethodName: "DeleteArtifact",
			Handler:    _ArtifactApi_DeleteArtifact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/solo-io/solo-projects/projects/grpcserver/api/v1/artifact.proto",
}

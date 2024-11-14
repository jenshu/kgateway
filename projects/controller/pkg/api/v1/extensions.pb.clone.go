// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/extensions.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_structpb "google.golang.org/protobuf/types/known/structpb"
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
func (m *Extensions) Clone() proto.Message {
	var target *Extensions
	if m == nil {
		return target
	}
	target = &Extensions{}

	if m.GetConfigs() != nil {
		target.Configs = make(map[string]*google_golang_org_protobuf_types_known_structpb.Struct, len(m.GetConfigs()))
		for k, v := range m.GetConfigs() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.Configs[k] = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Struct)
			} else {
				target.Configs[k] = proto.Clone(v).(*google_golang_org_protobuf_types_known_structpb.Struct)
			}

		}
	}

	return target
}

// Clone function
func (m *Extension) Clone() proto.Message {
	var target *Extension
	if m == nil {
		return target
	}
	target = &Extension{}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*google_golang_org_protobuf_types_known_structpb.Struct)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*google_golang_org_protobuf_types_known_structpb.Struct)
	}

	return target
}

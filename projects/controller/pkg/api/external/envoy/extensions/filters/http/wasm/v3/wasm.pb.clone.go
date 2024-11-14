// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/extensions/filters/http/wasm/v3/wasm.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_wasm_v3 "github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/extensions/wasm/v3"
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
func (m *Wasm) Clone() proto.Message {
	var target *Wasm
	if m == nil {
		return target
	}
	target = &Wasm{}

	if h, ok := interface{}(m.GetConfig()).(clone.Cloner); ok {
		target.Config = h.Clone().(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_wasm_v3.PluginConfig)
	} else {
		target.Config = proto.Clone(m.GetConfig()).(*github_com_solo_io_gloo_projects_gloo_pkg_api_external_envoy_extensions_wasm_v3.PluginConfig)
	}

	return target
}

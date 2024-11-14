// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/rest/rest.proto

package rest

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *ServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ServiceSpec)
	if !ok {
		that2, ok := that.(ServiceSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetTransformations()) != len(target.GetTransformations()) {
		return false
	}
	for k, v := range m.GetTransformations() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetTransformations()[k]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetTransformations()[k]) {
				return false
			}
		}

	}

	if h, ok := interface{}(m.GetSwaggerInfo()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSwaggerInfo()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSwaggerInfo(), target.GetSwaggerInfo()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *DestinationSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*DestinationSpec)
	if !ok {
		that2, ok := that.(DestinationSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetFunctionName(), target.GetFunctionName()) != 0 {
		return false
	}

	if h, ok := interface{}(m.GetParameters()).(equality.Equalizer); ok {
		if !h.Equal(target.GetParameters()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetParameters(), target.GetParameters()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetResponseTransformation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetResponseTransformation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetResponseTransformation(), target.GetResponseTransformation()) {
			return false
		}
	}

	return true
}

// Equal function
func (m *ServiceSpec_SwaggerInfo) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ServiceSpec_SwaggerInfo)
	if !ok {
		that2, ok := that.(ServiceSpec_SwaggerInfo)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	switch m.SwaggerSpec.(type) {

	case *ServiceSpec_SwaggerInfo_Url:
		if _, ok := target.SwaggerSpec.(*ServiceSpec_SwaggerInfo_Url); !ok {
			return false
		}

		if strings.Compare(m.GetUrl(), target.GetUrl()) != 0 {
			return false
		}

	case *ServiceSpec_SwaggerInfo_Inline:
		if _, ok := target.SwaggerSpec.(*ServiceSpec_SwaggerInfo_Inline); !ok {
			return false
		}

		if strings.Compare(m.GetInline(), target.GetInline()) != 0 {
			return false
		}

	default:
		// m is nil but target is not nil
		if m.SwaggerSpec != target.SwaggerSpec {
			return false
		}
	}

	return true
}

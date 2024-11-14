// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/core/v3/address.proto

package v3

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/solo-io/protoc-gen-ext/pkg/clone"
	"google.golang.org/protobuf/proto"

	google_golang_org_protobuf_types_known_wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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
func (m *Pipe) Clone() proto.Message {
	var target *Pipe
	if m == nil {
		return target
	}
	target = &Pipe{}

	target.Path = m.GetPath()

	target.Mode = m.GetMode()

	return target
}

// Clone function
func (m *SocketAddress) Clone() proto.Message {
	var target *SocketAddress
	if m == nil {
		return target
	}
	target = &SocketAddress{}

	target.Protocol = m.GetProtocol()

	target.Address = m.GetAddress()

	target.ResolverName = m.GetResolverName()

	target.Ipv4Compat = m.GetIpv4Compat()

	switch m.PortSpecifier.(type) {

	case *SocketAddress_PortValue:

		target.PortSpecifier = &SocketAddress_PortValue{
			PortValue: m.GetPortValue(),
		}

	case *SocketAddress_NamedPort:

		target.PortSpecifier = &SocketAddress_NamedPort{
			NamedPort: m.GetNamedPort(),
		}

	}

	return target
}

// Clone function
func (m *TcpKeepalive) Clone() proto.Message {
	var target *TcpKeepalive
	if m == nil {
		return target
	}
	target = &TcpKeepalive{}

	if h, ok := interface{}(m.GetKeepaliveProbes()).(clone.Cloner); ok {
		target.KeepaliveProbes = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.KeepaliveProbes = proto.Clone(m.GetKeepaliveProbes()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetKeepaliveTime()).(clone.Cloner); ok {
		target.KeepaliveTime = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.KeepaliveTime = proto.Clone(m.GetKeepaliveTime()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	if h, ok := interface{}(m.GetKeepaliveInterval()).(clone.Cloner); ok {
		target.KeepaliveInterval = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.KeepaliveInterval = proto.Clone(m.GetKeepaliveInterval()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}

// Clone function
func (m *BindConfig) Clone() proto.Message {
	var target *BindConfig
	if m == nil {
		return target
	}
	target = &BindConfig{}

	if h, ok := interface{}(m.GetSourceAddress()).(clone.Cloner); ok {
		target.SourceAddress = h.Clone().(*SocketAddress)
	} else {
		target.SourceAddress = proto.Clone(m.GetSourceAddress()).(*SocketAddress)
	}

	if h, ok := interface{}(m.GetFreebind()).(clone.Cloner); ok {
		target.Freebind = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	} else {
		target.Freebind = proto.Clone(m.GetFreebind()).(*google_golang_org_protobuf_types_known_wrapperspb.BoolValue)
	}

	if m.GetSocketOptions() != nil {
		target.SocketOptions = make([]*SocketOption, len(m.GetSocketOptions()))
		for idx, v := range m.GetSocketOptions() {

			if h, ok := interface{}(v).(clone.Cloner); ok {
				target.SocketOptions[idx] = h.Clone().(*SocketOption)
			} else {
				target.SocketOptions[idx] = proto.Clone(v).(*SocketOption)
			}

		}
	}

	return target
}

// Clone function
func (m *Address) Clone() proto.Message {
	var target *Address
	if m == nil {
		return target
	}
	target = &Address{}

	switch m.Address.(type) {

	case *Address_SocketAddress:

		if h, ok := interface{}(m.GetSocketAddress()).(clone.Cloner); ok {
			target.Address = &Address_SocketAddress{
				SocketAddress: h.Clone().(*SocketAddress),
			}
		} else {
			target.Address = &Address_SocketAddress{
				SocketAddress: proto.Clone(m.GetSocketAddress()).(*SocketAddress),
			}
		}

	case *Address_Pipe:

		if h, ok := interface{}(m.GetPipe()).(clone.Cloner); ok {
			target.Address = &Address_Pipe{
				Pipe: h.Clone().(*Pipe),
			}
		} else {
			target.Address = &Address_Pipe{
				Pipe: proto.Clone(m.GetPipe()).(*Pipe),
			}
		}

	}

	return target
}

// Clone function
func (m *CidrRange) Clone() proto.Message {
	var target *CidrRange
	if m == nil {
		return target
	}
	target = &CidrRange{}

	target.AddressPrefix = m.GetAddressPrefix()

	if h, ok := interface{}(m.GetPrefixLen()).(clone.Cloner); ok {
		target.PrefixLen = h.Clone().(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	} else {
		target.PrefixLen = proto.Clone(m.GetPrefixLen()).(*google_golang_org_protobuf_types_known_wrapperspb.UInt32Value)
	}

	return target
}

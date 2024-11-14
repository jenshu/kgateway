// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/options/grpc_json/grpc_json.proto

package grpc_json

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"
	"strconv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = strconv.Itoa
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *GrpcJsonTranscoder) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("grpc_json.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/grpc_json.GrpcJsonTranscoder")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Services")); err != nil {
		return 0, err
	}
	for i, v := range m.GetServices() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if h, ok := interface{}(m.GetPrintOptions()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PrintOptions")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPrintOptions(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PrintOptions")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("MatchIncomingRequestRoute")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetMatchIncomingRequestRoute())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("IgnoredQueryParameters")); err != nil {
		return 0, err
	}
	for i, v := range m.GetIgnoredQueryParameters() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if _, err = hasher.Write([]byte("v")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	if _, err = hasher.Write([]byte("AutoMapping")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAutoMapping())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("IgnoreUnknownQueryParameters")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetIgnoreUnknownQueryParameters())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ConvertGrpcStatus")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetConvertGrpcStatus())
	if err != nil {
		return 0, err
	}

	switch m.DescriptorSet.(type) {

	case *GrpcJsonTranscoder_ProtoDescriptor:

		if _, err = hasher.Write([]byte("ProtoDescriptor")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetProtoDescriptor())); err != nil {
			return 0, err
		}

	case *GrpcJsonTranscoder_ProtoDescriptorBin:

		if _, err = hasher.Write([]byte("ProtoDescriptorBin")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write(m.GetProtoDescriptorBin()); err != nil {
			return 0, err
		}

	case *GrpcJsonTranscoder_ProtoDescriptorConfigMap:

		if h, ok := interface{}(m.GetProtoDescriptorConfigMap()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ProtoDescriptorConfigMap")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetProtoDescriptorConfigMap(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ProtoDescriptorConfigMap")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *GrpcJsonTranscoder_PrintOptions) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("grpc_json.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/grpc_json.GrpcJsonTranscoder_PrintOptions")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AddWhitespace")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAddWhitespace())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AlwaysPrintPrimitiveFields")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAlwaysPrintPrimitiveFields())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("AlwaysPrintEnumsAsInts")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAlwaysPrintEnumsAsInts())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("PreserveProtoFieldNames")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetPreserveProtoFieldNames())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *GrpcJsonTranscoder_DescriptorConfigMap) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("grpc_json.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/options/grpc_json.GrpcJsonTranscoder_DescriptorConfigMap")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetConfigMapRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ConfigMapRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetConfigMapRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ConfigMapRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Key")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/type/tracing/v3/custom_tag.proto

package v3

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

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
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *CustomTag) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.tracing.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/tracing/v3.CustomTag")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetTag())); err != nil {
		return 0, err
	}

	switch m.Type.(type) {

	case *CustomTag_Literal_:

		if h, ok := interface{}(m.GetLiteral()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Literal")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetLiteral(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Literal")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *CustomTag_Environment_:

		if h, ok := interface{}(m.GetEnvironment()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Environment")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetEnvironment(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Environment")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *CustomTag_RequestHeader:

		if h, ok := interface{}(m.GetRequestHeader()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestHeader")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestHeader(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestHeader")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *CustomTag_Metadata_:

		if h, ok := interface{}(m.GetMetadata()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Metadata")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetMetadata(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Metadata")); err != nil {
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

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *CustomTag_Literal) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.tracing.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/tracing/v3.CustomTag_Literal")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetValue())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *CustomTag_Environment) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.tracing.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/tracing/v3.CustomTag_Environment")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetDefaultValue())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *CustomTag_Header) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.tracing.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/tracing/v3.CustomTag_Header")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetDefaultValue())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *CustomTag_Metadata) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.type.tracing.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/type/tracing/v3.CustomTag_Metadata")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetKind()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Kind")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetKind(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Kind")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetMetadataKey()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MetadataKey")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadataKey(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MetadataKey")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetDefaultValue())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

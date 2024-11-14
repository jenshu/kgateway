// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/config/core/v3/base.proto

package v3

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
func (m *Locality) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.Locality")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Region")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRegion())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Zone")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetZone())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("SubZone")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetSubZone())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *BuildVersion) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.BuildVersion")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetVersion()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Version")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetVersion(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Version")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Extension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.Extension")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Category")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetCategory())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("TypeDescriptor")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetTypeDescriptor())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetVersion()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Version")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetVersion(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Version")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Disabled")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetDisabled())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Node) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.Node")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Id")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetId())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Cluster")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetCluster())); err != nil {
		return 0, err
	}

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

	if h, ok := interface{}(m.GetLocality()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Locality")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLocality(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Locality")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("UserAgentName")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetUserAgentName())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Extensions")); err != nil {
		return 0, err
	}
	for i, v := range m.GetExtensions() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if _, err = hasher.Write([]byte("ClientFeatures")); err != nil {
		return 0, err
	}
	for i, v := range m.GetClientFeatures() {
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

	if _, err = hasher.Write([]byte("ListeningAddresses")); err != nil {
		return 0, err
	}
	for i, v := range m.GetListeningAddresses() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	switch m.UserAgentVersionType.(type) {

	case *Node_UserAgentVersion:

		if _, err = hasher.Write([]byte("UserAgentVersion")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetUserAgentVersion())); err != nil {
			return 0, err
		}

	case *Node_UserAgentBuildVersion:

		if h, ok := interface{}(m.GetUserAgentBuildVersion()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("UserAgentBuildVersion")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetUserAgentBuildVersion(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("UserAgentBuildVersion")); err != nil {
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
func (m *Metadata) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.Metadata")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetFilterMetadata() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("v")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("v")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
			}

			if _, err = innerHash.Write([]byte("k")); err != nil {
				return 0, err
			}
			if _, err = innerHash.Write([]byte(k)); err != nil {
				return 0, err
			}

			result = result ^ innerHash.Sum64()
		}
		err = binary.Write(hasher, binary.LittleEndian, result)
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RuntimeUInt32) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.RuntimeUInt32")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetDefaultValue())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RuntimeKey")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRuntimeKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RuntimeDouble) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.RuntimeDouble")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetDefaultValue())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("RuntimeKey")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRuntimeKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RuntimeFeatureFlag) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.RuntimeFeatureFlag")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDefaultValue()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDefaultValue(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("RuntimeKey")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRuntimeKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *HeaderValue) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.HeaderValue")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Key")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Value")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetValue())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *HeaderValueOption) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.HeaderValueOption")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHeader()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Header")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeader(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Header")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAppend()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Append")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAppend(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Append")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *HeaderMap) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.HeaderMap")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Headers")); err != nil {
		return 0, err
	}
	for i, v := range m.GetHeaders() {
		if _, err = hasher.Write([]byte(strconv.Itoa(i))); err != nil {
			return 0, err
		}

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("v")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("v")); err != nil {
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
func (m *DataSource) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.DataSource")); err != nil {
		return 0, err
	}

	switch m.Specifier.(type) {

	case *DataSource_Filename:

		if _, err = hasher.Write([]byte("Filename")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetFilename())); err != nil {
			return 0, err
		}

	case *DataSource_InlineBytes:

		if _, err = hasher.Write([]byte("InlineBytes")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write(m.GetInlineBytes()); err != nil {
			return 0, err
		}

	case *DataSource_InlineString:

		if _, err = hasher.Write([]byte("InlineString")); err != nil {
			return 0, err
		}
		if _, err = hasher.Write([]byte(m.GetInlineString())); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RetryPolicy) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.RetryPolicy")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRetryBackOff()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RetryBackOff")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRetryBackOff(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RetryBackOff")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetNumRetries()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("NumRetries")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetNumRetries(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("NumRetries")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RemoteDataSource) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.RemoteDataSource")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHttpUri()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HttpUri")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHttpUri(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HttpUri")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Sha256")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetSha256())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRetryPolicy()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RetryPolicy")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRetryPolicy(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RetryPolicy")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *AsyncDataSource) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.AsyncDataSource")); err != nil {
		return 0, err
	}

	switch m.Specifier.(type) {

	case *AsyncDataSource_Local:

		if h, ok := interface{}(m.GetLocal()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Local")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetLocal(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Local")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *AsyncDataSource_Remote:

		if h, ok := interface{}(m.GetRemote()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Remote")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRemote(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Remote")); err != nil {
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
func (m *TransportSocket) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.TransportSocket")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Name")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	switch m.ConfigType.(type) {

	case *TransportSocket_TypedConfig:

		if h, ok := interface{}(m.GetTypedConfig()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TypedConfig")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTypedConfig(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TypedConfig")); err != nil {
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
func (m *RuntimeFractionalPercent) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.RuntimeFractionalPercent")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDefaultValue()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDefaultValue(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DefaultValue")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("RuntimeKey")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetRuntimeKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ControlPlane) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.config.core.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/config/core/v3.ControlPlane")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Identifier")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetIdentifier())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

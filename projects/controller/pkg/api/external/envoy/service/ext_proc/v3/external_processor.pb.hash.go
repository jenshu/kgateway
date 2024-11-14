// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/external/envoy/service/ext_proc/v3/external_processor.proto

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
func (m *ProcessingRequest) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.ProcessingRequest")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMetadataContext()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("MetadataContext")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMetadataContext(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("MetadataContext")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetAttributes() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	err = binary.Write(hasher, binary.LittleEndian, m.GetObservabilityMode())
	if err != nil {
		return 0, err
	}

	switch m.Request.(type) {

	case *ProcessingRequest_RequestHeaders:

		if h, ok := interface{}(m.GetRequestHeaders()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestHeaders")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestHeaders(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestHeaders")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingRequest_ResponseHeaders:

		if h, ok := interface{}(m.GetResponseHeaders()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ResponseHeaders")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetResponseHeaders(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ResponseHeaders")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingRequest_RequestBody:

		if h, ok := interface{}(m.GetRequestBody()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestBody")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestBody(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestBody")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingRequest_ResponseBody:

		if h, ok := interface{}(m.GetResponseBody()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ResponseBody")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetResponseBody(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ResponseBody")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingRequest_RequestTrailers:

		if h, ok := interface{}(m.GetRequestTrailers()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestTrailers")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestTrailers(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestTrailers")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingRequest_ResponseTrailers:

		if h, ok := interface{}(m.GetResponseTrailers()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ResponseTrailers")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetResponseTrailers(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ResponseTrailers")); err != nil {
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
func (m *ProcessingResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.ProcessingResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetDynamicMetadata()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("DynamicMetadata")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetDynamicMetadata(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("DynamicMetadata")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetModeOverride()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ModeOverride")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetModeOverride(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ModeOverride")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetOverrideMessageTimeout()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("OverrideMessageTimeout")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetOverrideMessageTimeout(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("OverrideMessageTimeout")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.Response.(type) {

	case *ProcessingResponse_RequestHeaders:

		if h, ok := interface{}(m.GetRequestHeaders()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestHeaders")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestHeaders(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestHeaders")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingResponse_ResponseHeaders:

		if h, ok := interface{}(m.GetResponseHeaders()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ResponseHeaders")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetResponseHeaders(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ResponseHeaders")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingResponse_RequestBody:

		if h, ok := interface{}(m.GetRequestBody()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestBody")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestBody(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestBody")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingResponse_ResponseBody:

		if h, ok := interface{}(m.GetResponseBody()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ResponseBody")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetResponseBody(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ResponseBody")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingResponse_RequestTrailers:

		if h, ok := interface{}(m.GetRequestTrailers()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RequestTrailers")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRequestTrailers(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RequestTrailers")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingResponse_ResponseTrailers:

		if h, ok := interface{}(m.GetResponseTrailers()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ResponseTrailers")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetResponseTrailers(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ResponseTrailers")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *ProcessingResponse_ImmediateResponse:

		if h, ok := interface{}(m.GetImmediateResponse()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ImmediateResponse")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetImmediateResponse(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ImmediateResponse")); err != nil {
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
func (m *HttpHeaders) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.HttpHeaders")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHeaders()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Headers")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeaders(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Headers")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetAttributes() {
			innerHash.Reset()

			if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
				if _, err = innerHash.Write([]byte("")); err != nil {
					return 0, err
				}
				if _, err = h.Hash(innerHash); err != nil {
					return 0, err
				}
			} else {
				if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
					return 0, err
				} else {
					if _, err = innerHash.Write([]byte("")); err != nil {
						return 0, err
					}
					if err := binary.Write(innerHash, binary.LittleEndian, fieldValue); err != nil {
						return 0, err
					}
				}
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

	err = binary.Write(hasher, binary.LittleEndian, m.GetEndOfStream())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *HttpBody) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.HttpBody")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write(m.GetBody()); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetEndOfStream())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *HttpTrailers) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.HttpTrailers")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTrailers()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Trailers")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTrailers(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Trailers")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
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
func (m *HeadersResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.HeadersResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetResponse()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Response")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponse(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Response")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
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
func (m *TrailersResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.TrailersResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHeaderMutation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HeaderMutation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeaderMutation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HeaderMutation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
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
func (m *BodyResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.BodyResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetResponse()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Response")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResponse(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Response")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
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
func (m *CommonResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.CommonResponse")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetStatus())
	if err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetHeaderMutation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HeaderMutation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeaderMutation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HeaderMutation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetBodyMutation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("BodyMutation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBodyMutation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("BodyMutation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetTrailers()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Trailers")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTrailers(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Trailers")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetClearRouteCache())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *ImmediateResponse) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.ImmediateResponse")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Status")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Status")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHeaders()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Headers")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeaders(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Headers")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write(m.GetBody()); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetGrpcStatus()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("GrpcStatus")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetGrpcStatus(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("GrpcStatus")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte(m.GetDetails())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *GrpcStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.GrpcStatus")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetStatus())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *HeaderMutation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.HeaderMutation")); err != nil {
		return 0, err
	}

	for _, v := range m.GetSetHeaders() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetRemoveHeaders() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
//
// Deprecated: due to hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
// Prefer the HashUnique function instead.
func (m *BodyMutation) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("solo.io.envoy.service.ext_proc.v3.github.com/solo-io/gloo/projects/controller/pkg/api/external/envoy/service/ext_proc/v3.BodyMutation")); err != nil {
		return 0, err
	}

	switch m.Mutation.(type) {

	case *BodyMutation_Body:

		if _, err = hasher.Write(m.GetBody()); err != nil {
			return 0, err
		}

	case *BodyMutation_ClearBody:

		err = binary.Write(hasher, binary.LittleEndian, m.GetClearBody())
		if err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

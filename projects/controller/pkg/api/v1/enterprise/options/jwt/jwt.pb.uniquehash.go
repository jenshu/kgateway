// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/controller/api/v1/enterprise/options/jwt/jwt.proto

package jwt

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
func (m *JwtStagedVhostExtension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.JwtStagedVhostExtension")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("BeforeExtAuth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBeforeExtAuth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("BeforeExtAuth")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AfterExtAuth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAfterExtAuth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AfterExtAuth")); err != nil {
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
func (m *JwtStagedRouteProvidersExtension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.JwtStagedRouteProvidersExtension")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("BeforeExtAuth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBeforeExtAuth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("BeforeExtAuth")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AfterExtAuth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAfterExtAuth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AfterExtAuth")); err != nil {
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
func (m *JwtStagedRouteExtension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.JwtStagedRouteExtension")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetBeforeExtAuth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("BeforeExtAuth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBeforeExtAuth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("BeforeExtAuth")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAfterExtAuth()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AfterExtAuth")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAfterExtAuth(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AfterExtAuth")); err != nil {
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
func (m *VhostExtension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.VhostExtension")); err != nil {
		return 0, err
	}

	{
		var result uint64
		innerHash := fnv.New64()
		for k, v := range m.GetProviders() {
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

	if _, err = hasher.Write([]byte("AllowMissingOrFailedJwt")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAllowMissingOrFailedJwt())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ValidationPolicy")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetValidationPolicy())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RouteExtension) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.RouteExtension")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Disable")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetDisable())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *Provider) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.Provider")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetJwks()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Jwks")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetJwks(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Jwks")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("Audiences")); err != nil {
		return 0, err
	}
	for i, v := range m.GetAudiences() {
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

	if _, err = hasher.Write([]byte("Issuer")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetIssuer())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetTokenSource()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("TokenSource")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetTokenSource(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("TokenSource")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if _, err = hasher.Write([]byte("KeepToken")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetKeepToken())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("ClaimsToHeaders")); err != nil {
		return 0, err
	}
	for i, v := range m.GetClaimsToHeaders() {
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

	if h, ok := interface{}(m.GetClockSkewSeconds()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("ClockSkewSeconds")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetClockSkewSeconds(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("ClockSkewSeconds")); err != nil {
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
func (m *Jwks) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.Jwks")); err != nil {
		return 0, err
	}

	switch m.Jwks.(type) {

	case *Jwks_Remote:

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

	case *Jwks_Local:

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

	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *RemoteJwks) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.RemoteJwks")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Url")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetUrl())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetUpstreamRef(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("UpstreamRef")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetCacheDuration()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CacheDuration")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCacheDuration(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CacheDuration")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAsyncFetch()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AsyncFetch")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAsyncFetch(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AsyncFetch")); err != nil {
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
func (m *LocalJwks) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.LocalJwks")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Key")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetKey())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *TokenSource) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.TokenSource")); err != nil {
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

	if _, err = hasher.Write([]byte("QueryParams")); err != nil {
		return 0, err
	}
	for i, v := range m.GetQueryParams() {
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

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *ClaimToHeader) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.ClaimToHeader")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Claim")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetClaim())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Header")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetHeader())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Append")); err != nil {
		return 0, err
	}
	err = binary.Write(hasher, binary.LittleEndian, m.GetAppend())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// HashUnique function generates a hash of the object that is unique to the object by
// hashing field name and value pairs.
// Replaces Hash due to original hashing implemention only using field values. The omission
// of the field name in the hash calculation can lead to hash collisions.
func (m *TokenSource_HeaderSource) HashUnique(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("jwt.options.gloo.solo.io.github.com/solo-io/gloo/projects/controller/pkg/api/v1/enterprise/options/jwt.TokenSource_HeaderSource")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Header")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetHeader())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte("Prefix")); err != nil {
		return 0, err
	}
	if _, err = hasher.Write([]byte(m.GetPrefix())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

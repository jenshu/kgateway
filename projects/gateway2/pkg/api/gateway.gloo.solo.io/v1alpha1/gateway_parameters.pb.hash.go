// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/gateway_parameters.proto

package v1alpha1

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
func (m *GatewayParametersSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.GatewayParametersSpec")); err != nil {
		return 0, err
	}

	switch m.EnvironmentType.(type) {

	case *GatewayParametersSpec_Kube:

		if h, ok := interface{}(m.GetKube()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Kube")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetKube(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Kube")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *GatewayParametersSpec_SelfManaged:

		if h, ok := interface{}(m.GetSelfManaged()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("SelfManaged")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetSelfManaged(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("SelfManaged")); err != nil {
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
func (m *KubernetesProxyConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.KubernetesProxyConfig")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEnvoyContainer()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnvoyContainer")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnvoyContainer(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnvoyContainer")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSdsContainer()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SdsContainer")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSdsContainer(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SdsContainer")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetPodTemplate()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("PodTemplate")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetPodTemplate(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("PodTemplate")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetService()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Service")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetService(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Service")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetIstio()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Istio")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIstio(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Istio")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStats()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Stats")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStats(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Stats")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetAiExtension()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("AiExtension")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetAiExtension(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("AiExtension")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.WorkloadType.(type) {

	case *KubernetesProxyConfig_Deployment:

		if h, ok := interface{}(m.GetDeployment()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Deployment")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDeployment(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Deployment")); err != nil {
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
func (m *ProxyDeployment) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.ProxyDeployment")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetReplicas()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Replicas")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetReplicas(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Replicas")); err != nil {
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
func (m *EnvoyContainer) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.EnvoyContainer")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetBootstrap()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Bootstrap")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBootstrap(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Bootstrap")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetImage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Image")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetImage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Image")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SecurityContext")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSecurityContext(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SecurityContext")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResources()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Resources")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResources(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Resources")); err != nil {
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
func (m *EnvoyBootstrap) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.EnvoyBootstrap")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetLogLevel()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LogLevel")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLogLevel(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LogLevel")); err != nil {
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
		for k, v := range m.GetComponentLogLevels() {
			innerHash.Reset()

			if _, err = innerHash.Write([]byte(v)); err != nil {
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

// Hash function
func (m *IstioIntegration) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.IstioIntegration")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetIstioProxyContainer()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IstioProxyContainer")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIstioProxyContainer(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IstioProxyContainer")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetCustomSidecars() {

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

	return hasher.Sum64(), nil
}

// Hash function
func (m *SdsContainer) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.SdsContainer")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetImage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Image")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetImage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Image")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SecurityContext")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSecurityContext(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SecurityContext")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResources()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Resources")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResources(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Resources")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetBootstrap()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Bootstrap")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetBootstrap(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Bootstrap")); err != nil {
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
func (m *SdsBootstrap) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.SdsBootstrap")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetLogLevel()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LogLevel")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLogLevel(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LogLevel")); err != nil {
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
func (m *IstioContainer) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.IstioContainer")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetImage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Image")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetImage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Image")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SecurityContext")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSecurityContext(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SecurityContext")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResources()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Resources")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResources(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Resources")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetLogLevel()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("LogLevel")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetLogLevel(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("LogLevel")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetIstioDiscoveryAddress()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IstioDiscoveryAddress")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIstioDiscoveryAddress(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IstioDiscoveryAddress")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetIstioMetaMeshId()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IstioMetaMeshId")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIstioMetaMeshId(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IstioMetaMeshId")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetIstioMetaClusterId()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("IstioMetaClusterId")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetIstioMetaClusterId(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("IstioMetaClusterId")); err != nil {
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
func (m *GatewayParametersStatus) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.GatewayParametersStatus")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *StatsConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.StatsConfig")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEnabled()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Enabled")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnabled(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Enabled")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetRoutePrefixRewrite()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RoutePrefixRewrite")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRoutePrefixRewrite(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RoutePrefixRewrite")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetEnableStatsRoute()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("EnableStatsRoute")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnableStatsRoute(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("EnableStatsRoute")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetStatsRoutePrefixRewrite()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("StatsRoutePrefixRewrite")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetStatsRoutePrefixRewrite(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("StatsRoutePrefixRewrite")); err != nil {
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
func (m *AiExtension) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gateway.gloo.solo.io.github.com/solo-io/gloo/projects/gateway2/pkg/api/gateway.gloo.solo.io/v1alpha1.AiExtension")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetEnabled()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Enabled")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetEnabled(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Enabled")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetImage()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Image")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetImage(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Image")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetSecurityContext()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("SecurityContext")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetSecurityContext(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("SecurityContext")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetResources()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Resources")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetResources(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Resources")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetEnv() {

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

	for _, v := range m.GetPorts() {

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

	return hasher.Sum64(), nil
}

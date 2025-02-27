package wellknown

// DefaultXdsService is the default name of the Kubernetes Service that serves xDS config.
// This value should stay in sync with:
// - the default value of `XdsServiceName` in internal/kgateway/extensions2/settings/settings.go
// - the default Service name in install/helm/kgateway/templates/service.yaml
const DefaultXdsService = "kgateway"

// DefaultXdsPort is the default xDS port. This value should stay in sync with:
// - the default value of `XdsServicePort` in internal/kgateway/extensions2/settings/settings.go
// - the `controller.service.ports.grpc` value in install/helm/kgateway/values.yaml
var DefaultXdsPort uint32 = 9977

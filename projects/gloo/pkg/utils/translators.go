package utils

import (
	"fmt"
	"strings"

	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

const (
	// label key applied to Proxies generated by the Gloo Edge translator
	ProxyTypeKey = "created_by"
	// label value applied to Proxies generated by the Gloo Edge translator
	GlooEdgeProxyValue = "gloo-gateway"
	// label value applied to Proxies generated by the Gloo Gateway translator
	GlooGatewayProxyValue = "gloo-kube-gateway-api"
	// label value applied to Proxies generated by the Gloo Knative translator
	KnativeProxyValue = "gloo-knative"
	// label value applied to Proxies generated by the Gloo Ingress translator
	IngressProxyValue = "gloo-ingress"

	// Gloo Gateway proxies can live in different namespaces from writeNamespace
	NamespaceLabel = "proxy_namespace"
)

func GetTranslatorSelectorExpression(translators ...string) string {
	return fmt.Sprintf("%s in (%s)", ProxyTypeKey, strings.Join(translators, ", "))
}

func GetTranslatorValue(meta *core.Metadata) string {
	return meta.GetLabels()[ProxyTypeKey]
}

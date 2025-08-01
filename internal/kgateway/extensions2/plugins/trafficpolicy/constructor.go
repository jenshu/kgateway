package trafficpolicy

import (
	"context"
	"fmt"

	"istio.io/istio/pkg/kube/krt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/kgateway-dev/kgateway/v2/api/v1alpha1"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/extensions2/common"
	"github.com/kgateway-dev/kgateway/v2/internal/kgateway/ir"
)

// FetchGatewayExtensionFunc defines the signature for fetching gateway extensions
type FetchGatewayExtensionFunc func(krtctx krt.HandlerContext, extensionRef *corev1.LocalObjectReference, ns string) (*TrafficPolicyGatewayExtensionIR, error)

type TrafficPolicyConstructor struct {
	commoncol         *common.CommonCollections
	gatewayExtensions krt.Collection[TrafficPolicyGatewayExtensionIR]
	extBuilder        func(krtctx krt.HandlerContext, gExt ir.GatewayExtension) *TrafficPolicyGatewayExtensionIR
}

func NewTrafficPolicyConstructor(
	ctx context.Context,
	commoncol *common.CommonCollections,
) *TrafficPolicyConstructor {
	extBuilder := TranslateGatewayExtensionBuilder(commoncol)
	defaultExtBuilder := func(krtctx krt.HandlerContext, gExt ir.GatewayExtension) *TrafficPolicyGatewayExtensionIR {
		return extBuilder(krtctx, gExt)
	}
	gatewayExtensions := krt.NewCollection(commoncol.GatewayExtensions, defaultExtBuilder)
	return &TrafficPolicyConstructor{
		commoncol:         commoncol,
		gatewayExtensions: gatewayExtensions,
		extBuilder:        extBuilder,
	}
}

func (c *TrafficPolicyConstructor) ConstructIR(
	krtctx krt.HandlerContext,
	policyCR *v1alpha1.TrafficPolicy,
) (*TrafficPolicy, []error) {
	policyIr := TrafficPolicy{
		ct: policyCR.CreationTimestamp.Time,
	}
	outSpec := trafficPolicySpecIr{}

	var errors []error
	// Construct AI specific IR
	if err := constructAI(krtctx, policyCR, c.commoncol.Secrets, &outSpec); err != nil {
		errors = append(errors, err)
	}
	// Construct transformation specific IR
	if err := constructTransformation(policyCR, &outSpec); err != nil {
		errors = append(errors, err)
	}
	// Construct rustformation specific IR
	if err := constructRustformation(policyCR, &outSpec); err != nil {
		errors = append(errors, err)
	}
	// Construct extproc specific IR
	if err := constructExtProc(krtctx, policyCR, c.FetchGatewayExtension, &outSpec); err != nil {
		errors = append(errors, err)
	}
	// Construct extauth specific IR
	if err := constructExtAuth(krtctx, policyCR, c.FetchGatewayExtension, &outSpec); err != nil {
		errors = append(errors, err)
	}
	// Construct local rate limit specific IR
	if err := constructLocalRateLimit(policyCR, &outSpec); err != nil {
		errors = append(errors, err)
	}
	// Construct global rate limit specific IR
	if err := constructGlobalRateLimit(krtctx, policyCR, c.FetchGatewayExtension, &outSpec); err != nil {
		errors = append(errors, err)
	}
	// Construct cors specific IR
	if err := constructCORS(policyCR, &outSpec); err != nil {
		errors = append(errors, err)
	}
	// Construct csrf specific IR
	if err := constructCSRF(policyCR.Spec, &outSpec); err != nil {
		errors = append(errors, err)
	}

	// Construct hash policy specific IR
	constructHashPolicy(policyCR.Spec, &outSpec)
	// Construct auto host rewrite specific IR
	constructAutoHostRewrite(policyCR.Spec, &outSpec)
	// Construct buffer specific IR
	constructBuffer(policyCR.Spec, &outSpec)

	for _, err := range errors {
		logger.Error("error translating gateway extension", "namespace", policyCR.GetNamespace(), "name", policyCR.GetName(), "error", err)
	}
	policyIr.spec = outSpec

	return &policyIr, errors
}

func (c *TrafficPolicyConstructor) FetchGatewayExtension(krtctx krt.HandlerContext, extensionRef *corev1.LocalObjectReference, ns string) (*TrafficPolicyGatewayExtensionIR, error) {
	var gatewayExtension *TrafficPolicyGatewayExtensionIR
	if extensionRef != nil {
		gwExtName := types.NamespacedName{Name: extensionRef.Name, Namespace: ns}
		gatewayExtension = krt.FetchOne(krtctx, c.gatewayExtensions, krt.FilterObjectName(gwExtName))
	}
	if gatewayExtension == nil {
		return nil, fmt.Errorf("extension not found")
	}
	if gatewayExtension.Err != nil {
		return gatewayExtension, gatewayExtension.Err
	}
	return gatewayExtension, nil
}

func (c *TrafficPolicyConstructor) HasSynced() bool {
	return c.gatewayExtensions.HasSynced()
}

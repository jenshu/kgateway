// Code generated by skv2. DO NOT EDIT.

package gateway_resource_handler

import (
	"context"
	"sort"

	"github.com/ghodss/yaml"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"

	"github.com/solo-io/go-utils/contextutils"
	skv2v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	gateway_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gateway.solo.io/v1"
	rpc_edge_v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1"
	"github.com/solo-io/solo-projects/projects/apiserver/server/apiserverutils"
	fedv1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewFedGatewayResourceHandler(
	instanceClient fedv1.GlooInstanceClient,
	mcGatewayCRDClientset gateway_solo_io_v1.MulticlusterClientset,

) rpc_edge_v1.GatewayResourceApiServer {
	return &fedGatewayResourceHandler{
		instanceClient:        instanceClient,
		mcGatewayCRDClientset: mcGatewayCRDClientset,
	}
}

type fedGatewayResourceHandler struct {
	instanceClient        fedv1.GlooInstanceClient
	mcGatewayCRDClientset gateway_solo_io_v1.MulticlusterClientset
}

func (k *fedGatewayResourceHandler) ListGateways(ctx context.Context, request *rpc_edge_v1.ListGatewaysRequest) (*rpc_edge_v1.ListGatewaysResponse, error) {

	var rpcGateways []*rpc_edge_v1.Gateway
	if request.GetGlooInstanceRef() == nil || request.GetGlooInstanceRef().GetName() == "" || request.GetGlooInstanceRef().GetNamespace() == "" {
		// List gateways across all gloo edge instances
		instanceList, err := k.instanceClient.ListGlooInstance(ctx)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, instance := range instanceList.Items {
			rpcGatewayList, err := k.listGatewaysForGlooInstance(ctx, &instance)
			if err != nil {
				wrapped := eris.Wrapf(err, "Failed to list gateways for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
				contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
				return nil, wrapped
			}
			rpcGateways = append(rpcGateways, rpcGatewayList...)
		}
	} else {
		// List gateways for a specific gloo edge instance
		instance, err := k.instanceClient.GetGlooInstance(ctx, types.NamespacedName{
			Name:      request.GetGlooInstanceRef().GetName(),
			Namespace: request.GetGlooInstanceRef().GetNamespace(),
		})
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		rpcGateways, err = k.listGatewaysForGlooInstance(ctx, instance)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gateways for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
	}

	return &rpc_edge_v1.ListGatewaysResponse{
		Gateways: rpcGateways,
	}, nil
}

func (k *fedGatewayResourceHandler) listGatewaysForGlooInstance(ctx context.Context, instance *fedv1.GlooInstance) ([]*rpc_edge_v1.Gateway, error) {

	gatewayCRDClientset, err := k.mcGatewayCRDClientset.Cluster(instance.Spec.GetCluster())
	if err != nil {
		return nil, err
	}
	gatewayClient := gatewayCRDClientset.Gateways()

	var gatewayGatewayList []*gateway_solo_io_v1.Gateway
	watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
	if len(watchedNamespaces) != 0 {
		for _, ns := range watchedNamespaces {
			list, err := gatewayClient.ListGateway(ctx, client.InNamespace(ns))
			if err != nil {
				return nil, err
			}
			for i, _ := range list.Items {
				gatewayGatewayList = append(gatewayGatewayList, &list.Items[i])
			}
		}
	} else {
		list, err := gatewayClient.ListGateway(ctx)
		if err != nil {
			return nil, err
		}
		for i, _ := range list.Items {
			gatewayGatewayList = append(gatewayGatewayList, &list.Items[i])
		}
	}
	sort.Slice(gatewayGatewayList, func(i, j int) bool {
		x := gatewayGatewayList[i]
		y := gatewayGatewayList[j]
		return x.GetNamespace()+x.GetName() < y.GetNamespace()+y.GetName()
	})

	var rpcGateways []*rpc_edge_v1.Gateway
	for _, gateway := range gatewayGatewayList {
		rpcGateways = append(rpcGateways, BuildRpcGateway(gateway, &skv2v1.ObjectRef{
			Name:      instance.GetName(),
			Namespace: instance.GetNamespace(),
		}, instance.Spec.GetCluster()))
	}
	return rpcGateways, nil
}

func BuildRpcGateway(gateway *gateway_solo_io_v1.Gateway, glooInstance *skv2v1.ObjectRef, cluster string) *rpc_edge_v1.Gateway {
	m := &rpc_edge_v1.Gateway{
		Metadata:     apiserverutils.ToMetadata(gateway.ObjectMeta),
		GlooInstance: glooInstance,
		Spec:         &gateway.Spec,
		Status:       &gateway.Status,
	}
	m.Metadata.ClusterName = cluster
	return m
}

func (k *fedGatewayResourceHandler) GetGatewayYaml(ctx context.Context, request *rpc_edge_v1.GetGatewayYamlRequest) (*rpc_edge_v1.GetGatewayYamlResponse, error) {
	gatewayClientSet, err := k.mcGatewayCRDClientset.Cluster(request.GetGatewayRef().GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get gateway client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	gateway, err := gatewayClientSet.Gateways().GetGateway(ctx, client.ObjectKey{
		Namespace: request.GetGatewayRef().GetNamespace(),
		Name:      request.GetGatewayRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get gateway")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(gateway)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_edge_v1.GetGatewayYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

func (k *fedGatewayResourceHandler) ListVirtualServices(ctx context.Context, request *rpc_edge_v1.ListVirtualServicesRequest) (*rpc_edge_v1.ListVirtualServicesResponse, error) {

	var rpcVirtualServices []*rpc_edge_v1.VirtualService
	if request.GetGlooInstanceRef() == nil || request.GetGlooInstanceRef().GetName() == "" || request.GetGlooInstanceRef().GetNamespace() == "" {
		// List virtualServices across all gloo edge instances
		instanceList, err := k.instanceClient.ListGlooInstance(ctx)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, instance := range instanceList.Items {
			rpcVirtualServiceList, err := k.listVirtualServicesForGlooInstance(ctx, &instance)
			if err != nil {
				wrapped := eris.Wrapf(err, "Failed to list virtualServices for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
				contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
				return nil, wrapped
			}
			rpcVirtualServices = append(rpcVirtualServices, rpcVirtualServiceList...)
		}
	} else {
		// List virtualServices for a specific gloo edge instance
		instance, err := k.instanceClient.GetGlooInstance(ctx, types.NamespacedName{
			Name:      request.GetGlooInstanceRef().GetName(),
			Namespace: request.GetGlooInstanceRef().GetNamespace(),
		})
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		rpcVirtualServices, err = k.listVirtualServicesForGlooInstance(ctx, instance)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list virtualServices for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
	}

	return &rpc_edge_v1.ListVirtualServicesResponse{
		VirtualServices: rpcVirtualServices,
	}, nil
}

func (k *fedGatewayResourceHandler) listVirtualServicesForGlooInstance(ctx context.Context, instance *fedv1.GlooInstance) ([]*rpc_edge_v1.VirtualService, error) {

	gatewayCRDClientset, err := k.mcGatewayCRDClientset.Cluster(instance.Spec.GetCluster())
	if err != nil {
		return nil, err
	}
	virtualServiceClient := gatewayCRDClientset.VirtualServices()

	var gatewayVirtualServiceList []*gateway_solo_io_v1.VirtualService
	watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
	if len(watchedNamespaces) != 0 {
		for _, ns := range watchedNamespaces {
			list, err := virtualServiceClient.ListVirtualService(ctx, client.InNamespace(ns))
			if err != nil {
				return nil, err
			}
			for i, _ := range list.Items {
				gatewayVirtualServiceList = append(gatewayVirtualServiceList, &list.Items[i])
			}
		}
	} else {
		list, err := virtualServiceClient.ListVirtualService(ctx)
		if err != nil {
			return nil, err
		}
		for i, _ := range list.Items {
			gatewayVirtualServiceList = append(gatewayVirtualServiceList, &list.Items[i])
		}
	}
	sort.Slice(gatewayVirtualServiceList, func(i, j int) bool {
		x := gatewayVirtualServiceList[i]
		y := gatewayVirtualServiceList[j]
		return x.GetNamespace()+x.GetName() < y.GetNamespace()+y.GetName()
	})

	var rpcVirtualServices []*rpc_edge_v1.VirtualService
	for _, virtualService := range gatewayVirtualServiceList {
		rpcVirtualServices = append(rpcVirtualServices, BuildRpcVirtualService(virtualService, &skv2v1.ObjectRef{
			Name:      instance.GetName(),
			Namespace: instance.GetNamespace(),
		}, instance.Spec.GetCluster()))
	}
	return rpcVirtualServices, nil
}

func BuildRpcVirtualService(virtualService *gateway_solo_io_v1.VirtualService, glooInstance *skv2v1.ObjectRef, cluster string) *rpc_edge_v1.VirtualService {
	m := &rpc_edge_v1.VirtualService{
		Metadata:     apiserverutils.ToMetadata(virtualService.ObjectMeta),
		GlooInstance: glooInstance,
		Spec:         &virtualService.Spec,
		Status:       &virtualService.Status,
	}
	m.Metadata.ClusterName = cluster
	return m
}

func (k *fedGatewayResourceHandler) GetVirtualServiceYaml(ctx context.Context, request *rpc_edge_v1.GetVirtualServiceYamlRequest) (*rpc_edge_v1.GetVirtualServiceYamlResponse, error) {
	gatewayClientSet, err := k.mcGatewayCRDClientset.Cluster(request.GetVirtualServiceRef().GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get gateway client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	virtualService, err := gatewayClientSet.VirtualServices().GetVirtualService(ctx, client.ObjectKey{
		Namespace: request.GetVirtualServiceRef().GetNamespace(),
		Name:      request.GetVirtualServiceRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get virtualService")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(virtualService)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_edge_v1.GetVirtualServiceYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

func (k *fedGatewayResourceHandler) ListRouteTables(ctx context.Context, request *rpc_edge_v1.ListRouteTablesRequest) (*rpc_edge_v1.ListRouteTablesResponse, error) {

	var rpcRouteTables []*rpc_edge_v1.RouteTable
	if request.GetGlooInstanceRef() == nil || request.GetGlooInstanceRef().GetName() == "" || request.GetGlooInstanceRef().GetNamespace() == "" {
		// List routeTables across all gloo edge instances
		instanceList, err := k.instanceClient.ListGlooInstance(ctx)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, instance := range instanceList.Items {
			rpcRouteTableList, err := k.listRouteTablesForGlooInstance(ctx, &instance)
			if err != nil {
				wrapped := eris.Wrapf(err, "Failed to list routeTables for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
				contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
				return nil, wrapped
			}
			rpcRouteTables = append(rpcRouteTables, rpcRouteTableList...)
		}
	} else {
		// List routeTables for a specific gloo edge instance
		instance, err := k.instanceClient.GetGlooInstance(ctx, types.NamespacedName{
			Name:      request.GetGlooInstanceRef().GetName(),
			Namespace: request.GetGlooInstanceRef().GetNamespace(),
		})
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		rpcRouteTables, err = k.listRouteTablesForGlooInstance(ctx, instance)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list routeTables for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
	}

	return &rpc_edge_v1.ListRouteTablesResponse{
		RouteTables: rpcRouteTables,
	}, nil
}

func (k *fedGatewayResourceHandler) listRouteTablesForGlooInstance(ctx context.Context, instance *fedv1.GlooInstance) ([]*rpc_edge_v1.RouteTable, error) {

	gatewayCRDClientset, err := k.mcGatewayCRDClientset.Cluster(instance.Spec.GetCluster())
	if err != nil {
		return nil, err
	}
	routeTableClient := gatewayCRDClientset.RouteTables()

	var gatewayRouteTableList []*gateway_solo_io_v1.RouteTable
	watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
	if len(watchedNamespaces) != 0 {
		for _, ns := range watchedNamespaces {
			list, err := routeTableClient.ListRouteTable(ctx, client.InNamespace(ns))
			if err != nil {
				return nil, err
			}
			for i, _ := range list.Items {
				gatewayRouteTableList = append(gatewayRouteTableList, &list.Items[i])
			}
		}
	} else {
		list, err := routeTableClient.ListRouteTable(ctx)
		if err != nil {
			return nil, err
		}
		for i, _ := range list.Items {
			gatewayRouteTableList = append(gatewayRouteTableList, &list.Items[i])
		}
	}
	sort.Slice(gatewayRouteTableList, func(i, j int) bool {
		x := gatewayRouteTableList[i]
		y := gatewayRouteTableList[j]
		return x.GetNamespace()+x.GetName() < y.GetNamespace()+y.GetName()
	})

	var rpcRouteTables []*rpc_edge_v1.RouteTable
	for _, routeTable := range gatewayRouteTableList {
		rpcRouteTables = append(rpcRouteTables, BuildRpcRouteTable(routeTable, &skv2v1.ObjectRef{
			Name:      instance.GetName(),
			Namespace: instance.GetNamespace(),
		}, instance.Spec.GetCluster()))
	}
	return rpcRouteTables, nil
}

func BuildRpcRouteTable(routeTable *gateway_solo_io_v1.RouteTable, glooInstance *skv2v1.ObjectRef, cluster string) *rpc_edge_v1.RouteTable {
	m := &rpc_edge_v1.RouteTable{
		Metadata:     apiserverutils.ToMetadata(routeTable.ObjectMeta),
		GlooInstance: glooInstance,
		Spec:         &routeTable.Spec,
		Status:       &routeTable.Status,
	}
	m.Metadata.ClusterName = cluster
	return m
}

func (k *fedGatewayResourceHandler) GetRouteTableYaml(ctx context.Context, request *rpc_edge_v1.GetRouteTableYamlRequest) (*rpc_edge_v1.GetRouteTableYamlResponse, error) {
	gatewayClientSet, err := k.mcGatewayCRDClientset.Cluster(request.GetRouteTableRef().GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get gateway client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	routeTable, err := gatewayClientSet.RouteTables().GetRouteTable(ctx, client.ObjectKey{
		Namespace: request.GetRouteTableRef().GetNamespace(),
		Name:      request.GetRouteTableRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get routeTable")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(routeTable)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_edge_v1.GetRouteTableYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

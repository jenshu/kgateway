// Code generated by skv2. DO NOT EDIT.

package gloo_resource_handler

import (
	"context"
	"sort"

	"github.com/ghodss/yaml"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"

	"github.com/solo-io/go-utils/contextutils"
	skv2v1 "github.com/solo-io/skv2/pkg/api/core.skv2.solo.io/v1"
	gloo_solo_io_v1 "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1"
	rpc_edge_v1 "github.com/solo-io/solo-projects/projects/apiserver/pkg/api/rpc.edge.gloo/v1"
	"github.com/solo-io/solo-projects/projects/apiserver/server/apiserverutils"
	fedv1 "github.com/solo-io/solo-projects/projects/gloo-fed/pkg/api/fed.solo.io/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewGlooResourceHandler(
	instanceClient fedv1.GlooInstanceClient,
	mcGlooCRDClientset gloo_solo_io_v1.MulticlusterClientset,

) rpc_edge_v1.GlooResourceApiServer {
	return &glooResourceHandler{
		instanceClient:     instanceClient,
		mcGlooCRDClientset: mcGlooCRDClientset,
	}
}

type glooResourceHandler struct {
	instanceClient     fedv1.GlooInstanceClient
	mcGlooCRDClientset gloo_solo_io_v1.MulticlusterClientset
}

func (k *glooResourceHandler) ListUpstreams(ctx context.Context, request *rpc_edge_v1.ListUpstreamsRequest) (*rpc_edge_v1.ListUpstreamsResponse, error) {

	var rpcUpstreams []*rpc_edge_v1.Upstream
	if request.GetGlooInstanceRef() == nil || request.GetGlooInstanceRef().GetName() == "" || request.GetGlooInstanceRef().GetNamespace() == "" {
		// List upstreams across all gloo edge instances
		instanceList, err := k.instanceClient.ListGlooInstance(ctx)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, instance := range instanceList.Items {
			rpcUpstreamList, err := k.listUpstreamsForGlooInstance(ctx, &instance)
			if err != nil {
				wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
				contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
				return nil, wrapped
			}
			rpcUpstreams = append(rpcUpstreams, rpcUpstreamList...)
		}
	} else {
		// List upstreams for a specific gloo edge instance
		instance, err := k.instanceClient.GetGlooInstance(ctx, types.NamespacedName{
			Name:      request.GetGlooInstanceRef().GetName(),
			Namespace: request.GetGlooInstanceRef().GetNamespace(),
		})
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		rpcUpstreams, err = k.listUpstreamsForGlooInstance(ctx, instance)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list upstreams for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
	}

	return &rpc_edge_v1.ListUpstreamsResponse{
		Upstreams: rpcUpstreams,
	}, nil
}

func (k *glooResourceHandler) listUpstreamsForGlooInstance(ctx context.Context, instance *fedv1.GlooInstance) ([]*rpc_edge_v1.Upstream, error) {

	glooCRDClientset, err := k.mcGlooCRDClientset.Cluster(instance.Spec.GetCluster())
	if err != nil {
		return nil, err
	}
	upstreamClient := glooCRDClientset.Upstreams()

	var glooUpstreamList []*gloo_solo_io_v1.Upstream
	watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
	if len(watchedNamespaces) != 0 {
		for _, ns := range watchedNamespaces {
			list, err := upstreamClient.ListUpstream(ctx, client.InNamespace(ns))
			if err != nil {
				return nil, err
			}
			for i, _ := range list.Items {
				glooUpstreamList = append(glooUpstreamList, &list.Items[i])
			}
		}
	} else {
		list, err := upstreamClient.ListUpstream(ctx)
		if err != nil {
			return nil, err
		}
		for i, _ := range list.Items {
			glooUpstreamList = append(glooUpstreamList, &list.Items[i])
		}
	}
	sort.Slice(glooUpstreamList, func(i, j int) bool {
		x := glooUpstreamList[i]
		y := glooUpstreamList[j]
		return x.GetNamespace()+x.GetName() < y.GetNamespace()+y.GetName()
	})

	var rpcUpstreams []*rpc_edge_v1.Upstream
	for _, upstream := range glooUpstreamList {
		rpcUpstreams = append(rpcUpstreams, BuildRpcUpstream(upstream, &skv2v1.ObjectRef{
			Name:      instance.GetName(),
			Namespace: instance.GetNamespace(),
		}, instance.Spec.GetCluster()))
	}
	return rpcUpstreams, nil
}

func BuildRpcUpstream(upstream *gloo_solo_io_v1.Upstream, glooInstance *skv2v1.ObjectRef, cluster string) *rpc_edge_v1.Upstream {
	m := &rpc_edge_v1.Upstream{
		Metadata:     apiserverutils.ToMetadata(upstream.ObjectMeta),
		GlooInstance: glooInstance,
		Spec:         &upstream.Spec,
		Status:       &upstream.Status,
	}
	m.Metadata.ClusterName = cluster
	return m
}

func (k *glooResourceHandler) GetUpstreamYaml(ctx context.Context, request *rpc_edge_v1.GetUpstreamYamlRequest) (*rpc_edge_v1.GetUpstreamYamlResponse, error) {
	glooClientSet, err := k.mcGlooCRDClientset.Cluster(request.GetUpstreamRef().GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get gloo client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	upstream, err := glooClientSet.Upstreams().GetUpstream(ctx, client.ObjectKey{
		Namespace: request.GetUpstreamRef().GetNamespace(),
		Name:      request.GetUpstreamRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get upstream")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(upstream)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_edge_v1.GetUpstreamYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

func (k *glooResourceHandler) ListUpstreamGroups(ctx context.Context, request *rpc_edge_v1.ListUpstreamGroupsRequest) (*rpc_edge_v1.ListUpstreamGroupsResponse, error) {

	var rpcUpstreamGroups []*rpc_edge_v1.UpstreamGroup
	if request.GetGlooInstanceRef() == nil || request.GetGlooInstanceRef().GetName() == "" || request.GetGlooInstanceRef().GetNamespace() == "" {
		// List upstreamGroups across all gloo edge instances
		instanceList, err := k.instanceClient.ListGlooInstance(ctx)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, instance := range instanceList.Items {
			rpcUpstreamGroupList, err := k.listUpstreamGroupsForGlooInstance(ctx, &instance)
			if err != nil {
				wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
				contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
				return nil, wrapped
			}
			rpcUpstreamGroups = append(rpcUpstreamGroups, rpcUpstreamGroupList...)
		}
	} else {
		// List upstreamGroups for a specific gloo edge instance
		instance, err := k.instanceClient.GetGlooInstance(ctx, types.NamespacedName{
			Name:      request.GetGlooInstanceRef().GetName(),
			Namespace: request.GetGlooInstanceRef().GetNamespace(),
		})
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		rpcUpstreamGroups, err = k.listUpstreamGroupsForGlooInstance(ctx, instance)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list upstreamGroups for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
	}

	return &rpc_edge_v1.ListUpstreamGroupsResponse{
		UpstreamGroups: rpcUpstreamGroups,
	}, nil
}

func (k *glooResourceHandler) listUpstreamGroupsForGlooInstance(ctx context.Context, instance *fedv1.GlooInstance) ([]*rpc_edge_v1.UpstreamGroup, error) {

	glooCRDClientset, err := k.mcGlooCRDClientset.Cluster(instance.Spec.GetCluster())
	if err != nil {
		return nil, err
	}
	upstreamGroupClient := glooCRDClientset.UpstreamGroups()

	var glooUpstreamGroupList []*gloo_solo_io_v1.UpstreamGroup
	watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
	if len(watchedNamespaces) != 0 {
		for _, ns := range watchedNamespaces {
			list, err := upstreamGroupClient.ListUpstreamGroup(ctx, client.InNamespace(ns))
			if err != nil {
				return nil, err
			}
			for i, _ := range list.Items {
				glooUpstreamGroupList = append(glooUpstreamGroupList, &list.Items[i])
			}
		}
	} else {
		list, err := upstreamGroupClient.ListUpstreamGroup(ctx)
		if err != nil {
			return nil, err
		}
		for i, _ := range list.Items {
			glooUpstreamGroupList = append(glooUpstreamGroupList, &list.Items[i])
		}
	}
	sort.Slice(glooUpstreamGroupList, func(i, j int) bool {
		x := glooUpstreamGroupList[i]
		y := glooUpstreamGroupList[j]
		return x.GetNamespace()+x.GetName() < y.GetNamespace()+y.GetName()
	})

	var rpcUpstreamGroups []*rpc_edge_v1.UpstreamGroup
	for _, upstreamGroup := range glooUpstreamGroupList {
		rpcUpstreamGroups = append(rpcUpstreamGroups, BuildRpcUpstreamGroup(upstreamGroup, &skv2v1.ObjectRef{
			Name:      instance.GetName(),
			Namespace: instance.GetNamespace(),
		}, instance.Spec.GetCluster()))
	}
	return rpcUpstreamGroups, nil
}

func BuildRpcUpstreamGroup(upstreamGroup *gloo_solo_io_v1.UpstreamGroup, glooInstance *skv2v1.ObjectRef, cluster string) *rpc_edge_v1.UpstreamGroup {
	m := &rpc_edge_v1.UpstreamGroup{
		Metadata:     apiserverutils.ToMetadata(upstreamGroup.ObjectMeta),
		GlooInstance: glooInstance,
		Spec:         &upstreamGroup.Spec,
		Status:       &upstreamGroup.Status,
	}
	m.Metadata.ClusterName = cluster
	return m
}

func (k *glooResourceHandler) GetUpstreamGroupYaml(ctx context.Context, request *rpc_edge_v1.GetUpstreamGroupYamlRequest) (*rpc_edge_v1.GetUpstreamGroupYamlResponse, error) {
	glooClientSet, err := k.mcGlooCRDClientset.Cluster(request.GetUpstreamGroupRef().GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get gloo client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	upstreamGroup, err := glooClientSet.UpstreamGroups().GetUpstreamGroup(ctx, client.ObjectKey{
		Namespace: request.GetUpstreamGroupRef().GetNamespace(),
		Name:      request.GetUpstreamGroupRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get upstreamGroup")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(upstreamGroup)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_edge_v1.GetUpstreamGroupYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

func (k *glooResourceHandler) ListSettings(ctx context.Context, request *rpc_edge_v1.ListSettingsRequest) (*rpc_edge_v1.ListSettingsResponse, error) {

	var rpcSettings []*rpc_edge_v1.Settings
	if request.GetGlooInstanceRef() == nil || request.GetGlooInstanceRef().GetName() == "" || request.GetGlooInstanceRef().GetNamespace() == "" {
		// List settings across all gloo edge instances
		instanceList, err := k.instanceClient.ListGlooInstance(ctx)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, instance := range instanceList.Items {
			rpcSettingsList, err := k.listSettingsForGlooInstance(ctx, &instance)
			if err != nil {
				wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
				contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
				return nil, wrapped
			}
			rpcSettings = append(rpcSettings, rpcSettingsList...)
		}
	} else {
		// List settings for a specific gloo edge instance
		instance, err := k.instanceClient.GetGlooInstance(ctx, types.NamespacedName{
			Name:      request.GetGlooInstanceRef().GetName(),
			Namespace: request.GetGlooInstanceRef().GetNamespace(),
		})
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		rpcSettings, err = k.listSettingsForGlooInstance(ctx, instance)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list settings for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
	}

	return &rpc_edge_v1.ListSettingsResponse{
		Settings: rpcSettings,
	}, nil
}

func (k *glooResourceHandler) listSettingsForGlooInstance(ctx context.Context, instance *fedv1.GlooInstance) ([]*rpc_edge_v1.Settings, error) {

	glooCRDClientset, err := k.mcGlooCRDClientset.Cluster(instance.Spec.GetCluster())
	if err != nil {
		return nil, err
	}
	settingsClient := glooCRDClientset.Settings()

	var glooSettingsList []*gloo_solo_io_v1.Settings
	watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
	if len(watchedNamespaces) != 0 {
		for _, ns := range watchedNamespaces {
			list, err := settingsClient.ListSettings(ctx, client.InNamespace(ns))
			if err != nil {
				return nil, err
			}
			for i, _ := range list.Items {
				glooSettingsList = append(glooSettingsList, &list.Items[i])
			}
		}
	} else {
		list, err := settingsClient.ListSettings(ctx)
		if err != nil {
			return nil, err
		}
		for i, _ := range list.Items {
			glooSettingsList = append(glooSettingsList, &list.Items[i])
		}
	}
	sort.Slice(glooSettingsList, func(i, j int) bool {
		x := glooSettingsList[i]
		y := glooSettingsList[j]
		return x.GetNamespace()+x.GetName() < y.GetNamespace()+y.GetName()
	})

	var rpcSettings []*rpc_edge_v1.Settings
	for _, settings := range glooSettingsList {
		rpcSettings = append(rpcSettings, BuildRpcSettings(settings, &skv2v1.ObjectRef{
			Name:      instance.GetName(),
			Namespace: instance.GetNamespace(),
		}, instance.Spec.GetCluster()))
	}
	return rpcSettings, nil
}

func BuildRpcSettings(settings *gloo_solo_io_v1.Settings, glooInstance *skv2v1.ObjectRef, cluster string) *rpc_edge_v1.Settings {
	m := &rpc_edge_v1.Settings{
		Metadata:     apiserverutils.ToMetadata(settings.ObjectMeta),
		GlooInstance: glooInstance,
		Spec:         &settings.Spec,
		Status:       &settings.Status,
	}
	m.Metadata.ClusterName = cluster
	return m
}

func (k *glooResourceHandler) GetSettingsYaml(ctx context.Context, request *rpc_edge_v1.GetSettingsYamlRequest) (*rpc_edge_v1.GetSettingsYamlResponse, error) {
	glooClientSet, err := k.mcGlooCRDClientset.Cluster(request.GetSettingsRef().GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get gloo client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	settings, err := glooClientSet.Settings().GetSettings(ctx, client.ObjectKey{
		Namespace: request.GetSettingsRef().GetNamespace(),
		Name:      request.GetSettingsRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get settings")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(settings)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_edge_v1.GetSettingsYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

func (k *glooResourceHandler) ListProxies(ctx context.Context, request *rpc_edge_v1.ListProxiesRequest) (*rpc_edge_v1.ListProxiesResponse, error) {

	var rpcProxies []*rpc_edge_v1.Proxy
	if request.GetGlooInstanceRef() == nil || request.GetGlooInstanceRef().GetName() == "" || request.GetGlooInstanceRef().GetNamespace() == "" {
		// List proxies across all gloo edge instances
		instanceList, err := k.instanceClient.ListGlooInstance(ctx)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list gloo edge instances")
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		for _, instance := range instanceList.Items {
			rpcProxyList, err := k.listProxiesForGlooInstance(ctx, &instance)
			if err != nil {
				wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
				contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
				return nil, wrapped
			}
			rpcProxies = append(rpcProxies, rpcProxyList...)
		}
	} else {
		// List proxies for a specific gloo edge instance
		instance, err := k.instanceClient.GetGlooInstance(ctx, types.NamespacedName{
			Name:      request.GetGlooInstanceRef().GetName(),
			Namespace: request.GetGlooInstanceRef().GetNamespace(),
		})
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to get gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
		rpcProxies, err = k.listProxiesForGlooInstance(ctx, instance)
		if err != nil {
			wrapped := eris.Wrapf(err, "Failed to list proxies for gloo edge instance %s.%s", instance.GetNamespace(), instance.GetName())
			contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
			return nil, wrapped
		}
	}

	return &rpc_edge_v1.ListProxiesResponse{
		Proxies: rpcProxies,
	}, nil
}

func (k *glooResourceHandler) listProxiesForGlooInstance(ctx context.Context, instance *fedv1.GlooInstance) ([]*rpc_edge_v1.Proxy, error) {

	glooCRDClientset, err := k.mcGlooCRDClientset.Cluster(instance.Spec.GetCluster())
	if err != nil {
		return nil, err
	}
	proxyClient := glooCRDClientset.Proxies()

	var glooProxyList []*gloo_solo_io_v1.Proxy
	watchedNamespaces := instance.Spec.GetControlPlane().GetWatchedNamespaces()
	if len(watchedNamespaces) != 0 {
		for _, ns := range watchedNamespaces {
			list, err := proxyClient.ListProxy(ctx, client.InNamespace(ns))
			if err != nil {
				return nil, err
			}
			for i, _ := range list.Items {
				glooProxyList = append(glooProxyList, &list.Items[i])
			}
		}
	} else {
		list, err := proxyClient.ListProxy(ctx)
		if err != nil {
			return nil, err
		}
		for i, _ := range list.Items {
			glooProxyList = append(glooProxyList, &list.Items[i])
		}
	}
	sort.Slice(glooProxyList, func(i, j int) bool {
		x := glooProxyList[i]
		y := glooProxyList[j]
		return x.GetNamespace()+x.GetName() < y.GetNamespace()+y.GetName()
	})

	var rpcProxies []*rpc_edge_v1.Proxy
	for _, proxy := range glooProxyList {
		rpcProxies = append(rpcProxies, BuildRpcProxy(proxy, &skv2v1.ObjectRef{
			Name:      instance.GetName(),
			Namespace: instance.GetNamespace(),
		}, instance.Spec.GetCluster()))
	}
	return rpcProxies, nil
}

func BuildRpcProxy(proxy *gloo_solo_io_v1.Proxy, glooInstance *skv2v1.ObjectRef, cluster string) *rpc_edge_v1.Proxy {
	m := &rpc_edge_v1.Proxy{
		Metadata:     apiserverutils.ToMetadata(proxy.ObjectMeta),
		GlooInstance: glooInstance,
		Spec:         &proxy.Spec,
		Status:       &proxy.Status,
	}
	m.Metadata.ClusterName = cluster
	return m
}

func (k *glooResourceHandler) GetProxyYaml(ctx context.Context, request *rpc_edge_v1.GetProxyYamlRequest) (*rpc_edge_v1.GetProxyYamlResponse, error) {
	glooClientSet, err := k.mcGlooCRDClientset.Cluster(request.GetProxyRef().GetClusterName())
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get gloo client set")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	proxy, err := glooClientSet.Proxies().GetProxy(ctx, client.ObjectKey{
		Namespace: request.GetProxyRef().GetNamespace(),
		Name:      request.GetProxyRef().GetName(),
	})
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to get proxy")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	content, err := yaml.Marshal(proxy)
	if err != nil {
		wrapped := eris.Wrapf(err, "Failed to marshal kube resource into yaml")
		contextutils.LoggerFrom(ctx).Errorw(wrapped.Error(), zap.Error(err), zap.Any("request", request))
		return nil, wrapped
	}
	return &rpc_edge_v1.GetProxyYamlResponse{
		YamlData: &rpc_edge_v1.ResourceYaml{
			Yaml: string(content),
		},
	}, nil
}

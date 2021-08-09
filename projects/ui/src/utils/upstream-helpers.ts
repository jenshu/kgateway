import { Upstream } from 'proto/github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/gloo_resources_pb';
import { WeightedDestination } from 'proto/github.com/solo-io/solo-apis/api/gloo/gloo/v1/proxy_pb';

export const TYPE_AWS = 'AWS';
export const TYPE_AZURE = 'Azure';
export const TYPE_CONSUL = 'Consul';
export const TYPE_KUBE = 'Kubernetes';
export const TYPE_AWS_EC2 = 'AWS EC2';
export const TYPE_STATIC = 'Static';
export const TYPE_REST = 'Rest';
export const TYPE_GRPC = 'GRPC';
export const TYPE_OTHER = 'Other'; //just for consistency

export const getUpstreamType = (upstream?: Upstream.AsObject) => {
  if (upstream?.spec?.aws !== undefined) {
    return TYPE_AWS;
  }

  if (upstream?.spec?.azure !== undefined) {
    return TYPE_AZURE;
  }

  if (upstream?.spec?.consul !== undefined) {
    return TYPE_CONSUL;
  }

  if (upstream?.spec?.kube !== undefined) {
    return TYPE_KUBE;
  }

  if (upstream?.spec?.awsEc2 !== undefined) {
    return TYPE_AWS_EC2;
  }

  if (upstream?.spec?.pb_static !== undefined) {
    return TYPE_STATIC;
  }

  return TYPE_OTHER;
};

export const getWeightedDestinationType = (
  weightedDestination?: WeightedDestination.AsObject
) => {
  if (weightedDestination?.destination?.consul) {
    return TYPE_CONSUL;
  }
  if (weightedDestination?.destination?.kube) {
    return TYPE_KUBE;
  }

  /*const destSpec = weightedDestination?.destination?.destinationSpec;
  if (destSpec) {
    if (destSpec.aws) {
      return TYPE_AWS;
    }
    if (destSpec.azure) {
      return TYPE_AZURE;
    }
    if (destSpec.grpc) {
      return TYPE_GRPC;
    }
    if (destSpec.rest) {
      return TYPE_REST;
    }
  }*/
  if (weightedDestination?.destination?.upstream) {
    return TYPE_STATIC;
  }

  return TYPE_OTHER;
};

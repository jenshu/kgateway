/* eslint-disable */
// package: ratelimit.options.gloo.solo.io
// file: github.com/solo-io/solo-apis/api/gloo/gloo/v1/enterprise/options/ratelimit/ratelimit.proto

import * as jspb from "google-protobuf";
import * as github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb from "../../../../../../../../../../github.com/solo-io/solo-apis/api/rate-limiter/v1alpha1/ratelimit_pb";
import * as github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_local_ratelimit_local_ratelimit_pb from "../../../../../../../../../../github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/local_ratelimit/local_ratelimit_pb";
import * as github_com_solo_io_solo_kit_api_v1_ref_pb from "../../../../../../../../../../github.com/solo-io/solo-kit/api/v1/ref_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as google_protobuf_duration_pb from "google-protobuf/google/protobuf/duration_pb";
import * as extproto_ext_pb from "../../../../../../../../../../extproto/ext_pb";

export class IngressRateLimit extends jspb.Message {
  hasAuthorizedLimits(): boolean;
  clearAuthorizedLimits(): void;
  getAuthorizedLimits(): github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimit | undefined;
  setAuthorizedLimits(value?: github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimit): void;

  hasAnonymousLimits(): boolean;
  clearAnonymousLimits(): void;
  getAnonymousLimits(): github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimit | undefined;
  setAnonymousLimits(value?: github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimit): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): IngressRateLimit.AsObject;
  static toObject(includeInstance: boolean, msg: IngressRateLimit): IngressRateLimit.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: IngressRateLimit, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): IngressRateLimit;
  static deserializeBinaryFromReader(message: IngressRateLimit, reader: jspb.BinaryReader): IngressRateLimit;
}

export namespace IngressRateLimit {
  export type AsObject = {
    authorizedLimits?: github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimit.AsObject,
    anonymousLimits?: github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimit.AsObject,
  }
}

export class Settings extends jspb.Message {
  hasRatelimitServerRef(): boolean;
  clearRatelimitServerRef(): void;
  getRatelimitServerRef(): github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef | undefined;
  setRatelimitServerRef(value?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef): void;

  hasRequestTimeout(): boolean;
  clearRequestTimeout(): void;
  getRequestTimeout(): google_protobuf_duration_pb.Duration | undefined;
  setRequestTimeout(value?: google_protobuf_duration_pb.Duration): void;

  getDenyOnFail(): boolean;
  setDenyOnFail(value: boolean): void;

  getEnableXRatelimitHeaders(): boolean;
  setEnableXRatelimitHeaders(value: boolean): void;

  getRateLimitBeforeAuth(): boolean;
  setRateLimitBeforeAuth(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Settings.AsObject;
  static toObject(includeInstance: boolean, msg: Settings): Settings.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Settings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Settings;
  static deserializeBinaryFromReader(message: Settings, reader: jspb.BinaryReader): Settings;
}

export namespace Settings {
  export type AsObject = {
    ratelimitServerRef?: github_com_solo_io_solo_kit_api_v1_ref_pb.ResourceRef.AsObject,
    requestTimeout?: google_protobuf_duration_pb.Duration.AsObject,
    denyOnFail: boolean,
    enableXRatelimitHeaders: boolean,
    rateLimitBeforeAuth: boolean,
  }
}

export class ServiceSettings extends jspb.Message {
  clearDescriptorsList(): void;
  getDescriptorsList(): Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.Descriptor>;
  setDescriptorsList(value: Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.Descriptor>): void;
  addDescriptors(value?: github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.Descriptor, index?: number): github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.Descriptor;

  clearSetDescriptorsList(): void;
  getSetDescriptorsList(): Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.SetDescriptor>;
  setSetDescriptorsList(value: Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.SetDescriptor>): void;
  addSetDescriptors(value?: github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.SetDescriptor, index?: number): github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.SetDescriptor;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServiceSettings.AsObject;
  static toObject(includeInstance: boolean, msg: ServiceSettings): ServiceSettings.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServiceSettings, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServiceSettings;
  static deserializeBinaryFromReader(message: ServiceSettings, reader: jspb.BinaryReader): ServiceSettings;
}

export namespace ServiceSettings {
  export type AsObject = {
    descriptorsList: Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.Descriptor.AsObject>,
    setDescriptorsList: Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.SetDescriptor.AsObject>,
  }
}

export class RateLimitConfigRefs extends jspb.Message {
  clearRefsList(): void;
  getRefsList(): Array<RateLimitConfigRef>;
  setRefsList(value: Array<RateLimitConfigRef>): void;
  addRefs(value?: RateLimitConfigRef, index?: number): RateLimitConfigRef;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RateLimitConfigRefs.AsObject;
  static toObject(includeInstance: boolean, msg: RateLimitConfigRefs): RateLimitConfigRefs.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RateLimitConfigRefs, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RateLimitConfigRefs;
  static deserializeBinaryFromReader(message: RateLimitConfigRefs, reader: jspb.BinaryReader): RateLimitConfigRefs;
}

export namespace RateLimitConfigRefs {
  export type AsObject = {
    refsList: Array<RateLimitConfigRef.AsObject>,
  }
}

export class RateLimitConfigRef extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  getNamespace(): string;
  setNamespace(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RateLimitConfigRef.AsObject;
  static toObject(includeInstance: boolean, msg: RateLimitConfigRef): RateLimitConfigRef.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RateLimitConfigRef, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RateLimitConfigRef;
  static deserializeBinaryFromReader(message: RateLimitConfigRef, reader: jspb.BinaryReader): RateLimitConfigRef;
}

export namespace RateLimitConfigRef {
  export type AsObject = {
    name: string,
    namespace: string,
  }
}

export class RateLimitVhostExtension extends jspb.Message {
  clearRateLimitsList(): void;
  getRateLimitsList(): Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions>;
  setRateLimitsList(value: Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions>): void;
  addRateLimits(value?: github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions, index?: number): github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions;

  hasLocalRatelimit(): boolean;
  clearLocalRatelimit(): void;
  getLocalRatelimit(): github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_local_ratelimit_local_ratelimit_pb.TokenBucket | undefined;
  setLocalRatelimit(value?: github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_local_ratelimit_local_ratelimit_pb.TokenBucket): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RateLimitVhostExtension.AsObject;
  static toObject(includeInstance: boolean, msg: RateLimitVhostExtension): RateLimitVhostExtension.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RateLimitVhostExtension, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RateLimitVhostExtension;
  static deserializeBinaryFromReader(message: RateLimitVhostExtension, reader: jspb.BinaryReader): RateLimitVhostExtension;
}

export namespace RateLimitVhostExtension {
  export type AsObject = {
    rateLimitsList: Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions.AsObject>,
    localRatelimit?: github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_local_ratelimit_local_ratelimit_pb.TokenBucket.AsObject,
  }
}

export class RateLimitRouteExtension extends jspb.Message {
  getIncludeVhRateLimits(): boolean;
  setIncludeVhRateLimits(value: boolean): void;

  clearRateLimitsList(): void;
  getRateLimitsList(): Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions>;
  setRateLimitsList(value: Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions>): void;
  addRateLimits(value?: github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions, index?: number): github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions;

  hasLocalRatelimit(): boolean;
  clearLocalRatelimit(): void;
  getLocalRatelimit(): github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_local_ratelimit_local_ratelimit_pb.TokenBucket | undefined;
  setLocalRatelimit(value?: github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_local_ratelimit_local_ratelimit_pb.TokenBucket): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): RateLimitRouteExtension.AsObject;
  static toObject(includeInstance: boolean, msg: RateLimitRouteExtension): RateLimitRouteExtension.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: RateLimitRouteExtension, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): RateLimitRouteExtension;
  static deserializeBinaryFromReader(message: RateLimitRouteExtension, reader: jspb.BinaryReader): RateLimitRouteExtension;
}

export namespace RateLimitRouteExtension {
  export type AsObject = {
    includeVhRateLimits: boolean,
    rateLimitsList: Array<github_com_solo_io_solo_apis_api_rate_limiter_v1alpha1_ratelimit_pb.RateLimitActions.AsObject>,
    localRatelimit?: github_com_solo_io_solo_apis_api_gloo_gloo_v1_options_local_ratelimit_local_ratelimit_pb.TokenBucket.AsObject,
  }
}

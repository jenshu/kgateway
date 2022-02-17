/* eslint-disable */
// package: rpc.edge.gloo.solo.io
// file: github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/bootstrap.proto

import * as jspb from "google-protobuf";

export class GlooFedCheckRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GlooFedCheckRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GlooFedCheckRequest): GlooFedCheckRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GlooFedCheckRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GlooFedCheckRequest;
  static deserializeBinaryFromReader(message: GlooFedCheckRequest, reader: jspb.BinaryReader): GlooFedCheckRequest;
}

export namespace GlooFedCheckRequest {
  export type AsObject = {
  }
}

export class GlooFedCheckResponse extends jspb.Message {
  getEnabled(): boolean;
  setEnabled(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GlooFedCheckResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GlooFedCheckResponse): GlooFedCheckResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GlooFedCheckResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GlooFedCheckResponse;
  static deserializeBinaryFromReader(message: GlooFedCheckResponse, reader: jspb.BinaryReader): GlooFedCheckResponse;
}

export namespace GlooFedCheckResponse {
  export type AsObject = {
    enabled: boolean,
  }
}

export class GraphqlCheckRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GraphqlCheckRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GraphqlCheckRequest): GraphqlCheckRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GraphqlCheckRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GraphqlCheckRequest;
  static deserializeBinaryFromReader(message: GraphqlCheckRequest, reader: jspb.BinaryReader): GraphqlCheckRequest;
}

export namespace GraphqlCheckRequest {
  export type AsObject = {
  }
}

export class GraphqlCheckResponse extends jspb.Message {
  getEnabled(): boolean;
  setEnabled(value: boolean): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GraphqlCheckResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GraphqlCheckResponse): GraphqlCheckResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GraphqlCheckResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GraphqlCheckResponse;
  static deserializeBinaryFromReader(message: GraphqlCheckResponse, reader: jspb.BinaryReader): GraphqlCheckResponse;
}

export namespace GraphqlCheckResponse {
  export type AsObject = {
    enabled: boolean,
  }
}

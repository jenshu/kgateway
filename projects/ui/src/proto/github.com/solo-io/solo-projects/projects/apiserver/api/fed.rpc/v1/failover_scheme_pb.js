/* eslint-disable */
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var extproto_ext_pb = require('../../../../../../../../extproto/ext_pb.js');
var github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb = require('../../../../../../../../github.com/solo-io/solo-projects/projects/gloo-fed/api/fed/v1/failover_pb.js');
var github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb = require('../../../../../../../../github.com/solo-io/solo-projects/projects/apiserver/api/rpc.edge.gloo/v1/common_pb.js');
var github_com_solo$io_skv2_api_core_v1_core_pb = require('../../../../../../../../github.com/solo-io/skv2/api/core/v1/core_pb.js');
goog.exportSymbol('proto.fed.rpc.solo.io.FailoverScheme', null, global);
goog.exportSymbol('proto.fed.rpc.solo.io.GetFailoverSchemeRequest', null, global);
goog.exportSymbol('proto.fed.rpc.solo.io.GetFailoverSchemeResponse', null, global);
goog.exportSymbol('proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest', null, global);
goog.exportSymbol('proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fed.rpc.solo.io.FailoverScheme = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.rpc.solo.io.FailoverScheme, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.rpc.solo.io.FailoverScheme.displayName = 'proto.fed.rpc.solo.io.FailoverScheme';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fed.rpc.solo.io.FailoverScheme.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.rpc.solo.io.FailoverScheme.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.rpc.solo.io.FailoverScheme} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.FailoverScheme.toObject = function(includeInstance, msg) {
  var f, obj = {
    metadata: (f = msg.getMetadata()) && github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ObjectMeta.toObject(includeInstance, f),
    spec: (f = msg.getSpec()) && github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeSpec.toObject(includeInstance, f),
    status: (f = msg.getStatus()) && github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeStatus.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fed.rpc.solo.io.FailoverScheme}
 */
proto.fed.rpc.solo.io.FailoverScheme.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.rpc.solo.io.FailoverScheme;
  return proto.fed.rpc.solo.io.FailoverScheme.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.rpc.solo.io.FailoverScheme} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.rpc.solo.io.FailoverScheme}
 */
proto.fed.rpc.solo.io.FailoverScheme.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ObjectMeta;
      reader.readMessage(value,github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ObjectMeta.deserializeBinaryFromReader);
      msg.setMetadata(value);
      break;
    case 2:
      var value = new github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeSpec;
      reader.readMessage(value,github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeSpec.deserializeBinaryFromReader);
      msg.setSpec(value);
      break;
    case 3:
      var value = new github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeStatus;
      reader.readMessage(value,github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeStatus.deserializeBinaryFromReader);
      msg.setStatus(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fed.rpc.solo.io.FailoverScheme.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.rpc.solo.io.FailoverScheme.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.rpc.solo.io.FailoverScheme} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.FailoverScheme.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getMetadata();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ObjectMeta.serializeBinaryToWriter
    );
  }
  f = message.getSpec();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeSpec.serializeBinaryToWriter
    );
  }
  f = message.getStatus();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeStatus.serializeBinaryToWriter
    );
  }
};


/**
 * optional rpc.edge.gloo.solo.io.ObjectMeta metadata = 1;
 * @return {?proto.rpc.edge.gloo.solo.io.ObjectMeta}
 */
proto.fed.rpc.solo.io.FailoverScheme.prototype.getMetadata = function() {
  return /** @type{?proto.rpc.edge.gloo.solo.io.ObjectMeta} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ObjectMeta, 1));
};


/** @param {?proto.rpc.edge.gloo.solo.io.ObjectMeta|undefined} value */
proto.fed.rpc.solo.io.FailoverScheme.prototype.setMetadata = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.rpc.solo.io.FailoverScheme.prototype.clearMetadata = function() {
  this.setMetadata(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.rpc.solo.io.FailoverScheme.prototype.hasMetadata = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional fed.solo.io.FailoverSchemeSpec spec = 2;
 * @return {?proto.fed.solo.io.FailoverSchemeSpec}
 */
proto.fed.rpc.solo.io.FailoverScheme.prototype.getSpec = function() {
  return /** @type{?proto.fed.solo.io.FailoverSchemeSpec} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeSpec, 2));
};


/** @param {?proto.fed.solo.io.FailoverSchemeSpec|undefined} value */
proto.fed.rpc.solo.io.FailoverScheme.prototype.setSpec = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.fed.rpc.solo.io.FailoverScheme.prototype.clearSpec = function() {
  this.setSpec(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.rpc.solo.io.FailoverScheme.prototype.hasSpec = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional fed.solo.io.FailoverSchemeStatus status = 3;
 * @return {?proto.fed.solo.io.FailoverSchemeStatus}
 */
proto.fed.rpc.solo.io.FailoverScheme.prototype.getStatus = function() {
  return /** @type{?proto.fed.solo.io.FailoverSchemeStatus} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_gloo$fed_api_fed_v1_failover_pb.FailoverSchemeStatus, 3));
};


/** @param {?proto.fed.solo.io.FailoverSchemeStatus|undefined} value */
proto.fed.rpc.solo.io.FailoverScheme.prototype.setStatus = function(value) {
  jspb.Message.setWrapperField(this, 3, value);
};


proto.fed.rpc.solo.io.FailoverScheme.prototype.clearStatus = function() {
  this.setStatus(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.rpc.solo.io.FailoverScheme.prototype.hasStatus = function() {
  return jspb.Message.getField(this, 3) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.rpc.solo.io.GetFailoverSchemeRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.rpc.solo.io.GetFailoverSchemeRequest.displayName = 'proto.fed.rpc.solo.io.GetFailoverSchemeRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.rpc.solo.io.GetFailoverSchemeRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    upstreamRef: (f = msg.getUpstreamRef()) && github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fed.rpc.solo.io.GetFailoverSchemeRequest}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.rpc.solo.io.GetFailoverSchemeRequest;
  return proto.fed.rpc.solo.io.GetFailoverSchemeRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.rpc.solo.io.GetFailoverSchemeRequest}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef;
      reader.readMessage(value,github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef.deserializeBinaryFromReader);
      msg.setUpstreamRef(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.rpc.solo.io.GetFailoverSchemeRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUpstreamRef();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef.serializeBinaryToWriter
    );
  }
};


/**
 * optional core.skv2.solo.io.ClusterObjectRef upstream_ref = 1;
 * @return {?proto.core.skv2.solo.io.ClusterObjectRef}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.prototype.getUpstreamRef = function() {
  return /** @type{?proto.core.skv2.solo.io.ClusterObjectRef} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_skv2_api_core_v1_core_pb.ClusterObjectRef, 1));
};


/** @param {?proto.core.skv2.solo.io.ClusterObjectRef|undefined} value */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.prototype.setUpstreamRef = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.rpc.solo.io.GetFailoverSchemeRequest.prototype.clearUpstreamRef = function() {
  this.setUpstreamRef(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeRequest.prototype.hasUpstreamRef = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.rpc.solo.io.GetFailoverSchemeResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.rpc.solo.io.GetFailoverSchemeResponse.displayName = 'proto.fed.rpc.solo.io.GetFailoverSchemeResponse';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.rpc.solo.io.GetFailoverSchemeResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    failoverScheme: (f = msg.getFailoverScheme()) && proto.fed.rpc.solo.io.FailoverScheme.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fed.rpc.solo.io.GetFailoverSchemeResponse}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.rpc.solo.io.GetFailoverSchemeResponse;
  return proto.fed.rpc.solo.io.GetFailoverSchemeResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.rpc.solo.io.GetFailoverSchemeResponse}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.fed.rpc.solo.io.FailoverScheme;
      reader.readMessage(value,proto.fed.rpc.solo.io.FailoverScheme.deserializeBinaryFromReader);
      msg.setFailoverScheme(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.rpc.solo.io.GetFailoverSchemeResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFailoverScheme();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.fed.rpc.solo.io.FailoverScheme.serializeBinaryToWriter
    );
  }
};


/**
 * optional FailoverScheme failover_scheme = 1;
 * @return {?proto.fed.rpc.solo.io.FailoverScheme}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.prototype.getFailoverScheme = function() {
  return /** @type{?proto.fed.rpc.solo.io.FailoverScheme} */ (
    jspb.Message.getWrapperField(this, proto.fed.rpc.solo.io.FailoverScheme, 1));
};


/** @param {?proto.fed.rpc.solo.io.FailoverScheme|undefined} value */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.prototype.setFailoverScheme = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.rpc.solo.io.GetFailoverSchemeResponse.prototype.clearFailoverScheme = function() {
  this.setFailoverScheme(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeResponse.prototype.hasFailoverScheme = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.displayName = 'proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    failoverSchemeRef: (f = msg.getFailoverSchemeRef()) && github_com_solo$io_skv2_api_core_v1_core_pb.ObjectRef.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest;
  return proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_skv2_api_core_v1_core_pb.ObjectRef;
      reader.readMessage(value,github_com_solo$io_skv2_api_core_v1_core_pb.ObjectRef.deserializeBinaryFromReader);
      msg.setFailoverSchemeRef(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFailoverSchemeRef();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_skv2_api_core_v1_core_pb.ObjectRef.serializeBinaryToWriter
    );
  }
};


/**
 * optional core.skv2.solo.io.ObjectRef failover_scheme_ref = 1;
 * @return {?proto.core.skv2.solo.io.ObjectRef}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.prototype.getFailoverSchemeRef = function() {
  return /** @type{?proto.core.skv2.solo.io.ObjectRef} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_skv2_api_core_v1_core_pb.ObjectRef, 1));
};


/** @param {?proto.core.skv2.solo.io.ObjectRef|undefined} value */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.prototype.setFailoverSchemeRef = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.prototype.clearFailoverSchemeRef = function() {
  this.setFailoverSchemeRef(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlRequest.prototype.hasFailoverSchemeRef = function() {
  return jspb.Message.getField(this, 1) != null;
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.displayName = 'proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.toObject = function(includeInstance, msg) {
  var f, obj = {
    yamlData: (f = msg.getYamlData()) && github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ResourceYaml.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse;
  return proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ResourceYaml;
      reader.readMessage(value,github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ResourceYaml.deserializeBinaryFromReader);
      msg.setYamlData(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getYamlData();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ResourceYaml.serializeBinaryToWriter
    );
  }
};


/**
 * optional rpc.edge.gloo.solo.io.ResourceYaml yaml_data = 1;
 * @return {?proto.rpc.edge.gloo.solo.io.ResourceYaml}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.prototype.getYamlData = function() {
  return /** @type{?proto.rpc.edge.gloo.solo.io.ResourceYaml} */ (
    jspb.Message.getWrapperField(this, github_com_solo$io_solo$projects_projects_apiserver_api_rpc_edge_gloo_v1_common_pb.ResourceYaml, 1));
};


/** @param {?proto.rpc.edge.gloo.solo.io.ResourceYaml|undefined} value */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.prototype.setYamlData = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.prototype.clearYamlData = function() {
  this.setYamlData(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.fed.rpc.solo.io.GetFailoverSchemeYamlResponse.prototype.hasYamlData = function() {
  return jspb.Message.getField(this, 1) != null;
};


goog.object.extend(exports, proto.fed.rpc.solo.io);

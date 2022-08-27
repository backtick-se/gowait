// GENERATED CODE -- DO NOT EDIT!

'use strict';
var cowait_pb = require('./cowait_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');

function serialize_AwaitTaskReply(arg) {
  if (!(arg instanceof cowait_pb.AwaitTaskReply)) {
    throw new Error('Expected argument of type AwaitTaskReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_AwaitTaskReply(buffer_arg) {
  return cowait_pb.AwaitTaskReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_AwaitTaskReq(arg) {
  if (!(arg instanceof cowait_pb.AwaitTaskReq)) {
    throw new Error('Expected argument of type AwaitTaskReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_AwaitTaskReq(buffer_arg) {
  return cowait_pb.AwaitTaskReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ClusterEvent(arg) {
  if (!(arg instanceof cowait_pb.ClusterEvent)) {
    throw new Error('Expected argument of type ClusterEvent');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ClusterEvent(buffer_arg) {
  return cowait_pb.ClusterEvent.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ClusterInfoReply(arg) {
  if (!(arg instanceof cowait_pb.ClusterInfoReply)) {
    throw new Error('Expected argument of type ClusterInfoReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ClusterInfoReply(buffer_arg) {
  return cowait_pb.ClusterInfoReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ClusterInfoReq(arg) {
  if (!(arg instanceof cowait_pb.ClusterInfoReq)) {
    throw new Error('Expected argument of type ClusterInfoReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ClusterInfoReq(buffer_arg) {
  return cowait_pb.ClusterInfoReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_ClusterSubscribeReq(arg) {
  if (!(arg instanceof cowait_pb.ClusterSubscribeReq)) {
    throw new Error('Expected argument of type ClusterSubscribeReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_ClusterSubscribeReq(buffer_arg) {
  return cowait_pb.ClusterSubscribeReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreateTaskReply(arg) {
  if (!(arg instanceof cowait_pb.CreateTaskReply)) {
    throw new Error('Expected argument of type CreateTaskReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateTaskReply(buffer_arg) {
  return cowait_pb.CreateTaskReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_CreateTaskReq(arg) {
  if (!(arg instanceof cowait_pb.CreateTaskReq)) {
    throw new Error('Expected argument of type CreateTaskReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_CreateTaskReq(buffer_arg) {
  return cowait_pb.CreateTaskReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_KillTaskReply(arg) {
  if (!(arg instanceof cowait_pb.KillTaskReply)) {
    throw new Error('Expected argument of type KillTaskReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_KillTaskReply(buffer_arg) {
  return cowait_pb.KillTaskReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_KillTaskReq(arg) {
  if (!(arg instanceof cowait_pb.KillTaskReq)) {
    throw new Error('Expected argument of type KillTaskReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_KillTaskReq(buffer_arg) {
  return cowait_pb.KillTaskReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_LogEntry(arg) {
  if (!(arg instanceof cowait_pb.LogEntry)) {
    throw new Error('Expected argument of type LogEntry');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_LogEntry(buffer_arg) {
  return cowait_pb.LogEntry.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_LogSummary(arg) {
  if (!(arg instanceof cowait_pb.LogSummary)) {
    throw new Error('Expected argument of type LogSummary');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_LogSummary(buffer_arg) {
  return cowait_pb.LogSummary.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_QueryTasksReply(arg) {
  if (!(arg instanceof cowait_pb.QueryTasksReply)) {
    throw new Error('Expected argument of type QueryTasksReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_QueryTasksReply(buffer_arg) {
  return cowait_pb.QueryTasksReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_QueryTasksReq(arg) {
  if (!(arg instanceof cowait_pb.QueryTasksReq)) {
    throw new Error('Expected argument of type QueryTasksReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_QueryTasksReq(buffer_arg) {
  return cowait_pb.QueryTasksReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_TaskCompleteReply(arg) {
  if (!(arg instanceof cowait_pb.TaskCompleteReply)) {
    throw new Error('Expected argument of type TaskCompleteReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_TaskCompleteReply(buffer_arg) {
  return cowait_pb.TaskCompleteReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_TaskCompleteReq(arg) {
  if (!(arg instanceof cowait_pb.TaskCompleteReq)) {
    throw new Error('Expected argument of type TaskCompleteReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_TaskCompleteReq(buffer_arg) {
  return cowait_pb.TaskCompleteReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_TaskFailureReply(arg) {
  if (!(arg instanceof cowait_pb.TaskFailureReply)) {
    throw new Error('Expected argument of type TaskFailureReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_TaskFailureReply(buffer_arg) {
  return cowait_pb.TaskFailureReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_TaskFailureReq(arg) {
  if (!(arg instanceof cowait_pb.TaskFailureReq)) {
    throw new Error('Expected argument of type TaskFailureReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_TaskFailureReq(buffer_arg) {
  return cowait_pb.TaskFailureReq.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_TaskInitReply(arg) {
  if (!(arg instanceof cowait_pb.TaskInitReply)) {
    throw new Error('Expected argument of type TaskInitReply');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_TaskInitReply(buffer_arg) {
  return cowait_pb.TaskInitReply.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_TaskInitReq(arg) {
  if (!(arg instanceof cowait_pb.TaskInitReq)) {
    throw new Error('Expected argument of type TaskInitReq');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_TaskInitReq(buffer_arg) {
  return cowait_pb.TaskInitReq.deserializeBinary(new Uint8Array(buffer_arg));
}


var ExecutorService = exports['Executor'] = {
  taskInit: {
    path: '/Executor/TaskInit',
    requestStream: false,
    responseStream: false,
    requestType: cowait_pb.TaskInitReq,
    responseType: cowait_pb.TaskInitReply,
    requestSerialize: serialize_TaskInitReq,
    requestDeserialize: deserialize_TaskInitReq,
    responseSerialize: serialize_TaskInitReply,
    responseDeserialize: deserialize_TaskInitReply,
  },
  taskFailure: {
    path: '/Executor/TaskFailure',
    requestStream: false,
    responseStream: false,
    requestType: cowait_pb.TaskFailureReq,
    responseType: cowait_pb.TaskFailureReply,
    requestSerialize: serialize_TaskFailureReq,
    requestDeserialize: deserialize_TaskFailureReq,
    responseSerialize: serialize_TaskFailureReply,
    responseDeserialize: deserialize_TaskFailureReply,
  },
  taskComplete: {
    path: '/Executor/TaskComplete',
    requestStream: false,
    responseStream: false,
    requestType: cowait_pb.TaskCompleteReq,
    responseType: cowait_pb.TaskCompleteReply,
    requestSerialize: serialize_TaskCompleteReq,
    requestDeserialize: deserialize_TaskCompleteReq,
    responseSerialize: serialize_TaskCompleteReply,
    responseDeserialize: deserialize_TaskCompleteReply,
  },
  taskLog: {
    path: '/Executor/TaskLog',
    requestStream: true,
    responseStream: false,
    requestType: cowait_pb.LogEntry,
    responseType: cowait_pb.LogSummary,
    requestSerialize: serialize_LogEntry,
    requestDeserialize: deserialize_LogEntry,
    responseSerialize: serialize_LogSummary,
    responseDeserialize: deserialize_LogSummary,
  },
};

// Api Service
//
var CowaitService = exports['Cowait'] = {
  createTask: {
    path: '/Cowait/CreateTask',
    requestStream: false,
    responseStream: false,
    requestType: cowait_pb.CreateTaskReq,
    responseType: cowait_pb.CreateTaskReply,
    requestSerialize: serialize_CreateTaskReq,
    requestDeserialize: deserialize_CreateTaskReq,
    responseSerialize: serialize_CreateTaskReply,
    responseDeserialize: deserialize_CreateTaskReply,
  },
  queryTasks: {
    path: '/Cowait/QueryTasks',
    requestStream: false,
    responseStream: false,
    requestType: cowait_pb.QueryTasksReq,
    responseType: cowait_pb.QueryTasksReply,
    requestSerialize: serialize_QueryTasksReq,
    requestDeserialize: deserialize_QueryTasksReq,
    responseSerialize: serialize_QueryTasksReply,
    responseDeserialize: deserialize_QueryTasksReply,
  },
  killTask: {
    path: '/Cowait/KillTask',
    requestStream: false,
    responseStream: false,
    requestType: cowait_pb.KillTaskReq,
    responseType: cowait_pb.KillTaskReply,
    requestSerialize: serialize_KillTaskReq,
    requestDeserialize: deserialize_KillTaskReq,
    responseSerialize: serialize_KillTaskReply,
    responseDeserialize: deserialize_KillTaskReply,
  },
  awaitTask: {
    path: '/Cowait/AwaitTask',
    requestStream: false,
    responseStream: true,
    requestType: cowait_pb.AwaitTaskReq,
    responseType: cowait_pb.AwaitTaskReply,
    requestSerialize: serialize_AwaitTaskReq,
    requestDeserialize: deserialize_AwaitTaskReq,
    responseSerialize: serialize_AwaitTaskReply,
    responseDeserialize: deserialize_AwaitTaskReply,
  },
};

//
// Cluster Uplink
//
//
var ClusterService = exports['Cluster'] = {
  info: {
    path: '/Cluster/Info',
    requestStream: false,
    responseStream: false,
    requestType: cowait_pb.ClusterInfoReq,
    responseType: cowait_pb.ClusterInfoReply,
    requestSerialize: serialize_ClusterInfoReq,
    requestDeserialize: deserialize_ClusterInfoReq,
    responseSerialize: serialize_ClusterInfoReply,
    responseDeserialize: deserialize_ClusterInfoReply,
  },
  createTask: {
    path: '/Cluster/CreateTask',
    requestStream: false,
    responseStream: false,
    requestType: cowait_pb.CreateTaskReq,
    responseType: cowait_pb.CreateTaskReply,
    requestSerialize: serialize_CreateTaskReq,
    requestDeserialize: deserialize_CreateTaskReq,
    responseSerialize: serialize_CreateTaskReply,
    responseDeserialize: deserialize_CreateTaskReply,
  },
  subscribe: {
    path: '/Cluster/Subscribe',
    requestStream: false,
    responseStream: true,
    requestType: cowait_pb.ClusterSubscribeReq,
    responseType: cowait_pb.ClusterEvent,
    requestSerialize: serialize_ClusterSubscribeReq,
    requestDeserialize: deserialize_ClusterSubscribeReq,
    responseSerialize: serialize_ClusterEvent,
    responseDeserialize: deserialize_ClusterEvent,
  },
};


// package: 
// file: cowait.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as cowait_pb from "./cowait_pb";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

interface IExecutorService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    taskInit: IExecutorService_ITaskInit;
    taskFailure: IExecutorService_ITaskFailure;
    taskComplete: IExecutorService_ITaskComplete;
    taskLog: IExecutorService_ITaskLog;
}

interface IExecutorService_ITaskInit extends grpc.MethodDefinition<cowait_pb.TaskInitReq, cowait_pb.TaskInitReply> {
    path: "/Executor/TaskInit";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.TaskInitReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.TaskInitReq>;
    responseSerialize: grpc.serialize<cowait_pb.TaskInitReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.TaskInitReply>;
}
interface IExecutorService_ITaskFailure extends grpc.MethodDefinition<cowait_pb.TaskFailureReq, cowait_pb.TaskFailureReply> {
    path: "/Executor/TaskFailure";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.TaskFailureReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.TaskFailureReq>;
    responseSerialize: grpc.serialize<cowait_pb.TaskFailureReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.TaskFailureReply>;
}
interface IExecutorService_ITaskComplete extends grpc.MethodDefinition<cowait_pb.TaskCompleteReq, cowait_pb.TaskCompleteReply> {
    path: "/Executor/TaskComplete";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.TaskCompleteReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.TaskCompleteReq>;
    responseSerialize: grpc.serialize<cowait_pb.TaskCompleteReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.TaskCompleteReply>;
}
interface IExecutorService_ITaskLog extends grpc.MethodDefinition<cowait_pb.LogEntry, cowait_pb.LogSummary> {
    path: "/Executor/TaskLog";
    requestStream: true;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.LogEntry>;
    requestDeserialize: grpc.deserialize<cowait_pb.LogEntry>;
    responseSerialize: grpc.serialize<cowait_pb.LogSummary>;
    responseDeserialize: grpc.deserialize<cowait_pb.LogSummary>;
}

export const ExecutorService: IExecutorService;

export interface IExecutorServer extends grpc.UntypedServiceImplementation {
    taskInit: grpc.handleUnaryCall<cowait_pb.TaskInitReq, cowait_pb.TaskInitReply>;
    taskFailure: grpc.handleUnaryCall<cowait_pb.TaskFailureReq, cowait_pb.TaskFailureReply>;
    taskComplete: grpc.handleUnaryCall<cowait_pb.TaskCompleteReq, cowait_pb.TaskCompleteReply>;
    taskLog: grpc.handleClientStreamingCall<cowait_pb.LogEntry, cowait_pb.LogSummary>;
}

export interface IExecutorClient {
    taskInit(request: cowait_pb.TaskInitReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskInitReply) => void): grpc.ClientUnaryCall;
    taskInit(request: cowait_pb.TaskInitReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskInitReply) => void): grpc.ClientUnaryCall;
    taskInit(request: cowait_pb.TaskInitReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskInitReply) => void): grpc.ClientUnaryCall;
    taskFailure(request: cowait_pb.TaskFailureReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskFailureReply) => void): grpc.ClientUnaryCall;
    taskFailure(request: cowait_pb.TaskFailureReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskFailureReply) => void): grpc.ClientUnaryCall;
    taskFailure(request: cowait_pb.TaskFailureReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskFailureReply) => void): grpc.ClientUnaryCall;
    taskComplete(request: cowait_pb.TaskCompleteReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskCompleteReply) => void): grpc.ClientUnaryCall;
    taskComplete(request: cowait_pb.TaskCompleteReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskCompleteReply) => void): grpc.ClientUnaryCall;
    taskComplete(request: cowait_pb.TaskCompleteReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskCompleteReply) => void): grpc.ClientUnaryCall;
    taskLog(callback: (error: grpc.ServiceError | null, response: cowait_pb.LogSummary) => void): grpc.ClientWritableStream<cowait_pb.LogEntry>;
    taskLog(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.LogSummary) => void): grpc.ClientWritableStream<cowait_pb.LogEntry>;
    taskLog(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.LogSummary) => void): grpc.ClientWritableStream<cowait_pb.LogEntry>;
    taskLog(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.LogSummary) => void): grpc.ClientWritableStream<cowait_pb.LogEntry>;
}

export class ExecutorClient extends grpc.Client implements IExecutorClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public taskInit(request: cowait_pb.TaskInitReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskInitReply) => void): grpc.ClientUnaryCall;
    public taskInit(request: cowait_pb.TaskInitReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskInitReply) => void): grpc.ClientUnaryCall;
    public taskInit(request: cowait_pb.TaskInitReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskInitReply) => void): grpc.ClientUnaryCall;
    public taskFailure(request: cowait_pb.TaskFailureReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskFailureReply) => void): grpc.ClientUnaryCall;
    public taskFailure(request: cowait_pb.TaskFailureReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskFailureReply) => void): grpc.ClientUnaryCall;
    public taskFailure(request: cowait_pb.TaskFailureReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskFailureReply) => void): grpc.ClientUnaryCall;
    public taskComplete(request: cowait_pb.TaskCompleteReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskCompleteReply) => void): grpc.ClientUnaryCall;
    public taskComplete(request: cowait_pb.TaskCompleteReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskCompleteReply) => void): grpc.ClientUnaryCall;
    public taskComplete(request: cowait_pb.TaskCompleteReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.TaskCompleteReply) => void): grpc.ClientUnaryCall;
    public taskLog(callback: (error: grpc.ServiceError | null, response: cowait_pb.LogSummary) => void): grpc.ClientWritableStream<cowait_pb.LogEntry>;
    public taskLog(metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.LogSummary) => void): grpc.ClientWritableStream<cowait_pb.LogEntry>;
    public taskLog(options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.LogSummary) => void): grpc.ClientWritableStream<cowait_pb.LogEntry>;
    public taskLog(metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.LogSummary) => void): grpc.ClientWritableStream<cowait_pb.LogEntry>;
}

interface ICowaitService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    createTask: ICowaitService_ICreateTask;
    queryTasks: ICowaitService_IQueryTasks;
    killTask: ICowaitService_IKillTask;
    awaitTask: ICowaitService_IAwaitTask;
}

interface ICowaitService_ICreateTask extends grpc.MethodDefinition<cowait_pb.CreateTaskReq, cowait_pb.CreateTaskReply> {
    path: "/Cowait/CreateTask";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.CreateTaskReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.CreateTaskReq>;
    responseSerialize: grpc.serialize<cowait_pb.CreateTaskReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.CreateTaskReply>;
}
interface ICowaitService_IQueryTasks extends grpc.MethodDefinition<cowait_pb.QueryTasksReq, cowait_pb.QueryTasksReply> {
    path: "/Cowait/QueryTasks";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.QueryTasksReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.QueryTasksReq>;
    responseSerialize: grpc.serialize<cowait_pb.QueryTasksReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.QueryTasksReply>;
}
interface ICowaitService_IKillTask extends grpc.MethodDefinition<cowait_pb.KillTaskReq, cowait_pb.KillTaskReply> {
    path: "/Cowait/KillTask";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.KillTaskReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.KillTaskReq>;
    responseSerialize: grpc.serialize<cowait_pb.KillTaskReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.KillTaskReply>;
}
interface ICowaitService_IAwaitTask extends grpc.MethodDefinition<cowait_pb.AwaitTaskReq, cowait_pb.AwaitTaskReply> {
    path: "/Cowait/AwaitTask";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<cowait_pb.AwaitTaskReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.AwaitTaskReq>;
    responseSerialize: grpc.serialize<cowait_pb.AwaitTaskReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.AwaitTaskReply>;
}

export const CowaitService: ICowaitService;

export interface ICowaitServer extends grpc.UntypedServiceImplementation {
    createTask: grpc.handleUnaryCall<cowait_pb.CreateTaskReq, cowait_pb.CreateTaskReply>;
    queryTasks: grpc.handleUnaryCall<cowait_pb.QueryTasksReq, cowait_pb.QueryTasksReply>;
    killTask: grpc.handleUnaryCall<cowait_pb.KillTaskReq, cowait_pb.KillTaskReply>;
    awaitTask: grpc.handleServerStreamingCall<cowait_pb.AwaitTaskReq, cowait_pb.AwaitTaskReply>;
}

export interface ICowaitClient {
    createTask(request: cowait_pb.CreateTaskReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    createTask(request: cowait_pb.CreateTaskReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    createTask(request: cowait_pb.CreateTaskReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    queryTasks(request: cowait_pb.QueryTasksReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.QueryTasksReply) => void): grpc.ClientUnaryCall;
    queryTasks(request: cowait_pb.QueryTasksReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.QueryTasksReply) => void): grpc.ClientUnaryCall;
    queryTasks(request: cowait_pb.QueryTasksReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.QueryTasksReply) => void): grpc.ClientUnaryCall;
    killTask(request: cowait_pb.KillTaskReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.KillTaskReply) => void): grpc.ClientUnaryCall;
    killTask(request: cowait_pb.KillTaskReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.KillTaskReply) => void): grpc.ClientUnaryCall;
    killTask(request: cowait_pb.KillTaskReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.KillTaskReply) => void): grpc.ClientUnaryCall;
    awaitTask(request: cowait_pb.AwaitTaskReq, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<cowait_pb.AwaitTaskReply>;
    awaitTask(request: cowait_pb.AwaitTaskReq, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<cowait_pb.AwaitTaskReply>;
}

export class CowaitClient extends grpc.Client implements ICowaitClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public createTask(request: cowait_pb.CreateTaskReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    public createTask(request: cowait_pb.CreateTaskReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    public createTask(request: cowait_pb.CreateTaskReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    public queryTasks(request: cowait_pb.QueryTasksReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.QueryTasksReply) => void): grpc.ClientUnaryCall;
    public queryTasks(request: cowait_pb.QueryTasksReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.QueryTasksReply) => void): grpc.ClientUnaryCall;
    public queryTasks(request: cowait_pb.QueryTasksReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.QueryTasksReply) => void): grpc.ClientUnaryCall;
    public killTask(request: cowait_pb.KillTaskReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.KillTaskReply) => void): grpc.ClientUnaryCall;
    public killTask(request: cowait_pb.KillTaskReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.KillTaskReply) => void): grpc.ClientUnaryCall;
    public killTask(request: cowait_pb.KillTaskReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.KillTaskReply) => void): grpc.ClientUnaryCall;
    public awaitTask(request: cowait_pb.AwaitTaskReq, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<cowait_pb.AwaitTaskReply>;
    public awaitTask(request: cowait_pb.AwaitTaskReq, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<cowait_pb.AwaitTaskReply>;
}

interface IClusterService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    info: IClusterService_IInfo;
    createTask: IClusterService_ICreateTask;
    subscribe: IClusterService_ISubscribe;
}

interface IClusterService_IInfo extends grpc.MethodDefinition<cowait_pb.ClusterInfoReq, cowait_pb.ClusterInfoReply> {
    path: "/Cluster/Info";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.ClusterInfoReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.ClusterInfoReq>;
    responseSerialize: grpc.serialize<cowait_pb.ClusterInfoReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.ClusterInfoReply>;
}
interface IClusterService_ICreateTask extends grpc.MethodDefinition<cowait_pb.CreateTaskReq, cowait_pb.CreateTaskReply> {
    path: "/Cluster/CreateTask";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<cowait_pb.CreateTaskReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.CreateTaskReq>;
    responseSerialize: grpc.serialize<cowait_pb.CreateTaskReply>;
    responseDeserialize: grpc.deserialize<cowait_pb.CreateTaskReply>;
}
interface IClusterService_ISubscribe extends grpc.MethodDefinition<cowait_pb.ClusterSubscribeReq, cowait_pb.ClusterEvent> {
    path: "/Cluster/Subscribe";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<cowait_pb.ClusterSubscribeReq>;
    requestDeserialize: grpc.deserialize<cowait_pb.ClusterSubscribeReq>;
    responseSerialize: grpc.serialize<cowait_pb.ClusterEvent>;
    responseDeserialize: grpc.deserialize<cowait_pb.ClusterEvent>;
}

export const ClusterService: IClusterService;

export interface IClusterServer extends grpc.UntypedServiceImplementation {
    info: grpc.handleUnaryCall<cowait_pb.ClusterInfoReq, cowait_pb.ClusterInfoReply>;
    createTask: grpc.handleUnaryCall<cowait_pb.CreateTaskReq, cowait_pb.CreateTaskReply>;
    subscribe: grpc.handleServerStreamingCall<cowait_pb.ClusterSubscribeReq, cowait_pb.ClusterEvent>;
}

export interface IClusterClient {
    info(request: cowait_pb.ClusterInfoReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.ClusterInfoReply) => void): grpc.ClientUnaryCall;
    info(request: cowait_pb.ClusterInfoReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.ClusterInfoReply) => void): grpc.ClientUnaryCall;
    info(request: cowait_pb.ClusterInfoReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.ClusterInfoReply) => void): grpc.ClientUnaryCall;
    createTask(request: cowait_pb.CreateTaskReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    createTask(request: cowait_pb.CreateTaskReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    createTask(request: cowait_pb.CreateTaskReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    subscribe(request: cowait_pb.ClusterSubscribeReq, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<cowait_pb.ClusterEvent>;
    subscribe(request: cowait_pb.ClusterSubscribeReq, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<cowait_pb.ClusterEvent>;
}

export class ClusterClient extends grpc.Client implements IClusterClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public info(request: cowait_pb.ClusterInfoReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.ClusterInfoReply) => void): grpc.ClientUnaryCall;
    public info(request: cowait_pb.ClusterInfoReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.ClusterInfoReply) => void): grpc.ClientUnaryCall;
    public info(request: cowait_pb.ClusterInfoReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.ClusterInfoReply) => void): grpc.ClientUnaryCall;
    public createTask(request: cowait_pb.CreateTaskReq, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    public createTask(request: cowait_pb.CreateTaskReq, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    public createTask(request: cowait_pb.CreateTaskReq, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: cowait_pb.CreateTaskReply) => void): grpc.ClientUnaryCall;
    public subscribe(request: cowait_pb.ClusterSubscribeReq, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<cowait_pb.ClusterEvent>;
    public subscribe(request: cowait_pb.ClusterSubscribeReq, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<cowait_pb.ClusterEvent>;
}

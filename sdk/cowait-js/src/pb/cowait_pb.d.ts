// package: 
// file: cowait.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Header extends jspb.Message { 
    getId(): string;
    setId(value: string): Header;

    hasTime(): boolean;
    clearTime(): void;
    getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setTime(value?: google_protobuf_timestamp_pb.Timestamp): Header;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Header.AsObject;
    static toObject(includeInstance: boolean, msg: Header): Header.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Header, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Header;
    static deserializeBinaryFromReader(message: Header, reader: jspb.BinaryReader): Header;
}

export namespace Header {
    export type AsObject = {
        id: string,
        time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
}

export class TaskInitReq extends jspb.Message { 

    hasHeader(): boolean;
    clearHeader(): void;
    getHeader(): Header | undefined;
    setHeader(value?: Header): TaskInitReq;
    getVersion(): string;
    setVersion(value: string): TaskInitReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TaskInitReq.AsObject;
    static toObject(includeInstance: boolean, msg: TaskInitReq): TaskInitReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TaskInitReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TaskInitReq;
    static deserializeBinaryFromReader(message: TaskInitReq, reader: jspb.BinaryReader): TaskInitReq;
}

export namespace TaskInitReq {
    export type AsObject = {
        header?: Header.AsObject,
        version: string,
    }
}

export class TaskInitReply extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TaskInitReply.AsObject;
    static toObject(includeInstance: boolean, msg: TaskInitReply): TaskInitReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TaskInitReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TaskInitReply;
    static deserializeBinaryFromReader(message: TaskInitReply, reader: jspb.BinaryReader): TaskInitReply;
}

export namespace TaskInitReply {
    export type AsObject = {
    }
}

export class TaskFailureReq extends jspb.Message { 

    hasHeader(): boolean;
    clearHeader(): void;
    getHeader(): Header | undefined;
    setHeader(value?: Header): TaskFailureReq;
    getError(): string;
    setError(value: string): TaskFailureReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TaskFailureReq.AsObject;
    static toObject(includeInstance: boolean, msg: TaskFailureReq): TaskFailureReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TaskFailureReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TaskFailureReq;
    static deserializeBinaryFromReader(message: TaskFailureReq, reader: jspb.BinaryReader): TaskFailureReq;
}

export namespace TaskFailureReq {
    export type AsObject = {
        header?: Header.AsObject,
        error: string,
    }
}

export class TaskFailureReply extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TaskFailureReply.AsObject;
    static toObject(includeInstance: boolean, msg: TaskFailureReply): TaskFailureReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TaskFailureReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TaskFailureReply;
    static deserializeBinaryFromReader(message: TaskFailureReply, reader: jspb.BinaryReader): TaskFailureReply;
}

export namespace TaskFailureReply {
    export type AsObject = {
    }
}

export class TaskCompleteReq extends jspb.Message { 

    hasHeader(): boolean;
    clearHeader(): void;
    getHeader(): Header | undefined;
    setHeader(value?: Header): TaskCompleteReq;
    getResult(): string;
    setResult(value: string): TaskCompleteReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TaskCompleteReq.AsObject;
    static toObject(includeInstance: boolean, msg: TaskCompleteReq): TaskCompleteReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TaskCompleteReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TaskCompleteReq;
    static deserializeBinaryFromReader(message: TaskCompleteReq, reader: jspb.BinaryReader): TaskCompleteReq;
}

export namespace TaskCompleteReq {
    export type AsObject = {
        header?: Header.AsObject,
        result: string,
    }
}

export class TaskCompleteReply extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TaskCompleteReply.AsObject;
    static toObject(includeInstance: boolean, msg: TaskCompleteReply): TaskCompleteReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TaskCompleteReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TaskCompleteReply;
    static deserializeBinaryFromReader(message: TaskCompleteReply, reader: jspb.BinaryReader): TaskCompleteReply;
}

export namespace TaskCompleteReply {
    export type AsObject = {
    }
}

export class LogEntry extends jspb.Message { 

    hasHeader(): boolean;
    clearHeader(): void;
    getHeader(): Header | undefined;
    setHeader(value?: Header): LogEntry;
    getFile(): string;
    setFile(value: string): LogEntry;
    getData(): string;
    setData(value: string): LogEntry;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LogEntry.AsObject;
    static toObject(includeInstance: boolean, msg: LogEntry): LogEntry.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LogEntry, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LogEntry;
    static deserializeBinaryFromReader(message: LogEntry, reader: jspb.BinaryReader): LogEntry;
}

export namespace LogEntry {
    export type AsObject = {
        header?: Header.AsObject,
        file: string,
        data: string,
    }
}

export class LogSummary extends jspb.Message { 
    getRecords(): number;
    setRecords(value: number): LogSummary;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): LogSummary.AsObject;
    static toObject(includeInstance: boolean, msg: LogSummary): LogSummary.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: LogSummary, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): LogSummary;
    static deserializeBinaryFromReader(message: LogSummary, reader: jspb.BinaryReader): LogSummary;
}

export namespace LogSummary {
    export type AsObject = {
        records: number,
    }
}

export class Task extends jspb.Message { 
    getTaskId(): string;
    setTaskId(value: string): Task;
    getParent(): string;
    setParent(value: string): Task;
    getStatus(): string;
    setStatus(value: string): Task;

    hasSpec(): boolean;
    clearSpec(): void;
    getSpec(): TaskSpec | undefined;
    setSpec(value?: TaskSpec): Task;

    hasScheduled(): boolean;
    clearScheduled(): void;
    getScheduled(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setScheduled(value?: google_protobuf_timestamp_pb.Timestamp): Task;

    hasStarted(): boolean;
    clearStarted(): void;
    getStarted(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setStarted(value?: google_protobuf_timestamp_pb.Timestamp): Task;

    hasCompleted(): boolean;
    clearCompleted(): void;
    getCompleted(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setCompleted(value?: google_protobuf_timestamp_pb.Timestamp): Task;
    getResult(): string;
    setResult(value: string): Task;
    getError(): string;
    setError(value: string): Task;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Task.AsObject;
    static toObject(includeInstance: boolean, msg: Task): Task.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Task, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Task;
    static deserializeBinaryFromReader(message: Task, reader: jspb.BinaryReader): Task;
}

export namespace Task {
    export type AsObject = {
        taskId: string,
        parent: string,
        status: string,
        spec?: TaskSpec.AsObject,
        scheduled?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        started?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        completed?: google_protobuf_timestamp_pb.Timestamp.AsObject,
        result: string,
        error: string,
    }
}

export class TaskSpec extends jspb.Message { 
    getImage(): string;
    setImage(value: string): TaskSpec;
    getName(): string;
    setName(value: string): TaskSpec;
    clearCommandList(): void;
    getCommandList(): Array<string>;
    setCommandList(value: Array<string>): TaskSpec;
    addCommand(value: string, index?: number): string;
    getInput(): string;
    setInput(value: string): TaskSpec;
    getTimeout(): number;
    setTimeout(value: number): TaskSpec;

    hasTime(): boolean;
    clearTime(): void;
    getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
    setTime(value?: google_protobuf_timestamp_pb.Timestamp): TaskSpec;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TaskSpec.AsObject;
    static toObject(includeInstance: boolean, msg: TaskSpec): TaskSpec.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TaskSpec, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TaskSpec;
    static deserializeBinaryFromReader(message: TaskSpec, reader: jspb.BinaryReader): TaskSpec;
}

export namespace TaskSpec {
    export type AsObject = {
        image: string,
        name: string,
        commandList: Array<string>,
        input: string,
        timeout: number,
        time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    }
}

export class CreateTaskReq extends jspb.Message { 

    hasSpec(): boolean;
    clearSpec(): void;
    getSpec(): TaskSpec | undefined;
    setSpec(value?: TaskSpec): CreateTaskReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateTaskReq.AsObject;
    static toObject(includeInstance: boolean, msg: CreateTaskReq): CreateTaskReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateTaskReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateTaskReq;
    static deserializeBinaryFromReader(message: CreateTaskReq, reader: jspb.BinaryReader): CreateTaskReq;
}

export namespace CreateTaskReq {
    export type AsObject = {
        spec?: TaskSpec.AsObject,
    }
}

export class CreateTaskReply extends jspb.Message { 

    hasTask(): boolean;
    clearTask(): void;
    getTask(): Task | undefined;
    setTask(value?: Task): CreateTaskReply;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateTaskReply.AsObject;
    static toObject(includeInstance: boolean, msg: CreateTaskReply): CreateTaskReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateTaskReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateTaskReply;
    static deserializeBinaryFromReader(message: CreateTaskReply, reader: jspb.BinaryReader): CreateTaskReply;
}

export namespace CreateTaskReply {
    export type AsObject = {
        task?: Task.AsObject,
    }
}

export class QueryTasksReq extends jspb.Message { 
    getId(): string;
    setId(value: string): QueryTasksReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryTasksReq.AsObject;
    static toObject(includeInstance: boolean, msg: QueryTasksReq): QueryTasksReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryTasksReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryTasksReq;
    static deserializeBinaryFromReader(message: QueryTasksReq, reader: jspb.BinaryReader): QueryTasksReq;
}

export namespace QueryTasksReq {
    export type AsObject = {
        id: string,
    }
}

export class QueryTasksReply extends jspb.Message { 
    clearTasksList(): void;
    getTasksList(): Array<Task>;
    setTasksList(value: Array<Task>): QueryTasksReply;
    addTasks(value?: Task, index?: number): Task;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): QueryTasksReply.AsObject;
    static toObject(includeInstance: boolean, msg: QueryTasksReply): QueryTasksReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: QueryTasksReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): QueryTasksReply;
    static deserializeBinaryFromReader(message: QueryTasksReply, reader: jspb.BinaryReader): QueryTasksReply;
}

export namespace QueryTasksReply {
    export type AsObject = {
        tasksList: Array<Task.AsObject>,
    }
}

export class KillTaskReq extends jspb.Message { 
    getId(): string;
    setId(value: string): KillTaskReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): KillTaskReq.AsObject;
    static toObject(includeInstance: boolean, msg: KillTaskReq): KillTaskReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: KillTaskReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): KillTaskReq;
    static deserializeBinaryFromReader(message: KillTaskReq, reader: jspb.BinaryReader): KillTaskReq;
}

export namespace KillTaskReq {
    export type AsObject = {
        id: string,
    }
}

export class KillTaskReply extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): KillTaskReply.AsObject;
    static toObject(includeInstance: boolean, msg: KillTaskReply): KillTaskReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: KillTaskReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): KillTaskReply;
    static deserializeBinaryFromReader(message: KillTaskReply, reader: jspb.BinaryReader): KillTaskReply;
}

export namespace KillTaskReply {
    export type AsObject = {
    }
}

export class AwaitTaskReq extends jspb.Message { 
    getId(): string;
    setId(value: string): AwaitTaskReq;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AwaitTaskReq.AsObject;
    static toObject(includeInstance: boolean, msg: AwaitTaskReq): AwaitTaskReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AwaitTaskReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AwaitTaskReq;
    static deserializeBinaryFromReader(message: AwaitTaskReq, reader: jspb.BinaryReader): AwaitTaskReq;
}

export namespace AwaitTaskReq {
    export type AsObject = {
        id: string,
    }
}

export class AwaitTaskReply extends jspb.Message { 

    hasTask(): boolean;
    clearTask(): void;
    getTask(): Task | undefined;
    setTask(value?: Task): AwaitTaskReply;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AwaitTaskReply.AsObject;
    static toObject(includeInstance: boolean, msg: AwaitTaskReply): AwaitTaskReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AwaitTaskReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AwaitTaskReply;
    static deserializeBinaryFromReader(message: AwaitTaskReply, reader: jspb.BinaryReader): AwaitTaskReply;
}

export namespace AwaitTaskReply {
    export type AsObject = {
        task?: Task.AsObject,
    }
}

export class ClusterInfoReq extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterInfoReq.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterInfoReq): ClusterInfoReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterInfoReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterInfoReq;
    static deserializeBinaryFromReader(message: ClusterInfoReq, reader: jspb.BinaryReader): ClusterInfoReq;
}

export namespace ClusterInfoReq {
    export type AsObject = {
    }
}

export class ClusterInfoReply extends jspb.Message { 
    getClusterId(): string;
    setClusterId(value: string): ClusterInfoReply;
    getName(): string;
    setName(value: string): ClusterInfoReply;
    getKey(): string;
    setKey(value: string): ClusterInfoReply;
    getStatus(): string;
    setStatus(value: string): ClusterInfoReply;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterInfoReply.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterInfoReply): ClusterInfoReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterInfoReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterInfoReply;
    static deserializeBinaryFromReader(message: ClusterInfoReply, reader: jspb.BinaryReader): ClusterInfoReply;
}

export namespace ClusterInfoReply {
    export type AsObject = {
        clusterId: string,
        name: string,
        key: string,
        status: string,
    }
}

export class ClusterSpawnReq extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterSpawnReq.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterSpawnReq): ClusterSpawnReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterSpawnReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterSpawnReq;
    static deserializeBinaryFromReader(message: ClusterSpawnReq, reader: jspb.BinaryReader): ClusterSpawnReq;
}

export namespace ClusterSpawnReq {
    export type AsObject = {
    }
}

export class ClusterSpawnReply extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterSpawnReply.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterSpawnReply): ClusterSpawnReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterSpawnReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterSpawnReply;
    static deserializeBinaryFromReader(message: ClusterSpawnReply, reader: jspb.BinaryReader): ClusterSpawnReply;
}

export namespace ClusterSpawnReply {
    export type AsObject = {
    }
}

export class ClusterKillReq extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterKillReq.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterKillReq): ClusterKillReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterKillReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterKillReq;
    static deserializeBinaryFromReader(message: ClusterKillReq, reader: jspb.BinaryReader): ClusterKillReq;
}

export namespace ClusterKillReq {
    export type AsObject = {
    }
}

export class ClusterKillReply extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterKillReply.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterKillReply): ClusterKillReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterKillReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterKillReply;
    static deserializeBinaryFromReader(message: ClusterKillReply, reader: jspb.BinaryReader): ClusterKillReply;
}

export namespace ClusterKillReply {
    export type AsObject = {
    }
}

export class ClusterPokeReq extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterPokeReq.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterPokeReq): ClusterPokeReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterPokeReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterPokeReq;
    static deserializeBinaryFromReader(message: ClusterPokeReq, reader: jspb.BinaryReader): ClusterPokeReq;
}

export namespace ClusterPokeReq {
    export type AsObject = {
    }
}

export class ClusterPokeReply extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterPokeReply.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterPokeReply): ClusterPokeReply.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterPokeReply, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterPokeReply;
    static deserializeBinaryFromReader(message: ClusterPokeReply, reader: jspb.BinaryReader): ClusterPokeReply;
}

export namespace ClusterPokeReply {
    export type AsObject = {
    }
}

export class ClusterSubscribeReq extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterSubscribeReq.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterSubscribeReq): ClusterSubscribeReq.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterSubscribeReq, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterSubscribeReq;
    static deserializeBinaryFromReader(message: ClusterSubscribeReq, reader: jspb.BinaryReader): ClusterSubscribeReq;
}

export namespace ClusterSubscribeReq {
    export type AsObject = {
    }
}

export class ClusterEvent extends jspb.Message { 
    getClusterId(): string;
    setClusterId(value: string): ClusterEvent;
    getType(): string;
    setType(value: string): ClusterEvent;

    hasTask(): boolean;
    clearTask(): void;
    getTask(): Task | undefined;
    setTask(value?: Task): ClusterEvent;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): ClusterEvent.AsObject;
    static toObject(includeInstance: boolean, msg: ClusterEvent): ClusterEvent.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: ClusterEvent, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): ClusterEvent;
    static deserializeBinaryFromReader(message: ClusterEvent, reader: jspb.BinaryReader): ClusterEvent;
}

export namespace ClusterEvent {
    export type AsObject = {
        clusterId: string,
        type: string,
        task?: Task.AsObject,
    }
}

# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from . import cowait_pb2 as cowait__pb2


class ExecutorStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ExecInit = channel.unary_unary(
                '/Executor/ExecInit',
                request_serializer=cowait__pb2.ExecInitReq.SerializeToString,
                response_deserializer=cowait__pb2.ExecInitReply.FromString,
                )
        self.ExecAquire = channel.unary_unary(
                '/Executor/ExecAquire',
                request_serializer=cowait__pb2.ExecAquireReq.SerializeToString,
                response_deserializer=cowait__pb2.ExecAquireReply.FromString,
                )
        self.ExecStop = channel.unary_unary(
                '/Executor/ExecStop',
                request_serializer=cowait__pb2.ExecStopReq.SerializeToString,
                response_deserializer=cowait__pb2.ExecStopReply.FromString,
                )
        self.TaskInit = channel.unary_unary(
                '/Executor/TaskInit',
                request_serializer=cowait__pb2.TaskInitReq.SerializeToString,
                response_deserializer=cowait__pb2.TaskInitReply.FromString,
                )
        self.TaskFailure = channel.unary_unary(
                '/Executor/TaskFailure',
                request_serializer=cowait__pb2.TaskFailureReq.SerializeToString,
                response_deserializer=cowait__pb2.TaskFailureReply.FromString,
                )
        self.TaskComplete = channel.unary_unary(
                '/Executor/TaskComplete',
                request_serializer=cowait__pb2.TaskCompleteReq.SerializeToString,
                response_deserializer=cowait__pb2.TaskCompleteReply.FromString,
                )
        self.TaskLog = channel.stream_unary(
                '/Executor/TaskLog',
                request_serializer=cowait__pb2.LogEntry.SerializeToString,
                response_deserializer=cowait__pb2.LogSummary.FromString,
                )


class ExecutorServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ExecInit(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ExecAquire(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ExecStop(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TaskInit(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TaskFailure(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TaskComplete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TaskLog(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ExecutorServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ExecInit': grpc.unary_unary_rpc_method_handler(
                    servicer.ExecInit,
                    request_deserializer=cowait__pb2.ExecInitReq.FromString,
                    response_serializer=cowait__pb2.ExecInitReply.SerializeToString,
            ),
            'ExecAquire': grpc.unary_unary_rpc_method_handler(
                    servicer.ExecAquire,
                    request_deserializer=cowait__pb2.ExecAquireReq.FromString,
                    response_serializer=cowait__pb2.ExecAquireReply.SerializeToString,
            ),
            'ExecStop': grpc.unary_unary_rpc_method_handler(
                    servicer.ExecStop,
                    request_deserializer=cowait__pb2.ExecStopReq.FromString,
                    response_serializer=cowait__pb2.ExecStopReply.SerializeToString,
            ),
            'TaskInit': grpc.unary_unary_rpc_method_handler(
                    servicer.TaskInit,
                    request_deserializer=cowait__pb2.TaskInitReq.FromString,
                    response_serializer=cowait__pb2.TaskInitReply.SerializeToString,
            ),
            'TaskFailure': grpc.unary_unary_rpc_method_handler(
                    servicer.TaskFailure,
                    request_deserializer=cowait__pb2.TaskFailureReq.FromString,
                    response_serializer=cowait__pb2.TaskFailureReply.SerializeToString,
            ),
            'TaskComplete': grpc.unary_unary_rpc_method_handler(
                    servicer.TaskComplete,
                    request_deserializer=cowait__pb2.TaskCompleteReq.FromString,
                    response_serializer=cowait__pb2.TaskCompleteReply.SerializeToString,
            ),
            'TaskLog': grpc.stream_unary_rpc_method_handler(
                    servicer.TaskLog,
                    request_deserializer=cowait__pb2.LogEntry.FromString,
                    response_serializer=cowait__pb2.LogSummary.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Executor', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Executor(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ExecInit(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Executor/ExecInit',
            cowait__pb2.ExecInitReq.SerializeToString,
            cowait__pb2.ExecInitReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ExecAquire(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Executor/ExecAquire',
            cowait__pb2.ExecAquireReq.SerializeToString,
            cowait__pb2.ExecAquireReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ExecStop(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Executor/ExecStop',
            cowait__pb2.ExecStopReq.SerializeToString,
            cowait__pb2.ExecStopReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TaskInit(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Executor/TaskInit',
            cowait__pb2.TaskInitReq.SerializeToString,
            cowait__pb2.TaskInitReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TaskFailure(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Executor/TaskFailure',
            cowait__pb2.TaskFailureReq.SerializeToString,
            cowait__pb2.TaskFailureReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TaskComplete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Executor/TaskComplete',
            cowait__pb2.TaskCompleteReq.SerializeToString,
            cowait__pb2.TaskCompleteReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def TaskLog(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/Executor/TaskLog',
            cowait__pb2.LogEntry.SerializeToString,
            cowait__pb2.LogSummary.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class CowaitStub(object):
    """Api Service

    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateTask = channel.unary_unary(
                '/Cowait/CreateTask',
                request_serializer=cowait__pb2.CreateTaskReq.SerializeToString,
                response_deserializer=cowait__pb2.CreateTaskReply.FromString,
                )
        self.QueryTasks = channel.unary_unary(
                '/Cowait/QueryTasks',
                request_serializer=cowait__pb2.QueryTasksReq.SerializeToString,
                response_deserializer=cowait__pb2.QueryTasksReply.FromString,
                )
        self.KillTask = channel.unary_unary(
                '/Cowait/KillTask',
                request_serializer=cowait__pb2.KillTaskReq.SerializeToString,
                response_deserializer=cowait__pb2.KillTaskReply.FromString,
                )
        self.AwaitTask = channel.unary_stream(
                '/Cowait/AwaitTask',
                request_serializer=cowait__pb2.AwaitTaskReq.SerializeToString,
                response_deserializer=cowait__pb2.AwaitTaskReply.FromString,
                )


class CowaitServicer(object):
    """Api Service

    """

    def CreateTask(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def QueryTasks(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def KillTask(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def AwaitTask(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_CowaitServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateTask': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateTask,
                    request_deserializer=cowait__pb2.CreateTaskReq.FromString,
                    response_serializer=cowait__pb2.CreateTaskReply.SerializeToString,
            ),
            'QueryTasks': grpc.unary_unary_rpc_method_handler(
                    servicer.QueryTasks,
                    request_deserializer=cowait__pb2.QueryTasksReq.FromString,
                    response_serializer=cowait__pb2.QueryTasksReply.SerializeToString,
            ),
            'KillTask': grpc.unary_unary_rpc_method_handler(
                    servicer.KillTask,
                    request_deserializer=cowait__pb2.KillTaskReq.FromString,
                    response_serializer=cowait__pb2.KillTaskReply.SerializeToString,
            ),
            'AwaitTask': grpc.unary_stream_rpc_method_handler(
                    servicer.AwaitTask,
                    request_deserializer=cowait__pb2.AwaitTaskReq.FromString,
                    response_serializer=cowait__pb2.AwaitTaskReply.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Cowait', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Cowait(object):
    """Api Service

    """

    @staticmethod
    def CreateTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Cowait/CreateTask',
            cowait__pb2.CreateTaskReq.SerializeToString,
            cowait__pb2.CreateTaskReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def QueryTasks(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Cowait/QueryTasks',
            cowait__pb2.QueryTasksReq.SerializeToString,
            cowait__pb2.QueryTasksReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def KillTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Cowait/KillTask',
            cowait__pb2.KillTaskReq.SerializeToString,
            cowait__pb2.KillTaskReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def AwaitTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/Cowait/AwaitTask',
            cowait__pb2.AwaitTaskReq.SerializeToString,
            cowait__pb2.AwaitTaskReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)


class ClusterStub(object):
    """
    Cluster Uplink


    """

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Info = channel.unary_unary(
                '/Cluster/Info',
                request_serializer=cowait__pb2.ClusterInfoReq.SerializeToString,
                response_deserializer=cowait__pb2.ClusterInfoReply.FromString,
                )
        self.CreateTask = channel.unary_unary(
                '/Cluster/CreateTask',
                request_serializer=cowait__pb2.CreateTaskReq.SerializeToString,
                response_deserializer=cowait__pb2.CreateTaskReply.FromString,
                )
        self.KillTask = channel.unary_unary(
                '/Cluster/KillTask',
                request_serializer=cowait__pb2.KillTaskReq.SerializeToString,
                response_deserializer=cowait__pb2.KillTaskReply.FromString,
                )
        self.Subscribe = channel.unary_stream(
                '/Cluster/Subscribe',
                request_serializer=cowait__pb2.ClusterSubscribeReq.SerializeToString,
                response_deserializer=cowait__pb2.ClusterEvent.FromString,
                )


class ClusterServicer(object):
    """
    Cluster Uplink


    """

    def Info(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateTask(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def KillTask(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def Subscribe(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ClusterServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Info': grpc.unary_unary_rpc_method_handler(
                    servicer.Info,
                    request_deserializer=cowait__pb2.ClusterInfoReq.FromString,
                    response_serializer=cowait__pb2.ClusterInfoReply.SerializeToString,
            ),
            'CreateTask': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateTask,
                    request_deserializer=cowait__pb2.CreateTaskReq.FromString,
                    response_serializer=cowait__pb2.CreateTaskReply.SerializeToString,
            ),
            'KillTask': grpc.unary_unary_rpc_method_handler(
                    servicer.KillTask,
                    request_deserializer=cowait__pb2.KillTaskReq.FromString,
                    response_serializer=cowait__pb2.KillTaskReply.SerializeToString,
            ),
            'Subscribe': grpc.unary_stream_rpc_method_handler(
                    servicer.Subscribe,
                    request_deserializer=cowait__pb2.ClusterSubscribeReq.FromString,
                    response_serializer=cowait__pb2.ClusterEvent.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'Cluster', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Cluster(object):
    """
    Cluster Uplink


    """

    @staticmethod
    def Info(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Cluster/Info',
            cowait__pb2.ClusterInfoReq.SerializeToString,
            cowait__pb2.ClusterInfoReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CreateTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Cluster/CreateTask',
            cowait__pb2.CreateTaskReq.SerializeToString,
            cowait__pb2.CreateTaskReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def KillTask(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/Cluster/KillTask',
            cowait__pb2.KillTaskReq.SerializeToString,
            cowait__pb2.KillTaskReply.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def Subscribe(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/Cluster/Subscribe',
            cowait__pb2.ClusterSubscribeReq.SerializeToString,
            cowait__pb2.ClusterEvent.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

import json
import grpc
from google.protobuf.timestamp_pb2 import Timestamp
from .pb.cowait_pb2_grpc import ExecutorStub
from .pb.cowait_pb2 import ExecAquireReq, ExecInitReq, Header, TaskCompleteReq, TaskFailureReq, TaskInitReq
from .taskdef import Taskdef

VERSION = 'cowait-python/1.0'


class Client:
    def __init__(self, executor_id: str, host: str = '0.0.0.0', port: int = 1337):
        self.executor_id = executor_id
        self.host = host
        self.port = port
        self._channel = grpc.insecure_channel(f'{host}:{port}')
        self._client = ExecutorStub(self._channel)

    def _header(self, id) -> Header:
        return Header(
            id=id,
            time=Timestamp(),
        )

    def executor_init(self, image: str):
        self._client.ExecInit(ExecInitReq(
            header=self._header(self.executor_id),
            image=image,
        ))

    def dequeue(self) -> Taskdef:
        reply = self._client.ExecAquire(ExecAquireReq(
            header=self._header(self.executor_id),
        ))
        if not reply.next:
            return None
        return Taskdef(
            id=reply.next.task_id,
            name=reply.next.spec.name,
            input=json.loads(reply.next.spec.input or "{}"),
            image=reply.next.spec.image,
            time=reply.next.spec.time,
        )

    def init(self, task_id) -> None:
        self._client.TaskInit(TaskInitReq(
            header=self._header(task_id),
            version=VERSION,
            executor=self.executor_id,
        ))

    def complete(self, task_id, result: any) -> None:
        self._client.TaskComplete(TaskCompleteReq(
            header=self._header(task_id),
            result=json.dumps(result)
        ))

    def failure(self, task_id, error: str) -> None:
        self._client.TaskFailure(TaskFailureReq(
            header=self._header(task_id),
            error=error,
        ))

import json
import grpc
from google.protobuf.timestamp_pb2 import Timestamp
from .pb.cowait_pb2_grpc import TaskStub
from .pb.cowait_pb2 import Header, TaskCompleteReq, TaskFailureReq, TaskInitReq

VERSION = 'cowait-python/1.0'


class Client:
    def __init__(self, task_id: str, host: str = 'localhost', port: int = 1337):
        self.task_id = task_id
        self.host = host
        self.port = port
        self._channel = grpc.insecure_channel(f'{host}:{port}')
        self._client = TaskStub(self._channel)

    def _header(self) -> Header:
        return Header(
            id=self.task_id,
            time=Timestamp(),
        )

    def init(self):
        self._client.TaskInit(TaskInitReq(
            header=self._header(),
            version=VERSION,
        ))

    def complete(self, result: any):
        self._client.TaskComplete(TaskCompleteReq(
            header=self._header(),
            result=json.dumps(result)
        ))

    def failure(self, error: str):
        self._client.TaskFailure(TaskFailureReq(
            header=self._header(),
            error=error,
        ))

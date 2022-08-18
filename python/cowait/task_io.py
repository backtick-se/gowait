import os
import json
from io import TextIOWrapper


class Upstream:
    def __init__(self, fd: TextIOWrapper):
        self.fd = fd

    def _write(self, type: str, msg: dict):
        data = json.dumps({
            'Type': type,
            'Body': msg,
        })
        self.fd.write(data)
        self.fd.write('\n')
        self.fd.flush()

    def exit(self, result: dict):
        self._write('cowait/result', result)
        
    def invoke(self, taskdef: dict):
        self._write('cowait/invoke', taskdef)

    @staticmethod
    def open() -> 'Upstream':
        return Upstream(os.fdopen(3, "w"))


class Downstream:
    def __init__(self, fd: TextIOWrapper):
        self.fd = fd

    def _read(self) -> dict:
        data = self.fd.readline()
        msg = json.loads(data)
        if 'Type' not in msg:
            raise ValueError('Invalid message')
        if 'Type' == 'cowait/error':
            raise RuntimeError(msg['Body'])
        return msg

    def read_result(self) -> dict:
        msg = self._read()
        if msg['Type'] != 'cowait/result':
            raise ValueError('unexpected message type')
        return msg['Body']

    @staticmethod
    def open() -> 'Downstream':
        return Downstream(os.fdopen(4, "r"))

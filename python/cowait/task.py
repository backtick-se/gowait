import os
import sys
import json
from datetime import datetime
from .client import Client

_client: Client = None
_taskdef: dict = None


def _excepthook(type, value, trace):
    global _client
    sys.__excepthook__(type, value, trace)
    _client.failure(f'{type.__name__}: {value}')


def _init():
    global _client, _taskdef
    _taskdef = taskdef_from_env()
    _client = Client(os.getenv('COWAIT_ID'))
    _client.init()
    sys.excepthook = _excepthook


def taskdef_from_env():
    global _taskdef
    if _taskdef == None:
        taskjson = os.getenv('COWAIT_TASK')
        _taskdef = json.loads(taskjson)
        _taskdef['Time'] = datetime.strptime(_taskdef['Time'].split(".")[0] + "Z", "%Y-%m-%dT%H:%M:%SZ")
    return _taskdef


def inputs() -> dict:
    td = taskdef_from_env()
    return td['Input']


def error(error: str):
    global _client
    _client.failure(error)


def exit(result: dict):
    global _client
    _client.complete(result)


def time() -> datetime:
    """ logical execution time """
    global _taskdef
    return _taskdef['Time']


_init()

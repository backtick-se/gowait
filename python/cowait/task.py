import os
import sys
import json
from .task_io import Upstream, Downstream


_upstream = Upstream.open()
_downstream = Downstream.open()
_taskdef = None


def taskdef_from_env():
    global _taskdef
    if _taskdef == None:
        taskjson = os.getenv('COWAIT_TASK')
        _taskdef = json.loads(taskjson)
    return _taskdef


def inputs() -> dict:
    td = taskdef_from_env()
    return td['Input']


def exit(result: dict):
    global _upstream
    _upstream.exit(result)
    sys.exit(0)


def invoke(name, **inputs):
    global _upstream, _downstream
    _upstream.invoke({
        'Name': name,
        'Image': taskdef_from_env()['Image'],
        'Input': json.dumps(inputs),
    })
    return _downstream.read_result()

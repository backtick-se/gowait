import os
import sys
import json
import functools

upstream = os.fdopen(3, "w")
downstream = os.fdopen(4, "r")

__taskdef__ = None

def taskdef_from_env():
    global __taskdef__
    if __taskdef__ == None:
        taskjson = os.getenv('COWAIT_TASK')
        __taskdef__ = json.loads(taskjson)
        __taskdef__['Input'] = json.loads(__taskdef__.get('Input', '{}'))
    return __taskdef__


def inputs() -> dict:
    td = taskdef_from_env()
    return td['Input']

def exit(result: dict):
    r = json.dumps(result)
    msg = json.dumps({
        'type': 'task/return',
        'result': r,
    })
    upstream.write(msg)
    upstream.write('\n')
    sys.exit(0)

def task(f):
    @functools.wraps(f)
    def wrapper(*args, **kwargs):
        print('calling task', f.__name__)
        return f(*args, **kwargs)
    return wrapper

def invoke(name, **inputs):
    upstream.write(json.dumps({
        'type': 'task/call',
        'def': {
            'Name': name,
            'Image': taskdef_from_env()['Image'],
            'Input': json.dumps(inputs),
        },
    }))
    result = json.loads(downstream.readline())
    return result['result']

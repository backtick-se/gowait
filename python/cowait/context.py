from typing import List
from .client import Client
from .task import Task
from .taskdef import Taskdef


class Context:
    task: Task
    client: Client
    taskdef: Taskdef

    def __init__(self, client, taskdef, task):
        self.client = client
        self.taskdef = taskdef
        self.task = task


_context: List[Context] = []


def push_context(context: Context):
    global _context
    _context.append(context)


def pop_context():
    global _context
    _context.pop()


def get_context() -> Context:
    global _context
    if len(_context) == 0:
        raise RuntimeError("No context set")
    return _context[-1]

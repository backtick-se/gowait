from typing import Union
from .client import Client
from .context import Context, push_context, pop_context
from .task import Task, get_task
from .taskdef import Taskdef


Executable = Union[Task, str]

def execute(client: Client, taskdef: Taskdef, task: Executable) -> any:
    if isinstance(task, str):
        task = get_task(task)

    context = Context(client, taskdef, task)

    # prepare arguments
    args = {
        **taskdef.input,
    }
    if task.accepts_context:
        args['context'] = context

    # execute task
    push_context(context)
    try:
        return task(**args)
    finally:
        pop_context()

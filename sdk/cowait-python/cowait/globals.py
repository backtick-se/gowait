import sys
from datetime import datetime
from .context import get_context
from .executor import execute, Executable
from .taskdef import Taskdef


def invoke(task: Executable, **input) -> any:
    ctx = get_context()
    taskdef = Taskdef(
        id="subtask",
        input=input,
        image=ctx.taskdef.image,
        time=datetime.now(),
    )
    return execute(ctx.client, taskdef, task)


def input() -> dict:
    ctx = get_context()
    return ctx.taskdef.input


def exit(result: any):
    ctx = get_context()
    ctx.client.complete(result)
    sys.exit(0)

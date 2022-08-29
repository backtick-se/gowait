import os
import inspect
import importlib
from typing import Callable, Dict, List


class Task:
    def __init__(self, func: Callable):
        if not inspect.isfunction(func):
            raise ValueError('Task must be a function')

        self.func = func
        self.name = f'{func.__module__}.{func.__name__}'
        self.desc = func.__doc__.strip() if func.__doc__ else ''

        sig = inspect.signature(func)
        self.input = sig.parameters
        self.output = sig.return_annotation

    def __call__(self, *args, **kwargs) -> any:
        return self.func(*args, **kwargs)

    @property
    def accepts_context(self) -> bool:
        return 'context' in self.input

    def json(self) -> dict:
        return {
            'Name': self.name,
            'Description': self.desc,
            'Input': [
                {
                    'Name': name,
                    'Type': str(param.annotation),
                    'Default': str(param.default),
                }
                for name, param in self.input.items()
            ],
            'Output': str(self.output),
        }
    

_tasks: Dict[str, Task] = {}


def get_task(name: str) -> Task:
    if name not in _tasks:
        raise ValueError(f'No such task: {name}')
    return _tasks[name]


def get_tasks() -> List[Task]:
    return _tasks.values()


def find_tasks(path: str, recursive: bool = False):
    for f in os.listdir(path):
        if os.path.isdir(f):
            if recursive:
                find_tasks(f, recursive=True)
        elif f.endswith('.py'):
            try:
                importlib.import_module(f[:f.rindex('.py')])
            except Exception as e:
                print(f'failed to import {f}: {str(e)}')


def task():
    """
    Task function decorator
    """
    def deco(func):
        global _tasks
        task = Task(func)
        _tasks[task.name] = task
        return func
    return deco

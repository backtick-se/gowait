from typing import List
from .task import task, get_tasks


@task()
def enumerate(**input) -> List[dict]:
    """
    Enumerate all available tasks.
    """
    results = []
    for task in get_tasks():
        print(task.name, task)
        if task.desc:
            print('\t', task.desc)
        print('\t', 'input:')
        for input in task.input:
            print('\t\t', input)
        print('\t', 'output:', task.output)

        results.append(task.json())

    return results

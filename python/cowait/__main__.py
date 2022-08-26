import os
import sys
from cowait.client import Client
from cowait.task import find_tasks
from cowait.executor import execute
from cowait.taskdef import Taskdef


def main() -> int:
    client = Client(os.getenv('COWAIT_ID'))
    client.init()

    # install a global exception hook as a last resort to report all uncaught exceptions
    def _excepthook(type, value, trace):
        sys.__excepthook__(type, value, trace)
        try:
            client.failure(f'{type.__name__}: {value}')
        except Exception as e:
            print('failed to report exception:')
            print(e)

    sys.excepthook = _excepthook

    if len(sys.argv) != 3:
        print("usage: python -m cowait exec <task-name>")
        print(sys.argv)
        return 1

    # discover tasks
    find_tasks(os.getcwd())

    taskdef = Taskdef.from_env()
    name = sys.argv[2]

    try:
        result = execute(client, taskdef, name)
        client.complete(result)

    except Exception as e:
        client.failure(str(e))


if __name__ == '__main__':
    r = main()
    sys.exit(r or 0)

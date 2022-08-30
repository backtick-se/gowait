import os
import sys
import traceback
from cowait.client import Client
from cowait.task import find_tasks
from cowait.executor import execute
from cowait.taskdef import Taskdef


def main() -> int:
    client = Client(os.getenv('COWAIT_ID'))
    client.executor_init(os.getenv('COWAIT_IMAGE'))
    taskdef = client.dequeue()

    # install a global exception hook as a last resort to report all uncaught exceptions
    def _excepthook(type, value, trace):
        sys.__excepthook__(type, value, trace)
        try:
            trace = traceback.format_exc()
            client.failure(f'{type.__name__}: {value}')
            client.close()
        except Exception as e:
            print('failed to report exception:')
            print(e)

    sys.excepthook = _excepthook

    # discover tasks
    find_tasks(os.getcwd())

    client.init(taskdef.id)
    try:
        result = execute(client, taskdef, taskdef.name)
        client.complete(taskdef.id, result)

    except Exception as e:
        client.failure(taskdef.id, str(e))

    finally:
        client.close()


if __name__ == '__main__':
    r = main()
    sys.exit(r or 0)

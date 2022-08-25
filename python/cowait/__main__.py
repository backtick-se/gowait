import os
import sys
import importlib



def main() -> int:
    if len(sys.argv) < 2:
        print("no command given")
        return 1

    # discover tasks
    pyfiles = [f[:f.index('.py')] for f in os.listdir('.') if f.lower().endswith('.py')]
    for f in pyfiles:
        importlib.import_module(f)

    from cowait.task import _init
    _init()


if __name__ == '__main__':
    r = main()
    sys.exit(r)

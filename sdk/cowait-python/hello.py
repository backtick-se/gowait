import cowait
from cowait import Context


@cowait.task()
def my_task(context: Context, **input):
    """
    A very undescriptive description of what the task does.
    """
    print("hello world")
    print("input:", input)

    print("evil\nthis line does not end", end='')
    print(" ... until here")

    result = cowait.invoke('hello.square', value=5)
    # raise RuntimeError('something went to shit')

    return {
        'hello': 'world',
        'squared': result,
    }


@cowait.task()
def square(value: int) -> int:
    return value**2

import cowait


@cowait.task(rpc=False)
def my_task():
    print("hello world")

    print("evil\nthis line does not end", end='')
    print(" ... until here")

    args = cowait.inputs()
    # result = cowait.invoke('subtask', value=5)

    raise RuntimeError('something went to shit')

    cowait.exit({
        'hello': 'world',
    })

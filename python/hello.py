print("hello world")

print("evil\nthis line does not end", end='')
print(" ... until here")

import cowait

@cowait.task
def subtask(a):
    return a + a


args = cowait.inputs()
result = cowait.invoke('subtask', value=5)

cowait.exit({
    'result': result,
})
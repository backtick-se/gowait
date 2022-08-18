import cowait

print("hello world")

print("evil\nthis line does not end", end='')
print(" ... until here")

args = cowait.inputs()
result = cowait.invoke('subtask', value=5)

cowait.exit({
    'result': result,
})

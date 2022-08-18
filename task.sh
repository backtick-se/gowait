#!/bin/bash

set -e

command='["echo", "hello world"]'
command='["python3", "-u", "/Users/johanh/dev/cowait2/python/hello.py"]'

export COWAIT_TASK=$(
cat << EOF
{
    "ID": "test-task",
    "Name": "python",
    "Namespace": "default",
    "Image": "cowait/gowait", 
    "Command": ${command},
    "Input": {
        "value": 2
    }
} 
EOF
)

go run ./cmd/executor
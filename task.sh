#!/bin/bash

set -e

command='["sleep", "10"]'
# command='["python3", "-u", "/Users/johanh/dev/cowait2/python/hello.py"]'

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
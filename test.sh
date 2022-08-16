#!/bin/bash

set -e

taskdef=$(
cat << EOF
{
    "Name": "gowait-task",
    "Namespace": "default",
    "Image": "cowait/gowait-python", 
    "Command":["python", "-u", "hello.py"],
    "Input": "{\"value\":2}"
} 
EOF
)

docker compose build base
docker compose build python
docker run --rm -e COWAIT_TASK="$taskdef" cowait/gowait-python

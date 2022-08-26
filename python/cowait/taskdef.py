import os
import json
from datetime import datetime


class Taskdef:
    image: str
    input: dict
    time: datetime

    def __init__(self, image, input, time):
        self.image = image
        self.input = input
        self.time = time

    @staticmethod
    def from_json(t: dict) -> 'Taskdef':
        return Taskdef(
            image=t['Image'],
            input=t['Input'],
            time=datetime.strptime(t['Time'].split(".")[0] + "Z", "%Y-%m-%dT%H:%M:%SZ"),
        )

    @staticmethod
    def from_env():
        taskjson = os.getenv('COWAIT_TASK')
        taskdef = json.loads(taskjson)
        return Taskdef.from_json(taskdef)

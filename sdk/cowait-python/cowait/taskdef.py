from datetime import datetime


class Taskdef:
    id: str
    image: str
    input: dict
    time: datetime

    def __init__(self, id, image, input, time):
        self.id = id
        self.image = image
        self.input = input
        self.time = time

    @staticmethod
    def from_json(t: dict) -> 'Taskdef':
        return Taskdef(
            id=t['ID'],
            image=t['Image'],
            input=t['Input'],
            time=datetime.strptime(t['Time'].split(".")[0] + "Z", "%Y-%m-%dT%H:%M:%SZ"),
        )

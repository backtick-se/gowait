from datetime import datetime


class Taskdef:
    id: str
    name: str
    image: str
    input: dict
    time: datetime

    def __init__(self, id, name, image, input, time):
        self.id = id
        self.name = name
        self.image = image
        self.input = input
        self.time = time

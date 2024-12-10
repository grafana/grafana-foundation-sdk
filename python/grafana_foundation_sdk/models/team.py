# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Team:
    email: typing.Optional[str]
    name: str

    def __init__(self, email: typing.Optional[str] = None, name: str = ""):
        self.email = email
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
        }
        if self.email is not None:
            payload["email"] = self.email
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "email" in data:
            args["email"] = data["email"]
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)

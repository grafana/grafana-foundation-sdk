# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Team:
    # Name of the team.
    name: str
    # Email of the team.
    email: typing.Optional[str]

    def __init__(self, name: str = "", email: typing.Optional[str] = None):
        self.name = name
        self.email = email

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
        
        if "name" in data:
            args["name"] = data["name"]
        if "email" in data:
            args["email"] = data["email"]        

        return cls(**args)

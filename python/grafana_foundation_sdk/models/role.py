# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Role:
    # The role identifier `managed:builtins:editor:permissions`
    name: str
    # Optional display
    display_name: typing.Optional[str]
    # Name of the team.
    group_name: typing.Optional[str]
    # Role description
    description: typing.Optional[str]
    # Do not show this role
    hidden: bool

    def __init__(self, name: str = "", display_name: typing.Optional[str] = None, group_name: typing.Optional[str] = None, description: typing.Optional[str] = None, hidden: bool = False):
        self.name = name
        self.display_name = display_name
        self.group_name = group_name
        self.description = description
        self.hidden = hidden

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "hidden": self.hidden,
        }
        if self.display_name is not None:
            payload["displayName"] = self.display_name
        if self.group_name is not None:
            payload["groupName"] = self.group_name
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "displayName" in data:
            args["display_name"] = data["displayName"]
        if "groupName" in data:
            args["group_name"] = data["groupName"]
        if "description" in data:
            args["description"] = data["description"]
        if "hidden" in data:
            args["hidden"] = data["hidden"]        

        return cls(**args)

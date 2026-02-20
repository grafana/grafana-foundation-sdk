# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Resource:
    group: str
    kind: str
    # The set of resources
    # +listType=set
    names: list[str]

    def __init__(self, group: str = "", kind: str = "", names: typing.Optional[list[str]] = None) -> None:
        self.group = group
        self.kind = kind
        self.names = names if names is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "group": self.group,
            "kind": self.kind,
            "names": self.names,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "group" in data:
            args["group"] = data["group"]
        if "kind" in data:
            args["kind"] = data["kind"]
        if "names" in data:
            args["names"] = data["names"]        

        return cls(**args)


class Stars:
    resource: list['Resource']

    def __init__(self, resource: typing.Optional[list['Resource']] = None) -> None:
        self.resource = resource if resource is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "resource": self.resource,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "resource" in data:
            args["resource"] = [Resource.from_json(item) for item in data["resource"]]        

        return cls(**args)


StarsKind: typing.Literal["Stars"] = "Stars"




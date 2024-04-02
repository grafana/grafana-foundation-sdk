# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class RoleBinding:
    # The role we are discussing
    role: typing.Union['BuiltinRoleRef', 'CustomRoleRef']
    # The team or user that has the specified role
    subject: 'RoleBindingSubject'

    def __init__(self, role: typing.Union['BuiltinRoleRef', 'CustomRoleRef'] = BuiltinRoleRef(), subject: typing.Optional['RoleBindingSubject'] = None):
        self.role = role
        self.subject = subject if subject is not None else RoleBindingSubject()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "role": self.role,
            "subject": self.subject,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "role" in data:
            decoding_map: dict[str, typing.Union[typing.Type[BuiltinRoleRef], typing.Type[CustomRoleRef]]] = {"BuiltinRole": BuiltinRoleRef, "Role": CustomRoleRef}
            args["role"] = decoding_map[data["role"]["kind"]].from_json(data["role"])
        if "subject" in data:
            args["subject"] = RoleBindingSubject.from_json(data["subject"])        

        return cls(**args)


class CustomRoleRef:
    kind: typing.Literal["Role"]
    name: str

    def __init__(self, name: str = ""):
        self.kind = "Role"
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "name": self.name,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class BuiltinRoleRef:
    kind: typing.Literal["BuiltinRole"]
    name: typing.Literal["viewer", "editor", "admin"]

    def __init__(self, name: typing.Optional[typing.Literal["viewer", "editor", "admin"]] = None):
        self.kind = "BuiltinRole"
        self.name = name if name is not None else "viewer"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "name": self.name,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class RoleBindingSubject:
    kind: typing.Literal["Team", "User"]
    # The team/user identifier name
    name: str

    def __init__(self, kind: typing.Optional[typing.Literal["Team", "User"]] = None, name: str = ""):
        self.kind = kind if kind is not None else "Team"
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "name": self.name,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "kind" in data:
            args["kind"] = data["kind"]
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)




# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class AccessPolicy:
    # The scope where these policies should apply
    scope: 'ResourceRef'
    # The role that must apply this policy
    role: 'RoleRef'
    # The set of rules to apply.  Note that * is required to modify
    # access policy rules, and that "none" will reject all actions
    rules: list['AccessRule']

    def __init__(self, scope: typing.Optional['ResourceRef'] = None, role: typing.Optional['RoleRef'] = None, rules: typing.Optional[list['AccessRule']] = None):
        self.scope = scope if scope is not None else ResourceRef()
        self.role = role if role is not None else RoleRef()
        self.rules = rules if rules is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "scope": self.scope,
            "role": self.role,
            "rules": self.rules,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "scope" in data:
            args["scope"] = ResourceRef.from_json(data["scope"])
        if "role" in data:
            args["role"] = RoleRef.from_json(data["role"])
        if "rules" in data:
            args["rules"] = data["rules"]        

        return cls(**args)


class RoleRef:
    # Policies can apply to roles, teams, or users
    # Applying policies to individual users is supported, but discouraged
    kind: typing.Literal["Role", "BuiltinRole", "Team", "User"]
    name: str
    xname: str

    def __init__(self, kind: typing.Optional[typing.Literal["Role", "BuiltinRole", "Team", "User"]] = None, name: str = "", xname: str = ""):
        self.kind = kind if kind is not None else "Role"
        self.name = name
        self.xname = xname

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "name": self.name,
            "xname": self.xname,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "kind" in data:
            args["kind"] = data["kind"]
        if "name" in data:
            args["name"] = data["name"]
        if "xname" in data:
            args["xname"] = data["xname"]        

        return cls(**args)


class ResourceRef:
    kind: str
    name: str

    def __init__(self, kind: str = "", name: str = ""):
        self.kind = kind
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


class AccessRule:
    # The kind this rule applies to (dashboards, alert, etc)
    kind: str
    # READ, WRITE, CREATE, DELETE, ...
    # should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
    verb: typing.Union[typing.Literal["*"]]
    # Specific sub-elements like "alert.rules" or "dashboard.permissions"????
    target: typing.Optional[str]

    def __init__(self, kind: str = "*", verb: typing.Union[typing.Literal["*"]] = "*", target: typing.Optional[str] = None):
        self.kind = kind
        self.verb = verb
        self.target = target

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "verb": self.verb,
        }
        if self.target is not None:
            payload["target"] = self.target
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "kind" in data:
            args["kind"] = data["kind"]
        if "verb" in data:
            args["verb"] = data["verb"]
        if "target" in data:
            args["target"] = data["target"]        

        return cls(**args)




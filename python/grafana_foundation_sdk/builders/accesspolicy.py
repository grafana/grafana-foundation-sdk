# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import accesspolicy


class AccessPolicy(cogbuilder.Builder[accesspolicy.AccessPolicy]):    
    _internal: accesspolicy.AccessPolicy

    def __init__(self):
        self._internal = accesspolicy.AccessPolicy()

    def build(self) -> accesspolicy.AccessPolicy:
        """
        Builds the object.
        """
        return self._internal    
    
    def scope(self, scope: cogbuilder.Builder[accesspolicy.ResourceRef]) -> typing.Self:    
        """
        The scope where these policies should apply
        """
            
        scope_resource = scope.build()
        self._internal.scope = scope_resource
    
        return self
    
    def role(self, role: cogbuilder.Builder[accesspolicy.RoleRef]) -> typing.Self:    
        """
        The role that must apply this policy
        """
            
        role_resource = role.build()
        self._internal.role = role_resource
    
        return self
    
    def rules(self, rule: cogbuilder.Builder[accesspolicy.AccessRule]) -> typing.Self:    
        """
        The set of rules to apply.  Note that * is required to modify
        access policy rules, and that "none" will reject all actions
        """
            
        if self._internal.rules is None:
            self._internal.rules = []
        
        rule_resource = rule.build()
        self._internal.rules.append(rule_resource)
    
        return self
    

class RoleRef(cogbuilder.Builder[accesspolicy.RoleRef]):    
    _internal: accesspolicy.RoleRef

    def __init__(self):
        self._internal = accesspolicy.RoleRef()

    def build(self) -> accesspolicy.RoleRef:
        """
        Builds the object.
        """
        return self._internal    
    
    def kind(self, kind: typing.Literal["Role", "BuiltinRole", "Team", "User"]) -> typing.Self:    
        """
        Policies can apply to roles, teams, or users
        Applying policies to individual users is supported, but discouraged
        """
            
        self._internal.kind = kind
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def xname(self, xname: str) -> typing.Self:        
        self._internal.xname = xname
    
        return self
    

class ResourceRef(cogbuilder.Builder[accesspolicy.ResourceRef]):    
    _internal: accesspolicy.ResourceRef

    def __init__(self):
        self._internal = accesspolicy.ResourceRef()

    def build(self) -> accesspolicy.ResourceRef:
        """
        Builds the object.
        """
        return self._internal    
    
    def kind(self, kind: str) -> typing.Self:        
        self._internal.kind = kind
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    

class AccessRule(cogbuilder.Builder[accesspolicy.AccessRule]):    
    _internal: accesspolicy.AccessRule

    def __init__(self):
        self._internal = accesspolicy.AccessRule()

    def build(self) -> accesspolicy.AccessRule:
        """
        Builds the object.
        """
        return self._internal    
    
    def kind(self, kind: str) -> typing.Self:    
        """
        The kind this rule applies to (dashboards, alert, etc)
        """
            
        self._internal.kind = kind
    
        return self
    
    def verb(self, verb: typing.Union[typing.Literal["*"]]) -> typing.Self:    
        """
        READ, WRITE, CREATE, DELETE, ...
        should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
        """
            
        self._internal.verb = verb
    
        return self
    
    def target(self, target: str) -> typing.Self:    
        """
        Specific sub-elements like "alert.rules" or "dashboard.permissions"????
        """
            
        self._internal.target = target
    
        return self
    
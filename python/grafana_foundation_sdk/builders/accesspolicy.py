# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import accesspolicy


class AccessPolicy(cogbuilder.Builder[accesspolicy.AccessPolicy]):    
    __internal: accesspolicy.AccessPolicy

    def __init__(self):
        self.__internal = accesspolicy.AccessPolicy()

    def build(self) -> accesspolicy.AccessPolicy:
        return self.__internal    
    
    def scope(self, scope: cogbuilder.Builder[accesspolicy.ResourceRef]) -> typing.Self:    
        """
        The scope where these policies should apply
        """
            
        scope_resource = scope.build()
        self.__internal.scope = scope_resource
    
        return self
    
    def role(self, role: cogbuilder.Builder[accesspolicy.RoleRef]) -> typing.Self:    
        """
        The role that must apply this policy
        """
            
        role_resource = role.build()
        self.__internal.role = role_resource
    
        return self
    
    def rules(self, rules: cogbuilder.Builder[accesspolicy.AccessRule]) -> typing.Self:    
        """
        The set of rules to apply.  Note that * is required to modify
        access policy rules, and that "none" will reject all actions
        """
            
        if self.__internal.rules is None:
            self.__internal.rules = []
        
        rules_resource = rules.build()
        self.__internal.rules.append(rules_resource)
    
        return self
    

class RoleRef(cogbuilder.Builder[accesspolicy.RoleRef]):    
    __internal: accesspolicy.RoleRef

    def __init__(self):
        self.__internal = accesspolicy.RoleRef()

    def build(self) -> accesspolicy.RoleRef:
        return self.__internal    
    
    def kind(self, kind: typing.Literal["Role", "BuiltinRole", "Team", "User"]) -> typing.Self:    
        """
        Policies can apply to roles, teams, or users
        Applying policies to individual users is supported, but discouraged
        """
            
        self.__internal.kind = kind
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self.__internal.name = name
    
        return self
    
    def xname(self, xname: str) -> typing.Self:        
        self.__internal.xname = xname
    
        return self
    

class ResourceRef(cogbuilder.Builder[accesspolicy.ResourceRef]):    
    __internal: accesspolicy.ResourceRef

    def __init__(self):
        self.__internal = accesspolicy.ResourceRef()

    def build(self) -> accesspolicy.ResourceRef:
        return self.__internal    
    
    def kind(self, kind: str) -> typing.Self:        
        self.__internal.kind = kind
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self.__internal.name = name
    
        return self
    

class AccessRule(cogbuilder.Builder[accesspolicy.AccessRule]):    
    __internal: accesspolicy.AccessRule

    def __init__(self):
        self.__internal = accesspolicy.AccessRule()

    def build(self) -> accesspolicy.AccessRule:
        return self.__internal    
    
    def kind(self, kind: typing.Union[typing.Literal["*"]]) -> typing.Self:    
        """
        The kind this rule applies to (dashboards, alert, etc)
        """
            
        self.__internal.kind = kind
    
        return self
    
    def verb(self, verb: typing.Union[typing.Literal["*"]]) -> typing.Self:    
        """
        READ, WRITE, CREATE, DELETE, ...
        should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
        """
            
        self.__internal.verb = verb
    
        return self
    
    def target(self, target: str) -> typing.Self:    
        """
        Specific sub-elements like "alert.rules" or "dashboard.permissions"????
        """
            
        self.__internal.target = target
    
        return self
    
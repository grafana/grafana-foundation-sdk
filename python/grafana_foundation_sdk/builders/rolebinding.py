# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import rolebinding


class RoleBinding(cogbuilder.Builder[rolebinding.RoleBinding]):    
    _internal: rolebinding.RoleBinding

    def __init__(self):
        self._internal = rolebinding.RoleBinding()

    def build(self) -> rolebinding.RoleBinding:
        return self._internal    
    
    def role(self, role: typing.Union[cogbuilder.Builder[rolebinding.BuiltinRoleRef], cogbuilder.Builder[rolebinding.CustomRoleRef]]) -> typing.Self:    
        """
        The role we are discussing
        """
            
        role_resource = role.build()
        self._internal.role = role_resource
    
        return self
    
    def subject(self, subject: cogbuilder.Builder[rolebinding.RoleBindingSubject]) -> typing.Self:    
        """
        The team or user that has the specified role
        """
            
        subject_resource = subject.build()
        self._internal.subject = subject_resource
    
        return self
    

class CustomRoleRef(cogbuilder.Builder[rolebinding.CustomRoleRef]):    
    _internal: rolebinding.CustomRoleRef

    def __init__(self):
        self._internal = rolebinding.CustomRoleRef()        
        self._internal.kind = "Role"

    def build(self) -> rolebinding.CustomRoleRef:
        return self._internal    
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    

class BuiltinRoleRef(cogbuilder.Builder[rolebinding.BuiltinRoleRef]):    
    _internal: rolebinding.BuiltinRoleRef

    def __init__(self):
        self._internal = rolebinding.BuiltinRoleRef()        
        self._internal.kind = "BuiltinRole"

    def build(self) -> rolebinding.BuiltinRoleRef:
        return self._internal    
    
    def name(self, name: typing.Literal["viewer", "editor", "admin"]) -> typing.Self:        
        self._internal.name = name
    
        return self
    

class RoleBindingSubject(cogbuilder.Builder[rolebinding.RoleBindingSubject]):    
    _internal: rolebinding.RoleBindingSubject

    def __init__(self):
        self._internal = rolebinding.RoleBindingSubject()

    def build(self) -> rolebinding.RoleBindingSubject:
        return self._internal    
    
    def kind(self, kind: typing.Literal["Team", "User"]) -> typing.Self:        
        self._internal.kind = kind
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        """
        The team/user identifier name
        """
            
        self._internal.name = name
    
        return self
    
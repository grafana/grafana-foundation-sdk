# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import rolebinding


class RoleBinding(cogbuilder.Builder[rolebinding.RoleBinding]):    
    __internal: rolebinding.RoleBinding

    def __init__(self):
        self.__internal = rolebinding.RoleBinding()

    def build(self) -> rolebinding.RoleBinding:
        return self.__internal    
    
    def role(self, role: typing.Union[cogbuilder.Builder[rolebinding.BuiltinRoleRef], cogbuilder.Builder[rolebinding.CustomRoleRef]]) -> typing.Self:    
        """
        The role we are discussing
        """
            
        role_resource = role.build()
        self.__internal.role = role_resource
    
        return self
    
    def subject(self, subject: cogbuilder.Builder[rolebinding.RoleBindingSubject]) -> typing.Self:    
        """
        The team or user that has the specified role
        """
            
        subject_resource = subject.build()
        self.__internal.subject = subject_resource
    
        return self
    

class CustomRoleRef(cogbuilder.Builder[rolebinding.CustomRoleRef]):    
    __internal: rolebinding.CustomRoleRef

    def __init__(self):
        self.__internal = rolebinding.CustomRoleRef()        
        self.__internal.kind = "Role"

    def build(self) -> rolebinding.CustomRoleRef:
        return self.__internal    
    
    def name(self, name: str) -> typing.Self:        
        self.__internal.name = name
    
        return self
    

class BuiltinRoleRef(cogbuilder.Builder[rolebinding.BuiltinRoleRef]):    
    __internal: rolebinding.BuiltinRoleRef

    def __init__(self):
        self.__internal = rolebinding.BuiltinRoleRef()        
        self.__internal.kind = "BuiltinRole"

    def build(self) -> rolebinding.BuiltinRoleRef:
        return self.__internal    
    
    def name(self, name: typing.Literal["viewer", "editor", "admin"]) -> typing.Self:        
        self.__internal.name = name
    
        return self
    

class RoleBindingSubject(cogbuilder.Builder[rolebinding.RoleBindingSubject]):    
    __internal: rolebinding.RoleBindingSubject

    def __init__(self):
        self.__internal = rolebinding.RoleBindingSubject()

    def build(self) -> rolebinding.RoleBindingSubject:
        return self.__internal    
    
    def kind(self, kind: typing.Literal["Team", "User"]) -> typing.Self:        
        self.__internal.kind = kind
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        """
        The team/user identifier name
        """
            
        self.__internal.name = name
    
        return self
    
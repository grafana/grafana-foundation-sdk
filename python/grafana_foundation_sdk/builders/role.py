# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import role


class Role(cogbuilder.Builder[role.Role]):    
    __internal: role.Role

    def __init__(self):
        self.__internal = role.Role()

    def build(self) -> role.Role:
        return self.__internal    
    
    def name(self, name: str) -> typing.Self:    
        """
        The role identifier `managed:builtins:editor:permissions`
        """
            
        self.__internal.name = name
    
        return self
    
    def display_name(self, display_name: str) -> typing.Self:    
        """
        Optional display
        """
            
        self.__internal.display_name = display_name
    
        return self
    
    def group_name(self, group_name: str) -> typing.Self:    
        """
        Name of the team.
        """
            
        self.__internal.group_name = group_name
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Role description
        """
            
        self.__internal.description = description
    
        return self
    
    def hidden(self, hidden: typing.Union[bool]) -> typing.Self:    
        """
        Do not show this role
        """
            
        self.__internal.hidden = hidden
    
        return self
    
# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import role


class Role(cogbuilder.Builder[role.Role]):    
    _internal: role.Role

    def __init__(self):
        self._internal = role.Role()

    def build(self) -> role.Role:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        """
        The role identifier `managed:builtins:editor:permissions`
        """
            
        self._internal.name = name
    
        return self
    
    def display_name(self, display_name: str) -> typing.Self:    
        """
        Optional display
        """
            
        self._internal.display_name = display_name
    
        return self
    
    def group_name(self, group_name: str) -> typing.Self:    
        """
        Name of the team.
        """
            
        self._internal.group_name = group_name
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Role description
        """
            
        self._internal.description = description
    
        return self
    
    def hidden(self, hidden: bool) -> typing.Self:    
        """
        Do not show this role
        """
            
        self._internal.hidden = hidden
    
        return self
    
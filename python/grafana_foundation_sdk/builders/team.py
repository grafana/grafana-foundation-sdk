# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import team


class Team(cogbuilder.Builder[team.Team]):    
    _internal: team.Team

    def __init__(self, name: str):
        self._internal = team.Team()        
        self._internal.name = name

    def build(self) -> team.Team:
        """
        Builds the object.
        """
        return self._internal    
    
    def email(self, email: str) -> typing.Self:        
        self._internal.email = email
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import team


class Team(cogbuilder.Builder[team.Team]):    
    __internal: team.Team

    def __init__(self, name: str):
        self.__internal = team.Team()        
        self.__internal.name = name

    def build(self) -> team.Team:
        return self.__internal    
    
    def name(self, name: str) -> typing.Self:    
        """
        Name of the team.
        """
            
        self.__internal.name = name
    
        return self
    
    def email(self, email: str) -> typing.Self:    
        """
        Email of the team.
        """
            
        self.__internal.email = email
    
        return self
    
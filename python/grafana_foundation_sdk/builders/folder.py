# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import folder


class Folder(cogbuilder.Builder[folder.Folder]):    
    """
    TODO:
    common metadata will soon support setting the parent folder in the metadata
    """
    
    __internal: folder.Folder

    def __init__(self, title: str):
        self.__internal = folder.Folder()        
        self.__internal.title = title

    def build(self) -> folder.Folder:
        return self.__internal    
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Unique folder id. (will be k8s name)
        """
            
        self.__internal.uid = uid
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        """
        Folder title
        """
            
        self.__internal.title = title
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Description of the folder.
        """
            
        self.__internal.description = description
    
        return self
    
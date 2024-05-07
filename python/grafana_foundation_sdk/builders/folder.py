# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import folder


class Folder(cogbuilder.Builder[folder.Folder]):    
    """
    TODO:
    common metadata will soon support setting the parent folder in the metadata
    """
    
    _internal: folder.Folder

    def __init__(self):
        self._internal = folder.Folder()

    def build(self) -> folder.Folder:
        return self._internal    
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Unique folder id. (will be k8s name)
        """
            
        self._internal.uid = uid
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        """
        Folder title
        """
            
        self._internal.title = title
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Description of the folder.
        """
            
        self._internal.description = description
    
        return self
    
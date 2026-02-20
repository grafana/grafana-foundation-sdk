# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import folderv1beta1
from ..builders import resource as resource_builder
from ..models import resource


class Folder(cogbuilder.Builder[folderv1beta1.Folder]):
    _internal: folderv1beta1.Folder

    def __init__(self, title: str) -> None:
        self._internal = folderv1beta1.Folder()        
        self._internal.title = title

    def build(self) -> folderv1beta1.Folder:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        self._internal.description = description
    
        return self
    

"""
Creates a resource manifest from a Folder.
"""
def manifest(name: str, folder: cogbuilder.Builder[folderv1beta1.Folder]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("folder.grafana.app/v1beta1")
    builder.kind("Folder")
    builder.metadata(resource_builder.named(name))
    builder.spec(folder.build())

    return builder

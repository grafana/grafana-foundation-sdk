# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import starsv1alpha1
from ..builders import resource as resource_builder
from ..models import resource


class Resource(cogbuilder.Builder[starsv1alpha1.Resource]):
    _internal: starsv1alpha1.Resource

    def __init__(self) -> None:
        self._internal = starsv1alpha1.Resource()

    def build(self) -> starsv1alpha1.Resource:
        """
        Builds the object.
        """
        return self._internal    
    
    def group(self, group: str) -> typing.Self:    
        self._internal.group = group
    
        return self
    
    def kind(self, kind: str) -> typing.Self:    
        self._internal.kind = kind
    
        return self
    
    def names(self, names: list[str]) -> typing.Self:    
        """
        The set of resources
        +listType=set
        """
            
        self._internal.names = names
    
        return self
    


class Stars(cogbuilder.Builder[starsv1alpha1.Stars]):
    _internal: starsv1alpha1.Stars

    def __init__(self) -> None:
        self._internal = starsv1alpha1.Stars()

    def build(self) -> starsv1alpha1.Stars:
        """
        Builds the object.
        """
        return self._internal    
    
    def resources(self, resource: list[cogbuilder.Builder[starsv1alpha1.Resource]]) -> typing.Self:    
        resource_resources = [r1.build() for r1 in resource]
        self._internal.resource = resource_resources
    
        return self
    
    def resource(self, resource: cogbuilder.Builder[starsv1alpha1.Resource]) -> typing.Self:    
        if self._internal.resource is None:
            self._internal.resource = []
        
        resource_resource = resource.build()
        self._internal.resource.append(resource_resource)
    
        return self
    

"""
Creates a resource manifest from Stars.
"""
def manifest(name: str, stars: cogbuilder.Builder[starsv1alpha1.Stars]) -> cogbuilder.Builder[resource.Manifest]:
    builder = resource_builder.Manifest()
    builder.api_version("preferences.grafana.app/v1alpha1")
    builder.metadata(resource_builder.named(name))
    builder.kind("Stars")
    builder.spec(stars.build())

    return builder

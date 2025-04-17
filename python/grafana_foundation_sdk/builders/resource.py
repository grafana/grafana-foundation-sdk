# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import resource


class Manifest(cogbuilder.Builder[resource.Manifest]):
    _internal: resource.Manifest

    def __init__(self):
        self._internal = resource.Manifest()

    def build(self) -> resource.Manifest:
        """
        Builds the object.
        """
        return self._internal    
    
    def api_version(self, api_version: str) -> typing.Self:    
        self._internal.api_version = api_version
    
        return self
    
    def kind(self, kind: str) -> typing.Self:    
        self._internal.kind = kind
    
        return self
    
    def metadata(self, metadata: cogbuilder.Builder[resource.Metadata]) -> typing.Self:    
        metadata_resource = metadata.build()
        self._internal.metadata = metadata_resource
    
        return self
    
    def spec(self, spec: object) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    


class Metadata(cogbuilder.Builder[resource.Metadata]):
    _internal: resource.Metadata

    def __init__(self):
        self._internal = resource.Metadata()

    def build(self) -> resource.Metadata:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    
    def namespace(self, namespace: str) -> typing.Self:    
        self._internal.namespace = namespace
    
        return self
    
    def labels(self, labels: dict[str, str]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def annotations(self, annotations: dict[str, str]) -> typing.Self:    
        self._internal.annotations = annotations
    
        return self
    
    def uid(self, uid: str) -> typing.Self:    
        self._internal.uid = uid
    
        return self
    
    def resource_version(self, resource_version: str) -> typing.Self:    
        self._internal.resource_version = resource_version
    
        return self
    
    def generation(self, generation: int) -> typing.Self:    
        self._internal.generation = generation
    
        return self
    
    def creation_timestamp(self, creation_timestamp: str) -> typing.Self:    
        self._internal.creation_timestamp = creation_timestamp
    
        return self
    
    def update_timestamp(self, update_timestamp: str) -> typing.Self:    
        self._internal.update_timestamp = update_timestamp
    
        return self
    
    def deletion_timestamp(self, deletion_timestamp: str) -> typing.Self:    
        self._internal.deletion_timestamp = deletion_timestamp
    
        return self
    

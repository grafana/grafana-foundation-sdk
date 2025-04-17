# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Manifest:
    api_version: str
    kind: str
    metadata: 'Metadata'
    spec: object

    def __init__(self, api_version: str = "", kind: str = "", metadata: typing.Optional['Metadata'] = None, spec: object = None):
        self.api_version = api_version
        self.kind = kind
        self.metadata = metadata if metadata is not None else Metadata()
        self.spec = spec

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "apiVersion": self.api_version,
            "kind": self.kind,
            "metadata": self.metadata,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "apiVersion" in data:
            args["api_version"] = data["apiVersion"]
        if "kind" in data:
            args["kind"] = data["kind"]
        if "metadata" in data:
            args["metadata"] = Metadata.from_json(data["metadata"])
        if "spec" in data:
            args["spec"] = data["spec"]        

        return cls(**args)


class Metadata:
    name: str
    namespace: typing.Optional[str]
    labels: typing.Optional[dict[str, str]]
    annotations: typing.Optional[dict[str, str]]
    uid: typing.Optional[str]
    resource_version: typing.Optional[str]
    generation: typing.Optional[int]
    creation_timestamp: typing.Optional[str]
    update_timestamp: typing.Optional[str]
    deletion_timestamp: typing.Optional[str]

    def __init__(self, name: str = "", namespace: typing.Optional[str] = None, labels: typing.Optional[dict[str, str]] = None, annotations: typing.Optional[dict[str, str]] = None, uid: typing.Optional[str] = None, resource_version: typing.Optional[str] = None, generation: typing.Optional[int] = None, creation_timestamp: typing.Optional[str] = None, update_timestamp: typing.Optional[str] = None, deletion_timestamp: typing.Optional[str] = None):
        self.name = name
        self.namespace = namespace
        self.labels = labels
        self.annotations = annotations
        self.uid = uid
        self.resource_version = resource_version
        self.generation = generation
        self.creation_timestamp = creation_timestamp
        self.update_timestamp = update_timestamp
        self.deletion_timestamp = deletion_timestamp

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
        }
        if self.namespace is not None:
            payload["namespace"] = self.namespace
        if self.labels is not None:
            payload["labels"] = self.labels
        if self.annotations is not None:
            payload["annotations"] = self.annotations
        if self.uid is not None:
            payload["uid"] = self.uid
        if self.resource_version is not None:
            payload["resourceVersion"] = self.resource_version
        if self.generation is not None:
            payload["generation"] = self.generation
        if self.creation_timestamp is not None:
            payload["creationTimestamp"] = self.creation_timestamp
        if self.update_timestamp is not None:
            payload["updateTimestamp"] = self.update_timestamp
        if self.deletion_timestamp is not None:
            payload["deletionTimestamp"] = self.deletion_timestamp
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "namespace" in data:
            args["namespace"] = data["namespace"]
        if "labels" in data:
            args["labels"] = data["labels"]
        if "annotations" in data:
            args["annotations"] = data["annotations"]
        if "uid" in data:
            args["uid"] = data["uid"]
        if "resourceVersion" in data:
            args["resource_version"] = data["resourceVersion"]
        if "generation" in data:
            args["generation"] = data["generation"]
        if "creationTimestamp" in data:
            args["creation_timestamp"] = data["creationTimestamp"]
        if "updateTimestamp" in data:
            args["update_timestamp"] = data["updateTimestamp"]
        if "deletionTimestamp" in data:
            args["deletion_timestamp"] = data["deletionTimestamp"]        

        return cls(**args)




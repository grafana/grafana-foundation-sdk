# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Folder:
    title: str
    description: typing.Optional[str]

    def __init__(self, title: str = "", description: typing.Optional[str] = None) -> None:
        self.title = title
        self.description = description

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "title": self.title,
        }
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "description" in data:
            args["description"] = data["description"]        

        return cls(**args)


FolderV1Beta1: typing.Literal["folder.grafana.app/v1beta1"] = "folder.grafana.app/v1beta1"


FolderKind: typing.Literal["Folder"] = "Folder"




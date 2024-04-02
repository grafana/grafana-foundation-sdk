# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Folder:
    """
    TODO:
    common metadata will soon support setting the parent folder in the metadata
    """

    # Unique folder id. (will be k8s name)
    uid: str
    # Folder title
    title: str
    # Description of the folder.
    description: typing.Optional[str]

    def __init__(self, uid: str = "", title: str = "", description: typing.Optional[str] = None):
        self.uid = uid
        self.title = title
        self.description = description

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "uid": self.uid,
            "title": self.title,
        }
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "uid" in data:
            args["uid"] = data["uid"]
        if "title" in data:
            args["title"] = data["title"]
        if "description" in data:
            args["description"] = data["description"]        

        return cls(**args)

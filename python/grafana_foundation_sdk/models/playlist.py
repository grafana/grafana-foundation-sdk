# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Playlist:
    # Unique playlist identifier. Generated on creation, either by the
    # creator of the playlist of by the application.
    uid: str
    # Name of the playlist.
    name: str
    # Interval sets the time between switching views in a playlist.
    # FIXME: Is this based on a standardized format or what options are available? Can datemath be used?
    interval: str
    # The ordered list of items that the playlist will iterate over.
    # FIXME! This should not be optional, but changing it makes the godegen awkward
    items: typing.Optional[list['PlaylistItem']]

    def __init__(self, uid: str = "", name: str = "", interval: str = "5m", items: typing.Optional[list['PlaylistItem']] = None):
        self.uid = uid
        self.name = name
        self.interval = interval
        self.items = items

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "uid": self.uid,
            "name": self.name,
            "interval": self.interval,
        }
        if self.items is not None:
            payload["items"] = self.items
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "uid" in data:
            args["uid"] = data["uid"]
        if "name" in data:
            args["name"] = data["name"]
        if "interval" in data:
            args["interval"] = data["interval"]
        if "items" in data:
            args["items"] = data["items"]        

        return cls(**args)


class PlaylistItem:
    # Type of the item.
    type_val: typing.Literal["dashboard_by_uid", "dashboard_by_id", "dashboard_by_tag"]
    # Value depends on type and describes the playlist item.
    # 
    #  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
    #  is not portable as the numerical identifier is non-deterministic between different instances.
    #  Will be replaced by dashboard_by_uid in the future. (deprecated)
    #  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
    #  dashboards behind the tag will be added to the playlist.
    #  - dashboard_by_uid: The value is the dashboard UID
    value: str
    # Title is an unused property -- it will be removed in the future
    title: typing.Optional[str]

    def __init__(self, type_val: typing.Optional[typing.Literal["dashboard_by_uid", "dashboard_by_id", "dashboard_by_tag"]] = None, value: str = "", title: typing.Optional[str] = None):
        self.type_val = type_val if type_val is not None else "dashboard_by_uid"
        self.value = value
        self.title = title

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "value": self.value,
        }
        if self.title is not None:
            payload["title"] = self.title
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "value" in data:
            args["value"] = data["value"]
        if "title" in data:
            args["title"] = data["title"]        

        return cls(**args)




# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class Item:
    # type of the item.
    type_val: typing.Literal["dashboard_by_tag", "dashboard_by_uid", "dashboard_by_id"]
    # Value depends on type and describes the playlist item.
    #  - dashboard_by_id: The value is an internal numerical identifier set by Grafana. This
    #  is not portable as the numerical identifier is non-deterministic between different instances.
    #  Will be replaced by dashboard_by_uid in the future. (deprecated)
    #  - dashboard_by_tag: The value is a tag which is set on any number of dashboards. All
    #  dashboards behind the tag will be added to the playlist.
    #  - dashboard_by_uid: The value is the dashboard UID
    value: str

    def __init__(self, type_val: typing.Optional[typing.Literal["dashboard_by_tag", "dashboard_by_uid", "dashboard_by_id"]] = None, value: str = "") -> None:
        self.type_val = type_val if type_val is not None else "dashboard_by_tag"
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class Playlist:
    title: str
    interval: str
    items: list['Item']

    def __init__(self, title: str = "", interval: str = "", items: typing.Optional[list['Item']] = None) -> None:
        self.title = title
        self.interval = interval
        self.items = items if items is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "title": self.title,
            "interval": self.interval,
            "items": self.items,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "interval" in data:
            args["interval"] = data["interval"]
        if "items" in data:
            args["items"] = [Item.from_json(item) for item in data["items"]]        

        return cls(**args)




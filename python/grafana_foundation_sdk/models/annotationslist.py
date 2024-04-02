# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import runtime as cogruntime


class Options:
    only_from_this_dashboard: bool
    only_in_time_range: bool
    tags: list[str]
    limit: int
    show_user: bool
    show_time: bool
    show_tags: bool
    navigate_to_panel: bool
    navigate_before: str
    navigate_after: str

    def __init__(self, only_from_this_dashboard: bool = False, only_in_time_range: bool = False, tags: typing.Optional[list[str]] = None, limit: int = 10, show_user: bool = True, show_time: bool = True, show_tags: bool = True, navigate_to_panel: bool = True, navigate_before: str = "10m", navigate_after: str = "10m"):
        self.only_from_this_dashboard = only_from_this_dashboard
        self.only_in_time_range = only_in_time_range
        self.tags = tags if tags is not None else []
        self.limit = limit
        self.show_user = show_user
        self.show_time = show_time
        self.show_tags = show_tags
        self.navigate_to_panel = navigate_to_panel
        self.navigate_before = navigate_before
        self.navigate_after = navigate_after

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "onlyFromThisDashboard": self.only_from_this_dashboard,
            "onlyInTimeRange": self.only_in_time_range,
            "tags": self.tags,
            "limit": self.limit,
            "showUser": self.show_user,
            "showTime": self.show_time,
            "showTags": self.show_tags,
            "navigateToPanel": self.navigate_to_panel,
            "navigateBefore": self.navigate_before,
            "navigateAfter": self.navigate_after,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "onlyFromThisDashboard" in data:
            args["only_from_this_dashboard"] = data["onlyFromThisDashboard"]
        if "onlyInTimeRange" in data:
            args["only_in_time_range"] = data["onlyInTimeRange"]
        if "tags" in data:
            args["tags"] = data["tags"]
        if "limit" in data:
            args["limit"] = data["limit"]
        if "showUser" in data:
            args["show_user"] = data["showUser"]
        if "showTime" in data:
            args["show_time"] = data["showTime"]
        if "showTags" in data:
            args["show_tags"] = data["showTags"]
        if "navigateToPanel" in data:
            args["navigate_to_panel"] = data["navigateToPanel"]
        if "navigateBefore" in data:
            args["navigate_before"] = data["navigateBefore"]
        if "navigateAfter" in data:
            args["navigate_after"] = data["navigateAfter"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="annolist",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

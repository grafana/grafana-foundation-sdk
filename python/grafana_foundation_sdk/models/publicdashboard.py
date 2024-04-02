# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing


class PublicDashboard:
    # Unique public dashboard identifier
    uid: str
    # Dashboard unique identifier referenced by this public dashboard
    dashboard_uid: str
    # Unique public access token
    access_token: typing.Optional[str]
    # Flag that indicates if the public dashboard is enabled
    is_enabled: bool
    # Flag that indicates if annotations are enabled
    annotations_enabled: bool
    # Flag that indicates if the time range picker is enabled
    time_selection_enabled: bool

    def __init__(self, uid: str = "", dashboard_uid: str = "", access_token: typing.Optional[str] = None, is_enabled: bool = False, annotations_enabled: bool = False, time_selection_enabled: bool = False):
        self.uid = uid
        self.dashboard_uid = dashboard_uid
        self.access_token = access_token
        self.is_enabled = is_enabled
        self.annotations_enabled = annotations_enabled
        self.time_selection_enabled = time_selection_enabled

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "uid": self.uid,
            "dashboardUid": self.dashboard_uid,
            "isEnabled": self.is_enabled,
            "annotationsEnabled": self.annotations_enabled,
            "timeSelectionEnabled": self.time_selection_enabled,
        }
        if self.access_token is not None:
            payload["accessToken"] = self.access_token
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "uid" in data:
            args["uid"] = data["uid"]
        if "dashboardUid" in data:
            args["dashboard_uid"] = data["dashboardUid"]
        if "accessToken" in data:
            args["access_token"] = data["accessToken"]
        if "isEnabled" in data:
            args["is_enabled"] = data["isEnabled"]
        if "annotationsEnabled" in data:
            args["annotations_enabled"] = data["annotationsEnabled"]
        if "timeSelectionEnabled" in data:
            args["time_selection_enabled"] = data["timeSelectionEnabled"]        

        return cls(**args)

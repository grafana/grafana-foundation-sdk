# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import runtime as cogruntime


class Options:
    # Comma-separated list of values used to filter alert results
    labels: str
    # Name of the alertmanager used as a source for alerts
    alertmanager: str
    # Expand all alert groups by default
    expand_all: bool

    def __init__(self, labels: str = "", alertmanager: str = "", expand_all: bool = False):
        self.labels = labels
        self.alertmanager = alertmanager
        self.expand_all = expand_all

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "labels": self.labels,
            "alertmanager": self.alertmanager,
            "expandAll": self.expand_all,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "labels" in data:
            args["labels"] = data["labels"]
        if "alertmanager" in data:
            args["alertmanager"] = data["alertmanager"]
        if "expandAll" in data:
            args["expand_all"] = data["expandAll"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="alertGroups",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

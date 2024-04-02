# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import runtime as cogruntime


class Options:
    selected_series: int

    def __init__(self, selected_series: int = 0):
        self.selected_series = selected_series

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "selectedSeries": self.selected_series,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "selectedSeries" in data:
            args["selected_series"] = data["selectedSeries"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="datagrid",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    display_mode: common.BarGaugeDisplayMode
    value_mode: common.BarGaugeValueMode
    show_unfilled: bool
    min_viz_width: int
    reduce_options: common.ReduceDataOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    min_viz_height: int
    orientation: common.VizOrientation

    def __init__(self, display_mode: typing.Optional[common.BarGaugeDisplayMode] = None, value_mode: typing.Optional[common.BarGaugeValueMode] = None, show_unfilled: bool = True, min_viz_width: int = 0, reduce_options: typing.Optional[common.ReduceDataOptions] = None, text: typing.Optional[common.VizTextDisplayOptions] = None, min_viz_height: int = 10, orientation: typing.Optional[common.VizOrientation] = None):
        self.display_mode = display_mode if display_mode is not None else common.BarGaugeDisplayMode.GRADIENT
        self.value_mode = value_mode if value_mode is not None else common.BarGaugeValueMode.COLOR
        self.show_unfilled = show_unfilled
        self.min_viz_width = min_viz_width
        self.reduce_options = reduce_options if reduce_options is not None else common.ReduceDataOptions()
        self.text = text
        self.min_viz_height = min_viz_height
        self.orientation = orientation if orientation is not None else common.VizOrientation.AUTO

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "displayMode": self.display_mode,
            "valueMode": self.value_mode,
            "showUnfilled": self.show_unfilled,
            "minVizWidth": self.min_viz_width,
            "reduceOptions": self.reduce_options,
            "minVizHeight": self.min_viz_height,
            "orientation": self.orientation,
        }
        if self.text is not None:
            payload["text"] = self.text
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "displayMode" in data:
            args["display_mode"] = data["displayMode"]
        if "valueMode" in data:
            args["value_mode"] = data["valueMode"]
        if "showUnfilled" in data:
            args["show_unfilled"] = data["showUnfilled"]
        if "minVizWidth" in data:
            args["min_viz_width"] = data["minVizWidth"]
        if "reduceOptions" in data:
            args["reduce_options"] = common.ReduceDataOptions.from_json(data["reduceOptions"])
        if "text" in data:
            args["text"] = common.VizTextDisplayOptions.from_json(data["text"])
        if "minVizHeight" in data:
            args["min_viz_height"] = data["minVizHeight"]
        if "orientation" in data:
            args["orientation"] = data["orientation"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="bargauge",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

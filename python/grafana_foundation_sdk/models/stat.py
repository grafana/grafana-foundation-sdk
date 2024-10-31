# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    graph_mode: common.BigValueGraphMode
    color_mode: common.BigValueColorMode
    justify_mode: common.BigValueJustifyMode
    text_mode: common.BigValueTextMode
    wide_layout: bool
    show_percent_change: bool
    reduce_options: common.ReduceDataOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    percent_change_color_mode: common.PercentChangeColorMode
    orientation: common.VizOrientation

    def __init__(self, graph_mode: typing.Optional[common.BigValueGraphMode] = None, color_mode: typing.Optional[common.BigValueColorMode] = None, justify_mode: typing.Optional[common.BigValueJustifyMode] = None, text_mode: typing.Optional[common.BigValueTextMode] = None, wide_layout: bool = True, show_percent_change: bool = False, reduce_options: typing.Optional[common.ReduceDataOptions] = None, text: typing.Optional[common.VizTextDisplayOptions] = None, percent_change_color_mode: typing.Optional[common.PercentChangeColorMode] = None, orientation: typing.Optional[common.VizOrientation] = None):
        self.graph_mode = graph_mode if graph_mode is not None else common.BigValueGraphMode.AREA
        self.color_mode = color_mode if color_mode is not None else common.BigValueColorMode.VALUE
        self.justify_mode = justify_mode if justify_mode is not None else common.BigValueJustifyMode.AUTO
        self.text_mode = text_mode if text_mode is not None else common.BigValueTextMode.AUTO
        self.wide_layout = wide_layout
        self.show_percent_change = show_percent_change
        self.reduce_options = reduce_options if reduce_options is not None else common.ReduceDataOptions()
        self.text = text
        self.percent_change_color_mode = percent_change_color_mode if percent_change_color_mode is not None else common.PercentChangeColorMode.STANDARD
        self.orientation = orientation if orientation is not None else common.VizOrientation.AUTO

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "graphMode": self.graph_mode,
            "colorMode": self.color_mode,
            "justifyMode": self.justify_mode,
            "textMode": self.text_mode,
            "wideLayout": self.wide_layout,
            "showPercentChange": self.show_percent_change,
            "reduceOptions": self.reduce_options,
            "percentChangeColorMode": self.percent_change_color_mode,
            "orientation": self.orientation,
        }
        if self.text is not None:
            payload["text"] = self.text
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "graphMode" in data:
            args["graph_mode"] = data["graphMode"]
        if "colorMode" in data:
            args["color_mode"] = data["colorMode"]
        if "justifyMode" in data:
            args["justify_mode"] = data["justifyMode"]
        if "textMode" in data:
            args["text_mode"] = data["textMode"]
        if "wideLayout" in data:
            args["wide_layout"] = data["wideLayout"]
        if "showPercentChange" in data:
            args["show_percent_change"] = data["showPercentChange"]
        if "reduceOptions" in data:
            args["reduce_options"] = common.ReduceDataOptions.from_json(data["reduceOptions"])
        if "text" in data:
            args["text"] = common.VizTextDisplayOptions.from_json(data["text"])
        if "percentChangeColorMode" in data:
            args["percent_change_color_mode"] = data["percentChangeColorMode"]
        if "orientation" in data:
            args["orientation"] = data["orientation"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="stat",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

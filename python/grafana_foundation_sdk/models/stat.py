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
    reduce_options: common.ReduceDataOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    show_percent_change: bool
    orientation: common.VizOrientation

    def __init__(self, graph_mode: typing.Optional[common.BigValueGraphMode] = None, color_mode: typing.Optional[common.BigValueColorMode] = None, justify_mode: typing.Optional[common.BigValueJustifyMode] = None, text_mode: typing.Optional[common.BigValueTextMode] = None, wide_layout: bool = True, reduce_options: typing.Optional[common.ReduceDataOptions] = None, text: typing.Optional[common.VizTextDisplayOptions] = None, show_percent_change: bool = False, orientation: typing.Optional[common.VizOrientation] = None):
        self.graph_mode = graph_mode if graph_mode is not None else common.BigValueGraphMode.AREA
        self.color_mode = color_mode if color_mode is not None else common.BigValueColorMode.VALUE
        self.justify_mode = justify_mode if justify_mode is not None else common.BigValueJustifyMode.AUTO
        self.text_mode = text_mode if text_mode is not None else common.BigValueTextMode.AUTO
        self.wide_layout = wide_layout
        self.reduce_options = reduce_options if reduce_options is not None else common.ReduceDataOptions()
        self.text = text
        self.show_percent_change = show_percent_change
        self.orientation = orientation if orientation is not None else common.VizOrientation.AUTO

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "graphMode": self.graph_mode,
            "colorMode": self.color_mode,
            "justifyMode": self.justify_mode,
            "textMode": self.text_mode,
            "wideLayout": self.wide_layout,
            "reduceOptions": self.reduce_options,
            "showPercentChange": self.show_percent_change,
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
        if "reduceOptions" in data:
            args["reduce_options"] = common.ReduceDataOptions.from_json(data["reduceOptions"])
        if "text" in data:
            args["text"] = common.VizTextDisplayOptions.from_json(data["text"])
        if "showPercentChange" in data:
            args["show_percent_change"] = data["showPercentChange"]
        if "orientation" in data:
            args["orientation"] = data["orientation"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="stat",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

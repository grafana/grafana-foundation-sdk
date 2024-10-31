# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    display_mode: common.BarGaugeDisplayMode
    value_mode: common.BarGaugeValueMode
    name_placement: common.BarGaugeNamePlacement
    show_unfilled: bool
    sizing: common.BarGaugeSizing
    min_viz_width: int
    min_viz_height: int
    legend: common.VizLegendOptions
    reduce_options: common.ReduceDataOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    max_viz_height: int
    orientation: common.VizOrientation

    def __init__(self, display_mode: typing.Optional[common.BarGaugeDisplayMode] = None, value_mode: typing.Optional[common.BarGaugeValueMode] = None, name_placement: typing.Optional[common.BarGaugeNamePlacement] = None, show_unfilled: bool = True, sizing: typing.Optional[common.BarGaugeSizing] = None, min_viz_width: int = 8, min_viz_height: int = 16, legend: typing.Optional[common.VizLegendOptions] = None, reduce_options: typing.Optional[common.ReduceDataOptions] = None, text: typing.Optional[common.VizTextDisplayOptions] = None, max_viz_height: int = 300, orientation: typing.Optional[common.VizOrientation] = None):
        self.display_mode = display_mode if display_mode is not None else common.BarGaugeDisplayMode.GRADIENT
        self.value_mode = value_mode if value_mode is not None else common.BarGaugeValueMode.COLOR
        self.name_placement = name_placement if name_placement is not None else common.BarGaugeNamePlacement.AUTO
        self.show_unfilled = show_unfilled
        self.sizing = sizing if sizing is not None else common.BarGaugeSizing.AUTO
        self.min_viz_width = min_viz_width
        self.min_viz_height = min_viz_height
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.reduce_options = reduce_options if reduce_options is not None else common.ReduceDataOptions()
        self.text = text
        self.max_viz_height = max_viz_height
        self.orientation = orientation if orientation is not None else common.VizOrientation.AUTO

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "displayMode": self.display_mode,
            "valueMode": self.value_mode,
            "namePlacement": self.name_placement,
            "showUnfilled": self.show_unfilled,
            "sizing": self.sizing,
            "minVizWidth": self.min_viz_width,
            "minVizHeight": self.min_viz_height,
            "legend": self.legend,
            "reduceOptions": self.reduce_options,
            "maxVizHeight": self.max_viz_height,
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
        if "namePlacement" in data:
            args["name_placement"] = data["namePlacement"]
        if "showUnfilled" in data:
            args["show_unfilled"] = data["showUnfilled"]
        if "sizing" in data:
            args["sizing"] = data["sizing"]
        if "minVizWidth" in data:
            args["min_viz_width"] = data["minVizWidth"]
        if "minVizHeight" in data:
            args["min_viz_height"] = data["minVizHeight"]
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "reduceOptions" in data:
            args["reduce_options"] = common.ReduceDataOptions.from_json(data["reduceOptions"])
        if "text" in data:
            args["text"] = common.VizTextDisplayOptions.from_json(data["text"])
        if "maxVizHeight" in data:
            args["max_viz_height"] = data["maxVizHeight"]
        if "orientation" in data:
            args["orientation"] = data["orientation"]        

        return cls(**args)


def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="bargauge",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

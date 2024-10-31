# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    show_threshold_labels: bool
    show_threshold_markers: bool
    sizing: common.BarGaugeSizing
    min_viz_width: int
    reduce_options: common.ReduceDataOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    min_viz_height: int
    orientation: common.VizOrientation

    def __init__(self, show_threshold_labels: bool = False, show_threshold_markers: bool = True, sizing: typing.Optional[common.BarGaugeSizing] = None, min_viz_width: int = 75, reduce_options: typing.Optional[common.ReduceDataOptions] = None, text: typing.Optional[common.VizTextDisplayOptions] = None, min_viz_height: int = 75, orientation: typing.Optional[common.VizOrientation] = None):
        self.show_threshold_labels = show_threshold_labels
        self.show_threshold_markers = show_threshold_markers
        self.sizing = sizing if sizing is not None else common.BarGaugeSizing.AUTO
        self.min_viz_width = min_viz_width
        self.reduce_options = reduce_options if reduce_options is not None else common.ReduceDataOptions()
        self.text = text
        self.min_viz_height = min_viz_height
        self.orientation = orientation if orientation is not None else common.VizOrientation.AUTO

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "showThresholdLabels": self.show_threshold_labels,
            "showThresholdMarkers": self.show_threshold_markers,
            "sizing": self.sizing,
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
        
        if "showThresholdLabels" in data:
            args["show_threshold_labels"] = data["showThresholdLabels"]
        if "showThresholdMarkers" in data:
            args["show_threshold_markers"] = data["showThresholdMarkers"]
        if "sizing" in data:
            args["sizing"] = data["sizing"]
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
        identifier="gauge",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=None,
    )

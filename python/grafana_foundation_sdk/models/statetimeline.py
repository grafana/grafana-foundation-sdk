# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    # Show timeline values on chart
    show_value: common.VisibilityMode
    # Controls the row height
    row_height: float
    # Merge equal consecutive values
    merge_values: typing.Optional[bool]
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    timezone: typing.Optional[list[common.TimeZone]]
    # Controls value alignment on the timelines
    align_value: typing.Optional[common.TimelineValueAlignment]

    def __init__(self, show_value: typing.Optional[common.VisibilityMode] = None, row_height: float = 0.9, merge_values: typing.Optional[bool] = True, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, timezone: typing.Optional[list[common.TimeZone]] = None, align_value: typing.Optional[common.TimelineValueAlignment] = None):
        self.show_value = show_value if show_value is not None else common.VisibilityMode.AUTO
        self.row_height = row_height
        self.merge_values = merge_values
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.timezone = timezone
        self.align_value = align_value if align_value is not None else common.TimelineValueAlignment.LEFT

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "showValue": self.show_value,
            "rowHeight": self.row_height,
            "legend": self.legend,
            "tooltip": self.tooltip,
        }
        if self.merge_values is not None:
            payload["mergeValues"] = self.merge_values
        if self.timezone is not None:
            payload["timezone"] = self.timezone
        if self.align_value is not None:
            payload["alignValue"] = self.align_value
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "showValue" in data:
            args["show_value"] = data["showValue"]
        if "rowHeight" in data:
            args["row_height"] = data["rowHeight"]
        if "mergeValues" in data:
            args["merge_values"] = data["mergeValues"]
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "timezone" in data:
            args["timezone"] = data["timezone"]
        if "alignValue" in data:
            args["align_value"] = data["alignValue"]        

        return cls(**args)


class FieldConfig:
    line_width: typing.Optional[int]
    hide_from: typing.Optional[common.HideSeriesConfig]
    fill_opacity: typing.Optional[int]

    def __init__(self, line_width: typing.Optional[int] = 0, hide_from: typing.Optional[common.HideSeriesConfig] = None, fill_opacity: typing.Optional[int] = 70):
        self.line_width = line_width
        self.hide_from = hide_from
        self.fill_opacity = fill_opacity

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
        if self.fill_opacity is not None:
            payload["fillOpacity"] = self.fill_opacity
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
        if "hideFrom" in data:
            args["hide_from"] = common.HideSeriesConfig.from_json(data["hideFrom"])
        if "fillOpacity" in data:
            args["fill_opacity"] = data["fillOpacity"]        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="state-timeline",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )

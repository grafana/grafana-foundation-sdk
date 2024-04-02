# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    # Set the height of the rows
    row_height: float
    # Show values on the columns
    show_value: common.VisibilityMode
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    timezone: typing.Optional[list[common.TimeZone]]
    # Controls the column width
    col_width: typing.Optional[float]

    def __init__(self, row_height: float = 0.9, show_value: typing.Optional[common.VisibilityMode] = None, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, timezone: typing.Optional[list[common.TimeZone]] = None, col_width: typing.Optional[float] = 0.9):
        self.row_height = row_height
        self.show_value = show_value if show_value is not None else common.VisibilityMode.AUTO
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.timezone = timezone
        self.col_width = col_width

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "rowHeight": self.row_height,
            "showValue": self.show_value,
            "legend": self.legend,
            "tooltip": self.tooltip,
        }
        if self.timezone is not None:
            payload["timezone"] = self.timezone
        if self.col_width is not None:
            payload["colWidth"] = self.col_width
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rowHeight" in data:
            args["row_height"] = data["rowHeight"]
        if "showValue" in data:
            args["show_value"] = data["showValue"]
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "timezone" in data:
            args["timezone"] = data["timezone"]
        if "colWidth" in data:
            args["col_width"] = data["colWidth"]        

        return cls(**args)


class FieldConfig:
    line_width: typing.Optional[int]
    hide_from: typing.Optional[common.HideSeriesConfig]
    fill_opacity: typing.Optional[int]

    def __init__(self, line_width: typing.Optional[int] = 1, hide_from: typing.Optional[common.HideSeriesConfig] = None, fill_opacity: typing.Optional[int] = 70):
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
        identifier="status-history",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )

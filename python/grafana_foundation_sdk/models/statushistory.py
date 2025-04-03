# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..models import common
import typing
from ..cog import runtime as cogruntime


class Options:
    # Set the height of the rows
    row_height: float
    # Show values on the columns
    show_value: common.VisibilityMode
    # Controls the column width
    col_width: typing.Optional[float]
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    timezone: typing.Optional[list[common.TimeZone]]
    # Enables pagination when > 0
    per_page: typing.Optional[float]

    def __init__(self, row_height: float = 0.9, show_value: typing.Optional[common.VisibilityMode] = None, col_width: typing.Optional[float] = 0.9, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, timezone: typing.Optional[list[common.TimeZone]] = None, per_page: typing.Optional[float] = 20):
        self.row_height = row_height
        self.show_value = show_value if show_value is not None else common.VisibilityMode.AUTO
        self.col_width = col_width
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.timezone = timezone
        self.per_page = per_page

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "rowHeight": self.row_height,
            "showValue": self.show_value,
            "legend": self.legend,
            "tooltip": self.tooltip,
        }
        if self.col_width is not None:
            payload["colWidth"] = self.col_width
        if self.timezone is not None:
            payload["timezone"] = self.timezone
        if self.per_page is not None:
            payload["perPage"] = self.per_page
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rowHeight" in data:
            args["row_height"] = data["rowHeight"]
        if "showValue" in data:
            args["show_value"] = data["showValue"]
        if "colWidth" in data:
            args["col_width"] = data["colWidth"]
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "timezone" in data:
            args["timezone"] = [item for item in data["timezone"]]
        if "perPage" in data:
            args["per_page"] = data["perPage"]        

        return cls(**args)


class FieldConfig:
    line_width: typing.Optional[int]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    axis_centered_zero: typing.Optional[bool]
    hide_from: typing.Optional[common.HideSeriesConfig]
    fill_opacity: typing.Optional[int]
    axis_border_show: typing.Optional[bool]

    def __init__(self, line_width: typing.Optional[int] = 1, axis_placement: typing.Optional[common.AxisPlacement] = None, axis_color_mode: typing.Optional[common.AxisColorMode] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional[common.ScaleDistributionConfig] = None, axis_centered_zero: typing.Optional[bool] = None, hide_from: typing.Optional[common.HideSeriesConfig] = None, fill_opacity: typing.Optional[int] = 70, axis_border_show: typing.Optional[bool] = None):
        self.line_width = line_width
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.axis_centered_zero = axis_centered_zero
        self.hide_from = hide_from
        self.fill_opacity = fill_opacity
        self.axis_border_show = axis_border_show

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.axis_placement is not None:
            payload["axisPlacement"] = self.axis_placement
        if self.axis_color_mode is not None:
            payload["axisColorMode"] = self.axis_color_mode
        if self.axis_label is not None:
            payload["axisLabel"] = self.axis_label
        if self.axis_width is not None:
            payload["axisWidth"] = self.axis_width
        if self.axis_soft_min is not None:
            payload["axisSoftMin"] = self.axis_soft_min
        if self.axis_soft_max is not None:
            payload["axisSoftMax"] = self.axis_soft_max
        if self.axis_grid_show is not None:
            payload["axisGridShow"] = self.axis_grid_show
        if self.scale_distribution is not None:
            payload["scaleDistribution"] = self.scale_distribution
        if self.axis_centered_zero is not None:
            payload["axisCenteredZero"] = self.axis_centered_zero
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
        if self.fill_opacity is not None:
            payload["fillOpacity"] = self.fill_opacity
        if self.axis_border_show is not None:
            payload["axisBorderShow"] = self.axis_border_show
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
        if "axisPlacement" in data:
            args["axis_placement"] = data["axisPlacement"]
        if "axisColorMode" in data:
            args["axis_color_mode"] = data["axisColorMode"]
        if "axisLabel" in data:
            args["axis_label"] = data["axisLabel"]
        if "axisWidth" in data:
            args["axis_width"] = data["axisWidth"]
        if "axisSoftMin" in data:
            args["axis_soft_min"] = data["axisSoftMin"]
        if "axisSoftMax" in data:
            args["axis_soft_max"] = data["axisSoftMax"]
        if "axisGridShow" in data:
            args["axis_grid_show"] = data["axisGridShow"]
        if "scaleDistribution" in data:
            args["scale_distribution"] = common.ScaleDistributionConfig.from_json(data["scaleDistribution"])
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]
        if "hideFrom" in data:
            args["hide_from"] = common.HideSeriesConfig.from_json(data["hideFrom"])
        if "fillOpacity" in data:
            args["fill_opacity"] = data["fillOpacity"]
        if "axisBorderShow" in data:
            args["axis_border_show"] = data["axisBorderShow"]        

        return cls(**args)





def variant_config() -> cogruntime.PanelCfgConfig:
    return cogruntime.PanelCfgConfig(
        identifier="status-history",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )


# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing
from ..models import common
from ..cog import runtime as cogruntime


class SeriesMapping(enum.StrEnum):
    AUTO = "auto"
    MANUAL = "manual"


class ScatterShow(enum.StrEnum):
    POINTS = "points"
    LINES = "lines"
    POINTS_AND_LINES = "points+lines"


class XYDimensionConfig:
    frame: int
    x: typing.Optional[str]
    exclude: typing.Optional[list[str]]

    def __init__(self, frame: int = 0, x: typing.Optional[str] = None, exclude: typing.Optional[list[str]] = None):
        self.frame = frame
        self.x = x
        self.exclude = exclude

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "frame": self.frame,
        }
        if self.x is not None:
            payload["x"] = self.x
        if self.exclude is not None:
            payload["exclude"] = self.exclude
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "frame" in data:
            args["frame"] = data["frame"]
        if "x" in data:
            args["x"] = data["x"]
        if "exclude" in data:
            args["exclude"] = data["exclude"]        

        return cls(**args)


class FieldConfig:
    show: typing.Optional['ScatterShow']
    point_size: typing.Optional[common.ScaleDimensionConfig]
    point_color: typing.Optional[common.ColorDimensionConfig]
    line_color: typing.Optional[common.ColorDimensionConfig]
    line_width: typing.Optional[int]
    line_style: typing.Optional[common.LineStyle]
    label: typing.Optional[common.VisibilityMode]
    hide_from: typing.Optional[common.HideSeriesConfig]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    label_value: typing.Optional[common.TextDimensionConfig]
    axis_centered_zero: typing.Optional[bool]

    def __init__(self, show: typing.Optional['ScatterShow'] = None, point_size: typing.Optional[common.ScaleDimensionConfig] = None, point_color: typing.Optional[common.ColorDimensionConfig] = None, line_color: typing.Optional[common.ColorDimensionConfig] = None, line_width: typing.Optional[int] = None, line_style: typing.Optional[common.LineStyle] = None, label: typing.Optional[common.VisibilityMode] = None, hide_from: typing.Optional[common.HideSeriesConfig] = None, axis_placement: typing.Optional[common.AxisPlacement] = None, axis_color_mode: typing.Optional[common.AxisColorMode] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional[common.ScaleDistributionConfig] = None, label_value: typing.Optional[common.TextDimensionConfig] = None, axis_centered_zero: typing.Optional[bool] = None):
        self.show = show if show is not None else ScatterShow.POINTS
        self.point_size = point_size
        self.point_color = point_color
        self.line_color = line_color
        self.line_width = line_width
        self.line_style = line_style
        self.label = label if label is not None else common.VisibilityMode.AUTO
        self.hide_from = hide_from
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.label_value = label_value
        self.axis_centered_zero = axis_centered_zero

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.show is not None:
            payload["show"] = self.show
        if self.point_size is not None:
            payload["pointSize"] = self.point_size
        if self.point_color is not None:
            payload["pointColor"] = self.point_color
        if self.line_color is not None:
            payload["lineColor"] = self.line_color
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.line_style is not None:
            payload["lineStyle"] = self.line_style
        if self.label is not None:
            payload["label"] = self.label
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
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
        if self.label_value is not None:
            payload["labelValue"] = self.label_value
        if self.axis_centered_zero is not None:
            payload["axisCenteredZero"] = self.axis_centered_zero
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "show" in data:
            args["show"] = data["show"]
        if "pointSize" in data:
            args["point_size"] = common.ScaleDimensionConfig.from_json(data["pointSize"])
        if "pointColor" in data:
            args["point_color"] = common.ColorDimensionConfig.from_json(data["pointColor"])
        if "lineColor" in data:
            args["line_color"] = common.ColorDimensionConfig.from_json(data["lineColor"])
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
        if "lineStyle" in data:
            args["line_style"] = common.LineStyle.from_json(data["lineStyle"])
        if "label" in data:
            args["label"] = data["label"]
        if "hideFrom" in data:
            args["hide_from"] = common.HideSeriesConfig.from_json(data["hideFrom"])
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
        if "labelValue" in data:
            args["label_value"] = common.TextDimensionConfig.from_json(data["labelValue"])
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]        

        return cls(**args)


class ScatterSeriesConfig:
    x: typing.Optional[str]
    y: typing.Optional[str]
    show: typing.Optional['ScatterShow']
    point_size: typing.Optional[common.ScaleDimensionConfig]
    point_color: typing.Optional[common.ColorDimensionConfig]
    line_color: typing.Optional[common.ColorDimensionConfig]
    line_width: typing.Optional[int]
    line_style: typing.Optional[common.LineStyle]
    label: typing.Optional[common.VisibilityMode]
    hide_from: typing.Optional[common.HideSeriesConfig]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    name: typing.Optional[str]
    label_value: typing.Optional[common.TextDimensionConfig]
    axis_centered_zero: typing.Optional[bool]

    def __init__(self, x: typing.Optional[str] = None, y: typing.Optional[str] = None, show: typing.Optional['ScatterShow'] = None, point_size: typing.Optional[common.ScaleDimensionConfig] = None, point_color: typing.Optional[common.ColorDimensionConfig] = None, line_color: typing.Optional[common.ColorDimensionConfig] = None, line_width: typing.Optional[int] = None, line_style: typing.Optional[common.LineStyle] = None, label: typing.Optional[common.VisibilityMode] = None, hide_from: typing.Optional[common.HideSeriesConfig] = None, axis_placement: typing.Optional[common.AxisPlacement] = None, axis_color_mode: typing.Optional[common.AxisColorMode] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional[common.ScaleDistributionConfig] = None, name: typing.Optional[str] = None, label_value: typing.Optional[common.TextDimensionConfig] = None, axis_centered_zero: typing.Optional[bool] = None):
        self.x = x
        self.y = y
        self.show = show if show is not None else ScatterShow.POINTS
        self.point_size = point_size
        self.point_color = point_color
        self.line_color = line_color
        self.line_width = line_width
        self.line_style = line_style
        self.label = label if label is not None else common.VisibilityMode.AUTO
        self.hide_from = hide_from
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.name = name
        self.label_value = label_value
        self.axis_centered_zero = axis_centered_zero

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.x is not None:
            payload["x"] = self.x
        if self.y is not None:
            payload["y"] = self.y
        if self.show is not None:
            payload["show"] = self.show
        if self.point_size is not None:
            payload["pointSize"] = self.point_size
        if self.point_color is not None:
            payload["pointColor"] = self.point_color
        if self.line_color is not None:
            payload["lineColor"] = self.line_color
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.line_style is not None:
            payload["lineStyle"] = self.line_style
        if self.label is not None:
            payload["label"] = self.label
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
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
        if self.name is not None:
            payload["name"] = self.name
        if self.label_value is not None:
            payload["labelValue"] = self.label_value
        if self.axis_centered_zero is not None:
            payload["axisCenteredZero"] = self.axis_centered_zero
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "x" in data:
            args["x"] = data["x"]
        if "y" in data:
            args["y"] = data["y"]
        if "show" in data:
            args["show"] = data["show"]
        if "pointSize" in data:
            args["point_size"] = common.ScaleDimensionConfig.from_json(data["pointSize"])
        if "pointColor" in data:
            args["point_color"] = common.ColorDimensionConfig.from_json(data["pointColor"])
        if "lineColor" in data:
            args["line_color"] = common.ColorDimensionConfig.from_json(data["lineColor"])
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
        if "lineStyle" in data:
            args["line_style"] = common.LineStyle.from_json(data["lineStyle"])
        if "label" in data:
            args["label"] = data["label"]
        if "hideFrom" in data:
            args["hide_from"] = common.HideSeriesConfig.from_json(data["hideFrom"])
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
        if "name" in data:
            args["name"] = data["name"]
        if "labelValue" in data:
            args["label_value"] = common.TextDimensionConfig.from_json(data["labelValue"])
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]        

        return cls(**args)


class Options:
    series_mapping: typing.Optional['SeriesMapping']
    dims: 'XYDimensionConfig'
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    series: list['ScatterSeriesConfig']

    def __init__(self, series_mapping: typing.Optional['SeriesMapping'] = None, dims: typing.Optional['XYDimensionConfig'] = None, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, series: typing.Optional[list['ScatterSeriesConfig']] = None):
        self.series_mapping = series_mapping
        self.dims = dims if dims is not None else XYDimensionConfig()
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.series = series if series is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "dims": self.dims,
            "legend": self.legend,
            "tooltip": self.tooltip,
            "series": self.series,
        }
        if self.series_mapping is not None:
            payload["seriesMapping"] = self.series_mapping
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "seriesMapping" in data:
            args["series_mapping"] = data["seriesMapping"]
        if "dims" in data:
            args["dims"] = XYDimensionConfig.from_json(data["dims"])
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "series" in data:
            args["series"] = data["series"]        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="xychart",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )

# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..models import common
from ..cog import runtime as cogruntime


class Options:
    # Size of each bucket
    bucket_size: typing.Optional[int]
    # Offset buckets by this amount
    bucket_offset: typing.Optional[int]
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    # Combines multiple series into a single histogram
    combine: typing.Optional[bool]

    def __init__(self, bucket_size: typing.Optional[int] = None, bucket_offset: typing.Optional[int] = 0, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, combine: typing.Optional[bool] = None):
        self.bucket_size = bucket_size
        self.bucket_offset = bucket_offset
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.combine = combine

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "legend": self.legend,
            "tooltip": self.tooltip,
        }
        if self.bucket_size is not None:
            payload["bucketSize"] = self.bucket_size
        if self.bucket_offset is not None:
            payload["bucketOffset"] = self.bucket_offset
        if self.combine is not None:
            payload["combine"] = self.combine
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "bucketSize" in data:
            args["bucket_size"] = data["bucketSize"]
        if "bucketOffset" in data:
            args["bucket_offset"] = data["bucketOffset"]
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "combine" in data:
            args["combine"] = data["combine"]        

        return cls(**args)


class FieldConfig:
    # Controls line width of the bars.
    line_width: typing.Optional[int]
    # Controls the fill opacity of the bars.
    fill_opacity: typing.Optional[int]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    hide_from: typing.Optional[common.HideSeriesConfig]
    # Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
    # Gradient appearance is influenced by the Fill opacity setting.
    gradient_mode: typing.Optional[common.GraphGradientMode]
    axis_centered_zero: typing.Optional[bool]

    def __init__(self, line_width: typing.Optional[int] = 1, fill_opacity: typing.Optional[int] = 80, axis_placement: typing.Optional[common.AxisPlacement] = None, axis_color_mode: typing.Optional[common.AxisColorMode] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional[common.ScaleDistributionConfig] = None, hide_from: typing.Optional[common.HideSeriesConfig] = None, gradient_mode: typing.Optional[common.GraphGradientMode] = None, axis_centered_zero: typing.Optional[bool] = None):
        self.line_width = line_width
        self.fill_opacity = fill_opacity
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.hide_from = hide_from
        self.gradient_mode = gradient_mode if gradient_mode is not None else common.GraphGradientMode.NONE
        self.axis_centered_zero = axis_centered_zero

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.fill_opacity is not None:
            payload["fillOpacity"] = self.fill_opacity
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
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
        if self.gradient_mode is not None:
            payload["gradientMode"] = self.gradient_mode
        if self.axis_centered_zero is not None:
            payload["axisCenteredZero"] = self.axis_centered_zero
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
        if "fillOpacity" in data:
            args["fill_opacity"] = data["fillOpacity"]
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
        if "hideFrom" in data:
            args["hide_from"] = common.HideSeriesConfig.from_json(data["hideFrom"])
        if "gradientMode" in data:
            args["gradient_mode"] = data["gradientMode"]
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="histogram",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )

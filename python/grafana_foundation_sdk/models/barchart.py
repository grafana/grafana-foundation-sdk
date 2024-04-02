# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..models import common
from ..cog import runtime as cogruntime


class Options:
    # Manually select which field from the dataset to represent the x field.
    x_field: typing.Optional[str]
    # Use the color value for a sibling field to color each bar value.
    color_by_field: typing.Optional[str]
    # Controls the orientation of the bar chart, either vertical or horizontal.
    orientation: common.VizOrientation
    # Controls the radius of each bar.
    bar_radius: typing.Optional[float]
    # Controls the rotation of the x axis labels.
    x_tick_label_rotation: int
    # Sets the max length that a label can have before it is truncated.
    x_tick_label_max_length: int
    # Controls the spacing between x axis labels.
    # negative values indicate backwards skipping behavior
    x_tick_label_spacing: typing.Optional[int]
    # Controls whether bars are stacked or not, either normally or in percent mode.
    stacking: common.StackingMode
    # This controls whether values are shown on top or to the left of bars.
    show_value: common.VisibilityMode
    # Controls the width of bars. 1 = Max width, 0 = Min width.
    bar_width: float
    # Controls the width of groups. 1 = max with, 0 = min width.
    group_width: float
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    text: typing.Optional[common.VizTextDisplayOptions]
    # Enables mode which highlights the entire bar area and shows tooltip when cursor
    # hovers over highlighted area
    full_highlight: bool

    def __init__(self, x_field: typing.Optional[str] = None, color_by_field: typing.Optional[str] = None, orientation: typing.Optional[common.VizOrientation] = None, bar_radius: typing.Optional[float] = 0, x_tick_label_rotation: int = 0, x_tick_label_max_length: int = 0, x_tick_label_spacing: typing.Optional[int] = 0, stacking: typing.Optional[common.StackingMode] = None, show_value: typing.Optional[common.VisibilityMode] = None, bar_width: float = 0.97, group_width: float = 0.7, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, text: typing.Optional[common.VizTextDisplayOptions] = None, full_highlight: bool = False):
        self.x_field = x_field
        self.color_by_field = color_by_field
        self.orientation = orientation if orientation is not None else common.VizOrientation.AUTO
        self.bar_radius = bar_radius
        self.x_tick_label_rotation = x_tick_label_rotation
        self.x_tick_label_max_length = x_tick_label_max_length
        self.x_tick_label_spacing = x_tick_label_spacing
        self.stacking = stacking if stacking is not None else common.StackingMode.NONE
        self.show_value = show_value if show_value is not None else common.VisibilityMode.AUTO
        self.bar_width = bar_width
        self.group_width = group_width
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.text = text
        self.full_highlight = full_highlight

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "orientation": self.orientation,
            "xTickLabelRotation": self.x_tick_label_rotation,
            "xTickLabelMaxLength": self.x_tick_label_max_length,
            "stacking": self.stacking,
            "showValue": self.show_value,
            "barWidth": self.bar_width,
            "groupWidth": self.group_width,
            "legend": self.legend,
            "tooltip": self.tooltip,
            "fullHighlight": self.full_highlight,
        }
        if self.x_field is not None:
            payload["xField"] = self.x_field
        if self.color_by_field is not None:
            payload["colorByField"] = self.color_by_field
        if self.bar_radius is not None:
            payload["barRadius"] = self.bar_radius
        if self.x_tick_label_spacing is not None:
            payload["xTickLabelSpacing"] = self.x_tick_label_spacing
        if self.text is not None:
            payload["text"] = self.text
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "xField" in data:
            args["x_field"] = data["xField"]
        if "colorByField" in data:
            args["color_by_field"] = data["colorByField"]
        if "orientation" in data:
            args["orientation"] = data["orientation"]
        if "barRadius" in data:
            args["bar_radius"] = data["barRadius"]
        if "xTickLabelRotation" in data:
            args["x_tick_label_rotation"] = data["xTickLabelRotation"]
        if "xTickLabelMaxLength" in data:
            args["x_tick_label_max_length"] = data["xTickLabelMaxLength"]
        if "xTickLabelSpacing" in data:
            args["x_tick_label_spacing"] = data["xTickLabelSpacing"]
        if "stacking" in data:
            args["stacking"] = data["stacking"]
        if "showValue" in data:
            args["show_value"] = data["showValue"]
        if "barWidth" in data:
            args["bar_width"] = data["barWidth"]
        if "groupWidth" in data:
            args["group_width"] = data["groupWidth"]
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "text" in data:
            args["text"] = common.VizTextDisplayOptions.from_json(data["text"])
        if "fullHighlight" in data:
            args["full_highlight"] = data["fullHighlight"]        

        return cls(**args)


class FieldConfig:
    # Controls line width of the bars.
    line_width: typing.Optional[int]
    # Controls the fill opacity of the bars.
    fill_opacity: typing.Optional[int]
    # Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.
    # Gradient appearance is influenced by the Fill opacity setting.
    gradient_mode: typing.Optional[common.GraphGradientMode]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    hide_from: typing.Optional[common.HideSeriesConfig]
    # Threshold rendering
    thresholds_style: typing.Optional[common.GraphThresholdsStyleConfig]
    axis_centered_zero: typing.Optional[bool]

    def __init__(self, line_width: typing.Optional[int] = 1, fill_opacity: typing.Optional[int] = 80, gradient_mode: typing.Optional[common.GraphGradientMode] = None, axis_placement: typing.Optional[common.AxisPlacement] = None, axis_color_mode: typing.Optional[common.AxisColorMode] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional[common.ScaleDistributionConfig] = None, hide_from: typing.Optional[common.HideSeriesConfig] = None, thresholds_style: typing.Optional[common.GraphThresholdsStyleConfig] = None, axis_centered_zero: typing.Optional[bool] = None):
        self.line_width = line_width
        self.fill_opacity = fill_opacity
        self.gradient_mode = gradient_mode if gradient_mode is not None else common.GraphGradientMode.NONE
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.hide_from = hide_from
        self.thresholds_style = thresholds_style
        self.axis_centered_zero = axis_centered_zero

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
        if self.fill_opacity is not None:
            payload["fillOpacity"] = self.fill_opacity
        if self.gradient_mode is not None:
            payload["gradientMode"] = self.gradient_mode
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
        if self.thresholds_style is not None:
            payload["thresholdsStyle"] = self.thresholds_style
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
        if "gradientMode" in data:
            args["gradient_mode"] = data["gradientMode"]
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
        if "thresholdsStyle" in data:
            args["thresholds_style"] = common.GraphThresholdsStyleConfig.from_json(data["thresholdsStyle"])
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="barchart",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )

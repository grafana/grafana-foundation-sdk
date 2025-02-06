# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing
from ..models import common
from ..cog import runtime as cogruntime


class PointShape(enum.StrEnum):
    CIRCLE = "circle"
    SQUARE = "square"


class SeriesMapping(enum.StrEnum):
    AUTO = "auto"
    MANUAL = "manual"


class XYShowMode(enum.StrEnum):
    POINTS = "points"
    LINES = "lines"
    POINTS_AND_LINES = "points+lines"


class MatcherConfig:
    """
    NOTE: (copied from dashboard_kind.cue, since not exported)
    Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
    It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
    """

    # The matcher id. This is used to find the matcher implementation from registry.
    id_val: str
    # The matcher options. This is specific to the matcher implementation.
    options: typing.Optional[object]

    def __init__(self, id_val: str = "", options: typing.Optional[object] = None):
        self.id_val = id_val
        self.options = options

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
        }
        if self.options is not None:
            payload["options"] = self.options
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "options" in data:
            args["options"] = data["options"]        

        return cls(**args)


class FieldConfig:
    show: typing.Optional['XYShowMode']
    point_size: typing.Optional['XychartFieldConfigPointSize']
    point_shape: typing.Optional['PointShape']
    point_stroke_width: typing.Optional[int]
    fill_opacity: typing.Optional[int]
    line_width: typing.Optional[int]
    hide_from: typing.Optional[common.HideSeriesConfig]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    axis_centered_zero: typing.Optional[bool]
    line_style: typing.Optional[common.LineStyle]
    axis_border_show: typing.Optional[bool]

    def __init__(self, show: typing.Optional['XYShowMode'] = None, point_size: typing.Optional['XychartFieldConfigPointSize'] = None, point_shape: typing.Optional['PointShape'] = None, point_stroke_width: typing.Optional[int] = None, fill_opacity: typing.Optional[int] = 50, line_width: typing.Optional[int] = None, hide_from: typing.Optional[common.HideSeriesConfig] = None, axis_placement: typing.Optional[common.AxisPlacement] = None, axis_color_mode: typing.Optional[common.AxisColorMode] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional[common.ScaleDistributionConfig] = None, axis_centered_zero: typing.Optional[bool] = None, line_style: typing.Optional[common.LineStyle] = None, axis_border_show: typing.Optional[bool] = None):
        self.show = show if show is not None else XYShowMode.POINTS
        self.point_size = point_size
        self.point_shape = point_shape
        self.point_stroke_width = point_stroke_width
        self.fill_opacity = fill_opacity
        self.line_width = line_width
        self.hide_from = hide_from
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.axis_centered_zero = axis_centered_zero
        self.line_style = line_style
        self.axis_border_show = axis_border_show

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.show is not None:
            payload["show"] = self.show
        if self.point_size is not None:
            payload["pointSize"] = self.point_size
        if self.point_shape is not None:
            payload["pointShape"] = self.point_shape
        if self.point_stroke_width is not None:
            payload["pointStrokeWidth"] = self.point_stroke_width
        if self.fill_opacity is not None:
            payload["fillOpacity"] = self.fill_opacity
        if self.line_width is not None:
            payload["lineWidth"] = self.line_width
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
        if self.axis_centered_zero is not None:
            payload["axisCenteredZero"] = self.axis_centered_zero
        if self.line_style is not None:
            payload["lineStyle"] = self.line_style
        if self.axis_border_show is not None:
            payload["axisBorderShow"] = self.axis_border_show
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "show" in data:
            args["show"] = data["show"]
        if "pointSize" in data:
            args["point_size"] = XychartFieldConfigPointSize.from_json(data["pointSize"])
        if "pointShape" in data:
            args["point_shape"] = data["pointShape"]
        if "pointStrokeWidth" in data:
            args["point_stroke_width"] = data["pointStrokeWidth"]
        if "fillOpacity" in data:
            args["fill_opacity"] = data["fillOpacity"]
        if "lineWidth" in data:
            args["line_width"] = data["lineWidth"]
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
        if "axisCenteredZero" in data:
            args["axis_centered_zero"] = data["axisCenteredZero"]
        if "lineStyle" in data:
            args["line_style"] = common.LineStyle.from_json(data["lineStyle"])
        if "axisBorderShow" in data:
            args["axis_border_show"] = data["axisBorderShow"]        

        return cls(**args)


class XYSeriesConfig:
    name: typing.Optional['XychartXYSeriesConfigName']
    frame: typing.Optional['XychartXYSeriesConfigFrame']
    x: typing.Optional['XychartXYSeriesConfigX']
    y: typing.Optional['XychartXYSeriesConfigY']
    color: typing.Optional['XychartXYSeriesConfigColor']
    size: typing.Optional['XychartXYSeriesConfigSize']

    def __init__(self, name: typing.Optional['XychartXYSeriesConfigName'] = None, frame: typing.Optional['XychartXYSeriesConfigFrame'] = None, x: typing.Optional['XychartXYSeriesConfigX'] = None, y: typing.Optional['XychartXYSeriesConfigY'] = None, color: typing.Optional['XychartXYSeriesConfigColor'] = None, size: typing.Optional['XychartXYSeriesConfigSize'] = None):
        self.name = name
        self.frame = frame
        self.x = x
        self.y = y
        self.color = color
        self.size = size

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.name is not None:
            payload["name"] = self.name
        if self.frame is not None:
            payload["frame"] = self.frame
        if self.x is not None:
            payload["x"] = self.x
        if self.y is not None:
            payload["y"] = self.y
        if self.color is not None:
            payload["color"] = self.color
        if self.size is not None:
            payload["size"] = self.size
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = XychartXYSeriesConfigName.from_json(data["name"])
        if "frame" in data:
            args["frame"] = XychartXYSeriesConfigFrame.from_json(data["frame"])
        if "x" in data:
            args["x"] = XychartXYSeriesConfigX.from_json(data["x"])
        if "y" in data:
            args["y"] = XychartXYSeriesConfigY.from_json(data["y"])
        if "color" in data:
            args["color"] = XychartXYSeriesConfigColor.from_json(data["color"])
        if "size" in data:
            args["size"] = XychartXYSeriesConfigSize.from_json(data["size"])        

        return cls(**args)


class Options:
    mapping: 'SeriesMapping'
    legend: common.VizLegendOptions
    tooltip: common.VizTooltipOptions
    series: list['XYSeriesConfig']

    def __init__(self, mapping: typing.Optional['SeriesMapping'] = None, legend: typing.Optional[common.VizLegendOptions] = None, tooltip: typing.Optional[common.VizTooltipOptions] = None, series: typing.Optional[list['XYSeriesConfig']] = None):
        self.mapping = mapping if mapping is not None else SeriesMapping.AUTO
        self.legend = legend if legend is not None else common.VizLegendOptions()
        self.tooltip = tooltip if tooltip is not None else common.VizTooltipOptions()
        self.series = series if series is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mapping": self.mapping,
            "legend": self.legend,
            "tooltip": self.tooltip,
            "series": self.series,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mapping" in data:
            args["mapping"] = data["mapping"]
        if "legend" in data:
            args["legend"] = common.VizLegendOptions.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = common.VizTooltipOptions.from_json(data["tooltip"])
        if "series" in data:
            args["series"] = [XYSeriesConfig.from_json(item) for item in data["series"]]        

        return cls(**args)


class XychartFieldConfigPointSize:
    fixed: typing.Optional[int]
    min_val: typing.Optional[int]
    max_val: typing.Optional[int]

    def __init__(self, fixed: typing.Optional[int] = None, min_val: typing.Optional[int] = None, max_val: typing.Optional[int] = None):
        self.fixed = fixed
        self.min_val = min_val
        self.max_val = max_val

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.fixed is not None:
            payload["fixed"] = self.fixed
        if self.min_val is not None:
            payload["min"] = self.min_val
        if self.max_val is not None:
            payload["max"] = self.max_val
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "fixed" in data:
            args["fixed"] = data["fixed"]
        if "min" in data:
            args["min_val"] = data["min"]
        if "max" in data:
            args["max_val"] = data["max"]        

        return cls(**args)


class XychartXYSeriesConfigName:
    fixed: typing.Optional[str]

    def __init__(self, fixed: typing.Optional[str] = None):
        self.fixed = fixed

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.fixed is not None:
            payload["fixed"] = self.fixed
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "fixed" in data:
            args["fixed"] = data["fixed"]        

        return cls(**args)


class XychartXYSeriesConfigFrame:
    matcher: 'MatcherConfig'

    def __init__(self, matcher: typing.Optional['MatcherConfig'] = None):
        self.matcher = matcher if matcher is not None else MatcherConfig()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "matcher": self.matcher,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "matcher" in data:
            args["matcher"] = MatcherConfig.from_json(data["matcher"])        

        return cls(**args)


class XychartXYSeriesConfigX:
    matcher: 'MatcherConfig'

    def __init__(self, matcher: typing.Optional['MatcherConfig'] = None):
        self.matcher = matcher if matcher is not None else MatcherConfig()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "matcher": self.matcher,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "matcher" in data:
            args["matcher"] = MatcherConfig.from_json(data["matcher"])        

        return cls(**args)


class XychartXYSeriesConfigY:
    matcher: 'MatcherConfig'

    def __init__(self, matcher: typing.Optional['MatcherConfig'] = None):
        self.matcher = matcher if matcher is not None else MatcherConfig()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "matcher": self.matcher,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "matcher" in data:
            args["matcher"] = MatcherConfig.from_json(data["matcher"])        

        return cls(**args)


class XychartXYSeriesConfigColor:
    matcher: 'MatcherConfig'

    def __init__(self, matcher: typing.Optional['MatcherConfig'] = None):
        self.matcher = matcher if matcher is not None else MatcherConfig()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "matcher": self.matcher,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "matcher" in data:
            args["matcher"] = MatcherConfig.from_json(data["matcher"])        

        return cls(**args)


class XychartXYSeriesConfigSize:
    matcher: 'MatcherConfig'

    def __init__(self, matcher: typing.Optional['MatcherConfig'] = None):
        self.matcher = matcher if matcher is not None else MatcherConfig()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "matcher": self.matcher,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "matcher" in data:
            args["matcher"] = MatcherConfig.from_json(data["matcher"])        

        return cls(**args)





def variant_config() -> cogruntime.PanelCfgConfig:
    return cogruntime.PanelCfgConfig(
        identifier="xychart",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )


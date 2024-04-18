# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing
from ..models import common
from ..cog import runtime as cogruntime


class HeatmapColorMode(enum.StrEnum):
    """
    Controls the color mode of the heatmap
    """

    OPACITY = "opacity"
    SCHEME = "scheme"


class HeatmapColorScale(enum.StrEnum):
    """
    Controls the color scale of the heatmap
    """

    LINEAR = "linear"
    EXPONENTIAL = "exponential"


class HeatmapColorOptions:
    """
    Controls various color options
    """

    # Sets the color mode
    mode: typing.Optional['HeatmapColorMode']
    # Controls the color scheme used
    scheme: str
    # Controls the color fill when in opacity mode
    fill: str
    # Controls the color scale
    scale: typing.Optional['HeatmapColorScale']
    # Controls the exponent when scale is set to exponential
    exponent: float
    # Controls the number of color steps
    steps: int
    # Reverses the color scheme
    reverse: bool
    # Sets the minimum value for the color scale
    min_val: typing.Optional[float]
    # Sets the maximum value for the color scale
    max_val: typing.Optional[float]

    def __init__(self, mode: typing.Optional['HeatmapColorMode'] = None, scheme: str = "", fill: str = "", scale: typing.Optional['HeatmapColorScale'] = None, exponent: float = 0, steps: int = 0, reverse: bool = False, min_val: typing.Optional[float] = None, max_val: typing.Optional[float] = None):
        self.mode = mode
        self.scheme = scheme
        self.fill = fill
        self.scale = scale
        self.exponent = exponent
        self.steps = steps
        self.reverse = reverse
        self.min_val = min_val
        self.max_val = max_val

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "scheme": self.scheme,
            "fill": self.fill,
            "exponent": self.exponent,
            "steps": self.steps,
            "reverse": self.reverse,
        }
        if self.mode is not None:
            payload["mode"] = self.mode
        if self.scale is not None:
            payload["scale"] = self.scale
        if self.min_val is not None:
            payload["min"] = self.min_val
        if self.max_val is not None:
            payload["max"] = self.max_val
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "scheme" in data:
            args["scheme"] = data["scheme"]
        if "fill" in data:
            args["fill"] = data["fill"]
        if "scale" in data:
            args["scale"] = data["scale"]
        if "exponent" in data:
            args["exponent"] = data["exponent"]
        if "steps" in data:
            args["steps"] = data["steps"]
        if "reverse" in data:
            args["reverse"] = data["reverse"]
        if "min" in data:
            args["min_val"] = data["min"]
        if "max" in data:
            args["max_val"] = data["max"]        

        return cls(**args)


class YAxisConfig:
    """
    Configuration options for the yAxis
    """

    # Sets the yAxis unit
    unit: typing.Optional[str]
    # Reverses the yAxis
    reverse: typing.Optional[bool]
    # Controls the number of decimals for yAxis values
    decimals: typing.Optional[float]
    # Sets the minimum value for the yAxis
    min_val: typing.Optional[float]
    axis_placement: typing.Optional[common.AxisPlacement]
    axis_color_mode: typing.Optional[common.AxisColorMode]
    axis_label: typing.Optional[str]
    axis_width: typing.Optional[float]
    axis_soft_min: typing.Optional[float]
    axis_soft_max: typing.Optional[float]
    axis_grid_show: typing.Optional[bool]
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    axis_centered_zero: typing.Optional[bool]
    # Sets the maximum value for the yAxis
    max_val: typing.Optional[float]
    axis_border_show: typing.Optional[bool]

    def __init__(self, unit: typing.Optional[str] = None, reverse: typing.Optional[bool] = None, decimals: typing.Optional[float] = None, min_val: typing.Optional[float] = None, axis_placement: typing.Optional[common.AxisPlacement] = None, axis_color_mode: typing.Optional[common.AxisColorMode] = None, axis_label: typing.Optional[str] = None, axis_width: typing.Optional[float] = None, axis_soft_min: typing.Optional[float] = None, axis_soft_max: typing.Optional[float] = None, axis_grid_show: typing.Optional[bool] = None, scale_distribution: typing.Optional[common.ScaleDistributionConfig] = None, axis_centered_zero: typing.Optional[bool] = None, max_val: typing.Optional[float] = None, axis_border_show: typing.Optional[bool] = None):
        self.unit = unit
        self.reverse = reverse
        self.decimals = decimals
        self.min_val = min_val
        self.axis_placement = axis_placement
        self.axis_color_mode = axis_color_mode
        self.axis_label = axis_label
        self.axis_width = axis_width
        self.axis_soft_min = axis_soft_min
        self.axis_soft_max = axis_soft_max
        self.axis_grid_show = axis_grid_show
        self.scale_distribution = scale_distribution
        self.axis_centered_zero = axis_centered_zero
        self.max_val = max_val
        self.axis_border_show = axis_border_show

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.unit is not None:
            payload["unit"] = self.unit
        if self.reverse is not None:
            payload["reverse"] = self.reverse
        if self.decimals is not None:
            payload["decimals"] = self.decimals
        if self.min_val is not None:
            payload["min"] = self.min_val
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
        if self.max_val is not None:
            payload["max"] = self.max_val
        if self.axis_border_show is not None:
            payload["axisBorderShow"] = self.axis_border_show
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "unit" in data:
            args["unit"] = data["unit"]
        if "reverse" in data:
            args["reverse"] = data["reverse"]
        if "decimals" in data:
            args["decimals"] = data["decimals"]
        if "min" in data:
            args["min_val"] = data["min"]
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
        if "max" in data:
            args["max_val"] = data["max"]
        if "axisBorderShow" in data:
            args["axis_border_show"] = data["axisBorderShow"]        

        return cls(**args)


class CellValues:
    """
    Controls cell value options
    """

    # Controls the cell value unit
    unit: typing.Optional[str]
    # Controls the number of decimals for cell values
    decimals: typing.Optional[float]

    def __init__(self, unit: typing.Optional[str] = None, decimals: typing.Optional[float] = None):
        self.unit = unit
        self.decimals = decimals

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.unit is not None:
            payload["unit"] = self.unit
        if self.decimals is not None:
            payload["decimals"] = self.decimals
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "unit" in data:
            args["unit"] = data["unit"]
        if "decimals" in data:
            args["decimals"] = data["decimals"]        

        return cls(**args)


class FilterValueRange:
    """
    Controls the value filter range
    """

    # Sets the filter range to values less than or equal to the given value
    le: typing.Optional[float]
    # Sets the filter range to values greater than or equal to the given value
    ge: typing.Optional[float]

    def __init__(self, le: typing.Optional[float] = None, ge: typing.Optional[float] = None):
        self.le = le
        self.ge = ge

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.le is not None:
            payload["le"] = self.le
        if self.ge is not None:
            payload["ge"] = self.ge
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "le" in data:
            args["le"] = data["le"]
        if "ge" in data:
            args["ge"] = data["ge"]        

        return cls(**args)


class HeatmapTooltip:
    """
    Controls tooltip options
    """

    # Controls how the tooltip is shown
    mode: common.TooltipDisplayMode
    max_height: typing.Optional[float]
    max_width: typing.Optional[float]
    # Controls if the tooltip shows a histogram of the y-axis values
    y_histogram: typing.Optional[bool]
    # Controls if the tooltip shows a color scale in header
    show_color_scale: typing.Optional[bool]

    def __init__(self, mode: typing.Optional[common.TooltipDisplayMode] = None, max_height: typing.Optional[float] = None, max_width: typing.Optional[float] = None, y_histogram: typing.Optional[bool] = None, show_color_scale: typing.Optional[bool] = None):
        self.mode = mode if mode is not None else common.TooltipDisplayMode.SINGLE
        self.max_height = max_height
        self.max_width = max_width
        self.y_histogram = y_histogram
        self.show_color_scale = show_color_scale

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        if self.max_height is not None:
            payload["maxHeight"] = self.max_height
        if self.max_width is not None:
            payload["maxWidth"] = self.max_width
        if self.y_histogram is not None:
            payload["yHistogram"] = self.y_histogram
        if self.show_color_scale is not None:
            payload["showColorScale"] = self.show_color_scale
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "maxHeight" in data:
            args["max_height"] = data["maxHeight"]
        if "maxWidth" in data:
            args["max_width"] = data["maxWidth"]
        if "yHistogram" in data:
            args["y_histogram"] = data["yHistogram"]
        if "showColorScale" in data:
            args["show_color_scale"] = data["showColorScale"]        

        return cls(**args)


class HeatmapLegend:
    """
    Controls legend options
    """

    # Controls if the legend is shown
    show: bool

    def __init__(self, show: bool = False):
        self.show = show

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "show": self.show,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "show" in data:
            args["show"] = data["show"]        

        return cls(**args)


class ExemplarConfig:
    """
    Controls exemplar options
    """

    # Sets the color of the exemplar markers
    color: str

    def __init__(self, color: str = ""):
        self.color = color

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "color": self.color,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "color" in data:
            args["color"] = data["color"]        

        return cls(**args)


class RowsHeatmapOptions:
    """
    Controls frame rows options
    """

    # Sets the name of the cell when not calculating from data
    value: typing.Optional[str]
    # Controls tick alignment when not calculating from data
    layout: typing.Optional[common.HeatmapCellLayout]

    def __init__(self, value: typing.Optional[str] = None, layout: typing.Optional[common.HeatmapCellLayout] = None):
        self.value = value
        self.layout = layout

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.value is not None:
            payload["value"] = self.value
        if self.layout is not None:
            payload["layout"] = self.layout
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]
        if "layout" in data:
            args["layout"] = data["layout"]        

        return cls(**args)


class Options:
    # Controls if the heatmap should be calculated from data
    calculate: typing.Optional[bool]
    # Calculation options for the heatmap
    calculation: typing.Optional[common.HeatmapCalculationOptions]
    # Controls the color options
    color: 'HeatmapColorOptions'
    # Filters values between a given range
    filter_values: typing.Optional['FilterValueRange']
    # Controls tick alignment and value name when not calculating from data
    rows_frame: typing.Optional['RowsHeatmapOptions']
    # | *{
    # 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    # }
    # Controls the display of the value in the cell
    show_value: common.VisibilityMode
    # Controls gap between cells
    cell_gap: typing.Optional[int]
    # Controls cell radius
    cell_radius: typing.Optional[float]
    # Controls cell value unit
    cell_values: typing.Optional['CellValues']
    # Controls yAxis placement
    y_axis: 'YAxisConfig'
    # | *{
    # 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
    # }
    # Controls legend options
    legend: 'HeatmapLegend'
    # Controls tooltip options
    tooltip: 'HeatmapTooltip'
    # Controls exemplar options
    exemplars: 'ExemplarConfig'

    def __init__(self, calculate: typing.Optional[bool] = False, calculation: typing.Optional[common.HeatmapCalculationOptions] = None, color: typing.Optional['HeatmapColorOptions'] = None, filter_values: typing.Optional['FilterValueRange'] = None, rows_frame: typing.Optional['RowsHeatmapOptions'] = None, show_value: typing.Optional[common.VisibilityMode] = None, cell_gap: typing.Optional[int] = 1, cell_radius: typing.Optional[float] = None, cell_values: typing.Optional['CellValues'] = None, y_axis: typing.Optional['YAxisConfig'] = None, legend: typing.Optional['HeatmapLegend'] = None, tooltip: typing.Optional['HeatmapTooltip'] = None, exemplars: typing.Optional['ExemplarConfig'] = None):
        self.calculate = calculate
        self.calculation = calculation
        self.color = color if color is not None else HeatmapColorOptions(exponent=0.5, fill="dark-orange", reverse=False, scheme="Oranges", steps=64)
        self.filter_values = filter_values if filter_values is not None else FilterValueRange(le=1e-09)
        self.rows_frame = rows_frame
        self.show_value = show_value if show_value is not None else common.VisibilityMode.AUTO
        self.cell_gap = cell_gap
        self.cell_radius = cell_radius
        self.cell_values = cell_values
        self.y_axis = y_axis if y_axis is not None else YAxisConfig()
        self.legend = legend if legend is not None else HeatmapLegend(show=True)
        self.tooltip = tooltip if tooltip is not None else HeatmapTooltip()
        self.exemplars = exemplars if exemplars is not None else ExemplarConfig(color="rgba(255,0,255,0.7)")

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "color": self.color,
            "showValue": self.show_value,
            "yAxis": self.y_axis,
            "legend": self.legend,
            "tooltip": self.tooltip,
            "exemplars": self.exemplars,
        }
        if self.calculate is not None:
            payload["calculate"] = self.calculate
        if self.calculation is not None:
            payload["calculation"] = self.calculation
        if self.filter_values is not None:
            payload["filterValues"] = self.filter_values
        if self.rows_frame is not None:
            payload["rowsFrame"] = self.rows_frame
        if self.cell_gap is not None:
            payload["cellGap"] = self.cell_gap
        if self.cell_radius is not None:
            payload["cellRadius"] = self.cell_radius
        if self.cell_values is not None:
            payload["cellValues"] = self.cell_values
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "calculate" in data:
            args["calculate"] = data["calculate"]
        if "calculation" in data:
            args["calculation"] = common.HeatmapCalculationOptions.from_json(data["calculation"])
        if "color" in data:
            args["color"] = HeatmapColorOptions.from_json(data["color"])
        if "filterValues" in data:
            args["filter_values"] = FilterValueRange.from_json(data["filterValues"])
        if "rowsFrame" in data:
            args["rows_frame"] = RowsHeatmapOptions.from_json(data["rowsFrame"])
        if "showValue" in data:
            args["show_value"] = data["showValue"]
        if "cellGap" in data:
            args["cell_gap"] = data["cellGap"]
        if "cellRadius" in data:
            args["cell_radius"] = data["cellRadius"]
        if "cellValues" in data:
            args["cell_values"] = CellValues.from_json(data["cellValues"])
        if "yAxis" in data:
            args["y_axis"] = YAxisConfig.from_json(data["yAxis"])
        if "legend" in data:
            args["legend"] = HeatmapLegend.from_json(data["legend"])
        if "tooltip" in data:
            args["tooltip"] = HeatmapTooltip.from_json(data["tooltip"])
        if "exemplars" in data:
            args["exemplars"] = ExemplarConfig.from_json(data["exemplars"])        

        return cls(**args)


class FieldConfig:
    scale_distribution: typing.Optional[common.ScaleDistributionConfig]
    hide_from: typing.Optional[common.HideSeriesConfig]

    def __init__(self, scale_distribution: typing.Optional[common.ScaleDistributionConfig] = None, hide_from: typing.Optional[common.HideSeriesConfig] = None):
        self.scale_distribution = scale_distribution
        self.hide_from = hide_from

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.scale_distribution is not None:
            payload["scaleDistribution"] = self.scale_distribution
        if self.hide_from is not None:
            payload["hideFrom"] = self.hide_from
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "scaleDistribution" in data:
            args["scale_distribution"] = common.ScaleDistributionConfig.from_json(data["scaleDistribution"])
        if "hideFrom" in data:
            args["hide_from"] = common.HideSeriesConfig.from_json(data["hideFrom"])        

        return cls(**args)





def variant_config():
    return cogruntime.PanelCfgConfig(
        identifier="heatmap",
        options_from_json_hook=Options.from_json,
        field_config_from_json_hook=FieldConfig.from_json,
    )

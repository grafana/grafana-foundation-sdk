# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import common


class DataSourceJsonData(cogbuilder.Builder[common.DataSourceJsonData]):    
    """
    TODO docs
    """
    
    __internal: common.DataSourceJsonData

    def __init__(self):
        self.__internal = common.DataSourceJsonData()

    def build(self) -> common.DataSourceJsonData:
        return self.__internal    
    
    def auth_type(self, auth_type: str) -> typing.Self:        
        self.__internal.auth_type = auth_type
    
        return self
    
    def default_region(self, default_region: str) -> typing.Self:        
        self.__internal.default_region = default_region
    
        return self
    
    def profile(self, profile: str) -> typing.Self:        
        self.__internal.profile = profile
    
        return self
    
    def manage_alerts(self, manage_alerts: bool) -> typing.Self:        
        self.__internal.manage_alerts = manage_alerts
    
        return self
    
    def alertmanager_uid(self, alertmanager_uid: str) -> typing.Self:        
        self.__internal.alertmanager_uid = alertmanager_uid
    
        return self
    

class DataQuery(cogbuilder.Builder[common.DataQuery]):    
    """
    These are the common properties available to all queries in all datasources.
    Specific implementations will *extend* this interface, adding the required
    properties for the given context.
    """
    
    __internal: common.DataQuery

    def __init__(self):
        self.__internal = common.DataQuery()

    def build(self) -> common.DataQuery:
        return self.__internal    
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        self.__internal.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        Note this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self.__internal.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        self.__internal.query_type = query_type
    
        return self
    
    def datasource(self, datasource: object) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        self.__internal.datasource = datasource
    
        return self
    

class BaseDimensionConfig(cogbuilder.Builder[common.BaseDimensionConfig]):    
    __internal: common.BaseDimensionConfig

    def __init__(self):
        self.__internal = common.BaseDimensionConfig()

    def build(self) -> common.BaseDimensionConfig:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self.__internal.field = field
    
        return self
    

class ScaleDimensionConfig(cogbuilder.Builder[common.ScaleDimensionConfig]):    
    __internal: common.ScaleDimensionConfig

    def __init__(self):
        self.__internal = common.ScaleDimensionConfig()

    def build(self) -> common.ScaleDimensionConfig:
        return self.__internal    
    
    def min_val(self, min_val: float) -> typing.Self:        
        self.__internal.min_val = min_val
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:        
        self.__internal.max_val = max_val
    
        return self
    
    def fixed(self, fixed: float) -> typing.Self:        
        self.__internal.fixed = fixed
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self.__internal.field = field
    
        return self
    
    def mode(self, mode: common.ScaleDimensionMode) -> typing.Self:    
        """
        | *"linear"
        """
            
        self.__internal.mode = mode
    
        return self
    

class ColorDimensionConfig(cogbuilder.Builder[common.ColorDimensionConfig]):    
    __internal: common.ColorDimensionConfig

    def __init__(self):
        self.__internal = common.ColorDimensionConfig()

    def build(self) -> common.ColorDimensionConfig:
        return self.__internal    
    
    def fixed(self, fixed: str) -> typing.Self:    
        """
        color value
        """
            
        self.__internal.fixed = fixed
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self.__internal.field = field
    
        return self
    

class ScalarDimensionConfig(cogbuilder.Builder[common.ScalarDimensionConfig]):    
    __internal: common.ScalarDimensionConfig

    def __init__(self):
        self.__internal = common.ScalarDimensionConfig()

    def build(self) -> common.ScalarDimensionConfig:
        return self.__internal    
    
    def min_val(self, min_val: float) -> typing.Self:        
        self.__internal.min_val = min_val
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:        
        self.__internal.max_val = max_val
    
        return self
    
    def fixed(self, fixed: float) -> typing.Self:        
        self.__internal.fixed = fixed
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self.__internal.field = field
    
        return self
    
    def mode(self, mode: common.ScalarDimensionMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    

class TextDimensionConfig(cogbuilder.Builder[common.TextDimensionConfig]):    
    __internal: common.TextDimensionConfig

    def __init__(self):
        self.__internal = common.TextDimensionConfig()

    def build(self) -> common.TextDimensionConfig:
        return self.__internal    
    
    def mode(self, mode: common.TextDimensionMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self.__internal.field = field
    
        return self
    
    def fixed(self, fixed: str) -> typing.Self:        
        self.__internal.fixed = fixed
    
        return self
    

class MapLayerOptions(cogbuilder.Builder[common.MapLayerOptions]):    
    __internal: common.MapLayerOptions

    def __init__(self):
        self.__internal = common.MapLayerOptions()

    def build(self) -> common.MapLayerOptions:
        return self.__internal    
    
    def type_val(self, type_val: str) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        """
        configured unique display name
        """
            
        self.__internal.name = name
    
        return self
    
    def config(self, config: object) -> typing.Self:    
        """
        Custom options depending on the type
        """
            
        self.__internal.config = config
    
        return self
    
    def location(self, location: cogbuilder.Builder[common.FrameGeometrySource]) -> typing.Self:    
        """
        Common method to define geometry fields
        """
            
        location_resource = location.build()
        self.__internal.location = location_resource
    
        return self
    
    def filter_data(self, filter_data: object) -> typing.Self:    
        """
        Defines a frame MatcherConfig that may filter data for the given layer
        """
            
        self.__internal.filter_data = filter_data
    
        return self
    
    def opacity(self, opacity: int) -> typing.Self:    
        """
        Common properties:
        https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
        Layer opacity (0-1)
        """
            
        self.__internal.opacity = opacity
    
        return self
    
    def tooltip(self, tooltip: bool) -> typing.Self:    
        """
        Check tooltip (defaults to true)
        """
            
        self.__internal.tooltip = tooltip
    
        return self
    

class HeatmapCalculationBucketConfig(cogbuilder.Builder[common.HeatmapCalculationBucketConfig]):    
    __internal: common.HeatmapCalculationBucketConfig

    def __init__(self):
        self.__internal = common.HeatmapCalculationBucketConfig()

    def build(self) -> common.HeatmapCalculationBucketConfig:
        return self.__internal    
    
    def mode(self, mode: common.HeatmapCalculationMode) -> typing.Self:    
        """
        Sets the bucket calculation mode
        """
            
        self.__internal.mode = mode
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        """
        The number of buckets to use for the axis in the heatmap
        """
            
        self.__internal.value = value
    
        return self
    
    def scale(self, scale: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:    
        """
        Controls the scale of the buckets
        """
            
        scale_resource = scale.build()
        self.__internal.scale = scale_resource
    
        return self
    

class LineStyle(cogbuilder.Builder[common.LineStyle]):    
    """
    TODO docs
    """
    
    __internal: common.LineStyle

    def __init__(self):
        self.__internal = common.LineStyle()

    def build(self) -> common.LineStyle:
        return self.__internal    
    
    def fill(self, fill: typing.Literal["solid", "dash", "dot", "square"]) -> typing.Self:        
        self.__internal.fill = fill
    
        return self
    
    def dash(self, dash: list[float]) -> typing.Self:        
        self.__internal.dash = dash
    
        return self
    

class LineConfig(cogbuilder.Builder[common.LineConfig]):    
    """
    TODO docs
    """
    
    __internal: common.LineConfig

    def __init__(self):
        self.__internal = common.LineConfig()

    def build(self) -> common.LineConfig:
        return self.__internal    
    
    def line_color(self, line_color: str) -> typing.Self:        
        self.__internal.line_color = line_color
    
        return self
    
    def line_width(self, line_width: float) -> typing.Self:        
        self.__internal.line_width = line_width
    
        return self
    
    def line_interpolation(self, line_interpolation: common.LineInterpolation) -> typing.Self:        
        self.__internal.line_interpolation = line_interpolation
    
        return self
    
    def line_style(self, line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self:        
        line_style_resource = line_style.build()
        self.__internal.line_style = line_style_resource
    
        return self
    
    def span_nulls(self, span_nulls: typing.Union[bool, float]) -> typing.Self:    
        """
        Indicate if null values should be treated as gaps or connected.
        When the value is a number, it represents the maximum delta in the
        X axis that should be considered connected.  For timeseries, this is milliseconds
        """
            
        self.__internal.span_nulls = span_nulls
    
        return self
    

class BarConfig(cogbuilder.Builder[common.BarConfig]):    
    """
    TODO docs
    """
    
    __internal: common.BarConfig

    def __init__(self):
        self.__internal = common.BarConfig()

    def build(self) -> common.BarConfig:
        return self.__internal    
    
    def bar_alignment(self, bar_alignment: common.BarAlignment) -> typing.Self:        
        self.__internal.bar_alignment = bar_alignment
    
        return self
    
    def bar_width_factor(self, bar_width_factor: float) -> typing.Self:        
        self.__internal.bar_width_factor = bar_width_factor
    
        return self
    
    def bar_max_width(self, bar_max_width: float) -> typing.Self:        
        self.__internal.bar_max_width = bar_max_width
    
        return self
    

class FillConfig(cogbuilder.Builder[common.FillConfig]):    
    """
    TODO docs
    """
    
    __internal: common.FillConfig

    def __init__(self):
        self.__internal = common.FillConfig()

    def build(self) -> common.FillConfig:
        return self.__internal    
    
    def fill_color(self, fill_color: str) -> typing.Self:        
        self.__internal.fill_color = fill_color
    
        return self
    
    def fill_opacity(self, fill_opacity: float) -> typing.Self:        
        self.__internal.fill_opacity = fill_opacity
    
        return self
    
    def fill_below_to(self, fill_below_to: str) -> typing.Self:        
        self.__internal.fill_below_to = fill_below_to
    
        return self
    

class PointsConfig(cogbuilder.Builder[common.PointsConfig]):    
    """
    TODO docs
    """
    
    __internal: common.PointsConfig

    def __init__(self):
        self.__internal = common.PointsConfig()

    def build(self) -> common.PointsConfig:
        return self.__internal    
    
    def show_points(self, show_points: common.VisibilityMode) -> typing.Self:        
        self.__internal.show_points = show_points
    
        return self
    
    def point_size(self, point_size: float) -> typing.Self:        
        self.__internal.point_size = point_size
    
        return self
    
    def point_color(self, point_color: str) -> typing.Self:        
        self.__internal.point_color = point_color
    
        return self
    
    def point_symbol(self, point_symbol: str) -> typing.Self:        
        self.__internal.point_symbol = point_symbol
    
        return self
    

class ScaleDistributionConfig(cogbuilder.Builder[common.ScaleDistributionConfig]):    
    """
    TODO docs
    """
    
    __internal: common.ScaleDistributionConfig

    def __init__(self):
        self.__internal = common.ScaleDistributionConfig()

    def build(self) -> common.ScaleDistributionConfig:
        return self.__internal    
    
    def type_val(self, type_val: common.ScaleDistribution) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def log(self, log: float) -> typing.Self:        
        self.__internal.log = log
    
        return self
    
    def linear_threshold(self, linear_threshold: float) -> typing.Self:        
        self.__internal.linear_threshold = linear_threshold
    
        return self
    

class AxisConfig(cogbuilder.Builder[common.AxisConfig]):    
    """
    TODO docs
    """
    
    __internal: common.AxisConfig

    def __init__(self):
        self.__internal = common.AxisConfig()

    def build(self) -> common.AxisConfig:
        return self.__internal    
    
    def axis_placement(self, axis_placement: common.AxisPlacement) -> typing.Self:        
        self.__internal.axis_placement = axis_placement
    
        return self
    
    def axis_color_mode(self, axis_color_mode: common.AxisColorMode) -> typing.Self:        
        self.__internal.axis_color_mode = axis_color_mode
    
        return self
    
    def axis_label(self, axis_label: str) -> typing.Self:        
        self.__internal.axis_label = axis_label
    
        return self
    
    def axis_width(self, axis_width: float) -> typing.Self:        
        self.__internal.axis_width = axis_width
    
        return self
    
    def axis_soft_min(self, axis_soft_min: float) -> typing.Self:        
        self.__internal.axis_soft_min = axis_soft_min
    
        return self
    
    def axis_soft_max(self, axis_soft_max: float) -> typing.Self:        
        self.__internal.axis_soft_max = axis_soft_max
    
        return self
    
    def axis_grid_show(self, axis_grid_show: bool) -> typing.Self:        
        self.__internal.axis_grid_show = axis_grid_show
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:        
        scale_distribution_resource = scale_distribution.build()
        self.__internal.scale_distribution = scale_distribution_resource
    
        return self
    
    def axis_centered_zero(self, axis_centered_zero: bool) -> typing.Self:        
        self.__internal.axis_centered_zero = axis_centered_zero
    
        return self
    

class HideSeriesConfig(cogbuilder.Builder[common.HideSeriesConfig]):    
    """
    TODO docs
    """
    
    __internal: common.HideSeriesConfig

    def __init__(self):
        self.__internal = common.HideSeriesConfig()

    def build(self) -> common.HideSeriesConfig:
        return self.__internal    
    
    def tooltip(self, tooltip: bool) -> typing.Self:        
        self.__internal.tooltip = tooltip
    
        return self
    
    def legend(self, legend: bool) -> typing.Self:        
        self.__internal.legend = legend
    
        return self
    
    def viz(self, viz: bool) -> typing.Self:        
        self.__internal.viz = viz
    
        return self
    

class StackingConfig(cogbuilder.Builder[common.StackingConfig]):    
    """
    TODO docs
    """
    
    __internal: common.StackingConfig

    def __init__(self):
        self.__internal = common.StackingConfig()

    def build(self) -> common.StackingConfig:
        return self.__internal    
    
    def mode(self, mode: common.StackingMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    
    def group(self, group: str) -> typing.Self:        
        self.__internal.group = group
    
        return self
    

class StackableFieldConfig(cogbuilder.Builder[common.StackableFieldConfig]):    
    """
    TODO docs
    """
    
    __internal: common.StackableFieldConfig

    def __init__(self):
        self.__internal = common.StackableFieldConfig()

    def build(self) -> common.StackableFieldConfig:
        return self.__internal    
    
    def stacking(self, stacking: cogbuilder.Builder[common.StackingConfig]) -> typing.Self:        
        stacking_resource = stacking.build()
        self.__internal.stacking = stacking_resource
    
        return self
    

class HideableFieldConfig(cogbuilder.Builder[common.HideableFieldConfig]):    
    """
    TODO docs
    """
    
    __internal: common.HideableFieldConfig

    def __init__(self):
        self.__internal = common.HideableFieldConfig()

    def build(self) -> common.HideableFieldConfig:
        return self.__internal    
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:        
        hide_from_resource = hide_from.build()
        self.__internal.hide_from = hide_from_resource
    
        return self
    

class GraphThresholdsStyleConfig(cogbuilder.Builder[common.GraphThresholdsStyleConfig]):    
    """
    TODO docs
    """
    
    __internal: common.GraphThresholdsStyleConfig

    def __init__(self):
        self.__internal = common.GraphThresholdsStyleConfig()

    def build(self) -> common.GraphThresholdsStyleConfig:
        return self.__internal    
    
    def mode(self, mode: common.GraphTresholdsStyleMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    

class SingleStatBaseOptions(cogbuilder.Builder[common.SingleStatBaseOptions]):    
    """
    TODO docs
    """
    
    __internal: common.SingleStatBaseOptions

    def __init__(self):
        self.__internal = common.SingleStatBaseOptions()

    def build(self) -> common.SingleStatBaseOptions:
        return self.__internal    
    
    def reduce_options(self, reduce_options: cogbuilder.Builder[common.ReduceDataOptions]) -> typing.Self:        
        reduce_options_resource = reduce_options.build()
        self.__internal.reduce_options = reduce_options_resource
    
        return self
    
    def text(self, text: cogbuilder.Builder[common.VizTextDisplayOptions]) -> typing.Self:        
        text_resource = text.build()
        self.__internal.text = text_resource
    
        return self
    
    def orientation(self, orientation: common.VizOrientation) -> typing.Self:        
        self.__internal.orientation = orientation
    
        return self
    

class ReduceDataOptions(cogbuilder.Builder[common.ReduceDataOptions]):    
    """
    TODO docs
    """
    
    __internal: common.ReduceDataOptions

    def __init__(self):
        self.__internal = common.ReduceDataOptions()

    def build(self) -> common.ReduceDataOptions:
        return self.__internal    
    
    def values(self, values: bool) -> typing.Self:    
        """
        If true show each row value
        """
            
        self.__internal.values = values
    
        return self
    
    def limit(self, limit: float) -> typing.Self:    
        """
        if showing all values limit
        """
            
        self.__internal.limit = limit
    
        return self
    
    def calcs(self, calcs: list[str]) -> typing.Self:    
        """
        When !values, pick one value for the whole field
        """
            
        self.__internal.calcs = calcs
    
        return self
    
    def fields(self, fields: str) -> typing.Self:    
        """
        Which fields to show.  By default this is only numeric fields
        """
            
        self.__internal.fields = fields
    
        return self
    

class OptionsWithTooltip(cogbuilder.Builder[common.OptionsWithTooltip]):    
    """
    TODO docs
    """
    
    __internal: common.OptionsWithTooltip

    def __init__(self):
        self.__internal = common.OptionsWithTooltip()

    def build(self) -> common.OptionsWithTooltip:
        return self.__internal    
    
    def tooltip(self, tooltip: cogbuilder.Builder[common.VizTooltipOptions]) -> typing.Self:        
        tooltip_resource = tooltip.build()
        self.__internal.tooltip = tooltip_resource
    
        return self
    

class OptionsWithLegend(cogbuilder.Builder[common.OptionsWithLegend]):    
    """
    TODO docs
    """
    
    __internal: common.OptionsWithLegend

    def __init__(self):
        self.__internal = common.OptionsWithLegend()

    def build(self) -> common.OptionsWithLegend:
        return self.__internal    
    
    def legend(self, legend: cogbuilder.Builder[common.VizLegendOptions]) -> typing.Self:        
        legend_resource = legend.build()
        self.__internal.legend = legend_resource
    
        return self
    

class OptionsWithTimezones(cogbuilder.Builder[common.OptionsWithTimezones]):    
    """
    TODO docs
    """
    
    __internal: common.OptionsWithTimezones

    def __init__(self):
        self.__internal = common.OptionsWithTimezones()

    def build(self) -> common.OptionsWithTimezones:
        return self.__internal    
    
    def timezone(self, timezone: list[common.TimeZone]) -> typing.Self:        
        self.__internal.timezone = timezone
    
        return self
    

class OptionsWithTextFormatting(cogbuilder.Builder[common.OptionsWithTextFormatting]):    
    """
    TODO docs
    """
    
    __internal: common.OptionsWithTextFormatting

    def __init__(self):
        self.__internal = common.OptionsWithTextFormatting()

    def build(self) -> common.OptionsWithTextFormatting:
        return self.__internal    
    
    def text(self, text: cogbuilder.Builder[common.VizTextDisplayOptions]) -> typing.Self:        
        text_resource = text.build()
        self.__internal.text = text_resource
    
        return self
    

class VizTextDisplayOptions(cogbuilder.Builder[common.VizTextDisplayOptions]):    
    """
    TODO docs
    """
    
    __internal: common.VizTextDisplayOptions

    def __init__(self):
        self.__internal = common.VizTextDisplayOptions()

    def build(self) -> common.VizTextDisplayOptions:
        return self.__internal    
    
    def title_size(self, title_size: float) -> typing.Self:    
        """
        Explicit title text size
        """
            
        self.__internal.title_size = title_size
    
        return self
    
    def value_size(self, value_size: float) -> typing.Self:    
        """
        Explicit value text size
        """
            
        self.__internal.value_size = value_size
    
        return self
    

class GraphFieldConfig(cogbuilder.Builder[common.GraphFieldConfig]):    
    """
    TODO docs
    """
    
    __internal: common.GraphFieldConfig

    def __init__(self):
        self.__internal = common.GraphFieldConfig()

    def build(self) -> common.GraphFieldConfig:
        return self.__internal    
    
    def draw_style(self, draw_style: common.GraphDrawStyle) -> typing.Self:        
        self.__internal.draw_style = draw_style
    
        return self
    
    def gradient_mode(self, gradient_mode: common.GraphGradientMode) -> typing.Self:        
        self.__internal.gradient_mode = gradient_mode
    
        return self
    
    def thresholds_style(self, thresholds_style: cogbuilder.Builder[common.GraphThresholdsStyleConfig]) -> typing.Self:        
        thresholds_style_resource = thresholds_style.build()
        self.__internal.thresholds_style = thresholds_style_resource
    
        return self
    
    def line_color(self, line_color: str) -> typing.Self:        
        self.__internal.line_color = line_color
    
        return self
    
    def line_width(self, line_width: float) -> typing.Self:        
        self.__internal.line_width = line_width
    
        return self
    
    def line_interpolation(self, line_interpolation: common.LineInterpolation) -> typing.Self:        
        self.__internal.line_interpolation = line_interpolation
    
        return self
    
    def line_style(self, line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self:        
        line_style_resource = line_style.build()
        self.__internal.line_style = line_style_resource
    
        return self
    
    def fill_color(self, fill_color: str) -> typing.Self:        
        self.__internal.fill_color = fill_color
    
        return self
    
    def fill_opacity(self, fill_opacity: float) -> typing.Self:        
        self.__internal.fill_opacity = fill_opacity
    
        return self
    
    def show_points(self, show_points: common.VisibilityMode) -> typing.Self:        
        self.__internal.show_points = show_points
    
        return self
    
    def point_size(self, point_size: float) -> typing.Self:        
        self.__internal.point_size = point_size
    
        return self
    
    def point_color(self, point_color: str) -> typing.Self:        
        self.__internal.point_color = point_color
    
        return self
    
    def axis_placement(self, axis_placement: common.AxisPlacement) -> typing.Self:        
        self.__internal.axis_placement = axis_placement
    
        return self
    
    def axis_color_mode(self, axis_color_mode: common.AxisColorMode) -> typing.Self:        
        self.__internal.axis_color_mode = axis_color_mode
    
        return self
    
    def axis_label(self, axis_label: str) -> typing.Self:        
        self.__internal.axis_label = axis_label
    
        return self
    
    def axis_width(self, axis_width: float) -> typing.Self:        
        self.__internal.axis_width = axis_width
    
        return self
    
    def axis_soft_min(self, axis_soft_min: float) -> typing.Self:        
        self.__internal.axis_soft_min = axis_soft_min
    
        return self
    
    def axis_soft_max(self, axis_soft_max: float) -> typing.Self:        
        self.__internal.axis_soft_max = axis_soft_max
    
        return self
    
    def axis_grid_show(self, axis_grid_show: bool) -> typing.Self:        
        self.__internal.axis_grid_show = axis_grid_show
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:        
        scale_distribution_resource = scale_distribution.build()
        self.__internal.scale_distribution = scale_distribution_resource
    
        return self
    
    def bar_alignment(self, bar_alignment: common.BarAlignment) -> typing.Self:        
        self.__internal.bar_alignment = bar_alignment
    
        return self
    
    def bar_width_factor(self, bar_width_factor: float) -> typing.Self:        
        self.__internal.bar_width_factor = bar_width_factor
    
        return self
    
    def stacking(self, stacking: cogbuilder.Builder[common.StackingConfig]) -> typing.Self:        
        stacking_resource = stacking.build()
        self.__internal.stacking = stacking_resource
    
        return self
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:        
        hide_from_resource = hide_from.build()
        self.__internal.hide_from = hide_from_resource
    
        return self
    
    def transform(self, transform: common.GraphTransform) -> typing.Self:        
        self.__internal.transform = transform
    
        return self
    
    def span_nulls(self, span_nulls: typing.Union[bool, float]) -> typing.Self:    
        """
        Indicate if null values should be treated as gaps or connected.
        When the value is a number, it represents the maximum delta in the
        X axis that should be considered connected.  For timeseries, this is milliseconds
        """
            
        self.__internal.span_nulls = span_nulls
    
        return self
    
    def fill_below_to(self, fill_below_to: str) -> typing.Self:        
        self.__internal.fill_below_to = fill_below_to
    
        return self
    
    def point_symbol(self, point_symbol: str) -> typing.Self:        
        self.__internal.point_symbol = point_symbol
    
        return self
    
    def axis_centered_zero(self, axis_centered_zero: bool) -> typing.Self:        
        self.__internal.axis_centered_zero = axis_centered_zero
    
        return self
    
    def bar_max_width(self, bar_max_width: float) -> typing.Self:        
        self.__internal.bar_max_width = bar_max_width
    
        return self
    

class VizLegendOptions(cogbuilder.Builder[common.VizLegendOptions]):    
    """
    TODO docs
    """
    
    __internal: common.VizLegendOptions

    def __init__(self):
        self.__internal = common.VizLegendOptions()

    def build(self) -> common.VizLegendOptions:
        return self.__internal    
    
    def display_mode(self, display_mode: common.LegendDisplayMode) -> typing.Self:        
        self.__internal.display_mode = display_mode
    
        return self
    
    def placement(self, placement: common.LegendPlacement) -> typing.Self:        
        self.__internal.placement = placement
    
        return self
    
    def show_legend(self, show_legend: bool) -> typing.Self:        
        self.__internal.show_legend = show_legend
    
        return self
    
    def as_table(self, as_table: bool) -> typing.Self:        
        self.__internal.as_table = as_table
    
        return self
    
    def is_visible(self, is_visible: bool) -> typing.Self:        
        self.__internal.is_visible = is_visible
    
        return self
    
    def sort_by(self, sort_by: str) -> typing.Self:        
        self.__internal.sort_by = sort_by
    
        return self
    
    def sort_desc(self, sort_desc: bool) -> typing.Self:        
        self.__internal.sort_desc = sort_desc
    
        return self
    
    def width(self, width: float) -> typing.Self:        
        self.__internal.width = width
    
        return self
    
    def calcs(self, calcs: list[str]) -> typing.Self:        
        self.__internal.calcs = calcs
    
        return self
    

class VizTooltipOptions(cogbuilder.Builder[common.VizTooltipOptions]):    
    """
    TODO docs
    """
    
    __internal: common.VizTooltipOptions

    def __init__(self):
        self.__internal = common.VizTooltipOptions()

    def build(self) -> common.VizTooltipOptions:
        return self.__internal    
    
    def mode(self, mode: common.TooltipDisplayMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    
    def sort(self, sort: common.SortOrder) -> typing.Self:        
        self.__internal.sort = sort
    
        return self
    

class TableSortByFieldState(cogbuilder.Builder[common.TableSortByFieldState]):    
    """
    Sort by field state
    """
    
    __internal: common.TableSortByFieldState

    def __init__(self):
        self.__internal = common.TableSortByFieldState()

    def build(self) -> common.TableSortByFieldState:
        return self.__internal    
    
    def display_name(self, display_name: str) -> typing.Self:    
        """
        Sets the display name of the field to sort by
        """
            
        self.__internal.display_name = display_name
    
        return self
    
    def desc(self, desc: bool) -> typing.Self:    
        """
        Flag used to indicate descending sort order
        """
            
        self.__internal.desc = desc
    
        return self
    

class TableFooterOptions(cogbuilder.Builder[common.TableFooterOptions]):    
    """
    Footer options
    """
    
    __internal: common.TableFooterOptions

    def __init__(self):
        self.__internal = common.TableFooterOptions()

    def build(self) -> common.TableFooterOptions:
        return self.__internal    
    
    def show(self, show: bool) -> typing.Self:        
        self.__internal.show = show
    
        return self
    
    def reducer(self, reducer: list[str]) -> typing.Self:    
        """
        actually 1 value
        """
            
        self.__internal.reducer = reducer
    
        return self
    
    def fields(self, fields: list[str]) -> typing.Self:        
        self.__internal.fields = fields
    
        return self
    
    def enable_pagination(self, enable_pagination: bool) -> typing.Self:        
        self.__internal.enable_pagination = enable_pagination
    
        return self
    
    def count_rows(self, count_rows: bool) -> typing.Self:        
        self.__internal.count_rows = count_rows
    
        return self
    

class TableBarGaugeCellOptions(cogbuilder.Builder[common.TableBarGaugeCellOptions]):    
    """
    Gauge cell options
    """
    
    __internal: common.TableBarGaugeCellOptions

    def __init__(self):
        self.__internal = common.TableBarGaugeCellOptions()        
        self.__internal.type_val = "gauge"

    def build(self) -> common.TableBarGaugeCellOptions:
        return self.__internal    
    
    def mode(self, mode: common.BarGaugeDisplayMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    
    def value_display_mode(self, value_display_mode: common.BarGaugeValueMode) -> typing.Self:        
        self.__internal.value_display_mode = value_display_mode
    
        return self
    

class TableSparklineCellOptions(cogbuilder.Builder[common.TableSparklineCellOptions]):    
    """
    Sparkline cell options
    """
    
    __internal: common.TableSparklineCellOptions

    def __init__(self):
        self.__internal = common.TableSparklineCellOptions()        
        self.__internal.type_val = "sparkline"

    def build(self) -> common.TableSparklineCellOptions:
        return self.__internal    
    
    def draw_style(self, draw_style: common.GraphDrawStyle) -> typing.Self:        
        self.__internal.draw_style = draw_style
    
        return self
    
    def gradient_mode(self, gradient_mode: common.GraphGradientMode) -> typing.Self:        
        self.__internal.gradient_mode = gradient_mode
    
        return self
    
    def thresholds_style(self, thresholds_style: cogbuilder.Builder[common.GraphThresholdsStyleConfig]) -> typing.Self:        
        thresholds_style_resource = thresholds_style.build()
        self.__internal.thresholds_style = thresholds_style_resource
    
        return self
    
    def line_color(self, line_color: str) -> typing.Self:        
        self.__internal.line_color = line_color
    
        return self
    
    def line_width(self, line_width: float) -> typing.Self:        
        self.__internal.line_width = line_width
    
        return self
    
    def line_interpolation(self, line_interpolation: common.LineInterpolation) -> typing.Self:        
        self.__internal.line_interpolation = line_interpolation
    
        return self
    
    def line_style(self, line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self:        
        line_style_resource = line_style.build()
        self.__internal.line_style = line_style_resource
    
        return self
    
    def fill_color(self, fill_color: str) -> typing.Self:        
        self.__internal.fill_color = fill_color
    
        return self
    
    def fill_opacity(self, fill_opacity: float) -> typing.Self:        
        self.__internal.fill_opacity = fill_opacity
    
        return self
    
    def show_points(self, show_points: common.VisibilityMode) -> typing.Self:        
        self.__internal.show_points = show_points
    
        return self
    
    def point_size(self, point_size: float) -> typing.Self:        
        self.__internal.point_size = point_size
    
        return self
    
    def point_color(self, point_color: str) -> typing.Self:        
        self.__internal.point_color = point_color
    
        return self
    
    def axis_placement(self, axis_placement: common.AxisPlacement) -> typing.Self:        
        self.__internal.axis_placement = axis_placement
    
        return self
    
    def axis_color_mode(self, axis_color_mode: common.AxisColorMode) -> typing.Self:        
        self.__internal.axis_color_mode = axis_color_mode
    
        return self
    
    def axis_label(self, axis_label: str) -> typing.Self:        
        self.__internal.axis_label = axis_label
    
        return self
    
    def axis_width(self, axis_width: float) -> typing.Self:        
        self.__internal.axis_width = axis_width
    
        return self
    
    def axis_soft_min(self, axis_soft_min: float) -> typing.Self:        
        self.__internal.axis_soft_min = axis_soft_min
    
        return self
    
    def axis_soft_max(self, axis_soft_max: float) -> typing.Self:        
        self.__internal.axis_soft_max = axis_soft_max
    
        return self
    
    def axis_grid_show(self, axis_grid_show: bool) -> typing.Self:        
        self.__internal.axis_grid_show = axis_grid_show
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:        
        scale_distribution_resource = scale_distribution.build()
        self.__internal.scale_distribution = scale_distribution_resource
    
        return self
    
    def bar_alignment(self, bar_alignment: common.BarAlignment) -> typing.Self:        
        self.__internal.bar_alignment = bar_alignment
    
        return self
    
    def bar_width_factor(self, bar_width_factor: float) -> typing.Self:        
        self.__internal.bar_width_factor = bar_width_factor
    
        return self
    
    def stacking(self, stacking: cogbuilder.Builder[common.StackingConfig]) -> typing.Self:        
        stacking_resource = stacking.build()
        self.__internal.stacking = stacking_resource
    
        return self
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:        
        hide_from_resource = hide_from.build()
        self.__internal.hide_from = hide_from_resource
    
        return self
    
    def transform(self, transform: common.GraphTransform) -> typing.Self:        
        self.__internal.transform = transform
    
        return self
    
    def span_nulls(self, span_nulls: typing.Union[bool, float]) -> typing.Self:    
        """
        Indicate if null values should be treated as gaps or connected.
        When the value is a number, it represents the maximum delta in the
        X axis that should be considered connected.  For timeseries, this is milliseconds
        """
            
        self.__internal.span_nulls = span_nulls
    
        return self
    
    def fill_below_to(self, fill_below_to: str) -> typing.Self:        
        self.__internal.fill_below_to = fill_below_to
    
        return self
    
    def point_symbol(self, point_symbol: str) -> typing.Self:        
        self.__internal.point_symbol = point_symbol
    
        return self
    
    def axis_centered_zero(self, axis_centered_zero: bool) -> typing.Self:        
        self.__internal.axis_centered_zero = axis_centered_zero
    
        return self
    
    def bar_max_width(self, bar_max_width: float) -> typing.Self:        
        self.__internal.bar_max_width = bar_max_width
    
        return self
    

class TableColoredBackgroundCellOptions(cogbuilder.Builder[common.TableColoredBackgroundCellOptions]):    
    """
    Colored background cell options
    """
    
    __internal: common.TableColoredBackgroundCellOptions

    def __init__(self):
        self.__internal = common.TableColoredBackgroundCellOptions()        
        self.__internal.type_val = "color-background"

    def build(self) -> common.TableColoredBackgroundCellOptions:
        return self.__internal    
    
    def mode(self, mode: common.TableCellBackgroundDisplayMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    

class DataSourceRef(cogbuilder.Builder[common.DataSourceRef]):    
    __internal: common.DataSourceRef

    def __init__(self):
        self.__internal = common.DataSourceRef()

    def build(self) -> common.DataSourceRef:
        return self.__internal    
    
    def type_val(self, type_val: str) -> typing.Self:    
        """
        The plugin type-id
        """
            
        self.__internal.type_val = type_val
    
        return self
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Specific datasource instance
        """
            
        self.__internal.uid = uid
    
        return self
    

class ResourceDimensionConfig(cogbuilder.Builder[common.ResourceDimensionConfig]):    
    """
    Links to a resource (image/svg path)
    """
    
    __internal: common.ResourceDimensionConfig

    def __init__(self):
        self.__internal = common.ResourceDimensionConfig()

    def build(self) -> common.ResourceDimensionConfig:
        return self.__internal    
    
    def mode(self, mode: common.ResourceDimensionMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self.__internal.field = field
    
        return self
    
    def fixed(self, fixed: str) -> typing.Self:        
        self.__internal.fixed = fixed
    
        return self
    

class FrameGeometrySource(cogbuilder.Builder[common.FrameGeometrySource]):    
    __internal: common.FrameGeometrySource

    def __init__(self):
        self.__internal = common.FrameGeometrySource()

    def build(self) -> common.FrameGeometrySource:
        return self.__internal    
    
    def mode(self, mode: common.FrameGeometrySourceMode) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    
    def geohash(self, geohash: str) -> typing.Self:    
        """
        Field mappings
        """
            
        self.__internal.geohash = geohash
    
        return self
    
    def latitude(self, latitude: str) -> typing.Self:        
        self.__internal.latitude = latitude
    
        return self
    
    def longitude(self, longitude: str) -> typing.Self:        
        self.__internal.longitude = longitude
    
        return self
    
    def wkt(self, wkt: str) -> typing.Self:        
        self.__internal.wkt = wkt
    
        return self
    
    def lookup(self, lookup: str) -> typing.Self:        
        self.__internal.lookup = lookup
    
        return self
    
    def gazetteer(self, gazetteer: str) -> typing.Self:    
        """
        Path to Gazetteer
        """
            
        self.__internal.gazetteer = gazetteer
    
        return self
    

class HeatmapCalculationOptions(cogbuilder.Builder[common.HeatmapCalculationOptions]):    
    __internal: common.HeatmapCalculationOptions

    def __init__(self):
        self.__internal = common.HeatmapCalculationOptions()

    def build(self) -> common.HeatmapCalculationOptions:
        return self.__internal    
    
    def x_buckets(self, x_buckets: cogbuilder.Builder[common.HeatmapCalculationBucketConfig]) -> typing.Self:    
        """
        The number of buckets to use for the xAxis in the heatmap
        """
            
        x_buckets_resource = x_buckets.build()
        self.__internal.x_buckets = x_buckets_resource
    
        return self
    
    def y_buckets(self, y_buckets: cogbuilder.Builder[common.HeatmapCalculationBucketConfig]) -> typing.Self:    
        """
        The number of buckets to use for the yAxis in the heatmap
        """
            
        y_buckets_resource = y_buckets.build()
        self.__internal.y_buckets = y_buckets_resource
    
        return self
    

class TableFieldOptions(cogbuilder.Builder[common.TableFieldOptions]):    
    """
    Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
    Generally defines alignment, filtering capabilties, display options, etc.
    """
    
    __internal: common.TableFieldOptions

    def __init__(self):
        self.__internal = common.TableFieldOptions()

    def build(self) -> common.TableFieldOptions:
        return self.__internal    
    
    def width(self, width: float) -> typing.Self:        
        self.__internal.width = width
    
        return self
    
    def min_width(self, min_width: float) -> typing.Self:        
        self.__internal.min_width = min_width
    
        return self
    
    def align(self, align: common.FieldTextAlignment) -> typing.Self:        
        self.__internal.align = align
    
        return self
    
    def display_mode(self, display_mode: common.TableCellDisplayMode) -> typing.Self:    
        """
        This field is deprecated in favor of using cellOptions
        """
            
        self.__internal.display_mode = display_mode
    
        return self
    
    def cell_options(self, cell_options: common.TableCellOptions) -> typing.Self:        
        self.__internal.cell_options = cell_options
    
        return self
    
    def hidden(self, hidden: bool) -> typing.Self:    
        """
        ?? default is missing or false ??
        """
            
        self.__internal.hidden = hidden
    
        return self
    
    def inspect(self, inspect: bool) -> typing.Self:        
        self.__internal.inspect = inspect
    
        return self
    
    def filterable(self, filterable: bool) -> typing.Self:        
        self.__internal.filterable = filterable
    
        return self
    
    def hide_header(self, hide_header: bool) -> typing.Self:    
        """
        Hides any header for a column, usefull for columns that show some static content or buttons.
        """
            
        self.__internal.hide_header = hide_header
    
        return self
    
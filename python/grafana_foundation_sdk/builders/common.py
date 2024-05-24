# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import common


class DataSourceJsonData(cogbuilder.Builder[common.DataSourceJsonData]):    
    """
    TODO docs
    """
    
    _internal: common.DataSourceJsonData

    def __init__(self):
        self._internal = common.DataSourceJsonData()

    def build(self) -> common.DataSourceJsonData:
        return self._internal    
    
    def auth_type(self, auth_type: str) -> typing.Self:        
        self._internal.auth_type = auth_type
    
        return self
    
    def default_region(self, default_region: str) -> typing.Self:        
        self._internal.default_region = default_region
    
        return self
    
    def profile(self, profile: str) -> typing.Self:        
        self._internal.profile = profile
    
        return self
    
    def manage_alerts(self, manage_alerts: bool) -> typing.Self:        
        self._internal.manage_alerts = manage_alerts
    
        return self
    
    def alertmanager_uid(self, alertmanager_uid: str) -> typing.Self:        
        self._internal.alertmanager_uid = alertmanager_uid
    
        return self
    

class DataQuery(cogbuilder.Builder[common.DataQuery]):    
    """
    These are the common properties available to all queries in all datasources.
    Specific implementations will *extend* this interface, adding the required
    properties for the given context.
    """
    
    _internal: common.DataQuery

    def __init__(self):
        self._internal = common.DataQuery()

    def build(self) -> common.DataQuery:
        return self._internal    
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        Note this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def datasource(self, datasource: object) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        self._internal.datasource = datasource
    
        return self
    

class BaseDimensionConfig(cogbuilder.Builder[common.BaseDimensionConfig]):    
    _internal: common.BaseDimensionConfig

    def __init__(self):
        self._internal = common.BaseDimensionConfig()

    def build(self) -> common.BaseDimensionConfig:
        return self._internal    
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self._internal.field = field
    
        return self
    

class ScaleDimensionConfig(cogbuilder.Builder[common.ScaleDimensionConfig]):    
    _internal: common.ScaleDimensionConfig

    def __init__(self):
        self._internal = common.ScaleDimensionConfig()

    def build(self) -> common.ScaleDimensionConfig:
        return self._internal    
    
    def min_val(self, min_val: float) -> typing.Self:        
        self._internal.min_val = min_val
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:        
        self._internal.max_val = max_val
    
        return self
    
    def fixed(self, fixed: float) -> typing.Self:        
        self._internal.fixed = fixed
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self._internal.field = field
    
        return self
    
    def mode(self, mode: common.ScaleDimensionMode) -> typing.Self:    
        """
        | *"linear"
        """
            
        self._internal.mode = mode
    
        return self
    

class ColorDimensionConfig(cogbuilder.Builder[common.ColorDimensionConfig]):    
    _internal: common.ColorDimensionConfig

    def __init__(self):
        self._internal = common.ColorDimensionConfig()

    def build(self) -> common.ColorDimensionConfig:
        return self._internal    
    
    def fixed(self, fixed: str) -> typing.Self:    
        """
        color value
        """
            
        self._internal.fixed = fixed
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self._internal.field = field
    
        return self
    

class ScalarDimensionConfig(cogbuilder.Builder[common.ScalarDimensionConfig]):    
    _internal: common.ScalarDimensionConfig

    def __init__(self):
        self._internal = common.ScalarDimensionConfig()

    def build(self) -> common.ScalarDimensionConfig:
        return self._internal    
    
    def min_val(self, min_val: float) -> typing.Self:        
        self._internal.min_val = min_val
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:        
        self._internal.max_val = max_val
    
        return self
    
    def fixed(self, fixed: float) -> typing.Self:        
        self._internal.fixed = fixed
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self._internal.field = field
    
        return self
    
    def mode(self, mode: common.ScalarDimensionMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    

class TextDimensionConfig(cogbuilder.Builder[common.TextDimensionConfig]):    
    _internal: common.TextDimensionConfig

    def __init__(self):
        self._internal = common.TextDimensionConfig()

    def build(self) -> common.TextDimensionConfig:
        return self._internal    
    
    def mode(self, mode: common.TextDimensionMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self._internal.field = field
    
        return self
    
    def fixed(self, fixed: str) -> typing.Self:        
        self._internal.fixed = fixed
    
        return self
    

class MapLayerOptions(cogbuilder.Builder[common.MapLayerOptions]):    
    _internal: common.MapLayerOptions

    def __init__(self):
        self._internal = common.MapLayerOptions()

    def build(self) -> common.MapLayerOptions:
        return self._internal    
    
    def type_val(self, type_val: str) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        """
        configured unique display name
        """
            
        self._internal.name = name
    
        return self
    
    def config(self, config: object) -> typing.Self:    
        """
        Custom options depending on the type
        """
            
        self._internal.config = config
    
        return self
    
    def location(self, location: cogbuilder.Builder[common.FrameGeometrySource]) -> typing.Self:    
        """
        Common method to define geometry fields
        """
            
        location_resource = location.build()
        self._internal.location = location_resource
    
        return self
    
    def filter_data(self, filter_data: object) -> typing.Self:    
        """
        Defines a frame MatcherConfig that may filter data for the given layer
        """
            
        self._internal.filter_data = filter_data
    
        return self
    
    def opacity(self, opacity: int) -> typing.Self:    
        """
        Common properties:
        https://openlayers.org/en/latest/apidoc/module-ol_layer_Base-BaseLayer.html
        Layer opacity (0-1)
        """
            
        self._internal.opacity = opacity
    
        return self
    
    def tooltip(self, tooltip: bool) -> typing.Self:    
        """
        Check tooltip (defaults to true)
        """
            
        self._internal.tooltip = tooltip
    
        return self
    

class HeatmapCalculationBucketConfig(cogbuilder.Builder[common.HeatmapCalculationBucketConfig]):    
    _internal: common.HeatmapCalculationBucketConfig

    def __init__(self):
        self._internal = common.HeatmapCalculationBucketConfig()

    def build(self) -> common.HeatmapCalculationBucketConfig:
        return self._internal    
    
    def mode(self, mode: common.HeatmapCalculationMode) -> typing.Self:    
        """
        Sets the bucket calculation mode
        """
            
        self._internal.mode = mode
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        """
        The number of buckets to use for the axis in the heatmap
        """
            
        self._internal.value = value
    
        return self
    
    def scale(self, scale: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:    
        """
        Controls the scale of the buckets
        """
            
        scale_resource = scale.build()
        self._internal.scale = scale_resource
    
        return self
    

class LineStyle(cogbuilder.Builder[common.LineStyle]):    
    """
    TODO docs
    """
    
    _internal: common.LineStyle

    def __init__(self):
        self._internal = common.LineStyle()

    def build(self) -> common.LineStyle:
        return self._internal    
    
    def fill(self, fill: typing.Literal["solid", "dash", "dot", "square"]) -> typing.Self:        
        self._internal.fill = fill
    
        return self
    
    def dash(self, dash: list[float]) -> typing.Self:        
        self._internal.dash = dash
    
        return self
    

class LineConfig(cogbuilder.Builder[common.LineConfig]):    
    """
    TODO docs
    """
    
    _internal: common.LineConfig

    def __init__(self):
        self._internal = common.LineConfig()

    def build(self) -> common.LineConfig:
        return self._internal    
    
    def line_color(self, line_color: str) -> typing.Self:        
        self._internal.line_color = line_color
    
        return self
    
    def line_width(self, line_width: float) -> typing.Self:        
        self._internal.line_width = line_width
    
        return self
    
    def line_interpolation(self, line_interpolation: common.LineInterpolation) -> typing.Self:        
        self._internal.line_interpolation = line_interpolation
    
        return self
    
    def line_style(self, line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self:        
        line_style_resource = line_style.build()
        self._internal.line_style = line_style_resource
    
        return self
    
    def span_nulls(self, span_nulls: typing.Union[bool, float]) -> typing.Self:    
        """
        Indicate if null values should be treated as gaps or connected.
        When the value is a number, it represents the maximum delta in the
        X axis that should be considered connected.  For timeseries, this is milliseconds
        """
            
        self._internal.span_nulls = span_nulls
    
        return self
    

class BarConfig(cogbuilder.Builder[common.BarConfig]):    
    """
    TODO docs
    """
    
    _internal: common.BarConfig

    def __init__(self):
        self._internal = common.BarConfig()

    def build(self) -> common.BarConfig:
        return self._internal    
    
    def bar_alignment(self, bar_alignment: common.BarAlignment) -> typing.Self:        
        self._internal.bar_alignment = bar_alignment
    
        return self
    
    def bar_width_factor(self, bar_width_factor: float) -> typing.Self:        
        self._internal.bar_width_factor = bar_width_factor
    
        return self
    
    def bar_max_width(self, bar_max_width: float) -> typing.Self:        
        self._internal.bar_max_width = bar_max_width
    
        return self
    

class FillConfig(cogbuilder.Builder[common.FillConfig]):    
    """
    TODO docs
    """
    
    _internal: common.FillConfig

    def __init__(self):
        self._internal = common.FillConfig()

    def build(self) -> common.FillConfig:
        return self._internal    
    
    def fill_color(self, fill_color: str) -> typing.Self:        
        self._internal.fill_color = fill_color
    
        return self
    
    def fill_opacity(self, fill_opacity: float) -> typing.Self:        
        self._internal.fill_opacity = fill_opacity
    
        return self
    
    def fill_below_to(self, fill_below_to: str) -> typing.Self:        
        self._internal.fill_below_to = fill_below_to
    
        return self
    

class PointsConfig(cogbuilder.Builder[common.PointsConfig]):    
    """
    TODO docs
    """
    
    _internal: common.PointsConfig

    def __init__(self):
        self._internal = common.PointsConfig()

    def build(self) -> common.PointsConfig:
        return self._internal    
    
    def show_points(self, show_points: common.VisibilityMode) -> typing.Self:        
        self._internal.show_points = show_points
    
        return self
    
    def point_size(self, point_size: float) -> typing.Self:        
        self._internal.point_size = point_size
    
        return self
    
    def point_color(self, point_color: str) -> typing.Self:        
        self._internal.point_color = point_color
    
        return self
    
    def point_symbol(self, point_symbol: str) -> typing.Self:        
        self._internal.point_symbol = point_symbol
    
        return self
    

class ScaleDistributionConfig(cogbuilder.Builder[common.ScaleDistributionConfig]):    
    """
    TODO docs
    """
    
    _internal: common.ScaleDistributionConfig

    def __init__(self):
        self._internal = common.ScaleDistributionConfig()

    def build(self) -> common.ScaleDistributionConfig:
        return self._internal    
    
    def type_val(self, type_val: common.ScaleDistribution) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def log(self, log: float) -> typing.Self:        
        self._internal.log = log
    
        return self
    
    def linear_threshold(self, linear_threshold: float) -> typing.Self:        
        self._internal.linear_threshold = linear_threshold
    
        return self
    

class AxisConfig(cogbuilder.Builder[common.AxisConfig]):    
    """
    TODO docs
    """
    
    _internal: common.AxisConfig

    def __init__(self):
        self._internal = common.AxisConfig()

    def build(self) -> common.AxisConfig:
        return self._internal    
    
    def axis_placement(self, axis_placement: common.AxisPlacement) -> typing.Self:        
        self._internal.axis_placement = axis_placement
    
        return self
    
    def axis_color_mode(self, axis_color_mode: common.AxisColorMode) -> typing.Self:        
        self._internal.axis_color_mode = axis_color_mode
    
        return self
    
    def axis_label(self, axis_label: str) -> typing.Self:        
        self._internal.axis_label = axis_label
    
        return self
    
    def axis_width(self, axis_width: float) -> typing.Self:        
        self._internal.axis_width = axis_width
    
        return self
    
    def axis_soft_min(self, axis_soft_min: float) -> typing.Self:        
        self._internal.axis_soft_min = axis_soft_min
    
        return self
    
    def axis_soft_max(self, axis_soft_max: float) -> typing.Self:        
        self._internal.axis_soft_max = axis_soft_max
    
        return self
    
    def axis_grid_show(self, axis_grid_show: bool) -> typing.Self:        
        self._internal.axis_grid_show = axis_grid_show
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:        
        scale_distribution_resource = scale_distribution.build()
        self._internal.scale_distribution = scale_distribution_resource
    
        return self
    
    def axis_centered_zero(self, axis_centered_zero: bool) -> typing.Self:        
        self._internal.axis_centered_zero = axis_centered_zero
    
        return self
    

class HideSeriesConfig(cogbuilder.Builder[common.HideSeriesConfig]):    
    """
    TODO docs
    """
    
    _internal: common.HideSeriesConfig

    def __init__(self):
        self._internal = common.HideSeriesConfig()

    def build(self) -> common.HideSeriesConfig:
        return self._internal    
    
    def tooltip(self, tooltip: bool) -> typing.Self:        
        self._internal.tooltip = tooltip
    
        return self
    
    def legend(self, legend: bool) -> typing.Self:        
        self._internal.legend = legend
    
        return self
    
    def viz(self, viz: bool) -> typing.Self:        
        self._internal.viz = viz
    
        return self
    

class StackingConfig(cogbuilder.Builder[common.StackingConfig]):    
    """
    TODO docs
    """
    
    _internal: common.StackingConfig

    def __init__(self):
        self._internal = common.StackingConfig()

    def build(self) -> common.StackingConfig:
        return self._internal    
    
    def mode(self, mode: common.StackingMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    
    def group(self, group: str) -> typing.Self:        
        self._internal.group = group
    
        return self
    

class StackableFieldConfig(cogbuilder.Builder[common.StackableFieldConfig]):    
    """
    TODO docs
    """
    
    _internal: common.StackableFieldConfig

    def __init__(self):
        self._internal = common.StackableFieldConfig()

    def build(self) -> common.StackableFieldConfig:
        return self._internal    
    
    def stacking(self, stacking: cogbuilder.Builder[common.StackingConfig]) -> typing.Self:        
        stacking_resource = stacking.build()
        self._internal.stacking = stacking_resource
    
        return self
    

class HideableFieldConfig(cogbuilder.Builder[common.HideableFieldConfig]):    
    """
    TODO docs
    """
    
    _internal: common.HideableFieldConfig

    def __init__(self):
        self._internal = common.HideableFieldConfig()

    def build(self) -> common.HideableFieldConfig:
        return self._internal    
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:        
        hide_from_resource = hide_from.build()
        self._internal.hide_from = hide_from_resource
    
        return self
    

class GraphThresholdsStyleConfig(cogbuilder.Builder[common.GraphThresholdsStyleConfig]):    
    """
    TODO docs
    """
    
    _internal: common.GraphThresholdsStyleConfig

    def __init__(self):
        self._internal = common.GraphThresholdsStyleConfig()

    def build(self) -> common.GraphThresholdsStyleConfig:
        return self._internal    
    
    def mode(self, mode: common.GraphTresholdsStyleMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    

class SingleStatBaseOptions(cogbuilder.Builder[common.SingleStatBaseOptions]):    
    """
    TODO docs
    """
    
    _internal: common.SingleStatBaseOptions

    def __init__(self):
        self._internal = common.SingleStatBaseOptions()

    def build(self) -> common.SingleStatBaseOptions:
        return self._internal    
    
    def reduce_options(self, reduce_options: cogbuilder.Builder[common.ReduceDataOptions]) -> typing.Self:        
        reduce_options_resource = reduce_options.build()
        self._internal.reduce_options = reduce_options_resource
    
        return self
    
    def text(self, text: cogbuilder.Builder[common.VizTextDisplayOptions]) -> typing.Self:        
        text_resource = text.build()
        self._internal.text = text_resource
    
        return self
    
    def orientation(self, orientation: common.VizOrientation) -> typing.Self:        
        self._internal.orientation = orientation
    
        return self
    

class ReduceDataOptions(cogbuilder.Builder[common.ReduceDataOptions]):    
    """
    TODO docs
    """
    
    _internal: common.ReduceDataOptions

    def __init__(self):
        self._internal = common.ReduceDataOptions()

    def build(self) -> common.ReduceDataOptions:
        return self._internal    
    
    def values(self, values: bool) -> typing.Self:    
        """
        If true show each row value
        """
            
        self._internal.values = values
    
        return self
    
    def limit(self, limit: float) -> typing.Self:    
        """
        if showing all values limit
        """
            
        self._internal.limit = limit
    
        return self
    
    def calcs(self, calcs: list[str]) -> typing.Self:    
        """
        When !values, pick one value for the whole field
        """
            
        self._internal.calcs = calcs
    
        return self
    
    def fields(self, fields: str) -> typing.Self:    
        """
        Which fields to show.  By default this is only numeric fields
        """
            
        self._internal.fields = fields
    
        return self
    

class OptionsWithTooltip(cogbuilder.Builder[common.OptionsWithTooltip]):    
    """
    TODO docs
    """
    
    _internal: common.OptionsWithTooltip

    def __init__(self):
        self._internal = common.OptionsWithTooltip()

    def build(self) -> common.OptionsWithTooltip:
        return self._internal    
    
    def tooltip(self, tooltip: cogbuilder.Builder[common.VizTooltipOptions]) -> typing.Self:        
        tooltip_resource = tooltip.build()
        self._internal.tooltip = tooltip_resource
    
        return self
    

class OptionsWithLegend(cogbuilder.Builder[common.OptionsWithLegend]):    
    """
    TODO docs
    """
    
    _internal: common.OptionsWithLegend

    def __init__(self):
        self._internal = common.OptionsWithLegend()

    def build(self) -> common.OptionsWithLegend:
        return self._internal    
    
    def legend(self, legend: cogbuilder.Builder[common.VizLegendOptions]) -> typing.Self:        
        legend_resource = legend.build()
        self._internal.legend = legend_resource
    
        return self
    

class OptionsWithTimezones(cogbuilder.Builder[common.OptionsWithTimezones]):    
    """
    TODO docs
    """
    
    _internal: common.OptionsWithTimezones

    def __init__(self):
        self._internal = common.OptionsWithTimezones()

    def build(self) -> common.OptionsWithTimezones:
        return self._internal    
    
    def timezone(self, timezone: list[common.TimeZone]) -> typing.Self:        
        self._internal.timezone = timezone
    
        return self
    

class OptionsWithTextFormatting(cogbuilder.Builder[common.OptionsWithTextFormatting]):    
    """
    TODO docs
    """
    
    _internal: common.OptionsWithTextFormatting

    def __init__(self):
        self._internal = common.OptionsWithTextFormatting()

    def build(self) -> common.OptionsWithTextFormatting:
        return self._internal    
    
    def text(self, text: cogbuilder.Builder[common.VizTextDisplayOptions]) -> typing.Self:        
        text_resource = text.build()
        self._internal.text = text_resource
    
        return self
    

class VizTextDisplayOptions(cogbuilder.Builder[common.VizTextDisplayOptions]):    
    """
    TODO docs
    """
    
    _internal: common.VizTextDisplayOptions

    def __init__(self):
        self._internal = common.VizTextDisplayOptions()

    def build(self) -> common.VizTextDisplayOptions:
        return self._internal    
    
    def title_size(self, title_size: float) -> typing.Self:    
        """
        Explicit title text size
        """
            
        self._internal.title_size = title_size
    
        return self
    
    def value_size(self, value_size: float) -> typing.Self:    
        """
        Explicit value text size
        """
            
        self._internal.value_size = value_size
    
        return self
    

class GraphFieldConfig(cogbuilder.Builder[common.GraphFieldConfig]):    
    """
    TODO docs
    """
    
    _internal: common.GraphFieldConfig

    def __init__(self):
        self._internal = common.GraphFieldConfig()

    def build(self) -> common.GraphFieldConfig:
        return self._internal    
    
    def draw_style(self, draw_style: common.GraphDrawStyle) -> typing.Self:        
        self._internal.draw_style = draw_style
    
        return self
    
    def gradient_mode(self, gradient_mode: common.GraphGradientMode) -> typing.Self:        
        self._internal.gradient_mode = gradient_mode
    
        return self
    
    def thresholds_style(self, thresholds_style: cogbuilder.Builder[common.GraphThresholdsStyleConfig]) -> typing.Self:        
        thresholds_style_resource = thresholds_style.build()
        self._internal.thresholds_style = thresholds_style_resource
    
        return self
    
    def line_color(self, line_color: str) -> typing.Self:        
        self._internal.line_color = line_color
    
        return self
    
    def line_width(self, line_width: float) -> typing.Self:        
        self._internal.line_width = line_width
    
        return self
    
    def line_interpolation(self, line_interpolation: common.LineInterpolation) -> typing.Self:        
        self._internal.line_interpolation = line_interpolation
    
        return self
    
    def line_style(self, line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self:        
        line_style_resource = line_style.build()
        self._internal.line_style = line_style_resource
    
        return self
    
    def fill_color(self, fill_color: str) -> typing.Self:        
        self._internal.fill_color = fill_color
    
        return self
    
    def fill_opacity(self, fill_opacity: float) -> typing.Self:        
        self._internal.fill_opacity = fill_opacity
    
        return self
    
    def show_points(self, show_points: common.VisibilityMode) -> typing.Self:        
        self._internal.show_points = show_points
    
        return self
    
    def point_size(self, point_size: float) -> typing.Self:        
        self._internal.point_size = point_size
    
        return self
    
    def point_color(self, point_color: str) -> typing.Self:        
        self._internal.point_color = point_color
    
        return self
    
    def axis_placement(self, axis_placement: common.AxisPlacement) -> typing.Self:        
        self._internal.axis_placement = axis_placement
    
        return self
    
    def axis_color_mode(self, axis_color_mode: common.AxisColorMode) -> typing.Self:        
        self._internal.axis_color_mode = axis_color_mode
    
        return self
    
    def axis_label(self, axis_label: str) -> typing.Self:        
        self._internal.axis_label = axis_label
    
        return self
    
    def axis_width(self, axis_width: float) -> typing.Self:        
        self._internal.axis_width = axis_width
    
        return self
    
    def axis_soft_min(self, axis_soft_min: float) -> typing.Self:        
        self._internal.axis_soft_min = axis_soft_min
    
        return self
    
    def axis_soft_max(self, axis_soft_max: float) -> typing.Self:        
        self._internal.axis_soft_max = axis_soft_max
    
        return self
    
    def axis_grid_show(self, axis_grid_show: bool) -> typing.Self:        
        self._internal.axis_grid_show = axis_grid_show
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:        
        scale_distribution_resource = scale_distribution.build()
        self._internal.scale_distribution = scale_distribution_resource
    
        return self
    
    def bar_alignment(self, bar_alignment: common.BarAlignment) -> typing.Self:        
        self._internal.bar_alignment = bar_alignment
    
        return self
    
    def bar_width_factor(self, bar_width_factor: float) -> typing.Self:        
        self._internal.bar_width_factor = bar_width_factor
    
        return self
    
    def stacking(self, stacking: cogbuilder.Builder[common.StackingConfig]) -> typing.Self:        
        stacking_resource = stacking.build()
        self._internal.stacking = stacking_resource
    
        return self
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:        
        hide_from_resource = hide_from.build()
        self._internal.hide_from = hide_from_resource
    
        return self
    
    def transform(self, transform: common.GraphTransform) -> typing.Self:        
        self._internal.transform = transform
    
        return self
    
    def span_nulls(self, span_nulls: typing.Union[bool, float]) -> typing.Self:    
        """
        Indicate if null values should be treated as gaps or connected.
        When the value is a number, it represents the maximum delta in the
        X axis that should be considered connected.  For timeseries, this is milliseconds
        """
            
        self._internal.span_nulls = span_nulls
    
        return self
    
    def fill_below_to(self, fill_below_to: str) -> typing.Self:        
        self._internal.fill_below_to = fill_below_to
    
        return self
    
    def point_symbol(self, point_symbol: str) -> typing.Self:        
        self._internal.point_symbol = point_symbol
    
        return self
    
    def axis_centered_zero(self, axis_centered_zero: bool) -> typing.Self:        
        self._internal.axis_centered_zero = axis_centered_zero
    
        return self
    
    def bar_max_width(self, bar_max_width: float) -> typing.Self:        
        self._internal.bar_max_width = bar_max_width
    
        return self
    
    def insert_nulls(self, insert_nulls: typing.Union[bool, int]) -> typing.Self:        
        self._internal.insert_nulls = insert_nulls
    
        return self
    

class VizLegendOptions(cogbuilder.Builder[common.VizLegendOptions]):    
    """
    TODO docs
    """
    
    _internal: common.VizLegendOptions

    def __init__(self):
        self._internal = common.VizLegendOptions()

    def build(self) -> common.VizLegendOptions:
        return self._internal    
    
    def display_mode(self, display_mode: common.LegendDisplayMode) -> typing.Self:        
        self._internal.display_mode = display_mode
    
        return self
    
    def placement(self, placement: common.LegendPlacement) -> typing.Self:        
        self._internal.placement = placement
    
        return self
    
    def show_legend(self, show_legend: bool) -> typing.Self:        
        self._internal.show_legend = show_legend
    
        return self
    
    def as_table(self, as_table: bool) -> typing.Self:        
        self._internal.as_table = as_table
    
        return self
    
    def is_visible(self, is_visible: bool) -> typing.Self:        
        self._internal.is_visible = is_visible
    
        return self
    
    def sort_by(self, sort_by: str) -> typing.Self:        
        self._internal.sort_by = sort_by
    
        return self
    
    def sort_desc(self, sort_desc: bool) -> typing.Self:        
        self._internal.sort_desc = sort_desc
    
        return self
    
    def width(self, width: float) -> typing.Self:        
        self._internal.width = width
    
        return self
    
    def calcs(self, calcs: list[str]) -> typing.Self:        
        self._internal.calcs = calcs
    
        return self
    

class VizTooltipOptions(cogbuilder.Builder[common.VizTooltipOptions]):    
    """
    TODO docs
    """
    
    _internal: common.VizTooltipOptions

    def __init__(self):
        self._internal = common.VizTooltipOptions()

    def build(self) -> common.VizTooltipOptions:
        return self._internal    
    
    def mode(self, mode: common.TooltipDisplayMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    
    def sort(self, sort: common.SortOrder) -> typing.Self:        
        self._internal.sort = sort
    
        return self
    

class TableSortByFieldState(cogbuilder.Builder[common.TableSortByFieldState]):    
    """
    Sort by field state
    """
    
    _internal: common.TableSortByFieldState

    def __init__(self):
        self._internal = common.TableSortByFieldState()

    def build(self) -> common.TableSortByFieldState:
        return self._internal    
    
    def display_name(self, display_name: str) -> typing.Self:    
        """
        Sets the display name of the field to sort by
        """
            
        self._internal.display_name = display_name
    
        return self
    
    def desc(self, desc: bool) -> typing.Self:    
        """
        Flag used to indicate descending sort order
        """
            
        self._internal.desc = desc
    
        return self
    

class TableFooterOptions(cogbuilder.Builder[common.TableFooterOptions]):    
    """
    Footer options
    """
    
    _internal: common.TableFooterOptions

    def __init__(self):
        self._internal = common.TableFooterOptions()

    def build(self) -> common.TableFooterOptions:
        return self._internal    
    
    def show(self, show: bool) -> typing.Self:        
        self._internal.show = show
    
        return self
    
    def reducer(self, reducer: list[str]) -> typing.Self:    
        """
        actually 1 value
        """
            
        self._internal.reducer = reducer
    
        return self
    
    def fields(self, fields: list[str]) -> typing.Self:        
        self._internal.fields = fields
    
        return self
    
    def enable_pagination(self, enable_pagination: bool) -> typing.Self:        
        self._internal.enable_pagination = enable_pagination
    
        return self
    
    def count_rows(self, count_rows: bool) -> typing.Self:        
        self._internal.count_rows = count_rows
    
        return self
    

class TableBarGaugeCellOptions(cogbuilder.Builder[common.TableBarGaugeCellOptions]):    
    """
    Gauge cell options
    """
    
    _internal: common.TableBarGaugeCellOptions

    def __init__(self):
        self._internal = common.TableBarGaugeCellOptions()        
        self._internal.type_val = "gauge"

    def build(self) -> common.TableBarGaugeCellOptions:
        return self._internal    
    
    def mode(self, mode: common.BarGaugeDisplayMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    
    def value_display_mode(self, value_display_mode: common.BarGaugeValueMode) -> typing.Self:        
        self._internal.value_display_mode = value_display_mode
    
        return self
    

class TableSparklineCellOptions(cogbuilder.Builder[common.TableSparklineCellOptions]):    
    """
    Sparkline cell options
    """
    
    _internal: common.TableSparklineCellOptions

    def __init__(self):
        self._internal = common.TableSparklineCellOptions()        
        self._internal.type_val = "sparkline"

    def build(self) -> common.TableSparklineCellOptions:
        return self._internal    
    
    def draw_style(self, draw_style: common.GraphDrawStyle) -> typing.Self:        
        self._internal.draw_style = draw_style
    
        return self
    
    def gradient_mode(self, gradient_mode: common.GraphGradientMode) -> typing.Self:        
        self._internal.gradient_mode = gradient_mode
    
        return self
    
    def thresholds_style(self, thresholds_style: cogbuilder.Builder[common.GraphThresholdsStyleConfig]) -> typing.Self:        
        thresholds_style_resource = thresholds_style.build()
        self._internal.thresholds_style = thresholds_style_resource
    
        return self
    
    def line_color(self, line_color: str) -> typing.Self:        
        self._internal.line_color = line_color
    
        return self
    
    def line_width(self, line_width: float) -> typing.Self:        
        self._internal.line_width = line_width
    
        return self
    
    def line_interpolation(self, line_interpolation: common.LineInterpolation) -> typing.Self:        
        self._internal.line_interpolation = line_interpolation
    
        return self
    
    def line_style(self, line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self:        
        line_style_resource = line_style.build()
        self._internal.line_style = line_style_resource
    
        return self
    
    def fill_color(self, fill_color: str) -> typing.Self:        
        self._internal.fill_color = fill_color
    
        return self
    
    def fill_opacity(self, fill_opacity: float) -> typing.Self:        
        self._internal.fill_opacity = fill_opacity
    
        return self
    
    def show_points(self, show_points: common.VisibilityMode) -> typing.Self:        
        self._internal.show_points = show_points
    
        return self
    
    def point_size(self, point_size: float) -> typing.Self:        
        self._internal.point_size = point_size
    
        return self
    
    def point_color(self, point_color: str) -> typing.Self:        
        self._internal.point_color = point_color
    
        return self
    
    def axis_placement(self, axis_placement: common.AxisPlacement) -> typing.Self:        
        self._internal.axis_placement = axis_placement
    
        return self
    
    def axis_color_mode(self, axis_color_mode: common.AxisColorMode) -> typing.Self:        
        self._internal.axis_color_mode = axis_color_mode
    
        return self
    
    def axis_label(self, axis_label: str) -> typing.Self:        
        self._internal.axis_label = axis_label
    
        return self
    
    def axis_width(self, axis_width: float) -> typing.Self:        
        self._internal.axis_width = axis_width
    
        return self
    
    def axis_soft_min(self, axis_soft_min: float) -> typing.Self:        
        self._internal.axis_soft_min = axis_soft_min
    
        return self
    
    def axis_soft_max(self, axis_soft_max: float) -> typing.Self:        
        self._internal.axis_soft_max = axis_soft_max
    
        return self
    
    def axis_grid_show(self, axis_grid_show: bool) -> typing.Self:        
        self._internal.axis_grid_show = axis_grid_show
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:        
        scale_distribution_resource = scale_distribution.build()
        self._internal.scale_distribution = scale_distribution_resource
    
        return self
    
    def bar_alignment(self, bar_alignment: common.BarAlignment) -> typing.Self:        
        self._internal.bar_alignment = bar_alignment
    
        return self
    
    def bar_width_factor(self, bar_width_factor: float) -> typing.Self:        
        self._internal.bar_width_factor = bar_width_factor
    
        return self
    
    def stacking(self, stacking: cogbuilder.Builder[common.StackingConfig]) -> typing.Self:        
        stacking_resource = stacking.build()
        self._internal.stacking = stacking_resource
    
        return self
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:        
        hide_from_resource = hide_from.build()
        self._internal.hide_from = hide_from_resource
    
        return self
    
    def transform(self, transform: common.GraphTransform) -> typing.Self:        
        self._internal.transform = transform
    
        return self
    
    def span_nulls(self, span_nulls: typing.Union[bool, float]) -> typing.Self:    
        """
        Indicate if null values should be treated as gaps or connected.
        When the value is a number, it represents the maximum delta in the
        X axis that should be considered connected.  For timeseries, this is milliseconds
        """
            
        self._internal.span_nulls = span_nulls
    
        return self
    
    def fill_below_to(self, fill_below_to: str) -> typing.Self:        
        self._internal.fill_below_to = fill_below_to
    
        return self
    
    def point_symbol(self, point_symbol: str) -> typing.Self:        
        self._internal.point_symbol = point_symbol
    
        return self
    
    def axis_centered_zero(self, axis_centered_zero: bool) -> typing.Self:        
        self._internal.axis_centered_zero = axis_centered_zero
    
        return self
    
    def bar_max_width(self, bar_max_width: float) -> typing.Self:        
        self._internal.bar_max_width = bar_max_width
    
        return self
    

class TableColoredBackgroundCellOptions(cogbuilder.Builder[common.TableColoredBackgroundCellOptions]):    
    """
    Colored background cell options
    """
    
    _internal: common.TableColoredBackgroundCellOptions

    def __init__(self):
        self._internal = common.TableColoredBackgroundCellOptions()        
        self._internal.type_val = "color-background"

    def build(self) -> common.TableColoredBackgroundCellOptions:
        return self._internal    
    
    def mode(self, mode: common.TableCellBackgroundDisplayMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    

class DataSourceRef(cogbuilder.Builder[common.DataSourceRef]):    
    _internal: common.DataSourceRef

    def __init__(self):
        self._internal = common.DataSourceRef()

    def build(self) -> common.DataSourceRef:
        return self._internal    
    
    def type_val(self, type_val: str) -> typing.Self:    
        """
        The plugin type-id
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Specific datasource instance
        """
            
        self._internal.uid = uid
    
        return self
    

class ResourceDimensionConfig(cogbuilder.Builder[common.ResourceDimensionConfig]):    
    """
    Links to a resource (image/svg path)
    """
    
    _internal: common.ResourceDimensionConfig

    def __init__(self):
        self._internal = common.ResourceDimensionConfig()

    def build(self) -> common.ResourceDimensionConfig:
        return self._internal    
    
    def mode(self, mode: common.ResourceDimensionMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    
    def field(self, field: str) -> typing.Self:    
        """
        fixed: T -- will be added by each element
        """
            
        self._internal.field = field
    
        return self
    
    def fixed(self, fixed: str) -> typing.Self:        
        self._internal.fixed = fixed
    
        return self
    

class FrameGeometrySource(cogbuilder.Builder[common.FrameGeometrySource]):    
    _internal: common.FrameGeometrySource

    def __init__(self):
        self._internal = common.FrameGeometrySource()

    def build(self) -> common.FrameGeometrySource:
        return self._internal    
    
    def mode(self, mode: common.FrameGeometrySourceMode) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    
    def geohash(self, geohash: str) -> typing.Self:    
        """
        Field mappings
        """
            
        self._internal.geohash = geohash
    
        return self
    
    def latitude(self, latitude: str) -> typing.Self:        
        self._internal.latitude = latitude
    
        return self
    
    def longitude(self, longitude: str) -> typing.Self:        
        self._internal.longitude = longitude
    
        return self
    
    def wkt(self, wkt: str) -> typing.Self:        
        self._internal.wkt = wkt
    
        return self
    
    def lookup(self, lookup: str) -> typing.Self:        
        self._internal.lookup = lookup
    
        return self
    
    def gazetteer(self, gazetteer: str) -> typing.Self:    
        """
        Path to Gazetteer
        """
            
        self._internal.gazetteer = gazetteer
    
        return self
    

class HeatmapCalculationOptions(cogbuilder.Builder[common.HeatmapCalculationOptions]):    
    _internal: common.HeatmapCalculationOptions

    def __init__(self):
        self._internal = common.HeatmapCalculationOptions()

    def build(self) -> common.HeatmapCalculationOptions:
        return self._internal    
    
    def x_buckets(self, x_buckets: cogbuilder.Builder[common.HeatmapCalculationBucketConfig]) -> typing.Self:    
        """
        The number of buckets to use for the xAxis in the heatmap
        """
            
        x_buckets_resource = x_buckets.build()
        self._internal.x_buckets = x_buckets_resource
    
        return self
    
    def y_buckets(self, y_buckets: cogbuilder.Builder[common.HeatmapCalculationBucketConfig]) -> typing.Self:    
        """
        The number of buckets to use for the yAxis in the heatmap
        """
            
        y_buckets_resource = y_buckets.build()
        self._internal.y_buckets = y_buckets_resource
    
        return self
    

class TableFieldOptions(cogbuilder.Builder[common.TableFieldOptions]):    
    """
    Field options for each field within a table (e.g 10, "The String", 64.20, etc.)
    Generally defines alignment, filtering capabilties, display options, etc.
    """
    
    _internal: common.TableFieldOptions

    def __init__(self):
        self._internal = common.TableFieldOptions()

    def build(self) -> common.TableFieldOptions:
        return self._internal    
    
    def width(self, width: float) -> typing.Self:        
        self._internal.width = width
    
        return self
    
    def min_width(self, min_width: float) -> typing.Self:        
        self._internal.min_width = min_width
    
        return self
    
    def align(self, align: common.FieldTextAlignment) -> typing.Self:        
        self._internal.align = align
    
        return self
    
    def display_mode(self, display_mode: common.TableCellDisplayMode) -> typing.Self:    
        """
        This field is deprecated in favor of using cellOptions
        """
            
        self._internal.display_mode = display_mode
    
        return self
    
    def cell_options(self, cell_options: common.TableCellOptions) -> typing.Self:        
        self._internal.cell_options = cell_options
    
        return self
    
    def hidden(self, hidden: bool) -> typing.Self:    
        """
        ?? default is missing or false ??
        """
            
        self._internal.hidden = hidden
    
        return self
    
    def inspect(self, inspect: bool) -> typing.Self:        
        self._internal.inspect = inspect
    
        return self
    
    def filterable(self, filterable: bool) -> typing.Self:        
        self._internal.filterable = filterable
    
        return self
    
    def hide_header(self, hide_header: bool) -> typing.Self:    
        """
        Hides any header for a column, usefull for columns that show some static content or buttons.
        """
            
        self._internal.hide_header = hide_header
    
        return self
    
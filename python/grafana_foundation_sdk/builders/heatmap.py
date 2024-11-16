# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import heatmap
from ..models import common
from ..models import dashboard
from ..cog import variants as cogvariants


class HeatmapColorOptions(cogbuilder.Builder[heatmap.HeatmapColorOptions]):    
    """
    Controls various color options
    """
    
    _internal: heatmap.HeatmapColorOptions

    def __init__(self):
        self._internal = heatmap.HeatmapColorOptions()

    def build(self) -> heatmap.HeatmapColorOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def mode(self, mode: heatmap.HeatmapColorMode) -> typing.Self:    
        """
        Sets the color mode
        """
            
        self._internal.mode = mode
    
        return self
    
    def scheme(self, scheme: str) -> typing.Self:    
        """
        Controls the color scheme used
        """
            
        self._internal.scheme = scheme
    
        return self
    
    def fill(self, fill: str) -> typing.Self:    
        """
        Controls the color fill when in opacity mode
        """
            
        self._internal.fill = fill
    
        return self
    
    def scale(self, scale: heatmap.HeatmapColorScale) -> typing.Self:    
        """
        Controls the color scale
        """
            
        self._internal.scale = scale
    
        return self
    
    def exponent(self, exponent: float) -> typing.Self:    
        """
        Controls the exponent when scale is set to exponential
        """
            
        self._internal.exponent = exponent
    
        return self
    
    def steps(self, steps: int) -> typing.Self:    
        """
        Controls the number of color steps
        """
            
        if not steps >= 2:
            raise ValueError("steps must be >= 2")
        if not steps <= 128:
            raise ValueError("steps must be <= 128")
        self._internal.steps = steps
    
        return self
    
    def reverse(self, reverse: bool) -> typing.Self:    
        """
        Reverses the color scheme
        """
            
        self._internal.reverse = reverse
    
        return self
    
    def min_val(self, min_val: float) -> typing.Self:    
        """
        Sets the minimum value for the color scale
        """
            
        self._internal.min_val = min_val
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:    
        """
        Sets the maximum value for the color scale
        """
            
        self._internal.max_val = max_val
    
        return self
    

class YAxisConfig(cogbuilder.Builder[heatmap.YAxisConfig]):    
    """
    Configuration options for the yAxis
    """
    
    _internal: heatmap.YAxisConfig

    def __init__(self):
        self._internal = heatmap.YAxisConfig()

    def build(self) -> heatmap.YAxisConfig:
        """
        Builds the object.
        """
        return self._internal    
    
    def unit(self, unit: str) -> typing.Self:    
        """
        Sets the yAxis unit
        """
            
        self._internal.unit = unit
    
        return self
    
    def reverse(self, reverse: bool) -> typing.Self:    
        """
        Reverses the yAxis
        """
            
        self._internal.reverse = reverse
    
        return self
    
    def decimals(self, decimals: float) -> typing.Self:    
        """
        Controls the number of decimals for yAxis values
        """
            
        self._internal.decimals = decimals
    
        return self
    
    def min_val(self, min_val: float) -> typing.Self:    
        """
        Sets the minimum value for the yAxis
        """
            
        self._internal.min_val = min_val
    
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
    
    def axis_centered_zero(self, axis_centered_zero: bool) -> typing.Self:        
        self._internal.axis_centered_zero = axis_centered_zero
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:    
        """
        Sets the maximum value for the yAxis
        """
            
        self._internal.max_val = max_val
    
        return self
    
    def axis_border_show(self, axis_border_show: bool) -> typing.Self:        
        self._internal.axis_border_show = axis_border_show
    
        return self
    

class CellValues(cogbuilder.Builder[heatmap.CellValues]):    
    """
    Controls cell value options
    """
    
    _internal: heatmap.CellValues

    def __init__(self):
        self._internal = heatmap.CellValues()

    def build(self) -> heatmap.CellValues:
        """
        Builds the object.
        """
        return self._internal    
    
    def unit(self, unit: str) -> typing.Self:    
        """
        Controls the cell value unit
        """
            
        self._internal.unit = unit
    
        return self
    
    def decimals(self, decimals: float) -> typing.Self:    
        """
        Controls the number of decimals for cell values
        """
            
        self._internal.decimals = decimals
    
        return self
    

class FilterValueRange(cogbuilder.Builder[heatmap.FilterValueRange]):    
    """
    Controls the value filter range
    """
    
    _internal: heatmap.FilterValueRange

    def __init__(self):
        self._internal = heatmap.FilterValueRange()

    def build(self) -> heatmap.FilterValueRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def le(self, le: float) -> typing.Self:    
        """
        Sets the filter range to values less than or equal to the given value
        """
            
        self._internal.le = le
    
        return self
    
    def ge(self, ge: float) -> typing.Self:    
        """
        Sets the filter range to values greater than or equal to the given value
        """
            
        self._internal.ge = ge
    
        return self
    

class HeatmapTooltip(cogbuilder.Builder[heatmap.HeatmapTooltip]):    
    """
    Controls tooltip options
    """
    
    _internal: heatmap.HeatmapTooltip

    def __init__(self):
        self._internal = heatmap.HeatmapTooltip()

    def build(self) -> heatmap.HeatmapTooltip:
        """
        Builds the object.
        """
        return self._internal    
    
    def mode(self, mode: common.TooltipDisplayMode) -> typing.Self:    
        """
        Controls how the tooltip is shown
        """
            
        self._internal.mode = mode
    
        return self
    
    def y_histogram(self, y_histogram: bool) -> typing.Self:    
        """
        Controls if the tooltip shows a histogram of the y-axis values
        """
            
        self._internal.y_histogram = y_histogram
    
        return self
    
    def show_color_scale(self, show_color_scale: bool) -> typing.Self:    
        """
        Controls if the tooltip shows a color scale in header
        """
            
        self._internal.show_color_scale = show_color_scale
    
        return self
    

class HeatmapLegend(cogbuilder.Builder[heatmap.HeatmapLegend]):    
    """
    Controls legend options
    """
    
    _internal: heatmap.HeatmapLegend

    def __init__(self):
        self._internal = heatmap.HeatmapLegend()

    def build(self) -> heatmap.HeatmapLegend:
        """
        Builds the object.
        """
        return self._internal    
    
    def show(self, show: bool) -> typing.Self:    
        """
        Controls if the legend is shown
        """
            
        self._internal.show = show
    
        return self
    

class ExemplarConfig(cogbuilder.Builder[heatmap.ExemplarConfig]):    
    """
    Controls exemplar options
    """
    
    _internal: heatmap.ExemplarConfig

    def __init__(self):
        self._internal = heatmap.ExemplarConfig()

    def build(self) -> heatmap.ExemplarConfig:
        """
        Builds the object.
        """
        return self._internal    
    
    def color(self, color: str) -> typing.Self:    
        """
        Sets the color of the exemplar markers
        """
            
        self._internal.color = color
    
        return self
    

class RowsHeatmapOptions(cogbuilder.Builder[heatmap.RowsHeatmapOptions]):    
    """
    Controls frame rows options
    """
    
    _internal: heatmap.RowsHeatmapOptions

    def __init__(self):
        self._internal = heatmap.RowsHeatmapOptions()

    def build(self) -> heatmap.RowsHeatmapOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: str) -> typing.Self:    
        """
        Sets the name of the cell when not calculating from data
        """
            
        self._internal.value = value
    
        return self
    
    def layout(self, layout: common.HeatmapCellLayout) -> typing.Self:    
        """
        Controls tick alignment when not calculating from data
        """
            
        self._internal.layout = layout
    
        return self
    

class Panel(cogbuilder.Builder[dashboard.Panel]):    
    """
    Dashboard panels are the basic visualization building blocks.
    """
    
    _internal: dashboard.Panel

    def __init__(self):
        self._internal = dashboard.Panel()        
        self._internal.type_val = "heatmap"

    def build(self) -> dashboard.Panel:
        """
        Builds the object.
        """
        return self._internal    
    
    def id_val(self, id_val: int) -> typing.Self:    
        """
        Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
        """
            
        self._internal.id_val = id_val
    
        return self
    
    def targets(self, targets: list[cogbuilder.Builder[cogvariants.Dataquery]]) -> typing.Self:    
        """
        Depends on the panel plugin. See the plugin documentation for details.
        """
            
        targets_resources = [r1.build() for r1 in targets]
        self._internal.targets = targets_resources
    
        return self
    
    def with_target(self, target: cogbuilder.Builder[cogvariants.Dataquery]) -> typing.Self:    
        """
        Depends on the panel plugin. See the plugin documentation for details.
        """
            
        if self._internal.targets is None:
            self._internal.targets = []
        
        target_resource = target.build()
        self._internal.targets.append(target_resource)
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        """
        Panel title.
        """
            
        self._internal.title = title
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Panel description.
        """
            
        self._internal.description = description
    
        return self
    
    def transparent(self, transparent: bool) -> typing.Self:    
        """
        Whether to display the panel without a background.
        """
            
        self._internal.transparent = transparent
    
        return self
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        The datasource used in all targets.
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def grid_pos(self, grid_pos: dashboard.GridPos) -> typing.Self:    
        """
        Grid position.
        """
            
        self._internal.grid_pos = grid_pos
    
        return self
    
    def height(self, h: int) -> typing.Self:    
        """
        Panel height. The height is the number of rows from the top edge of the panel.
        """
            
        if not h > 0:
            raise ValueError("h must be > 0")
        if self._internal.grid_pos is None:
            self._internal.grid_pos = dashboard.GridPos()
        assert isinstance(self._internal.grid_pos, dashboard.GridPos)
        self._internal.grid_pos.h = h
    
        return self
    
    def span(self, w: int) -> typing.Self:    
        """
        Panel width. The width is the number of columns from the left edge of the panel.
        """
            
        if not w > 0:
            raise ValueError("w must be > 0")
        if not w <= 24:
            raise ValueError("w must be <= 24")
        if self._internal.grid_pos is None:
            self._internal.grid_pos = dashboard.GridPos()
        assert isinstance(self._internal.grid_pos, dashboard.GridPos)
        self._internal.grid_pos.w = w
    
        return self
    
    def links(self, links: list[cogbuilder.Builder[dashboard.DashboardLink]]) -> typing.Self:    
        """
        Panel links.
        """
            
        links_resources = [r1.build() for r1 in links]
        self._internal.links = links_resources
    
        return self
    
    def repeat(self, repeat: str) -> typing.Self:    
        """
        Name of template variable to repeat for.
        """
            
        self._internal.repeat = repeat
    
        return self
    
    def repeat_direction(self, repeat_direction: typing.Literal["h", "v"]) -> typing.Self:    
        """
        Direction to repeat in if 'repeat' is set.
        `h` for horizontal, `v` for vertical.
        """
            
        self._internal.repeat_direction = repeat_direction
    
        return self
    
    def max_per_row(self, max_per_row: float) -> typing.Self:    
        """
        Option for repeated panels that controls max items per row
        Only relevant for horizontally repeated panels
        """
            
        self._internal.max_per_row = max_per_row
    
        return self
    
    def max_data_points(self, max_data_points: float) -> typing.Self:    
        """
        The maximum number of data points that the panel queries are retrieving.
        """
            
        self._internal.max_data_points = max_data_points
    
        return self
    
    def transformations(self, transformations: list[dashboard.DataTransformerConfig]) -> typing.Self:    
        """
        List of transformations that are applied to the panel data before rendering.
        When there are multiple transformations, Grafana applies them in the order they are listed.
        Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
        """
            
        self._internal.transformations = transformations
    
        return self
    
    def with_transformation(self, transformation: dashboard.DataTransformerConfig) -> typing.Self:    
        """
        List of transformations that are applied to the panel data before rendering.
        When there are multiple transformations, Grafana applies them in the order they are listed.
        Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
        """
            
        if self._internal.transformations is None:
            self._internal.transformations = []
        
        self._internal.transformations.append(transformation)
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        """
        The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
        This value must be formatted as a number followed by a valid time
        identifier like: "40s", "3d", etc.
        See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
        """
            
        self._internal.interval = interval
    
        return self
    
    def time_from(self, time_from: str) -> typing.Self:    
        """
        Overrides the relative time range for individual panels,
        which causes them to be different than what is selected in
        the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
        time periods or days on the same dashboard.
        The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
        `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
        Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
        See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
        """
            
        self._internal.time_from = time_from
    
        return self
    
    def time_shift(self, time_shift: str) -> typing.Self:    
        """
        Overrides the time range for individual panels by shifting its start and end relative to the time picker.
        For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
        Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
        See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
        """
            
        self._internal.time_shift = time_shift
    
        return self
    
    def hide_time_override(self, hide_time_override: bool) -> typing.Self:    
        """
        Controls if the timeFrom or timeShift overrides are shown in the panel header
        """
            
        self._internal.hide_time_override = hide_time_override
    
        return self
    
    def library_panel(self, library_panel: dashboard.LibraryPanelRef) -> typing.Self:    
        """
        Dynamically load the panel
        """
            
        self._internal.library_panel = library_panel
    
        return self
    
    def display_name(self, display_name: str) -> typing.Self:    
        """
        The display value for this field.  This supports template variables blank is auto
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        self._internal.field_config.defaults.display_name = display_name
    
        return self
    
    def unit(self, unit: str) -> typing.Self:    
        """
        Unit a field should use. The unit you select is applied to all fields except time.
        You can use the units ID availables in Grafana or a custom unit.
        Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
        As custom unit, you can use the following formats:
        `suffix:<suffix>` for custom unit that should go after value.
        `prefix:<prefix>` for custom unit that should go before value.
        `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
        `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
        `count:<unit>` for a custom count unit.
        `currency:<unit>` for custom a currency unit.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        self._internal.field_config.defaults.unit = unit
    
        return self
    
    def decimals(self, decimals: float) -> typing.Self:    
        """
        Specify the number of decimals Grafana includes in the rendered value.
        If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
        For example 1.1234 will display as 1.12 and 100.456 will display as 100.
        To display all decimals, set the unit to `String`.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        self._internal.field_config.defaults.decimals = decimals
    
        return self
    
    def min_val(self, min_val: float) -> typing.Self:    
        """
        The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        self._internal.field_config.defaults.min_val = min_val
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:    
        """
        The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        self._internal.field_config.defaults.max_val = max_val
    
        return self
    
    def mappings(self, mappings: list[dashboard.ValueMapping]) -> typing.Self:    
        """
        Convert input values into a display string
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        self._internal.field_config.defaults.mappings = mappings
    
        return self
    
    def thresholds(self, thresholds: cogbuilder.Builder[dashboard.ThresholdsConfig]) -> typing.Self:    
        """
        Map numeric values to states
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        thresholds_resource = thresholds.build()
        self._internal.field_config.defaults.thresholds = thresholds_resource
    
        return self
    
    def color_scheme(self, color: cogbuilder.Builder[dashboard.FieldColor]) -> typing.Self:    
        """
        Panel color configuration
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        color_resource = color.build()
        self._internal.field_config.defaults.color = color_resource
    
        return self
    
    def data_links(self, links: list[cogbuilder.Builder[dashboard.DashboardLink]]) -> typing.Self:    
        """
        The behavior when clicking on a result
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        links_resources = [r1.build() for r1 in links]
        self._internal.field_config.defaults.links = links_resources
    
        return self
    
    def no_value(self, no_value: str) -> typing.Self:    
        """
        Alternative to empty string
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        self._internal.field_config.defaults.no_value = no_value
    
        return self
    
    def overrides(self, overrides: list[cogbuilder.Builder[dashboard.DashboardFieldConfigSourceOverrides]]) -> typing.Self:    
        """
        Overrides are the options applied to specific fields overriding the defaults.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        overrides_resources = [r1.build() for r1 in overrides]
        self._internal.field_config.overrides = overrides_resources
    
        return self
    
    def with_override(self, matcher: dashboard.MatcherConfig, properties: list[dashboard.DynamicConfigValue]) -> typing.Self:    
        """
        Overrides are the options applied to specific fields overriding the defaults.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.overrides is None:
            self._internal.field_config.overrides = []
        
        self._internal.field_config.overrides.append(dashboard.DashboardFieldConfigSourceOverrides(
            matcher=matcher,
            properties=properties,
        ))
    
        return self
    
    def calculate(self, calculate: bool) -> typing.Self:    
        """
        Controls if the heatmap should be calculated from data
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        self._internal.options.calculate = calculate
    
        return self
    
    def calculation(self, calculation: cogbuilder.Builder[common.HeatmapCalculationOptions]) -> typing.Self:    
        """
        Calculation options for the heatmap
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        calculation_resource = calculation.build()
        self._internal.options.calculation = calculation_resource
    
        return self
    
    def color(self, color: cogbuilder.Builder[heatmap.HeatmapColorOptions]) -> typing.Self:    
        """
        Controls the color options
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        color_resource = color.build()
        self._internal.options.color = color_resource
    
        return self
    
    def filter_values(self, filter_values: cogbuilder.Builder[heatmap.FilterValueRange]) -> typing.Self:    
        """
        Filters values between a given range
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        filter_values_resource = filter_values.build()
        self._internal.options.filter_values = filter_values_resource
    
        return self
    
    def rows_frame(self, rows_frame: cogbuilder.Builder[heatmap.RowsHeatmapOptions]) -> typing.Self:    
        """
        Controls tick alignment and value name when not calculating from data
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        rows_frame_resource = rows_frame.build()
        self._internal.options.rows_frame = rows_frame_resource
    
        return self
    
    def show_value(self, show_value: common.VisibilityMode) -> typing.Self:    
        """
        | *{
        	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
        }
        Controls the display of the value in the cell
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        self._internal.options.show_value = show_value
    
        return self
    
    def cell_gap(self, cell_gap: int) -> typing.Self:    
        """
        Controls gap between cells
        """
            
        if not cell_gap <= 25:
            raise ValueError("cell_gap must be <= 25")
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        self._internal.options.cell_gap = cell_gap
    
        return self
    
    def cell_radius(self, cell_radius: float) -> typing.Self:    
        """
        Controls cell radius
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        self._internal.options.cell_radius = cell_radius
    
        return self
    
    def cell_values(self, cell_values: cogbuilder.Builder[heatmap.CellValues]) -> typing.Self:    
        """
        Controls cell value unit
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        cell_values_resource = cell_values.build()
        self._internal.options.cell_values = cell_values_resource
    
        return self
    
    def y_axis(self, y_axis: cogbuilder.Builder[heatmap.YAxisConfig]) -> typing.Self:    
        """
        Controls yAxis placement
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        y_axis_resource = y_axis.build()
        self._internal.options.y_axis = y_axis_resource
    
        return self
    
    def show_legend(self) -> typing.Self:    
        """
        | *{
        	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
        }
        Controls legend options
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        if self._internal.options.legend is None:
            self._internal.options.legend = heatmap.HeatmapLegend()
        assert isinstance(self._internal.options.legend, heatmap.HeatmapLegend)
        self._internal.options.legend.show = True
    
        return self
    
    def hide_legend(self) -> typing.Self:    
        """
        | *{
        	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
        }
        Controls legend options
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        if self._internal.options.legend is None:
            self._internal.options.legend = heatmap.HeatmapLegend()
        assert isinstance(self._internal.options.legend, heatmap.HeatmapLegend)
        self._internal.options.legend.show = False
    
        return self
    
    def mode(self, mode: common.TooltipDisplayMode) -> typing.Self:    
        """
        Controls how the tooltip is shown
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        if self._internal.options.tooltip is None:
            self._internal.options.tooltip = heatmap.HeatmapTooltip()
        assert isinstance(self._internal.options.tooltip, heatmap.HeatmapTooltip)
        self._internal.options.tooltip.mode = mode
    
        return self
    
    def show_y_histogram(self) -> typing.Self:    
        """
        Controls if the tooltip shows a histogram of the y-axis values
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        if self._internal.options.tooltip is None:
            self._internal.options.tooltip = heatmap.HeatmapTooltip()
        assert isinstance(self._internal.options.tooltip, heatmap.HeatmapTooltip)
        self._internal.options.tooltip.y_histogram = True
    
        return self
    
    def hide_y_histogram(self) -> typing.Self:    
        """
        Controls if the tooltip shows a histogram of the y-axis values
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        if self._internal.options.tooltip is None:
            self._internal.options.tooltip = heatmap.HeatmapTooltip()
        assert isinstance(self._internal.options.tooltip, heatmap.HeatmapTooltip)
        self._internal.options.tooltip.y_histogram = False
    
        return self
    
    def show_color_scale(self, show_color_scale: bool) -> typing.Self:    
        """
        Controls if the tooltip shows a color scale in header
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        if self._internal.options.tooltip is None:
            self._internal.options.tooltip = heatmap.HeatmapTooltip()
        assert isinstance(self._internal.options.tooltip, heatmap.HeatmapTooltip)
        self._internal.options.tooltip.show_color_scale = show_color_scale
    
        return self
    
    def exemplars_color(self, color: str) -> typing.Self:    
        """
        Controls exemplar options
        """
            
        if self._internal.options is None:
            self._internal.options = heatmap.Options()
        assert isinstance(self._internal.options, heatmap.Options)
        if self._internal.options.exemplars is None:
            self._internal.options.exemplars = heatmap.ExemplarConfig()
        assert isinstance(self._internal.options.exemplars, heatmap.ExemplarConfig)
        self._internal.options.exemplars.color = color
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:        
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = heatmap.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, heatmap.FieldConfig)
        scale_distribution_resource = scale_distribution.build()
        self._internal.field_config.defaults.custom.scale_distribution = scale_distribution_resource
    
        return self
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:        
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = heatmap.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, heatmap.FieldConfig)
        hide_from_resource = hide_from.build()
        self._internal.field_config.defaults.custom.hide_from = hide_from_resource
    
        return self
    
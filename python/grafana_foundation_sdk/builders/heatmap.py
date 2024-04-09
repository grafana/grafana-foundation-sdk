# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import dashboard
from ..models import heatmap
from ..cog import variants as cogvariants
from ..models import common


class Panel(cogbuilder.Builder[dashboard.Panel]):    
    """
    Dashboard panels are the basic visualization building blocks.
    """
    
    __internal: dashboard.Panel

    def __init__(self):
        self.__internal = dashboard.Panel()        
        self.__internal.type_val = "heatmap"

    def build(self) -> dashboard.Panel:
        return self.__internal    
    
    def id_val(self, id_val: int) -> typing.Self:    
        """
        Unique identifier of the panel. Generated by Grafana when creating a new panel. It must be unique within a dashboard, but not globally.
        """
            
        self.__internal.id_val = id_val
    
        return self
    
    def with_target(self, targets: cogbuilder.Builder[cogvariants.Dataquery]) -> typing.Self:    
        """
        Depends on the panel plugin. See the plugin documentation for details.
        """
            
        if self.__internal.targets is None:
            self.__internal.targets = []
        
        targets_resource = targets.build()
        self.__internal.targets.append(targets_resource)
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        """
        Panel title.
        """
            
        self.__internal.title = title
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Panel description.
        """
            
        self.__internal.description = description
    
        return self
    
    def transparent(self, transparent: bool) -> typing.Self:    
        """
        Whether to display the panel without a background.
        """
            
        self.__internal.transparent = transparent
    
        return self
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        The datasource used in all targets.
        """
            
        self.__internal.datasource = datasource
    
        return self
    
    def height(self, h: int) -> typing.Self:    
        """
        Panel height. The height is the number of rows from the top edge of the panel.
        """
            
        if not h > 0:
            raise ValueError("h must be > 0")
        if self.__internal.grid_pos is None:
            self.__internal.grid_pos = dashboard.GridPos()
        
        assert isinstance(self.__internal.grid_pos, dashboard.GridPos)
        
        self.__internal.grid_pos.h = h
    
        return self
    
    def span(self, w: int) -> typing.Self:    
        """
        Panel width. The width is the number of columns from the left edge of the panel.
        """
            
        if not w > 0:
            raise ValueError("w must be > 0")
        if not w <= 24:
            raise ValueError("w must be <= 24")
        if self.__internal.grid_pos is None:
            self.__internal.grid_pos = dashboard.GridPos()
        
        assert isinstance(self.__internal.grid_pos, dashboard.GridPos)
        
        self.__internal.grid_pos.w = w
    
        return self
    
    def links(self, links: list[cogbuilder.Builder[dashboard.DashboardLink]]) -> typing.Self:    
        """
        Panel links.
        """
            
        links_resources = [r1.build() for r1 in links]
        self.__internal.links = links_resources
    
        return self
    
    def repeat(self, repeat: str) -> typing.Self:    
        """
        Name of template variable to repeat for.
        """
            
        self.__internal.repeat = repeat
    
        return self
    
    def repeat_direction(self, repeat_direction: typing.Literal["h", "v"]) -> typing.Self:    
        """
        Direction to repeat in if 'repeat' is set.
        `h` for horizontal, `v` for vertical.
        """
            
        self.__internal.repeat_direction = repeat_direction
    
        return self
    
    def max_per_row(self, max_per_row: float) -> typing.Self:    
        """
        Option for repeated panels that controls max items per row
        Only relevant for horizontally repeated panels
        """
            
        self.__internal.max_per_row = max_per_row
    
        return self
    
    def max_data_points(self, max_data_points: float) -> typing.Self:    
        """
        The maximum number of data points that the panel queries are retrieving.
        """
            
        self.__internal.max_data_points = max_data_points
    
        return self
    
    def with_transformation(self, transformations: dashboard.DataTransformerConfig) -> typing.Self:    
        """
        List of transformations that are applied to the panel data before rendering.
        When there are multiple transformations, Grafana applies them in the order they are listed.
        Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
        """
            
        if self.__internal.transformations is None:
            self.__internal.transformations = []
        
        self.__internal.transformations.append(transformations)
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        """
        The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
        This value must be formatted as a number followed by a valid time
        identifier like: "40s", "3d", etc.
        See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
        """
            
        self.__internal.interval = interval
    
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
            
        self.__internal.time_from = time_from
    
        return self
    
    def time_shift(self, time_shift: str) -> typing.Self:    
        """
        Overrides the time range for individual panels by shifting its start and end relative to the time picker.
        For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
        Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
        See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
        """
            
        self.__internal.time_shift = time_shift
    
        return self
    
    def hide_time_override(self, hide_time_override: bool) -> typing.Self:    
        """
        Controls if the timeFrom or timeShift overrides are shown in the panel header
        """
            
        self.__internal.hide_time_override = hide_time_override
    
        return self
    
    def library_panel(self, library_panel: dashboard.LibraryPanelRef) -> typing.Self:    
        """
        Dynamically load the panel
        """
            
        self.__internal.library_panel = library_panel
    
        return self
    
    def cache_timeout(self, cache_timeout: str) -> typing.Self:    
        """
        Sets panel queries cache timeout.
        """
            
        self.__internal.cache_timeout = cache_timeout
    
        return self
    
    def query_caching_ttl(self, query_caching_ttl: float) -> typing.Self:    
        """
        Overrides the data source configured time-to-live for a query cache item in milliseconds
        """
            
        self.__internal.query_caching_ttl = query_caching_ttl
    
        return self
    
    def display_name(self, display_name: str) -> typing.Self:    
        """
        The display value for this field.  This supports template variables blank is auto
        """
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        self.__internal.field_config.defaults.display_name = display_name
    
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
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        self.__internal.field_config.defaults.unit = unit
    
        return self
    
    def decimals(self, decimals: float) -> typing.Self:    
        """
        Specify the number of decimals Grafana includes in the rendered value.
        If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
        For example 1.1234 will display as 1.12 and 100.456 will display as 100.
        To display all decimals, set the unit to `String`.
        """
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        self.__internal.field_config.defaults.decimals = decimals
    
        return self
    
    def min_val(self, min_val: float) -> typing.Self:    
        """
        The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
        """
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        self.__internal.field_config.defaults.min_val = min_val
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:    
        """
        The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
        """
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        self.__internal.field_config.defaults.max_val = max_val
    
        return self
    
    def mappings(self, mappings: list[dashboard.ValueMapping]) -> typing.Self:    
        """
        Convert input values into a display string
        """
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        self.__internal.field_config.defaults.mappings = mappings
    
        return self
    
    def thresholds(self, thresholds: cogbuilder.Builder[dashboard.ThresholdsConfig]) -> typing.Self:    
        """
        Map numeric values to states
        """
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        thresholds_resource = thresholds.build()
        self.__internal.field_config.defaults.thresholds = thresholds_resource
    
        return self
    
    def no_value(self, no_value: str) -> typing.Self:    
        """
        Alternative to empty string
        """
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        self.__internal.field_config.defaults.no_value = no_value
    
        return self
    
    def with_override(self, matcher: dashboard.MatcherConfig, properties: list[dashboard.DynamicConfigValue]) -> typing.Self:    
        """
        Overrides are the options applied to specific fields overriding the defaults.
        """
            
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.overrides is None:
            self.__internal.field_config.overrides = []
        
        self.__internal.field_config.overrides.append(dashboard.DashboardFieldConfigSourceOverrides(
            matcher=matcher,
            properties=properties,
        ))
    
        return self
    
    def calculate(self, calculate: bool) -> typing.Self:    
        """
        Controls if the heatmap should be calculated from data
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.calculate = calculate
    
        return self
    
    def calculation(self, calculation: cogbuilder.Builder[common.HeatmapCalculationOptions]) -> typing.Self:    
        """
        Calculation options for the heatmap
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        calculation_resource = calculation.build()
        self.__internal.options.calculation = calculation_resource
    
        return self
    
    def color(self, color: heatmap.HeatmapColorOptions) -> typing.Self:    
        """
        Controls the color options
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.color = color
    
        return self
    
    def filter_values(self, filter_values: heatmap.FilterValueRange) -> typing.Self:    
        """
        Filters values between a given range
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.filter_values = filter_values
    
        return self
    
    def rows_frame(self, rows_frame: heatmap.RowsHeatmapOptions) -> typing.Self:    
        """
        Controls tick alignment and value name when not calculating from data
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.rows_frame = rows_frame
    
        return self
    
    def show_value(self, show_value: common.VisibilityMode) -> typing.Self:    
        """
        | *{
        	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
        }
        Controls the display of the value in the cell
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.show_value = show_value
    
        return self
    
    def cell_gap(self, cell_gap: int) -> typing.Self:    
        """
        Controls gap between cells
        """
            
        if not cell_gap <= 25:
            raise ValueError("cell_gap must be <= 25")
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.cell_gap = cell_gap
    
        return self
    
    def cell_radius(self, cell_radius: float) -> typing.Self:    
        """
        Controls cell radius
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.cell_radius = cell_radius
    
        return self
    
    def cell_values(self, cell_values: heatmap.CellValues) -> typing.Self:    
        """
        Controls cell value unit
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.cell_values = cell_values
    
        return self
    
    def y_axis(self, y_axis: heatmap.YAxisConfig) -> typing.Self:    
        """
        Controls yAxis placement
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.y_axis = y_axis
    
        return self
    
    def legend(self, legend: heatmap.HeatmapLegend) -> typing.Self:    
        """
        | *{
        	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
        }
        Controls legend options
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.legend = legend
    
        return self
    
    def tooltip(self, tooltip: heatmap.HeatmapTooltip) -> typing.Self:    
        """
        Controls tooltip options
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.tooltip = tooltip
    
        return self
    
    def exemplars_color(self, exemplars: heatmap.ExemplarConfig) -> typing.Self:    
        """
        Controls exemplar options
        """
            
        if self.__internal.options is None:
            self.__internal.options = heatmap.Options()
        
        assert isinstance(self.__internal.options, heatmap.Options)
        
        self.__internal.options.exemplars = exemplars
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:        
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        if self.__internal.field_config.defaults.custom is None:
            self.__internal.field_config.defaults.custom = heatmap.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults.custom, heatmap.FieldConfig)
        
        scale_distribution_resource = scale_distribution.build()
        self.__internal.field_config.defaults.custom.scale_distribution = scale_distribution_resource
    
        return self
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:        
        if self.__internal.field_config is None:
            self.__internal.field_config = dashboard.FieldConfigSource()
        
        assert isinstance(self.__internal.field_config, dashboard.FieldConfigSource)
        
        if self.__internal.field_config.defaults is None:
            self.__internal.field_config.defaults = dashboard.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults, dashboard.FieldConfig)
        
        if self.__internal.field_config.defaults.custom is None:
            self.__internal.field_config.defaults.custom = heatmap.FieldConfig()
        
        assert isinstance(self.__internal.field_config.defaults.custom, heatmap.FieldConfig)
        
        hide_from_resource = hide_from.build()
        self.__internal.field_config.defaults.custom.hide_from = hide_from_resource
    
        return self
    
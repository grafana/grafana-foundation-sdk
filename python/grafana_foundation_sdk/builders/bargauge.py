# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import dashboard
from ..cog import variants as cogvariants
from ..models import common
from ..models import bargauge


class Panel(cogbuilder.Builder[dashboard.Panel]):    
    """
    Dashboard panels are the basic visualization building blocks.
    """
    
    _internal: dashboard.Panel

    def __init__(self):
        self._internal = dashboard.Panel()        
        self._internal.type_val = "bargauge"

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
    
    def override_by_name(self, name: str, properties: list[dashboard.DynamicConfigValue]) -> typing.Self:    
        """
        Adds override rules for a specific field, referred to by its name.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.overrides is None:
            self._internal.field_config.overrides = []
        
        self._internal.field_config.overrides.append(dashboard.DashboardFieldConfigSourceOverrides(
            matcher=dashboard.MatcherConfig(
            id_val="byName",
            options=name,
        ),
            properties=properties,
        ))
    
        return self
    
    def override_by_regexp(self, regexp: str, properties: list[dashboard.DynamicConfigValue]) -> typing.Self:    
        """
        Adds override rules for the fields whose name match the given regexp.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.overrides is None:
            self._internal.field_config.overrides = []
        
        self._internal.field_config.overrides.append(dashboard.DashboardFieldConfigSourceOverrides(
            matcher=dashboard.MatcherConfig(
            id_val="byRegexp",
            options=regexp,
        ),
            properties=properties,
        ))
    
        return self
    
    def override_by_field_type(self, field_type: str, properties: list[dashboard.DynamicConfigValue]) -> typing.Self:    
        """
        Adds override rules for all the fields of the given type.
        """
            
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.overrides is None:
            self._internal.field_config.overrides = []
        
        self._internal.field_config.overrides.append(dashboard.DashboardFieldConfigSourceOverrides(
            matcher=dashboard.MatcherConfig(
            id_val="byType",
            options=field_type,
        ),
            properties=properties,
        ))
    
        return self
    
    def override_by_query(self, query_ref_id: str, properties: list[dashboard.DynamicConfigValue]) -> typing.Self:        
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.overrides is None:
            self._internal.field_config.overrides = []
        
        self._internal.field_config.overrides.append(dashboard.DashboardFieldConfigSourceOverrides(
            matcher=dashboard.MatcherConfig(
            id_val="byFrameRefID",
            options=query_ref_id,
        ),
            properties=properties,
        ))
    
        return self
    
    def display_mode(self, display_mode: common.BarGaugeDisplayMode) -> typing.Self:        
        if self._internal.options is None:
            self._internal.options = bargauge.Options()
        assert isinstance(self._internal.options, bargauge.Options)
        self._internal.options.display_mode = display_mode
    
        return self
    
    def value_mode(self, value_mode: common.BarGaugeValueMode) -> typing.Self:        
        if self._internal.options is None:
            self._internal.options = bargauge.Options()
        assert isinstance(self._internal.options, bargauge.Options)
        self._internal.options.value_mode = value_mode
    
        return self
    
    def show_unfilled(self, show_unfilled: bool) -> typing.Self:        
        if self._internal.options is None:
            self._internal.options = bargauge.Options()
        assert isinstance(self._internal.options, bargauge.Options)
        self._internal.options.show_unfilled = show_unfilled
    
        return self
    
    def min_viz_width(self, min_viz_width: int) -> typing.Self:        
        if self._internal.options is None:
            self._internal.options = bargauge.Options()
        assert isinstance(self._internal.options, bargauge.Options)
        self._internal.options.min_viz_width = min_viz_width
    
        return self
    
    def reduce_options(self, reduce_options: cogbuilder.Builder[common.ReduceDataOptions]) -> typing.Self:        
        if self._internal.options is None:
            self._internal.options = bargauge.Options()
        assert isinstance(self._internal.options, bargauge.Options)
        reduce_options_resource = reduce_options.build()
        self._internal.options.reduce_options = reduce_options_resource
    
        return self
    
    def text(self, text: cogbuilder.Builder[common.VizTextDisplayOptions]) -> typing.Self:        
        if self._internal.options is None:
            self._internal.options = bargauge.Options()
        assert isinstance(self._internal.options, bargauge.Options)
        text_resource = text.build()
        self._internal.options.text = text_resource
    
        return self
    
    def min_viz_height(self, min_viz_height: int) -> typing.Self:        
        if self._internal.options is None:
            self._internal.options = bargauge.Options()
        assert isinstance(self._internal.options, bargauge.Options)
        self._internal.options.min_viz_height = min_viz_height
    
        return self
    
    def orientation(self, orientation: common.VizOrientation) -> typing.Self:        
        if self._internal.options is None:
            self._internal.options = bargauge.Options()
        assert isinstance(self._internal.options, bargauge.Options)
        self._internal.options.orientation = orientation
    
        return self
    
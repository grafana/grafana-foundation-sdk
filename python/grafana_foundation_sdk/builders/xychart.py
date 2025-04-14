# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import xychart
from ..models import dashboard
from ..cog import variants as cogvariants
from ..models import common


class MatcherConfig(cogbuilder.Builder[xychart.MatcherConfig]):    
    """
    NOTE: (copied from dashboard_kind.cue, since not exported)
    Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
    It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
    """
    
    _internal: xychart.MatcherConfig

    def __init__(self):
        self._internal = xychart.MatcherConfig()

    def build(self) -> xychart.MatcherConfig:
        """
        Builds the object.
        """
        return self._internal    
    
    def id(self, id_val: str) -> typing.Self:    
        """
        The matcher id. This is used to find the matcher implementation from registry.
        """
            
        self._internal.id_val = id_val
    
        return self
    
    def options(self, options: object) -> typing.Self:    
        """
        The matcher options. This is specific to the matcher implementation.
        """
            
        self._internal.options = options
    
        return self
    


class XYSeriesConfig(cogbuilder.Builder[xychart.XYSeriesConfig]):
    _internal: xychart.XYSeriesConfig

    def __init__(self):
        self._internal = xychart.XYSeriesConfig()

    def build(self) -> xychart.XYSeriesConfig:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: cogbuilder.Builder[xychart.XychartXYSeriesConfigName]) -> typing.Self:    
        name_resource = name.build()
        self._internal.name = name_resource
    
        return self
    
    def frame(self, frame: cogbuilder.Builder[xychart.XychartXYSeriesConfigFrame]) -> typing.Self:    
        frame_resource = frame.build()
        self._internal.frame = frame_resource
    
        return self
    
    def x(self, x: cogbuilder.Builder[xychart.XychartXYSeriesConfigX]) -> typing.Self:    
        x_resource = x.build()
        self._internal.x = x_resource
    
        return self
    
    def y(self, y: cogbuilder.Builder[xychart.XychartXYSeriesConfigY]) -> typing.Self:    
        y_resource = y.build()
        self._internal.y = y_resource
    
        return self
    
    def color(self, color: cogbuilder.Builder[xychart.XychartXYSeriesConfigColor]) -> typing.Self:    
        color_resource = color.build()
        self._internal.color = color_resource
    
        return self
    
    def size(self, size: cogbuilder.Builder[xychart.XychartXYSeriesConfigSize]) -> typing.Self:    
        size_resource = size.build()
        self._internal.size = size_resource
    
        return self
    


class XychartFieldConfigPointSize(cogbuilder.Builder[xychart.XychartFieldConfigPointSize]):
    _internal: xychart.XychartFieldConfigPointSize

    def __init__(self):
        self._internal = xychart.XychartFieldConfigPointSize()

    def build(self) -> xychart.XychartFieldConfigPointSize:
        """
        Builds the object.
        """
        return self._internal    
    
    def fixed(self, fixed: int) -> typing.Self:    
        if not fixed >= 0:
            raise ValueError("fixed must be >= 0")
        self._internal.fixed = fixed
    
        return self
    
    def min(self, min_val: int) -> typing.Self:    
        if not min_val >= 0:
            raise ValueError("min_val must be >= 0")
        self._internal.min_val = min_val
    
        return self
    
    def max(self, max_val: int) -> typing.Self:    
        if not max_val >= 0:
            raise ValueError("max_val must be >= 0")
        self._internal.max_val = max_val
    
        return self
    


class XychartXYSeriesConfigName(cogbuilder.Builder[xychart.XychartXYSeriesConfigName]):
    _internal: xychart.XychartXYSeriesConfigName

    def __init__(self):
        self._internal = xychart.XychartXYSeriesConfigName()

    def build(self) -> xychart.XychartXYSeriesConfigName:
        """
        Builds the object.
        """
        return self._internal    
    
    def fixed(self, fixed: str) -> typing.Self:    
        self._internal.fixed = fixed
    
        return self
    


class XychartXYSeriesConfigFrame(cogbuilder.Builder[xychart.XychartXYSeriesConfigFrame]):
    _internal: xychart.XychartXYSeriesConfigFrame

    def __init__(self):
        self._internal = xychart.XychartXYSeriesConfigFrame()

    def build(self) -> xychart.XychartXYSeriesConfigFrame:
        """
        Builds the object.
        """
        return self._internal    
    
    def matcher(self, matcher: cogbuilder.Builder[xychart.MatcherConfig]) -> typing.Self:    
        matcher_resource = matcher.build()
        self._internal.matcher = matcher_resource
    
        return self
    


class XychartXYSeriesConfigX(cogbuilder.Builder[xychart.XychartXYSeriesConfigX]):
    _internal: xychart.XychartXYSeriesConfigX

    def __init__(self):
        self._internal = xychart.XychartXYSeriesConfigX()

    def build(self) -> xychart.XychartXYSeriesConfigX:
        """
        Builds the object.
        """
        return self._internal    
    
    def matcher(self, matcher: cogbuilder.Builder[xychart.MatcherConfig]) -> typing.Self:    
        matcher_resource = matcher.build()
        self._internal.matcher = matcher_resource
    
        return self
    


class XychartXYSeriesConfigY(cogbuilder.Builder[xychart.XychartXYSeriesConfigY]):
    _internal: xychart.XychartXYSeriesConfigY

    def __init__(self):
        self._internal = xychart.XychartXYSeriesConfigY()

    def build(self) -> xychart.XychartXYSeriesConfigY:
        """
        Builds the object.
        """
        return self._internal    
    
    def matcher(self, matcher: cogbuilder.Builder[xychart.MatcherConfig]) -> typing.Self:    
        matcher_resource = matcher.build()
        self._internal.matcher = matcher_resource
    
        return self
    


class XychartXYSeriesConfigColor(cogbuilder.Builder[xychart.XychartXYSeriesConfigColor]):
    _internal: xychart.XychartXYSeriesConfigColor

    def __init__(self):
        self._internal = xychart.XychartXYSeriesConfigColor()

    def build(self) -> xychart.XychartXYSeriesConfigColor:
        """
        Builds the object.
        """
        return self._internal    
    
    def matcher(self, matcher: cogbuilder.Builder[xychart.MatcherConfig]) -> typing.Self:    
        matcher_resource = matcher.build()
        self._internal.matcher = matcher_resource
    
        return self
    


class XychartXYSeriesConfigSize(cogbuilder.Builder[xychart.XychartXYSeriesConfigSize]):
    _internal: xychart.XychartXYSeriesConfigSize

    def __init__(self):
        self._internal = xychart.XychartXYSeriesConfigSize()

    def build(self) -> xychart.XychartXYSeriesConfigSize:
        """
        Builds the object.
        """
        return self._internal    
    
    def matcher(self, matcher: cogbuilder.Builder[xychart.MatcherConfig]) -> typing.Self:    
        matcher_resource = matcher.build()
        self._internal.matcher = matcher_resource
    
        return self
    


class Panel(cogbuilder.Builder[dashboard.Panel]):    
    """
    Dashboard panels are the basic visualization building blocks.
    """
    
    _internal: dashboard.Panel

    def __init__(self):
        self._internal = dashboard.Panel()        
        self._internal.type_val = "xychart"

    def build(self) -> dashboard.Panel:
        """
        Builds the object.
        """
        return self._internal    
    
    def id(self, id_val: int) -> typing.Self:    
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
    
    def cache_timeout(self, cache_timeout: str) -> typing.Self:    
        """
        Sets panel queries cache timeout.
        """
            
        self._internal.cache_timeout = cache_timeout
    
        return self
    
    def query_caching_ttl(self, query_caching_ttl: float) -> typing.Self:    
        """
        Overrides the data source configured time-to-live for a query cache item in milliseconds
        """
            
        self._internal.query_caching_ttl = query_caching_ttl
    
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
    
    def min(self, min_val: float) -> typing.Self:    
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
    
    def max(self, max_val: float) -> typing.Self:    
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
    
    def show(self, show: xychart.XYShowMode) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.show = show
    
        return self
    
    def point_size(self, point_size: cogbuilder.Builder[xychart.XychartFieldConfigPointSize]) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        point_size_resource = point_size.build()
        self._internal.field_config.defaults.custom.point_size = point_size_resource
    
        return self
    
    def point_shape(self, point_shape: xychart.PointShape) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.point_shape = point_shape
    
        return self
    
    def point_stroke_width(self, point_stroke_width: int) -> typing.Self:    
        if not point_stroke_width >= 0:
            raise ValueError("point_stroke_width must be >= 0")
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.point_stroke_width = point_stroke_width
    
        return self
    
    def fill_opacity(self, fill_opacity: int) -> typing.Self:    
        if not fill_opacity <= 100:
            raise ValueError("fill_opacity must be <= 100")
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.fill_opacity = fill_opacity
    
        return self
    
    def line_width(self, line_width: int) -> typing.Self:    
        if not line_width >= 0:
            raise ValueError("line_width must be >= 0")
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.line_width = line_width
    
        return self
    
    def hide_from(self, hide_from: cogbuilder.Builder[common.HideSeriesConfig]) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        hide_from_resource = hide_from.build()
        self._internal.field_config.defaults.custom.hide_from = hide_from_resource
    
        return self
    
    def axis_placement(self, axis_placement: common.AxisPlacement) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_placement = axis_placement
    
        return self
    
    def axis_color_mode(self, axis_color_mode: common.AxisColorMode) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_color_mode = axis_color_mode
    
        return self
    
    def axis_label(self, axis_label: str) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_label = axis_label
    
        return self
    
    def axis_width(self, axis_width: float) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_width = axis_width
    
        return self
    
    def axis_soft_min(self, axis_soft_min: float) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_soft_min = axis_soft_min
    
        return self
    
    def axis_soft_max(self, axis_soft_max: float) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_soft_max = axis_soft_max
    
        return self
    
    def axis_grid_show(self, axis_grid_show: bool) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_grid_show = axis_grid_show
    
        return self
    
    def scale_distribution(self, scale_distribution: cogbuilder.Builder[common.ScaleDistributionConfig]) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        scale_distribution_resource = scale_distribution.build()
        self._internal.field_config.defaults.custom.scale_distribution = scale_distribution_resource
    
        return self
    
    def axis_centered_zero(self, axis_centered_zero: bool) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_centered_zero = axis_centered_zero
    
        return self
    
    def line_style(self, line_style: cogbuilder.Builder[common.LineStyle]) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        line_style_resource = line_style.build()
        self._internal.field_config.defaults.custom.line_style = line_style_resource
    
        return self
    
    def axis_border_show(self, axis_border_show: bool) -> typing.Self:    
        if self._internal.field_config is None:
            self._internal.field_config = dashboard.FieldConfigSource()
        assert isinstance(self._internal.field_config, dashboard.FieldConfigSource)
        if self._internal.field_config.defaults is None:
            self._internal.field_config.defaults = dashboard.FieldConfig()
        assert isinstance(self._internal.field_config.defaults, dashboard.FieldConfig)
        if self._internal.field_config.defaults.custom is None:
            self._internal.field_config.defaults.custom = xychart.FieldConfig()
        assert isinstance(self._internal.field_config.defaults.custom, xychart.FieldConfig)
        self._internal.field_config.defaults.custom.axis_border_show = axis_border_show
    
        return self
    
    def mapping(self, mapping: xychart.SeriesMapping) -> typing.Self:    
        if self._internal.options is None:
            self._internal.options = xychart.Options()
        assert isinstance(self._internal.options, xychart.Options)
        self._internal.options.mapping = mapping
    
        return self
    
    def legend(self, legend: cogbuilder.Builder[common.VizLegendOptions]) -> typing.Self:    
        if self._internal.options is None:
            self._internal.options = xychart.Options()
        assert isinstance(self._internal.options, xychart.Options)
        legend_resource = legend.build()
        self._internal.options.legend = legend_resource
    
        return self
    
    def tooltip(self, tooltip: cogbuilder.Builder[common.VizTooltipOptions]) -> typing.Self:    
        if self._internal.options is None:
            self._internal.options = xychart.Options()
        assert isinstance(self._internal.options, xychart.Options)
        tooltip_resource = tooltip.build()
        self._internal.options.tooltip = tooltip_resource
    
        return self
    
    def series(self, series: list[cogbuilder.Builder[xychart.XYSeriesConfig]]) -> typing.Self:    
        if self._internal.options is None:
            self._internal.options = xychart.Options()
        assert isinstance(self._internal.options, xychart.Options)
        series_resources = [r1.build() for r1 in series]
        self._internal.options.series = series_resources
    
        return self
    

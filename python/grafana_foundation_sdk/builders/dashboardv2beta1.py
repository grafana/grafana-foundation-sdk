# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import dashboardv2beta1


class Dashboard(cogbuilder.Builder[dashboardv2beta1.Dashboard]):
    _internal: dashboardv2beta1.Dashboard

    def __init__(self, title: str) -> None:
        self._internal = dashboardv2beta1.Dashboard()        
        self._internal.title = title

    def build(self) -> dashboardv2beta1.Dashboard:
        """
        Builds the object.
        """
        return self._internal    
    
    def annotations(self, annotations: list[cogbuilder.Builder[dashboardv2beta1.AnnotationQueryKind]]) -> typing.Self:    
        annotations_resources = [r1.build() for r1 in annotations]
        self._internal.annotations = annotations_resources
    
        return self
    
    def cursor_sync(self, cursor_sync: dashboardv2beta1.DashboardCursorSync) -> typing.Self:    
        """
        Configuration of dashboard cursor sync behavior.
        "Off" for no shared crosshair or tooltip (default).
        "Crosshair" for shared crosshair.
        "Tooltip" for shared crosshair AND shared tooltip.
        """
            
        self._internal.cursor_sync = cursor_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Description of dashboard.
        """
            
        self._internal.description = description
    
        return self
    
    def editable(self, editable: bool) -> typing.Self:    
        """
        Whether a dashboard is editable or not.
        """
            
        self._internal.editable = editable
    
        return self
    
    def elements(self, elements: dict[str, cogbuilder.Builder[dashboardv2beta1.Element]]) -> typing.Self:    
        elements_resources = { key1: val1.build() for (key1, val1) in elements.items() }
        self._internal.elements = elements_resources
    
        return self
    
    def element(self, key: str, element: cogbuilder.Builder[dashboardv2beta1.Element]) -> typing.Self:    
        if self._internal.elements is None:
            self._internal.elements = {}
        
        element_resource = element.build()
        self._internal.elements[key] = element_resource
    
        return self
    
    def grid_layout(self, grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.GridLayoutKind]) -> typing.Self:    
        grid_layout_kind_resource = grid_layout_kind.build()
        self._internal.layout = grid_layout_kind_resource
    
        return self
    
    def rows_layout(self, rows_layout_kind: cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind]) -> typing.Self:    
        rows_layout_kind_resource = rows_layout_kind.build()
        self._internal.layout = rows_layout_kind_resource
    
        return self
    
    def auto_grid_layout(self, auto_grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind]) -> typing.Self:    
        auto_grid_layout_kind_resource = auto_grid_layout_kind.build()
        self._internal.layout = auto_grid_layout_kind_resource
    
        return self
    
    def tabs_layout(self, tabs_layout_kind: cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind]) -> typing.Self:    
        tabs_layout_kind_resource = tabs_layout_kind.build()
        self._internal.layout = tabs_layout_kind_resource
    
        return self
    
    def links(self, links: list[cogbuilder.Builder[dashboardv2beta1.DashboardLink]]) -> typing.Self:    
        """
        Links with references to other dashboards or external websites.
        """
            
        links_resources = [r1.build() for r1 in links]
        self._internal.links = links_resources
    
        return self
    
    def live_now(self, live_now: bool) -> typing.Self:    
        """
        When set to true, the dashboard will redraw panels at an interval matching the pixel width.
        This will keep data "moving left" regardless of the query refresh rate. This setting helps
        avoid dashboards presenting stale live data.
        """
            
        self._internal.live_now = live_now
    
        return self
    
    def preload(self, preload: bool) -> typing.Self:    
        """
        When set to true, the dashboard will load all panels in the dashboard when it's loaded.
        """
            
        self._internal.preload = preload
    
        return self
    
    def revision(self, revision: int) -> typing.Self:    
        """
        Plugins only. The version of the dashboard installed together with the plugin.
        This is used to determine if the dashboard should be updated when the plugin is updated.
        """
            
        self._internal.revision = revision
    
        return self
    
    def tags(self, tags: list[str]) -> typing.Self:    
        """
        Tags associated with dashboard.
        """
            
        self._internal.tags = tags
    
        return self
    
    def time_settings(self, time_settings: cogbuilder.Builder[dashboardv2beta1.TimeSettingsSpec]) -> typing.Self:    
        time_settings_resource = time_settings.build()
        self._internal.time_settings = time_settings_resource
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        """
        Title of dashboard.
        """
            
        self._internal.title = title
    
        return self
    
    def variables(self, variables: list[cogbuilder.Builder[dashboardv2beta1.VariableKind]]) -> typing.Self:    
        """
        Configured template variables.
        """
            
        variables_resources = [r1.build() for r1 in variables]
        self._internal.variables = variables_resources
    
        return self
    
    def variable(self, variable: cogbuilder.Builder[dashboardv2beta1.VariableKind]) -> typing.Self:    
        """
        Configured template variables.
        """
            
        if self._internal.variables is None:
            self._internal.variables = []
        
        variable_resource = variable.build()
        self._internal.variables.append(variable_resource)
    
        return self
    


class AnnotationQuery(cogbuilder.Builder[dashboardv2beta1.AnnotationQueryKind]):
    _internal: dashboardv2beta1.AnnotationQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AnnotationQueryKind()        
        self._internal.kind = "AnnotationQuery"

    def build(self) -> dashboardv2beta1.AnnotationQueryKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: cogbuilder.Builder[dashboardv2beta1.AnnotationQuerySpec]) -> typing.Self:    
        spec_resource = spec.build()
        self._internal.spec = spec_resource
    
        return self
    
    def query(self, query: cogbuilder.Builder[dashboardv2beta1.DataQueryKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        query_resource = query.build()
        self._internal.spec.query = query_resource
    
        return self
    
    def enable(self, enable: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        self._internal.spec.enable = enable
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        self._internal.spec.hide = hide
    
        return self
    
    def icon_color(self, icon_color: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        self._internal.spec.icon_color = icon_color
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        self._internal.spec.name = name
    
        return self
    
    def built_in(self, built_in: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        self._internal.spec.built_in = built_in
    
        return self
    
    def filter(self, filter_val: cogbuilder.Builder[dashboardv2beta1.AnnotationPanelFilter]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        filter_val_resource = filter_val.build()
        self._internal.spec.filter_val = filter_val_resource
    
        return self
    
    def placement(self, placement: str) -> typing.Self:    
        """
        Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        self._internal.spec.placement = placement
    
        return self
    
    def mappings(self, mappings: dict[str, cogbuilder.Builder[dashboardv2beta1.AnnotationEventFieldMapping]]) -> typing.Self:    
        """
        Mappings define how to convert data frame fields to annotation event fields.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        mappings_resources = { key1: val1.build() for (key1, val1) in mappings.items() }
        self._internal.spec.mappings = mappings_resources
    
        return self
    
    def legacy_options(self, legacy_options: dict[str, object]) -> typing.Self:    
        """
        Catch-all field for datasource-specific properties. Should not be available in as code tooling.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AnnotationQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AnnotationQuerySpec)
        self._internal.spec.legacy_options = legacy_options
    
        return self
    


class AnnotationQuerySpec(cogbuilder.Builder[dashboardv2beta1.AnnotationQuerySpec]):
    _internal: dashboardv2beta1.AnnotationQuerySpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AnnotationQuerySpec()

    def build(self) -> dashboardv2beta1.AnnotationQuerySpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def query(self, query: cogbuilder.Builder[dashboardv2beta1.DataQueryKind]) -> typing.Self:    
        query_resource = query.build()
        self._internal.query = query_resource
    
        return self
    
    def enable(self, enable: bool) -> typing.Self:    
        self._internal.enable = enable
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        self._internal.hide = hide
    
        return self
    
    def icon_color(self, icon_color: str) -> typing.Self:    
        self._internal.icon_color = icon_color
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    
    def built_in(self, built_in: bool) -> typing.Self:    
        self._internal.built_in = built_in
    
        return self
    
    def filter(self, filter_val: cogbuilder.Builder[dashboardv2beta1.AnnotationPanelFilter]) -> typing.Self:    
        filter_val_resource = filter_val.build()
        self._internal.filter_val = filter_val_resource
    
        return self
    
    def placement(self, placement: str) -> typing.Self:    
        """
        Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
        """
            
        self._internal.placement = placement
    
        return self
    
    def mappings(self, mappings: dict[str, cogbuilder.Builder[dashboardv2beta1.AnnotationEventFieldMapping]]) -> typing.Self:    
        """
        Mappings define how to convert data frame fields to annotation event fields.
        """
            
        mappings_resources = { key1: val1.build() for (key1, val1) in mappings.items() }
        self._internal.mappings = mappings_resources
    
        return self
    
    def legacy_options(self, legacy_options: dict[str, object]) -> typing.Self:    
        """
        Catch-all field for datasource-specific properties. Should not be available in as code tooling.
        """
            
        self._internal.legacy_options = legacy_options
    
        return self
    


class DataQueryKind(cogbuilder.Builder[dashboardv2beta1.DataQueryKind]):
    _internal: dashboardv2beta1.DataQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.DataQueryKind()        
        self._internal.kind = "DataQuery"

    def build(self) -> dashboardv2beta1.DataQueryKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def group(self, group: str) -> typing.Self:    
        self._internal.group = group
    
        return self
    
    def version(self, version: str) -> typing.Self:    
        self._internal.version = version
    
        return self
    
    def datasource(self, ref: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) -> typing.Self:    
        """
        New type for datasource reference
        Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
        """
            
        ref_resource = ref.build()
        self._internal.datasource = ref_resource
    
        return self
    
    def spec(self, spec: object) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    


class AnnotationPanelFilter(cogbuilder.Builder[dashboardv2beta1.AnnotationPanelFilter]):
    _internal: dashboardv2beta1.AnnotationPanelFilter

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AnnotationPanelFilter()

    def build(self) -> dashboardv2beta1.AnnotationPanelFilter:
        """
        Builds the object.
        """
        return self._internal    
    
    def exclude(self, exclude: bool) -> typing.Self:    
        """
        Should the specified panels be included or excluded
        """
            
        self._internal.exclude = exclude
    
        return self
    
    def ids(self, ids: list[int]) -> typing.Self:    
        """
        Panel IDs that should be included or excluded
        """
            
        self._internal.ids = ids
    
        return self
    


class AnnotationEventFieldMapping(cogbuilder.Builder[dashboardv2beta1.AnnotationEventFieldMapping]):    
    """
    Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
    """
    
    _internal: dashboardv2beta1.AnnotationEventFieldMapping

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AnnotationEventFieldMapping()

    def build(self) -> dashboardv2beta1.AnnotationEventFieldMapping:
        """
        Builds the object.
        """
        return self._internal    
    
    def source(self, source: str) -> typing.Self:    
        """
        Source type for the field value
        """
            
        self._internal.source = source
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        """
        Constant value to use when source is "text"
        """
            
        self._internal.value = value
    
        return self
    
    def regex(self, regex: str) -> typing.Self:    
        """
        Regular expression to apply to the field value
        """
            
        self._internal.regex = regex
    
        return self
    


class Panel(cogbuilder.Builder[dashboardv2beta1.PanelKind]):
    _internal: dashboardv2beta1.PanelKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.PanelKind()        
        self._internal.kind = "Panel"

    def build(self) -> dashboardv2beta1.PanelKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def id(self, id_val: float) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelSpec)
        self._internal.spec.id_val = id_val
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelSpec)
        self._internal.spec.title = title
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelSpec)
        self._internal.spec.description = description
    
        return self
    
    def links(self, links: list[cogbuilder.Builder[dashboardv2beta1.DataLink]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelSpec)
        links_resources = [r1.build() for r1 in links]
        self._internal.spec.links = links_resources
    
        return self
    
    def data(self, data: cogbuilder.Builder[dashboardv2beta1.QueryGroupKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelSpec)
        data_resource = data.build()
        self._internal.spec.data = data_resource
    
        return self
    
    def visualization(self, viz_config: cogbuilder.Builder[dashboardv2beta1.VizConfigKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelSpec)
        viz_config_resource = viz_config.build()
        self._internal.spec.viz_config = viz_config_resource
    
        return self
    
    def transparent(self, transparent: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelSpec)
        self._internal.spec.transparent = transparent
    
        return self
    


class DataLink(cogbuilder.Builder[dashboardv2beta1.DataLink]):
    _internal: dashboardv2beta1.DataLink

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.DataLink()

    def build(self) -> dashboardv2beta1.DataLink:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def url(self, url: str) -> typing.Self:    
        self._internal.url = url
    
        return self
    
    def target_blank(self, target_blank: bool) -> typing.Self:    
        self._internal.target_blank = target_blank
    
        return self
    


class QueryGroup(cogbuilder.Builder[dashboardv2beta1.QueryGroupKind]):
    _internal: dashboardv2beta1.QueryGroupKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.QueryGroupKind()        
        self._internal.kind = "QueryGroup"

    def build(self) -> dashboardv2beta1.QueryGroupKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def targets(self, queries: list[cogbuilder.Builder[dashboardv2beta1.PanelQueryKind]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryGroupSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryGroupSpec)
        queries_resources = [r1.build() for r1 in queries]
        self._internal.spec.queries = queries_resources
    
        return self
    
    def target(self, querie: cogbuilder.Builder[dashboardv2beta1.PanelQueryKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryGroupSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryGroupSpec)
        if self._internal.spec.queries is None:
            self._internal.spec.queries = []
        
        querie_resource = querie.build()
        self._internal.spec.queries.append(querie_resource)
    
        return self
    
    def transformations(self, transformations: list[cogbuilder.Builder[dashboardv2beta1.TransformationKind]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryGroupSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryGroupSpec)
        transformations_resources = [r1.build() for r1 in transformations]
        self._internal.spec.transformations = transformations_resources
    
        return self
    
    def transformation(self, transformation: cogbuilder.Builder[dashboardv2beta1.TransformationKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryGroupSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryGroupSpec)
        if self._internal.spec.transformations is None:
            self._internal.spec.transformations = []
        
        transformation_resource = transformation.build()
        self._internal.spec.transformations.append(transformation_resource)
    
        return self
    
    def query_options(self, query_options: cogbuilder.Builder[dashboardv2beta1.QueryOptionsSpec]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryGroupSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryGroupSpec)
        query_options_resource = query_options.build()
        self._internal.spec.query_options = query_options_resource
    
        return self
    


class Target(cogbuilder.Builder[dashboardv2beta1.PanelQueryKind]):
    _internal: dashboardv2beta1.PanelQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.PanelQueryKind()        
        self._internal.kind = "PanelQuery"

    def build(self) -> dashboardv2beta1.PanelQueryKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def query(self, query: cogbuilder.Builder[dashboardv2beta1.DataQueryKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelQuerySpec)
        query_resource = query.build()
        self._internal.spec.query = query_resource
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelQuerySpec)
        self._internal.spec.ref_id = ref_id
    
        return self
    
    def hidden(self, hidden: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.PanelQuerySpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.PanelQuerySpec)
        self._internal.spec.hidden = hidden
    
        return self
    


class Transformation(cogbuilder.Builder[dashboardv2beta1.TransformationKind]):
    _internal: dashboardv2beta1.TransformationKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.TransformationKind()

    def build(self) -> dashboardv2beta1.TransformationKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def kind(self, kind: str) -> typing.Self:    
        """
        The kind of a TransformationKind is the transformation ID
        """
            
        self._internal.kind = kind
    
        return self
    
    def id(self, id_val: str) -> typing.Self:    
        """
        Unique identifier of transformer
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DataTransformerConfig()
        assert isinstance(self._internal.spec, dashboardv2beta1.DataTransformerConfig)
        self._internal.spec.id_val = id_val
    
        return self
    
    def disabled(self, disabled: bool) -> typing.Self:    
        """
        Disabled transformations are skipped
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DataTransformerConfig()
        assert isinstance(self._internal.spec, dashboardv2beta1.DataTransformerConfig)
        self._internal.spec.disabled = disabled
    
        return self
    
    def filter(self, filter_val: dashboardv2beta1.MatcherConfig) -> typing.Self:    
        """
        Optional frame matcher. When missing it will be applied to all results
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DataTransformerConfig()
        assert isinstance(self._internal.spec, dashboardv2beta1.DataTransformerConfig)
        self._internal.spec.filter_val = filter_val
    
        return self
    
    def topic(self, topic: dashboardv2beta1.DataTopic) -> typing.Self:    
        """
        Where to pull DataFrames from as input to transformation
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DataTransformerConfig()
        assert isinstance(self._internal.spec, dashboardv2beta1.DataTransformerConfig)
        self._internal.spec.topic = topic
    
        return self
    
    def options(self, options: object) -> typing.Self:    
        """
        Options to be passed to the transformer
        Valid options depend on the transformer id
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DataTransformerConfig()
        assert isinstance(self._internal.spec, dashboardv2beta1.DataTransformerConfig)
        self._internal.spec.options = options
    
        return self
    


class QueryOptionsSpec(cogbuilder.Builder[dashboardv2beta1.QueryOptionsSpec]):
    _internal: dashboardv2beta1.QueryOptionsSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.QueryOptionsSpec()

    def build(self) -> dashboardv2beta1.QueryOptionsSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def time_from(self, time_from: str) -> typing.Self:    
        self._internal.time_from = time_from
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        self._internal.max_data_points = max_data_points
    
        return self
    
    def time_shift(self, time_shift: str) -> typing.Self:    
        self._internal.time_shift = time_shift
    
        return self
    
    def query_caching_ttl(self, query_caching_ttl: int) -> typing.Self:    
        self._internal.query_caching_ttl = query_caching_ttl
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        self._internal.interval = interval
    
        return self
    
    def cache_timeout(self, cache_timeout: str) -> typing.Self:    
        self._internal.cache_timeout = cache_timeout
    
        return self
    
    def hide_time_override(self, hide_time_override: bool) -> typing.Self:    
        self._internal.hide_time_override = hide_time_override
    
        return self
    
    def time_compare(self, time_compare: str) -> typing.Self:    
        self._internal.time_compare = time_compare
    
        return self
    


class VizConfigKind(cogbuilder.Builder[dashboardv2beta1.VizConfigKind]):
    _internal: dashboardv2beta1.VizConfigKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.VizConfigKind()        
        self._internal.kind = "VizConfig"

    def build(self) -> dashboardv2beta1.VizConfigKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def group(self, group: str) -> typing.Self:    
        """
        The group is the plugin ID
        """
            
        self._internal.group = group
    
        return self
    
    def version(self, version: str) -> typing.Self:    
        self._internal.version = version
    
        return self
    
    def spec(self, spec: dashboardv2beta1.VizConfigSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def options(self, options: object) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        self._internal.spec.options = options
    
        return self
    
    def field_config(self, field_config: dashboardv2beta1.FieldConfigSource) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        self._internal.spec.field_config = field_config
    
        return self
    
    def display_name(self, display_name: str) -> typing.Self:    
        """
        The display value for this field.  This supports template variables blank is auto
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.display_name = display_name
    
        return self
    
    def display_name_from_ds(self, display_name_from_ds: str) -> typing.Self:    
        """
        This can be used by data sources that return and explicit naming structure for values and labels
        When this property is configured, this value is used rather than the default naming strategy.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.display_name_from_ds = display_name_from_ds
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Human readable field metadata
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.description = description
    
        return self
    
    def path(self, path: str) -> typing.Self:    
        """
        An explicit path to the field in the datasource.  When the frame meta includes a path,
        This will default to `${frame.meta.path}/${field.name}
        
        When defined, this value can be used as an identifier within the datasource scope, and
        may be used to update the results
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.path = path
    
        return self
    
    def writeable(self, writeable: bool) -> typing.Self:    
        """
        True if data source can write a value to the path. Auth/authz are supported separately
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.writeable = writeable
    
        return self
    
    def filterable(self, filterable: bool) -> typing.Self:    
        """
        True if data source field supports ad-hoc filters
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.filterable = filterable
    
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
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.unit = unit
    
        return self
    
    def decimals(self, decimals: float) -> typing.Self:    
        """
        Specify the number of decimals Grafana includes in the rendered value.
        If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
        For example 1.1234 will display as 1.12 and 100.456 will display as 100.
        To display all decimals, set the unit to `String`.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.decimals = decimals
    
        return self
    
    def min(self, min_val: float) -> typing.Self:    
        """
        The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.min_val = min_val
    
        return self
    
    def max(self, max_val: float) -> typing.Self:    
        """
        The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.max_val = max_val
    
        return self
    
    def mappings(self, mappings: list[dashboardv2beta1.ValueMapping]) -> typing.Self:    
        """
        Convert input values into a display string
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.mappings = mappings
    
        return self
    
    def thresholds(self, thresholds: cogbuilder.Builder[dashboardv2beta1.ThresholdsConfig]) -> typing.Self:    
        """
        Map numeric values to states
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        thresholds_resource = thresholds.build()
        self._internal.spec.field_config.defaults.thresholds = thresholds_resource
    
        return self
    
    def color_scheme(self, color: cogbuilder.Builder[dashboardv2beta1.FieldColor]) -> typing.Self:    
        """
        Panel color configuration
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        color_resource = color.build()
        self._internal.spec.field_config.defaults.color = color_resource
    
        return self
    
    def data_links(self, links: list[object]) -> typing.Self:    
        """
        The behavior when clicking on a result
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.links = links
    
        return self
    
    def actions(self, actions: list[cogbuilder.Builder[dashboardv2beta1.Action]]) -> typing.Self:    
        """
        Define interactive HTTP requests that can be triggered from data visualizations.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        actions_resources = [r1.build() for r1 in actions]
        self._internal.spec.field_config.defaults.actions = actions_resources
    
        return self
    
    def no_value(self, no_value: str) -> typing.Self:    
        """
        Alternative to empty string
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.no_value = no_value
    
        return self
    
    def custom(self, custom: object) -> typing.Self:    
        """
        custom is specified by the FieldConfig field
        in panel plugin schemas.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.custom = custom
    
        return self
    
    def field_min_max(self, field_min_max: bool) -> typing.Self:    
        """
        Calculate min max per field
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.field_min_max = field_min_max
    
        return self
    
    def null_value_mode(self, null_value_mode: dashboardv2beta1.NullValueMode) -> typing.Self:    
        """
        How null values should be handled when calculating field stats
        "null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.defaults is None:
            self._internal.spec.field_config.defaults = dashboardv2beta1.FieldConfig()
        assert isinstance(self._internal.spec.field_config.defaults, dashboardv2beta1.FieldConfig)
        self._internal.spec.field_config.defaults.null_value_mode = null_value_mode
    
        return self
    
    def defaults(self, defaults: dashboardv2beta1.FieldConfig) -> typing.Self:    
        """
        Defaults are the options applied to all fields.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        self._internal.spec.field_config.defaults = defaults
    
        return self
    
    def overrides(self, overrides: list[cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]]) -> typing.Self:    
        """
        Overrides are the options applied to specific fields overriding the defaults.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        overrides_resources = [r1.build() for r1 in overrides]
        self._internal.spec.field_config.overrides = overrides_resources
    
        return self
    
    def override(self, override: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]) -> typing.Self:    
        """
        Overrides are the options applied to specific fields overriding the defaults.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.overrides is None:
            self._internal.spec.field_config.overrides = []
        
        override_resource = override.build()
        self._internal.spec.field_config.overrides.append(override_resource)
    
        return self
    
    def override_by_name(self, name: str, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self:    
        """
        Adds override rules for a specific field, referred to by its name.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.overrides is None:
            self._internal.spec.field_config.overrides = []
        
        self._internal.spec.field_config.overrides.append(dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides(
            matcher=dashboardv2beta1.MatcherConfig(
            id_val="byName",
            options=name,
        ),
            properties=properties,
        ))
    
        return self
    
    def override_by_regexp(self, regexp: str, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self:    
        """
        Adds override rules for the fields whose name match the given regexp.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.overrides is None:
            self._internal.spec.field_config.overrides = []
        
        self._internal.spec.field_config.overrides.append(dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides(
            matcher=dashboardv2beta1.MatcherConfig(
            id_val="byRegexp",
            options=regexp,
        ),
            properties=properties,
        ))
    
        return self
    
    def override_by_field_type(self, field_type: str, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self:    
        """
        Adds override rules for all the fields of the given type.
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.overrides is None:
            self._internal.spec.field_config.overrides = []
        
        self._internal.spec.field_config.overrides.append(dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides(
            matcher=dashboardv2beta1.MatcherConfig(
            id_val="byType",
            options=field_type,
        ),
            properties=properties,
        ))
    
        return self
    
    def override_by_query(self, query_ref_id: str, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.VizConfigSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.VizConfigSpec)
        if self._internal.spec.field_config is None:
            self._internal.spec.field_config = dashboardv2beta1.FieldConfigSource()
        assert isinstance(self._internal.spec.field_config, dashboardv2beta1.FieldConfigSource)
        if self._internal.spec.field_config.overrides is None:
            self._internal.spec.field_config.overrides = []
        
        self._internal.spec.field_config.overrides.append(dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides(
            matcher=dashboardv2beta1.MatcherConfig(
            id_val="byFrameRefID",
            options=query_ref_id,
        ),
            properties=properties,
        ))
    
        return self
    


class ValueMappingResult(cogbuilder.Builder[dashboardv2beta1.ValueMappingResult]):    
    """
    Result used as replacement with text and color when the value matches
    """
    
    _internal: dashboardv2beta1.ValueMappingResult

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ValueMappingResult()

    def build(self) -> dashboardv2beta1.ValueMappingResult:
        """
        Builds the object.
        """
        return self._internal    
    
    def text(self, text: str) -> typing.Self:    
        """
        Text to display when the value matches
        """
            
        self._internal.text = text
    
        return self
    
    def color(self, color: str) -> typing.Self:    
        """
        Text to use when the value matches
        """
            
        self._internal.color = color
    
        return self
    
    def icon(self, icon: str) -> typing.Self:    
        """
        Icon to display when the value matches. Only specific visualizations.
        """
            
        self._internal.icon = icon
    
        return self
    
    def index(self, index: int) -> typing.Self:    
        """
        Position in the mapping array. Only used internally.
        """
            
        self._internal.index = index
    
        return self
    


class ThresholdsConfig(cogbuilder.Builder[dashboardv2beta1.ThresholdsConfig]):
    _internal: dashboardv2beta1.ThresholdsConfig

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ThresholdsConfig()

    def build(self) -> dashboardv2beta1.ThresholdsConfig:
        """
        Builds the object.
        """
        return self._internal    
    
    def mode(self, mode: dashboardv2beta1.ThresholdsMode) -> typing.Self:    
        self._internal.mode = mode
    
        return self
    
    def steps(self, steps: list[dashboardv2beta1.Threshold]) -> typing.Self:    
        self._internal.steps = steps
    
        return self
    


class FieldColor(cogbuilder.Builder[dashboardv2beta1.FieldColor]):    
    """
    Map a field to a color.
    """
    
    _internal: dashboardv2beta1.FieldColor

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.FieldColor()

    def build(self) -> dashboardv2beta1.FieldColor:
        """
        Builds the object.
        """
        return self._internal    
    
    def mode(self, mode: dashboardv2beta1.FieldColorModeId) -> typing.Self:    
        """
        The main color scheme mode.
        """
            
        self._internal.mode = mode
    
        return self
    
    def fixed_color(self, fixed_color: str) -> typing.Self:    
        """
        The fixed color value for fixed or shades color modes.
        """
            
        self._internal.fixed_color = fixed_color
    
        return self
    
    def series_by(self, series_by: dashboardv2beta1.FieldColorSeriesByMode) -> typing.Self:    
        """
        Some visualizations need to know how to assign a series color from by value color schemes.
        """
            
        self._internal.series_by = series_by
    
        return self
    


class Action(cogbuilder.Builder[dashboardv2beta1.Action]):
    _internal: dashboardv2beta1.Action

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Action()

    def build(self) -> dashboardv2beta1.Action:
        """
        Builds the object.
        """
        return self._internal    
    
    def type(self, type_val: dashboardv2beta1.ActionType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def fetch(self, fetch: cogbuilder.Builder[dashboardv2beta1.FetchOptions]) -> typing.Self:    
        fetch_resource = fetch.build()
        self._internal.fetch = fetch_resource
    
        return self
    
    def infinity(self, infinity: cogbuilder.Builder[dashboardv2beta1.InfinityOptions]) -> typing.Self:    
        infinity_resource = infinity.build()
        self._internal.infinity = infinity_resource
    
        return self
    
    def confirmation(self, confirmation: str) -> typing.Self:    
        self._internal.confirmation = confirmation
    
        return self
    
    def one_click(self, one_click: bool) -> typing.Self:    
        self._internal.one_click = one_click
    
        return self
    
    def variables(self, variables: list[cogbuilder.Builder[dashboardv2beta1.ActionVariable]]) -> typing.Self:    
        variables_resources = [r1.build() for r1 in variables]
        self._internal.variables = variables_resources
    
        return self
    
    def style(self, style: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1ActionStyle]) -> typing.Self:    
        style_resource = style.build()
        self._internal.style = style_resource
    
        return self
    


class FetchOptions(cogbuilder.Builder[dashboardv2beta1.FetchOptions]):
    _internal: dashboardv2beta1.FetchOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.FetchOptions()

    def build(self) -> dashboardv2beta1.FetchOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def method(self, method: dashboardv2beta1.HttpRequestMethod) -> typing.Self:    
        self._internal.method = method
    
        return self
    
    def url(self, url: str) -> typing.Self:    
        self._internal.url = url
    
        return self
    
    def body(self, body: str) -> typing.Self:    
        self._internal.body = body
    
        return self
    
    def query_params(self, query_params: list[list[str]]) -> typing.Self:    
        """
        These are 2D arrays of strings, each representing a key-value pair
        We are defining them this way because we can't generate a go struct that
        that would have exactly two strings in each sub-array
        """
            
        self._internal.query_params = query_params
    
        return self
    
    def headers(self, headers: list[list[str]]) -> typing.Self:    
        self._internal.headers = headers
    
        return self
    


class InfinityOptions(cogbuilder.Builder[dashboardv2beta1.InfinityOptions]):
    _internal: dashboardv2beta1.InfinityOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.InfinityOptions()

    def build(self) -> dashboardv2beta1.InfinityOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def method(self, method: dashboardv2beta1.HttpRequestMethod) -> typing.Self:    
        self._internal.method = method
    
        return self
    
    def url(self, url: str) -> typing.Self:    
        self._internal.url = url
    
        return self
    
    def body(self, body: str) -> typing.Self:    
        self._internal.body = body
    
        return self
    
    def query_params(self, query_params: list[list[str]]) -> typing.Self:    
        """
        These are 2D arrays of strings, each representing a key-value pair
        We are defining them this way because we can't generate a go struct that
        that would have exactly two strings in each sub-array
        """
            
        self._internal.query_params = query_params
    
        return self
    
    def datasource_uid(self, datasource_uid: str) -> typing.Self:    
        self._internal.datasource_uid = datasource_uid
    
        return self
    
    def headers(self, headers: list[list[str]]) -> typing.Self:    
        self._internal.headers = headers
    
        return self
    


class ActionVariable(cogbuilder.Builder[dashboardv2beta1.ActionVariable]):
    _internal: dashboardv2beta1.ActionVariable

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ActionVariable()

    def build(self) -> dashboardv2beta1.ActionVariable:
        """
        Builds the object.
        """
        return self._internal    
    
    def key(self, key: str) -> typing.Self:    
        self._internal.key = key
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    


class LibraryPanel(cogbuilder.Builder[dashboardv2beta1.LibraryPanelKind]):
    _internal: dashboardv2beta1.LibraryPanelKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.LibraryPanelKind()        
        self._internal.kind = "LibraryPanel"

    def build(self) -> dashboardv2beta1.LibraryPanelKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: cogbuilder.Builder[dashboardv2beta1.LibraryPanelKindSpec]) -> typing.Self:    
        spec_resource = spec.build()
        self._internal.spec = spec_resource
    
        return self
    
    def id(self, id_val: float) -> typing.Self:    
        """
        Panel ID for the library panel in the dashboard
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.LibraryPanelKindSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.LibraryPanelKindSpec)
        self._internal.spec.id_val = id_val
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        """
        Title for the library panel in the dashboard
        """
            
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.LibraryPanelKindSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.LibraryPanelKindSpec)
        self._internal.spec.title = title
    
        return self
    
    def library_panel(self, library_panel: cogbuilder.Builder[dashboardv2beta1.LibraryPanelRef]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.LibraryPanelKindSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.LibraryPanelKindSpec)
        library_panel_resource = library_panel.build()
        self._internal.spec.library_panel = library_panel_resource
    
        return self
    


class LibraryPanelKindSpec(cogbuilder.Builder[dashboardv2beta1.LibraryPanelKindSpec]):
    _internal: dashboardv2beta1.LibraryPanelKindSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.LibraryPanelKindSpec()

    def build(self) -> dashboardv2beta1.LibraryPanelKindSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def id(self, id_val: float) -> typing.Self:    
        """
        Panel ID for the library panel in the dashboard
        """
            
        self._internal.id_val = id_val
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        """
        Title for the library panel in the dashboard
        """
            
        self._internal.title = title
    
        return self
    
    def library_panel(self, library_panel: cogbuilder.Builder[dashboardv2beta1.LibraryPanelRef]) -> typing.Self:    
        library_panel_resource = library_panel.build()
        self._internal.library_panel = library_panel_resource
    
        return self
    


class LibraryPanelRef(cogbuilder.Builder[dashboardv2beta1.LibraryPanelRef]):    
    """
    A library panel is a reusable panel that you can use in any dashboard.
    When you make a change to a library panel, that change propagates to all instances of where the panel is used.
    Library panels streamline reuse of panels across multiple dashboards.
    """
    
    _internal: dashboardv2beta1.LibraryPanelRef

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.LibraryPanelRef()

    def build(self) -> dashboardv2beta1.LibraryPanelRef:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        """
        Library panel name
        """
            
        self._internal.name = name
    
        return self
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Library panel uid
        """
            
        self._internal.uid = uid
    
        return self
    


class GridLayout(cogbuilder.Builder[dashboardv2beta1.GridLayoutKind]):
    _internal: dashboardv2beta1.GridLayoutKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.GridLayoutKind()        
        self._internal.kind = "GridLayout"

    def build(self) -> dashboardv2beta1.GridLayoutKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def items(self, items: list[cogbuilder.Builder[dashboardv2beta1.GridLayoutItemKind]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutSpec)
        items_resources = [r1.build() for r1 in items]
        self._internal.spec.items = items_resources
    
        return self
    
    def item(self, item: cogbuilder.Builder[dashboardv2beta1.GridLayoutItemKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutSpec)
        if self._internal.spec.items is None:
            self._internal.spec.items = []
        
        item_resource = item.build()
        self._internal.spec.items.append(item_resource)
    
        return self
    


class GridLayoutItem(cogbuilder.Builder[dashboardv2beta1.GridLayoutItemKind]):
    _internal: dashboardv2beta1.GridLayoutItemKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.GridLayoutItemKind()        
        self._internal.kind = "GridLayoutItem"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutItemSpec)
        if self._internal.spec.element is None:
            self._internal.spec.element = dashboardv2beta1.ElementReference()
        assert isinstance(self._internal.spec.element, dashboardv2beta1.ElementReference)
        self._internal.spec.element.kind = "ElementReference"        
        self._internal.spec.element.name = name

    def build(self) -> dashboardv2beta1.GridLayoutItemKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def x(self, x: int) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutItemSpec)
        self._internal.spec.x = x
    
        return self
    
    def y(self, y: int) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutItemSpec)
        self._internal.spec.y = y
    
        return self
    
    def width(self, width: int) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutItemSpec)
        self._internal.spec.width = width
    
        return self
    
    def height(self, height: int) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutItemSpec)
        self._internal.spec.height = height
    
        return self
    
    def repeat(self, repeat: cogbuilder.Builder[dashboardv2beta1.RepeatOptions]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutItemSpec)
        repeat_resource = repeat.build()
        self._internal.spec.repeat = repeat_resource
    
        return self
    
    def element(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GridLayoutItemSpec)
        if self._internal.spec.element is None:
            self._internal.spec.element = dashboardv2beta1.ElementReference()
        assert isinstance(self._internal.spec.element, dashboardv2beta1.ElementReference)
        self._internal.spec.element.name = name
    
        return self
    


class ElementReference(cogbuilder.Builder[dashboardv2beta1.ElementReference]):
    _internal: dashboardv2beta1.ElementReference

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ElementReference()        
        self._internal.kind = "ElementReference"

    def build(self) -> dashboardv2beta1.ElementReference:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    


class RepeatOptions(cogbuilder.Builder[dashboardv2beta1.RepeatOptions]):
    _internal: dashboardv2beta1.RepeatOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.RepeatOptions()

    def build(self) -> dashboardv2beta1.RepeatOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    
    def direction(self, direction: typing.Literal["h", "v"]) -> typing.Self:    
        self._internal.direction = direction
    
        return self
    
    def max_per_row(self, max_per_row: int) -> typing.Self:    
        self._internal.max_per_row = max_per_row
    
        return self
    


class RowsLayout(cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind]):
    _internal: dashboardv2beta1.RowsLayoutKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.RowsLayoutKind()        
        self._internal.kind = "RowsLayout"

    def build(self) -> dashboardv2beta1.RowsLayoutKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def rows(self, rows: list[cogbuilder.Builder[dashboardv2beta1.RowsLayoutRowKind]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutSpec)
        rows_resources = [r1.build() for r1 in rows]
        self._internal.spec.rows = rows_resources
    
        return self
    
    def row(self, row: cogbuilder.Builder[dashboardv2beta1.RowsLayoutRowKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutSpec)
        if self._internal.spec.rows is None:
            self._internal.spec.rows = []
        
        row_resource = row.build()
        self._internal.spec.rows.append(row_resource)
    
        return self
    


class RowsLayoutSpec(cogbuilder.Builder[dashboardv2beta1.RowsLayoutSpec]):
    _internal: dashboardv2beta1.RowsLayoutSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.RowsLayoutSpec()

    def build(self) -> dashboardv2beta1.RowsLayoutSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def rows(self, rows: list[cogbuilder.Builder[dashboardv2beta1.RowsLayoutRowKind]]) -> typing.Self:    
        rows_resources = [r1.build() for r1 in rows]
        self._internal.rows = rows_resources
    
        return self
    


class RowsLayoutRow(cogbuilder.Builder[dashboardv2beta1.RowsLayoutRowKind]):
    _internal: dashboardv2beta1.RowsLayoutRowKind

    def __init__(self, title: str) -> None:
        self._internal = dashboardv2beta1.RowsLayoutRowKind()        
        self._internal.kind = "RowsLayoutRow"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        self._internal.spec.title = title

    def build(self) -> dashboardv2beta1.RowsLayoutRowKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        self._internal.spec.title = title
    
        return self
    
    def collapse(self, collapse: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        self._internal.spec.collapse = collapse
    
        return self
    
    def hide_header(self, hide_header: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        self._internal.spec.hide_header = hide_header
    
        return self
    
    def fill_screen(self, fill_screen: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        self._internal.spec.fill_screen = fill_screen
    
        return self
    
    def conditional_rendering(self, conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        conditional_rendering_resource = conditional_rendering.build()
        self._internal.spec.conditional_rendering = conditional_rendering_resource
    
        return self
    
    def repeat(self, repeat: cogbuilder.Builder[dashboardv2beta1.RowRepeatOptions]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        repeat_resource = repeat.build()
        self._internal.spec.repeat = repeat_resource
    
        return self
    
    def grid_layout(self, grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.GridLayoutKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        grid_layout_kind_resource = grid_layout_kind.build()
        self._internal.spec.layout = grid_layout_kind_resource
    
        return self
    
    def auto_grid_layout(self, auto_grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        auto_grid_layout_kind_resource = auto_grid_layout_kind.build()
        self._internal.spec.layout = auto_grid_layout_kind_resource
    
        return self
    
    def tabs_layout(self, tabs_layout_kind: cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        tabs_layout_kind_resource = tabs_layout_kind.build()
        self._internal.spec.layout = tabs_layout_kind_resource
    
        return self
    
    def rows_layout(self, rows_layout_kind: cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.RowsLayoutRowSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.RowsLayoutRowSpec)
        rows_layout_kind_resource = rows_layout_kind.build()
        self._internal.spec.layout = rows_layout_kind_resource
    
        return self
    


class RowsLayoutRowSpec(cogbuilder.Builder[dashboardv2beta1.RowsLayoutRowSpec]):
    _internal: dashboardv2beta1.RowsLayoutRowSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.RowsLayoutRowSpec()

    def build(self) -> dashboardv2beta1.RowsLayoutRowSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def collapse(self, collapse: bool) -> typing.Self:    
        self._internal.collapse = collapse
    
        return self
    
    def hide_header(self, hide_header: bool) -> typing.Self:    
        self._internal.hide_header = hide_header
    
        return self
    
    def fill_screen(self, fill_screen: bool) -> typing.Self:    
        self._internal.fill_screen = fill_screen
    
        return self
    
    def conditional_rendering(self, conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self:    
        conditional_rendering_resource = conditional_rendering.build()
        self._internal.conditional_rendering = conditional_rendering_resource
    
        return self
    
    def repeat(self, repeat: cogbuilder.Builder[dashboardv2beta1.RowRepeatOptions]) -> typing.Self:    
        repeat_resource = repeat.build()
        self._internal.repeat = repeat_resource
    
        return self
    
    def layout(self, layout: typing.Union[cogbuilder.Builder[dashboardv2beta1.GridLayoutKind], cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind], cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind], cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind]]) -> typing.Self:    
        layout_resource = layout.build()
        self._internal.layout = layout_resource
    
        return self
    


class ConditionalRenderingGroup(cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]):
    _internal: dashboardv2beta1.ConditionalRenderingGroupKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ConditionalRenderingGroupKind()        
        self._internal.kind = "ConditionalRenderingGroup"

    def build(self) -> dashboardv2beta1.ConditionalRenderingGroupKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupSpec]) -> typing.Self:    
        spec_resource = spec.build()
        self._internal.spec = spec_resource
    
        return self
    
    def visibility(self, visibility: typing.Literal["show", "hide"]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConditionalRenderingGroupSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConditionalRenderingGroupSpec)
        self._internal.spec.visibility = visibility
    
        return self
    
    def condition(self, condition: typing.Literal["and", "or"]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConditionalRenderingGroupSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConditionalRenderingGroupSpec)
        self._internal.spec.condition = condition
    
        return self
    
    def items(self, items: list[typing.Union[cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingVariableKind], cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingDataKind], cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind]]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConditionalRenderingGroupSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConditionalRenderingGroupSpec)
        items_resources = [r1.build() for r1 in items]
        self._internal.spec.items = items_resources
    
        return self
    


class ConditionalRenderingGroupSpec(cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupSpec]):
    _internal: dashboardv2beta1.ConditionalRenderingGroupSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ConditionalRenderingGroupSpec()

    def build(self) -> dashboardv2beta1.ConditionalRenderingGroupSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def visibility(self, visibility: typing.Literal["show", "hide"]) -> typing.Self:    
        self._internal.visibility = visibility
    
        return self
    
    def condition(self, condition: typing.Literal["and", "or"]) -> typing.Self:    
        self._internal.condition = condition
    
        return self
    
    def items(self, items: list[typing.Union[cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingVariableKind], cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingDataKind], cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind]]]) -> typing.Self:    
        items_resources = [r1.build() for r1 in items]
        self._internal.items = items_resources
    
        return self
    


class ConditionalRenderingVariable(cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingVariableKind]):
    _internal: dashboardv2beta1.ConditionalRenderingVariableKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ConditionalRenderingVariableKind()        
        self._internal.kind = "ConditionalRenderingVariable"

    def build(self) -> dashboardv2beta1.ConditionalRenderingVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingVariableSpec]) -> typing.Self:    
        spec_resource = spec.build()
        self._internal.spec = spec_resource
    
        return self
    
    def variable(self, variable: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConditionalRenderingVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConditionalRenderingVariableSpec)
        self._internal.spec.variable = variable
    
        return self
    
    def operator(self, operator: typing.Literal["equals", "notEquals", "matches", "notMatches"]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConditionalRenderingVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConditionalRenderingVariableSpec)
        self._internal.spec.operator = operator
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConditionalRenderingVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConditionalRenderingVariableSpec)
        self._internal.spec.value = value
    
        return self
    


class ConditionalRenderingVariableSpec(cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingVariableSpec]):
    _internal: dashboardv2beta1.ConditionalRenderingVariableSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ConditionalRenderingVariableSpec()

    def build(self) -> dashboardv2beta1.ConditionalRenderingVariableSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def variable(self, variable: str) -> typing.Self:    
        self._internal.variable = variable
    
        return self
    
    def operator(self, operator: typing.Literal["equals", "notEquals", "matches", "notMatches"]) -> typing.Self:    
        self._internal.operator = operator
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    


class ConditionalRenderingData(cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingDataKind]):
    _internal: dashboardv2beta1.ConditionalRenderingDataKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ConditionalRenderingDataKind()        
        self._internal.kind = "ConditionalRenderingData"

    def build(self) -> dashboardv2beta1.ConditionalRenderingDataKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingDataSpec]) -> typing.Self:    
        spec_resource = spec.build()
        self._internal.spec = spec_resource
    
        return self
    
    def value(self, value: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConditionalRenderingDataSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConditionalRenderingDataSpec)
        self._internal.spec.value = value
    
        return self
    


class ConditionalRenderingDataSpec(cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingDataSpec]):
    _internal: dashboardv2beta1.ConditionalRenderingDataSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ConditionalRenderingDataSpec()

    def build(self) -> dashboardv2beta1.ConditionalRenderingDataSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: bool) -> typing.Self:    
        self._internal.value = value
    
        return self
    


class ConditionalRenderingTimeRangeSize(cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind]):
    _internal: dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind()        
        self._internal.kind = "ConditionalRenderingTimeRangeSize"

    def build(self) -> dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec]) -> typing.Self:    
        spec_resource = spec.build()
        self._internal.spec = spec_resource
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec)
        self._internal.spec.value = value
    
        return self
    


class ConditionalRenderingTimeRangeSizeSpec(cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec]):
    _internal: dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec()

    def build(self) -> dashboardv2beta1.ConditionalRenderingTimeRangeSizeSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    


class RowRepeatOptions(cogbuilder.Builder[dashboardv2beta1.RowRepeatOptions]):
    _internal: dashboardv2beta1.RowRepeatOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.RowRepeatOptions()

    def build(self) -> dashboardv2beta1.RowRepeatOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    


class AutoGridLayout(cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind]):
    _internal: dashboardv2beta1.AutoGridLayoutKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AutoGridLayoutKind()        
        self._internal.kind = "AutoGridLayout"

    def build(self) -> dashboardv2beta1.AutoGridLayoutKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def max_column_count(self, max_column_count: float) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutSpec)
        self._internal.spec.max_column_count = max_column_count
    
        return self
    
    def column_width_mode(self, column_width_mode: typing.Literal["narrow", "standard", "wide", "custom"]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutSpec)
        self._internal.spec.column_width_mode = column_width_mode
    
        return self
    
    def column_width(self, column_width: float) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutSpec)
        self._internal.spec.column_width = column_width
    
        return self
    
    def row_height_mode(self, row_height_mode: typing.Literal["short", "standard", "tall", "custom"]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutSpec)
        self._internal.spec.row_height_mode = row_height_mode
    
        return self
    
    def row_height(self, row_height: float) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutSpec)
        self._internal.spec.row_height = row_height
    
        return self
    
    def fill_screen(self, fill_screen: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutSpec)
        self._internal.spec.fill_screen = fill_screen
    
        return self
    
    def items(self, items: list[cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutItemKind]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutSpec)
        items_resources = [r1.build() for r1 in items]
        self._internal.spec.items = items_resources
    
        return self
    
    def item(self, item: cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutItemKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutSpec)
        if self._internal.spec.items is None:
            self._internal.spec.items = []
        
        item_resource = item.build()
        self._internal.spec.items.append(item_resource)
    
        return self
    


class AutoGridLayoutSpec(cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutSpec]):
    _internal: dashboardv2beta1.AutoGridLayoutSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AutoGridLayoutSpec()

    def build(self) -> dashboardv2beta1.AutoGridLayoutSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def max_column_count(self, max_column_count: float) -> typing.Self:    
        self._internal.max_column_count = max_column_count
    
        return self
    
    def column_width_mode(self, column_width_mode: typing.Literal["narrow", "standard", "wide", "custom"]) -> typing.Self:    
        self._internal.column_width_mode = column_width_mode
    
        return self
    
    def column_width(self, column_width: float) -> typing.Self:    
        self._internal.column_width = column_width
    
        return self
    
    def row_height_mode(self, row_height_mode: typing.Literal["short", "standard", "tall", "custom"]) -> typing.Self:    
        self._internal.row_height_mode = row_height_mode
    
        return self
    
    def row_height(self, row_height: float) -> typing.Self:    
        self._internal.row_height = row_height
    
        return self
    
    def fill_screen(self, fill_screen: bool) -> typing.Self:    
        self._internal.fill_screen = fill_screen
    
        return self
    
    def items(self, items: list[cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutItemKind]]) -> typing.Self:    
        items_resources = [r1.build() for r1 in items]
        self._internal.items = items_resources
    
        return self
    


class AutoGridLayoutItem(cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutItemKind]):
    _internal: dashboardv2beta1.AutoGridLayoutItemKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.AutoGridLayoutItemKind()        
        self._internal.kind = "AutoGridLayoutItem"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutItemSpec)
        if self._internal.spec.element is None:
            self._internal.spec.element = dashboardv2beta1.ElementReference()
        assert isinstance(self._internal.spec.element, dashboardv2beta1.ElementReference)
        self._internal.spec.element.kind = "ElementReference"        
        self._internal.spec.element.name = name

    def build(self) -> dashboardv2beta1.AutoGridLayoutItemKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def element(self, element: cogbuilder.Builder[dashboardv2beta1.ElementReference]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutItemSpec)
        element_resource = element.build()
        self._internal.spec.element = element_resource
    
        return self
    
    def repeat(self, repeat: cogbuilder.Builder[dashboardv2beta1.AutoGridRepeatOptions]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutItemSpec)
        repeat_resource = repeat.build()
        self._internal.spec.repeat = repeat_resource
    
        return self
    
    def conditional_rendering(self, conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutItemSpec)
        conditional_rendering_resource = conditional_rendering.build()
        self._internal.spec.conditional_rendering = conditional_rendering_resource
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AutoGridLayoutItemSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AutoGridLayoutItemSpec)
        if self._internal.spec.element is None:
            self._internal.spec.element = dashboardv2beta1.ElementReference()
        assert isinstance(self._internal.spec.element, dashboardv2beta1.ElementReference)
        self._internal.spec.element.name = name
    
        return self
    


class AutoGridLayoutItemSpec(cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutItemSpec]):
    _internal: dashboardv2beta1.AutoGridLayoutItemSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AutoGridLayoutItemSpec()

    def build(self) -> dashboardv2beta1.AutoGridLayoutItemSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def element(self, element: cogbuilder.Builder[dashboardv2beta1.ElementReference]) -> typing.Self:    
        element_resource = element.build()
        self._internal.element = element_resource
    
        return self
    
    def repeat(self, repeat: cogbuilder.Builder[dashboardv2beta1.AutoGridRepeatOptions]) -> typing.Self:    
        repeat_resource = repeat.build()
        self._internal.repeat = repeat_resource
    
        return self
    
    def conditional_rendering(self, conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self:    
        conditional_rendering_resource = conditional_rendering.build()
        self._internal.conditional_rendering = conditional_rendering_resource
    
        return self
    


class AutoGridRepeatOptions(cogbuilder.Builder[dashboardv2beta1.AutoGridRepeatOptions]):
    _internal: dashboardv2beta1.AutoGridRepeatOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AutoGridRepeatOptions()

    def build(self) -> dashboardv2beta1.AutoGridRepeatOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    


class TabsLayout(cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind]):
    _internal: dashboardv2beta1.TabsLayoutKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.TabsLayoutKind()        
        self._internal.kind = "TabsLayout"

    def build(self) -> dashboardv2beta1.TabsLayoutKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def tabs(self, tabs: list[cogbuilder.Builder[dashboardv2beta1.TabsLayoutTabKind]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutSpec)
        tabs_resources = [r1.build() for r1 in tabs]
        self._internal.spec.tabs = tabs_resources
    
        return self
    
    def tab(self, tab: cogbuilder.Builder[dashboardv2beta1.TabsLayoutTabKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutSpec)
        if self._internal.spec.tabs is None:
            self._internal.spec.tabs = []
        
        tab_resource = tab.build()
        self._internal.spec.tabs.append(tab_resource)
    
        return self
    


class TabsLayoutSpec(cogbuilder.Builder[dashboardv2beta1.TabsLayoutSpec]):
    _internal: dashboardv2beta1.TabsLayoutSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.TabsLayoutSpec()

    def build(self) -> dashboardv2beta1.TabsLayoutSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def tabs(self, tabs: list[cogbuilder.Builder[dashboardv2beta1.TabsLayoutTabKind]]) -> typing.Self:    
        tabs_resources = [r1.build() for r1 in tabs]
        self._internal.tabs = tabs_resources
    
        return self
    


class TabsLayoutTab(cogbuilder.Builder[dashboardv2beta1.TabsLayoutTabKind]):
    _internal: dashboardv2beta1.TabsLayoutTabKind

    def __init__(self, title: str) -> None:
        self._internal = dashboardv2beta1.TabsLayoutTabKind()        
        self._internal.kind = "TabsLayoutTab"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutTabSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutTabSpec)
        self._internal.spec.title = title

    def build(self) -> dashboardv2beta1.TabsLayoutTabKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutTabSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutTabSpec)
        self._internal.spec.title = title
    
        return self
    
    def grid_layout(self, grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.GridLayoutKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutTabSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutTabSpec)
        grid_layout_kind_resource = grid_layout_kind.build()
        self._internal.spec.layout = grid_layout_kind_resource
    
        return self
    
    def rows_layout(self, rows_layout_kind: cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutTabSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutTabSpec)
        rows_layout_kind_resource = rows_layout_kind.build()
        self._internal.spec.layout = rows_layout_kind_resource
    
        return self
    
    def auto_grid_layout(self, auto_grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutTabSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutTabSpec)
        auto_grid_layout_kind_resource = auto_grid_layout_kind.build()
        self._internal.spec.layout = auto_grid_layout_kind_resource
    
        return self
    
    def tabs_layout(self, tabs_layout_kind: cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutTabSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutTabSpec)
        tabs_layout_kind_resource = tabs_layout_kind.build()
        self._internal.spec.layout = tabs_layout_kind_resource
    
        return self
    
    def conditional_rendering(self, conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutTabSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutTabSpec)
        conditional_rendering_resource = conditional_rendering.build()
        self._internal.spec.conditional_rendering = conditional_rendering_resource
    
        return self
    
    def repeat(self, repeat: cogbuilder.Builder[dashboardv2beta1.TabRepeatOptions]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TabsLayoutTabSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TabsLayoutTabSpec)
        repeat_resource = repeat.build()
        self._internal.spec.repeat = repeat_resource
    
        return self
    


class TabsLayoutTabSpec(cogbuilder.Builder[dashboardv2beta1.TabsLayoutTabSpec]):
    _internal: dashboardv2beta1.TabsLayoutTabSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.TabsLayoutTabSpec()

    def build(self) -> dashboardv2beta1.TabsLayoutTabSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    
    def layout(self, layout: typing.Union[cogbuilder.Builder[dashboardv2beta1.GridLayoutKind], cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind], cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind], cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind]]) -> typing.Self:    
        layout_resource = layout.build()
        self._internal.layout = layout_resource
    
        return self
    
    def conditional_rendering(self, conditional_rendering: cogbuilder.Builder[dashboardv2beta1.ConditionalRenderingGroupKind]) -> typing.Self:    
        conditional_rendering_resource = conditional_rendering.build()
        self._internal.conditional_rendering = conditional_rendering_resource
    
        return self
    
    def repeat(self, repeat: cogbuilder.Builder[dashboardv2beta1.TabRepeatOptions]) -> typing.Self:    
        repeat_resource = repeat.build()
        self._internal.repeat = repeat_resource
    
        return self
    


class TabRepeatOptions(cogbuilder.Builder[dashboardv2beta1.TabRepeatOptions]):
    _internal: dashboardv2beta1.TabRepeatOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.TabRepeatOptions()

    def build(self) -> dashboardv2beta1.TabRepeatOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    


class DashboardLink(cogbuilder.Builder[dashboardv2beta1.DashboardLink]):    
    """
    Links with references to other dashboards or external resources
    """
    
    _internal: dashboardv2beta1.DashboardLink

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.DashboardLink()

    def build(self) -> dashboardv2beta1.DashboardLink:
        """
        Builds the object.
        """
        return self._internal    
    
    def title(self, title: str) -> typing.Self:    
        """
        Title to display with the link
        """
            
        self._internal.title = title
    
        return self
    
    def type(self, type_val: dashboardv2beta1.DashboardLinkType) -> typing.Self:    
        """
        Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
        FIXME: The type is generated as `type: DashboardLinkType | dashboardLinkType.Link;` but it should be `type: DashboardLinkType`
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def icon(self, icon: str) -> typing.Self:    
        """
        Icon name to be displayed with the link
        """
            
        self._internal.icon = icon
    
        return self
    
    def tooltip(self, tooltip: str) -> typing.Self:    
        """
        Tooltip to display when the user hovers their mouse over it
        """
            
        self._internal.tooltip = tooltip
    
        return self
    
    def url(self, url: str) -> typing.Self:    
        """
        Link URL. Only required/valid if the type is link
        """
            
        self._internal.url = url
    
        return self
    
    def tags(self, tags: list[str]) -> typing.Self:    
        """
        List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
        """
            
        self._internal.tags = tags
    
        return self
    
    def as_dropdown(self, as_dropdown: bool) -> typing.Self:    
        """
        If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
        """
            
        self._internal.as_dropdown = as_dropdown
    
        return self
    
    def target_blank(self, target_blank: bool) -> typing.Self:    
        """
        If true, the link will be opened in a new tab
        """
            
        self._internal.target_blank = target_blank
    
        return self
    
    def include_vars(self, include_vars: bool) -> typing.Self:    
        """
        If true, includes current template variables values in the link as query params
        """
            
        self._internal.include_vars = include_vars
    
        return self
    
    def keep_time(self, keep_time: bool) -> typing.Self:    
        """
        If true, includes current time range in the link as query params
        """
            
        self._internal.keep_time = keep_time
    
        return self
    
    def placement(self, placement: str) -> typing.Self:    
        """
        Placement can be used to display the link somewhere else on the dashboard other than above the visualisations.
        """
            
        self._internal.placement = placement
    
        return self
    


class TimeSettings(cogbuilder.Builder[dashboardv2beta1.TimeSettingsSpec]):    
    """
    Time configuration
    It defines the default time config for the time picker, the refresh picker for the specific dashboard.
    """
    
    _internal: dashboardv2beta1.TimeSettingsSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.TimeSettingsSpec()

    def build(self) -> dashboardv2beta1.TimeSettingsSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def timezone(self, timezone: str) -> typing.Self:    
        """
        Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
        """
            
        self._internal.timezone = timezone
    
        return self
    
    def from_val(self, from_val: str) -> typing.Self:    
        """
        Start time range for dashboard.
        Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:    
        """
        End time range for dashboard.
        Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
        """
            
        self._internal.to = to
    
        return self
    
    def auto_refresh(self, auto_refresh: str) -> typing.Self:    
        """
        Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
        v1: refresh
        """
            
        self._internal.auto_refresh = auto_refresh
    
        return self
    
    def auto_refresh_intervals(self, auto_refresh_intervals: list[str]) -> typing.Self:    
        """
        Interval options available in the refresh picker dropdown.
        v1: timepicker.refresh_intervals
        """
            
        self._internal.auto_refresh_intervals = auto_refresh_intervals
    
        return self
    
    def quick_ranges(self, quick_ranges: list[cogbuilder.Builder[dashboardv2beta1.TimeRangeOption]]) -> typing.Self:    
        """
        Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
        v1: timepicker.quick_ranges , not exposed in the UI
        """
            
        quick_ranges_resources = [r1.build() for r1 in quick_ranges]
        self._internal.quick_ranges = quick_ranges_resources
    
        return self
    
    def hide_timepicker(self, hide_timepicker: bool) -> typing.Self:    
        """
        Whether timepicker is visible or not.
        v1: timepicker.hidden
        """
            
        self._internal.hide_timepicker = hide_timepicker
    
        return self
    
    def week_start(self, week_start: typing.Literal["saturday", "monday", "sunday"]) -> typing.Self:    
        """
        Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
        """
            
        self._internal.week_start = week_start
    
        return self
    
    def fiscal_year_start_month(self, fiscal_year_start_month: int) -> typing.Self:    
        """
        The month that the fiscal year starts on. 0 = January, 11 = December
        """
            
        self._internal.fiscal_year_start_month = fiscal_year_start_month
    
        return self
    
    def now_delay(self, now_delay: str) -> typing.Self:    
        """
        Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
        v1: timepicker.nowDelay
        """
            
        self._internal.now_delay = now_delay
    
        return self
    


class TimeRangeOption(cogbuilder.Builder[dashboardv2beta1.TimeRangeOption]):
    _internal: dashboardv2beta1.TimeRangeOption

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.TimeRangeOption()

    def build(self) -> dashboardv2beta1.TimeRangeOption:
        """
        Builds the object.
        """
        return self._internal    
    
    def display(self, display: str) -> typing.Self:    
        self._internal.display = display
    
        return self
    
    def from_val(self, from_val: str) -> typing.Self:    
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:    
        self._internal.to = to
    
        return self
    


class QueryVariable(cogbuilder.Builder[dashboardv2beta1.QueryVariableKind]):    
    """
    Query variable kind
    """
    
    _internal: dashboardv2beta1.QueryVariableKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.QueryVariableKind()        
        self._internal.kind = "QueryVariable"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.name = name

    def build(self) -> dashboardv2beta1.QueryVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: dashboardv2beta1.QueryVariableSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.name = name
    
        return self
    
    def current(self, current: dashboardv2beta1.VariableOption) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.current = current
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.hide = hide
    
        return self
    
    def refresh(self, refresh: dashboardv2beta1.VariableRefresh) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.refresh = refresh
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.description = description
    
        return self
    
    def query(self, query: cogbuilder.Builder[dashboardv2beta1.DataQueryKind]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        query_resource = query.build()
        self._internal.spec.query = query_resource
    
        return self
    
    def regex(self, regex: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.regex = regex
    
        return self
    
    def regex_apply_to(self, regex_apply_to: dashboardv2beta1.VariableRegexApplyTo) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.regex_apply_to = regex_apply_to
    
        return self
    
    def sort(self, sort: dashboardv2beta1.VariableSort) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.sort = sort
    
        return self
    
    def definition(self, definition: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.definition = definition
    
        return self
    
    def options(self, options: list[dashboardv2beta1.VariableOption]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.options = options
    
        return self
    
    def multi(self, multi: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.multi = multi
    
        return self
    
    def include_all(self, include_all: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.include_all = include_all
    
        return self
    
    def all_value(self, all_value: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.all_value = all_value
    
        return self
    
    def placeholder(self, placeholder: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.placeholder = placeholder
    
        return self
    
    def allow_custom_value(self, allow_custom_value: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.allow_custom_value = allow_custom_value
    
        return self
    
    def static_options(self, static_options: list[dashboardv2beta1.VariableOption]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.static_options = static_options
    
        return self
    
    def static_options_order(self, static_options_order: typing.Literal["before", "after", "sorted"]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.QueryVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.QueryVariableSpec)
        self._internal.spec.static_options_order = static_options_order
    
        return self
    


class TextVariable(cogbuilder.Builder[dashboardv2beta1.TextVariableKind]):    
    """
    Text variable kind
    """
    
    _internal: dashboardv2beta1.TextVariableKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.TextVariableKind()        
        self._internal.kind = "TextVariable"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TextVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TextVariableSpec)
        self._internal.spec.name = name

    def build(self) -> dashboardv2beta1.TextVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: dashboardv2beta1.TextVariableSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TextVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TextVariableSpec)
        self._internal.spec.name = name
    
        return self
    
    def current(self, current: dashboardv2beta1.VariableOption) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TextVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TextVariableSpec)
        self._internal.spec.current = current
    
        return self
    
    def query(self, query: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TextVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TextVariableSpec)
        self._internal.spec.query = query
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TextVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TextVariableSpec)
        self._internal.spec.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TextVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TextVariableSpec)
        self._internal.spec.hide = hide
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TextVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TextVariableSpec)
        self._internal.spec.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.TextVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.TextVariableSpec)
        self._internal.spec.description = description
    
        return self
    


class ConstantVariable(cogbuilder.Builder[dashboardv2beta1.ConstantVariableKind]):    
    """
    Constant variable kind
    """
    
    _internal: dashboardv2beta1.ConstantVariableKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.ConstantVariableKind()        
        self._internal.kind = "ConstantVariable"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConstantVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConstantVariableSpec)
        self._internal.spec.name = name

    def build(self) -> dashboardv2beta1.ConstantVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: dashboardv2beta1.ConstantVariableSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConstantVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConstantVariableSpec)
        self._internal.spec.name = name
    
        return self
    
    def query(self, query: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConstantVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConstantVariableSpec)
        self._internal.spec.query = query
    
        return self
    
    def current(self, current: dashboardv2beta1.VariableOption) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConstantVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConstantVariableSpec)
        self._internal.spec.current = current
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConstantVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConstantVariableSpec)
        self._internal.spec.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConstantVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConstantVariableSpec)
        self._internal.spec.hide = hide
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConstantVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConstantVariableSpec)
        self._internal.spec.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.ConstantVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.ConstantVariableSpec)
        self._internal.spec.description = description
    
        return self
    


class DatasourceVariable(cogbuilder.Builder[dashboardv2beta1.DatasourceVariableKind]):    
    """
    Datasource variable kind
    """
    
    _internal: dashboardv2beta1.DatasourceVariableKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.DatasourceVariableKind()        
        self._internal.kind = "DatasourceVariable"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.name = name

    def build(self) -> dashboardv2beta1.DatasourceVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: dashboardv2beta1.DatasourceVariableSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.name = name
    
        return self
    
    def plugin_id(self, plugin_id: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.plugin_id = plugin_id
    
        return self
    
    def refresh(self, refresh: dashboardv2beta1.VariableRefresh) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.refresh = refresh
    
        return self
    
    def regex(self, regex: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.regex = regex
    
        return self
    
    def current(self, current: dashboardv2beta1.VariableOption) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.current = current
    
        return self
    
    def options(self, options: list[dashboardv2beta1.VariableOption]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.options = options
    
        return self
    
    def multi(self, multi: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.multi = multi
    
        return self
    
    def include_all(self, include_all: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.include_all = include_all
    
        return self
    
    def all_value(self, all_value: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.all_value = all_value
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.hide = hide
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.description = description
    
        return self
    
    def allow_custom_value(self, allow_custom_value: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.DatasourceVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.DatasourceVariableSpec)
        self._internal.spec.allow_custom_value = allow_custom_value
    
        return self
    


class IntervalVariable(cogbuilder.Builder[dashboardv2beta1.IntervalVariableKind]):    
    """
    Interval variable kind
    """
    
    _internal: dashboardv2beta1.IntervalVariableKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.IntervalVariableKind()        
        self._internal.kind = "IntervalVariable"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.refresh = "onTimeRangeChanged"        
        self._internal.spec.name = name

    def build(self) -> dashboardv2beta1.IntervalVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: dashboardv2beta1.IntervalVariableSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.name = name
    
        return self
    
    def query(self, query: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.query = query
    
        return self
    
    def current(self, current: dashboardv2beta1.VariableOption) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.current = current
    
        return self
    
    def options(self, options: list[dashboardv2beta1.VariableOption]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.options = options
    
        return self
    
    def auto(self, auto: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.auto = auto
    
        return self
    
    def auto_min(self, auto_min: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.auto_min = auto_min
    
        return self
    
    def auto_count(self, auto_count: int) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.auto_count = auto_count
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.hide = hide
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.IntervalVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.IntervalVariableSpec)
        self._internal.spec.description = description
    
        return self
    


class CustomVariable(cogbuilder.Builder[dashboardv2beta1.CustomVariableKind]):    
    """
    Custom variable kind
    """
    
    _internal: dashboardv2beta1.CustomVariableKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.CustomVariableKind()        
        self._internal.kind = "CustomVariable"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.name = name

    def build(self) -> dashboardv2beta1.CustomVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: dashboardv2beta1.CustomVariableSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.name = name
    
        return self
    
    def query(self, query: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.query = query
    
        return self
    
    def current(self, current: dashboardv2beta1.VariableOption) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.current = current
    
        return self
    
    def options(self, options: list[dashboardv2beta1.VariableOption]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.options = options
    
        return self
    
    def multi(self, multi: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.multi = multi
    
        return self
    
    def include_all(self, include_all: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.include_all = include_all
    
        return self
    
    def all_value(self, all_value: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.all_value = all_value
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.hide = hide
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.description = description
    
        return self
    
    def allow_custom_value(self, allow_custom_value: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.allow_custom_value = allow_custom_value
    
        return self
    
    def values_format(self, values_format: typing.Literal["csv", "json"]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.CustomVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.CustomVariableSpec)
        self._internal.spec.values_format = values_format
    
        return self
    


class GroupByVariable(cogbuilder.Builder[dashboardv2beta1.GroupByVariableKind]):    
    """
    Group variable kind
    """
    
    _internal: dashboardv2beta1.GroupByVariableKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.GroupByVariableKind()        
        self._internal.kind = "GroupByVariable"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.name = name

    def build(self) -> dashboardv2beta1.GroupByVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def group(self, group: str) -> typing.Self:    
        self._internal.group = group
    
        return self
    
    def datasource(self, datasource: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource]) -> typing.Self:    
        datasource_resource = datasource.build()
        self._internal.datasource = datasource_resource
    
        return self
    
    def spec(self, spec: dashboardv2beta1.GroupByVariableSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.name = name
    
        return self
    
    def default_value(self, default_value: dashboardv2beta1.VariableOption) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.default_value = default_value
    
        return self
    
    def current(self, current: dashboardv2beta1.VariableOption) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.current = current
    
        return self
    
    def options(self, options: list[dashboardv2beta1.VariableOption]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.options = options
    
        return self
    
    def multi(self, multi: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.multi = multi
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.hide = hide
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.GroupByVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.GroupByVariableSpec)
        self._internal.spec.description = description
    
        return self
    


class AdhocVariable(cogbuilder.Builder[dashboardv2beta1.AdhocVariableKind]):    
    """
    Adhoc variable kind
    """
    
    _internal: dashboardv2beta1.AdhocVariableKind

    def __init__(self, name: str) -> None:
        self._internal = dashboardv2beta1.AdhocVariableKind()        
        self._internal.kind = "AdhocVariable"        
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        self._internal.spec.name = name

    def build(self) -> dashboardv2beta1.AdhocVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def group(self, group: str) -> typing.Self:    
        self._internal.group = group
    
        return self
    
    def datasource(self, datasource: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1AdhocVariableKindDatasource]) -> typing.Self:    
        datasource_resource = datasource.build()
        self._internal.datasource = datasource_resource
    
        return self
    
    def spec(self, spec: dashboardv2beta1.AdhocVariableSpec) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        self._internal.spec.name = name
    
        return self
    
    def base_filters(self, base_filters: list[cogbuilder.Builder[dashboardv2beta1.AdHocFilterWithLabels]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        base_filters_resources = [r1.build() for r1 in base_filters]
        self._internal.spec.base_filters = base_filters_resources
    
        return self
    
    def filters(self, filters: list[cogbuilder.Builder[dashboardv2beta1.AdHocFilterWithLabels]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        filters_resources = [r1.build() for r1 in filters]
        self._internal.spec.filters = filters_resources
    
        return self
    
    def default_keys(self, default_keys: list[cogbuilder.Builder[dashboardv2beta1.MetricFindValue]]) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        default_keys_resources = [r1.build() for r1 in default_keys]
        self._internal.spec.default_keys = default_keys_resources
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        self._internal.spec.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        self._internal.spec.hide = hide
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        self._internal.spec.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        self._internal.spec.description = description
    
        return self
    
    def allow_custom_value(self, allow_custom_value: bool) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = dashboardv2beta1.AdhocVariableSpec()
        assert isinstance(self._internal.spec, dashboardv2beta1.AdhocVariableSpec)
        self._internal.spec.allow_custom_value = allow_custom_value
    
        return self
    


class AdHocFilterWithLabels(cogbuilder.Builder[dashboardv2beta1.AdHocFilterWithLabels]):    
    """
    Define the AdHocFilterWithLabels type
    """
    
    _internal: dashboardv2beta1.AdHocFilterWithLabels

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.AdHocFilterWithLabels()

    def build(self) -> dashboardv2beta1.AdHocFilterWithLabels:
        """
        Builds the object.
        """
        return self._internal    
    
    def key(self, key: str) -> typing.Self:    
        self._internal.key = key
    
        return self
    
    def operator(self, operator: str) -> typing.Self:    
        self._internal.operator = operator
    
        return self
    
    def value(self, value: str) -> typing.Self:    
        self._internal.value = value
    
        return self
    
    def values(self, values: list[str]) -> typing.Self:    
        self._internal.values = values
    
        return self
    
    def key_label(self, key_label: str) -> typing.Self:    
        self._internal.key_label = key_label
    
        return self
    
    def value_labels(self, value_labels: list[str]) -> typing.Self:    
        self._internal.value_labels = value_labels
    
        return self
    
    def force_edit(self, force_edit: bool) -> typing.Self:    
        self._internal.force_edit = force_edit
    
        return self
    
    def origin(self, origin: str) -> typing.Self:    
        self._internal.origin = origin
    
        return self
    
    def condition(self, condition: str) -> typing.Self:    
        """
        @deprecated
        """
            
        self._internal.condition = condition
    
        return self
    


class MetricFindValue(cogbuilder.Builder[dashboardv2beta1.MetricFindValue]):    
    """
    Define the MetricFindValue type
    """
    
    _internal: dashboardv2beta1.MetricFindValue

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.MetricFindValue()

    def build(self) -> dashboardv2beta1.MetricFindValue:
        """
        Builds the object.
        """
        return self._internal    
    
    def text(self, text: str) -> typing.Self:    
        self._internal.text = text
    
        return self
    
    def value(self, value: typing.Union[str, float]) -> typing.Self:    
        self._internal.value = value
    
        return self
    
    def group(self, group: str) -> typing.Self:    
        self._internal.group = group
    
        return self
    
    def expandable(self, expandable: bool) -> typing.Self:    
        self._internal.expandable = expandable
    
        return self
    


class SwitchVariableKind(cogbuilder.Builder[dashboardv2beta1.SwitchVariableKind]):
    _internal: dashboardv2beta1.SwitchVariableKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.SwitchVariableKind()        
        self._internal.kind = "SwitchVariable"

    def build(self) -> dashboardv2beta1.SwitchVariableKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def spec(self, spec: cogbuilder.Builder[dashboardv2beta1.SwitchVariableSpec]) -> typing.Self:    
        spec_resource = spec.build()
        self._internal.spec = spec_resource
    
        return self
    


class SwitchVariableSpec(cogbuilder.Builder[dashboardv2beta1.SwitchVariableSpec]):
    _internal: dashboardv2beta1.SwitchVariableSpec

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.SwitchVariableSpec()

    def build(self) -> dashboardv2beta1.SwitchVariableSpec:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    
    def current(self, current: str) -> typing.Self:    
        self._internal.current = current
    
        return self
    
    def enabled_value(self, enabled_value: str) -> typing.Self:    
        self._internal.enabled_value = enabled_value
    
        return self
    
    def disabled_value(self, disabled_value: str) -> typing.Self:    
        self._internal.disabled_value = disabled_value
    
        return self
    
    def label(self, label: str) -> typing.Self:    
        self._internal.label = label
    
        return self
    
    def hide(self, hide: dashboardv2beta1.VariableHide) -> typing.Self:    
        self._internal.hide = hide
    
        return self
    
    def skip_url_sync(self, skip_url_sync: bool) -> typing.Self:    
        self._internal.skip_url_sync = skip_url_sync
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        self._internal.description = description
    
        return self
    


class Kind(cogbuilder.Builder[dashboardv2beta1.Kind]):    
    """
    --- Common types ---
    """
    
    _internal: dashboardv2beta1.Kind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Kind()

    def build(self) -> dashboardv2beta1.Kind:
        """
        Builds the object.
        """
        return self._internal    
    
    def kind(self, kind: str) -> typing.Self:    
        self._internal.kind = kind
    
        return self
    
    def spec(self, spec: object) -> typing.Self:    
        self._internal.spec = spec
    
        return self
    
    def metadata(self, metadata: object) -> typing.Self:    
        self._internal.metadata = metadata
    
        return self
    


class CustomFormatterVariable(cogbuilder.Builder[dashboardv2beta1.CustomFormatterVariable]):    
    """
    Custom formatter variable
    """
    
    _internal: dashboardv2beta1.CustomFormatterVariable

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.CustomFormatterVariable()

    def build(self) -> dashboardv2beta1.CustomFormatterVariable:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    
    def type(self, type_val: dashboardv2beta1.VariableType) -> typing.Self:    
        self._internal.type_val = type_val
    
        return self
    
    def multi(self, multi: bool) -> typing.Self:    
        self._internal.multi = multi
    
        return self
    
    def include_all(self, include_all: bool) -> typing.Self:    
        self._internal.include_all = include_all
    
        return self
    


class VariableValueOption(cogbuilder.Builder[dashboardv2beta1.VariableValueOption]):    
    """
    FIXME: should we introduce this? --- Variable value option
    """
    
    _internal: dashboardv2beta1.VariableValueOption

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.VariableValueOption()

    def build(self) -> dashboardv2beta1.VariableValueOption:
        """
        Builds the object.
        """
        return self._internal    
    
    def label(self, label: str) -> typing.Self:    
        self._internal.label = label
    
        return self
    
    def value(self, value: dashboardv2beta1.VariableValueSingle) -> typing.Self:    
        self._internal.value = value
    
        return self
    
    def group(self, group: str) -> typing.Self:    
        self._internal.group = group
    
        return self
    


class Dashboardv2beta1DataQueryKindDatasource(cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]):
    _internal: dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource()

    def build(self) -> dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    


class Dashboardv2beta1FieldConfigSourceOverrides(cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides]):
    _internal: dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides()

    def build(self) -> dashboardv2beta1.Dashboardv2beta1FieldConfigSourceOverrides:
        """
        Builds the object.
        """
        return self._internal    
    
    def system_ref(self, system_ref: str) -> typing.Self:    
        """
        Describes config override rules created when interacting with Grafana.
        """
            
        self._internal.system_ref = system_ref
    
        return self
    
    def matcher(self, matcher: dashboardv2beta1.MatcherConfig) -> typing.Self:    
        self._internal.matcher = matcher
    
        return self
    
    def properties(self, properties: list[dashboardv2beta1.DynamicConfigValue]) -> typing.Self:    
        self._internal.properties = properties
    
        return self
    


class Dashboardv2beta1RangeMapOptions(cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1RangeMapOptions]):
    _internal: dashboardv2beta1.Dashboardv2beta1RangeMapOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Dashboardv2beta1RangeMapOptions()

    def build(self) -> dashboardv2beta1.Dashboardv2beta1RangeMapOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: float) -> typing.Self:    
        """
        Min value of the range. It can be null which means -Infinity
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: float) -> typing.Self:    
        """
        Max value of the range. It can be null which means +Infinity
        """
            
        self._internal.to = to
    
        return self
    
    def result(self, result: cogbuilder.Builder[dashboardv2beta1.ValueMappingResult]) -> typing.Self:    
        """
        Config to apply when the value is within the range
        """
            
        result_resource = result.build()
        self._internal.result = result_resource
    
        return self
    


class Dashboardv2beta1RegexMapOptions(cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1RegexMapOptions]):
    _internal: dashboardv2beta1.Dashboardv2beta1RegexMapOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Dashboardv2beta1RegexMapOptions()

    def build(self) -> dashboardv2beta1.Dashboardv2beta1RegexMapOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def pattern(self, pattern: str) -> typing.Self:    
        """
        Regular expression to match against
        """
            
        self._internal.pattern = pattern
    
        return self
    
    def result(self, result: cogbuilder.Builder[dashboardv2beta1.ValueMappingResult]) -> typing.Self:    
        """
        Config to apply when the value matches the regex
        """
            
        result_resource = result.build()
        self._internal.result = result_resource
    
        return self
    


class Dashboardv2beta1SpecialValueMapOptions(cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1SpecialValueMapOptions]):
    _internal: dashboardv2beta1.Dashboardv2beta1SpecialValueMapOptions

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Dashboardv2beta1SpecialValueMapOptions()

    def build(self) -> dashboardv2beta1.Dashboardv2beta1SpecialValueMapOptions:
        """
        Builds the object.
        """
        return self._internal    
    
    def match(self, match: dashboardv2beta1.SpecialValueMatch) -> typing.Self:    
        """
        Special value to match against
        """
            
        self._internal.match = match
    
        return self
    
    def result(self, result: cogbuilder.Builder[dashboardv2beta1.ValueMappingResult]) -> typing.Self:    
        """
        Config to apply when the value matches the special value
        """
            
        result_resource = result.build()
        self._internal.result = result_resource
    
        return self
    


class Dashboardv2beta1ActionStyle(cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1ActionStyle]):
    _internal: dashboardv2beta1.Dashboardv2beta1ActionStyle

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Dashboardv2beta1ActionStyle()

    def build(self) -> dashboardv2beta1.Dashboardv2beta1ActionStyle:
        """
        Builds the object.
        """
        return self._internal    
    
    def background_color(self, background_color: str) -> typing.Self:    
        self._internal.background_color = background_color
    
        return self
    


class Dashboardv2beta1GroupByVariableKindDatasource(cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource]):
    _internal: dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource()

    def build(self) -> dashboardv2beta1.Dashboardv2beta1GroupByVariableKindDatasource:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    


class Dashboardv2beta1AdhocVariableKindDatasource(cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1AdhocVariableKindDatasource]):
    _internal: dashboardv2beta1.Dashboardv2beta1AdhocVariableKindDatasource

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.Dashboardv2beta1AdhocVariableKindDatasource()

    def build(self) -> dashboardv2beta1.Dashboardv2beta1AdhocVariableKindDatasource:
        """
        Builds the object.
        """
        return self._internal    
    
    def name(self, name: str) -> typing.Self:    
        self._internal.name = name
    
        return self
    

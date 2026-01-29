# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import runtime as cogruntime
import enum


class Dashboard:
    annotations: list['AnnotationQueryKind']
    # Configuration of dashboard cursor sync behavior.
    # "Off" for no shared crosshair or tooltip (default).
    # "Crosshair" for shared crosshair.
    # "Tooltip" for shared crosshair AND shared tooltip.
    cursor_sync: 'DashboardCursorSync'
    # Description of dashboard.
    description: typing.Optional[str]
    # Whether a dashboard is editable or not.
    editable: typing.Optional[bool]
    elements: dict[str, 'Element']
    layout: typing.Union['GridLayoutKind', 'RowsLayoutKind', 'AutoGridLayoutKind', 'TabsLayoutKind']
    # Links with references to other dashboards or external websites.
    links: list['DashboardLink']
    # When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    # This will keep data "moving left" regardless of the query refresh rate. This setting helps
    # avoid dashboards presenting stale live data.
    live_now: typing.Optional[bool]
    # When set to true, the dashboard will load all panels in the dashboard when it's loaded.
    preload: bool
    # Plugins only. The version of the dashboard installed together with the plugin.
    # This is used to determine if the dashboard should be updated when the plugin is updated.
    revision: typing.Optional[int]
    # Tags associated with dashboard.
    tags: list[str]
    time_settings: 'TimeSettingsSpec'
    # Title of dashboard.
    title: str
    # Configured template variables.
    variables: list['VariableKind']

    def __init__(self, annotations: typing.Optional[list['AnnotationQueryKind']] = None, cursor_sync: typing.Optional['DashboardCursorSync'] = None, description: typing.Optional[str] = None, editable: typing.Optional[bool] = True, elements: typing.Optional[dict[str, 'Element']] = None, layout: typing.Optional[typing.Union['GridLayoutKind', 'RowsLayoutKind', 'AutoGridLayoutKind', 'TabsLayoutKind']] = None, links: typing.Optional[list['DashboardLink']] = None, live_now: typing.Optional[bool] = None, preload: bool = False, revision: typing.Optional[int] = None, tags: typing.Optional[list[str]] = None, time_settings: typing.Optional['TimeSettingsSpec'] = None, title: str = "", variables: typing.Optional[list['VariableKind']] = None) -> None:
        self.annotations = annotations if annotations is not None else []
        self.cursor_sync = cursor_sync if cursor_sync is not None else DashboardCursorSync.OFF
        self.description = description
        self.editable = editable
        self.elements = elements if elements is not None else {}
        self.layout = layout if layout is not None else GridLayoutKind()
        self.links = links if links is not None else []
        self.live_now = live_now
        self.preload = preload
        self.revision = revision
        self.tags = tags if tags is not None else []
        self.time_settings = time_settings if time_settings is not None else TimeSettingsSpec()
        self.title = title
        self.variables = variables if variables is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "annotations": self.annotations,
            "cursorSync": self.cursor_sync,
            "elements": self.elements,
            "layout": self.layout,
            "links": self.links,
            "preload": self.preload,
            "tags": self.tags,
            "timeSettings": self.time_settings,
            "title": self.title,
            "variables": self.variables,
        }
        if self.description is not None:
            payload["description"] = self.description
        if self.editable is not None:
            payload["editable"] = self.editable
        if self.live_now is not None:
            payload["liveNow"] = self.live_now
        if self.revision is not None:
            payload["revision"] = self.revision
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "annotations" in data:
            args["annotations"] = [AnnotationQueryKind.from_json(item) for item in data["annotations"]]
        if "cursorSync" in data:
            args["cursor_sync"] = data["cursorSync"]
        if "description" in data:
            args["description"] = data["description"]
        if "editable" in data:
            args["editable"] = data["editable"]
        if "elements" in data:
            decoding_map_elements_map_ref_union: dict[str, typing.Union[typing.Type[LibraryPanelKind], typing.Type[PanelKind]]] = {"LibraryPanel": LibraryPanelKind, "Panel": PanelKind}
            args["elements"] = {key: decoding_map_elements_map_ref_union[data["elements"][key]["kind"]].from_json(data["elements"][key]) for key in data["elements"].keys()}
        if "layout" in data:
            decoding_map_layout_union: dict[str, typing.Union[typing.Type[AutoGridLayoutKind], typing.Type[GridLayoutKind], typing.Type[RowsLayoutKind], typing.Type[TabsLayoutKind]]] = {"AutoGridLayout": AutoGridLayoutKind, "GridLayout": GridLayoutKind, "RowsLayout": RowsLayoutKind, "TabsLayout": TabsLayoutKind}
            args["layout"] = decoding_map_layout_union[data["layout"]["kind"]].from_json(data["layout"])
        if "links" in data:
            args["links"] = [DashboardLink.from_json(item) for item in data["links"]]
        if "liveNow" in data:
            args["live_now"] = data["liveNow"]
        if "preload" in data:
            args["preload"] = data["preload"]
        if "revision" in data:
            args["revision"] = data["revision"]
        if "tags" in data:
            args["tags"] = data["tags"]
        if "timeSettings" in data:
            args["time_settings"] = TimeSettingsSpec.from_json(data["timeSettings"])
        if "title" in data:
            args["title"] = data["title"]
        if "variables" in data:
            decoding_map_variables_array_ref_union: dict[str, typing.Union[typing.Type[AdhocVariableKind], typing.Type[ConstantVariableKind], typing.Type[CustomVariableKind], typing.Type[DatasourceVariableKind], typing.Type[GroupByVariableKind], typing.Type[IntervalVariableKind], typing.Type[QueryVariableKind], typing.Type[SwitchVariableKind], typing.Type[TextVariableKind]]] = {"AdhocVariable": AdhocVariableKind, "ConstantVariable": ConstantVariableKind, "CustomVariable": CustomVariableKind, "DatasourceVariable": DatasourceVariableKind, "GroupByVariable": GroupByVariableKind, "IntervalVariable": IntervalVariableKind, "QueryVariable": QueryVariableKind, "SwitchVariable": SwitchVariableKind, "TextVariable": TextVariableKind}
            args["variables"] = [decoding_map_variables_array_ref_union[item["kind"]].from_json(item) for item in data["variables"]]        

        return cls(**args)


class AnnotationQueryKind:
    kind: typing.Literal["AnnotationQuery"]
    spec: 'AnnotationQuerySpec'

    def __init__(self, spec: typing.Optional['AnnotationQuerySpec'] = None) -> None:
        self.kind = "AnnotationQuery"
        self.spec = spec if spec is not None else AnnotationQuerySpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = AnnotationQuerySpec.from_json(data["spec"])        

        return cls(**args)


class AnnotationQuerySpec:
    query: 'DataQueryKind'
    enable: bool
    hide: bool
    icon_color: str
    name: str
    built_in: typing.Optional[bool]
    filter_val: typing.Optional['AnnotationPanelFilter']
    # Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
    placement: str
    # Mappings define how to convert data frame fields to annotation event fields.
    mappings: typing.Optional[dict[str, 'AnnotationEventFieldMapping']]
    # Catch-all field for datasource-specific properties. Should not be available in as code tooling.
    legacy_options: typing.Optional[dict[str, object]]

    def __init__(self, query: typing.Optional['DataQueryKind'] = None, enable: bool = False, hide: bool = False, icon_color: str = "", name: str = "", built_in: typing.Optional[bool] = False, filter_val: typing.Optional['AnnotationPanelFilter'] = None, mappings: typing.Optional[dict[str, 'AnnotationEventFieldMapping']] = None, legacy_options: typing.Optional[dict[str, object]] = None) -> None:
        self.query = query if query is not None else DataQueryKind()
        self.enable = enable
        self.hide = hide
        self.icon_color = icon_color
        self.name = name
        self.built_in = built_in
        self.filter_val = filter_val
        self.placement = AnnotationQueryPlacement
        self.mappings = mappings
        self.legacy_options = legacy_options

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "query": self.query,
            "enable": self.enable,
            "hide": self.hide,
            "iconColor": self.icon_color,
            "name": self.name,
        }
        if self.built_in is not None:
            payload["builtIn"] = self.built_in
        if self.filter_val is not None:
            payload["filter"] = self.filter_val
        if self.placement is not None:
            payload["placement"] = self.placement
        if self.mappings is not None:
            payload["mappings"] = self.mappings
        if self.legacy_options is not None:
            payload["legacyOptions"] = self.legacy_options
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "query" in data:
            args["query"] = DataQueryKind.from_json(data["query"])
        if "enable" in data:
            args["enable"] = data["enable"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "iconColor" in data:
            args["icon_color"] = data["iconColor"]
        if "name" in data:
            args["name"] = data["name"]
        if "builtIn" in data:
            args["built_in"] = data["builtIn"]
        if "filter" in data:
            args["filter_val"] = AnnotationPanelFilter.from_json(data["filter"])
        if "mappings" in data:
            args["mappings"] = {key: AnnotationEventFieldMapping.from_json(data["mappings"][key]) for key in data["mappings"].keys()}
        if "legacyOptions" in data:
            args["legacy_options"] = data["legacyOptions"]        

        return cls(**args)


class DataQueryKind:
    kind: typing.Literal["DataQuery"]
    group: str
    version: str
    # New type for datasource reference
    # Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
    datasource: typing.Optional['Dashboardv2beta1DataQueryKindDatasource']
    spec: typing.Optional[object]

    def __init__(self, group: str = "", version: str = "v0", datasource: typing.Optional['Dashboardv2beta1DataQueryKindDatasource'] = None, spec: typing.Optional[object] = None) -> None:
        self.kind = "DataQuery"
        self.group = group
        self.version = version
        self.datasource = datasource
        self.spec = spec

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "group": self.group,
            "version": self.version,
            "spec": self.spec,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        return payload

    
    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
    
            
        if "group" in data:
            args["group"] = data["group"]
    
        if "version" in data:
            args["version"] = data["version"]
    
        if "datasource" in data:
            args["datasource"] = data["datasource"]
    
        if "spec" in data:
            args["spec"] = cogruntime.dataquery_from_json(data["spec"],  data["group"])
        return cls(**args)


class AnnotationPanelFilter:
    # Should the specified panels be included or excluded
    exclude: typing.Optional[bool]
    # Panel IDs that should be included or excluded
    ids: list[int]

    def __init__(self, exclude: typing.Optional[bool] = False, ids: typing.Optional[list[int]] = None) -> None:
        self.exclude = exclude
        self.ids = ids if ids is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "ids": self.ids,
        }
        if self.exclude is not None:
            payload["exclude"] = self.exclude
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "exclude" in data:
            args["exclude"] = data["exclude"]
        if "ids" in data:
            args["ids"] = data["ids"]        

        return cls(**args)


# Annotation Query placement. Defines where the annotation query should be displayed.
# - "inControlsMenu" renders the annotation query in the dashboard controls dropdown menu
AnnotationQueryPlacement: typing.Literal["inControlsMenu"] = "inControlsMenu"


class AnnotationEventFieldMapping:
    """
    Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
    """

    # Source type for the field value
    source: typing.Optional[str]
    # Constant value to use when source is "text"
    value: typing.Optional[str]
    # Regular expression to apply to the field value
    regex: typing.Optional[str]

    def __init__(self, source: typing.Optional[str] = "field", value: typing.Optional[str] = None, regex: typing.Optional[str] = None) -> None:
        self.source = source
        self.value = value
        self.regex = regex

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.source is not None:
            payload["source"] = self.source
        if self.value is not None:
            payload["value"] = self.value
        if self.regex is not None:
            payload["regex"] = self.regex
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "source" in data:
            args["source"] = data["source"]
        if "value" in data:
            args["value"] = data["value"]
        if "regex" in data:
            args["regex"] = data["regex"]        

        return cls(**args)


class DashboardCursorSync(enum.StrEnum):
    """
    "Off" for no shared crosshair or tooltip (default).
    "Crosshair" for shared crosshair.
    "Tooltip" for shared crosshair AND shared tooltip.
    """

    CROSSHAIR = "Crosshair"
    TOOLTIP = "Tooltip"
    OFF = "Off"


# Supported dashboard elements
# |* more element types in the future
Element: typing.TypeAlias = typing.Union['PanelKind', 'LibraryPanelKind']


class PanelKind:
    kind: typing.Literal["Panel"]
    spec: 'PanelSpec'

    def __init__(self, spec: typing.Optional['PanelSpec'] = None) -> None:
        self.kind = "Panel"
        self.spec = spec if spec is not None else PanelSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = PanelSpec.from_json(data["spec"])        

        return cls(**args)


class PanelSpec:
    id_val: float
    title: str
    description: str
    links: list['DataLink']
    data: 'QueryGroupKind'
    viz_config: 'VizConfigKind'
    transparent: typing.Optional[bool]

    def __init__(self, id_val: float = 0, title: str = "", description: str = "", links: typing.Optional[list['DataLink']] = None, data: typing.Optional['QueryGroupKind'] = None, viz_config: typing.Optional['VizConfigKind'] = None, transparent: typing.Optional[bool] = None) -> None:
        self.id_val = id_val
        self.title = title
        self.description = description
        self.links = links if links is not None else []
        self.data = data if data is not None else QueryGroupKind()
        self.viz_config = viz_config if viz_config is not None else VizConfigKind()
        self.transparent = transparent

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "title": self.title,
            "description": self.description,
            "links": self.links,
            "data": self.data,
            "vizConfig": self.viz_config,
        }
        if self.transparent is not None:
            payload["transparent"] = self.transparent
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "title" in data:
            args["title"] = data["title"]
        if "description" in data:
            args["description"] = data["description"]
        if "links" in data:
            args["links"] = [DataLink.from_json(item) for item in data["links"]]
        if "data" in data:
            args["data"] = QueryGroupKind.from_json(data["data"])
        if "vizConfig" in data:
            args["viz_config"] = VizConfigKind.from_json(data["vizConfig"])
        if "transparent" in data:
            args["transparent"] = data["transparent"]        

        return cls(**args)


class DataLink:
    title: str
    url: str
    target_blank: typing.Optional[bool]

    def __init__(self, title: str = "", url: str = "", target_blank: typing.Optional[bool] = None) -> None:
        self.title = title
        self.url = url
        self.target_blank = target_blank

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "title": self.title,
            "url": self.url,
        }
        if self.target_blank is not None:
            payload["targetBlank"] = self.target_blank
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "url" in data:
            args["url"] = data["url"]
        if "targetBlank" in data:
            args["target_blank"] = data["targetBlank"]        

        return cls(**args)


class QueryGroupKind:
    kind: typing.Literal["QueryGroup"]
    spec: 'QueryGroupSpec'

    def __init__(self, spec: typing.Optional['QueryGroupSpec'] = None) -> None:
        self.kind = "QueryGroup"
        self.spec = spec if spec is not None else QueryGroupSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = QueryGroupSpec.from_json(data["spec"])        

        return cls(**args)


class QueryGroupSpec:
    queries: list['PanelQueryKind']
    transformations: list['TransformationKind']
    query_options: 'QueryOptionsSpec'

    def __init__(self, queries: typing.Optional[list['PanelQueryKind']] = None, transformations: typing.Optional[list['TransformationKind']] = None, query_options: typing.Optional['QueryOptionsSpec'] = None) -> None:
        self.queries = queries if queries is not None else []
        self.transformations = transformations if transformations is not None else []
        self.query_options = query_options if query_options is not None else QueryOptionsSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "queries": self.queries,
            "transformations": self.transformations,
            "queryOptions": self.query_options,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "queries" in data:
            args["queries"] = [PanelQueryKind.from_json(item) for item in data["queries"]]
        if "transformations" in data:
            args["transformations"] = [TransformationKind.from_json(item) for item in data["transformations"]]
        if "queryOptions" in data:
            args["query_options"] = QueryOptionsSpec.from_json(data["queryOptions"])        

        return cls(**args)


class PanelQueryKind:
    kind: typing.Literal["PanelQuery"]
    spec: 'PanelQuerySpec'

    def __init__(self, spec: typing.Optional['PanelQuerySpec'] = None) -> None:
        self.kind = "PanelQuery"
        self.spec = spec if spec is not None else PanelQuerySpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = PanelQuerySpec.from_json(data["spec"])        

        return cls(**args)


class PanelQuerySpec:
    query: 'DataQueryKind'
    ref_id: str
    hidden: bool

    def __init__(self, query: typing.Optional['DataQueryKind'] = None, ref_id: str = "A", hidden: bool = False) -> None:
        self.query = query if query is not None else DataQueryKind()
        self.ref_id = ref_id
        self.hidden = hidden

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "query": self.query,
            "refId": self.ref_id,
            "hidden": self.hidden,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "query" in data:
            args["query"] = DataQueryKind.from_json(data["query"])
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hidden" in data:
            args["hidden"] = data["hidden"]        

        return cls(**args)


class TransformationKind:
    # The kind of a TransformationKind is the transformation ID
    kind: str
    spec: 'DataTransformerConfig'

    def __init__(self, kind: str = "", spec: typing.Optional['DataTransformerConfig'] = None) -> None:
        self.kind = kind
        self.spec = spec if spec is not None else DataTransformerConfig()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "kind" in data:
            args["kind"] = data["kind"]
        if "spec" in data:
            args["spec"] = DataTransformerConfig.from_json(data["spec"])        

        return cls(**args)


class DataTransformerConfig:
    """
    Transformations allow to manipulate data returned by a query before the system applies a visualization.
    Using transformations you can: rename fields, join time series data, perform mathematical operations across queries,
    use the output of one transformation as the input to another transformation, etc.
    """

    # Unique identifier of transformer
    id_val: str
    # Disabled transformations are skipped
    disabled: typing.Optional[bool]
    # Optional frame matcher. When missing it will be applied to all results
    filter_val: typing.Optional['MatcherConfig']
    # Where to pull DataFrames from as input to transformation
    topic: typing.Optional['DataTopic']
    # Options to be passed to the transformer
    # Valid options depend on the transformer id
    options: object

    def __init__(self, id_val: str = "", disabled: typing.Optional[bool] = None, filter_val: typing.Optional['MatcherConfig'] = None, topic: typing.Optional['DataTopic'] = None, options: object = None) -> None:
        self.id_val = id_val
        self.disabled = disabled
        self.filter_val = filter_val
        self.topic = topic
        self.options = options

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "options": self.options,
        }
        if self.disabled is not None:
            payload["disabled"] = self.disabled
        if self.filter_val is not None:
            payload["filter"] = self.filter_val
        if self.topic is not None:
            payload["topic"] = self.topic
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "disabled" in data:
            args["disabled"] = data["disabled"]
        if "filter" in data:
            args["filter_val"] = MatcherConfig.from_json(data["filter"])
        if "topic" in data:
            args["topic"] = data["topic"]
        if "options" in data:
            args["options"] = data["options"]        

        return cls(**args)


class MatcherConfig:
    """
    Matcher is a predicate configuration. Based on the config a set of field(s) or values is filtered in order to apply override / transformation.
    It comes with in id ( to resolve implementation from registry) and a configuration that’s specific to a particular matcher type.
    """

    # The matcher id. This is used to find the matcher implementation from registry.
    id_val: str
    # The matcher options. This is specific to the matcher implementation.
    options: typing.Optional[object]

    def __init__(self, id_val: str = "", options: typing.Optional[object] = None) -> None:
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


class DataTopic(enum.StrEnum):
    """
    A topic is attached to DataFrame metadata in query results.
    This specifies where the data should be used.
    """

    SERIES = "series"
    ANNOTATIONS = "annotations"
    ALERT_STATES = "alertStates"


class QueryOptionsSpec:
    time_from: typing.Optional[str]
    max_data_points: typing.Optional[int]
    time_shift: typing.Optional[str]
    query_caching_ttl: typing.Optional[int]
    interval: typing.Optional[str]
    cache_timeout: typing.Optional[str]
    hide_time_override: typing.Optional[bool]
    time_compare: typing.Optional[str]

    def __init__(self, time_from: typing.Optional[str] = None, max_data_points: typing.Optional[int] = None, time_shift: typing.Optional[str] = None, query_caching_ttl: typing.Optional[int] = None, interval: typing.Optional[str] = None, cache_timeout: typing.Optional[str] = None, hide_time_override: typing.Optional[bool] = None, time_compare: typing.Optional[str] = None) -> None:
        self.time_from = time_from
        self.max_data_points = max_data_points
        self.time_shift = time_shift
        self.query_caching_ttl = query_caching_ttl
        self.interval = interval
        self.cache_timeout = cache_timeout
        self.hide_time_override = hide_time_override
        self.time_compare = time_compare

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.time_from is not None:
            payload["timeFrom"] = self.time_from
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.time_shift is not None:
            payload["timeShift"] = self.time_shift
        if self.query_caching_ttl is not None:
            payload["queryCachingTTL"] = self.query_caching_ttl
        if self.interval is not None:
            payload["interval"] = self.interval
        if self.cache_timeout is not None:
            payload["cacheTimeout"] = self.cache_timeout
        if self.hide_time_override is not None:
            payload["hideTimeOverride"] = self.hide_time_override
        if self.time_compare is not None:
            payload["timeCompare"] = self.time_compare
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "timeFrom" in data:
            args["time_from"] = data["timeFrom"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "timeShift" in data:
            args["time_shift"] = data["timeShift"]
        if "queryCachingTTL" in data:
            args["query_caching_ttl"] = data["queryCachingTTL"]
        if "interval" in data:
            args["interval"] = data["interval"]
        if "cacheTimeout" in data:
            args["cache_timeout"] = data["cacheTimeout"]
        if "hideTimeOverride" in data:
            args["hide_time_override"] = data["hideTimeOverride"]
        if "timeCompare" in data:
            args["time_compare"] = data["timeCompare"]        

        return cls(**args)


class VizConfigKind:
    kind: typing.Literal["VizConfig"]
    # The group is the plugin ID
    group: str
    version: str
    spec: 'VizConfigSpec'

    def __init__(self, group: str = "", version: str = "", spec: typing.Optional['VizConfigSpec'] = None) -> None:
        self.kind = "VizConfig"
        self.group = group
        self.version = version
        self.spec = spec if spec is not None else VizConfigSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "group": self.group,
            "version": self.version,
            "spec": self.spec,
        }
        return payload

    
    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        defaults: dict[str, typing.Any] = {}
        custom: dict[str, typing.Any] = {}
        try:
            defaults = data["spec"]["fieldConfig"]["defaults"]
            custom = defaults.get("custom", {})
        except KeyError:
            pass
    
    
            
        if "group" in data:
            args["group"] = { data["group"]}
    
        if "version" in data:
            args["version"] = { data["version"]}
    
        if "spec" in data:
            args["spec"] = VizConfigSpec(options=cogruntime.panelcfg_options_from_json(data["spec"]["options"], data["kind"]) if data["spec"]["options"] is not None else None,field_config=FieldConfigSource.from_json({**data["spec"]["fieldConfig"], **{"defaults": {**defaults, **{"custom": cogruntime.panelcfg_field_config_from_json(custom, data["group"])}}}}) if data["spec"]["fieldConfig"] is not None else FieldConfigSource(),)
        return cls(**args)


class VizConfigSpec:
    """
    --- Kinds ---
    """

    options: typing.Optional[object]
    field_config: 'FieldConfigSource'

    def __init__(self, options: typing.Optional[object] = None, field_config: typing.Optional['FieldConfigSource'] = None) -> None:
        self.options = options
        self.field_config = field_config if field_config is not None else FieldConfigSource()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "options": self.options,
            "fieldConfig": self.field_config,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "options" in data:
            args["options"] = data["options"]
        if "fieldConfig" in data:
            args["field_config"] = FieldConfigSource.from_json(data["fieldConfig"])        

        return cls(**args)


class FieldConfigSource:
    """
    The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
    Each column within this structure is called a field. A field can represent a single time series or table column.
    Field options allow you to change how the data is displayed in your visualizations.
    """

    # Defaults are the options applied to all fields.
    defaults: 'FieldConfig'
    # Overrides are the options applied to specific fields overriding the defaults.
    overrides: list['Dashboardv2beta1FieldConfigSourceOverrides']

    def __init__(self, defaults: typing.Optional['FieldConfig'] = None, overrides: typing.Optional[list['Dashboardv2beta1FieldConfigSourceOverrides']] = None) -> None:
        self.defaults = defaults if defaults is not None else FieldConfig()
        self.overrides = overrides if overrides is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "defaults": self.defaults,
            "overrides": self.overrides,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "defaults" in data:
            args["defaults"] = FieldConfig.from_json(data["defaults"])
        if "overrides" in data:
            args["overrides"] = [Dashboardv2beta1FieldConfigSourceOverrides.from_json(item) for item in data["overrides"]]        

        return cls(**args)


class FieldConfig:
    """
    The data model used in Grafana, namely the data frame, is a columnar-oriented table structure that unifies both time series and table query results.
    Each column within this structure is called a field. A field can represent a single time series or table column.
    Field options allow you to change how the data is displayed in your visualizations.
    """

    # The display value for this field.  This supports template variables blank is auto
    display_name: typing.Optional[str]
    # This can be used by data sources that return and explicit naming structure for values and labels
    # When this property is configured, this value is used rather than the default naming strategy.
    display_name_from_ds: typing.Optional[str]
    # Human readable field metadata
    description: typing.Optional[str]
    # An explicit path to the field in the datasource.  When the frame meta includes a path,
    # This will default to `${frame.meta.path}/${field.name}
    # 
    # When defined, this value can be used as an identifier within the datasource scope, and
    # may be used to update the results
    path: typing.Optional[str]
    # True if data source can write a value to the path. Auth/authz are supported separately
    writeable: typing.Optional[bool]
    # True if data source field supports ad-hoc filters
    filterable: typing.Optional[bool]
    # Unit a field should use. The unit you select is applied to all fields except time.
    # You can use the units ID availables in Grafana or a custom unit.
    # Available units in Grafana: https://github.com/grafana/grafana/blob/main/packages/grafana-data/src/valueFormats/categories.ts
    # As custom unit, you can use the following formats:
    # `suffix:<suffix>` for custom unit that should go after value.
    # `prefix:<prefix>` for custom unit that should go before value.
    # `time:<format>` For custom date time formats type for example `time:YYYY-MM-DD`.
    # `si:<base scale><unit characters>` for custom SI units. For example: `si: mF`. This one is a bit more advanced as you can specify both a unit and the source data scale. So if your source data is represented as milli (thousands of) something prefix the unit with that SI scale character.
    # `count:<unit>` for a custom count unit.
    # `currency:<unit>` for custom a currency unit.
    unit: typing.Optional[str]
    # Specify the number of decimals Grafana includes in the rendered value.
    # If you leave this field blank, Grafana automatically truncates the number of decimals based on the value.
    # For example 1.1234 will display as 1.12 and 100.456 will display as 100.
    # To display all decimals, set the unit to `String`.
    decimals: typing.Optional[float]
    # The minimum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    min_val: typing.Optional[float]
    # The maximum value used in percentage threshold calculations. Leave blank for auto calculation based on all series and fields.
    max_val: typing.Optional[float]
    # Convert input values into a display string
    mappings: typing.Optional[list['ValueMapping']]
    # Map numeric values to states
    thresholds: typing.Optional['ThresholdsConfig']
    # Panel color configuration
    color: typing.Optional['FieldColor']
    # The behavior when clicking on a result
    links: typing.Optional[list[object]]
    # Define interactive HTTP requests that can be triggered from data visualizations.
    actions: typing.Optional[list['Action']]
    # Alternative to empty string
    no_value: typing.Optional[str]
    # custom is specified by the FieldConfig field
    # in panel plugin schemas.
    custom: typing.Optional[object]
    # Calculate min max per field
    field_min_max: typing.Optional[bool]
    # How null values should be handled when calculating field stats
    # "null" - Include null values, "connected" - Ignore nulls, "null as zero" - Treat nulls as zero
    null_value_mode: typing.Optional['NullValueMode']

    def __init__(self, display_name: typing.Optional[str] = None, display_name_from_ds: typing.Optional[str] = None, description: typing.Optional[str] = None, path: typing.Optional[str] = None, writeable: typing.Optional[bool] = None, filterable: typing.Optional[bool] = None, unit: typing.Optional[str] = None, decimals: typing.Optional[float] = None, min_val: typing.Optional[float] = None, max_val: typing.Optional[float] = None, mappings: typing.Optional[list['ValueMapping']] = None, thresholds: typing.Optional['ThresholdsConfig'] = None, color: typing.Optional['FieldColor'] = None, links: typing.Optional[list[object]] = None, actions: typing.Optional[list['Action']] = None, no_value: typing.Optional[str] = None, custom: typing.Optional[object] = None, field_min_max: typing.Optional[bool] = None, null_value_mode: typing.Optional['NullValueMode'] = None) -> None:
        self.display_name = display_name
        self.display_name_from_ds = display_name_from_ds
        self.description = description
        self.path = path
        self.writeable = writeable
        self.filterable = filterable
        self.unit = unit
        self.decimals = decimals
        self.min_val = min_val
        self.max_val = max_val
        self.mappings = mappings
        self.thresholds = thresholds
        self.color = color
        self.links = links
        self.actions = actions
        self.no_value = no_value
        self.custom = custom
        self.field_min_max = field_min_max
        self.null_value_mode = null_value_mode

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.display_name is not None:
            payload["displayName"] = self.display_name
        if self.display_name_from_ds is not None:
            payload["displayNameFromDS"] = self.display_name_from_ds
        if self.description is not None:
            payload["description"] = self.description
        if self.path is not None:
            payload["path"] = self.path
        if self.writeable is not None:
            payload["writeable"] = self.writeable
        if self.filterable is not None:
            payload["filterable"] = self.filterable
        if self.unit is not None:
            payload["unit"] = self.unit
        if self.decimals is not None:
            payload["decimals"] = self.decimals
        if self.min_val is not None:
            payload["min"] = self.min_val
        if self.max_val is not None:
            payload["max"] = self.max_val
        if self.mappings is not None:
            payload["mappings"] = self.mappings
        if self.thresholds is not None:
            payload["thresholds"] = self.thresholds
        if self.color is not None:
            payload["color"] = self.color
        if self.links is not None:
            payload["links"] = self.links
        if self.actions is not None:
            payload["actions"] = self.actions
        if self.no_value is not None:
            payload["noValue"] = self.no_value
        if self.custom is not None:
            payload["custom"] = self.custom
        if self.field_min_max is not None:
            payload["fieldMinMax"] = self.field_min_max
        if self.null_value_mode is not None:
            payload["nullValueMode"] = self.null_value_mode
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "displayName" in data:
            args["display_name"] = data["displayName"]
        if "displayNameFromDS" in data:
            args["display_name_from_ds"] = data["displayNameFromDS"]
        if "description" in data:
            args["description"] = data["description"]
        if "path" in data:
            args["path"] = data["path"]
        if "writeable" in data:
            args["writeable"] = data["writeable"]
        if "filterable" in data:
            args["filterable"] = data["filterable"]
        if "unit" in data:
            args["unit"] = data["unit"]
        if "decimals" in data:
            args["decimals"] = data["decimals"]
        if "min" in data:
            args["min_val"] = data["min"]
        if "max" in data:
            args["max_val"] = data["max"]
        if "mappings" in data:
            decoding_map_mappings_array_ref_union: dict[str, typing.Union[typing.Type[RangeMap], typing.Type[RegexMap], typing.Type[SpecialValueMap], typing.Type[ValueMap]]] = {"range": RangeMap, "regex": RegexMap, "special": SpecialValueMap, "value": ValueMap}
            args["mappings"] = [decoding_map_mappings_array_ref_union[item["type"]].from_json(item) for item in data["mappings"]]
        if "thresholds" in data:
            args["thresholds"] = ThresholdsConfig.from_json(data["thresholds"])
        if "color" in data:
            args["color"] = FieldColor.from_json(data["color"])
        if "links" in data:
            args["links"] = data["links"]
        if "actions" in data:
            args["actions"] = [Action.from_json(item) for item in data["actions"]]
        if "noValue" in data:
            args["no_value"] = data["noValue"]
        if "custom" in data:
            args["custom"] = data["custom"]
        if "fieldMinMax" in data:
            args["field_min_max"] = data["fieldMinMax"]
        if "nullValueMode" in data:
            args["null_value_mode"] = data["nullValueMode"]        

        return cls(**args)


ValueMapping: typing.TypeAlias = typing.Union['ValueMap', 'RangeMap', 'RegexMap', 'SpecialValueMap']


class ValueMap:
    """
    Maps text values to a color or different display text and color.
    For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
    """

    type_val: str
    # Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
    options: dict[str, 'ValueMappingResult']

    def __init__(self, options: typing.Optional[dict[str, 'ValueMappingResult']] = None) -> None:
        self.type_val = MappingType.VALUE
        self.options = options if options is not None else {}

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "options": self.options,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "options" in data:
            args["options"] = {key: ValueMappingResult.from_json(data["options"][key]) for key in data["options"].keys()}        

        return cls(**args)


class MappingType(enum.StrEnum):
    """
    Supported value mapping types
    `value`: Maps text values to a color or different display text and color. For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.
    `range`: Maps numerical ranges to a display text and color. For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
    `regex`: Maps regular expressions to replacement text and a color. For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
    `special`: Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color. See SpecialValueMatch to see the list of special values. For example, you can configure a special value mapping so that null values appear as N/A.
    """

    VALUE = "value"
    RANGE = "range"
    REGEX = "regex"
    SPECIAL = "special"


class ValueMappingResult:
    """
    Result used as replacement with text and color when the value matches
    """

    # Text to display when the value matches
    text: typing.Optional[str]
    # Text to use when the value matches
    color: typing.Optional[str]
    # Icon to display when the value matches. Only specific visualizations.
    icon: typing.Optional[str]
    # Position in the mapping array. Only used internally.
    index: typing.Optional[int]

    def __init__(self, text: typing.Optional[str] = None, color: typing.Optional[str] = None, icon: typing.Optional[str] = None, index: typing.Optional[int] = None) -> None:
        self.text = text
        self.color = color
        self.icon = icon
        self.index = index

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.text is not None:
            payload["text"] = self.text
        if self.color is not None:
            payload["color"] = self.color
        if self.icon is not None:
            payload["icon"] = self.icon
        if self.index is not None:
            payload["index"] = self.index
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "text" in data:
            args["text"] = data["text"]
        if "color" in data:
            args["color"] = data["color"]
        if "icon" in data:
            args["icon"] = data["icon"]
        if "index" in data:
            args["index"] = data["index"]        

        return cls(**args)


class RangeMap:
    """
    Maps numerical ranges to a display text and color.
    For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
    """

    type_val: str
    # Range to match against and the result to apply when the value is within the range
    options: 'Dashboardv2beta1RangeMapOptions'

    def __init__(self, options: typing.Optional['Dashboardv2beta1RangeMapOptions'] = None) -> None:
        self.type_val = MappingType.RANGE
        self.options = options if options is not None else Dashboardv2beta1RangeMapOptions()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "options": self.options,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "options" in data:
            args["options"] = Dashboardv2beta1RangeMapOptions.from_json(data["options"])        

        return cls(**args)


class RegexMap:
    """
    Maps regular expressions to replacement text and a color.
    For example, if a value is www.example.com, you can configure a regex value mapping so that Grafana displays www and truncates the domain.
    """

    type_val: str
    # Regular expression to match against and the result to apply when the value matches the regex
    options: 'Dashboardv2beta1RegexMapOptions'

    def __init__(self, options: typing.Optional['Dashboardv2beta1RegexMapOptions'] = None) -> None:
        self.type_val = MappingType.REGEX
        self.options = options if options is not None else Dashboardv2beta1RegexMapOptions()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "options": self.options,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "options" in data:
            args["options"] = Dashboardv2beta1RegexMapOptions.from_json(data["options"])        

        return cls(**args)


class SpecialValueMap:
    """
    Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.
    See SpecialValueMatch to see the list of special values.
    For example, you can configure a special value mapping so that null values appear as N/A.
    """

    type_val: str
    options: 'Dashboardv2beta1SpecialValueMapOptions'

    def __init__(self, options: typing.Optional['Dashboardv2beta1SpecialValueMapOptions'] = None) -> None:
        self.type_val = MappingType.SPECIAL
        self.options = options if options is not None else Dashboardv2beta1SpecialValueMapOptions()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "options": self.options,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "options" in data:
            args["options"] = Dashboardv2beta1SpecialValueMapOptions.from_json(data["options"])        

        return cls(**args)


class SpecialValueMatch(enum.StrEnum):
    """
    Special value types supported by the `SpecialValueMap`
    """

    TRUE = "true"
    FALSE = "false"
    NULL = "null"
    NA_N = "nan"
    NULL_AND_NA_N = "null+nan"
    EMPTY = "empty"


class ThresholdsConfig:
    mode: 'ThresholdsMode'
    steps: list['Threshold']

    def __init__(self, mode: typing.Optional['ThresholdsMode'] = None, steps: typing.Optional[list['Threshold']] = None) -> None:
        self.mode = mode if mode is not None else ThresholdsMode.ABSOLUTE
        self.steps = steps if steps is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
            "steps": self.steps,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "steps" in data:
            args["steps"] = [Threshold.from_json(item) for item in data["steps"]]        

        return cls(**args)


class ThresholdsMode(enum.StrEnum):
    ABSOLUTE = "absolute"
    PERCENTAGE = "percentage"


class Threshold:
    # Value null means -Infinity
    value: typing.Optional[float]
    color: str

    def __init__(self, value: typing.Optional[float] = None, color: str = "") -> None:
        self.value = value
        self.color = color

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "value": self.value,
            "color": self.color,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]
        if "color" in data:
            args["color"] = data["color"]        

        return cls(**args)


class FieldColor:
    """
    Map a field to a color.
    """

    # The main color scheme mode.
    mode: 'FieldColorModeId'
    # The fixed color value for fixed or shades color modes.
    fixed_color: typing.Optional[str]
    # Some visualizations need to know how to assign a series color from by value color schemes.
    series_by: typing.Optional['FieldColorSeriesByMode']

    def __init__(self, mode: typing.Optional['FieldColorModeId'] = None, fixed_color: typing.Optional[str] = None, series_by: typing.Optional['FieldColorSeriesByMode'] = None) -> None:
        self.mode = mode if mode is not None else FieldColorModeId.THRESHOLDS
        self.fixed_color = fixed_color
        self.series_by = series_by

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        if self.fixed_color is not None:
            payload["fixedColor"] = self.fixed_color
        if self.series_by is not None:
            payload["seriesBy"] = self.series_by
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "fixedColor" in data:
            args["fixed_color"] = data["fixedColor"]
        if "seriesBy" in data:
            args["series_by"] = data["seriesBy"]        

        return cls(**args)


class FieldColorModeId(enum.StrEnum):
    """
    Color mode for a field. You can specify a single color, or select a continuous (gradient) color schemes, based on a value.
    Continuous color interpolates a color using the percentage of a value relative to min and max.
    Accepted values are:
    `thresholds`: From thresholds. Informs Grafana to take the color from the matching threshold
    `palette-classic`: Classic palette. Grafana will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations
    `palette-classic-by-name`: Classic palette (by name). Grafana will assign color by looking up a color in a palette by series name. Useful for Graphs and pie charts and other categorical data visualizations
    `continuous-viridis`: Continuous Viridis palette mode
    `continuous-magma`: Continuous Magma palette mode
    `continuous-plasma`: Continuous Plasma palette mode
    `continuous-inferno`: Continuous Inferno palette mode
    `continuous-cividis`: Continuous Cividis palette mode
    `continuous-GrYlRd`: Continuous Green-Yellow-Red palette mode
    `continuous-RdYlGr`: Continuous Red-Yellow-Green palette mode
    `continuous-BlYlRd`: Continuous Blue-Yellow-Red palette mode
    `continuous-YlRd`: Continuous Yellow-Red palette mode
    `continuous-BlPu`: Continuous Blue-Purple palette mode
    `continuous-YlBl`: Continuous Yellow-Blue palette mode
    `continuous-blues`: Continuous Blue palette mode
    `continuous-reds`: Continuous Red palette mode
    `continuous-greens`: Continuous Green palette mode
    `continuous-purples`: Continuous Purple palette mode
    `shades`: Shades of a single color. Specify a single color, useful in an override rule.
    `fixed`: Fixed color mode. Specify a single color, useful in an override rule.
    """

    THRESHOLDS = "thresholds"
    PALETTE_CLASSIC = "palette-classic"
    PALETTE_CLASSIC_BY_NAME = "palette-classic-by-name"
    CONTINUOUS_VIRIDIS = "continuous-viridis"
    CONTINUOUS_MAGMA = "continuous-magma"
    CONTINUOUS_PLASMA = "continuous-plasma"
    CONTINUOUS_INFERNO = "continuous-inferno"
    CONTINUOUS_CIVIDIS = "continuous-cividis"
    CONTINUOUS_GR_YL_RD = "continuous-GrYlRd"
    CONTINUOUS_RD_YL_GR = "continuous-RdYlGr"
    CONTINUOUS_BL_YL_RD = "continuous-BlYlRd"
    CONTINUOUS_YL_RD = "continuous-YlRd"
    CONTINUOUS_BL_PU = "continuous-BlPu"
    CONTINUOUS_YL_BL = "continuous-YlBl"
    CONTINUOUS_BLUES = "continuous-blues"
    CONTINUOUS_REDS = "continuous-reds"
    CONTINUOUS_GREENS = "continuous-greens"
    CONTINUOUS_PURPLES = "continuous-purples"
    FIXED = "fixed"
    SHADES = "shades"


class FieldColorSeriesByMode(enum.StrEnum):
    """
    Defines how to assign a series color from "by value" color schemes. For example for an aggregated data points like a timeseries, the color can be assigned by the min, max or last value.
    """

    MIN = "min"
    MAX = "max"
    LAST = "last"


class Action:
    type_val: 'ActionType'
    title: str
    fetch: typing.Optional['FetchOptions']
    infinity: typing.Optional['InfinityOptions']
    confirmation: typing.Optional[str]
    one_click: typing.Optional[bool]
    variables: typing.Optional[list['ActionVariable']]
    style: typing.Optional['Dashboardv2beta1ActionStyle']

    def __init__(self, type_val: typing.Optional['ActionType'] = None, title: str = "", fetch: typing.Optional['FetchOptions'] = None, infinity: typing.Optional['InfinityOptions'] = None, confirmation: typing.Optional[str] = None, one_click: typing.Optional[bool] = None, variables: typing.Optional[list['ActionVariable']] = None, style: typing.Optional['Dashboardv2beta1ActionStyle'] = None) -> None:
        self.type_val = type_val if type_val is not None else ActionType.FETCH
        self.title = title
        self.fetch = fetch
        self.infinity = infinity
        self.confirmation = confirmation
        self.one_click = one_click
        self.variables = variables
        self.style = style

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "title": self.title,
        }
        if self.fetch is not None:
            payload["fetch"] = self.fetch
        if self.infinity is not None:
            payload["infinity"] = self.infinity
        if self.confirmation is not None:
            payload["confirmation"] = self.confirmation
        if self.one_click is not None:
            payload["oneClick"] = self.one_click
        if self.variables is not None:
            payload["variables"] = self.variables
        if self.style is not None:
            payload["style"] = self.style
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "title" in data:
            args["title"] = data["title"]
        if "fetch" in data:
            args["fetch"] = FetchOptions.from_json(data["fetch"])
        if "infinity" in data:
            args["infinity"] = InfinityOptions.from_json(data["infinity"])
        if "confirmation" in data:
            args["confirmation"] = data["confirmation"]
        if "oneClick" in data:
            args["one_click"] = data["oneClick"]
        if "variables" in data:
            args["variables"] = [ActionVariable.from_json(item) for item in data["variables"]]
        if "style" in data:
            args["style"] = Dashboardv2beta1ActionStyle.from_json(data["style"])        

        return cls(**args)


class ActionType(enum.StrEnum):
    FETCH = "fetch"
    INFINITY = "infinity"


class FetchOptions:
    method: 'HttpRequestMethod'
    url: str
    body: typing.Optional[str]
    # These are 2D arrays of strings, each representing a key-value pair
    # We are defining them this way because we can't generate a go struct that
    # that would have exactly two strings in each sub-array
    query_params: typing.Optional[list[list[str]]]
    headers: typing.Optional[list[list[str]]]

    def __init__(self, method: typing.Optional['HttpRequestMethod'] = None, url: str = "", body: typing.Optional[str] = None, query_params: typing.Optional[list[list[str]]] = None, headers: typing.Optional[list[list[str]]] = None) -> None:
        self.method = method if method is not None else HttpRequestMethod.GET
        self.url = url
        self.body = body
        self.query_params = query_params
        self.headers = headers

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "method": self.method,
            "url": self.url,
        }
        if self.body is not None:
            payload["body"] = self.body
        if self.query_params is not None:
            payload["queryParams"] = self.query_params
        if self.headers is not None:
            payload["headers"] = self.headers
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "method" in data:
            args["method"] = data["method"]
        if "url" in data:
            args["url"] = data["url"]
        if "body" in data:
            args["body"] = data["body"]
        if "queryParams" in data:
            args["query_params"] = [item for item in data["queryParams"]]
        if "headers" in data:
            args["headers"] = [item for item in data["headers"]]        

        return cls(**args)


class HttpRequestMethod(enum.StrEnum):
    GET = "GET"
    PUT = "PUT"
    POST = "POST"
    DELETE = "DELETE"
    PATCH = "PATCH"


class InfinityOptions:
    method: 'HttpRequestMethod'
    url: str
    body: typing.Optional[str]
    # These are 2D arrays of strings, each representing a key-value pair
    # We are defining them this way because we can't generate a go struct that
    # that would have exactly two strings in each sub-array
    query_params: typing.Optional[list[list[str]]]
    datasource_uid: str
    headers: typing.Optional[list[list[str]]]

    def __init__(self, method: typing.Optional['HttpRequestMethod'] = None, url: str = "", body: typing.Optional[str] = None, query_params: typing.Optional[list[list[str]]] = None, datasource_uid: str = "", headers: typing.Optional[list[list[str]]] = None) -> None:
        self.method = method if method is not None else HttpRequestMethod.GET
        self.url = url
        self.body = body
        self.query_params = query_params
        self.datasource_uid = datasource_uid
        self.headers = headers

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "method": self.method,
            "url": self.url,
            "datasourceUid": self.datasource_uid,
        }
        if self.body is not None:
            payload["body"] = self.body
        if self.query_params is not None:
            payload["queryParams"] = self.query_params
        if self.headers is not None:
            payload["headers"] = self.headers
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "method" in data:
            args["method"] = data["method"]
        if "url" in data:
            args["url"] = data["url"]
        if "body" in data:
            args["body"] = data["body"]
        if "queryParams" in data:
            args["query_params"] = [item for item in data["queryParams"]]
        if "datasourceUid" in data:
            args["datasource_uid"] = data["datasourceUid"]
        if "headers" in data:
            args["headers"] = [item for item in data["headers"]]        

        return cls(**args)


class ActionVariable:
    key: str
    name: str
    type_val: str

    def __init__(self, key: str = "", name: str = "") -> None:
        self.key = key
        self.name = name
        self.type_val = ActionVariableType

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "key": self.key,
            "name": self.name,
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "key" in data:
            args["key"] = data["key"]
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


# Action variable type
ActionVariableType: typing.Literal["string"] = "string"


class NullValueMode(enum.StrEnum):
    """
    How null values should be handled
    """

    NULL = "null"
    CONNECTED = "connected"
    NULL_AS_ZERO = "null as zero"


class DynamicConfigValue:
    id_val: str
    value: typing.Optional[object]

    def __init__(self, id_val: str = "", value: typing.Optional[object] = None) -> None:
        self.id_val = id_val
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
        }
        if self.value is not None:
            payload["value"] = self.value
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class LibraryPanelKind:
    kind: typing.Literal["LibraryPanel"]
    spec: 'LibraryPanelKindSpec'

    def __init__(self, spec: typing.Optional['LibraryPanelKindSpec'] = None) -> None:
        self.kind = "LibraryPanel"
        self.spec = spec if spec is not None else LibraryPanelKindSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = LibraryPanelKindSpec.from_json(data["spec"])        

        return cls(**args)


class LibraryPanelKindSpec:
    # Panel ID for the library panel in the dashboard
    id_val: float
    # Title for the library panel in the dashboard
    title: str
    library_panel: 'LibraryPanelRef'

    def __init__(self, id_val: float = 0, title: str = "", library_panel: typing.Optional['LibraryPanelRef'] = None) -> None:
        self.id_val = id_val
        self.title = title
        self.library_panel = library_panel if library_panel is not None else LibraryPanelRef()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "title": self.title,
            "libraryPanel": self.library_panel,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "title" in data:
            args["title"] = data["title"]
        if "libraryPanel" in data:
            args["library_panel"] = LibraryPanelRef.from_json(data["libraryPanel"])        

        return cls(**args)


class LibraryPanelRef:
    """
    A library panel is a reusable panel that you can use in any dashboard.
    When you make a change to a library panel, that change propagates to all instances of where the panel is used.
    Library panels streamline reuse of panels across multiple dashboards.
    """

    # Library panel name
    name: str
    # Library panel uid
    uid: str

    def __init__(self, name: str = "", uid: str = "") -> None:
        self.name = name
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "uid": self.uid,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class GridLayoutKind:
    kind: typing.Literal["GridLayout"]
    spec: 'GridLayoutSpec'

    def __init__(self, spec: typing.Optional['GridLayoutSpec'] = None) -> None:
        self.kind = "GridLayout"
        self.spec = spec if spec is not None else GridLayoutSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = GridLayoutSpec.from_json(data["spec"])        

        return cls(**args)


class GridLayoutSpec:
    items: list['GridLayoutItemKind']

    def __init__(self, items: typing.Optional[list['GridLayoutItemKind']] = None) -> None:
        self.items = items if items is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "items": self.items,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "items" in data:
            args["items"] = [GridLayoutItemKind.from_json(item) for item in data["items"]]        

        return cls(**args)


class GridLayoutItemKind:
    kind: typing.Literal["GridLayoutItem"]
    spec: 'GridLayoutItemSpec'

    def __init__(self, spec: typing.Optional['GridLayoutItemSpec'] = None) -> None:
        self.kind = "GridLayoutItem"
        self.spec = spec if spec is not None else GridLayoutItemSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = GridLayoutItemSpec.from_json(data["spec"])        

        return cls(**args)


class GridLayoutItemSpec:
    x: int
    y: int
    width: int
    height: int
    # reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
    element: 'ElementReference'
    repeat: typing.Optional['RepeatOptions']

    def __init__(self, x: int = 0, y: int = 0, width: int = 0, height: int = 0, element: typing.Optional['ElementReference'] = None, repeat: typing.Optional['RepeatOptions'] = None) -> None:
        self.x = x
        self.y = y
        self.width = width
        self.height = height
        self.element = element if element is not None else ElementReference()
        self.repeat = repeat

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "x": self.x,
            "y": self.y,
            "width": self.width,
            "height": self.height,
            "element": self.element,
        }
        if self.repeat is not None:
            payload["repeat"] = self.repeat
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "x" in data:
            args["x"] = data["x"]
        if "y" in data:
            args["y"] = data["y"]
        if "width" in data:
            args["width"] = data["width"]
        if "height" in data:
            args["height"] = data["height"]
        if "element" in data:
            args["element"] = ElementReference.from_json(data["element"])
        if "repeat" in data:
            args["repeat"] = RepeatOptions.from_json(data["repeat"])        

        return cls(**args)


class ElementReference:
    kind: typing.Literal["ElementReference"]
    name: str

    def __init__(self, name: str = "") -> None:
        self.kind = "ElementReference"
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "name": self.name,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class RepeatOptions:
    mode: str
    value: str
    direction: typing.Optional[typing.Literal["h", "v"]]
    max_per_row: typing.Optional[int]

    def __init__(self, value: str = "", direction: typing.Optional[typing.Literal["h", "v"]] = None, max_per_row: typing.Optional[int] = None) -> None:
        self.mode = RepeatMode.VARIABLE
        self.value = value
        self.direction = direction
        self.max_per_row = max_per_row

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
            "value": self.value,
        }
        if self.direction is not None:
            payload["direction"] = self.direction
        if self.max_per_row is not None:
            payload["maxPerRow"] = self.max_per_row
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]
        if "direction" in data:
            args["direction"] = data["direction"]
        if "maxPerRow" in data:
            args["max_per_row"] = data["maxPerRow"]        

        return cls(**args)


class RepeatMode(enum.StrEnum):
    """
    other repeat modes will be added in the future: label, frame
    """

    VARIABLE = "variable"


class RowsLayoutKind:
    kind: typing.Literal["RowsLayout"]
    spec: 'RowsLayoutSpec'

    def __init__(self, spec: typing.Optional['RowsLayoutSpec'] = None) -> None:
        self.kind = "RowsLayout"
        self.spec = spec if spec is not None else RowsLayoutSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = RowsLayoutSpec.from_json(data["spec"])        

        return cls(**args)


class RowsLayoutSpec:
    rows: list['RowsLayoutRowKind']

    def __init__(self, rows: typing.Optional[list['RowsLayoutRowKind']] = None) -> None:
        self.rows = rows if rows is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "rows": self.rows,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "rows" in data:
            args["rows"] = [RowsLayoutRowKind.from_json(item) for item in data["rows"]]        

        return cls(**args)


class RowsLayoutRowKind:
    kind: typing.Literal["RowsLayoutRow"]
    spec: 'RowsLayoutRowSpec'

    def __init__(self, spec: typing.Optional['RowsLayoutRowSpec'] = None) -> None:
        self.kind = "RowsLayoutRow"
        self.spec = spec if spec is not None else RowsLayoutRowSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = RowsLayoutRowSpec.from_json(data["spec"])        

        return cls(**args)


class RowsLayoutRowSpec:
    title: typing.Optional[str]
    collapse: typing.Optional[bool]
    hide_header: typing.Optional[bool]
    fill_screen: typing.Optional[bool]
    conditional_rendering: typing.Optional['ConditionalRenderingGroupKind']
    repeat: typing.Optional['RowRepeatOptions']
    layout: typing.Union['GridLayoutKind', 'AutoGridLayoutKind', 'TabsLayoutKind', 'RowsLayoutKind']

    def __init__(self, title: typing.Optional[str] = None, collapse: typing.Optional[bool] = None, hide_header: typing.Optional[bool] = None, fill_screen: typing.Optional[bool] = None, conditional_rendering: typing.Optional['ConditionalRenderingGroupKind'] = None, repeat: typing.Optional['RowRepeatOptions'] = None, layout: typing.Optional[typing.Union['GridLayoutKind', 'AutoGridLayoutKind', 'TabsLayoutKind', 'RowsLayoutKind']] = None) -> None:
        self.title = title
        self.collapse = collapse
        self.hide_header = hide_header
        self.fill_screen = fill_screen
        self.conditional_rendering = conditional_rendering
        self.repeat = repeat
        self.layout = layout if layout is not None else GridLayoutKind()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "layout": self.layout,
        }
        if self.title is not None:
            payload["title"] = self.title
        if self.collapse is not None:
            payload["collapse"] = self.collapse
        if self.hide_header is not None:
            payload["hideHeader"] = self.hide_header
        if self.fill_screen is not None:
            payload["fillScreen"] = self.fill_screen
        if self.conditional_rendering is not None:
            payload["conditionalRendering"] = self.conditional_rendering
        if self.repeat is not None:
            payload["repeat"] = self.repeat
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "collapse" in data:
            args["collapse"] = data["collapse"]
        if "hideHeader" in data:
            args["hide_header"] = data["hideHeader"]
        if "fillScreen" in data:
            args["fill_screen"] = data["fillScreen"]
        if "conditionalRendering" in data:
            args["conditional_rendering"] = ConditionalRenderingGroupKind.from_json(data["conditionalRendering"])
        if "repeat" in data:
            args["repeat"] = RowRepeatOptions.from_json(data["repeat"])
        if "layout" in data:
            decoding_map_layout_union: dict[str, typing.Union[typing.Type[AutoGridLayoutKind], typing.Type[GridLayoutKind], typing.Type[RowsLayoutKind], typing.Type[TabsLayoutKind]]] = {"AutoGridLayout": AutoGridLayoutKind, "GridLayout": GridLayoutKind, "RowsLayout": RowsLayoutKind, "TabsLayout": TabsLayoutKind}
            args["layout"] = decoding_map_layout_union[data["layout"]["kind"]].from_json(data["layout"])        

        return cls(**args)


class ConditionalRenderingGroupKind:
    kind: typing.Literal["ConditionalRenderingGroup"]
    spec: 'ConditionalRenderingGroupSpec'

    def __init__(self, spec: typing.Optional['ConditionalRenderingGroupSpec'] = None) -> None:
        self.kind = "ConditionalRenderingGroup"
        self.spec = spec if spec is not None else ConditionalRenderingGroupSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = ConditionalRenderingGroupSpec.from_json(data["spec"])        

        return cls(**args)


class ConditionalRenderingGroupSpec:
    visibility: typing.Literal["show", "hide"]
    condition: typing.Literal["and", "or"]
    items: list[typing.Union['ConditionalRenderingVariableKind', 'ConditionalRenderingDataKind', 'ConditionalRenderingTimeRangeSizeKind']]

    def __init__(self, visibility: typing.Optional[typing.Literal["show", "hide"]] = None, condition: typing.Optional[typing.Literal["and", "or"]] = None, items: typing.Optional[list[typing.Union['ConditionalRenderingVariableKind', 'ConditionalRenderingDataKind', 'ConditionalRenderingTimeRangeSizeKind']]] = None) -> None:
        self.visibility = visibility if visibility is not None else "show"
        self.condition = condition if condition is not None else "and"
        self.items = items if items is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "visibility": self.visibility,
            "condition": self.condition,
            "items": self.items,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "visibility" in data:
            args["visibility"] = data["visibility"]
        if "condition" in data:
            args["condition"] = data["condition"]
        if "items" in data:
            decoding_map_items_array_union: dict[str, typing.Union[typing.Type[ConditionalRenderingDataKind], typing.Type[ConditionalRenderingTimeRangeSizeKind], typing.Type[ConditionalRenderingVariableKind]]] = {"ConditionalRenderingData": ConditionalRenderingDataKind, "ConditionalRenderingTimeRangeSize": ConditionalRenderingTimeRangeSizeKind, "ConditionalRenderingVariable": ConditionalRenderingVariableKind}
            args["items"] = [decoding_map_items_array_union[item["kind"]].from_json(item) for item in data["items"]]        

        return cls(**args)


class ConditionalRenderingVariableKind:
    kind: typing.Literal["ConditionalRenderingVariable"]
    spec: 'ConditionalRenderingVariableSpec'

    def __init__(self, spec: typing.Optional['ConditionalRenderingVariableSpec'] = None) -> None:
        self.kind = "ConditionalRenderingVariable"
        self.spec = spec if spec is not None else ConditionalRenderingVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = ConditionalRenderingVariableSpec.from_json(data["spec"])        

        return cls(**args)


class ConditionalRenderingVariableSpec:
    variable: str
    operator: typing.Literal["equals", "notEquals", "matches", "notMatches"]
    value: str

    def __init__(self, variable: str = "", operator: typing.Optional[typing.Literal["equals", "notEquals", "matches", "notMatches"]] = None, value: str = "") -> None:
        self.variable = variable
        self.operator = operator if operator is not None else "equals"
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "variable": self.variable,
            "operator": self.operator,
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "variable" in data:
            args["variable"] = data["variable"]
        if "operator" in data:
            args["operator"] = data["operator"]
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class ConditionalRenderingDataKind:
    kind: typing.Literal["ConditionalRenderingData"]
    spec: 'ConditionalRenderingDataSpec'

    def __init__(self, spec: typing.Optional['ConditionalRenderingDataSpec'] = None) -> None:
        self.kind = "ConditionalRenderingData"
        self.spec = spec if spec is not None else ConditionalRenderingDataSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = ConditionalRenderingDataSpec.from_json(data["spec"])        

        return cls(**args)


class ConditionalRenderingDataSpec:
    value: bool

    def __init__(self, value: bool = False) -> None:
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class ConditionalRenderingTimeRangeSizeKind:
    kind: typing.Literal["ConditionalRenderingTimeRangeSize"]
    spec: 'ConditionalRenderingTimeRangeSizeSpec'

    def __init__(self, spec: typing.Optional['ConditionalRenderingTimeRangeSizeSpec'] = None) -> None:
        self.kind = "ConditionalRenderingTimeRangeSize"
        self.spec = spec if spec is not None else ConditionalRenderingTimeRangeSizeSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = ConditionalRenderingTimeRangeSizeSpec.from_json(data["spec"])        

        return cls(**args)


class ConditionalRenderingTimeRangeSizeSpec:
    value: str

    def __init__(self, value: str = "") -> None:
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class RowRepeatOptions:
    mode: str
    value: str

    def __init__(self, value: str = "") -> None:
        self.mode = RepeatMode.VARIABLE
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class AutoGridLayoutKind:
    kind: typing.Literal["AutoGridLayout"]
    spec: 'AutoGridLayoutSpec'

    def __init__(self, spec: typing.Optional['AutoGridLayoutSpec'] = None) -> None:
        self.kind = "AutoGridLayout"
        self.spec = spec if spec is not None else AutoGridLayoutSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = AutoGridLayoutSpec.from_json(data["spec"])        

        return cls(**args)


class AutoGridLayoutSpec:
    max_column_count: typing.Optional[float]
    column_width_mode: typing.Literal["narrow", "standard", "wide", "custom"]
    column_width: typing.Optional[float]
    row_height_mode: typing.Literal["short", "standard", "tall", "custom"]
    row_height: typing.Optional[float]
    fill_screen: typing.Optional[bool]
    items: list['AutoGridLayoutItemKind']

    def __init__(self, max_column_count: typing.Optional[float] = 3, column_width_mode: typing.Optional[typing.Literal["narrow", "standard", "wide", "custom"]] = None, column_width: typing.Optional[float] = None, row_height_mode: typing.Optional[typing.Literal["short", "standard", "tall", "custom"]] = None, row_height: typing.Optional[float] = None, fill_screen: typing.Optional[bool] = False, items: typing.Optional[list['AutoGridLayoutItemKind']] = None) -> None:
        self.max_column_count = max_column_count
        self.column_width_mode = column_width_mode if column_width_mode is not None else "standard"
        self.column_width = column_width
        self.row_height_mode = row_height_mode if row_height_mode is not None else "standard"
        self.row_height = row_height
        self.fill_screen = fill_screen
        self.items = items if items is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "columnWidthMode": self.column_width_mode,
            "rowHeightMode": self.row_height_mode,
            "items": self.items,
        }
        if self.max_column_count is not None:
            payload["maxColumnCount"] = self.max_column_count
        if self.column_width is not None:
            payload["columnWidth"] = self.column_width
        if self.row_height is not None:
            payload["rowHeight"] = self.row_height
        if self.fill_screen is not None:
            payload["fillScreen"] = self.fill_screen
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "maxColumnCount" in data:
            args["max_column_count"] = data["maxColumnCount"]
        if "columnWidthMode" in data:
            args["column_width_mode"] = data["columnWidthMode"]
        if "columnWidth" in data:
            args["column_width"] = data["columnWidth"]
        if "rowHeightMode" in data:
            args["row_height_mode"] = data["rowHeightMode"]
        if "rowHeight" in data:
            args["row_height"] = data["rowHeight"]
        if "fillScreen" in data:
            args["fill_screen"] = data["fillScreen"]
        if "items" in data:
            args["items"] = [AutoGridLayoutItemKind.from_json(item) for item in data["items"]]        

        return cls(**args)


class AutoGridLayoutItemKind:
    kind: typing.Literal["AutoGridLayoutItem"]
    spec: 'AutoGridLayoutItemSpec'

    def __init__(self, spec: typing.Optional['AutoGridLayoutItemSpec'] = None) -> None:
        self.kind = "AutoGridLayoutItem"
        self.spec = spec if spec is not None else AutoGridLayoutItemSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = AutoGridLayoutItemSpec.from_json(data["spec"])        

        return cls(**args)


class AutoGridLayoutItemSpec:
    element: 'ElementReference'
    repeat: typing.Optional['AutoGridRepeatOptions']
    conditional_rendering: typing.Optional['ConditionalRenderingGroupKind']

    def __init__(self, element: typing.Optional['ElementReference'] = None, repeat: typing.Optional['AutoGridRepeatOptions'] = None, conditional_rendering: typing.Optional['ConditionalRenderingGroupKind'] = None) -> None:
        self.element = element if element is not None else ElementReference()
        self.repeat = repeat
        self.conditional_rendering = conditional_rendering

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "element": self.element,
        }
        if self.repeat is not None:
            payload["repeat"] = self.repeat
        if self.conditional_rendering is not None:
            payload["conditionalRendering"] = self.conditional_rendering
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "element" in data:
            args["element"] = ElementReference.from_json(data["element"])
        if "repeat" in data:
            args["repeat"] = AutoGridRepeatOptions.from_json(data["repeat"])
        if "conditionalRendering" in data:
            args["conditional_rendering"] = ConditionalRenderingGroupKind.from_json(data["conditionalRendering"])        

        return cls(**args)


class AutoGridRepeatOptions:
    mode: str
    value: str

    def __init__(self, value: str = "") -> None:
        self.mode = RepeatMode.VARIABLE
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class TabsLayoutKind:
    kind: typing.Literal["TabsLayout"]
    spec: 'TabsLayoutSpec'

    def __init__(self, spec: typing.Optional['TabsLayoutSpec'] = None) -> None:
        self.kind = "TabsLayout"
        self.spec = spec if spec is not None else TabsLayoutSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = TabsLayoutSpec.from_json(data["spec"])        

        return cls(**args)


class TabsLayoutSpec:
    tabs: list['TabsLayoutTabKind']

    def __init__(self, tabs: typing.Optional[list['TabsLayoutTabKind']] = None) -> None:
        self.tabs = tabs if tabs is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "tabs": self.tabs,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "tabs" in data:
            args["tabs"] = [TabsLayoutTabKind.from_json(item) for item in data["tabs"]]        

        return cls(**args)


class TabsLayoutTabKind:
    kind: typing.Literal["TabsLayoutTab"]
    spec: 'TabsLayoutTabSpec'

    def __init__(self, spec: typing.Optional['TabsLayoutTabSpec'] = None) -> None:
        self.kind = "TabsLayoutTab"
        self.spec = spec if spec is not None else TabsLayoutTabSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = TabsLayoutTabSpec.from_json(data["spec"])        

        return cls(**args)


class TabsLayoutTabSpec:
    title: typing.Optional[str]
    layout: typing.Union['GridLayoutKind', 'RowsLayoutKind', 'AutoGridLayoutKind', 'TabsLayoutKind']
    conditional_rendering: typing.Optional['ConditionalRenderingGroupKind']
    repeat: typing.Optional['TabRepeatOptions']

    def __init__(self, title: typing.Optional[str] = None, layout: typing.Optional[typing.Union['GridLayoutKind', 'RowsLayoutKind', 'AutoGridLayoutKind', 'TabsLayoutKind']] = None, conditional_rendering: typing.Optional['ConditionalRenderingGroupKind'] = None, repeat: typing.Optional['TabRepeatOptions'] = None) -> None:
        self.title = title
        self.layout = layout if layout is not None else GridLayoutKind()
        self.conditional_rendering = conditional_rendering
        self.repeat = repeat

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "layout": self.layout,
        }
        if self.title is not None:
            payload["title"] = self.title
        if self.conditional_rendering is not None:
            payload["conditionalRendering"] = self.conditional_rendering
        if self.repeat is not None:
            payload["repeat"] = self.repeat
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "layout" in data:
            decoding_map_layout_union: dict[str, typing.Union[typing.Type[AutoGridLayoutKind], typing.Type[GridLayoutKind], typing.Type[RowsLayoutKind], typing.Type[TabsLayoutKind]]] = {"AutoGridLayout": AutoGridLayoutKind, "GridLayout": GridLayoutKind, "RowsLayout": RowsLayoutKind, "TabsLayout": TabsLayoutKind}
            args["layout"] = decoding_map_layout_union[data["layout"]["kind"]].from_json(data["layout"])
        if "conditionalRendering" in data:
            args["conditional_rendering"] = ConditionalRenderingGroupKind.from_json(data["conditionalRendering"])
        if "repeat" in data:
            args["repeat"] = TabRepeatOptions.from_json(data["repeat"])        

        return cls(**args)


class TabRepeatOptions:
    mode: str
    value: str

    def __init__(self, value: str = "") -> None:
        self.mode = RepeatMode.VARIABLE
        self.value = value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class DashboardLink:
    """
    Links with references to other dashboards or external resources
    """

    # Title to display with the link
    title: str
    # Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
    # FIXME: The type is generated as `type: DashboardLinkType | dashboardLinkType.Link;` but it should be `type: DashboardLinkType`
    type_val: 'DashboardLinkType'
    # Icon name to be displayed with the link
    icon: str
    # Tooltip to display when the user hovers their mouse over it
    tooltip: str
    # Link URL. Only required/valid if the type is link
    url: typing.Optional[str]
    # List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
    tags: list[str]
    # If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
    as_dropdown: bool
    # If true, the link will be opened in a new tab
    target_blank: bool
    # If true, includes current template variables values in the link as query params
    include_vars: bool
    # If true, includes current time range in the link as query params
    keep_time: bool
    # Placement can be used to display the link somewhere else on the dashboard other than above the visualisations.
    placement: str

    def __init__(self, title: str = "", type_val: typing.Optional['DashboardLinkType'] = None, icon: str = "", tooltip: str = "", url: typing.Optional[str] = None, tags: typing.Optional[list[str]] = None, as_dropdown: bool = False, target_blank: bool = False, include_vars: bool = False, keep_time: bool = False) -> None:
        self.title = title
        self.type_val = type_val if type_val is not None else DashboardLinkType.LINK
        self.icon = icon
        self.tooltip = tooltip
        self.url = url
        self.tags = tags if tags is not None else []
        self.as_dropdown = as_dropdown
        self.target_blank = target_blank
        self.include_vars = include_vars
        self.keep_time = keep_time
        self.placement = DashboardLinkPlacement

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "title": self.title,
            "type": self.type_val,
            "icon": self.icon,
            "tooltip": self.tooltip,
            "tags": self.tags,
            "asDropdown": self.as_dropdown,
            "targetBlank": self.target_blank,
            "includeVars": self.include_vars,
            "keepTime": self.keep_time,
        }
        if self.url is not None:
            payload["url"] = self.url
        if self.placement is not None:
            payload["placement"] = self.placement
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "title" in data:
            args["title"] = data["title"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "icon" in data:
            args["icon"] = data["icon"]
        if "tooltip" in data:
            args["tooltip"] = data["tooltip"]
        if "url" in data:
            args["url"] = data["url"]
        if "tags" in data:
            args["tags"] = data["tags"]
        if "asDropdown" in data:
            args["as_dropdown"] = data["asDropdown"]
        if "targetBlank" in data:
            args["target_blank"] = data["targetBlank"]
        if "includeVars" in data:
            args["include_vars"] = data["includeVars"]
        if "keepTime" in data:
            args["keep_time"] = data["keepTime"]        

        return cls(**args)


class DashboardLinkType(enum.StrEnum):
    """
    Dashboard Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
    """

    LINK = "link"
    DASHBOARDS = "dashboards"


# Dashboard Link placement. Defines where the link should be displayed.
# - "inControlsMenu" renders the link in bottom part of the dashboard controls dropdown menu
DashboardLinkPlacement: typing.Literal["inControlsMenu"] = "inControlsMenu"


class TimeSettingsSpec:
    """
    Time configuration
    It defines the default time config for the time picker, the refresh picker for the specific dashboard.
    """

    # Timezone of dashboard. Accepted values are IANA TZDB zone ID or "browser" or "utc".
    timezone: typing.Optional[str]
    # Start time range for dashboard.
    # Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    from_val: str
    # End time range for dashboard.
    # Accepted values are relative time strings like "now-6h" or absolute time strings like "2020-07-10T08:00:00.000Z".
    to: str
    # Refresh rate of dashboard. Represented via interval string, e.g. "5s", "1m", "1h", "1d".
    # v1: refresh
    auto_refresh: str
    # Interval options available in the refresh picker dropdown.
    # v1: timepicker.refresh_intervals
    auto_refresh_intervals: list[str]
    # Selectable options available in the time picker dropdown. Has no effect on provisioned dashboard.
    # v1: timepicker.quick_ranges , not exposed in the UI
    quick_ranges: typing.Optional[list['TimeRangeOption']]
    # Whether timepicker is visible or not.
    # v1: timepicker.hidden
    hide_timepicker: bool
    # Day when the week starts. Expressed by the name of the day in lowercase, e.g. "monday".
    week_start: typing.Optional[typing.Literal["saturday", "monday", "sunday"]]
    # The month that the fiscal year starts on. 0 = January, 11 = December
    fiscal_year_start_month: int
    # Override the now time by entering a time delay. Use this option to accommodate known delays in data aggregation to avoid null values.
    # v1: timepicker.nowDelay
    now_delay: typing.Optional[str]

    def __init__(self, timezone: typing.Optional[str] = "browser", from_val: str = "now-6h", to: str = "now", auto_refresh: str = "", auto_refresh_intervals: typing.Optional[list[str]] = None, quick_ranges: typing.Optional[list['TimeRangeOption']] = None, hide_timepicker: bool = False, week_start: typing.Optional[typing.Literal["saturday", "monday", "sunday"]] = None, fiscal_year_start_month: int = 0, now_delay: typing.Optional[str] = None) -> None:
        self.timezone = timezone
        self.from_val = from_val
        self.to = to
        self.auto_refresh = auto_refresh
        self.auto_refresh_intervals = auto_refresh_intervals if auto_refresh_intervals is not None else ["5s", "10s", "30s", "1m", "5m", "15m", "30m", "1h", "2h", "1d"]
        self.quick_ranges = quick_ranges
        self.hide_timepicker = hide_timepicker
        self.week_start = week_start
        self.fiscal_year_start_month = fiscal_year_start_month
        self.now_delay = now_delay

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "from": self.from_val,
            "to": self.to,
            "autoRefresh": self.auto_refresh,
            "autoRefreshIntervals": self.auto_refresh_intervals,
            "hideTimepicker": self.hide_timepicker,
            "fiscalYearStartMonth": self.fiscal_year_start_month,
        }
        if self.timezone is not None:
            payload["timezone"] = self.timezone
        if self.quick_ranges is not None:
            payload["quickRanges"] = self.quick_ranges
        if self.week_start is not None:
            payload["weekStart"] = self.week_start
        if self.now_delay is not None:
            payload["nowDelay"] = self.now_delay
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "timezone" in data:
            args["timezone"] = data["timezone"]
        if "from" in data:
            args["from_val"] = data["from"]
        if "to" in data:
            args["to"] = data["to"]
        if "autoRefresh" in data:
            args["auto_refresh"] = data["autoRefresh"]
        if "autoRefreshIntervals" in data:
            args["auto_refresh_intervals"] = data["autoRefreshIntervals"]
        if "quickRanges" in data:
            args["quick_ranges"] = [TimeRangeOption.from_json(item) for item in data["quickRanges"]]
        if "hideTimepicker" in data:
            args["hide_timepicker"] = data["hideTimepicker"]
        if "weekStart" in data:
            args["week_start"] = data["weekStart"]
        if "fiscalYearStartMonth" in data:
            args["fiscal_year_start_month"] = data["fiscalYearStartMonth"]
        if "nowDelay" in data:
            args["now_delay"] = data["nowDelay"]        

        return cls(**args)


class TimeRangeOption:
    display: str
    from_val: str
    to: str

    def __init__(self, display: str = "Last 6 hours", from_val: str = "now-6h", to: str = "now") -> None:
        self.display = display
        self.from_val = from_val
        self.to = to

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "display": self.display,
            "from": self.from_val,
            "to": self.to,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "display" in data:
            args["display"] = data["display"]
        if "from" in data:
            args["from_val"] = data["from"]
        if "to" in data:
            args["to"] = data["to"]        

        return cls(**args)


VariableKind: typing.TypeAlias = typing.Union['QueryVariableKind', 'TextVariableKind', 'ConstantVariableKind', 'DatasourceVariableKind', 'IntervalVariableKind', 'CustomVariableKind', 'GroupByVariableKind', 'AdhocVariableKind', 'SwitchVariableKind']


class QueryVariableKind:
    """
    Query variable kind
    """

    kind: typing.Literal["QueryVariable"]
    spec: 'QueryVariableSpec'

    def __init__(self, spec: typing.Optional['QueryVariableSpec'] = None) -> None:
        self.kind = "QueryVariable"
        self.spec = spec if spec is not None else QueryVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = QueryVariableSpec.from_json(data["spec"])        

        return cls(**args)


class QueryVariableSpec:
    """
    Query variable specification
    """

    name: str
    current: 'VariableOption'
    label: typing.Optional[str]
    hide: 'VariableHide'
    refresh: 'VariableRefresh'
    skip_url_sync: bool
    description: typing.Optional[str]
    query: 'DataQueryKind'
    regex: str
    regex_apply_to: typing.Optional['VariableRegexApplyTo']
    sort: 'VariableSort'
    definition: typing.Optional[str]
    options: list['VariableOption']
    multi: bool
    include_all: bool
    all_value: typing.Optional[str]
    placeholder: typing.Optional[str]
    allow_custom_value: bool
    static_options: typing.Optional[list['VariableOption']]
    static_options_order: typing.Optional[typing.Literal["before", "after", "sorted"]]

    def __init__(self, name: str = "", current: typing.Optional['VariableOption'] = None, label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, refresh: typing.Optional['VariableRefresh'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None, query: typing.Optional['DataQueryKind'] = None, regex: str = "", regex_apply_to: typing.Optional['VariableRegexApplyTo'] = None, sort: typing.Optional['VariableSort'] = None, definition: typing.Optional[str] = None, options: typing.Optional[list['VariableOption']] = None, multi: bool = False, include_all: bool = False, all_value: typing.Optional[str] = None, placeholder: typing.Optional[str] = None, allow_custom_value: bool = True, static_options: typing.Optional[list['VariableOption']] = None, static_options_order: typing.Optional[typing.Literal["before", "after", "sorted"]] = None) -> None:
        self.name = name
        self.current = current if current is not None else VariableOption(text="", value="")
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.refresh = refresh if refresh is not None else VariableRefresh.NEVER
        self.skip_url_sync = skip_url_sync
        self.description = description
        self.query = query if query is not None else DataQueryKind()
        self.regex = regex
        self.regex_apply_to = regex_apply_to if regex_apply_to is not None else VariableRegexApplyTo.VALUE
        self.sort = sort if sort is not None else VariableSort.DISABLED
        self.definition = definition
        self.options = options if options is not None else []
        self.multi = multi
        self.include_all = include_all
        self.all_value = all_value
        self.placeholder = placeholder
        self.allow_custom_value = allow_custom_value
        self.static_options = static_options
        self.static_options_order = static_options_order

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "current": self.current,
            "hide": self.hide,
            "refresh": self.refresh,
            "skipUrlSync": self.skip_url_sync,
            "query": self.query,
            "regex": self.regex,
            "sort": self.sort,
            "options": self.options,
            "multi": self.multi,
            "includeAll": self.include_all,
            "allowCustomValue": self.allow_custom_value,
        }
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        if self.regex_apply_to is not None:
            payload["regexApplyTo"] = self.regex_apply_to
        if self.definition is not None:
            payload["definition"] = self.definition
        if self.all_value is not None:
            payload["allValue"] = self.all_value
        if self.placeholder is not None:
            payload["placeholder"] = self.placeholder
        if self.static_options is not None:
            payload["staticOptions"] = self.static_options
        if self.static_options_order is not None:
            payload["staticOptionsOrder"] = self.static_options_order
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "current" in data:
            args["current"] = VariableOption.from_json(data["current"])
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "refresh" in data:
            args["refresh"] = data["refresh"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]
        if "query" in data:
            args["query"] = DataQueryKind.from_json(data["query"])
        if "regex" in data:
            args["regex"] = data["regex"]
        if "regexApplyTo" in data:
            args["regex_apply_to"] = data["regexApplyTo"]
        if "sort" in data:
            args["sort"] = data["sort"]
        if "definition" in data:
            args["definition"] = data["definition"]
        if "options" in data:
            args["options"] = [VariableOption.from_json(item) for item in data["options"]]
        if "multi" in data:
            args["multi"] = data["multi"]
        if "includeAll" in data:
            args["include_all"] = data["includeAll"]
        if "allValue" in data:
            args["all_value"] = data["allValue"]
        if "placeholder" in data:
            args["placeholder"] = data["placeholder"]
        if "allowCustomValue" in data:
            args["allow_custom_value"] = data["allowCustomValue"]
        if "staticOptions" in data:
            args["static_options"] = [VariableOption.from_json(item) for item in data["staticOptions"]]
        if "staticOptionsOrder" in data:
            args["static_options_order"] = data["staticOptionsOrder"]        

        return cls(**args)


class VariableOption:
    """
    Variable option specification
    """

    # Whether the option is selected or not
    selected: typing.Optional[bool]
    # Text to be displayed for the option
    text: typing.Union[str, list[str]]
    # Value of the option
    value: typing.Union[str, list[str]]

    def __init__(self, selected: typing.Optional[bool] = None, text: typing.Optional[typing.Union[str, list[str]]] = None, value: typing.Optional[typing.Union[str, list[str]]] = None) -> None:
        self.selected = selected
        self.text = text if text is not None else ""
        self.value = value if value is not None else ""

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "text": self.text,
            "value": self.value,
        }
        if self.selected is not None:
            payload["selected"] = self.selected
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "selected" in data:
            args["selected"] = data["selected"]
        if "text" in data:
            args["text"] = data["text"]
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class VariableHide(enum.StrEnum):
    """
    Determine if the variable shows on dashboard
    Accepted values are `dontHide` (show label and value), `hideLabel` (show value only), `hideVariable` (show nothing), `inControlsMenu` (show in a drop-down menu).
    """

    DONT_HIDE = "dontHide"
    HIDE_LABEL = "hideLabel"
    HIDE_VARIABLE = "hideVariable"
    IN_CONTROLS_MENU = "inControlsMenu"


class VariableRefresh(enum.StrEnum):
    """
    Options to config when to refresh a variable
    `never`: Never refresh the variable
    `onDashboardLoad`: Queries the data source every time the dashboard loads.
    `onTimeRangeChanged`: Queries the data source when the dashboard time range changes.
    """

    NEVER = "never"
    ON_DASHBOARD_LOAD = "onDashboardLoad"
    ON_TIME_RANGE_CHANGED = "onTimeRangeChanged"


class VariableRegexApplyTo(enum.StrEnum):
    """
    Determine whether regex applies to variable value or display text
    Accepted values are `value` (apply to value used in queries) or `text` (apply to display text shown to users)
    """

    VALUE = "value"
    TEXT = "text"


class VariableSort(enum.StrEnum):
    """
    Sort variable options
    Accepted values are:
    `disabled`: No sorting
    `alphabeticalAsc`: Alphabetical ASC
    `alphabeticalDesc`: Alphabetical DESC
    `numericalAsc`: Numerical ASC
    `numericalDesc`: Numerical DESC
    `alphabeticalCaseInsensitiveAsc`: Alphabetical Case Insensitive ASC
    `alphabeticalCaseInsensitiveDesc`: Alphabetical Case Insensitive DESC
    `naturalAsc`: Natural ASC
    `naturalDesc`: Natural DESC
    VariableSort enum with default value
    """

    DISABLED = "disabled"
    ALPHABETICAL_ASC = "alphabeticalAsc"
    ALPHABETICAL_DESC = "alphabeticalDesc"
    NUMERICAL_ASC = "numericalAsc"
    NUMERICAL_DESC = "numericalDesc"
    ALPHABETICAL_CASE_INSENSITIVE_ASC = "alphabeticalCaseInsensitiveAsc"
    ALPHABETICAL_CASE_INSENSITIVE_DESC = "alphabeticalCaseInsensitiveDesc"
    NATURAL_ASC = "naturalAsc"
    NATURAL_DESC = "naturalDesc"


class TextVariableKind:
    """
    Text variable kind
    """

    kind: typing.Literal["TextVariable"]
    spec: 'TextVariableSpec'

    def __init__(self, spec: typing.Optional['TextVariableSpec'] = None) -> None:
        self.kind = "TextVariable"
        self.spec = spec if spec is not None else TextVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = TextVariableSpec.from_json(data["spec"])        

        return cls(**args)


class TextVariableSpec:
    """
    Text variable specification
    """

    name: str
    current: 'VariableOption'
    query: str
    label: typing.Optional[str]
    hide: 'VariableHide'
    skip_url_sync: bool
    description: typing.Optional[str]

    def __init__(self, name: str = "", current: typing.Optional['VariableOption'] = None, query: str = "", label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None) -> None:
        self.name = name
        self.current = current if current is not None else VariableOption(text="", value="")
        self.query = query
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.skip_url_sync = skip_url_sync
        self.description = description

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "current": self.current,
            "query": self.query,
            "hide": self.hide,
            "skipUrlSync": self.skip_url_sync,
        }
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "current" in data:
            args["current"] = VariableOption.from_json(data["current"])
        if "query" in data:
            args["query"] = data["query"]
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]        

        return cls(**args)


class ConstantVariableKind:
    """
    Constant variable kind
    """

    kind: typing.Literal["ConstantVariable"]
    spec: 'ConstantVariableSpec'

    def __init__(self, spec: typing.Optional['ConstantVariableSpec'] = None) -> None:
        self.kind = "ConstantVariable"
        self.spec = spec if spec is not None else ConstantVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = ConstantVariableSpec.from_json(data["spec"])        

        return cls(**args)


class ConstantVariableSpec:
    """
    Constant variable specification
    """

    name: str
    query: str
    current: 'VariableOption'
    label: typing.Optional[str]
    hide: 'VariableHide'
    skip_url_sync: bool
    description: typing.Optional[str]

    def __init__(self, name: str = "", query: str = "", current: typing.Optional['VariableOption'] = None, label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None) -> None:
        self.name = name
        self.query = query
        self.current = current if current is not None else VariableOption(text="", value="")
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.skip_url_sync = skip_url_sync
        self.description = description

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "query": self.query,
            "current": self.current,
            "hide": self.hide,
            "skipUrlSync": self.skip_url_sync,
        }
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "query" in data:
            args["query"] = data["query"]
        if "current" in data:
            args["current"] = VariableOption.from_json(data["current"])
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]        

        return cls(**args)


class DatasourceVariableKind:
    """
    Datasource variable kind
    """

    kind: typing.Literal["DatasourceVariable"]
    spec: 'DatasourceVariableSpec'

    def __init__(self, spec: typing.Optional['DatasourceVariableSpec'] = None) -> None:
        self.kind = "DatasourceVariable"
        self.spec = spec if spec is not None else DatasourceVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = DatasourceVariableSpec.from_json(data["spec"])        

        return cls(**args)


class DatasourceVariableSpec:
    """
    Datasource variable specification
    """

    name: str
    plugin_id: str
    refresh: 'VariableRefresh'
    regex: str
    current: 'VariableOption'
    options: list['VariableOption']
    multi: bool
    include_all: bool
    all_value: typing.Optional[str]
    label: typing.Optional[str]
    hide: 'VariableHide'
    skip_url_sync: bool
    description: typing.Optional[str]
    allow_custom_value: bool

    def __init__(self, name: str = "", plugin_id: str = "", refresh: typing.Optional['VariableRefresh'] = None, regex: str = "", current: typing.Optional['VariableOption'] = None, options: typing.Optional[list['VariableOption']] = None, multi: bool = False, include_all: bool = False, all_value: typing.Optional[str] = None, label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None, allow_custom_value: bool = True) -> None:
        self.name = name
        self.plugin_id = plugin_id
        self.refresh = refresh if refresh is not None else VariableRefresh.NEVER
        self.regex = regex
        self.current = current if current is not None else VariableOption(text="", value="")
        self.options = options if options is not None else []
        self.multi = multi
        self.include_all = include_all
        self.all_value = all_value
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.skip_url_sync = skip_url_sync
        self.description = description
        self.allow_custom_value = allow_custom_value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "pluginId": self.plugin_id,
            "refresh": self.refresh,
            "regex": self.regex,
            "current": self.current,
            "options": self.options,
            "multi": self.multi,
            "includeAll": self.include_all,
            "hide": self.hide,
            "skipUrlSync": self.skip_url_sync,
            "allowCustomValue": self.allow_custom_value,
        }
        if self.all_value is not None:
            payload["allValue"] = self.all_value
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "pluginId" in data:
            args["plugin_id"] = data["pluginId"]
        if "refresh" in data:
            args["refresh"] = data["refresh"]
        if "regex" in data:
            args["regex"] = data["regex"]
        if "current" in data:
            args["current"] = VariableOption.from_json(data["current"])
        if "options" in data:
            args["options"] = [VariableOption.from_json(item) for item in data["options"]]
        if "multi" in data:
            args["multi"] = data["multi"]
        if "includeAll" in data:
            args["include_all"] = data["includeAll"]
        if "allValue" in data:
            args["all_value"] = data["allValue"]
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]
        if "allowCustomValue" in data:
            args["allow_custom_value"] = data["allowCustomValue"]        

        return cls(**args)


class IntervalVariableKind:
    """
    Interval variable kind
    """

    kind: typing.Literal["IntervalVariable"]
    spec: 'IntervalVariableSpec'

    def __init__(self, spec: typing.Optional['IntervalVariableSpec'] = None) -> None:
        self.kind = "IntervalVariable"
        self.spec = spec if spec is not None else IntervalVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = IntervalVariableSpec.from_json(data["spec"])        

        return cls(**args)


class IntervalVariableSpec:
    """
    Interval variable specification
    """

    name: str
    query: str
    current: 'VariableOption'
    options: list['VariableOption']
    auto: bool
    auto_min: str
    auto_count: int
    refresh: typing.Literal["onTimeRangeChanged"]
    label: typing.Optional[str]
    hide: 'VariableHide'
    skip_url_sync: bool
    description: typing.Optional[str]

    def __init__(self, name: str = "", query: str = "", current: typing.Optional['VariableOption'] = None, options: typing.Optional[list['VariableOption']] = None, auto: bool = False, auto_min: str = "", auto_count: int = 0, label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None) -> None:
        self.name = name
        self.query = query
        self.current = current if current is not None else VariableOption(text="", value="")
        self.options = options if options is not None else []
        self.auto = auto
        self.auto_min = auto_min
        self.auto_count = auto_count
        self.refresh = "onTimeRangeChanged"
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.skip_url_sync = skip_url_sync
        self.description = description

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "query": self.query,
            "current": self.current,
            "options": self.options,
            "auto": self.auto,
            "auto_min": self.auto_min,
            "auto_count": self.auto_count,
            "refresh": self.refresh,
            "hide": self.hide,
            "skipUrlSync": self.skip_url_sync,
        }
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "query" in data:
            args["query"] = data["query"]
        if "current" in data:
            args["current"] = VariableOption.from_json(data["current"])
        if "options" in data:
            args["options"] = [VariableOption.from_json(item) for item in data["options"]]
        if "auto" in data:
            args["auto"] = data["auto"]
        if "auto_min" in data:
            args["auto_min"] = data["auto_min"]
        if "auto_count" in data:
            args["auto_count"] = data["auto_count"]
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]        

        return cls(**args)


class CustomVariableKind:
    """
    Custom variable kind
    """

    kind: typing.Literal["CustomVariable"]
    spec: 'CustomVariableSpec'

    def __init__(self, spec: typing.Optional['CustomVariableSpec'] = None) -> None:
        self.kind = "CustomVariable"
        self.spec = spec if spec is not None else CustomVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = CustomVariableSpec.from_json(data["spec"])        

        return cls(**args)


class CustomVariableSpec:
    """
    Custom variable specification
    """

    name: str
    query: str
    current: 'VariableOption'
    options: list['VariableOption']
    multi: bool
    include_all: bool
    all_value: typing.Optional[str]
    label: typing.Optional[str]
    hide: 'VariableHide'
    skip_url_sync: bool
    description: typing.Optional[str]
    allow_custom_value: bool
    values_format: typing.Optional[typing.Literal["csv", "json"]]

    def __init__(self, name: str = "", query: str = "", current: typing.Optional['VariableOption'] = None, options: typing.Optional[list['VariableOption']] = None, multi: bool = False, include_all: bool = False, all_value: typing.Optional[str] = None, label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None, allow_custom_value: bool = True, values_format: typing.Optional[typing.Literal["csv", "json"]] = None) -> None:
        self.name = name
        self.query = query
        self.current = current if current is not None else VariableOption()
        self.options = options if options is not None else []
        self.multi = multi
        self.include_all = include_all
        self.all_value = all_value
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.skip_url_sync = skip_url_sync
        self.description = description
        self.allow_custom_value = allow_custom_value
        self.values_format = values_format

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "query": self.query,
            "current": self.current,
            "options": self.options,
            "multi": self.multi,
            "includeAll": self.include_all,
            "hide": self.hide,
            "skipUrlSync": self.skip_url_sync,
            "allowCustomValue": self.allow_custom_value,
        }
        if self.all_value is not None:
            payload["allValue"] = self.all_value
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        if self.values_format is not None:
            payload["valuesFormat"] = self.values_format
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "query" in data:
            args["query"] = data["query"]
        if "current" in data:
            args["current"] = VariableOption.from_json(data["current"])
        if "options" in data:
            args["options"] = [VariableOption.from_json(item) for item in data["options"]]
        if "multi" in data:
            args["multi"] = data["multi"]
        if "includeAll" in data:
            args["include_all"] = data["includeAll"]
        if "allValue" in data:
            args["all_value"] = data["allValue"]
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]
        if "allowCustomValue" in data:
            args["allow_custom_value"] = data["allowCustomValue"]
        if "valuesFormat" in data:
            args["values_format"] = data["valuesFormat"]        

        return cls(**args)


class GroupByVariableKind:
    """
    Group variable kind
    """

    kind: typing.Literal["GroupByVariable"]
    group: str
    datasource: typing.Optional['Dashboardv2beta1GroupByVariableKindDatasource']
    spec: 'GroupByVariableSpec'

    def __init__(self, group: str = "", datasource: typing.Optional['Dashboardv2beta1GroupByVariableKindDatasource'] = None, spec: typing.Optional['GroupByVariableSpec'] = None) -> None:
        self.kind = "GroupByVariable"
        self.group = group
        self.datasource = datasource
        self.spec = spec if spec is not None else GroupByVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "group": self.group,
            "spec": self.spec,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "group" in data:
            args["group"] = data["group"]
        if "datasource" in data:
            args["datasource"] = Dashboardv2beta1GroupByVariableKindDatasource.from_json(data["datasource"])
        if "spec" in data:
            args["spec"] = GroupByVariableSpec.from_json(data["spec"])        

        return cls(**args)


class GroupByVariableSpec:
    """
    GroupBy variable specification
    """

    name: str
    default_value: typing.Optional['VariableOption']
    current: 'VariableOption'
    options: list['VariableOption']
    multi: bool
    label: typing.Optional[str]
    hide: 'VariableHide'
    skip_url_sync: bool
    description: typing.Optional[str]

    def __init__(self, name: str = "", default_value: typing.Optional['VariableOption'] = None, current: typing.Optional['VariableOption'] = None, options: typing.Optional[list['VariableOption']] = None, multi: bool = False, label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None) -> None:
        self.name = name
        self.default_value = default_value
        self.current = current if current is not None else VariableOption(text="", value="")
        self.options = options if options is not None else []
        self.multi = multi
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.skip_url_sync = skip_url_sync
        self.description = description

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "current": self.current,
            "options": self.options,
            "multi": self.multi,
            "hide": self.hide,
            "skipUrlSync": self.skip_url_sync,
        }
        if self.default_value is not None:
            payload["defaultValue"] = self.default_value
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "defaultValue" in data:
            args["default_value"] = VariableOption.from_json(data["defaultValue"])
        if "current" in data:
            args["current"] = VariableOption.from_json(data["current"])
        if "options" in data:
            args["options"] = [VariableOption.from_json(item) for item in data["options"]]
        if "multi" in data:
            args["multi"] = data["multi"]
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]        

        return cls(**args)


class AdhocVariableKind:
    """
    Adhoc variable kind
    """

    kind: typing.Literal["AdhocVariable"]
    group: str
    datasource: typing.Optional['Dashboardv2beta1AdhocVariableKindDatasource']
    spec: 'AdhocVariableSpec'

    def __init__(self, group: str = "", datasource: typing.Optional['Dashboardv2beta1AdhocVariableKindDatasource'] = None, spec: typing.Optional['AdhocVariableSpec'] = None) -> None:
        self.kind = "AdhocVariable"
        self.group = group
        self.datasource = datasource
        self.spec = spec if spec is not None else AdhocVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "group": self.group,
            "spec": self.spec,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "group" in data:
            args["group"] = data["group"]
        if "datasource" in data:
            args["datasource"] = Dashboardv2beta1AdhocVariableKindDatasource.from_json(data["datasource"])
        if "spec" in data:
            args["spec"] = AdhocVariableSpec.from_json(data["spec"])        

        return cls(**args)


class AdhocVariableSpec:
    """
    Adhoc variable specification
    """

    name: str
    base_filters: list['AdHocFilterWithLabels']
    filters: list['AdHocFilterWithLabels']
    default_keys: list['MetricFindValue']
    label: typing.Optional[str]
    hide: 'VariableHide'
    skip_url_sync: bool
    description: typing.Optional[str]
    allow_custom_value: bool

    def __init__(self, name: str = "", base_filters: typing.Optional[list['AdHocFilterWithLabels']] = None, filters: typing.Optional[list['AdHocFilterWithLabels']] = None, default_keys: typing.Optional[list['MetricFindValue']] = None, label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None, allow_custom_value: bool = True) -> None:
        self.name = name
        self.base_filters = base_filters if base_filters is not None else []
        self.filters = filters if filters is not None else []
        self.default_keys = default_keys if default_keys is not None else []
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.skip_url_sync = skip_url_sync
        self.description = description
        self.allow_custom_value = allow_custom_value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "baseFilters": self.base_filters,
            "filters": self.filters,
            "defaultKeys": self.default_keys,
            "hide": self.hide,
            "skipUrlSync": self.skip_url_sync,
            "allowCustomValue": self.allow_custom_value,
        }
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "baseFilters" in data:
            args["base_filters"] = [AdHocFilterWithLabels.from_json(item) for item in data["baseFilters"]]
        if "filters" in data:
            args["filters"] = [AdHocFilterWithLabels.from_json(item) for item in data["filters"]]
        if "defaultKeys" in data:
            args["default_keys"] = [MetricFindValue.from_json(item) for item in data["defaultKeys"]]
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]
        if "allowCustomValue" in data:
            args["allow_custom_value"] = data["allowCustomValue"]        

        return cls(**args)


class AdHocFilterWithLabels:
    """
    Define the AdHocFilterWithLabels type
    """

    key: str
    operator: str
    value: str
    values: typing.Optional[list[str]]
    key_label: typing.Optional[str]
    value_labels: typing.Optional[list[str]]
    force_edit: typing.Optional[bool]
    origin: str
    # @deprecated
    condition: typing.Optional[str]

    def __init__(self, key: str = "", operator: str = "", value: str = "", values: typing.Optional[list[str]] = None, key_label: typing.Optional[str] = None, value_labels: typing.Optional[list[str]] = None, force_edit: typing.Optional[bool] = None, condition: typing.Optional[str] = None) -> None:
        self.key = key
        self.operator = operator
        self.value = value
        self.values = values
        self.key_label = key_label
        self.value_labels = value_labels
        self.force_edit = force_edit
        self.origin = FilterOrigin
        self.condition = condition

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "key": self.key,
            "operator": self.operator,
            "value": self.value,
        }
        if self.values is not None:
            payload["values"] = self.values
        if self.key_label is not None:
            payload["keyLabel"] = self.key_label
        if self.value_labels is not None:
            payload["valueLabels"] = self.value_labels
        if self.force_edit is not None:
            payload["forceEdit"] = self.force_edit
        if self.origin is not None:
            payload["origin"] = self.origin
        if self.condition is not None:
            payload["condition"] = self.condition
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "key" in data:
            args["key"] = data["key"]
        if "operator" in data:
            args["operator"] = data["operator"]
        if "value" in data:
            args["value"] = data["value"]
        if "values" in data:
            args["values"] = data["values"]
        if "keyLabel" in data:
            args["key_label"] = data["keyLabel"]
        if "valueLabels" in data:
            args["value_labels"] = data["valueLabels"]
        if "forceEdit" in data:
            args["force_edit"] = data["forceEdit"]
        if "condition" in data:
            args["condition"] = data["condition"]        

        return cls(**args)


# Determine the origin of the adhoc variable filter
FilterOrigin: typing.Literal["dashboard"] = "dashboard"


class MetricFindValue:
    """
    Define the MetricFindValue type
    """

    text: str
    value: typing.Optional[typing.Union[str, float]]
    group: typing.Optional[str]
    expandable: typing.Optional[bool]

    def __init__(self, text: str = "", value: typing.Optional[typing.Union[str, float]] = None, group: typing.Optional[str] = None, expandable: typing.Optional[bool] = None) -> None:
        self.text = text
        self.value = value
        self.group = group
        self.expandable = expandable

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "text": self.text,
        }
        if self.value is not None:
            payload["value"] = self.value
        if self.group is not None:
            payload["group"] = self.group
        if self.expandable is not None:
            payload["expandable"] = self.expandable
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "text" in data:
            args["text"] = data["text"]
        if "value" in data:
            args["value"] = data["value"]
        if "group" in data:
            args["group"] = data["group"]
        if "expandable" in data:
            args["expandable"] = data["expandable"]        

        return cls(**args)


class SwitchVariableKind:
    kind: typing.Literal["SwitchVariable"]
    spec: 'SwitchVariableSpec'

    def __init__(self, spec: typing.Optional['SwitchVariableSpec'] = None) -> None:
        self.kind = "SwitchVariable"
        self.spec = spec if spec is not None else SwitchVariableSpec()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "spec" in data:
            args["spec"] = SwitchVariableSpec.from_json(data["spec"])        

        return cls(**args)


class SwitchVariableSpec:
    name: str
    current: str
    enabled_value: str
    disabled_value: str
    label: typing.Optional[str]
    hide: 'VariableHide'
    skip_url_sync: bool
    description: typing.Optional[str]

    def __init__(self, name: str = "", current: str = "false", enabled_value: str = "true", disabled_value: str = "false", label: typing.Optional[str] = None, hide: typing.Optional['VariableHide'] = None, skip_url_sync: bool = False, description: typing.Optional[str] = None) -> None:
        self.name = name
        self.current = current
        self.enabled_value = enabled_value
        self.disabled_value = disabled_value
        self.label = label
        self.hide = hide if hide is not None else VariableHide.DONT_HIDE
        self.skip_url_sync = skip_url_sync
        self.description = description

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "current": self.current,
            "enabledValue": self.enabled_value,
            "disabledValue": self.disabled_value,
            "hide": self.hide,
            "skipUrlSync": self.skip_url_sync,
        }
        if self.label is not None:
            payload["label"] = self.label
        if self.description is not None:
            payload["description"] = self.description
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "current" in data:
            args["current"] = data["current"]
        if "enabledValue" in data:
            args["enabled_value"] = data["enabledValue"]
        if "disabledValue" in data:
            args["disabled_value"] = data["disabledValue"]
        if "label" in data:
            args["label"] = data["label"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "skipUrlSync" in data:
            args["skip_url_sync"] = data["skipUrlSync"]
        if "description" in data:
            args["description"] = data["description"]        

        return cls(**args)


class AnnotationEventFieldSource(enum.StrEnum):
    """
    Annotation event field source. Defines how to obtain the value for an annotation event field.
    - "field": Find the value with a matching key (default)
    - "text": Write a constant string into the value
    - "skip": Do not include the field
    """

    FIELD = "field"
    TEXT = "text"
    SKIP = "skip"


class Kind:
    """
    --- Common types ---
    """

    kind: str
    spec: object
    metadata: typing.Optional[object]

    def __init__(self, kind: str = "", spec: object = None, metadata: typing.Optional[object] = None) -> None:
        self.kind = kind
        self.spec = spec
        self.metadata = metadata

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "kind": self.kind,
            "spec": self.spec,
        }
        if self.metadata is not None:
            payload["metadata"] = self.metadata
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "kind" in data:
            args["kind"] = data["kind"]
        if "spec" in data:
            args["spec"] = data["spec"]
        if "metadata" in data:
            args["metadata"] = data["metadata"]        

        return cls(**args)


# Variable types
VariableValue: typing.TypeAlias = typing.Union[str, bool, float, 'CustomVariableValue', list['VariableValueSingle']]


VariableValueSingle: typing.TypeAlias = typing.Union[str, bool, float, 'CustomVariableValue']


class CustomVariableValue:
    """
    Custom variable value
    """

    # The format name or function used in the expression
    formatter: typing.Optional[str]

    def __init__(self, formatter: typing.Optional[str] = None) -> None:
        self.formatter = formatter

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "formatter": self.formatter,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "formatter" in data:
            args["formatter"] = data["formatter"]        

        return cls(**args)


class VariableType(enum.StrEnum):
    """
    Dashboard variable type
    `query`: Query-generated list of values such as metric names, server names, sensor IDs, data centers, and so on.
    `adhoc`: Key/value filters that are automatically added to all metric queries for a data source (Prometheus, Loki, InfluxDB, and Elasticsearch only).
    `constant`: 	Define a hidden constant.
    `datasource`: Quickly change the data source for an entire dashboard.
    `interval`: Interval variables represent time spans.
    `textbox`: Display a free text input field with an optional default value.
    `custom`: Define the variable options manually using a comma-separated list.
    `system`: Variables defined by Grafana. See: https://grafana.com/docs/grafana/latest/dashboards/variables/add-template-variables/#global-variables
    """

    QUERY = "query"
    ADHOC = "adhoc"
    GROUPBY = "groupby"
    CONSTANT = "constant"
    DATASOURCE = "datasource"
    INTERVAL = "interval"
    TEXTBOX = "textbox"
    CUSTOM = "custom"
    SYSTEM = "system"
    SNAPSHOT = "snapshot"
    SWITCH = "switch"


class CustomFormatterVariable:
    """
    Custom formatter variable
    """

    name: str
    type_val: 'VariableType'
    multi: bool
    include_all: bool

    def __init__(self, name: str = "", type_val: typing.Optional['VariableType'] = None, multi: bool = False, include_all: bool = False) -> None:
        self.name = name
        self.type_val = type_val if type_val is not None else VariableType.QUERY
        self.multi = multi
        self.include_all = include_all

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "type": self.type_val,
            "multi": self.multi,
            "includeAll": self.include_all,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "multi" in data:
            args["multi"] = data["multi"]
        if "includeAll" in data:
            args["include_all"] = data["includeAll"]        

        return cls(**args)


class VariableValueOption:
    """
    FIXME: should we introduce this? --- Variable value option
    """

    label: str
    value: 'VariableValueSingle'
    group: typing.Optional[str]

    def __init__(self, label: str = "", value: typing.Optional['VariableValueSingle'] = None, group: typing.Optional[str] = None) -> None:
        self.label = label
        self.value = value if value is not None else ""
        self.group = group

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "label": self.label,
            "value": self.value,
        }
        if self.group is not None:
            payload["group"] = self.group
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "label" in data:
            args["label"] = data["label"]
        if "value" in data:
            args["value"] = data["value"]
        if "group" in data:
            args["group"] = data["group"]        

        return cls(**args)


class Dashboardv2beta1DataQueryKindDatasource:
    name: typing.Optional[str]

    def __init__(self, name: typing.Optional[str] = None) -> None:
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.name is not None:
            payload["name"] = self.name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class Dashboardv2beta1FieldConfigSourceOverrides:
    # Describes config override rules created when interacting with Grafana.
    system_ref: typing.Optional[str]
    matcher: 'MatcherConfig'
    properties: list['DynamicConfigValue']

    def __init__(self, system_ref: typing.Optional[str] = None, matcher: typing.Optional['MatcherConfig'] = None, properties: typing.Optional[list['DynamicConfigValue']] = None) -> None:
        self.system_ref = system_ref
        self.matcher = matcher if matcher is not None else MatcherConfig()
        self.properties = properties if properties is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "matcher": self.matcher,
            "properties": self.properties,
        }
        if self.system_ref is not None:
            payload["__systemRef"] = self.system_ref
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "__systemRef" in data:
            args["system_ref"] = data["__systemRef"]
        if "matcher" in data:
            args["matcher"] = MatcherConfig.from_json(data["matcher"])
        if "properties" in data:
            args["properties"] = [DynamicConfigValue.from_json(item) for item in data["properties"]]        

        return cls(**args)


class Dashboardv2beta1RangeMapOptions:
    # Min value of the range. It can be null which means -Infinity
    from_val: typing.Optional[float]
    # Max value of the range. It can be null which means +Infinity
    to: typing.Optional[float]
    # Config to apply when the value is within the range
    result: 'ValueMappingResult'

    def __init__(self, from_val: typing.Optional[float] = None, to: typing.Optional[float] = None, result: typing.Optional['ValueMappingResult'] = None) -> None:
        self.from_val = from_val
        self.to = to
        self.result = result if result is not None else ValueMappingResult()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "from": self.from_val,
            "to": self.to,
            "result": self.result,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "from" in data:
            args["from_val"] = data["from"]
        if "to" in data:
            args["to"] = data["to"]
        if "result" in data:
            args["result"] = ValueMappingResult.from_json(data["result"])        

        return cls(**args)


class Dashboardv2beta1RegexMapOptions:
    # Regular expression to match against
    pattern: str
    # Config to apply when the value matches the regex
    result: 'ValueMappingResult'

    def __init__(self, pattern: str = "", result: typing.Optional['ValueMappingResult'] = None) -> None:
        self.pattern = pattern
        self.result = result if result is not None else ValueMappingResult()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "pattern": self.pattern,
            "result": self.result,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pattern" in data:
            args["pattern"] = data["pattern"]
        if "result" in data:
            args["result"] = ValueMappingResult.from_json(data["result"])        

        return cls(**args)


class Dashboardv2beta1SpecialValueMapOptions:
    # Special value to match against
    match: 'SpecialValueMatch'
    # Config to apply when the value matches the special value
    result: 'ValueMappingResult'

    def __init__(self, match: typing.Optional['SpecialValueMatch'] = None, result: typing.Optional['ValueMappingResult'] = None) -> None:
        self.match = match if match is not None else SpecialValueMatch.TRUE
        self.result = result if result is not None else ValueMappingResult()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "match": self.match,
            "result": self.result,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "match" in data:
            args["match"] = data["match"]
        if "result" in data:
            args["result"] = ValueMappingResult.from_json(data["result"])        

        return cls(**args)


class Dashboardv2beta1ActionStyle:
    background_color: typing.Optional[str]

    def __init__(self, background_color: typing.Optional[str] = None) -> None:
        self.background_color = background_color

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.background_color is not None:
            payload["backgroundColor"] = self.background_color
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "backgroundColor" in data:
            args["background_color"] = data["backgroundColor"]        

        return cls(**args)


class Dashboardv2beta1GroupByVariableKindDatasource:
    name: typing.Optional[str]

    def __init__(self, name: typing.Optional[str] = None) -> None:
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.name is not None:
            payload["name"] = self.name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class Dashboardv2beta1AdhocVariableKindDatasource:
    name: typing.Optional[str]

    def __init__(self, name: typing.Optional[str] = None) -> None:
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.name is not None:
            payload["name"] = self.name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)




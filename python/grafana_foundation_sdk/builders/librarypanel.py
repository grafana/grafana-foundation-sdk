# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import librarypanel
from ..cog import variants as cogvariants
from ..models import dashboard


class LibraryPanel(cogbuilder.Builder[librarypanel.LibraryPanel]):    
    _internal: librarypanel.LibraryPanel

    def __init__(self):
        self._internal = librarypanel.LibraryPanel()

    def build(self) -> librarypanel.LibraryPanel:
        return self._internal    
    
    def folder_uid(self, folder_uid: str) -> typing.Self:    
        """
        Folder UID
        """
            
        self._internal.folder_uid = folder_uid
    
        return self
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Library element UID
        """
            
        self._internal.uid = uid
    
        return self
    
    def name(self, name: str) -> typing.Self:    
        """
        Panel name (also saved in the model)
        """
            
        if not len(name) >= 1:
            raise ValueError("len(name) must be >= 1")
        self._internal.name = name
    
        return self
    
    def description(self, description: str) -> typing.Self:    
        """
        Panel description
        """
            
        self._internal.description = description
    
        return self
    
    def type_val(self, type_val: str) -> typing.Self:    
        """
        The panel type (from inside the model)
        """
            
        if not len(type_val) >= 1:
            raise ValueError("len(type_val) must be >= 1")
        self._internal.type_val = type_val
    
        return self
    
    def schema_version(self, schema_version: int) -> typing.Self:    
        """
        Dashboard version when this was saved (zero if unknown)
        """
            
        self._internal.schema_version = schema_version
    
        return self
    
    def version(self, version: int) -> typing.Self:    
        """
        panel version, incremented each time the dashboard is updated.
        """
            
        self._internal.version = version
    
        return self
    
    def model(self, model: cogbuilder.Builder[librarypanel.LibrarypanelLibraryPanelModel]) -> typing.Self:    
        """
        TODO: should be the same panel schema defined in dashboard
        Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
        """
            
        model_resource = model.build()
        self._internal.model = model_resource
    
        return self
    
    def meta(self, meta: cogbuilder.Builder[librarypanel.LibraryElementDTOMeta]) -> typing.Self:    
        """
        Object storage metadata
        """
            
        meta_resource = meta.build()
        self._internal.meta = meta_resource
    
        return self
    

class LibraryElementDTOMetaUser(cogbuilder.Builder[librarypanel.LibraryElementDTOMetaUser]):    
    _internal: librarypanel.LibraryElementDTOMetaUser

    def __init__(self):
        self._internal = librarypanel.LibraryElementDTOMetaUser()

    def build(self) -> librarypanel.LibraryElementDTOMetaUser:
        return self._internal    
    
    def id_val(self, id_val: int) -> typing.Self:        
        self._internal.id_val = id_val
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def avatar_url(self, avatar_url: str) -> typing.Self:        
        self._internal.avatar_url = avatar_url
    
        return self
    

class LibraryElementDTOMeta(cogbuilder.Builder[librarypanel.LibraryElementDTOMeta]):    
    _internal: librarypanel.LibraryElementDTOMeta

    def __init__(self):
        self._internal = librarypanel.LibraryElementDTOMeta()

    def build(self) -> librarypanel.LibraryElementDTOMeta:
        return self._internal    
    
    def folder_name(self, folder_name: str) -> typing.Self:        
        self._internal.folder_name = folder_name
    
        return self
    
    def folder_uid(self, folder_uid: str) -> typing.Self:        
        self._internal.folder_uid = folder_uid
    
        return self
    
    def connected_dashboards(self, connected_dashboards: int) -> typing.Self:        
        self._internal.connected_dashboards = connected_dashboards
    
        return self
    
    def created(self, created: str) -> typing.Self:        
        self._internal.created = created
    
        return self
    
    def updated(self, updated: str) -> typing.Self:        
        self._internal.updated = updated
    
        return self
    
    def created_by(self, created_by: cogbuilder.Builder[librarypanel.LibraryElementDTOMetaUser]) -> typing.Self:        
        created_by_resource = created_by.build()
        self._internal.created_by = created_by_resource
    
        return self
    
    def updated_by(self, updated_by: cogbuilder.Builder[librarypanel.LibraryElementDTOMetaUser]) -> typing.Self:        
        updated_by_resource = updated_by.build()
        self._internal.updated_by = updated_by_resource
    
        return self
    

class LibrarypanelLibraryPanelModel(cogbuilder.Builder[librarypanel.LibrarypanelLibraryPanelModel]):    
    _internal: librarypanel.LibrarypanelLibraryPanelModel

    def __init__(self):
        self._internal = librarypanel.LibrarypanelLibraryPanelModel()

    def build(self) -> librarypanel.LibrarypanelLibraryPanelModel:
        return self._internal    
    
    def type_val(self, type_val: str) -> typing.Self:    
        """
        The panel plugin type id. This is used to find the plugin to display the panel.
        """
            
        if not len(type_val) >= 1:
            raise ValueError("len(type_val) must be >= 1")
        self._internal.type_val = type_val
    
        return self
    
    def plugin_version(self, plugin_version: str) -> typing.Self:    
        """
        The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
        """
            
        self._internal.plugin_version = plugin_version
    
        return self
    
    def targets(self, targets: list[cogbuilder.Builder[cogvariants.Dataquery]]) -> typing.Self:    
        """
        Depends on the panel plugin. See the plugin documentation for details.
        """
            
        targets_resources = [r1.build() for r1 in targets]
        self._internal.targets = targets_resources
    
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
    
    def options(self, options: object) -> typing.Self:    
        """
        It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
        """
            
        self._internal.options = options
    
        return self
    
    def field_config(self, field_config: dashboard.FieldConfigSource) -> typing.Self:    
        """
        Field options allow you to change how the data is displayed in your visualizations.
        """
            
        self._internal.field_config = field_config
    
        return self
    
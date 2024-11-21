# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import variants as cogvariants
from ..models import dashboard
from ..cog import runtime as cogruntime


class LibraryPanel:
    # Folder UID
    folder_uid: typing.Optional[str]
    # Library element UID
    uid: str
    # Panel name (also saved in the model)
    name: str
    # Panel description
    description: typing.Optional[str]
    # The panel type (from inside the model)
    type_val: str
    # Dashboard version when this was saved (zero if unknown)
    schema_version: typing.Optional[int]
    # panel version, incremented each time the dashboard is updated.
    version: int
    # TODO: should be the same panel schema defined in dashboard
    # Typescript: Omit<Panel, 'gridPos' | 'id' | 'libraryPanel'>;
    model: 'PanelModel'
    # Object storage metadata
    meta: typing.Optional['LibraryElementDTOMeta']

    def __init__(self, folder_uid: typing.Optional[str] = None, uid: str = "", name: str = "", description: typing.Optional[str] = None, type_val: str = "", schema_version: typing.Optional[int] = None, version: int = 0, model: typing.Optional['PanelModel'] = None, meta: typing.Optional['LibraryElementDTOMeta'] = None):
        self.folder_uid = folder_uid
        self.uid = uid
        self.name = name
        self.description = description
        self.type_val = type_val
        self.schema_version = schema_version
        self.version = version
        self.model = model if model is not None else PanelModel()
        self.meta = meta

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "uid": self.uid,
            "name": self.name,
            "type": self.type_val,
            "version": self.version,
            "model": self.model,
        }
        if self.folder_uid is not None:
            payload["folderUid"] = self.folder_uid
        if self.description is not None:
            payload["description"] = self.description
        if self.schema_version is not None:
            payload["schemaVersion"] = self.schema_version
        if self.meta is not None:
            payload["meta"] = self.meta
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "folderUid" in data:
            args["folder_uid"] = data["folderUid"]
        if "uid" in data:
            args["uid"] = data["uid"]
        if "name" in data:
            args["name"] = data["name"]
        if "description" in data:
            args["description"] = data["description"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "schemaVersion" in data:
            args["schema_version"] = data["schemaVersion"]
        if "version" in data:
            args["version"] = data["version"]
        if "model" in data:
            args["model"] = PanelModel.from_json(data["model"])
        if "meta" in data:
            args["meta"] = LibraryElementDTOMeta.from_json(data["meta"])        

        return cls(**args)


class LibraryElementDTOMetaUser:
    id_val: int
    name: str
    avatar_url: str

    def __init__(self, id_val: int = 0, name: str = "", avatar_url: str = ""):
        self.id_val = id_val
        self.name = name
        self.avatar_url = avatar_url

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "name": self.name,
            "avatarUrl": self.avatar_url,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "name" in data:
            args["name"] = data["name"]
        if "avatarUrl" in data:
            args["avatar_url"] = data["avatarUrl"]        

        return cls(**args)


class LibraryElementDTOMeta:
    folder_name: str
    folder_uid: str
    connected_dashboards: int
    created: str
    updated: str
    created_by: 'LibraryElementDTOMetaUser'
    updated_by: 'LibraryElementDTOMetaUser'

    def __init__(self, folder_name: str = "", folder_uid: str = "", connected_dashboards: int = 0, created: str = "", updated: str = "", created_by: typing.Optional['LibraryElementDTOMetaUser'] = None, updated_by: typing.Optional['LibraryElementDTOMetaUser'] = None):
        self.folder_name = folder_name
        self.folder_uid = folder_uid
        self.connected_dashboards = connected_dashboards
        self.created = created
        self.updated = updated
        self.created_by = created_by if created_by is not None else LibraryElementDTOMetaUser()
        self.updated_by = updated_by if updated_by is not None else LibraryElementDTOMetaUser()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "folderName": self.folder_name,
            "folderUid": self.folder_uid,
            "connectedDashboards": self.connected_dashboards,
            "created": self.created,
            "updated": self.updated,
            "createdBy": self.created_by,
            "updatedBy": self.updated_by,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "folderName" in data:
            args["folder_name"] = data["folderName"]
        if "folderUid" in data:
            args["folder_uid"] = data["folderUid"]
        if "connectedDashboards" in data:
            args["connected_dashboards"] = data["connectedDashboards"]
        if "created" in data:
            args["created"] = data["created"]
        if "updated" in data:
            args["updated"] = data["updated"]
        if "createdBy" in data:
            args["created_by"] = LibraryElementDTOMetaUser.from_json(data["createdBy"])
        if "updatedBy" in data:
            args["updated_by"] = LibraryElementDTOMetaUser.from_json(data["updatedBy"])        

        return cls(**args)


class PanelModel:
    """
    Dashboard panels are the basic visualization building blocks.
    """

    # The panel plugin type id. This is used to find the plugin to display the panel.
    type_val: str
    # The version of the plugin that is used for this panel. This is used to find the plugin to display the panel and to migrate old panel configs.
    plugin_version: typing.Optional[str]
    # Tags for the panel.
    tags: typing.Optional[list[str]]
    # Depends on the panel plugin. See the plugin documentation for details.
    targets: typing.Optional[list[cogvariants.Dataquery]]
    # Panel title.
    title: typing.Optional[str]
    # Panel description.
    description: typing.Optional[str]
    # Whether to display the panel without a background.
    transparent: bool
    # The datasource used in all targets.
    datasource: typing.Optional[dashboard.DataSourceRef]
    # Panel links.
    links: typing.Optional[list[dashboard.DashboardLink]]
    # Name of template variable to repeat for.
    repeat: typing.Optional[str]
    # Direction to repeat in if 'repeat' is set.
    # `h` for horizontal, `v` for vertical.
    repeat_direction: typing.Optional[typing.Literal["h", "v"]]
    # Id of the repeating panel.
    repeat_panel_id: typing.Optional[int]
    # The maximum number of data points that the panel queries are retrieving.
    max_data_points: typing.Optional[float]
    # List of transformations that are applied to the panel data before rendering.
    # When there are multiple transformations, Grafana applies them in the order they are listed.
    # Each transformation creates a result set that then passes on to the next transformation in the processing pipeline.
    transformations: list[dashboard.DataTransformerConfig]
    # The min time interval setting defines a lower limit for the $__interval and $__interval_ms variables.
    # This value must be formatted as a number followed by a valid time
    # identifier like: "40s", "3d", etc.
    # See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    interval: typing.Optional[str]
    # Overrides the relative time range for individual panels,
    # which causes them to be different than what is selected in
    # the dashboard time picker in the top-right corner of the dashboard. You can use this to show metrics from different
    # time periods or days on the same dashboard.
    # The value is formatted as time operation like: `now-5m` (Last 5 minutes), `now/d` (the day so far),
    # `now-5d/d`(Last 5 days), `now/w` (This week so far), `now-2y/y` (Last 2 years).
    # Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    # See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    time_from: typing.Optional[str]
    # Overrides the time range for individual panels by shifting its start and end relative to the time picker.
    # For example, you can shift the time range for the panel to be two hours earlier than the dashboard time picker setting `2h`.
    # Note: Panel time overrides have no effect when the dashboard’s time range is absolute.
    # See: https://grafana.com/docs/grafana/latest/panels-visualizations/query-transform-data/#query-options
    time_shift: typing.Optional[str]
    # It depends on the panel plugin. They are specified by the Options field in panel plugin schemas.
    options: object
    # Field options allow you to change how the data is displayed in your visualizations.
    field_config: dashboard.FieldConfigSource

    def __init__(self, type_val: str = "", plugin_version: typing.Optional[str] = None, tags: typing.Optional[list[str]] = None, targets: typing.Optional[list[cogvariants.Dataquery]] = None, title: typing.Optional[str] = None, description: typing.Optional[str] = None, transparent: bool = False, datasource: typing.Optional[dashboard.DataSourceRef] = None, links: typing.Optional[list[dashboard.DashboardLink]] = None, repeat: typing.Optional[str] = None, repeat_direction: typing.Optional[typing.Literal["h", "v"]] = None, repeat_panel_id: typing.Optional[int] = None, max_data_points: typing.Optional[float] = None, transformations: typing.Optional[list[dashboard.DataTransformerConfig]] = None, interval: typing.Optional[str] = None, time_from: typing.Optional[str] = None, time_shift: typing.Optional[str] = None, options: object = None, field_config: typing.Optional[dashboard.FieldConfigSource] = None):
        self.type_val = type_val
        self.plugin_version = plugin_version
        self.tags = tags
        self.targets = targets
        self.title = title
        self.description = description
        self.transparent = transparent
        self.datasource = datasource
        self.links = links
        self.repeat = repeat
        self.repeat_direction = repeat_direction if repeat_direction is not None else "h"
        self.repeat_panel_id = repeat_panel_id
        self.max_data_points = max_data_points
        self.transformations = transformations if transformations is not None else []
        self.interval = interval
        self.time_from = time_from
        self.time_shift = time_shift
        self.options = options
        self.field_config = field_config if field_config is not None else dashboard.FieldConfigSource()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "transparent": self.transparent,
            "transformations": self.transformations,
            "options": self.options,
            "fieldConfig": self.field_config,
        }
        if self.plugin_version is not None:
            payload["pluginVersion"] = self.plugin_version
        if self.tags is not None:
            payload["tags"] = self.tags
        if self.targets is not None:
            payload["targets"] = self.targets
        if self.title is not None:
            payload["title"] = self.title
        if self.description is not None:
            payload["description"] = self.description
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.links is not None:
            payload["links"] = self.links
        if self.repeat is not None:
            payload["repeat"] = self.repeat
        if self.repeat_direction is not None:
            payload["repeatDirection"] = self.repeat_direction
        if self.repeat_panel_id is not None:
            payload["repeatPanelId"] = self.repeat_panel_id
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.interval is not None:
            payload["interval"] = self.interval
        if self.time_from is not None:
            payload["timeFrom"] = self.time_from
        if self.time_shift is not None:
            payload["timeShift"] = self.time_shift
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "pluginVersion" in data:
            args["plugin_version"] = data["pluginVersion"]
        if "tags" in data:
            args["tags"] = data["tags"]
        if "targets" in data:
            args["targets"] = [cogruntime.dataquery_from_json(dataquery_json, data["datasource"]["type"] if data.get("datasource") is not None and data["datasource"].get("type", "") != "" else "") for dataquery_json in data["targets"]]
        if "title" in data:
            args["title"] = data["title"]
        if "description" in data:
            args["description"] = data["description"]
        if "transparent" in data:
            args["transparent"] = data["transparent"]
        if "datasource" in data:
            args["datasource"] = dashboard.DataSourceRef.from_json(data["datasource"])
        if "links" in data:
            args["links"] = data["links"]
        if "repeat" in data:
            args["repeat"] = data["repeat"]
        if "repeatDirection" in data:
            args["repeat_direction"] = data["repeatDirection"]
        if "repeatPanelId" in data:
            args["repeat_panel_id"] = data["repeatPanelId"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "transformations" in data:
            args["transformations"] = data["transformations"]
        if "interval" in data:
            args["interval"] = data["interval"]
        if "timeFrom" in data:
            args["time_from"] = data["timeFrom"]
        if "timeShift" in data:
            args["time_shift"] = data["timeShift"]
        if "options" in data:
            args["options"] = data["options"]
        if "fieldConfig" in data:
            args["field_config"] = dashboard.FieldConfigSource.from_json(data["fieldConfig"])        

        return cls(**args)




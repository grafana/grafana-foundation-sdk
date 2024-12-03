# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..cog import variants as cogvariants
import typing
from ..models import dashboard
from ..cog import runtime as cogruntime
import enum


class Dataquery(cogvariants.Dataquery):
    """
    Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
    """

    format_val: 'FormatOptions'
    connection_args: 'ConnectionArgs'
    table: typing.Optional[str]
    column: typing.Optional[str]
    query_id: typing.Optional[str]
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # Note this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    raw_sql: str
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]

    def __init__(self, format_val: typing.Optional['FormatOptions'] = None, connection_args: typing.Optional['ConnectionArgs'] = None, table: typing.Optional[str] = None, column: typing.Optional[str] = None, query_id: typing.Optional[str] = None, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, raw_sql: str = "", datasource: typing.Optional[dashboard.DataSourceRef] = None):
        self.format_val = format_val if format_val is not None else FormatOptions.TIME_SERIES
        self.connection_args = connection_args if connection_args is not None else ConnectionArgs()
        self.table = table
        self.column = column
        self.query_id = query_id
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.raw_sql = raw_sql
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "format": self.format_val,
            "connectionArgs": self.connection_args,
            "refId": self.ref_id,
            "rawSQL": self.raw_sql,
        }
        if self.table is not None:
            payload["table"] = self.table
        if self.column is not None:
            payload["column"] = self.column
        if self.query_id is not None:
            payload["queryID"] = self.query_id
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "format" in data:
            args["format_val"] = data["format"]
        if "connectionArgs" in data:
            args["connection_args"] = ConnectionArgs.from_json(data["connectionArgs"])
        if "table" in data:
            args["table"] = data["table"]
        if "column" in data:
            args["column"] = data["column"]
        if "queryID" in data:
            args["query_id"] = data["queryID"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "rawSQL" in data:
            args["raw_sql"] = data["rawSQL"]
        if "datasource" in data:
            args["datasource"] = dashboard.DataSourceRef.from_json(data["datasource"])        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="grafana-athena-datasource",
        from_json_hook=Dataquery.from_json,
    )


DefaultKey: typing.Literal["__default"] = "__default"


class ConnectionArgs:
    region: typing.Optional[str]
    catalog: typing.Optional[str]
    database: typing.Optional[str]
    result_reuse_enabled: typing.Optional[bool]
    result_reuse_max_age_in_minutes: typing.Optional[float]

    def __init__(self, region: typing.Optional[str] = "__default", catalog: typing.Optional[str] = "__default", database: typing.Optional[str] = "__default", result_reuse_enabled: typing.Optional[bool] = False, result_reuse_max_age_in_minutes: typing.Optional[float] = 60):
        self.region = region
        self.catalog = catalog
        self.database = database
        self.result_reuse_enabled = result_reuse_enabled
        self.result_reuse_max_age_in_minutes = result_reuse_max_age_in_minutes

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.region is not None:
            payload["region"] = self.region
        if self.catalog is not None:
            payload["catalog"] = self.catalog
        if self.database is not None:
            payload["database"] = self.database
        if self.result_reuse_enabled is not None:
            payload["resultReuseEnabled"] = self.result_reuse_enabled
        if self.result_reuse_max_age_in_minutes is not None:
            payload["resultReuseMaxAgeInMinutes"] = self.result_reuse_max_age_in_minutes
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "region" in data:
            args["region"] = data["region"]
        if "catalog" in data:
            args["catalog"] = data["catalog"]
        if "database" in data:
            args["database"] = data["database"]
        if "resultReuseEnabled" in data:
            args["result_reuse_enabled"] = data["resultReuseEnabled"]
        if "resultReuseMaxAgeInMinutes" in data:
            args["result_reuse_max_age_in_minutes"] = data["resultReuseMaxAgeInMinutes"]        

        return cls(**args)


class FormatOptions(enum.IntEnum):
    TIME_SERIES = 0
    TABLE = 1
    LOGS = 2




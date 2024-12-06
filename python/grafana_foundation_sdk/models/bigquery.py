# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing
from ..cog import variants as cogvariants
from ..models import dashboard
from ..cog import runtime as cogruntime


class QueryFormat(enum.IntEnum):
    TIMESERIES = 0
    TABLE = 1


class QueryPriority(enum.StrEnum):
    INTERACTIVE = "INTERACTIVE"
    BATCH = "BATCH"


class EditorMode(enum.StrEnum):
    CODE = "code"
    BUILDER = "builder"


class SQLExpression:
    columns: typing.Optional[list['QueryEditorFunctionExpression']]
    from_val: typing.Optional[str]
    # whereJsonTree?: _
    where_string: typing.Optional[str]
    group_by: typing.Optional[list['QueryEditorGroupByExpression']]
    order_by: typing.Optional['QueryEditorPropertyExpression']
    order_by_direction: typing.Optional['OrderByDirection']
    limit: typing.Optional[int]
    offset: typing.Optional[int]

    def __init__(self, columns: typing.Optional[list['QueryEditorFunctionExpression']] = None, from_val: typing.Optional[str] = None, where_string: typing.Optional[str] = None, group_by: typing.Optional[list['QueryEditorGroupByExpression']] = None, order_by: typing.Optional['QueryEditorPropertyExpression'] = None, order_by_direction: typing.Optional['OrderByDirection'] = None, limit: typing.Optional[int] = None, offset: typing.Optional[int] = None):
        self.columns = columns
        self.from_val = from_val
        self.where_string = where_string
        self.group_by = group_by
        self.order_by = order_by
        self.order_by_direction = order_by_direction
        self.limit = limit
        self.offset = offset

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.columns is not None:
            payload["columns"] = self.columns
        if self.from_val is not None:
            payload["from"] = self.from_val
        if self.where_string is not None:
            payload["whereString"] = self.where_string
        if self.group_by is not None:
            payload["groupBy"] = self.group_by
        if self.order_by is not None:
            payload["orderBy"] = self.order_by
        if self.order_by_direction is not None:
            payload["orderByDirection"] = self.order_by_direction
        if self.limit is not None:
            payload["limit"] = self.limit
        if self.offset is not None:
            payload["offset"] = self.offset
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "columns" in data:
            args["columns"] = data["columns"]
        if "from" in data:
            args["from_val"] = data["from"]
        if "whereString" in data:
            args["where_string"] = data["whereString"]
        if "groupBy" in data:
            args["group_by"] = data["groupBy"]
        if "orderBy" in data:
            args["order_by"] = QueryEditorPropertyExpression.from_json(data["orderBy"])
        if "orderByDirection" in data:
            args["order_by_direction"] = data["orderByDirection"]
        if "limit" in data:
            args["limit"] = data["limit"]
        if "offset" in data:
            args["offset"] = data["offset"]        

        return cls(**args)


class QueryEditorExpressionType(enum.StrEnum):
    PROPERTY = "property"
    OPERATOR = "operator"
    OR = "or"
    AND = "and"
    GROUP_BY = "groupBy"
    FUNCTION = "function"
    FUNCTION_PARAMETER = "functionParameter"


class QueryEditorFunctionExpression:
    type_val: typing.Literal["function"]
    name: typing.Optional[str]
    parameters: typing.Optional[list['QueryEditorFunctionParameterExpression']]

    def __init__(self, name: typing.Optional[str] = None, parameters: typing.Optional[list['QueryEditorFunctionParameterExpression']] = None):
        self.type_val = "function"
        self.name = name
        self.parameters = parameters

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.name is not None:
            payload["name"] = self.name
        if self.parameters is not None:
            payload["parameters"] = self.parameters
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "parameters" in data:
            args["parameters"] = data["parameters"]        

        return cls(**args)


class QueryEditorFunctionParameterExpression:
    type_val: typing.Literal["functionParameter"]
    name: typing.Optional[str]

    def __init__(self, name: typing.Optional[str] = None):
        self.type_val = "functionParameter"
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
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


class QueryEditorGroupByExpression:
    type_val: typing.Literal["groupBy"]
    property_val: 'QueryEditorProperty'

    def __init__(self, property_val: typing.Optional['QueryEditorProperty'] = None):
        self.type_val = "groupBy"
        self.property_val = property_val if property_val is not None else QueryEditorProperty()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "property": self.property_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "property" in data:
            args["property_val"] = QueryEditorProperty.from_json(data["property"])        

        return cls(**args)


class QueryEditorProperty:
    type_val: 'QueryEditorPropertyType'
    name: typing.Optional[str]

    def __init__(self, type_val: typing.Optional['QueryEditorPropertyType'] = None, name: typing.Optional[str] = None):
        self.type_val = type_val if type_val is not None else QueryEditorPropertyType.STRING
        self.name = name

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.name is not None:
            payload["name"] = self.name
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "name" in data:
            args["name"] = data["name"]        

        return cls(**args)


class QueryEditorPropertyType(enum.StrEnum):
    STRING = "string"


class QueryEditorPropertyExpression:
    type_val: typing.Literal["property"]
    property_val: 'QueryEditorProperty'

    def __init__(self, property_val: typing.Optional['QueryEditorProperty'] = None):
        self.type_val = "property"
        self.property_val = property_val if property_val is not None else QueryEditorProperty()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "property": self.property_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "property" in data:
            args["property_val"] = QueryEditorProperty.from_json(data["property"])        

        return cls(**args)


class OrderByDirection(enum.StrEnum):
    ASC = "ASC"
    DESC = "DESC"


class Dataquery(cogvariants.Dataquery):
    """
    Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
    """

    dataset: typing.Optional[str]
    table: typing.Optional[str]
    project: typing.Optional[str]
    format_val: 'QueryFormat'
    raw_query: typing.Optional[bool]
    raw_sql: str
    location: typing.Optional[str]
    partitioned: typing.Optional[bool]
    partitioned_field: typing.Optional[str]
    convert_to_utc: typing.Optional[bool]
    sharded: typing.Optional[bool]
    query_priority: typing.Optional['QueryPriority']
    time_shift: typing.Optional[str]
    editor_mode: typing.Optional['EditorMode']
    sql: typing.Optional['SQLExpression']
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[dashboard.DataSourceRef]

    def __init__(self, dataset: typing.Optional[str] = None, table: typing.Optional[str] = None, project: typing.Optional[str] = None, format_val: typing.Optional['QueryFormat'] = None, raw_query: typing.Optional[bool] = None, raw_sql: str = "", location: typing.Optional[str] = None, partitioned: typing.Optional[bool] = None, partitioned_field: typing.Optional[str] = None, convert_to_utc: typing.Optional[bool] = None, sharded: typing.Optional[bool] = None, query_priority: typing.Optional['QueryPriority'] = None, time_shift: typing.Optional[str] = None, editor_mode: typing.Optional['EditorMode'] = None, sql: typing.Optional['SQLExpression'] = None, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[dashboard.DataSourceRef] = None):
        self.dataset = dataset
        self.table = table
        self.project = project
        self.format_val = format_val if format_val is not None else QueryFormat.TIMESERIES
        self.raw_query = raw_query
        self.raw_sql = raw_sql
        self.location = location
        self.partitioned = partitioned
        self.partitioned_field = partitioned_field
        self.convert_to_utc = convert_to_utc
        self.sharded = sharded
        self.query_priority = query_priority
        self.time_shift = time_shift
        self.editor_mode = editor_mode
        self.sql = sql
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "format": self.format_val,
            "rawSql": self.raw_sql,
            "refId": self.ref_id,
        }
        if self.dataset is not None:
            payload["dataset"] = self.dataset
        if self.table is not None:
            payload["table"] = self.table
        if self.project is not None:
            payload["project"] = self.project
        if self.raw_query is not None:
            payload["rawQuery"] = self.raw_query
        if self.location is not None:
            payload["location"] = self.location
        if self.partitioned is not None:
            payload["partitioned"] = self.partitioned
        if self.partitioned_field is not None:
            payload["partitionedField"] = self.partitioned_field
        if self.convert_to_utc is not None:
            payload["convertToUTC"] = self.convert_to_utc
        if self.sharded is not None:
            payload["sharded"] = self.sharded
        if self.query_priority is not None:
            payload["queryPriority"] = self.query_priority
        if self.time_shift is not None:
            payload["timeShift"] = self.time_shift
        if self.editor_mode is not None:
            payload["editorMode"] = self.editor_mode
        if self.sql is not None:
            payload["sql"] = self.sql
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
        
        if "dataset" in data:
            args["dataset"] = data["dataset"]
        if "table" in data:
            args["table"] = data["table"]
        if "project" in data:
            args["project"] = data["project"]
        if "format" in data:
            args["format_val"] = data["format"]
        if "rawQuery" in data:
            args["raw_query"] = data["rawQuery"]
        if "rawSql" in data:
            args["raw_sql"] = data["rawSql"]
        if "location" in data:
            args["location"] = data["location"]
        if "partitioned" in data:
            args["partitioned"] = data["partitioned"]
        if "partitionedField" in data:
            args["partitioned_field"] = data["partitionedField"]
        if "convertToUTC" in data:
            args["convert_to_utc"] = data["convertToUTC"]
        if "sharded" in data:
            args["sharded"] = data["sharded"]
        if "queryPriority" in data:
            args["query_priority"] = data["queryPriority"]
        if "timeShift" in data:
            args["time_shift"] = data["timeShift"]
        if "editorMode" in data:
            args["editor_mode"] = data["editorMode"]
        if "sql" in data:
            args["sql"] = SQLExpression.from_json(data["sql"])
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "datasource" in data:
            args["datasource"] = dashboard.DataSourceRef.from_json(data["datasource"])        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="grafana-bigquery-datasource",
        from_json_hook=Dataquery.from_json,
    )




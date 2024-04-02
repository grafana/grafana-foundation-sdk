# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..cog import variants as cogvariants
import typing
from ..cog import runtime as cogruntime
import enum


class TempoQuery(cogvariants.Dataquery):
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
    # TraceQL query or trace ID
    query: typing.Optional[str]
    # @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
    search: typing.Optional[str]
    # @deprecated Query traces by service name
    service_name: typing.Optional[str]
    # @deprecated Query traces by span name
    span_name: typing.Optional[str]
    # @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
    min_duration: typing.Optional[str]
    # @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
    max_duration: typing.Optional[str]
    # Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}
    service_map_query: typing.Optional[str]
    # Use service.namespace in addition to service.name to uniquely identify a service.
    service_map_include_namespace: typing.Optional[bool]
    # Defines the maximum number of traces that are returned from Tempo
    limit: typing.Optional[int]
    # Defines the maximum number of spans per spanset that are returned from Tempo
    spss: typing.Optional[int]
    filters: list['TraceqlFilter']
    # Filters that are used to query the metrics summary
    group_by: typing.Optional[list['TraceqlFilter']]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]
    # The type of the table that is used to display the search results
    table_type: typing.Optional['SearchTableType']

    def __init__(self, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, query: typing.Optional[str] = None, search: typing.Optional[str] = None, service_name: typing.Optional[str] = None, span_name: typing.Optional[str] = None, min_duration: typing.Optional[str] = None, max_duration: typing.Optional[str] = None, service_map_query: typing.Optional[str] = None, service_map_include_namespace: typing.Optional[bool] = None, limit: typing.Optional[int] = None, spss: typing.Optional[int] = None, filters: typing.Optional[list['TraceqlFilter']] = None, group_by: typing.Optional[list['TraceqlFilter']] = None, datasource: typing.Optional[object] = None, table_type: typing.Optional['SearchTableType'] = None):
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.query = query
        self.search = search
        self.service_name = service_name
        self.span_name = span_name
        self.min_duration = min_duration
        self.max_duration = max_duration
        self.service_map_query = service_map_query
        self.service_map_include_namespace = service_map_include_namespace
        self.limit = limit
        self.spss = spss
        self.filters = filters if filters is not None else []
        self.group_by = group_by
        self.datasource = datasource
        self.table_type = table_type

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "refId": self.ref_id,
            "filters": self.filters,
        }
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.query is not None:
            payload["query"] = self.query
        if self.search is not None:
            payload["search"] = self.search
        if self.service_name is not None:
            payload["serviceName"] = self.service_name
        if self.span_name is not None:
            payload["spanName"] = self.span_name
        if self.min_duration is not None:
            payload["minDuration"] = self.min_duration
        if self.max_duration is not None:
            payload["maxDuration"] = self.max_duration
        if self.service_map_query is not None:
            payload["serviceMapQuery"] = self.service_map_query
        if self.service_map_include_namespace is not None:
            payload["serviceMapIncludeNamespace"] = self.service_map_include_namespace
        if self.limit is not None:
            payload["limit"] = self.limit
        if self.spss is not None:
            payload["spss"] = self.spss
        if self.group_by is not None:
            payload["groupBy"] = self.group_by
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.table_type is not None:
            payload["tableType"] = self.table_type
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "query" in data:
            args["query"] = data["query"]
        if "search" in data:
            args["search"] = data["search"]
        if "serviceName" in data:
            args["service_name"] = data["serviceName"]
        if "spanName" in data:
            args["span_name"] = data["spanName"]
        if "minDuration" in data:
            args["min_duration"] = data["minDuration"]
        if "maxDuration" in data:
            args["max_duration"] = data["maxDuration"]
        if "serviceMapQuery" in data:
            args["service_map_query"] = data["serviceMapQuery"]
        if "serviceMapIncludeNamespace" in data:
            args["service_map_include_namespace"] = data["serviceMapIncludeNamespace"]
        if "limit" in data:
            args["limit"] = data["limit"]
        if "spss" in data:
            args["spss"] = data["spss"]
        if "filters" in data:
            args["filters"] = data["filters"]
        if "groupBy" in data:
            args["group_by"] = data["groupBy"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]
        if "tableType" in data:
            args["table_type"] = data["tableType"]        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="tempo",
        from_json_hook=TempoQuery.from_json,
    )


class TempoQueryType(enum.StrEnum):
    """
    search = Loki search, nativeSearch = Tempo search for backwards compatibility
    """

    TRACEQL = "traceql"
    TRACEQL_SEARCH = "traceqlSearch"
    SEARCH = "search"
    SERVICE_MAP = "serviceMap"
    UPLOAD = "upload"
    NATIVE_SEARCH = "nativeSearch"
    TRACE_ID = "traceId"
    CLEAR = "clear"


class SearchStreamingState(enum.StrEnum):
    """
    The state of the TraceQL streaming search query
    """

    PENDING = "pending"
    STREAMING = "streaming"
    DONE = "done"
    ERROR = "error"


class SearchTableType(enum.StrEnum):
    """
    The type of the table that is used to display the search results
    """

    TRACES = "traces"
    SPANS = "spans"


class TraceqlSearchScope(enum.StrEnum):
    """
    static fields are pre-set in the UI, dynamic fields are added by the user
    """

    INTRINSIC = "intrinsic"
    UNSCOPED = "unscoped"
    RESOURCE = "resource"
    SPAN = "span"


class TraceqlFilter:
    # Uniquely identify the filter, will not be used in the query generation
    id_val: str
    # The tag for the search filter, for example: .http.status_code, .service.name, status
    tag: typing.Optional[str]
    # The operator that connects the tag to the value, for example: =, >, !=, =~
    operator: typing.Optional[str]
    # The value for the search filter
    value: typing.Optional[typing.Union[str, list[str]]]
    # The type of the value, used for example to check whether we need to wrap the value in quotes when generating the query
    value_type: typing.Optional[str]
    # The scope of the filter, can either be unscoped/all scopes, resource or span
    scope: typing.Optional['TraceqlSearchScope']

    def __init__(self, id_val: str = "", tag: typing.Optional[str] = None, operator: typing.Optional[str] = None, value: typing.Optional[typing.Union[str, list[str]]] = None, value_type: typing.Optional[str] = None, scope: typing.Optional['TraceqlSearchScope'] = None):
        self.id_val = id_val
        self.tag = tag
        self.operator = operator
        self.value = value
        self.value_type = value_type
        self.scope = scope

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
        }
        if self.tag is not None:
            payload["tag"] = self.tag
        if self.operator is not None:
            payload["operator"] = self.operator
        if self.value is not None:
            payload["value"] = self.value
        if self.value_type is not None:
            payload["valueType"] = self.value_type
        if self.scope is not None:
            payload["scope"] = self.scope
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "tag" in data:
            args["tag"] = data["tag"]
        if "operator" in data:
            args["operator"] = data["operator"]
        if "value" in data:
            args["value"] = data["value"]
        if "valueType" in data:
            args["value_type"] = data["valueType"]
        if "scope" in data:
            args["scope"] = data["scope"]        

        return cls(**args)




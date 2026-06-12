# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import variants as cogvariants
from ..models import common
import enum
from ..cog import runtime as cogruntime


class AdhocFilters:
    key: str
    operator: str
    value: str
    values: typing.Optional[list[str]]

    def __init__(self, key: str = "", operator: str = "", value: str = "", values: typing.Optional[list[str]] = None) -> None:
        self.key = key
        self.operator = operator
        self.value = value
        self.values = values

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "key": self.key,
            "operator": self.operator,
            "value": self.value,
        }
        if self.values is not None:
            payload["values"] = self.values
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

        return cls(**args)


class ResultAssertions:
    # Maximum frame count
    max_frames: typing.Optional[int]
    # Type asserts that the frame matches a known type structure.
    # Possible enum values:
    #  - `""` 
    #  - `"timeseries-wide"` 
    #  - `"timeseries-long"` 
    #  - `"timeseries-many"` 
    #  - `"timeseries-multi"` 
    #  - `"directory-listing"` 
    #  - `"table"` 
    #  - `"numeric-wide"` 
    #  - `"numeric-multi"` 
    #  - `"numeric-long"` 
    #  - `"log-lines"` 
    type_val: typing.Optional['ResultAssertionsType']
    # TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    # contract documentation https://grafana.github.io/dataplane/contract/.
    type_version: list[int]

    def __init__(self, max_frames: typing.Optional[int] = None, type_val: typing.Optional['ResultAssertionsType'] = None, type_version: typing.Optional[list[int]] = None) -> None:
        self.max_frames = max_frames
        self.type_val = type_val
        self.type_version = type_version if type_version is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "typeVersion": self.type_version,
        }
        if self.max_frames is not None:
            payload["maxFrames"] = self.max_frames
        if self.type_val is not None:
            payload["type"] = self.type_val
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "maxFrames" in data:
            args["max_frames"] = data["maxFrames"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "typeVersion" in data:
            args["type_version"] = data["typeVersion"]        

        return cls(**args)


class ScopesFilters:
    key: str
    operator: str
    value: str
    values: typing.Optional[list[str]]

    def __init__(self, key: str = "", operator: str = "", value: str = "", values: typing.Optional[list[str]] = None) -> None:
        self.key = key
        self.operator = operator
        self.value = value
        self.values = values

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "key": self.key,
            "operator": self.operator,
            "value": self.value,
        }
        if self.values is not None:
            payload["values"] = self.values
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

        return cls(**args)


class Scopes:
    default_path: typing.Optional[list[str]]
    filters: typing.Optional[list['ScopesFilters']]
    title: str

    def __init__(self, default_path: typing.Optional[list[str]] = None, filters: typing.Optional[list['ScopesFilters']] = None, title: str = "") -> None:
        self.default_path = default_path
        self.filters = filters
        self.title = title

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "title": self.title,
        }
        if self.default_path is not None:
            payload["defaultPath"] = self.default_path
        if self.filters is not None:
            payload["filters"] = self.filters
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "defaultPath" in data:
            args["default_path"] = data["defaultPath"]
        if "filters" in data:
            args["filters"] = [ScopesFilters.from_json(item) for item in data["filters"]]
        if "title" in data:
            args["title"] = data["title"]        

        return cls(**args)


class TimeRange:
    # From is the start time of the query.
    from_val: str
    # To is the end time of the query.
    to: str

    def __init__(self, from_val: str = "now-6h", to: str = "now") -> None:
        self.from_val = from_val
        self.to = to

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "from": self.from_val,
            "to": self.to,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "from" in data:
            args["from_val"] = data["from"]
        if "to" in data:
            args["to"] = data["to"]        

        return cls(**args)


class Dataquery(cogvariants.Dataquery):
    # Additional Ad-hoc filters that take precedence over Scope on conflict.
    adhoc_filters: typing.Optional[list['AdhocFilters']]
    # The datasource
    datasource: typing.Optional[common.DataSourceRef]
    # what we should show in the editor
    # Possible enum values:
    #  - `"builder"` 
    #  - `"code"` 
    editor_mode: typing.Optional['QueryEditorMode']
    # Execute an additional query to identify interesting raw samples relevant for the given expr
    exemplar: typing.Optional[bool]
    # The actual expression/query that will be evaluated by Prometheus
    expr: str
    # The response format
    # Possible enum values:
    #  - `"time_series"` 
    #  - `"table"` 
    #  - `"heatmap"` 
    format_val: typing.Optional['PromQueryFormat']
    # Group By parameters to apply to aggregate expressions in the query
    group_by_keys: typing.Optional[list[str]]
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Returns only the latest value that Prometheus has scraped for the requested time series
    instant: typing.Optional[bool]
    # Used to specify how many times to divide max data points by. We use max data points under query options
    # See https://github.com/grafana/grafana/issues/48081
    # Deprecated: use interval
    interval_factor: typing.Optional[int]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    legend_format: typing.Optional[str]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    range_val: typing.Optional[bool]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: typing.Optional[str]
    # Optionally define expected query result behavior
    result_assertions: typing.Optional['ResultAssertions']
    # A set of filters applied to apply to the query
    scopes: typing.Optional[list['Scopes']]
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional['TimeRange']
    # An additional lower limit for the step parameter of the Prometheus query and for the
    # `$__interval` and `$__rate_interval` variables.
    interval: typing.Optional[str]

    def __init__(self, adhoc_filters: typing.Optional[list['AdhocFilters']] = None, datasource: typing.Optional[common.DataSourceRef] = None, editor_mode: typing.Optional['QueryEditorMode'] = None, exemplar: typing.Optional[bool] = None, expr: str = "", format_val: typing.Optional['PromQueryFormat'] = None, group_by_keys: typing.Optional[list[str]] = None, hide: typing.Optional[bool] = None, instant: typing.Optional[bool] = None, interval_factor: typing.Optional[int] = None, interval_ms: typing.Optional[float] = None, legend_format: typing.Optional[str] = None, max_data_points: typing.Optional[int] = None, query_type: typing.Optional[str] = None, range_val: typing.Optional[bool] = None, ref_id: typing.Optional[str] = None, result_assertions: typing.Optional['ResultAssertions'] = None, scopes: typing.Optional[list['Scopes']] = None, time_range: typing.Optional['TimeRange'] = None, interval: typing.Optional[str] = None) -> None:
        self.adhoc_filters = adhoc_filters
        self.datasource = datasource
        self.editor_mode = editor_mode
        self.exemplar = exemplar
        self.expr = expr
        self.format_val = format_val
        self.group_by_keys = group_by_keys
        self.hide = hide
        self.instant = instant
        self.interval_factor = interval_factor
        self.interval_ms = interval_ms
        self.legend_format = legend_format
        self.max_data_points = max_data_points
        self.query_type = query_type
        self.range_val = range_val
        self.ref_id = ref_id
        self.result_assertions = result_assertions
        self.scopes = scopes
        self.time_range = time_range
        self.interval = interval

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "expr": self.expr,
        }
        if self.adhoc_filters is not None:
            payload["adhocFilters"] = self.adhoc_filters
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.editor_mode is not None:
            payload["editorMode"] = self.editor_mode
        if self.exemplar is not None:
            payload["exemplar"] = self.exemplar
        if self.format_val is not None:
            payload["format"] = self.format_val
        if self.group_by_keys is not None:
            payload["groupByKeys"] = self.group_by_keys
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.instant is not None:
            payload["instant"] = self.instant
        if self.interval_factor is not None:
            payload["intervalFactor"] = self.interval_factor
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        if self.legend_format is not None:
            payload["legendFormat"] = self.legend_format
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.range_val is not None:
            payload["range"] = self.range_val
        if self.ref_id is not None:
            payload["refId"] = self.ref_id
        if self.result_assertions is not None:
            payload["resultAssertions"] = self.result_assertions
        if self.scopes is not None:
            payload["scopes"] = self.scopes
        if self.time_range is not None:
            payload["timeRange"] = self.time_range
        if self.interval is not None:
            payload["interval"] = self.interval
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "adhocFilters" in data:
            args["adhoc_filters"] = [AdhocFilters.from_json(item) for item in data["adhocFilters"]]
        if "datasource" in data:
            args["datasource"] = common.DataSourceRef.from_json(data["datasource"])
        if "editorMode" in data:
            args["editor_mode"] = data["editorMode"]
        if "exemplar" in data:
            args["exemplar"] = data["exemplar"]
        if "expr" in data:
            args["expr"] = data["expr"]
        if "format" in data:
            args["format_val"] = data["format"]
        if "groupByKeys" in data:
            args["group_by_keys"] = data["groupByKeys"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "instant" in data:
            args["instant"] = data["instant"]
        if "intervalFactor" in data:
            args["interval_factor"] = data["intervalFactor"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]
        if "legendFormat" in data:
            args["legend_format"] = data["legendFormat"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "range" in data:
            args["range_val"] = data["range"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "resultAssertions" in data:
            args["result_assertions"] = ResultAssertions.from_json(data["resultAssertions"])
        if "scopes" in data:
            args["scopes"] = [Scopes.from_json(item) for item in data["scopes"]]
        if "timeRange" in data:
            args["time_range"] = TimeRange.from_json(data["timeRange"])
        if "interval" in data:
            args["interval"] = data["interval"]        

        return cls(**args)


class ResultAssertionsType(enum.StrEnum):
    NONE = ""
    TIMESERIES_WIDE = "timeseries-wide"
    TIMESERIES_LONG = "timeseries-long"
    TIMESERIES_MANY = "timeseries-many"
    TIMESERIES_MULTI = "timeseries-multi"
    DIRECTORY_LISTING = "directory-listing"
    TABLE = "table"
    NUMERIC_WIDE = "numeric-wide"
    NUMERIC_MULTI = "numeric-multi"
    NUMERIC_LONG = "numeric-long"
    LOG_LINES = "log-lines"


class QueryEditorMode(enum.StrEnum):
    BUILDER = "builder"
    CODE = "code"


class PromQueryFormat(enum.StrEnum):
    TIME_SERIES = "time_series"
    TABLE = "table"
    HEATMAP = "heatmap"





def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="prometheus",
        from_json_hook=Dataquery.from_json,
    )


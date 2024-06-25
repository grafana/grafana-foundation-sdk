# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import runtime as cogruntime
from ..cog import variants as cogvariants


Expr: typing.TypeAlias = typing.Union['TypeMath', 'TypeReduce', 'TypeResample', 'TypeClassicConditions', 'TypeThreshold', 'TypeSql']


def variant_config() -> cogruntime.DataqueryConfig:
    decoding_map: dict[str, typing.Union[typing.Type[TypeSql], typing.Type[TypeMath], typing.Type[TypeReduce], typing.Type[TypeResample], typing.Type[TypeClassicConditions], typing.Type[TypeThreshold]]] = {"sql": TypeSql, "math": TypeMath, "reduce": TypeReduce, "resample": TypeResample, "classic_conditions": TypeClassicConditions, "threshold": TypeThreshold}
    return cogruntime.DataqueryConfig(
        identifier="__expr__",
        from_json_hook=lambda data: decoding_map[data["type"]].from_json(data),
    )


class TypeMath(cogvariants.Dataquery):
    # The datasource
    datasource: typing.Optional['ExprTypeMathDatasource']
    # General math expression
    expression: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: str
    # Optionally define expected query result behavior
    result_assertions: typing.Optional['ExprTypeMathResultAssertions']
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional['ExprTypeMathTimeRange']
    type_val: typing.Literal["math"]

    def __init__(self, datasource: typing.Optional['ExprTypeMathDatasource'] = None, expression: str = "", hide: typing.Optional[bool] = None, interval_ms: typing.Optional[float] = None, max_data_points: typing.Optional[int] = None, query_type: typing.Optional[str] = None, ref_id: str = "", result_assertions: typing.Optional['ExprTypeMathResultAssertions'] = None, time_range: typing.Optional['ExprTypeMathTimeRange'] = None):
        self.datasource = datasource
        self.expression = expression
        self.hide = hide
        self.interval_ms = interval_ms
        self.max_data_points = max_data_points
        self.query_type = query_type
        self.ref_id = ref_id
        self.result_assertions = result_assertions
        self.time_range = time_range
        self.type_val = "math"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "expression": self.expression,
            "refId": self.ref_id,
            "type": self.type_val,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.result_assertions is not None:
            payload["resultAssertions"] = self.result_assertions
        if self.time_range is not None:
            payload["timeRange"] = self.time_range
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "datasource" in data:
            args["datasource"] = ExprTypeMathDatasource.from_json(data["datasource"])
        if "expression" in data:
            args["expression"] = data["expression"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "resultAssertions" in data:
            args["result_assertions"] = ExprTypeMathResultAssertions.from_json(data["resultAssertions"])
        if "timeRange" in data:
            args["time_range"] = ExprTypeMathTimeRange.from_json(data["timeRange"])        

        return cls(**args)


class TypeReduce(cogvariants.Dataquery):
    # The datasource
    datasource: typing.Optional['ExprTypeReduceDatasource']
    # Reference to single query result
    expression: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # The reducer
    # Possible enum values:
    #  - `"sum"` 
    #  - `"mean"` 
    #  - `"min"` 
    #  - `"max"` 
    #  - `"count"` 
    #  - `"last"` 
    reducer: typing.Literal["sum", "mean", "min", "max", "count", "last"]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: str
    # Optionally define expected query result behavior
    result_assertions: typing.Optional['ExprTypeReduceResultAssertions']
    # Reducer Options
    settings: typing.Optional['ExprTypeReduceSettings']
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional['ExprTypeReduceTimeRange']
    type_val: typing.Literal["reduce"]

    def __init__(self, datasource: typing.Optional['ExprTypeReduceDatasource'] = None, expression: str = "", hide: typing.Optional[bool] = None, interval_ms: typing.Optional[float] = None, max_data_points: typing.Optional[int] = None, query_type: typing.Optional[str] = None, reducer: typing.Optional[typing.Literal["sum", "mean", "min", "max", "count", "last"]] = None, ref_id: str = "", result_assertions: typing.Optional['ExprTypeReduceResultAssertions'] = None, settings: typing.Optional['ExprTypeReduceSettings'] = None, time_range: typing.Optional['ExprTypeReduceTimeRange'] = None):
        self.datasource = datasource
        self.expression = expression
        self.hide = hide
        self.interval_ms = interval_ms
        self.max_data_points = max_data_points
        self.query_type = query_type
        self.reducer = reducer if reducer is not None else "sum"
        self.ref_id = ref_id
        self.result_assertions = result_assertions
        self.settings = settings
        self.time_range = time_range
        self.type_val = "reduce"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "expression": self.expression,
            "reducer": self.reducer,
            "refId": self.ref_id,
            "type": self.type_val,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.result_assertions is not None:
            payload["resultAssertions"] = self.result_assertions
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.time_range is not None:
            payload["timeRange"] = self.time_range
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "datasource" in data:
            args["datasource"] = ExprTypeReduceDatasource.from_json(data["datasource"])
        if "expression" in data:
            args["expression"] = data["expression"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "reducer" in data:
            args["reducer"] = data["reducer"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "resultAssertions" in data:
            args["result_assertions"] = ExprTypeReduceResultAssertions.from_json(data["resultAssertions"])
        if "settings" in data:
            args["settings"] = ExprTypeReduceSettings.from_json(data["settings"])
        if "timeRange" in data:
            args["time_range"] = ExprTypeReduceTimeRange.from_json(data["timeRange"])        

        return cls(**args)


class TypeResample(cogvariants.Dataquery):
    # The datasource
    datasource: typing.Optional['ExprTypeResampleDatasource']
    # The downsample function
    # Possible enum values:
    #  - `"sum"` 
    #  - `"mean"` 
    #  - `"min"` 
    #  - `"max"` 
    #  - `"count"` 
    #  - `"last"` 
    downsampler: typing.Literal["sum", "mean", "min", "max", "count", "last"]
    # The math expression
    expression: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: str
    # Optionally define expected query result behavior
    result_assertions: typing.Optional['ExprTypeResampleResultAssertions']
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional['ExprTypeResampleTimeRange']
    type_val: typing.Literal["resample"]
    # The upsample function
    # Possible enum values:
    #  - `"pad"` Use the last seen value
    #  - `"backfilling"` backfill
    #  - `"fillna"` Do not fill values (nill)
    upsampler: typing.Literal["pad", "backfilling", "fillna"]
    # The time duration
    window: str

    def __init__(self, datasource: typing.Optional['ExprTypeResampleDatasource'] = None, downsampler: typing.Optional[typing.Literal["sum", "mean", "min", "max", "count", "last"]] = None, expression: str = "", hide: typing.Optional[bool] = None, interval_ms: typing.Optional[float] = None, max_data_points: typing.Optional[int] = None, query_type: typing.Optional[str] = None, ref_id: str = "", result_assertions: typing.Optional['ExprTypeResampleResultAssertions'] = None, time_range: typing.Optional['ExprTypeResampleTimeRange'] = None, upsampler: typing.Optional[typing.Literal["pad", "backfilling", "fillna"]] = None, window: str = ""):
        self.datasource = datasource
        self.downsampler = downsampler if downsampler is not None else "sum"
        self.expression = expression
        self.hide = hide
        self.interval_ms = interval_ms
        self.max_data_points = max_data_points
        self.query_type = query_type
        self.ref_id = ref_id
        self.result_assertions = result_assertions
        self.time_range = time_range
        self.type_val = "resample"
        self.upsampler = upsampler if upsampler is not None else "pad"
        self.window = window

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "downsampler": self.downsampler,
            "expression": self.expression,
            "refId": self.ref_id,
            "type": self.type_val,
            "upsampler": self.upsampler,
            "window": self.window,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.result_assertions is not None:
            payload["resultAssertions"] = self.result_assertions
        if self.time_range is not None:
            payload["timeRange"] = self.time_range
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "datasource" in data:
            args["datasource"] = ExprTypeResampleDatasource.from_json(data["datasource"])
        if "downsampler" in data:
            args["downsampler"] = data["downsampler"]
        if "expression" in data:
            args["expression"] = data["expression"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "resultAssertions" in data:
            args["result_assertions"] = ExprTypeResampleResultAssertions.from_json(data["resultAssertions"])
        if "timeRange" in data:
            args["time_range"] = ExprTypeResampleTimeRange.from_json(data["timeRange"])
        if "upsampler" in data:
            args["upsampler"] = data["upsampler"]
        if "window" in data:
            args["window"] = data["window"]        

        return cls(**args)


class TypeClassicConditions(cogvariants.Dataquery):
    conditions: list['ExprTypeClassicConditionsConditions']
    # The datasource
    datasource: typing.Optional['ExprTypeClassicConditionsDatasource']
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: str
    # Optionally define expected query result behavior
    result_assertions: typing.Optional['ExprTypeClassicConditionsResultAssertions']
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional['ExprTypeClassicConditionsTimeRange']
    type_val: typing.Literal["classic_conditions"]

    def __init__(self, conditions: typing.Optional[list['ExprTypeClassicConditionsConditions']] = None, datasource: typing.Optional['ExprTypeClassicConditionsDatasource'] = None, hide: typing.Optional[bool] = None, interval_ms: typing.Optional[float] = None, max_data_points: typing.Optional[int] = None, query_type: typing.Optional[str] = None, ref_id: str = "", result_assertions: typing.Optional['ExprTypeClassicConditionsResultAssertions'] = None, time_range: typing.Optional['ExprTypeClassicConditionsTimeRange'] = None):
        self.conditions = conditions if conditions is not None else []
        self.datasource = datasource
        self.hide = hide
        self.interval_ms = interval_ms
        self.max_data_points = max_data_points
        self.query_type = query_type
        self.ref_id = ref_id
        self.result_assertions = result_assertions
        self.time_range = time_range
        self.type_val = "classic_conditions"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "conditions": self.conditions,
            "refId": self.ref_id,
            "type": self.type_val,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.result_assertions is not None:
            payload["resultAssertions"] = self.result_assertions
        if self.time_range is not None:
            payload["timeRange"] = self.time_range
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "conditions" in data:
            args["conditions"] = data["conditions"]
        if "datasource" in data:
            args["datasource"] = ExprTypeClassicConditionsDatasource.from_json(data["datasource"])
        if "hide" in data:
            args["hide"] = data["hide"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "resultAssertions" in data:
            args["result_assertions"] = ExprTypeClassicConditionsResultAssertions.from_json(data["resultAssertions"])
        if "timeRange" in data:
            args["time_range"] = ExprTypeClassicConditionsTimeRange.from_json(data["timeRange"])        

        return cls(**args)


class TypeThreshold(cogvariants.Dataquery):
    # Threshold Conditions
    conditions: list['ExprTypeThresholdConditions']
    # The datasource
    datasource: typing.Optional['ExprTypeThresholdDatasource']
    # Reference to single query result
    expression: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: str
    # Optionally define expected query result behavior
    result_assertions: typing.Optional['ExprTypeThresholdResultAssertions']
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional['ExprTypeThresholdTimeRange']
    type_val: typing.Literal["threshold"]

    def __init__(self, conditions: typing.Optional[list['ExprTypeThresholdConditions']] = None, datasource: typing.Optional['ExprTypeThresholdDatasource'] = None, expression: str = "", hide: typing.Optional[bool] = None, interval_ms: typing.Optional[float] = None, max_data_points: typing.Optional[int] = None, query_type: typing.Optional[str] = None, ref_id: str = "", result_assertions: typing.Optional['ExprTypeThresholdResultAssertions'] = None, time_range: typing.Optional['ExprTypeThresholdTimeRange'] = None):
        self.conditions = conditions if conditions is not None else []
        self.datasource = datasource
        self.expression = expression
        self.hide = hide
        self.interval_ms = interval_ms
        self.max_data_points = max_data_points
        self.query_type = query_type
        self.ref_id = ref_id
        self.result_assertions = result_assertions
        self.time_range = time_range
        self.type_val = "threshold"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "conditions": self.conditions,
            "expression": self.expression,
            "refId": self.ref_id,
            "type": self.type_val,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.result_assertions is not None:
            payload["resultAssertions"] = self.result_assertions
        if self.time_range is not None:
            payload["timeRange"] = self.time_range
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "conditions" in data:
            args["conditions"] = data["conditions"]
        if "datasource" in data:
            args["datasource"] = ExprTypeThresholdDatasource.from_json(data["datasource"])
        if "expression" in data:
            args["expression"] = data["expression"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "resultAssertions" in data:
            args["result_assertions"] = ExprTypeThresholdResultAssertions.from_json(data["resultAssertions"])
        if "timeRange" in data:
            args["time_range"] = ExprTypeThresholdTimeRange.from_json(data["timeRange"])        

        return cls(**args)


class TypeSql(cogvariants.Dataquery):
    # The datasource
    datasource: typing.Optional['ExprTypeSqlDatasource']
    expression: str
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: str
    # Optionally define expected query result behavior
    result_assertions: typing.Optional['ExprTypeSqlResultAssertions']
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional['ExprTypeSqlTimeRange']
    type_val: typing.Literal["sql"]

    def __init__(self, datasource: typing.Optional['ExprTypeSqlDatasource'] = None, expression: str = "", hide: typing.Optional[bool] = None, interval_ms: typing.Optional[float] = None, max_data_points: typing.Optional[int] = None, query_type: typing.Optional[str] = None, ref_id: str = "", result_assertions: typing.Optional['ExprTypeSqlResultAssertions'] = None, time_range: typing.Optional['ExprTypeSqlTimeRange'] = None):
        self.datasource = datasource
        self.expression = expression
        self.hide = hide
        self.interval_ms = interval_ms
        self.max_data_points = max_data_points
        self.query_type = query_type
        self.ref_id = ref_id
        self.result_assertions = result_assertions
        self.time_range = time_range
        self.type_val = "sql"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "expression": self.expression,
            "refId": self.ref_id,
            "type": self.type_val,
        }
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.result_assertions is not None:
            payload["resultAssertions"] = self.result_assertions
        if self.time_range is not None:
            payload["timeRange"] = self.time_range
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "datasource" in data:
            args["datasource"] = ExprTypeSqlDatasource.from_json(data["datasource"])
        if "expression" in data:
            args["expression"] = data["expression"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "resultAssertions" in data:
            args["result_assertions"] = ExprTypeSqlResultAssertions.from_json(data["resultAssertions"])
        if "timeRange" in data:
            args["time_range"] = ExprTypeSqlTimeRange.from_json(data["timeRange"])        

        return cls(**args)


class TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql:
    type_math: typing.Optional['TypeMath']
    type_reduce: typing.Optional['TypeReduce']
    type_resample: typing.Optional['TypeResample']
    type_classic_conditions: typing.Optional['TypeClassicConditions']
    type_threshold: typing.Optional['TypeThreshold']
    type_sql: typing.Optional['TypeSql']

    def __init__(self, type_math: typing.Optional['TypeMath'] = None, type_reduce: typing.Optional['TypeReduce'] = None, type_resample: typing.Optional['TypeResample'] = None, type_classic_conditions: typing.Optional['TypeClassicConditions'] = None, type_threshold: typing.Optional['TypeThreshold'] = None, type_sql: typing.Optional['TypeSql'] = None):
        self.type_math = type_math
        self.type_reduce = type_reduce
        self.type_resample = type_resample
        self.type_classic_conditions = type_classic_conditions
        self.type_threshold = type_threshold
        self.type_sql = type_sql

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.type_math is not None:
            payload["TypeMath"] = self.type_math
        if self.type_reduce is not None:
            payload["TypeReduce"] = self.type_reduce
        if self.type_resample is not None:
            payload["TypeResample"] = self.type_resample
        if self.type_classic_conditions is not None:
            payload["TypeClassicConditions"] = self.type_classic_conditions
        if self.type_threshold is not None:
            payload["TypeThreshold"] = self.type_threshold
        if self.type_sql is not None:
            payload["TypeSql"] = self.type_sql
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "TypeMath" in data:
            args["type_math"] = TypeMath.from_json(data["TypeMath"])
        if "TypeReduce" in data:
            args["type_reduce"] = TypeReduce.from_json(data["TypeReduce"])
        if "TypeResample" in data:
            args["type_resample"] = TypeResample.from_json(data["TypeResample"])
        if "TypeClassicConditions" in data:
            args["type_classic_conditions"] = TypeClassicConditions.from_json(data["TypeClassicConditions"])
        if "TypeThreshold" in data:
            args["type_threshold"] = TypeThreshold.from_json(data["TypeThreshold"])
        if "TypeSql" in data:
            args["type_sql"] = TypeSql.from_json(data["TypeSql"])        

        return cls(**args)


class ExprTypeMathDatasource:
    # The apiserver version
    api_version: typing.Optional[str]
    # The datasource plugin type
    type_val: typing.Literal["__expr__"]
    # Datasource UID (NOTE: name in k8s)
    uid: typing.Optional[str]

    def __init__(self, api_version: typing.Optional[str] = None, uid: typing.Optional[str] = None):
        self.api_version = api_version
        self.type_val = "__expr__"
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.api_version is not None:
            payload["apiVersion"] = self.api_version
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "apiVersion" in data:
            args["api_version"] = data["apiVersion"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class ExprTypeMathResultAssertions:
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
    type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]]
    # TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    # contract documentation https://grafana.github.io/dataplane/contract/.
    type_version: list[int]

    def __init__(self, max_frames: typing.Optional[int] = None, type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]] = None, type_version: typing.Optional[list[int]] = None):
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


class ExprTypeMathTimeRange:
    # From is the start time of the query.
    from_val: str
    # To is the end time of the query.
    to: str

    def __init__(self, from_val: str = "now-6h", to: str = "now"):
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


class ExprTypeReduceDatasource:
    # The apiserver version
    api_version: typing.Optional[str]
    # The datasource plugin type
    type_val: typing.Literal["__expr__"]
    # Datasource UID (NOTE: name in k8s)
    uid: typing.Optional[str]

    def __init__(self, api_version: typing.Optional[str] = None, uid: typing.Optional[str] = None):
        self.api_version = api_version
        self.type_val = "__expr__"
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.api_version is not None:
            payload["apiVersion"] = self.api_version
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "apiVersion" in data:
            args["api_version"] = data["apiVersion"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class ExprTypeReduceResultAssertions:
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
    type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]]
    # TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    # contract documentation https://grafana.github.io/dataplane/contract/.
    type_version: list[int]

    def __init__(self, max_frames: typing.Optional[int] = None, type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]] = None, type_version: typing.Optional[list[int]] = None):
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


class ExprTypeReduceSettings:
    # Non-number reduce behavior
    # Possible enum values:
    #  - `"dropNN"` Drop non-numbers
    #  - `"replaceNN"` Replace non-numbers
    mode: typing.Literal["dropNN", "replaceNN"]
    # Only valid when mode is replace
    replace_with_value: typing.Optional[float]

    def __init__(self, mode: typing.Optional[typing.Literal["dropNN", "replaceNN"]] = None, replace_with_value: typing.Optional[float] = None):
        self.mode = mode if mode is not None else "dropNN"
        self.replace_with_value = replace_with_value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "mode": self.mode,
        }
        if self.replace_with_value is not None:
            payload["replaceWithValue"] = self.replace_with_value
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "replaceWithValue" in data:
            args["replace_with_value"] = data["replaceWithValue"]        

        return cls(**args)


class ExprTypeReduceTimeRange:
    # From is the start time of the query.
    from_val: str
    # To is the end time of the query.
    to: str

    def __init__(self, from_val: str = "now-6h", to: str = "now"):
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


class ExprTypeResampleDatasource:
    # The apiserver version
    api_version: typing.Optional[str]
    # The datasource plugin type
    type_val: typing.Literal["__expr__"]
    # Datasource UID (NOTE: name in k8s)
    uid: typing.Optional[str]

    def __init__(self, api_version: typing.Optional[str] = None, uid: typing.Optional[str] = None):
        self.api_version = api_version
        self.type_val = "__expr__"
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.api_version is not None:
            payload["apiVersion"] = self.api_version
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "apiVersion" in data:
            args["api_version"] = data["apiVersion"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class ExprTypeResampleResultAssertions:
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
    type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]]
    # TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    # contract documentation https://grafana.github.io/dataplane/contract/.
    type_version: list[int]

    def __init__(self, max_frames: typing.Optional[int] = None, type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]] = None, type_version: typing.Optional[list[int]] = None):
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


class ExprTypeResampleTimeRange:
    # From is the start time of the query.
    from_val: str
    # To is the end time of the query.
    to: str

    def __init__(self, from_val: str = "now-6h", to: str = "now"):
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


class ExprTypeClassicConditionsConditionsEvaluator:
    params: list[float]
    # e.g. "gt"
    type_val: str

    def __init__(self, params: typing.Optional[list[float]] = None, type_val: str = ""):
        self.params = params if params is not None else []
        self.type_val = type_val

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "params": self.params,
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "params" in data:
            args["params"] = data["params"]
        if "type" in data:
            args["type_val"] = data["type"]        

        return cls(**args)


class ExprTypeClassicConditionsConditionsOperator:
    type_val: typing.Literal["and", "or"]

    def __init__(self, type_val: typing.Optional[typing.Literal["and", "or"]] = None):
        self.type_val = type_val if type_val is not None else "and"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]        

        return cls(**args)


class ExprTypeClassicConditionsConditionsQuery:
    params: list[str]

    def __init__(self, params: typing.Optional[list[str]] = None):
        self.params = params if params is not None else []

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "params": self.params,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "params" in data:
            args["params"] = data["params"]        

        return cls(**args)


class ExprTypeClassicConditionsConditionsReducer:
    type_val: str

    def __init__(self, type_val: str = ""):
        self.type_val = type_val

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]        

        return cls(**args)


class ExprTypeClassicConditionsConditions:
    evaluator: 'ExprTypeClassicConditionsConditionsEvaluator'
    operator: 'ExprTypeClassicConditionsConditionsOperator'
    query: 'ExprTypeClassicConditionsConditionsQuery'
    reducer: 'ExprTypeClassicConditionsConditionsReducer'

    def __init__(self, evaluator: typing.Optional['ExprTypeClassicConditionsConditionsEvaluator'] = None, operator: typing.Optional['ExprTypeClassicConditionsConditionsOperator'] = None, query: typing.Optional['ExprTypeClassicConditionsConditionsQuery'] = None, reducer: typing.Optional['ExprTypeClassicConditionsConditionsReducer'] = None):
        self.evaluator = evaluator if evaluator is not None else ExprTypeClassicConditionsConditionsEvaluator()
        self.operator = operator if operator is not None else ExprTypeClassicConditionsConditionsOperator()
        self.query = query if query is not None else ExprTypeClassicConditionsConditionsQuery()
        self.reducer = reducer if reducer is not None else ExprTypeClassicConditionsConditionsReducer()

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "evaluator": self.evaluator,
            "operator": self.operator,
            "query": self.query,
            "reducer": self.reducer,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "evaluator" in data:
            args["evaluator"] = ExprTypeClassicConditionsConditionsEvaluator.from_json(data["evaluator"])
        if "operator" in data:
            args["operator"] = ExprTypeClassicConditionsConditionsOperator.from_json(data["operator"])
        if "query" in data:
            args["query"] = ExprTypeClassicConditionsConditionsQuery.from_json(data["query"])
        if "reducer" in data:
            args["reducer"] = ExprTypeClassicConditionsConditionsReducer.from_json(data["reducer"])        

        return cls(**args)


class ExprTypeClassicConditionsDatasource:
    # The apiserver version
    api_version: typing.Optional[str]
    # The datasource plugin type
    type_val: typing.Literal["__expr__"]
    # Datasource UID (NOTE: name in k8s)
    uid: typing.Optional[str]

    def __init__(self, api_version: typing.Optional[str] = None, uid: typing.Optional[str] = None):
        self.api_version = api_version
        self.type_val = "__expr__"
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.api_version is not None:
            payload["apiVersion"] = self.api_version
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "apiVersion" in data:
            args["api_version"] = data["apiVersion"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class ExprTypeClassicConditionsResultAssertions:
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
    type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]]
    # TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    # contract documentation https://grafana.github.io/dataplane/contract/.
    type_version: list[int]

    def __init__(self, max_frames: typing.Optional[int] = None, type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]] = None, type_version: typing.Optional[list[int]] = None):
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


class ExprTypeClassicConditionsTimeRange:
    # From is the start time of the query.
    from_val: str
    # To is the end time of the query.
    to: str

    def __init__(self, from_val: str = "now-6h", to: str = "now"):
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


class ExprTypeThresholdConditionsEvaluator:
    params: list[float]
    # e.g. "gt"
    type_val: typing.Literal["gt", "lt", "within_range", "outside_range"]

    def __init__(self, params: typing.Optional[list[float]] = None, type_val: typing.Optional[typing.Literal["gt", "lt", "within_range", "outside_range"]] = None):
        self.params = params if params is not None else []
        self.type_val = type_val if type_val is not None else "gt"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "params": self.params,
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "params" in data:
            args["params"] = data["params"]
        if "type" in data:
            args["type_val"] = data["type"]        

        return cls(**args)


class ExprTypeThresholdConditionsUnloadEvaluator:
    params: list[float]
    # e.g. "gt"
    type_val: typing.Literal["gt", "lt", "within_range", "outside_range"]

    def __init__(self, params: typing.Optional[list[float]] = None, type_val: typing.Optional[typing.Literal["gt", "lt", "within_range", "outside_range"]] = None):
        self.params = params if params is not None else []
        self.type_val = type_val if type_val is not None else "gt"

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "params": self.params,
            "type": self.type_val,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "params" in data:
            args["params"] = data["params"]
        if "type" in data:
            args["type_val"] = data["type"]        

        return cls(**args)


class ExprTypeThresholdConditions:
    evaluator: 'ExprTypeThresholdConditionsEvaluator'
    loaded_dimensions: typing.Optional[object]
    unload_evaluator: typing.Optional['ExprTypeThresholdConditionsUnloadEvaluator']

    def __init__(self, evaluator: typing.Optional['ExprTypeThresholdConditionsEvaluator'] = None, loaded_dimensions: typing.Optional[object] = None, unload_evaluator: typing.Optional['ExprTypeThresholdConditionsUnloadEvaluator'] = None):
        self.evaluator = evaluator if evaluator is not None else ExprTypeThresholdConditionsEvaluator()
        self.loaded_dimensions = loaded_dimensions
        self.unload_evaluator = unload_evaluator

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "evaluator": self.evaluator,
        }
        if self.loaded_dimensions is not None:
            payload["loadedDimensions"] = self.loaded_dimensions
        if self.unload_evaluator is not None:
            payload["unloadEvaluator"] = self.unload_evaluator
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "evaluator" in data:
            args["evaluator"] = ExprTypeThresholdConditionsEvaluator.from_json(data["evaluator"])
        if "loadedDimensions" in data:
            args["loaded_dimensions"] = data["loadedDimensions"]
        if "unloadEvaluator" in data:
            args["unload_evaluator"] = ExprTypeThresholdConditionsUnloadEvaluator.from_json(data["unloadEvaluator"])        

        return cls(**args)


class ExprTypeThresholdDatasource:
    # The apiserver version
    api_version: typing.Optional[str]
    # The datasource plugin type
    type_val: typing.Literal["__expr__"]
    # Datasource UID (NOTE: name in k8s)
    uid: typing.Optional[str]

    def __init__(self, api_version: typing.Optional[str] = None, uid: typing.Optional[str] = None):
        self.api_version = api_version
        self.type_val = "__expr__"
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.api_version is not None:
            payload["apiVersion"] = self.api_version
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "apiVersion" in data:
            args["api_version"] = data["apiVersion"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class ExprTypeThresholdResultAssertions:
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
    type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]]
    # TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    # contract documentation https://grafana.github.io/dataplane/contract/.
    type_version: list[int]

    def __init__(self, max_frames: typing.Optional[int] = None, type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]] = None, type_version: typing.Optional[list[int]] = None):
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


class ExprTypeThresholdTimeRange:
    # From is the start time of the query.
    from_val: str
    # To is the end time of the query.
    to: str

    def __init__(self, from_val: str = "now-6h", to: str = "now"):
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


class ExprTypeSqlDatasource:
    # The apiserver version
    api_version: typing.Optional[str]
    # The datasource plugin type
    type_val: typing.Literal["__expr__"]
    # Datasource UID (NOTE: name in k8s)
    uid: typing.Optional[str]

    def __init__(self, api_version: typing.Optional[str] = None, uid: typing.Optional[str] = None):
        self.api_version = api_version
        self.type_val = "__expr__"
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.api_version is not None:
            payload["apiVersion"] = self.api_version
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "apiVersion" in data:
            args["api_version"] = data["apiVersion"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class ExprTypeSqlResultAssertions:
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
    type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]]
    # TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
    # contract documentation https://grafana.github.io/dataplane/contract/.
    type_version: list[int]

    def __init__(self, max_frames: typing.Optional[int] = None, type_val: typing.Optional[typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]] = None, type_version: typing.Optional[list[int]] = None):
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


class ExprTypeSqlTimeRange:
    # From is the start time of the query.
    from_val: str
    # To is the end time of the query.
    to: str

    def __init__(self, from_val: str = "now-6h", to: str = "now"):
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




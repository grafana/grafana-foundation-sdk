# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import expr
from ..models import dashboard


class TypeMath(cogbuilder.Builder[expr.TypeMath]):    
    _internal: expr.TypeMath

    def __init__(self):
        self._internal = expr.TypeMath()        
        self._internal.type_val = "math"

    def build(self) -> expr.TypeMath:
        """
        Builds the object.
        """
        return self._internal    
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        The datasource
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def expression(self, expression: str) -> typing.Self:    
        """
        General math expression
        """
            
        if not len(expression) >= 1:
            raise ValueError("len(expression) must be >= 1")
        self._internal.expression = expression
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        self._internal.interval_ms = interval_ms
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        self._internal.max_data_points = max_data_points
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[expr.ExprTypeMathResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        result_assertions_resource = result_assertions.build()
        self._internal.result_assertions = result_assertions_resource
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[expr.ExprTypeMathTimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        time_range_resource = time_range.build()
        self._internal.time_range = time_range_resource
    
        return self
    

class TypeReduce(cogbuilder.Builder[expr.TypeReduce]):    
    _internal: expr.TypeReduce

    def __init__(self):
        self._internal = expr.TypeReduce()        
        self._internal.type_val = "reduce"

    def build(self) -> expr.TypeReduce:
        """
        Builds the object.
        """
        return self._internal    
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        The datasource
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def expression(self, expression: str) -> typing.Self:    
        """
        Reference to single query result
        """
            
        if not len(expression) >= 1:
            raise ValueError("len(expression) must be >= 1")
        self._internal.expression = expression
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        self._internal.interval_ms = interval_ms
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        self._internal.max_data_points = max_data_points
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def reducer(self, reducer: typing.Literal["sum", "mean", "min", "max", "count", "last", "median"]) -> typing.Self:    
        """
        The reducer
        Possible enum values:
         - `"sum"` 
         - `"mean"` 
         - `"min"` 
         - `"max"` 
         - `"count"` 
         - `"last"` 
         - `"median"` 
        """
            
        self._internal.reducer = reducer
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[expr.ExprTypeReduceResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        result_assertions_resource = result_assertions.build()
        self._internal.result_assertions = result_assertions_resource
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[expr.ExprTypeReduceSettings]) -> typing.Self:    
        """
        Reducer Options
        """
            
        settings_resource = settings.build()
        self._internal.settings = settings_resource
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[expr.ExprTypeReduceTimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        time_range_resource = time_range.build()
        self._internal.time_range = time_range_resource
    
        return self
    

class TypeResample(cogbuilder.Builder[expr.TypeResample]):    
    _internal: expr.TypeResample

    def __init__(self):
        self._internal = expr.TypeResample()        
        self._internal.type_val = "resample"

    def build(self) -> expr.TypeResample:
        """
        Builds the object.
        """
        return self._internal    
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        The datasource
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def downsampler(self, downsampler: typing.Literal["sum", "mean", "min", "max", "count", "last", "median"]) -> typing.Self:    
        """
        The downsample function
        Possible enum values:
         - `"sum"` 
         - `"mean"` 
         - `"min"` 
         - `"max"` 
         - `"count"` 
         - `"last"` 
         - `"median"` 
        """
            
        self._internal.downsampler = downsampler
    
        return self
    
    def expression(self, expression: str) -> typing.Self:    
        """
        The math expression
        """
            
        if not len(expression) >= 1:
            raise ValueError("len(expression) must be >= 1")
        self._internal.expression = expression
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        self._internal.interval_ms = interval_ms
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        self._internal.max_data_points = max_data_points
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[expr.ExprTypeResampleResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        result_assertions_resource = result_assertions.build()
        self._internal.result_assertions = result_assertions_resource
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[expr.ExprTypeResampleTimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        time_range_resource = time_range.build()
        self._internal.time_range = time_range_resource
    
        return self
    
    def upsampler(self, upsampler: typing.Literal["pad", "backfilling", "fillna"]) -> typing.Self:    
        """
        The upsample function
        Possible enum values:
         - `"pad"` Use the last seen value
         - `"backfilling"` backfill
         - `"fillna"` Do not fill values (nill)
        """
            
        self._internal.upsampler = upsampler
    
        return self
    
    def window(self, window: str) -> typing.Self:    
        """
        The time duration
        """
            
        if not len(window) >= 1:
            raise ValueError("len(window) must be >= 1")
        self._internal.window = window
    
        return self
    

class TypeClassicConditions(cogbuilder.Builder[expr.TypeClassicConditions]):    
    _internal: expr.TypeClassicConditions

    def __init__(self):
        self._internal = expr.TypeClassicConditions()        
        self._internal.type_val = "classic_conditions"

    def build(self) -> expr.TypeClassicConditions:
        """
        Builds the object.
        """
        return self._internal    
    
    def conditions(self, conditions: list[cogbuilder.Builder[expr.ExprTypeClassicConditionsConditions]]) -> typing.Self:        
        conditions_resources = [r1.build() for r1 in conditions]
        self._internal.conditions = conditions_resources
    
        return self
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        The datasource
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        self._internal.interval_ms = interval_ms
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        self._internal.max_data_points = max_data_points
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[expr.ExprTypeClassicConditionsResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        result_assertions_resource = result_assertions.build()
        self._internal.result_assertions = result_assertions_resource
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[expr.ExprTypeClassicConditionsTimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        time_range_resource = time_range.build()
        self._internal.time_range = time_range_resource
    
        return self
    

class TypeThreshold(cogbuilder.Builder[expr.TypeThreshold]):    
    _internal: expr.TypeThreshold

    def __init__(self):
        self._internal = expr.TypeThreshold()        
        self._internal.type_val = "threshold"

    def build(self) -> expr.TypeThreshold:
        """
        Builds the object.
        """
        return self._internal    
    
    def conditions(self, conditions: list[cogbuilder.Builder[expr.ExprTypeThresholdConditions]]) -> typing.Self:    
        """
        Threshold Conditions
        """
            
        conditions_resources = [r1.build() for r1 in conditions]
        self._internal.conditions = conditions_resources
    
        return self
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        The datasource
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def expression(self, expression: str) -> typing.Self:    
        """
        Reference to single query result
        """
            
        if not len(expression) >= 1:
            raise ValueError("len(expression) must be >= 1")
        self._internal.expression = expression
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        self._internal.interval_ms = interval_ms
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        self._internal.max_data_points = max_data_points
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[expr.ExprTypeThresholdResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        result_assertions_resource = result_assertions.build()
        self._internal.result_assertions = result_assertions_resource
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[expr.ExprTypeThresholdTimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        time_range_resource = time_range.build()
        self._internal.time_range = time_range_resource
    
        return self
    

class TypeSql(cogbuilder.Builder[expr.TypeSql]):    
    _internal: expr.TypeSql

    def __init__(self):
        self._internal = expr.TypeSql()        
        self._internal.type_val = "sql"

    def build(self) -> expr.TypeSql:
        """
        Builds the object.
        """
        return self._internal    
    
    def datasource(self, datasource: dashboard.DataSourceRef) -> typing.Self:    
        """
        The datasource
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def expression(self, expression: str) -> typing.Self:        
        if not len(expression) >= 1:
            raise ValueError("len(expression) must be >= 1")
        self._internal.expression = expression
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        self._internal.interval_ms = interval_ms
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        self._internal.max_data_points = max_data_points
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[expr.ExprTypeSqlResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        result_assertions_resource = result_assertions.build()
        self._internal.result_assertions = result_assertions_resource
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[expr.ExprTypeSqlTimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        time_range_resource = time_range.build()
        self._internal.time_range = time_range_resource
    
        return self
    

class ExprTypeMathResultAssertions(cogbuilder.Builder[expr.ExprTypeMathResultAssertions]):    
    _internal: expr.ExprTypeMathResultAssertions

    def __init__(self):
        self._internal = expr.ExprTypeMathResultAssertions()

    def build(self) -> expr.ExprTypeMathResultAssertions:
        """
        Builds the object.
        """
        return self._internal    
    
    def max_frames(self, max_frames: int) -> typing.Self:    
        """
        Maximum frame count
        """
            
        self._internal.max_frames = max_frames
    
        return self
    
    def type_val(self, type_val: typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]) -> typing.Self:    
        """
        Type asserts that the frame matches a known type structure.
        Possible enum values:
         - `""` 
         - `"timeseries-wide"` 
         - `"timeseries-long"` 
         - `"timeseries-many"` 
         - `"timeseries-multi"` 
         - `"directory-listing"` 
         - `"table"` 
         - `"numeric-wide"` 
         - `"numeric-multi"` 
         - `"numeric-long"` 
         - `"log-lines"` 
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def type_version(self, type_version: list[int]) -> typing.Self:    
        """
        TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
        contract documentation https://grafana.github.io/dataplane/contract/.
        """
            
        self._internal.type_version = type_version
    
        return self
    

class ExprTypeMathTimeRange(cogbuilder.Builder[expr.ExprTypeMathTimeRange]):    
    _internal: expr.ExprTypeMathTimeRange

    def __init__(self):
        self._internal = expr.ExprTypeMathTimeRange()

    def build(self) -> expr.ExprTypeMathTimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: str) -> typing.Self:    
        """
        From is the start time of the query.
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:    
        """
        To is the end time of the query.
        """
            
        self._internal.to = to
    
        return self
    

class ExprTypeReduceResultAssertions(cogbuilder.Builder[expr.ExprTypeReduceResultAssertions]):    
    _internal: expr.ExprTypeReduceResultAssertions

    def __init__(self):
        self._internal = expr.ExprTypeReduceResultAssertions()

    def build(self) -> expr.ExprTypeReduceResultAssertions:
        """
        Builds the object.
        """
        return self._internal    
    
    def max_frames(self, max_frames: int) -> typing.Self:    
        """
        Maximum frame count
        """
            
        self._internal.max_frames = max_frames
    
        return self
    
    def type_val(self, type_val: typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]) -> typing.Self:    
        """
        Type asserts that the frame matches a known type structure.
        Possible enum values:
         - `""` 
         - `"timeseries-wide"` 
         - `"timeseries-long"` 
         - `"timeseries-many"` 
         - `"timeseries-multi"` 
         - `"directory-listing"` 
         - `"table"` 
         - `"numeric-wide"` 
         - `"numeric-multi"` 
         - `"numeric-long"` 
         - `"log-lines"` 
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def type_version(self, type_version: list[int]) -> typing.Self:    
        """
        TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
        contract documentation https://grafana.github.io/dataplane/contract/.
        """
            
        self._internal.type_version = type_version
    
        return self
    

class ExprTypeReduceSettings(cogbuilder.Builder[expr.ExprTypeReduceSettings]):    
    _internal: expr.ExprTypeReduceSettings

    def __init__(self):
        self._internal = expr.ExprTypeReduceSettings()

    def build(self) -> expr.ExprTypeReduceSettings:
        """
        Builds the object.
        """
        return self._internal    
    
    def mode(self, mode: typing.Literal["dropNN", "replaceNN"]) -> typing.Self:    
        """
        Non-number reduce behavior
        Possible enum values:
         - `"dropNN"` Drop non-numbers
         - `"replaceNN"` Replace non-numbers
        """
            
        self._internal.mode = mode
    
        return self
    
    def replace_with_value(self, replace_with_value: float) -> typing.Self:    
        """
        Only valid when mode is replace
        """
            
        self._internal.replace_with_value = replace_with_value
    
        return self
    

class ExprTypeReduceTimeRange(cogbuilder.Builder[expr.ExprTypeReduceTimeRange]):    
    _internal: expr.ExprTypeReduceTimeRange

    def __init__(self):
        self._internal = expr.ExprTypeReduceTimeRange()

    def build(self) -> expr.ExprTypeReduceTimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: str) -> typing.Self:    
        """
        From is the start time of the query.
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:    
        """
        To is the end time of the query.
        """
            
        self._internal.to = to
    
        return self
    

class ExprTypeResampleResultAssertions(cogbuilder.Builder[expr.ExprTypeResampleResultAssertions]):    
    _internal: expr.ExprTypeResampleResultAssertions

    def __init__(self):
        self._internal = expr.ExprTypeResampleResultAssertions()

    def build(self) -> expr.ExprTypeResampleResultAssertions:
        """
        Builds the object.
        """
        return self._internal    
    
    def max_frames(self, max_frames: int) -> typing.Self:    
        """
        Maximum frame count
        """
            
        self._internal.max_frames = max_frames
    
        return self
    
    def type_val(self, type_val: typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]) -> typing.Self:    
        """
        Type asserts that the frame matches a known type structure.
        Possible enum values:
         - `""` 
         - `"timeseries-wide"` 
         - `"timeseries-long"` 
         - `"timeseries-many"` 
         - `"timeseries-multi"` 
         - `"directory-listing"` 
         - `"table"` 
         - `"numeric-wide"` 
         - `"numeric-multi"` 
         - `"numeric-long"` 
         - `"log-lines"` 
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def type_version(self, type_version: list[int]) -> typing.Self:    
        """
        TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
        contract documentation https://grafana.github.io/dataplane/contract/.
        """
            
        self._internal.type_version = type_version
    
        return self
    

class ExprTypeResampleTimeRange(cogbuilder.Builder[expr.ExprTypeResampleTimeRange]):    
    _internal: expr.ExprTypeResampleTimeRange

    def __init__(self):
        self._internal = expr.ExprTypeResampleTimeRange()

    def build(self) -> expr.ExprTypeResampleTimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: str) -> typing.Self:    
        """
        From is the start time of the query.
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:    
        """
        To is the end time of the query.
        """
            
        self._internal.to = to
    
        return self
    

class ExprTypeClassicConditionsConditionsEvaluator(cogbuilder.Builder[expr.ExprTypeClassicConditionsConditionsEvaluator]):    
    _internal: expr.ExprTypeClassicConditionsConditionsEvaluator

    def __init__(self):
        self._internal = expr.ExprTypeClassicConditionsConditionsEvaluator()

    def build(self) -> expr.ExprTypeClassicConditionsConditionsEvaluator:
        """
        Builds the object.
        """
        return self._internal    
    
    def params(self, params: list[float]) -> typing.Self:        
        self._internal.params = params
    
        return self
    
    def type_val(self, type_val: str) -> typing.Self:    
        """
        e.g. "gt"
        """
            
        self._internal.type_val = type_val
    
        return self
    

class ExprTypeClassicConditionsConditionsOperator(cogbuilder.Builder[expr.ExprTypeClassicConditionsConditionsOperator]):    
    _internal: expr.ExprTypeClassicConditionsConditionsOperator

    def __init__(self):
        self._internal = expr.ExprTypeClassicConditionsConditionsOperator()

    def build(self) -> expr.ExprTypeClassicConditionsConditionsOperator:
        """
        Builds the object.
        """
        return self._internal    
    
    def type_val(self, type_val: typing.Literal["and", "or", "logic-or"]) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    

class ExprTypeClassicConditionsConditionsQuery(cogbuilder.Builder[expr.ExprTypeClassicConditionsConditionsQuery]):    
    _internal: expr.ExprTypeClassicConditionsConditionsQuery

    def __init__(self):
        self._internal = expr.ExprTypeClassicConditionsConditionsQuery()

    def build(self) -> expr.ExprTypeClassicConditionsConditionsQuery:
        """
        Builds the object.
        """
        return self._internal    
    
    def params(self, params: list[str]) -> typing.Self:        
        self._internal.params = params
    
        return self
    

class ExprTypeClassicConditionsConditionsReducer(cogbuilder.Builder[expr.ExprTypeClassicConditionsConditionsReducer]):    
    _internal: expr.ExprTypeClassicConditionsConditionsReducer

    def __init__(self):
        self._internal = expr.ExprTypeClassicConditionsConditionsReducer()

    def build(self) -> expr.ExprTypeClassicConditionsConditionsReducer:
        """
        Builds the object.
        """
        return self._internal    
    
    def type_val(self, type_val: str) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    

class ExprTypeClassicConditionsConditions(cogbuilder.Builder[expr.ExprTypeClassicConditionsConditions]):    
    _internal: expr.ExprTypeClassicConditionsConditions

    def __init__(self):
        self._internal = expr.ExprTypeClassicConditionsConditions()

    def build(self) -> expr.ExprTypeClassicConditionsConditions:
        """
        Builds the object.
        """
        return self._internal    
    
    def evaluator(self, evaluator: cogbuilder.Builder[expr.ExprTypeClassicConditionsConditionsEvaluator]) -> typing.Self:        
        evaluator_resource = evaluator.build()
        self._internal.evaluator = evaluator_resource
    
        return self
    
    def operator(self, operator: cogbuilder.Builder[expr.ExprTypeClassicConditionsConditionsOperator]) -> typing.Self:        
        operator_resource = operator.build()
        self._internal.operator = operator_resource
    
        return self
    
    def query(self, query: cogbuilder.Builder[expr.ExprTypeClassicConditionsConditionsQuery]) -> typing.Self:        
        query_resource = query.build()
        self._internal.query = query_resource
    
        return self
    
    def reducer(self, reducer: cogbuilder.Builder[expr.ExprTypeClassicConditionsConditionsReducer]) -> typing.Self:        
        reducer_resource = reducer.build()
        self._internal.reducer = reducer_resource
    
        return self
    

class ExprTypeClassicConditionsResultAssertions(cogbuilder.Builder[expr.ExprTypeClassicConditionsResultAssertions]):    
    _internal: expr.ExprTypeClassicConditionsResultAssertions

    def __init__(self):
        self._internal = expr.ExprTypeClassicConditionsResultAssertions()

    def build(self) -> expr.ExprTypeClassicConditionsResultAssertions:
        """
        Builds the object.
        """
        return self._internal    
    
    def max_frames(self, max_frames: int) -> typing.Self:    
        """
        Maximum frame count
        """
            
        self._internal.max_frames = max_frames
    
        return self
    
    def type_val(self, type_val: typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]) -> typing.Self:    
        """
        Type asserts that the frame matches a known type structure.
        Possible enum values:
         - `""` 
         - `"timeseries-wide"` 
         - `"timeseries-long"` 
         - `"timeseries-many"` 
         - `"timeseries-multi"` 
         - `"directory-listing"` 
         - `"table"` 
         - `"numeric-wide"` 
         - `"numeric-multi"` 
         - `"numeric-long"` 
         - `"log-lines"` 
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def type_version(self, type_version: list[int]) -> typing.Self:    
        """
        TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
        contract documentation https://grafana.github.io/dataplane/contract/.
        """
            
        self._internal.type_version = type_version
    
        return self
    

class ExprTypeClassicConditionsTimeRange(cogbuilder.Builder[expr.ExprTypeClassicConditionsTimeRange]):    
    _internal: expr.ExprTypeClassicConditionsTimeRange

    def __init__(self):
        self._internal = expr.ExprTypeClassicConditionsTimeRange()

    def build(self) -> expr.ExprTypeClassicConditionsTimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: str) -> typing.Self:    
        """
        From is the start time of the query.
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:    
        """
        To is the end time of the query.
        """
            
        self._internal.to = to
    
        return self
    

class ExprTypeThresholdConditionsEvaluator(cogbuilder.Builder[expr.ExprTypeThresholdConditionsEvaluator]):    
    _internal: expr.ExprTypeThresholdConditionsEvaluator

    def __init__(self):
        self._internal = expr.ExprTypeThresholdConditionsEvaluator()

    def build(self) -> expr.ExprTypeThresholdConditionsEvaluator:
        """
        Builds the object.
        """
        return self._internal    
    
    def params(self, params: list[float]) -> typing.Self:        
        self._internal.params = params
    
        return self
    
    def type_val(self, type_val: typing.Literal["gt", "lt", "within_range", "outside_range"]) -> typing.Self:    
        """
        e.g. "gt"
        """
            
        self._internal.type_val = type_val
    
        return self
    

class ExprTypeThresholdConditionsUnloadEvaluator(cogbuilder.Builder[expr.ExprTypeThresholdConditionsUnloadEvaluator]):    
    _internal: expr.ExprTypeThresholdConditionsUnloadEvaluator

    def __init__(self):
        self._internal = expr.ExprTypeThresholdConditionsUnloadEvaluator()

    def build(self) -> expr.ExprTypeThresholdConditionsUnloadEvaluator:
        """
        Builds the object.
        """
        return self._internal    
    
    def params(self, params: list[float]) -> typing.Self:        
        self._internal.params = params
    
        return self
    
    def type_val(self, type_val: typing.Literal["gt", "lt", "within_range", "outside_range"]) -> typing.Self:    
        """
        e.g. "gt"
        """
            
        self._internal.type_val = type_val
    
        return self
    

class ExprTypeThresholdConditions(cogbuilder.Builder[expr.ExprTypeThresholdConditions]):    
    _internal: expr.ExprTypeThresholdConditions

    def __init__(self):
        self._internal = expr.ExprTypeThresholdConditions()

    def build(self) -> expr.ExprTypeThresholdConditions:
        """
        Builds the object.
        """
        return self._internal    
    
    def evaluator(self, evaluator: cogbuilder.Builder[expr.ExprTypeThresholdConditionsEvaluator]) -> typing.Self:        
        evaluator_resource = evaluator.build()
        self._internal.evaluator = evaluator_resource
    
        return self
    
    def loaded_dimensions(self, loaded_dimensions: object) -> typing.Self:        
        self._internal.loaded_dimensions = loaded_dimensions
    
        return self
    
    def unload_evaluator(self, unload_evaluator: cogbuilder.Builder[expr.ExprTypeThresholdConditionsUnloadEvaluator]) -> typing.Self:        
        unload_evaluator_resource = unload_evaluator.build()
        self._internal.unload_evaluator = unload_evaluator_resource
    
        return self
    

class ExprTypeThresholdResultAssertions(cogbuilder.Builder[expr.ExprTypeThresholdResultAssertions]):    
    _internal: expr.ExprTypeThresholdResultAssertions

    def __init__(self):
        self._internal = expr.ExprTypeThresholdResultAssertions()

    def build(self) -> expr.ExprTypeThresholdResultAssertions:
        """
        Builds the object.
        """
        return self._internal    
    
    def max_frames(self, max_frames: int) -> typing.Self:    
        """
        Maximum frame count
        """
            
        self._internal.max_frames = max_frames
    
        return self
    
    def type_val(self, type_val: typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]) -> typing.Self:    
        """
        Type asserts that the frame matches a known type structure.
        Possible enum values:
         - `""` 
         - `"timeseries-wide"` 
         - `"timeseries-long"` 
         - `"timeseries-many"` 
         - `"timeseries-multi"` 
         - `"directory-listing"` 
         - `"table"` 
         - `"numeric-wide"` 
         - `"numeric-multi"` 
         - `"numeric-long"` 
         - `"log-lines"` 
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def type_version(self, type_version: list[int]) -> typing.Self:    
        """
        TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
        contract documentation https://grafana.github.io/dataplane/contract/.
        """
            
        self._internal.type_version = type_version
    
        return self
    

class ExprTypeThresholdTimeRange(cogbuilder.Builder[expr.ExprTypeThresholdTimeRange]):    
    _internal: expr.ExprTypeThresholdTimeRange

    def __init__(self):
        self._internal = expr.ExprTypeThresholdTimeRange()

    def build(self) -> expr.ExprTypeThresholdTimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: str) -> typing.Self:    
        """
        From is the start time of the query.
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:    
        """
        To is the end time of the query.
        """
            
        self._internal.to = to
    
        return self
    

class ExprTypeSqlResultAssertions(cogbuilder.Builder[expr.ExprTypeSqlResultAssertions]):    
    _internal: expr.ExprTypeSqlResultAssertions

    def __init__(self):
        self._internal = expr.ExprTypeSqlResultAssertions()

    def build(self) -> expr.ExprTypeSqlResultAssertions:
        """
        Builds the object.
        """
        return self._internal    
    
    def max_frames(self, max_frames: int) -> typing.Self:    
        """
        Maximum frame count
        """
            
        self._internal.max_frames = max_frames
    
        return self
    
    def type_val(self, type_val: typing.Literal["", "timeseries-wide", "timeseries-long", "timeseries-many", "timeseries-multi", "directory-listing", "table", "numeric-wide", "numeric-multi", "numeric-long", "log-lines"]) -> typing.Self:    
        """
        Type asserts that the frame matches a known type structure.
        Possible enum values:
         - `""` 
         - `"timeseries-wide"` 
         - `"timeseries-long"` 
         - `"timeseries-many"` 
         - `"timeseries-multi"` 
         - `"directory-listing"` 
         - `"table"` 
         - `"numeric-wide"` 
         - `"numeric-multi"` 
         - `"numeric-long"` 
         - `"log-lines"` 
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def type_version(self, type_version: list[int]) -> typing.Self:    
        """
        TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
        contract documentation https://grafana.github.io/dataplane/contract/.
        """
            
        self._internal.type_version = type_version
    
        return self
    

class ExprTypeSqlTimeRange(cogbuilder.Builder[expr.ExprTypeSqlTimeRange]):    
    _internal: expr.ExprTypeSqlTimeRange

    def __init__(self):
        self._internal = expr.ExprTypeSqlTimeRange()

    def build(self) -> expr.ExprTypeSqlTimeRange:
        """
        Builds the object.
        """
        return self._internal    
    
    def from_val(self, from_val: str) -> typing.Self:    
        """
        From is the start time of the query.
        """
            
        self._internal.from_val = from_val
    
        return self
    
    def to(self, to: str) -> typing.Self:    
        """
        To is the end time of the query.
        """
            
        self._internal.to = to
    
        return self
    
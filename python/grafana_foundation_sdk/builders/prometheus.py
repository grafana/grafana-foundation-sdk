# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import dashboardv2
from ..models import prometheus
from ..models import common
from ..models import dashboardv2beta1


class QueryV2(cogbuilder.Builder[dashboardv2.DataQueryKind]):
    _internal: dashboardv2.DataQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2.DataQueryKind()        
        self._internal.kind = "DataQuery"        
        self._internal.group = "prometheus"

    def build(self) -> dashboardv2.DataQueryKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def version(self, version: str) -> typing.Self:    
        self._internal.version = version
    
        return self
    
    def labels(self, labels: dict[str, str]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def datasource(self, ref: cogbuilder.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) -> typing.Self:    
        """
        New type for datasource reference
        Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
        """
            
        ref_resource = ref.build()
        self._internal.datasource = ref_resource
    
        return self
    
    def adhoc_filters(self, adhoc_filters: list[cogbuilder.Builder[prometheus.AdhocFilters]]) -> typing.Self:    
        """
        Additional Ad-hoc filters that take precedence over Scope on conflict.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        adhoc_filters_resources = [r1.build() for r1 in adhoc_filters]
        self._internal.spec.adhoc_filters = adhoc_filters_resources
    
        return self
    
    def editor_mode(self, editor_mode: prometheus.QueryEditorMode) -> typing.Self:    
        """
        what we should show in the editor
        Possible enum values:
         - `"builder"` 
         - `"code"` 
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.editor_mode = editor_mode
    
        return self
    
    def exemplar(self, exemplar: bool) -> typing.Self:    
        """
        Execute an additional query to identify interesting raw samples relevant for the given expr
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.exemplar = exemplar
    
        return self
    
    def expr(self, expr: str) -> typing.Self:    
        """
        The actual expression/query that will be evaluated by Prometheus
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.expr = expr
    
        return self
    
    def format(self, format_val: prometheus.PromQueryFormat) -> typing.Self:    
        """
        The response format
        Possible enum values:
         - `"time_series"` 
         - `"table"` 
         - `"heatmap"` 
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.format_val = format_val
    
        return self
    
    def group_by_keys(self, group_by_keys: list[str]) -> typing.Self:    
        """
        Group By parameters to apply to aggregate expressions in the query
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.group_by_keys = group_by_keys
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.hide = hide
    
        return self
    
    def instant(self, instant: bool) -> typing.Self:    
        """
        Returns only the latest value that Prometheus has scraped for the requested time series
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.instant = instant
    
        return self
    
    def interval_factor(self, interval_factor: int) -> typing.Self:    
        """
        Used to specify how many times to divide max data points by. We use max data points under query options
        See https://github.com/grafana/grafana/issues/48081
        Deprecated: use interval
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.interval_factor = interval_factor
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.interval_ms = interval_ms
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:    
        """
        Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.legend_format = legend_format
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.max_data_points = max_data_points
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.query_type = query_type
    
        return self
    
    def range(self, range_val: bool) -> typing.Self:    
        """
        Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.range_val = range_val
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[prometheus.ResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        result_assertions_resource = result_assertions.build()
        self._internal.spec.result_assertions = result_assertions_resource
    
        return self
    
    def scopes(self, scopes: list[cogbuilder.Builder[prometheus.Scopes]]) -> typing.Self:    
        """
        A set of filters applied to apply to the query
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        scopes_resources = [r1.build() for r1 in scopes]
        self._internal.spec.scopes = scopes_resources
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[prometheus.TimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        time_range_resource = time_range.build()
        self._internal.spec.time_range = time_range_resource
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        """
        An additional lower limit for the step parameter of the Prometheus query and for the
        `$__interval` and `$__rate_interval` variables.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.interval = interval
    
        return self
    


class AdhocFilters(cogbuilder.Builder[prometheus.AdhocFilters]):
    _internal: prometheus.AdhocFilters

    def __init__(self) -> None:
        self._internal = prometheus.AdhocFilters()

    def build(self) -> prometheus.AdhocFilters:
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
    


class ResultAssertions(cogbuilder.Builder[prometheus.ResultAssertions]):
    _internal: prometheus.ResultAssertions

    def __init__(self) -> None:
        self._internal = prometheus.ResultAssertions()

    def build(self) -> prometheus.ResultAssertions:
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
    
    def type(self, type_val: prometheus.ResultAssertionsType) -> typing.Self:    
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
    


class ScopesFilters(cogbuilder.Builder[prometheus.ScopesFilters]):
    _internal: prometheus.ScopesFilters

    def __init__(self) -> None:
        self._internal = prometheus.ScopesFilters()

    def build(self) -> prometheus.ScopesFilters:
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
    


class Scopes(cogbuilder.Builder[prometheus.Scopes]):
    _internal: prometheus.Scopes

    def __init__(self) -> None:
        self._internal = prometheus.Scopes()

    def build(self) -> prometheus.Scopes:
        """
        Builds the object.
        """
        return self._internal    
    
    def default_path(self, default_path: list[str]) -> typing.Self:    
        self._internal.default_path = default_path
    
        return self
    
    def filters(self, filters: list[cogbuilder.Builder[prometheus.ScopesFilters]]) -> typing.Self:    
        filters_resources = [r1.build() for r1 in filters]
        self._internal.filters = filters_resources
    
        return self
    
    def title(self, title: str) -> typing.Self:    
        self._internal.title = title
    
        return self
    


class TimeRange(cogbuilder.Builder[prometheus.TimeRange]):
    _internal: prometheus.TimeRange

    def __init__(self) -> None:
        self._internal = prometheus.TimeRange()

    def build(self) -> prometheus.TimeRange:
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
    


class Dataquery(cogbuilder.Builder[prometheus.Dataquery]):
    _internal: prometheus.Dataquery

    def __init__(self) -> None:
        self._internal = prometheus.Dataquery()

    def build(self) -> prometheus.Dataquery:
        """
        Builds the object.
        """
        return self._internal    
    
    def adhoc_filters(self, adhoc_filters: list[cogbuilder.Builder[prometheus.AdhocFilters]]) -> typing.Self:    
        """
        Additional Ad-hoc filters that take precedence over Scope on conflict.
        """
            
        adhoc_filters_resources = [r1.build() for r1 in adhoc_filters]
        self._internal.adhoc_filters = adhoc_filters_resources
    
        return self
    
    def datasource(self, datasource: common.DataSourceRef) -> typing.Self:    
        """
        The datasource
        """
            
        self._internal.datasource = datasource
    
        return self
    
    def editor_mode(self, editor_mode: prometheus.QueryEditorMode) -> typing.Self:    
        """
        what we should show in the editor
        Possible enum values:
         - `"builder"` 
         - `"code"` 
        """
            
        self._internal.editor_mode = editor_mode
    
        return self
    
    def exemplar(self, exemplar: bool) -> typing.Self:    
        """
        Execute an additional query to identify interesting raw samples relevant for the given expr
        """
            
        self._internal.exemplar = exemplar
    
        return self
    
    def expr(self, expr: str) -> typing.Self:    
        """
        The actual expression/query that will be evaluated by Prometheus
        """
            
        self._internal.expr = expr
    
        return self
    
    def format(self, format_val: prometheus.PromQueryFormat) -> typing.Self:    
        """
        The response format
        Possible enum values:
         - `"time_series"` 
         - `"table"` 
         - `"heatmap"` 
        """
            
        self._internal.format_val = format_val
    
        return self
    
    def group_by_keys(self, group_by_keys: list[str]) -> typing.Self:    
        """
        Group By parameters to apply to aggregate expressions in the query
        """
            
        self._internal.group_by_keys = group_by_keys
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def instant(self) -> typing.Self:    
        """
        Returns only the latest value that Prometheus has scraped for the requested time series
        """
            
        self._internal.instant = True    
        self._internal.range_val = False
    
        return self
    
    def interval_factor(self, interval_factor: int) -> typing.Self:    
        """
        Used to specify how many times to divide max data points by. We use max data points under query options
        See https://github.com/grafana/grafana/issues/48081
        Deprecated: use interval
        """
            
        self._internal.interval_factor = interval_factor
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        self._internal.interval_ms = interval_ms
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:    
        """
        Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
        """
            
        self._internal.legend_format = legend_format
    
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
    
    def range(self) -> typing.Self:    
        """
        Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
        """
            
        self._internal.range_val = True    
        self._internal.instant = False
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[prometheus.ResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        result_assertions_resource = result_assertions.build()
        self._internal.result_assertions = result_assertions_resource
    
        return self
    
    def scopes(self, scopes: list[cogbuilder.Builder[prometheus.Scopes]]) -> typing.Self:    
        """
        A set of filters applied to apply to the query
        """
            
        scopes_resources = [r1.build() for r1 in scopes]
        self._internal.scopes = scopes_resources
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[prometheus.TimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        time_range_resource = time_range.build()
        self._internal.time_range = time_range_resource
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        """
        An additional lower limit for the step parameter of the Prometheus query and for the
        `$__interval` and `$__rate_interval` variables.
        """
            
        self._internal.interval = interval
    
        return self
    
    def range_and_instant(self) -> typing.Self:    
        self._internal.range_val = True    
        self._internal.instant = True
    
        return self
    


class Query(cogbuilder.Builder[dashboardv2beta1.DataQueryKind]):
    _internal: dashboardv2beta1.DataQueryKind

    def __init__(self) -> None:
        self._internal = dashboardv2beta1.DataQueryKind()        
        self._internal.kind = "DataQuery"        
        self._internal.group = "prometheus"

    def build(self) -> dashboardv2beta1.DataQueryKind:
        """
        Builds the object.
        """
        return self._internal    
    
    def version(self, version: str) -> typing.Self:    
        self._internal.version = version
    
        return self
    
    def labels(self, labels: dict[str, str]) -> typing.Self:    
        self._internal.labels = labels
    
        return self
    
    def datasource(self, ref: cogbuilder.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) -> typing.Self:    
        """
        New type for datasource reference
        Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
        """
            
        ref_resource = ref.build()
        self._internal.datasource = ref_resource
    
        return self
    
    def adhoc_filters(self, adhoc_filters: list[cogbuilder.Builder[prometheus.AdhocFilters]]) -> typing.Self:    
        """
        Additional Ad-hoc filters that take precedence over Scope on conflict.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        adhoc_filters_resources = [r1.build() for r1 in adhoc_filters]
        self._internal.spec.adhoc_filters = adhoc_filters_resources
    
        return self
    
    def editor_mode(self, editor_mode: prometheus.QueryEditorMode) -> typing.Self:    
        """
        what we should show in the editor
        Possible enum values:
         - `"builder"` 
         - `"code"` 
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.editor_mode = editor_mode
    
        return self
    
    def exemplar(self, exemplar: bool) -> typing.Self:    
        """
        Execute an additional query to identify interesting raw samples relevant for the given expr
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.exemplar = exemplar
    
        return self
    
    def expr(self, expr: str) -> typing.Self:    
        """
        The actual expression/query that will be evaluated by Prometheus
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.expr = expr
    
        return self
    
    def format(self, format_val: prometheus.PromQueryFormat) -> typing.Self:    
        """
        The response format
        Possible enum values:
         - `"time_series"` 
         - `"table"` 
         - `"heatmap"` 
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.format_val = format_val
    
        return self
    
    def group_by_keys(self, group_by_keys: list[str]) -> typing.Self:    
        """
        Group By parameters to apply to aggregate expressions in the query
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.group_by_keys = group_by_keys
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        NOTE: this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.hide = hide
    
        return self
    
    def instant(self) -> typing.Self:    
        """
        Returns only the latest value that Prometheus has scraped for the requested time series
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.instant = True    
        self._internal.spec.range_val = False
    
        return self
    
    def interval_factor(self, interval_factor: int) -> typing.Self:    
        """
        Used to specify how many times to divide max data points by. We use max data points under query options
        See https://github.com/grafana/grafana/issues/48081
        Deprecated: use interval
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.interval_factor = interval_factor
    
        return self
    
    def interval_ms(self, interval_ms: float) -> typing.Self:    
        """
        Interval is the suggested duration between time points in a time series query.
        NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
        from the interval required to fill a pixels in the visualization
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.interval_ms = interval_ms
    
        return self
    
    def legend_format(self, legend_format: str) -> typing.Self:    
        """
        Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.legend_format = legend_format
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.max_data_points = max_data_points
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.query_type = query_type
    
        return self
    
    def range(self) -> typing.Self:    
        """
        Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.range_val = True    
        self._internal.spec.instant = False
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[prometheus.ResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        result_assertions_resource = result_assertions.build()
        self._internal.spec.result_assertions = result_assertions_resource
    
        return self
    
    def scopes(self, scopes: list[cogbuilder.Builder[prometheus.Scopes]]) -> typing.Self:    
        """
        A set of filters applied to apply to the query
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        scopes_resources = [r1.build() for r1 in scopes]
        self._internal.spec.scopes = scopes_resources
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[prometheus.TimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        time_range_resource = time_range.build()
        self._internal.spec.time_range = time_range_resource
    
        return self
    
    def interval(self, interval: str) -> typing.Self:    
        """
        An additional lower limit for the step parameter of the Prometheus query and for the
        `$__interval` and `$__rate_interval` variables.
        """
            
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.interval = interval
    
        return self
    
    def range_and_instant(self) -> typing.Self:    
        if self._internal.spec is None:
            self._internal.spec = prometheus.Dataquery()
        assert isinstance(self._internal.spec, prometheus.Dataquery)
        self._internal.spec.range_val = True    
        self._internal.spec.instant = True
    
        return self
    

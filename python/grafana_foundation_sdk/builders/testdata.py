# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import testdata


class CSVWave(cogbuilder.Builder[testdata.CSVWave]):    
    _internal: testdata.CSVWave

    def __init__(self):
        self._internal = testdata.CSVWave()

    def build(self) -> testdata.CSVWave:
        return self._internal    
    
    def labels(self, labels: str) -> typing.Self:        
        self._internal.labels = labels
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def time_step(self, time_step: int) -> typing.Self:        
        self._internal.time_step = time_step
    
        return self
    
    def values_csv(self, values_csv: str) -> typing.Self:        
        self._internal.values_csv = values_csv
    
        return self
    

class Datasource(cogbuilder.Builder[testdata.Datasource]):    
    _internal: testdata.Datasource

    def __init__(self):
        self._internal = testdata.Datasource()

    def build(self) -> testdata.Datasource:
        return self._internal    
    
    def type_val(self, type_val: str) -> typing.Self:    
        """
        The datasource plugin type
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def uid(self, uid: str) -> typing.Self:    
        """
        Datasource UID
        """
            
        self._internal.uid = uid
    
        return self
    

class NodesQuery(cogbuilder.Builder[testdata.NodesQuery]):    
    _internal: testdata.NodesQuery

    def __init__(self):
        self._internal = testdata.NodesQuery()

    def build(self) -> testdata.NodesQuery:
        return self._internal    
    
    def count(self, count: int) -> typing.Self:        
        self._internal.count = count
    
        return self
    
    def seed(self, seed: int) -> typing.Self:        
        self._internal.seed = seed
    
        return self
    
    def type_val(self, type_val: typing.Literal["random", "random edges", "response_medium", "response_small", "feature_showcase"]) -> typing.Self:    
        """
        Possible enum values:
         - `"random"` 
         - `"random edges"` 
         - `"response_medium"` 
         - `"response_small"` 
         - `"feature_showcase"` 
        """
            
        self._internal.type_val = type_val
    
        return self
    

class PulseWaveQuery(cogbuilder.Builder[testdata.PulseWaveQuery]):    
    _internal: testdata.PulseWaveQuery

    def __init__(self):
        self._internal = testdata.PulseWaveQuery()

    def build(self) -> testdata.PulseWaveQuery:
        return self._internal    
    
    def off_count(self, off_count: int) -> typing.Self:        
        self._internal.off_count = off_count
    
        return self
    
    def off_value(self, off_value: float) -> typing.Self:        
        self._internal.off_value = off_value
    
        return self
    
    def on_count(self, on_count: int) -> typing.Self:        
        self._internal.on_count = on_count
    
        return self
    
    def on_value(self, on_value: float) -> typing.Self:        
        self._internal.on_value = on_value
    
        return self
    
    def time_step(self, time_step: int) -> typing.Self:        
        self._internal.time_step = time_step
    
        return self
    

class ResultAssertions(cogbuilder.Builder[testdata.ResultAssertions]):    
    _internal: testdata.ResultAssertions

    def __init__(self):
        self._internal = testdata.ResultAssertions()

    def build(self) -> testdata.ResultAssertions:
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
    

class Key(cogbuilder.Builder[testdata.Key]):    
    _internal: testdata.Key

    def __init__(self):
        self._internal = testdata.Key()

    def build(self) -> testdata.Key:
        return self._internal    
    
    def tick(self, tick: float) -> typing.Self:        
        self._internal.tick = tick
    
        return self
    
    def type_val(self, type_val: str) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def uid(self, uid: str) -> typing.Self:        
        self._internal.uid = uid
    
        return self
    

class SimulationQuery(cogbuilder.Builder[testdata.SimulationQuery]):    
    _internal: testdata.SimulationQuery

    def __init__(self):
        self._internal = testdata.SimulationQuery()

    def build(self) -> testdata.SimulationQuery:
        return self._internal    
    
    def config(self, config: object) -> typing.Self:        
        self._internal.config = config
    
        return self
    
    def key(self, key: cogbuilder.Builder[testdata.Key]) -> typing.Self:        
        key_resource = key.build()
        self._internal.key = key_resource
    
        return self
    
    def last(self, last: bool) -> typing.Self:        
        self._internal.last = last
    
        return self
    
    def stream(self, stream: bool) -> typing.Self:        
        self._internal.stream = stream
    
        return self
    

class StreamingQuery(cogbuilder.Builder[testdata.StreamingQuery]):    
    _internal: testdata.StreamingQuery

    def __init__(self):
        self._internal = testdata.StreamingQuery()

    def build(self) -> testdata.StreamingQuery:
        return self._internal    
    
    def bands(self, bands: int) -> typing.Self:        
        self._internal.bands = bands
    
        return self
    
    def noise(self, noise: float) -> typing.Self:        
        self._internal.noise = noise
    
        return self
    
    def speed(self, speed: float) -> typing.Self:        
        self._internal.speed = speed
    
        return self
    
    def spread(self, spread: float) -> typing.Self:        
        self._internal.spread = spread
    
        return self
    
    def type_val(self, type_val: typing.Literal["fetch", "logs", "signal", "traces"]) -> typing.Self:    
        """
        Possible enum values:
         - `"fetch"` 
         - `"logs"` 
         - `"signal"` 
         - `"traces"` 
        """
            
        self._internal.type_val = type_val
    
        return self
    
    def url(self, url: str) -> typing.Self:        
        self._internal.url = url
    
        return self
    

class TimeRange(cogbuilder.Builder[testdata.TimeRange]):    
    _internal: testdata.TimeRange

    def __init__(self):
        self._internal = testdata.TimeRange()

    def build(self) -> testdata.TimeRange:
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
    

class USAQuery(cogbuilder.Builder[testdata.USAQuery]):    
    _internal: testdata.USAQuery

    def __init__(self):
        self._internal = testdata.USAQuery()

    def build(self) -> testdata.USAQuery:
        return self._internal    
    
    def fields(self, fields: list[str]) -> typing.Self:        
        self._internal.fields = fields
    
        return self
    
    def mode(self, mode: str) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    
    def period(self, period: str) -> typing.Self:        
        self._internal.period = period
    
        return self
    
    def states(self, states: list[str]) -> typing.Self:        
        self._internal.states = states
    
        return self
    

class Dataquery(cogbuilder.Builder[testdata.Dataquery]):    
    _internal: testdata.Dataquery

    def __init__(self):
        self._internal = testdata.Dataquery()

    def build(self) -> testdata.Dataquery:
        return self._internal    
    
    def alias(self, alias: str) -> typing.Self:        
        self._internal.alias = alias
    
        return self
    
    def channel(self, channel: str) -> typing.Self:    
        """
        Used for live query
        """
            
        self._internal.channel = channel
    
        return self
    
    def csv_content(self, csv_content: str) -> typing.Self:        
        self._internal.csv_content = csv_content
    
        return self
    
    def csv_file_name(self, csv_file_name: str) -> typing.Self:        
        self._internal.csv_file_name = csv_file_name
    
        return self
    
    def csv_wave(self, csv_wave: list[cogbuilder.Builder[testdata.CSVWave]]) -> typing.Self:        
        csv_wave_resources = [r1.build() for r1 in csv_wave]
        self._internal.csv_wave = csv_wave_resources
    
        return self
    
    def datasource(self, datasource: cogbuilder.Builder[testdata.Datasource]) -> typing.Self:    
        """
        The datasource
        """
            
        datasource_resource = datasource.build()
        self._internal.datasource = datasource_resource
    
        return self
    
    def drop_percent(self, drop_percent: float) -> typing.Self:    
        """
        Drop percentage (the chance we will lose a point 0-100)
        """
            
        self._internal.drop_percent = drop_percent
    
        return self
    
    def error_type(self, error_type: typing.Literal["frontend_exception", "frontend_observable", "server_panic"]) -> typing.Self:    
        """
        Possible enum values:
         - `"frontend_exception"` 
         - `"frontend_observable"` 
         - `"server_panic"` 
        """
            
        self._internal.error_type = error_type
    
        return self
    
    def flamegraph_diff(self, flamegraph_diff: bool) -> typing.Self:        
        self._internal.flamegraph_diff = flamegraph_diff
    
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
    
    def labels(self, labels: str) -> typing.Self:        
        self._internal.labels = labels
    
        return self
    
    def level_column(self, level_column: bool) -> typing.Self:        
        self._internal.level_column = level_column
    
        return self
    
    def lines(self, lines: int) -> typing.Self:        
        self._internal.lines = lines
    
        return self
    
    def max_val(self, max_val: float) -> typing.Self:        
        self._internal.max_val = max_val
    
        return self
    
    def max_data_points(self, max_data_points: int) -> typing.Self:    
        """
        MaxDataPoints is the maximum number of data points that should be returned from a time series query.
        NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
        from the number of pixels visible in a visualization
        """
            
        self._internal.max_data_points = max_data_points
    
        return self
    
    def min_val(self, min_val: float) -> typing.Self:        
        self._internal.min_val = min_val
    
        return self
    
    def nodes(self, nodes: cogbuilder.Builder[testdata.NodesQuery]) -> typing.Self:        
        nodes_resource = nodes.build()
        self._internal.nodes = nodes_resource
    
        return self
    
    def noise(self, noise: float) -> typing.Self:        
        self._internal.noise = noise
    
        return self
    
    def points(self, points: list[list[object]]) -> typing.Self:        
        self._internal.points = points
    
        return self
    
    def pulse_wave(self, pulse_wave: cogbuilder.Builder[testdata.PulseWaveQuery]) -> typing.Self:        
        pulse_wave_resource = pulse_wave.build()
        self._internal.pulse_wave = pulse_wave_resource
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        QueryType is an optional identifier for the type of query.
        It can be used to distinguish different types of queries.
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def raw_frame_content(self, raw_frame_content: str) -> typing.Self:        
        self._internal.raw_frame_content = raw_frame_content
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        RefID is the unique identifier of the query, set by the frontend call.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def result_assertions(self, result_assertions: cogbuilder.Builder[testdata.ResultAssertions]) -> typing.Self:    
        """
        Optionally define expected query result behavior
        """
            
        result_assertions_resource = result_assertions.build()
        self._internal.result_assertions = result_assertions_resource
    
        return self
    
    def scenario_id(self, scenario_id: typing.Literal["annotations", "arrow", "csv_content", "csv_file", "csv_metric_values", "datapoints_outside_range", "exponential_heatmap_bucket_data", "flame_graph", "grafana_api", "linear_heatmap_bucket_data", "live", "logs", "manual_entry", "no_data_points", "node_graph", "predictable_csv_wave", "predictable_pulse", "random_walk", "random_walk_table", "random_walk_with_error", "raw_frame", "server_error_500", "simulation", "slow_query", "streaming_client", "table_static", "trace", "usa", "variables-query"]) -> typing.Self:    
        """
        Possible enum values:
         - `"annotations"` 
         - `"arrow"` 
         - `"csv_content"` 
         - `"csv_file"` 
         - `"csv_metric_values"` 
         - `"datapoints_outside_range"` 
         - `"exponential_heatmap_bucket_data"` 
         - `"flame_graph"` 
         - `"grafana_api"` 
         - `"linear_heatmap_bucket_data"` 
         - `"live"` 
         - `"logs"` 
         - `"manual_entry"` 
         - `"no_data_points"` 
         - `"node_graph"` 
         - `"predictable_csv_wave"` 
         - `"predictable_pulse"` 
         - `"random_walk"` 
         - `"random_walk_table"` 
         - `"random_walk_with_error"` 
         - `"raw_frame"` 
         - `"server_error_500"` 
         - `"simulation"` 
         - `"slow_query"` 
         - `"streaming_client"` 
         - `"table_static"` 
         - `"trace"` 
         - `"usa"` 
         - `"variables-query"` 
        """
            
        self._internal.scenario_id = scenario_id
    
        return self
    
    def series_count(self, series_count: int) -> typing.Self:        
        self._internal.series_count = series_count
    
        return self
    
    def sim(self, sim: cogbuilder.Builder[testdata.SimulationQuery]) -> typing.Self:        
        sim_resource = sim.build()
        self._internal.sim = sim_resource
    
        return self
    
    def span_count(self, span_count: int) -> typing.Self:        
        self._internal.span_count = span_count
    
        return self
    
    def spread(self, spread: float) -> typing.Self:        
        self._internal.spread = spread
    
        return self
    
    def start_value(self, start_value: float) -> typing.Self:        
        self._internal.start_value = start_value
    
        return self
    
    def stream(self, stream: cogbuilder.Builder[testdata.StreamingQuery]) -> typing.Self:        
        stream_resource = stream.build()
        self._internal.stream = stream_resource
    
        return self
    
    def string_input(self, string_input: str) -> typing.Self:    
        """
        common parameter used by many query types
        """
            
        self._internal.string_input = string_input
    
        return self
    
    def time_range(self, time_range: cogbuilder.Builder[testdata.TimeRange]) -> typing.Self:    
        """
        TimeRange represents the query range
        NOTE: unlike generic /ds/query, we can now send explicit time values in each query
        NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
        """
            
        time_range_resource = time_range.build()
        self._internal.time_range = time_range_resource
    
        return self
    
    def usa(self, usa: cogbuilder.Builder[testdata.USAQuery]) -> typing.Self:        
        usa_resource = usa.build()
        self._internal.usa = usa_resource
    
        return self
    
    def with_nil(self, with_nil: bool) -> typing.Self:        
        self._internal.with_nil = with_nil
    
        return self
    
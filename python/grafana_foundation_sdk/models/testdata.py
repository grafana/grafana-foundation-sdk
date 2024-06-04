# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import variants as cogvariants
from ..cog import runtime as cogruntime


class CSVWave:
    labels: typing.Optional[str]
    name: typing.Optional[str]
    time_step: typing.Optional[int]
    values_csv: typing.Optional[str]

    def __init__(self, labels: typing.Optional[str] = None, name: typing.Optional[str] = None, time_step: typing.Optional[int] = None, values_csv: typing.Optional[str] = None):
        self.labels = labels
        self.name = name
        self.time_step = time_step
        self.values_csv = values_csv

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.labels is not None:
            payload["labels"] = self.labels
        if self.name is not None:
            payload["name"] = self.name
        if self.time_step is not None:
            payload["timeStep"] = self.time_step
        if self.values_csv is not None:
            payload["valuesCSV"] = self.values_csv
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "labels" in data:
            args["labels"] = data["labels"]
        if "name" in data:
            args["name"] = data["name"]
        if "timeStep" in data:
            args["time_step"] = data["timeStep"]
        if "valuesCSV" in data:
            args["values_csv"] = data["valuesCSV"]        

        return cls(**args)


class Datasource:
    # The datasource plugin type
    type_val: str
    # Datasource UID
    uid: typing.Optional[str]

    def __init__(self, type_val: str = "", uid: typing.Optional[str] = None):
        self.type_val = type_val
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
        }
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class NodesQuery:
    count: typing.Optional[int]
    seed: typing.Optional[int]
    # Possible enum values:
    #  - `"random"` 
    #  - `"random edges"` 
    #  - `"response_medium"` 
    #  - `"response_small"` 
    #  - `"feature_showcase"` 
    type_val: typing.Optional[typing.Literal["random", "random edges", "response_medium", "response_small", "feature_showcase"]]

    def __init__(self, count: typing.Optional[int] = None, seed: typing.Optional[int] = None, type_val: typing.Optional[typing.Literal["random", "random edges", "response_medium", "response_small", "feature_showcase"]] = None):
        self.count = count
        self.seed = seed
        self.type_val = type_val

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.count is not None:
            payload["count"] = self.count
        if self.seed is not None:
            payload["seed"] = self.seed
        if self.type_val is not None:
            payload["type"] = self.type_val
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "count" in data:
            args["count"] = data["count"]
        if "seed" in data:
            args["seed"] = data["seed"]
        if "type" in data:
            args["type_val"] = data["type"]        

        return cls(**args)


class PulseWaveQuery:
    off_count: typing.Optional[int]
    off_value: typing.Optional[float]
    on_count: typing.Optional[int]
    on_value: typing.Optional[float]
    time_step: typing.Optional[int]

    def __init__(self, off_count: typing.Optional[int] = None, off_value: typing.Optional[float] = None, on_count: typing.Optional[int] = None, on_value: typing.Optional[float] = None, time_step: typing.Optional[int] = None):
        self.off_count = off_count
        self.off_value = off_value
        self.on_count = on_count
        self.on_value = on_value
        self.time_step = time_step

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.off_count is not None:
            payload["offCount"] = self.off_count
        if self.off_value is not None:
            payload["offValue"] = self.off_value
        if self.on_count is not None:
            payload["onCount"] = self.on_count
        if self.on_value is not None:
            payload["onValue"] = self.on_value
        if self.time_step is not None:
            payload["timeStep"] = self.time_step
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "offCount" in data:
            args["off_count"] = data["offCount"]
        if "offValue" in data:
            args["off_value"] = data["offValue"]
        if "onCount" in data:
            args["on_count"] = data["onCount"]
        if "onValue" in data:
            args["on_value"] = data["onValue"]
        if "timeStep" in data:
            args["time_step"] = data["timeStep"]        

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


class Key:
    tick: float
    type_val: str
    uid: typing.Optional[str]

    def __init__(self, tick: float = 0, type_val: str = "", uid: typing.Optional[str] = None):
        self.tick = tick
        self.type_val = type_val
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "tick": self.tick,
            "type": self.type_val,
        }
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "tick" in data:
            args["tick"] = data["tick"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)


class SimulationQuery:
    config: typing.Optional[object]
    key: 'Key'
    last: typing.Optional[bool]
    stream: typing.Optional[bool]

    def __init__(self, config: typing.Optional[object] = None, key: typing.Optional['Key'] = None, last: typing.Optional[bool] = None, stream: typing.Optional[bool] = None):
        self.config = config
        self.key = key if key is not None else Key()
        self.last = last
        self.stream = stream

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "key": self.key,
        }
        if self.config is not None:
            payload["config"] = self.config
        if self.last is not None:
            payload["last"] = self.last
        if self.stream is not None:
            payload["stream"] = self.stream
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "config" in data:
            args["config"] = data["config"]
        if "key" in data:
            args["key"] = Key.from_json(data["key"])
        if "last" in data:
            args["last"] = data["last"]
        if "stream" in data:
            args["stream"] = data["stream"]        

        return cls(**args)


class StreamingQuery:
    bands: typing.Optional[int]
    noise: float
    speed: float
    spread: float
    # Possible enum values:
    #  - `"fetch"` 
    #  - `"logs"` 
    #  - `"signal"` 
    #  - `"traces"` 
    type_val: typing.Literal["fetch", "logs", "signal", "traces"]
    url: typing.Optional[str]

    def __init__(self, bands: typing.Optional[int] = None, noise: float = 0, speed: float = 0, spread: float = 0, type_val: typing.Optional[typing.Literal["fetch", "logs", "signal", "traces"]] = None, url: typing.Optional[str] = None):
        self.bands = bands
        self.noise = noise
        self.speed = speed
        self.spread = spread
        self.type_val = type_val if type_val is not None else "fetch"
        self.url = url

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "noise": self.noise,
            "speed": self.speed,
            "spread": self.spread,
            "type": self.type_val,
        }
        if self.bands is not None:
            payload["bands"] = self.bands
        if self.url is not None:
            payload["url"] = self.url
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "bands" in data:
            args["bands"] = data["bands"]
        if "noise" in data:
            args["noise"] = data["noise"]
        if "speed" in data:
            args["speed"] = data["speed"]
        if "spread" in data:
            args["spread"] = data["spread"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "url" in data:
            args["url"] = data["url"]        

        return cls(**args)


class TimeRange:
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


class USAQuery:
    fields: typing.Optional[list[str]]
    mode: typing.Optional[str]
    period: typing.Optional[str]
    states: typing.Optional[list[str]]

    def __init__(self, fields: typing.Optional[list[str]] = None, mode: typing.Optional[str] = None, period: typing.Optional[str] = None, states: typing.Optional[list[str]] = None):
        self.fields = fields
        self.mode = mode
        self.period = period
        self.states = states

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.fields is not None:
            payload["fields"] = self.fields
        if self.mode is not None:
            payload["mode"] = self.mode
        if self.period is not None:
            payload["period"] = self.period
        if self.states is not None:
            payload["states"] = self.states
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "fields" in data:
            args["fields"] = data["fields"]
        if "mode" in data:
            args["mode"] = data["mode"]
        if "period" in data:
            args["period"] = data["period"]
        if "states" in data:
            args["states"] = data["states"]        

        return cls(**args)


class Dataquery(cogvariants.Dataquery):
    alias: typing.Optional[str]
    # Used for live query
    channel: typing.Optional[str]
    csv_content: typing.Optional[str]
    csv_file_name: typing.Optional[str]
    csv_wave: typing.Optional[list['CSVWave']]
    # The datasource
    datasource: typing.Optional['Datasource']
    # Drop percentage (the chance we will lose a point 0-100)
    drop_percent: typing.Optional[float]
    # Possible enum values:
    #  - `"frontend_exception"` 
    #  - `"frontend_observable"` 
    #  - `"server_panic"` 
    error_type: typing.Optional[typing.Literal["frontend_exception", "frontend_observable", "server_panic"]]
    flamegraph_diff: typing.Optional[bool]
    # true if query is disabled (ie should not be returned to the dashboard)
    # NOTE: this does not always imply that the query should not be executed since
    # the results from a hidden query may be used as the input to other queries (SSE etc)
    hide: typing.Optional[bool]
    # Interval is the suggested duration between time points in a time series query.
    # NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
    # from the interval required to fill a pixels in the visualization
    interval_ms: typing.Optional[float]
    labels: typing.Optional[str]
    level_column: typing.Optional[bool]
    lines: typing.Optional[int]
    max_val: typing.Optional[float]
    # MaxDataPoints is the maximum number of data points that should be returned from a time series query.
    # NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
    # from the number of pixels visible in a visualization
    max_data_points: typing.Optional[int]
    min_val: typing.Optional[float]
    nodes: typing.Optional['NodesQuery']
    noise: typing.Optional[float]
    points: typing.Optional[list[list[object]]]
    pulse_wave: typing.Optional['PulseWaveQuery']
    # QueryType is an optional identifier for the type of query.
    # It can be used to distinguish different types of queries.
    query_type: typing.Optional[str]
    raw_frame_content: typing.Optional[str]
    # RefID is the unique identifier of the query, set by the frontend call.
    ref_id: typing.Optional[str]
    # Optionally define expected query result behavior
    result_assertions: typing.Optional['ResultAssertions']
    # Possible enum values:
    #  - `"annotations"` 
    #  - `"arrow"` 
    #  - `"csv_content"` 
    #  - `"csv_file"` 
    #  - `"csv_metric_values"` 
    #  - `"datapoints_outside_range"` 
    #  - `"exponential_heatmap_bucket_data"` 
    #  - `"flame_graph"` 
    #  - `"grafana_api"` 
    #  - `"linear_heatmap_bucket_data"` 
    #  - `"live"` 
    #  - `"logs"` 
    #  - `"manual_entry"` 
    #  - `"no_data_points"` 
    #  - `"node_graph"` 
    #  - `"predictable_csv_wave"` 
    #  - `"predictable_pulse"` 
    #  - `"random_walk"` 
    #  - `"random_walk_table"` 
    #  - `"random_walk_with_error"` 
    #  - `"raw_frame"` 
    #  - `"server_error_500"` 
    #  - `"simulation"` 
    #  - `"slow_query"` 
    #  - `"streaming_client"` 
    #  - `"table_static"` 
    #  - `"trace"` 
    #  - `"usa"` 
    #  - `"variables-query"` 
    scenario_id: typing.Optional[typing.Literal["annotations", "arrow", "csv_content", "csv_file", "csv_metric_values", "datapoints_outside_range", "exponential_heatmap_bucket_data", "flame_graph", "grafana_api", "linear_heatmap_bucket_data", "live", "logs", "manual_entry", "no_data_points", "node_graph", "predictable_csv_wave", "predictable_pulse", "random_walk", "random_walk_table", "random_walk_with_error", "raw_frame", "server_error_500", "simulation", "slow_query", "streaming_client", "table_static", "trace", "usa", "variables-query"]]
    series_count: typing.Optional[int]
    sim: typing.Optional['SimulationQuery']
    span_count: typing.Optional[int]
    spread: typing.Optional[float]
    start_value: typing.Optional[float]
    stream: typing.Optional['StreamingQuery']
    # common parameter used by many query types
    string_input: typing.Optional[str]
    # TimeRange represents the query range
    # NOTE: unlike generic /ds/query, we can now send explicit time values in each query
    # NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
    time_range: typing.Optional['TimeRange']
    usa: typing.Optional['USAQuery']
    with_nil: typing.Optional[bool]

    def __init__(self, alias: typing.Optional[str] = None, channel: typing.Optional[str] = None, csv_content: typing.Optional[str] = None, csv_file_name: typing.Optional[str] = None, csv_wave: typing.Optional[list['CSVWave']] = None, datasource: typing.Optional['Datasource'] = None, drop_percent: typing.Optional[float] = None, error_type: typing.Optional[typing.Literal["frontend_exception", "frontend_observable", "server_panic"]] = None, flamegraph_diff: typing.Optional[bool] = None, hide: typing.Optional[bool] = None, interval_ms: typing.Optional[float] = None, labels: typing.Optional[str] = None, level_column: typing.Optional[bool] = None, lines: typing.Optional[int] = None, max_val: typing.Optional[float] = None, max_data_points: typing.Optional[int] = None, min_val: typing.Optional[float] = None, nodes: typing.Optional['NodesQuery'] = None, noise: typing.Optional[float] = None, points: typing.Optional[list[list[object]]] = None, pulse_wave: typing.Optional['PulseWaveQuery'] = None, query_type: typing.Optional[str] = None, raw_frame_content: typing.Optional[str] = None, ref_id: typing.Optional[str] = None, result_assertions: typing.Optional['ResultAssertions'] = None, scenario_id: typing.Optional[typing.Literal["annotations", "arrow", "csv_content", "csv_file", "csv_metric_values", "datapoints_outside_range", "exponential_heatmap_bucket_data", "flame_graph", "grafana_api", "linear_heatmap_bucket_data", "live", "logs", "manual_entry", "no_data_points", "node_graph", "predictable_csv_wave", "predictable_pulse", "random_walk", "random_walk_table", "random_walk_with_error", "raw_frame", "server_error_500", "simulation", "slow_query", "streaming_client", "table_static", "trace", "usa", "variables-query"]] = None, series_count: typing.Optional[int] = None, sim: typing.Optional['SimulationQuery'] = None, span_count: typing.Optional[int] = None, spread: typing.Optional[float] = None, start_value: typing.Optional[float] = None, stream: typing.Optional['StreamingQuery'] = None, string_input: typing.Optional[str] = None, time_range: typing.Optional['TimeRange'] = None, usa: typing.Optional['USAQuery'] = None, with_nil: typing.Optional[bool] = None):
        self.alias = alias
        self.channel = channel
        self.csv_content = csv_content
        self.csv_file_name = csv_file_name
        self.csv_wave = csv_wave
        self.datasource = datasource
        self.drop_percent = drop_percent
        self.error_type = error_type
        self.flamegraph_diff = flamegraph_diff
        self.hide = hide
        self.interval_ms = interval_ms
        self.labels = labels
        self.level_column = level_column
        self.lines = lines
        self.max_val = max_val
        self.max_data_points = max_data_points
        self.min_val = min_val
        self.nodes = nodes
        self.noise = noise
        self.points = points
        self.pulse_wave = pulse_wave
        self.query_type = query_type
        self.raw_frame_content = raw_frame_content
        self.ref_id = ref_id
        self.result_assertions = result_assertions
        self.scenario_id = scenario_id
        self.series_count = series_count
        self.sim = sim
        self.span_count = span_count
        self.spread = spread
        self.start_value = start_value
        self.stream = stream
        self.string_input = string_input
        self.time_range = time_range
        self.usa = usa
        self.with_nil = with_nil

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.alias is not None:
            payload["alias"] = self.alias
        if self.channel is not None:
            payload["channel"] = self.channel
        if self.csv_content is not None:
            payload["csvContent"] = self.csv_content
        if self.csv_file_name is not None:
            payload["csvFileName"] = self.csv_file_name
        if self.csv_wave is not None:
            payload["csvWave"] = self.csv_wave
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.drop_percent is not None:
            payload["dropPercent"] = self.drop_percent
        if self.error_type is not None:
            payload["errorType"] = self.error_type
        if self.flamegraph_diff is not None:
            payload["flamegraphDiff"] = self.flamegraph_diff
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        if self.labels is not None:
            payload["labels"] = self.labels
        if self.level_column is not None:
            payload["levelColumn"] = self.level_column
        if self.lines is not None:
            payload["lines"] = self.lines
        if self.max_val is not None:
            payload["max"] = self.max_val
        if self.max_data_points is not None:
            payload["maxDataPoints"] = self.max_data_points
        if self.min_val is not None:
            payload["min"] = self.min_val
        if self.nodes is not None:
            payload["nodes"] = self.nodes
        if self.noise is not None:
            payload["noise"] = self.noise
        if self.points is not None:
            payload["points"] = self.points
        if self.pulse_wave is not None:
            payload["pulseWave"] = self.pulse_wave
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.raw_frame_content is not None:
            payload["rawFrameContent"] = self.raw_frame_content
        if self.ref_id is not None:
            payload["refId"] = self.ref_id
        if self.result_assertions is not None:
            payload["resultAssertions"] = self.result_assertions
        if self.scenario_id is not None:
            payload["scenarioId"] = self.scenario_id
        if self.series_count is not None:
            payload["seriesCount"] = self.series_count
        if self.sim is not None:
            payload["sim"] = self.sim
        if self.span_count is not None:
            payload["spanCount"] = self.span_count
        if self.spread is not None:
            payload["spread"] = self.spread
        if self.start_value is not None:
            payload["startValue"] = self.start_value
        if self.stream is not None:
            payload["stream"] = self.stream
        if self.string_input is not None:
            payload["stringInput"] = self.string_input
        if self.time_range is not None:
            payload["timeRange"] = self.time_range
        if self.usa is not None:
            payload["usa"] = self.usa
        if self.with_nil is not None:
            payload["withNil"] = self.with_nil
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "alias" in data:
            args["alias"] = data["alias"]
        if "channel" in data:
            args["channel"] = data["channel"]
        if "csvContent" in data:
            args["csv_content"] = data["csvContent"]
        if "csvFileName" in data:
            args["csv_file_name"] = data["csvFileName"]
        if "csvWave" in data:
            args["csv_wave"] = data["csvWave"]
        if "datasource" in data:
            args["datasource"] = Datasource.from_json(data["datasource"])
        if "dropPercent" in data:
            args["drop_percent"] = data["dropPercent"]
        if "errorType" in data:
            args["error_type"] = data["errorType"]
        if "flamegraphDiff" in data:
            args["flamegraph_diff"] = data["flamegraphDiff"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]
        if "labels" in data:
            args["labels"] = data["labels"]
        if "levelColumn" in data:
            args["level_column"] = data["levelColumn"]
        if "lines" in data:
            args["lines"] = data["lines"]
        if "max" in data:
            args["max_val"] = data["max"]
        if "maxDataPoints" in data:
            args["max_data_points"] = data["maxDataPoints"]
        if "min" in data:
            args["min_val"] = data["min"]
        if "nodes" in data:
            args["nodes"] = NodesQuery.from_json(data["nodes"])
        if "noise" in data:
            args["noise"] = data["noise"]
        if "points" in data:
            args["points"] = data["points"]
        if "pulseWave" in data:
            args["pulse_wave"] = PulseWaveQuery.from_json(data["pulseWave"])
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "rawFrameContent" in data:
            args["raw_frame_content"] = data["rawFrameContent"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "resultAssertions" in data:
            args["result_assertions"] = ResultAssertions.from_json(data["resultAssertions"])
        if "scenarioId" in data:
            args["scenario_id"] = data["scenarioId"]
        if "seriesCount" in data:
            args["series_count"] = data["seriesCount"]
        if "sim" in data:
            args["sim"] = SimulationQuery.from_json(data["sim"])
        if "spanCount" in data:
            args["span_count"] = data["spanCount"]
        if "spread" in data:
            args["spread"] = data["spread"]
        if "startValue" in data:
            args["start_value"] = data["startValue"]
        if "stream" in data:
            args["stream"] = StreamingQuery.from_json(data["stream"])
        if "stringInput" in data:
            args["string_input"] = data["stringInput"]
        if "timeRange" in data:
            args["time_range"] = TimeRange.from_json(data["timeRange"])
        if "usa" in data:
            args["usa"] = USAQuery.from_json(data["usa"])
        if "withNil" in data:
            args["with_nil"] = data["withNil"]        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="",
        from_json_hook=Dataquery.from_json,
    )




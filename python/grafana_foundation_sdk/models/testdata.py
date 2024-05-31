# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
import typing
from ..cog import variants as cogvariants
from ..cog import runtime as cogruntime


class TestDataQueryType(enum.StrEnum):
    RANDOM_WALK = "random_walk"
    SLOW_QUERY = "slow_query"
    RANDOM_WALK_WITH_ERROR = "random_walk_with_error"
    RANDOM_WALK_TABLE = "random_walk_table"
    EXPONENTIAL_HEATMAP_BUCKET_DATA = "exponential_heatmap_bucket_data"
    LINEAR_HEATMAP_BUCKET_DATA = "linear_heatmap_bucket_data"
    NO_DATA_POINTS = "no_data_points"
    DATA_POINTS_OUTSIDE_RANGE = "datapoints_outside_range"
    CSV_METRIC_VALUES = "csv_metric_values"
    PREDICTABLE_PULSE = "predictable_pulse"
    PREDICTABLE_CSV_WAVE = "predictable_csv_wave"
    STREAMING_CLIENT = "streaming_client"
    SIMULATION = "simulation"
    USA = "usa"
    LIVE = "live"
    GRAFANA_API = "grafana_api"
    ARROW = "arrow"
    ANNOTATIONS = "annotations"
    TABLE_STATIC = "table_static"
    SERVER_ERROR500 = "server_error_500"
    LOGS = "logs"
    NODE_GRAPH = "node_graph"
    FLAME_GRAPH = "flame_graph"
    RAW_FRAME = "raw_frame"
    CSV_FILE = "csv_file"
    CSV_CONTENT = "csv_content"
    TRACE = "trace"
    MANUAL_ENTRY = "manual_entry"
    VARIABLES_QUERY = "variables-query"


class StreamingQuery:
    type_val: typing.Literal["signal", "logs", "fetch", "traces"]
    speed: int
    spread: int
    noise: int
    bands: typing.Optional[int]
    url: typing.Optional[str]

    def __init__(self, type_val: typing.Optional[typing.Literal["signal", "logs", "fetch", "traces"]] = None, speed: int = 0, spread: int = 0, noise: int = 0, bands: typing.Optional[int] = None, url: typing.Optional[str] = None):
        self.type_val = type_val if type_val is not None else "signal"
        self.speed = speed
        self.spread = spread
        self.noise = noise
        self.bands = bands
        self.url = url

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "speed": self.speed,
            "spread": self.spread,
            "noise": self.noise,
        }
        if self.bands is not None:
            payload["bands"] = self.bands
        if self.url is not None:
            payload["url"] = self.url
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "speed" in data:
            args["speed"] = data["speed"]
        if "spread" in data:
            args["spread"] = data["spread"]
        if "noise" in data:
            args["noise"] = data["noise"]
        if "bands" in data:
            args["bands"] = data["bands"]
        if "url" in data:
            args["url"] = data["url"]        

        return cls(**args)


class PulseWaveQuery:
    time_step: typing.Optional[int]
    on_count: typing.Optional[int]
    off_count: typing.Optional[int]
    on_value: typing.Optional[float]
    off_value: typing.Optional[float]

    def __init__(self, time_step: typing.Optional[int] = None, on_count: typing.Optional[int] = None, off_count: typing.Optional[int] = None, on_value: typing.Optional[float] = None, off_value: typing.Optional[float] = None):
        self.time_step = time_step
        self.on_count = on_count
        self.off_count = off_count
        self.on_value = on_value
        self.off_value = off_value

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.time_step is not None:
            payload["timeStep"] = self.time_step
        if self.on_count is not None:
            payload["onCount"] = self.on_count
        if self.off_count is not None:
            payload["offCount"] = self.off_count
        if self.on_value is not None:
            payload["onValue"] = self.on_value
        if self.off_value is not None:
            payload["offValue"] = self.off_value
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "timeStep" in data:
            args["time_step"] = data["timeStep"]
        if "onCount" in data:
            args["on_count"] = data["onCount"]
        if "offCount" in data:
            args["off_count"] = data["offCount"]
        if "onValue" in data:
            args["on_value"] = data["onValue"]
        if "offValue" in data:
            args["off_value"] = data["offValue"]        

        return cls(**args)


class SimulationQuery:
    key: 'Key'
    config: typing.Optional[dict[str, object]]
    stream: typing.Optional[bool]
    last: typing.Optional[bool]

    def __init__(self, key: typing.Optional['Key'] = None, config: typing.Optional[dict[str, object]] = None, stream: typing.Optional[bool] = None, last: typing.Optional[bool] = None):
        self.key = key if key is not None else Key()
        self.config = config
        self.stream = stream
        self.last = last

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "key": self.key,
        }
        if self.config is not None:
            payload["config"] = self.config
        if self.stream is not None:
            payload["stream"] = self.stream
        if self.last is not None:
            payload["last"] = self.last
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "key" in data:
            args["key"] = Key.from_json(data["key"])
        if "config" in data:
            args["config"] = data["config"]
        if "stream" in data:
            args["stream"] = data["stream"]
        if "last" in data:
            args["last"] = data["last"]        

        return cls(**args)


class NodesQuery:
    type_val: typing.Optional[typing.Literal["random", "response_small", "response_medium", "random edges"]]
    count: typing.Optional[int]
    seed: typing.Optional[int]

    def __init__(self, type_val: typing.Optional[typing.Literal["random", "response_small", "response_medium", "random edges"]] = None, count: typing.Optional[int] = None, seed: typing.Optional[int] = None):
        self.type_val = type_val
        self.count = count
        self.seed = seed

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.type_val is not None:
            payload["type"] = self.type_val
        if self.count is not None:
            payload["count"] = self.count
        if self.seed is not None:
            payload["seed"] = self.seed
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "count" in data:
            args["count"] = data["count"]
        if "seed" in data:
            args["seed"] = data["seed"]        

        return cls(**args)


class USAQuery:
    mode: typing.Optional[str]
    period: typing.Optional[str]
    fields: typing.Optional[list[str]]
    states: typing.Optional[list[str]]

    def __init__(self, mode: typing.Optional[str] = None, period: typing.Optional[str] = None, fields: typing.Optional[list[str]] = None, states: typing.Optional[list[str]] = None):
        self.mode = mode
        self.period = period
        self.fields = fields
        self.states = states

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.mode is not None:
            payload["mode"] = self.mode
        if self.period is not None:
            payload["period"] = self.period
        if self.fields is not None:
            payload["fields"] = self.fields
        if self.states is not None:
            payload["states"] = self.states
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "mode" in data:
            args["mode"] = data["mode"]
        if "period" in data:
            args["period"] = data["period"]
        if "fields" in data:
            args["fields"] = data["fields"]
        if "states" in data:
            args["states"] = data["states"]        

        return cls(**args)


class CSVWave:
    time_step: typing.Optional[int]
    name: typing.Optional[str]
    values_csv: typing.Optional[str]
    labels: typing.Optional[str]

    def __init__(self, time_step: typing.Optional[int] = None, name: typing.Optional[str] = None, values_csv: typing.Optional[str] = None, labels: typing.Optional[str] = None):
        self.time_step = time_step
        self.name = name
        self.values_csv = values_csv
        self.labels = labels

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.time_step is not None:
            payload["timeStep"] = self.time_step
        if self.name is not None:
            payload["name"] = self.name
        if self.values_csv is not None:
            payload["valuesCSV"] = self.values_csv
        if self.labels is not None:
            payload["labels"] = self.labels
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "timeStep" in data:
            args["time_step"] = data["timeStep"]
        if "name" in data:
            args["name"] = data["name"]
        if "valuesCSV" in data:
            args["values_csv"] = data["valuesCSV"]
        if "labels" in data:
            args["labels"] = data["labels"]        

        return cls(**args)


class Scenario:
    """
    TODO: Should this live here given it's not used in the dataquery?
    """

    id_val: str
    name: str
    string_input: str
    description: typing.Optional[str]
    hide_alias_field: typing.Optional[bool]

    def __init__(self, id_val: str = "", name: str = "", string_input: str = "", description: typing.Optional[str] = None, hide_alias_field: typing.Optional[bool] = None):
        self.id_val = id_val
        self.name = name
        self.string_input = string_input
        self.description = description
        self.hide_alias_field = hide_alias_field

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "name": self.name,
            "stringInput": self.string_input,
        }
        if self.description is not None:
            payload["description"] = self.description
        if self.hide_alias_field is not None:
            payload["hideAliasField"] = self.hide_alias_field
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "name" in data:
            args["name"] = data["name"]
        if "stringInput" in data:
            args["string_input"] = data["stringInput"]
        if "description" in data:
            args["description"] = data["description"]
        if "hideAliasField" in data:
            args["hide_alias_field"] = data["hideAliasField"]        

        return cls(**args)


class Dataquery(cogvariants.Dataquery):
    alias: typing.Optional[str]
    scenario_id: typing.Optional['TestDataQueryType']
    string_input: typing.Optional[str]
    stream: typing.Optional['StreamingQuery']
    pulse_wave: typing.Optional['PulseWaveQuery']
    sim: typing.Optional['SimulationQuery']
    csv_wave: typing.Optional[list['CSVWave']]
    labels: typing.Optional[str]
    lines: typing.Optional[int]
    level_column: typing.Optional[bool]
    channel: typing.Optional[str]
    nodes: typing.Optional['NodesQuery']
    csv_file_name: typing.Optional[str]
    csv_content: typing.Optional[str]
    raw_frame_content: typing.Optional[str]
    series_count: typing.Optional[int]
    usa: typing.Optional['USAQuery']
    error_type: typing.Optional[typing.Literal["server_panic", "frontend_exception", "frontend_observable"]]
    span_count: typing.Optional[int]
    points: typing.Optional[list[list[typing.Union[str, int]]]]
    # Drop percentage (the chance we will lose a point 0-100)
    drop_percent: typing.Optional[float]
    flamegraph_diff: typing.Optional[bool]
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
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]

    def __init__(self, alias: typing.Optional[str] = None, scenario_id: typing.Optional['TestDataQueryType'] = None, string_input: typing.Optional[str] = None, stream: typing.Optional['StreamingQuery'] = None, pulse_wave: typing.Optional['PulseWaveQuery'] = None, sim: typing.Optional['SimulationQuery'] = None, csv_wave: typing.Optional[list['CSVWave']] = None, labels: typing.Optional[str] = None, lines: typing.Optional[int] = None, level_column: typing.Optional[bool] = None, channel: typing.Optional[str] = None, nodes: typing.Optional['NodesQuery'] = None, csv_file_name: typing.Optional[str] = None, csv_content: typing.Optional[str] = None, raw_frame_content: typing.Optional[str] = None, series_count: typing.Optional[int] = None, usa: typing.Optional['USAQuery'] = None, error_type: typing.Optional[typing.Literal["server_panic", "frontend_exception", "frontend_observable"]] = None, span_count: typing.Optional[int] = None, points: typing.Optional[list[list[typing.Union[str, int]]]] = None, drop_percent: typing.Optional[float] = None, flamegraph_diff: typing.Optional[bool] = None, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[object] = None):
        self.alias = alias
        self.scenario_id = scenario_id if scenario_id is not None else TestDataQueryType.RANDOM_WALK
        self.string_input = string_input
        self.stream = stream
        self.pulse_wave = pulse_wave
        self.sim = sim
        self.csv_wave = csv_wave
        self.labels = labels
        self.lines = lines
        self.level_column = level_column
        self.channel = channel
        self.nodes = nodes
        self.csv_file_name = csv_file_name
        self.csv_content = csv_content
        self.raw_frame_content = raw_frame_content
        self.series_count = series_count
        self.usa = usa
        self.error_type = error_type
        self.span_count = span_count
        self.points = points
        self.drop_percent = drop_percent
        self.flamegraph_diff = flamegraph_diff
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.datasource = datasource

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "refId": self.ref_id,
        }
        if self.alias is not None:
            payload["alias"] = self.alias
        if self.scenario_id is not None:
            payload["scenarioId"] = self.scenario_id
        if self.string_input is not None:
            payload["stringInput"] = self.string_input
        if self.stream is not None:
            payload["stream"] = self.stream
        if self.pulse_wave is not None:
            payload["pulseWave"] = self.pulse_wave
        if self.sim is not None:
            payload["sim"] = self.sim
        if self.csv_wave is not None:
            payload["csvWave"] = self.csv_wave
        if self.labels is not None:
            payload["labels"] = self.labels
        if self.lines is not None:
            payload["lines"] = self.lines
        if self.level_column is not None:
            payload["levelColumn"] = self.level_column
        if self.channel is not None:
            payload["channel"] = self.channel
        if self.nodes is not None:
            payload["nodes"] = self.nodes
        if self.csv_file_name is not None:
            payload["csvFileName"] = self.csv_file_name
        if self.csv_content is not None:
            payload["csvContent"] = self.csv_content
        if self.raw_frame_content is not None:
            payload["rawFrameContent"] = self.raw_frame_content
        if self.series_count is not None:
            payload["seriesCount"] = self.series_count
        if self.usa is not None:
            payload["usa"] = self.usa
        if self.error_type is not None:
            payload["errorType"] = self.error_type
        if self.span_count is not None:
            payload["spanCount"] = self.span_count
        if self.points is not None:
            payload["points"] = self.points
        if self.drop_percent is not None:
            payload["dropPercent"] = self.drop_percent
        if self.flamegraph_diff is not None:
            payload["flamegraphDiff"] = self.flamegraph_diff
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
        
        if "alias" in data:
            args["alias"] = data["alias"]
        if "scenarioId" in data:
            args["scenario_id"] = data["scenarioId"]
        if "stringInput" in data:
            args["string_input"] = data["stringInput"]
        if "stream" in data:
            args["stream"] = StreamingQuery.from_json(data["stream"])
        if "pulseWave" in data:
            args["pulse_wave"] = PulseWaveQuery.from_json(data["pulseWave"])
        if "sim" in data:
            args["sim"] = SimulationQuery.from_json(data["sim"])
        if "csvWave" in data:
            args["csv_wave"] = data["csvWave"]
        if "labels" in data:
            args["labels"] = data["labels"]
        if "lines" in data:
            args["lines"] = data["lines"]
        if "levelColumn" in data:
            args["level_column"] = data["levelColumn"]
        if "channel" in data:
            args["channel"] = data["channel"]
        if "nodes" in data:
            args["nodes"] = NodesQuery.from_json(data["nodes"])
        if "csvFileName" in data:
            args["csv_file_name"] = data["csvFileName"]
        if "csvContent" in data:
            args["csv_content"] = data["csvContent"]
        if "rawFrameContent" in data:
            args["raw_frame_content"] = data["rawFrameContent"]
        if "seriesCount" in data:
            args["series_count"] = data["seriesCount"]
        if "usa" in data:
            args["usa"] = USAQuery.from_json(data["usa"])
        if "errorType" in data:
            args["error_type"] = data["errorType"]
        if "spanCount" in data:
            args["span_count"] = data["spanCount"]
        if "points" in data:
            args["points"] = data["points"]
        if "dropPercent" in data:
            args["drop_percent"] = data["dropPercent"]
        if "flamegraphDiff" in data:
            args["flamegraph_diff"] = data["flamegraphDiff"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="testdata",
        from_json_hook=Dataquery.from_json,
    )


class Key:
    type_val: str
    tick: float
    uid: typing.Optional[str]

    def __init__(self, type_val: str = "", tick: float = 0, uid: typing.Optional[str] = None):
        self.type_val = type_val
        self.tick = tick
        self.uid = uid

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "tick": self.tick,
        }
        if self.uid is not None:
            payload["uid"] = self.uid
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "tick" in data:
            args["tick"] = data["tick"]
        if "uid" in data:
            args["uid"] = data["uid"]        

        return cls(**args)




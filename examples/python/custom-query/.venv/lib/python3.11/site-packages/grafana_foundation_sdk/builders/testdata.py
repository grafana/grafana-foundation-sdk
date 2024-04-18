# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import testdata


class StreamingQuery(cogbuilder.Builder[testdata.StreamingQuery]):    
    __internal: testdata.StreamingQuery

    def __init__(self):
        self.__internal = testdata.StreamingQuery()

    def build(self) -> testdata.StreamingQuery:
        return self.__internal    
    
    def type_val(self, type_val: typing.Literal["signal", "logs", "fetch", "traces"]) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def speed(self, speed: int) -> typing.Self:        
        self.__internal.speed = speed
    
        return self
    
    def spread(self, spread: int) -> typing.Self:        
        self.__internal.spread = spread
    
        return self
    
    def noise(self, noise: int) -> typing.Self:        
        self.__internal.noise = noise
    
        return self
    
    def bands(self, bands: int) -> typing.Self:        
        self.__internal.bands = bands
    
        return self
    
    def url(self, url: str) -> typing.Self:        
        self.__internal.url = url
    
        return self
    

class PulseWaveQuery(cogbuilder.Builder[testdata.PulseWaveQuery]):    
    __internal: testdata.PulseWaveQuery

    def __init__(self):
        self.__internal = testdata.PulseWaveQuery()

    def build(self) -> testdata.PulseWaveQuery:
        return self.__internal    
    
    def time_step(self, time_step: int) -> typing.Self:        
        self.__internal.time_step = time_step
    
        return self
    
    def on_count(self, on_count: int) -> typing.Self:        
        self.__internal.on_count = on_count
    
        return self
    
    def off_count(self, off_count: int) -> typing.Self:        
        self.__internal.off_count = off_count
    
        return self
    
    def on_value(self, on_value: float) -> typing.Self:        
        self.__internal.on_value = on_value
    
        return self
    
    def off_value(self, off_value: float) -> typing.Self:        
        self.__internal.off_value = off_value
    
        return self
    

class SimulationQuery(cogbuilder.Builder[testdata.SimulationQuery]):    
    __internal: testdata.SimulationQuery

    def __init__(self):
        self.__internal = testdata.SimulationQuery()

    def build(self) -> testdata.SimulationQuery:
        return self.__internal    
    
    def key(self, key: cogbuilder.Builder[testdata.Key]) -> typing.Self:        
        key_resource = key.build()
        self.__internal.key = key_resource
    
        return self
    
    def config(self, config: object) -> typing.Self:        
        self.__internal.config = config
    
        return self
    
    def stream(self, stream: bool) -> typing.Self:        
        self.__internal.stream = stream
    
        return self
    
    def last(self, last: bool) -> typing.Self:        
        self.__internal.last = last
    
        return self
    

class NodesQuery(cogbuilder.Builder[testdata.NodesQuery]):    
    __internal: testdata.NodesQuery

    def __init__(self):
        self.__internal = testdata.NodesQuery()

    def build(self) -> testdata.NodesQuery:
        return self.__internal    
    
    def type_val(self, type_val: typing.Literal["random", "response_small", "response_medium", "random edges"]) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def count(self, count: int) -> typing.Self:        
        self.__internal.count = count
    
        return self
    
    def seed(self, seed: int) -> typing.Self:        
        self.__internal.seed = seed
    
        return self
    

class USAQuery(cogbuilder.Builder[testdata.USAQuery]):    
    __internal: testdata.USAQuery

    def __init__(self):
        self.__internal = testdata.USAQuery()

    def build(self) -> testdata.USAQuery:
        return self.__internal    
    
    def mode(self, mode: str) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    
    def period(self, period: str) -> typing.Self:        
        self.__internal.period = period
    
        return self
    
    def fields(self, fields: list[str]) -> typing.Self:        
        self.__internal.fields = fields
    
        return self
    
    def states(self, states: list[str]) -> typing.Self:        
        self.__internal.states = states
    
        return self
    

class CSVWave(cogbuilder.Builder[testdata.CSVWave]):    
    __internal: testdata.CSVWave

    def __init__(self):
        self.__internal = testdata.CSVWave()

    def build(self) -> testdata.CSVWave:
        return self.__internal    
    
    def time_step(self, time_step: int) -> typing.Self:        
        self.__internal.time_step = time_step
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self.__internal.name = name
    
        return self
    
    def values_csv(self, values_csv: str) -> typing.Self:        
        self.__internal.values_csv = values_csv
    
        return self
    
    def labels(self, labels: str) -> typing.Self:        
        self.__internal.labels = labels
    
        return self
    

class Scenario(cogbuilder.Builder[testdata.Scenario]):    
    """
    TODO: Should this live here given it's not used in the dataquery?
    """
    
    __internal: testdata.Scenario

    def __init__(self):
        self.__internal = testdata.Scenario()

    def build(self) -> testdata.Scenario:
        return self.__internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self.__internal.name = name
    
        return self
    
    def string_input(self, string_input: str) -> typing.Self:        
        self.__internal.string_input = string_input
    
        return self
    
    def description(self, description: str) -> typing.Self:        
        self.__internal.description = description
    
        return self
    
    def hide_alias_field(self, hide_alias_field: bool) -> typing.Self:        
        self.__internal.hide_alias_field = hide_alias_field
    
        return self
    

class Dataquery(cogbuilder.Builder[testdata.Dataquery]):    
    __internal: testdata.Dataquery

    def __init__(self):
        self.__internal = testdata.Dataquery()

    def build(self) -> testdata.Dataquery:
        return self.__internal    
    
    def alias(self, alias: str) -> typing.Self:        
        self.__internal.alias = alias
    
        return self
    
    def scenario_id(self, scenario_id: testdata.TestDataQueryType) -> typing.Self:        
        self.__internal.scenario_id = scenario_id
    
        return self
    
    def string_input(self, string_input: str) -> typing.Self:        
        self.__internal.string_input = string_input
    
        return self
    
    def stream(self, stream: cogbuilder.Builder[testdata.StreamingQuery]) -> typing.Self:        
        stream_resource = stream.build()
        self.__internal.stream = stream_resource
    
        return self
    
    def pulse_wave(self, pulse_wave: cogbuilder.Builder[testdata.PulseWaveQuery]) -> typing.Self:        
        pulse_wave_resource = pulse_wave.build()
        self.__internal.pulse_wave = pulse_wave_resource
    
        return self
    
    def sim(self, sim: cogbuilder.Builder[testdata.SimulationQuery]) -> typing.Self:        
        sim_resource = sim.build()
        self.__internal.sim = sim_resource
    
        return self
    
    def csv_wave(self, csv_wave: list[cogbuilder.Builder[testdata.CSVWave]]) -> typing.Self:        
        csv_wave_resources = [r1.build() for r1 in csv_wave]
        self.__internal.csv_wave = csv_wave_resources
    
        return self
    
    def labels(self, labels: str) -> typing.Self:        
        self.__internal.labels = labels
    
        return self
    
    def lines(self, lines: int) -> typing.Self:        
        self.__internal.lines = lines
    
        return self
    
    def level_column(self, level_column: bool) -> typing.Self:        
        self.__internal.level_column = level_column
    
        return self
    
    def channel(self, channel: str) -> typing.Self:        
        self.__internal.channel = channel
    
        return self
    
    def nodes(self, nodes: cogbuilder.Builder[testdata.NodesQuery]) -> typing.Self:        
        nodes_resource = nodes.build()
        self.__internal.nodes = nodes_resource
    
        return self
    
    def csv_file_name(self, csv_file_name: str) -> typing.Self:        
        self.__internal.csv_file_name = csv_file_name
    
        return self
    
    def csv_content(self, csv_content: str) -> typing.Self:        
        self.__internal.csv_content = csv_content
    
        return self
    
    def raw_frame_content(self, raw_frame_content: str) -> typing.Self:        
        self.__internal.raw_frame_content = raw_frame_content
    
        return self
    
    def series_count(self, series_count: int) -> typing.Self:        
        self.__internal.series_count = series_count
    
        return self
    
    def usa(self, usa: cogbuilder.Builder[testdata.USAQuery]) -> typing.Self:        
        usa_resource = usa.build()
        self.__internal.usa = usa_resource
    
        return self
    
    def error_type(self, error_type: typing.Literal["server_panic", "frontend_exception", "frontend_observable"]) -> typing.Self:        
        self.__internal.error_type = error_type
    
        return self
    
    def span_count(self, span_count: int) -> typing.Self:        
        self.__internal.span_count = span_count
    
        return self
    
    def points(self, points: list[list[typing.Union[str, int]]]) -> typing.Self:        
        self.__internal.points = points
    
        return self
    
    def drop_percent(self, drop_percent: float) -> typing.Self:        
        self.__internal.drop_percent = drop_percent
    
        return self
    
    def flamegraph_diff(self, flamegraph_diff: bool) -> typing.Self:        
        self.__internal.flamegraph_diff = flamegraph_diff
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:        
        self.__internal.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:        
        self.__internal.query_type = query_type
    
        return self
    
    def datasource(self, datasource: object) -> typing.Self:        
        self.__internal.datasource = datasource
    
        return self
    

class Key(cogbuilder.Builder[testdata.Key]):    
    __internal: testdata.Key

    def __init__(self):
        self.__internal = testdata.Key()

    def build(self) -> testdata.Key:
        return self.__internal    
    
    def type_val(self, type_val: str) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def tick(self, tick: float) -> typing.Self:        
        self.__internal.tick = tick
    
        return self
    
    def uid(self, uid: str) -> typing.Self:        
        self.__internal.uid = uid
    
        return self
    
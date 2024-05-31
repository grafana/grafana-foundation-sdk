# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import testdata


class StreamingQuery(cogbuilder.Builder[testdata.StreamingQuery]):    
    _internal: testdata.StreamingQuery

    def __init__(self):
        self._internal = testdata.StreamingQuery()

    def build(self) -> testdata.StreamingQuery:
        return self._internal    
    
    def type_val(self, type_val: typing.Literal["signal", "logs", "fetch"]) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def speed(self, speed: int) -> typing.Self:        
        self._internal.speed = speed
    
        return self
    
    def spread(self, spread: int) -> typing.Self:        
        self._internal.spread = spread
    
        return self
    
    def noise(self, noise: int) -> typing.Self:        
        self._internal.noise = noise
    
        return self
    
    def bands(self, bands: int) -> typing.Self:        
        self._internal.bands = bands
    
        return self
    
    def url(self, url: str) -> typing.Self:        
        self._internal.url = url
    
        return self
    

class PulseWaveQuery(cogbuilder.Builder[testdata.PulseWaveQuery]):    
    _internal: testdata.PulseWaveQuery

    def __init__(self):
        self._internal = testdata.PulseWaveQuery()

    def build(self) -> testdata.PulseWaveQuery:
        return self._internal    
    
    def time_step(self, time_step: int) -> typing.Self:        
        self._internal.time_step = time_step
    
        return self
    
    def on_count(self, on_count: int) -> typing.Self:        
        self._internal.on_count = on_count
    
        return self
    
    def off_count(self, off_count: int) -> typing.Self:        
        self._internal.off_count = off_count
    
        return self
    
    def on_value(self, on_value: float) -> typing.Self:        
        self._internal.on_value = on_value
    
        return self
    
    def off_value(self, off_value: float) -> typing.Self:        
        self._internal.off_value = off_value
    
        return self
    

class SimulationQuery(cogbuilder.Builder[testdata.SimulationQuery]):    
    _internal: testdata.SimulationQuery

    def __init__(self):
        self._internal = testdata.SimulationQuery()

    def build(self) -> testdata.SimulationQuery:
        return self._internal    
    
    def key(self, key: cogbuilder.Builder[testdata.Key]) -> typing.Self:        
        key_resource = key.build()
        self._internal.key = key_resource
    
        return self
    
    def config(self, config: dict[str, object]) -> typing.Self:        
        self._internal.config = config
    
        return self
    
    def stream(self, stream: bool) -> typing.Self:        
        self._internal.stream = stream
    
        return self
    
    def last(self, last: bool) -> typing.Self:        
        self._internal.last = last
    
        return self
    

class NodesQuery(cogbuilder.Builder[testdata.NodesQuery]):    
    _internal: testdata.NodesQuery

    def __init__(self):
        self._internal = testdata.NodesQuery()

    def build(self) -> testdata.NodesQuery:
        return self._internal    
    
    def type_val(self, type_val: typing.Literal["random", "response", "random edges"]) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def count(self, count: int) -> typing.Self:        
        self._internal.count = count
    
        return self
    

class USAQuery(cogbuilder.Builder[testdata.USAQuery]):    
    _internal: testdata.USAQuery

    def __init__(self):
        self._internal = testdata.USAQuery()

    def build(self) -> testdata.USAQuery:
        return self._internal    
    
    def mode(self, mode: str) -> typing.Self:        
        self._internal.mode = mode
    
        return self
    
    def period(self, period: str) -> typing.Self:        
        self._internal.period = period
    
        return self
    
    def fields(self, fields: list[str]) -> typing.Self:        
        self._internal.fields = fields
    
        return self
    
    def states(self, states: list[str]) -> typing.Self:        
        self._internal.states = states
    
        return self
    

class CSVWave(cogbuilder.Builder[testdata.CSVWave]):    
    _internal: testdata.CSVWave

    def __init__(self):
        self._internal = testdata.CSVWave()

    def build(self) -> testdata.CSVWave:
        return self._internal    
    
    def time_step(self, time_step: int) -> typing.Self:        
        self._internal.time_step = time_step
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def values_csv(self, values_csv: str) -> typing.Self:        
        self._internal.values_csv = values_csv
    
        return self
    
    def labels(self, labels: str) -> typing.Self:        
        self._internal.labels = labels
    
        return self
    

class Scenario(cogbuilder.Builder[testdata.Scenario]):    
    """
    TODO: Should this live here given it's not used in the dataquery?
    """
    
    _internal: testdata.Scenario

    def __init__(self):
        self._internal = testdata.Scenario()

    def build(self) -> testdata.Scenario:
        return self._internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self._internal.id_val = id_val
    
        return self
    
    def name(self, name: str) -> typing.Self:        
        self._internal.name = name
    
        return self
    
    def string_input(self, string_input: str) -> typing.Self:        
        self._internal.string_input = string_input
    
        return self
    
    def description(self, description: str) -> typing.Self:        
        self._internal.description = description
    
        return self
    
    def hide_alias_field(self, hide_alias_field: bool) -> typing.Self:        
        self._internal.hide_alias_field = hide_alias_field
    
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
    
    def scenario_id(self, scenario_id: testdata.TestDataQueryType) -> typing.Self:        
        self._internal.scenario_id = scenario_id
    
        return self
    
    def string_input(self, string_input: str) -> typing.Self:        
        self._internal.string_input = string_input
    
        return self
    
    def stream(self, stream: cogbuilder.Builder[testdata.StreamingQuery]) -> typing.Self:        
        stream_resource = stream.build()
        self._internal.stream = stream_resource
    
        return self
    
    def pulse_wave(self, pulse_wave: cogbuilder.Builder[testdata.PulseWaveQuery]) -> typing.Self:        
        pulse_wave_resource = pulse_wave.build()
        self._internal.pulse_wave = pulse_wave_resource
    
        return self
    
    def sim(self, sim: cogbuilder.Builder[testdata.SimulationQuery]) -> typing.Self:        
        sim_resource = sim.build()
        self._internal.sim = sim_resource
    
        return self
    
    def csv_wave(self, csv_wave: list[cogbuilder.Builder[testdata.CSVWave]]) -> typing.Self:        
        csv_wave_resources = [r1.build() for r1 in csv_wave]
        self._internal.csv_wave = csv_wave_resources
    
        return self
    
    def labels(self, labels: str) -> typing.Self:        
        self._internal.labels = labels
    
        return self
    
    def lines(self, lines: int) -> typing.Self:        
        self._internal.lines = lines
    
        return self
    
    def level_column(self, level_column: bool) -> typing.Self:        
        self._internal.level_column = level_column
    
        return self
    
    def channel(self, channel: str) -> typing.Self:        
        self._internal.channel = channel
    
        return self
    
    def nodes(self, nodes: cogbuilder.Builder[testdata.NodesQuery]) -> typing.Self:        
        nodes_resource = nodes.build()
        self._internal.nodes = nodes_resource
    
        return self
    
    def csv_file_name(self, csv_file_name: str) -> typing.Self:        
        self._internal.csv_file_name = csv_file_name
    
        return self
    
    def csv_content(self, csv_content: str) -> typing.Self:        
        self._internal.csv_content = csv_content
    
        return self
    
    def raw_frame_content(self, raw_frame_content: str) -> typing.Self:        
        self._internal.raw_frame_content = raw_frame_content
    
        return self
    
    def series_count(self, series_count: int) -> typing.Self:        
        self._internal.series_count = series_count
    
        return self
    
    def usa(self, usa: cogbuilder.Builder[testdata.USAQuery]) -> typing.Self:        
        usa_resource = usa.build()
        self._internal.usa = usa_resource
    
        return self
    
    def error_type(self, error_type: typing.Literal["server_panic", "frontend_exception", "frontend_observable"]) -> typing.Self:        
        self._internal.error_type = error_type
    
        return self
    
    def span_count(self, span_count: int) -> typing.Self:        
        self._internal.span_count = span_count
    
        return self
    
    def points(self, points: list[list[typing.Union[str, int]]]) -> typing.Self:        
        self._internal.points = points
    
        return self
    
    def drop_percent(self, drop_percent: float) -> typing.Self:    
        """
        Drop percentage (the chance we will lose a point 0-100)
        """
            
        self._internal.drop_percent = drop_percent
    
        return self
    
    def flamegraph_diff(self, flamegraph_diff: bool) -> typing.Self:        
        self._internal.flamegraph_diff = flamegraph_diff
    
        return self
    
    def ref_id(self, ref_id: str) -> typing.Self:    
        """
        A unique identifier for the query within the list of targets.
        In server side expressions, the refId is used as a variable name to identify results.
        By default, the UI will assign A->Z; however setting meaningful names may be useful.
        """
            
        self._internal.ref_id = ref_id
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:    
        """
        true if query is disabled (ie should not be returned to the dashboard)
        Note this does not always imply that the query should not be executed since
        the results from a hidden query may be used as the input to other queries (SSE etc)
        """
            
        self._internal.hide = hide
    
        return self
    
    def query_type(self, query_type: str) -> typing.Self:    
        """
        Specify the query flavor
        TODO make this required and give it a default
        """
            
        self._internal.query_type = query_type
    
        return self
    
    def datasource(self, datasource: object) -> typing.Self:    
        """
        For mixed data sources the selected datasource is on the query level.
        For non mixed scenarios this is undefined.
        TODO find a better way to do this ^ that's friendly to schema
        TODO this shouldn't be unknown but DataSourceRef | null
        """
            
        self._internal.datasource = datasource
    
        return self
    

class Key(cogbuilder.Builder[testdata.Key]):    
    _internal: testdata.Key

    def __init__(self):
        self._internal = testdata.Key()

    def build(self) -> testdata.Key:
        return self._internal    
    
    def type_val(self, type_val: str) -> typing.Self:        
        self._internal.type_val = type_val
    
        return self
    
    def tick(self, tick: float) -> typing.Self:        
        self._internal.tick = tick
    
        return self
    
    def uid(self, uid: str) -> typing.Self:        
        self._internal.uid = uid
    
        return self
    
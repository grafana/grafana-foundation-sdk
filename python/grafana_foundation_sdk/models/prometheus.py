# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import enum
from ..cog import variants as cogvariants
import typing
from ..cog import runtime as cogruntime


class QueryEditorMode(enum.StrEnum):
    CODE = "code"
    BUILDER = "builder"


class PromQueryFormat(enum.StrEnum):
    TIME_SERIES = "time_series"
    TABLE = "table"
    HEATMAP = "heatmap"


class Dataquery(cogvariants.Dataquery):
    # The actual expression/query that will be evaluated by Prometheus
    expr: str
    # Returns only the latest value that Prometheus has scraped for the requested time series
    instant: typing.Optional[bool]
    # Returns a Range vector, comprised of a set of time series containing a range of data points over time for each time series
    range_val: typing.Optional[bool]
    # Execute an additional query to identify interesting raw samples relevant for the given expr
    exemplar: typing.Optional[bool]
    # Specifies which editor is being used to prepare the query. It can be "code" or "builder"
    editor_mode: typing.Optional['QueryEditorMode']
    # Query format to determine how to display data points in panel. It can be "time_series", "table", "heatmap"
    format_val: typing.Optional['PromQueryFormat']
    # Series name override or template. Ex. {{hostname}} will be replaced with label value for hostname
    legend_format: typing.Optional[str]
    # @deprecated Used to specify how many times to divide max data points by. We use max data points under query options
    # See https://github.com/grafana/grafana/issues/48081
    interval_factor: typing.Optional[float]
    # A unique identifier for the query within the list of targets.
    # In server side expressions, the refId is used as a variable name to identify results.
    # By default, the UI will assign A->Z; however setting meaningful names may be useful.
    ref_id: str
    # If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
    hide: typing.Optional[bool]
    # Specify the query flavor
    # TODO make this required and give it a default
    query_type: typing.Optional[str]
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]
    # An additional lower limit for the step parameter of the Prometheus query and for the
    # `$__interval` and `$__rate_interval` variables.
    interval: typing.Optional[str]

    def __init__(self, expr: str = "", instant: typing.Optional[bool] = None, range_val: typing.Optional[bool] = None, exemplar: typing.Optional[bool] = None, editor_mode: typing.Optional['QueryEditorMode'] = None, format_val: typing.Optional['PromQueryFormat'] = None, legend_format: typing.Optional[str] = None, interval_factor: typing.Optional[float] = None, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[object] = None, interval: typing.Optional[str] = None):
        self.expr = expr
        self.instant = instant
        self.range_val = range_val
        self.exemplar = exemplar
        self.editor_mode = editor_mode
        self.format_val = format_val
        self.legend_format = legend_format
        self.interval_factor = interval_factor
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.datasource = datasource
        self.interval = interval

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "expr": self.expr,
            "refId": self.ref_id,
        }
        if self.instant is not None:
            payload["instant"] = self.instant
        if self.range_val is not None:
            payload["range"] = self.range_val
        if self.exemplar is not None:
            payload["exemplar"] = self.exemplar
        if self.editor_mode is not None:
            payload["editorMode"] = self.editor_mode
        if self.format_val is not None:
            payload["format"] = self.format_val
        if self.legend_format is not None:
            payload["legendFormat"] = self.legend_format
        if self.interval_factor is not None:
            payload["intervalFactor"] = self.interval_factor
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.interval is not None:
            payload["interval"] = self.interval
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "expr" in data:
            args["expr"] = data["expr"]
        if "instant" in data:
            args["instant"] = data["instant"]
        if "range" in data:
            args["range_val"] = data["range"]
        if "exemplar" in data:
            args["exemplar"] = data["exemplar"]
        if "editorMode" in data:
            args["editor_mode"] = data["editorMode"]
        if "format" in data:
            args["format_val"] = data["format"]
        if "legendFormat" in data:
            args["legend_format"] = data["legendFormat"]
        if "intervalFactor" in data:
            args["interval_factor"] = data["intervalFactor"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "datasource" in data:
            args["datasource"] = data["datasource"]
        if "interval" in data:
            args["interval"] = data["interval"]        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="prometheus",
        from_json_hook=Dataquery.from_json,
    )




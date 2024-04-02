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
    expr: typing.Optional[str]
    instant: typing.Optional[bool]
    range_val: typing.Optional[bool]
    exemplar: typing.Optional[bool]
    editor_mode: typing.Optional['QueryEditorMode']
    format_val: typing.Optional['PromQueryFormat']
    legend_format: typing.Optional[str]
    interval_factor: typing.Optional[float]
    ref_id: typing.Optional[str]
    hide: typing.Optional[bool]
    query_type: typing.Optional[str]
    datasource: typing.Optional[object]
    # An additional lower limit for the step parameter of the Prometheus query and for the
    # `$__interval` and `$__rate_interval` variables.
    interval: typing.Optional[str]

    def __init__(self, expr: typing.Optional[str] = None, instant: typing.Optional[bool] = None, range_val: typing.Optional[bool] = None, exemplar: typing.Optional[bool] = None, editor_mode: typing.Optional['QueryEditorMode'] = None, format_val: typing.Optional['PromQueryFormat'] = None, legend_format: typing.Optional[str] = None, interval_factor: typing.Optional[float] = None, ref_id: typing.Optional[str] = None, hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[object] = None, interval: typing.Optional[str] = None):
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
        }
        if self.expr is not None:
            payload["expr"] = self.expr
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
        if self.ref_id is not None:
            payload["refId"] = self.ref_id
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




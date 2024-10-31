# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
import enum
from ..cog import variants as cogvariants
from ..models import dashboard
from ..cog import runtime as cogruntime


BucketAggregation: typing.TypeAlias = typing.Union['DateHistogram', 'Histogram', 'Terms', 'Filters', 'GeoHashGrid', 'Nested']


MetricAggregation: typing.TypeAlias = typing.Union['Count', 'MovingAverage', 'Derivative', 'CumulativeSum', 'BucketScript', 'SerialDiff', 'RawData', 'RawDocument', 'UniqueCount', 'Percentiles', 'ExtendedStats', 'Min', 'Max', 'Sum', 'Average', 'MovingFunction', 'Logs', 'Rate', 'TopMetrics']


class BucketAggregationType(enum.StrEnum):
    TERMS = "terms"
    FILTERS = "filters"
    GEOHASH_GRID = "geohash_grid"
    DATE_HISTOGRAM = "date_histogram"
    HISTOGRAM = "histogram"
    NESTED = "nested"


class BaseBucketAggregation:
    id_val: str
    type_val: 'BucketAggregationType'
    settings: typing.Optional[object]

    def __init__(self, id_val: str = "", type_val: typing.Optional['BucketAggregationType'] = None, settings: typing.Optional[object] = None):
        self.id_val = id_val
        self.type_val = type_val if type_val is not None else BucketAggregationType.TERMS
        self.settings = settings

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "type": self.type_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "settings" in data:
            args["settings"] = data["settings"]        

        return cls(**args)


class BucketAggregationWithField:
    field: typing.Optional[str]
    id_val: str
    type_val: 'BucketAggregationType'
    settings: typing.Optional[object]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", type_val: typing.Optional['BucketAggregationType'] = None, settings: typing.Optional[object] = None):
        self.field = field
        self.id_val = id_val
        self.type_val = type_val if type_val is not None else BucketAggregationType.TERMS
        self.settings = settings

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "type": self.type_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "settings" in data:
            args["settings"] = data["settings"]        

        return cls(**args)


class DateHistogram:
    field: typing.Optional[str]
    id_val: str
    type_val: typing.Literal["date_histogram"]
    settings: typing.Optional['ElasticsearchDateHistogramSettings']

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchDateHistogramSettings'] = None):
        self.field = field
        self.id_val = id_val
        self.type_val = "date_histogram"
        self.settings = settings

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "type": self.type_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchDateHistogramSettings.from_json(data["settings"])        

        return cls(**args)


class DateHistogramSettings:
    interval: typing.Optional[str]
    min_doc_count: typing.Optional[str]
    trim_edges: typing.Optional[str]
    offset: typing.Optional[str]
    time_zone: typing.Optional[str]

    def __init__(self, interval: typing.Optional[str] = None, min_doc_count: typing.Optional[str] = None, trim_edges: typing.Optional[str] = None, offset: typing.Optional[str] = None, time_zone: typing.Optional[str] = None):
        self.interval = interval
        self.min_doc_count = min_doc_count
        self.trim_edges = trim_edges
        self.offset = offset
        self.time_zone = time_zone

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.interval is not None:
            payload["interval"] = self.interval
        if self.min_doc_count is not None:
            payload["min_doc_count"] = self.min_doc_count
        if self.trim_edges is not None:
            payload["trimEdges"] = self.trim_edges
        if self.offset is not None:
            payload["offset"] = self.offset
        if self.time_zone is not None:
            payload["timeZone"] = self.time_zone
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "interval" in data:
            args["interval"] = data["interval"]
        if "min_doc_count" in data:
            args["min_doc_count"] = data["min_doc_count"]
        if "trimEdges" in data:
            args["trim_edges"] = data["trimEdges"]
        if "offset" in data:
            args["offset"] = data["offset"]
        if "timeZone" in data:
            args["time_zone"] = data["timeZone"]        

        return cls(**args)


class Histogram:
    field: typing.Optional[str]
    id_val: str
    type_val: typing.Literal["histogram"]
    settings: typing.Optional['ElasticsearchHistogramSettings']

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchHistogramSettings'] = None):
        self.field = field
        self.id_val = id_val
        self.type_val = "histogram"
        self.settings = settings

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "type": self.type_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchHistogramSettings.from_json(data["settings"])        

        return cls(**args)


class HistogramSettings:
    interval: typing.Optional[str]
    min_doc_count: typing.Optional[str]

    def __init__(self, interval: typing.Optional[str] = None, min_doc_count: typing.Optional[str] = None):
        self.interval = interval
        self.min_doc_count = min_doc_count

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.interval is not None:
            payload["interval"] = self.interval
        if self.min_doc_count is not None:
            payload["min_doc_count"] = self.min_doc_count
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "interval" in data:
            args["interval"] = data["interval"]
        if "min_doc_count" in data:
            args["min_doc_count"] = data["min_doc_count"]        

        return cls(**args)


class TermsOrder(enum.StrEnum):
    DESC = "desc"
    ASC = "asc"


class Nested:
    field: typing.Optional[str]
    id_val: str
    type_val: typing.Literal["nested"]
    settings: typing.Optional[object]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional[object] = None):
        self.field = field
        self.id_val = id_val
        self.type_val = "nested"
        self.settings = settings

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "type": self.type_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = data["settings"]        

        return cls(**args)


class Terms:
    field: typing.Optional[str]
    id_val: str
    type_val: typing.Literal["terms"]
    settings: typing.Optional['ElasticsearchTermsSettings']

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchTermsSettings'] = None):
        self.field = field
        self.id_val = id_val
        self.type_val = "terms"
        self.settings = settings

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "type": self.type_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchTermsSettings.from_json(data["settings"])        

        return cls(**args)


class TermsSettings:
    order: typing.Optional['TermsOrder']
    size: typing.Optional[str]
    min_doc_count: typing.Optional[str]
    order_by: typing.Optional[str]
    missing: typing.Optional[str]

    def __init__(self, order: typing.Optional['TermsOrder'] = None, size: typing.Optional[str] = None, min_doc_count: typing.Optional[str] = None, order_by: typing.Optional[str] = None, missing: typing.Optional[str] = None):
        self.order = order
        self.size = size
        self.min_doc_count = min_doc_count
        self.order_by = order_by
        self.missing = missing

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.order is not None:
            payload["order"] = self.order
        if self.size is not None:
            payload["size"] = self.size
        if self.min_doc_count is not None:
            payload["min_doc_count"] = self.min_doc_count
        if self.order_by is not None:
            payload["orderBy"] = self.order_by
        if self.missing is not None:
            payload["missing"] = self.missing
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "order" in data:
            args["order"] = data["order"]
        if "size" in data:
            args["size"] = data["size"]
        if "min_doc_count" in data:
            args["min_doc_count"] = data["min_doc_count"]
        if "orderBy" in data:
            args["order_by"] = data["orderBy"]
        if "missing" in data:
            args["missing"] = data["missing"]        

        return cls(**args)


class Filters:
    id_val: str
    type_val: typing.Literal["filters"]
    settings: typing.Optional['ElasticsearchFiltersSettings']

    def __init__(self, id_val: str = "", settings: typing.Optional['ElasticsearchFiltersSettings'] = None):
        self.id_val = id_val
        self.type_val = "filters"
        self.settings = settings

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "type": self.type_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchFiltersSettings.from_json(data["settings"])        

        return cls(**args)


class Filter:
    query: str
    label: str

    def __init__(self, query: str = "", label: str = ""):
        self.query = query
        self.label = label

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "query": self.query,
            "label": self.label,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "query" in data:
            args["query"] = data["query"]
        if "label" in data:
            args["label"] = data["label"]        

        return cls(**args)


class FiltersSettings:
    filters: typing.Optional[list['Filter']]

    def __init__(self, filters: typing.Optional[list['Filter']] = None):
        self.filters = filters

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.filters is not None:
            payload["filters"] = self.filters
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "filters" in data:
            args["filters"] = data["filters"]        

        return cls(**args)


class GeoHashGrid:
    field: typing.Optional[str]
    id_val: str
    type_val: typing.Literal["geohash_grid"]
    settings: typing.Optional['ElasticsearchGeoHashGridSettings']

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchGeoHashGridSettings'] = None):
        self.field = field
        self.id_val = id_val
        self.type_val = "geohash_grid"
        self.settings = settings

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "id": self.id_val,
            "type": self.type_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchGeoHashGridSettings.from_json(data["settings"])        

        return cls(**args)


class GeoHashGridSettings:
    precision: typing.Optional[str]

    def __init__(self, precision: typing.Optional[str] = None):
        self.precision = precision

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.precision is not None:
            payload["precision"] = self.precision
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "precision" in data:
            args["precision"] = data["precision"]        

        return cls(**args)


class PipelineMetricAggregationType(enum.StrEnum):
    MOVING_AVG = "moving_avg"
    MOVING_FN = "moving_fn"
    DERIVATIVE = "derivative"
    SERIAL_DIFF = "serial_diff"
    CUMULATIVE_SUM = "cumulative_sum"
    BUCKET_SCRIPT = "bucket_script"


MetricAggregationType: typing.TypeAlias = typing.Union[typing.Literal["count"], 'PipelineMetricAggregationType']


class BaseMetricAggregation:
    type_val: 'MetricAggregationType'
    id_val: str
    hide: typing.Optional[bool]

    def __init__(self, type_val: typing.Optional['MetricAggregationType'] = None, id_val: str = "", hide: typing.Optional[bool] = None):
        self.type_val = type_val if type_val is not None else "count"
        self.id_val = id_val
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "type" in data:
            args["type_val"] = data["type"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class PipelineVariable:
    name: str
    pipeline_agg: str

    def __init__(self, name: str = "", pipeline_agg: str = ""):
        self.name = name
        self.pipeline_agg = pipeline_agg

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "name": self.name,
            "pipelineAgg": self.pipeline_agg,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "name" in data:
            args["name"] = data["name"]
        if "pipelineAgg" in data:
            args["pipeline_agg"] = data["pipelineAgg"]        

        return cls(**args)


class MetricAggregationWithField:
    field: typing.Optional[str]
    type_val: 'MetricAggregationType'
    id_val: str
    hide: typing.Optional[bool]

    def __init__(self, field: typing.Optional[str] = None, type_val: typing.Optional['MetricAggregationType'] = None, id_val: str = "", hide: typing.Optional[bool] = None):
        self.field = field
        self.type_val = type_val if type_val is not None else "count"
        self.id_val = id_val
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class MetricAggregationWithMissingSupport:
    settings: typing.Optional['ElasticsearchMetricAggregationWithMissingSupportSettings']
    type_val: 'MetricAggregationType'
    id_val: str
    hide: typing.Optional[bool]

    def __init__(self, settings: typing.Optional['ElasticsearchMetricAggregationWithMissingSupportSettings'] = None, type_val: typing.Optional['MetricAggregationType'] = None, id_val: str = "", hide: typing.Optional[bool] = None):
        self.settings = settings
        self.type_val = type_val if type_val is not None else "count"
        self.id_val = id_val
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "settings" in data:
            args["settings"] = ElasticsearchMetricAggregationWithMissingSupportSettings.from_json(data["settings"])
        if "type" in data:
            args["type_val"] = data["type"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


InlineScript: typing.TypeAlias = typing.Union[str, 'ElasticsearchInlineScript']


class MetricAggregationWithInlineScript:
    settings: typing.Optional['ElasticsearchMetricAggregationWithInlineScriptSettings']
    type_val: 'MetricAggregationType'
    id_val: str
    hide: typing.Optional[bool]

    def __init__(self, settings: typing.Optional['ElasticsearchMetricAggregationWithInlineScriptSettings'] = None, type_val: typing.Optional['MetricAggregationType'] = None, id_val: str = "", hide: typing.Optional[bool] = None):
        self.settings = settings
        self.type_val = type_val if type_val is not None else "count"
        self.id_val = id_val
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "settings" in data:
            args["settings"] = ElasticsearchMetricAggregationWithInlineScriptSettings.from_json(data["settings"])
        if "type" in data:
            args["type_val"] = data["type"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Count:
    type_val: typing.Literal["count"]
    id_val: str
    hide: typing.Optional[bool]

    def __init__(self, id_val: str = "", hide: typing.Optional[bool] = None):
        self.type_val = "count"
        self.id_val = id_val
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Average:
    type_val: typing.Literal["avg"]
    field: typing.Optional[str]
    id_val: str
    settings: typing.Optional['ElasticsearchAverageSettings']
    hide: typing.Optional[bool]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchAverageSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "avg"
        self.field = field
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchAverageSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Sum:
    type_val: typing.Literal["sum"]
    field: typing.Optional[str]
    id_val: str
    settings: typing.Optional['ElasticsearchSumSettings']
    hide: typing.Optional[bool]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchSumSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "sum"
        self.field = field
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchSumSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Max:
    type_val: typing.Literal["max"]
    field: typing.Optional[str]
    id_val: str
    settings: typing.Optional['ElasticsearchMaxSettings']
    hide: typing.Optional[bool]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchMaxSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "max"
        self.field = field
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchMaxSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Min:
    type_val: typing.Literal["min"]
    field: typing.Optional[str]
    id_val: str
    settings: typing.Optional['ElasticsearchMinSettings']
    hide: typing.Optional[bool]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchMinSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "min"
        self.field = field
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchMinSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class ExtendedStatMetaType(enum.StrEnum):
    AVG = "avg"
    MIN = "min"
    MAX = "max"
    SUM = "sum"
    COUNT = "count"
    STD_DEVIATION = "std_deviation"
    STD_DEVIATION_BOUNDS_UPPER = "std_deviation_bounds_upper"
    STD_DEVIATION_BOUNDS_LOWER = "std_deviation_bounds_lower"


class ExtendedStat:
    label: str
    value: 'ExtendedStatMetaType'

    def __init__(self, label: str = "", value: typing.Optional['ExtendedStatMetaType'] = None):
        self.label = label
        self.value = value if value is not None else ExtendedStatMetaType.AVG

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "label": self.label,
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "label" in data:
            args["label"] = data["label"]
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class ExtendedStats:
    type_val: typing.Literal["extended_stats"]
    settings: typing.Optional['ElasticsearchExtendedStatsSettings']
    field: typing.Optional[str]
    id_val: str
    meta: typing.Optional[object]
    hide: typing.Optional[bool]

    def __init__(self, settings: typing.Optional['ElasticsearchExtendedStatsSettings'] = None, field: typing.Optional[str] = None, id_val: str = "", meta: typing.Optional[object] = None, hide: typing.Optional[bool] = None):
        self.type_val = "extended_stats"
        self.settings = settings
        self.field = field
        self.id_val = id_val
        self.meta = meta
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.field is not None:
            payload["field"] = self.field
        if self.meta is not None:
            payload["meta"] = self.meta
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "settings" in data:
            args["settings"] = ElasticsearchExtendedStatsSettings.from_json(data["settings"])
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "meta" in data:
            args["meta"] = data["meta"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Percentiles:
    type_val: typing.Literal["percentiles"]
    field: typing.Optional[str]
    id_val: str
    settings: typing.Optional['ElasticsearchPercentilesSettings']
    hide: typing.Optional[bool]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchPercentilesSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "percentiles"
        self.field = field
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchPercentilesSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class UniqueCount:
    type_val: typing.Literal["cardinality"]
    field: typing.Optional[str]
    id_val: str
    settings: typing.Optional['ElasticsearchUniqueCountSettings']
    hide: typing.Optional[bool]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchUniqueCountSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "cardinality"
        self.field = field
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchUniqueCountSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class RawDocument:
    type_val: typing.Literal["raw_document"]
    id_val: str
    settings: typing.Optional['ElasticsearchRawDocumentSettings']
    hide: typing.Optional[bool]

    def __init__(self, id_val: str = "", settings: typing.Optional['ElasticsearchRawDocumentSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "raw_document"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchRawDocumentSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class RawData:
    type_val: typing.Literal["raw_data"]
    id_val: str
    settings: typing.Optional['ElasticsearchRawDataSettings']
    hide: typing.Optional[bool]

    def __init__(self, id_val: str = "", settings: typing.Optional['ElasticsearchRawDataSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "raw_data"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchRawDataSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Logs:
    type_val: typing.Literal["logs"]
    id_val: str
    settings: typing.Optional['ElasticsearchLogsSettings']
    hide: typing.Optional[bool]

    def __init__(self, id_val: str = "", settings: typing.Optional['ElasticsearchLogsSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "logs"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchLogsSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Rate:
    type_val: typing.Literal["rate"]
    field: typing.Optional[str]
    id_val: str
    settings: typing.Optional['ElasticsearchRateSettings']
    hide: typing.Optional[bool]

    def __init__(self, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchRateSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "rate"
        self.field = field
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchRateSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class BasePipelineMetricAggregation:
    pipeline_agg: typing.Optional[str]
    field: typing.Optional[str]
    type_val: str
    id_val: str
    hide: typing.Optional[bool]

    def __init__(self, pipeline_agg: typing.Optional[str] = None, field: typing.Optional[str] = None, type_val: str = "", id_val: str = "", hide: typing.Optional[bool] = None):
        self.pipeline_agg = pipeline_agg
        self.field = field
        self.type_val = type_val
        self.id_val = id_val
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.pipeline_agg is not None:
            payload["pipelineAgg"] = self.pipeline_agg
        if self.field is not None:
            payload["field"] = self.field
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pipelineAgg" in data:
            args["pipeline_agg"] = data["pipelineAgg"]
        if "field" in data:
            args["field"] = data["field"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class PipelineMetricAggregationWithMultipleBucketPaths:
    pipeline_variables: typing.Optional[list['PipelineVariable']]
    type_val: 'MetricAggregationType'
    id_val: str
    hide: typing.Optional[bool]

    def __init__(self, pipeline_variables: typing.Optional[list['PipelineVariable']] = None, type_val: typing.Optional['MetricAggregationType'] = None, id_val: str = "", hide: typing.Optional[bool] = None):
        self.pipeline_variables = pipeline_variables
        self.type_val = type_val if type_val is not None else "count"
        self.id_val = id_val
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.pipeline_variables is not None:
            payload["pipelineVariables"] = self.pipeline_variables
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pipelineVariables" in data:
            args["pipeline_variables"] = data["pipelineVariables"]
        if "type" in data:
            args["type_val"] = data["type"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class MovingAverageModel(enum.StrEnum):
    SIMPLE = "simple"
    LINEAR = "linear"
    EWMA = "ewma"
    HOLT = "holt"
    HOLT_WINTERS = "holt_winters"


class MovingAverageModelOption:
    label: str
    value: 'MovingAverageModel'

    def __init__(self, label: str = "", value: typing.Optional['MovingAverageModel'] = None):
        self.label = label
        self.value = value if value is not None else MovingAverageModel.SIMPLE

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "label": self.label,
            "value": self.value,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "label" in data:
            args["label"] = data["label"]
        if "value" in data:
            args["value"] = data["value"]        

        return cls(**args)


class BaseMovingAverageModelSettings:
    model: 'MovingAverageModel'
    window: str
    predict: str

    def __init__(self, model: typing.Optional['MovingAverageModel'] = None, window: str = "", predict: str = ""):
        self.model = model if model is not None else MovingAverageModel.SIMPLE
        self.window = window
        self.predict = predict

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "model": self.model,
            "window": self.window,
            "predict": self.predict,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "model" in data:
            args["model"] = data["model"]
        if "window" in data:
            args["window"] = data["window"]
        if "predict" in data:
            args["predict"] = data["predict"]        

        return cls(**args)


class MovingAverageSimpleModelSettings:
    model: typing.Literal["simple"]
    window: str
    predict: str

    def __init__(self, window: str = "", predict: str = ""):
        self.model = "simple"
        self.window = window
        self.predict = predict

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "model": self.model,
            "window": self.window,
            "predict": self.predict,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "window" in data:
            args["window"] = data["window"]
        if "predict" in data:
            args["predict"] = data["predict"]        

        return cls(**args)


class MovingAverageLinearModelSettings:
    model: typing.Literal["linear"]
    window: str
    predict: str

    def __init__(self, window: str = "", predict: str = ""):
        self.model = "linear"
        self.window = window
        self.predict = predict

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "model": self.model,
            "window": self.window,
            "predict": self.predict,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "window" in data:
            args["window"] = data["window"]
        if "predict" in data:
            args["predict"] = data["predict"]        

        return cls(**args)


class MovingAverageEWMAModelSettings:
    model: typing.Literal["ewma"]
    settings: typing.Optional['ElasticsearchMovingAverageEWMAModelSettingsSettings']
    window: str
    minimize: bool
    predict: str

    def __init__(self, settings: typing.Optional['ElasticsearchMovingAverageEWMAModelSettingsSettings'] = None, window: str = "", minimize: bool = False, predict: str = ""):
        self.model = "ewma"
        self.settings = settings
        self.window = window
        self.minimize = minimize
        self.predict = predict

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "model": self.model,
            "window": self.window,
            "minimize": self.minimize,
            "predict": self.predict,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "settings" in data:
            args["settings"] = ElasticsearchMovingAverageEWMAModelSettingsSettings.from_json(data["settings"])
        if "window" in data:
            args["window"] = data["window"]
        if "minimize" in data:
            args["minimize"] = data["minimize"]
        if "predict" in data:
            args["predict"] = data["predict"]        

        return cls(**args)


class MovingAverageHoltModelSettings:
    model: typing.Literal["holt"]
    settings: 'ElasticsearchMovingAverageHoltModelSettingsSettings'
    window: str
    minimize: bool
    predict: str

    def __init__(self, settings: typing.Optional['ElasticsearchMovingAverageHoltModelSettingsSettings'] = None, window: str = "", minimize: bool = False, predict: str = ""):
        self.model = "holt"
        self.settings = settings if settings is not None else ElasticsearchMovingAverageHoltModelSettingsSettings()
        self.window = window
        self.minimize = minimize
        self.predict = predict

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "model": self.model,
            "settings": self.settings,
            "window": self.window,
            "minimize": self.minimize,
            "predict": self.predict,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "settings" in data:
            args["settings"] = ElasticsearchMovingAverageHoltModelSettingsSettings.from_json(data["settings"])
        if "window" in data:
            args["window"] = data["window"]
        if "minimize" in data:
            args["minimize"] = data["minimize"]
        if "predict" in data:
            args["predict"] = data["predict"]        

        return cls(**args)


class MovingAverageHoltWintersModelSettings:
    model: typing.Literal["holt_winters"]
    settings: 'ElasticsearchMovingAverageHoltWintersModelSettingsSettings'
    window: str
    minimize: bool
    predict: str

    def __init__(self, settings: typing.Optional['ElasticsearchMovingAverageHoltWintersModelSettingsSettings'] = None, window: str = "", minimize: bool = False, predict: str = ""):
        self.model = "holt_winters"
        self.settings = settings if settings is not None else ElasticsearchMovingAverageHoltWintersModelSettingsSettings()
        self.window = window
        self.minimize = minimize
        self.predict = predict

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "model": self.model,
            "settings": self.settings,
            "window": self.window,
            "minimize": self.minimize,
            "predict": self.predict,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "settings" in data:
            args["settings"] = ElasticsearchMovingAverageHoltWintersModelSettingsSettings.from_json(data["settings"])
        if "window" in data:
            args["window"] = data["window"]
        if "minimize" in data:
            args["minimize"] = data["minimize"]
        if "predict" in data:
            args["predict"] = data["predict"]        

        return cls(**args)


class MovingAverage:
    """
    #MovingAverage's settings are overridden in types.ts
    """

    pipeline_agg: typing.Optional[str]
    field: typing.Optional[str]
    type_val: typing.Literal["moving_avg"]
    id_val: str
    settings: typing.Optional[dict[str, object]]
    hide: typing.Optional[bool]

    def __init__(self, pipeline_agg: typing.Optional[str] = None, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional[dict[str, object]] = None, hide: typing.Optional[bool] = None):
        self.pipeline_agg = pipeline_agg
        self.field = field
        self.type_val = "moving_avg"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.pipeline_agg is not None:
            payload["pipelineAgg"] = self.pipeline_agg
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pipelineAgg" in data:
            args["pipeline_agg"] = data["pipelineAgg"]
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = data["settings"]
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class MovingFunction:
    pipeline_agg: typing.Optional[str]
    field: typing.Optional[str]
    type_val: typing.Literal["moving_fn"]
    id_val: str
    settings: typing.Optional['ElasticsearchMovingFunctionSettings']
    hide: typing.Optional[bool]

    def __init__(self, pipeline_agg: typing.Optional[str] = None, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchMovingFunctionSettings'] = None, hide: typing.Optional[bool] = None):
        self.pipeline_agg = pipeline_agg
        self.field = field
        self.type_val = "moving_fn"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.pipeline_agg is not None:
            payload["pipelineAgg"] = self.pipeline_agg
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pipelineAgg" in data:
            args["pipeline_agg"] = data["pipelineAgg"]
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchMovingFunctionSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class Derivative:
    pipeline_agg: typing.Optional[str]
    field: typing.Optional[str]
    type_val: typing.Literal["derivative"]
    id_val: str
    settings: typing.Optional['ElasticsearchDerivativeSettings']
    hide: typing.Optional[bool]

    def __init__(self, pipeline_agg: typing.Optional[str] = None, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchDerivativeSettings'] = None, hide: typing.Optional[bool] = None):
        self.pipeline_agg = pipeline_agg
        self.field = field
        self.type_val = "derivative"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.pipeline_agg is not None:
            payload["pipelineAgg"] = self.pipeline_agg
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pipelineAgg" in data:
            args["pipeline_agg"] = data["pipelineAgg"]
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchDerivativeSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class SerialDiff:
    pipeline_agg: typing.Optional[str]
    field: typing.Optional[str]
    type_val: typing.Literal["serial_diff"]
    id_val: str
    settings: typing.Optional['ElasticsearchSerialDiffSettings']
    hide: typing.Optional[bool]

    def __init__(self, pipeline_agg: typing.Optional[str] = None, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchSerialDiffSettings'] = None, hide: typing.Optional[bool] = None):
        self.pipeline_agg = pipeline_agg
        self.field = field
        self.type_val = "serial_diff"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.pipeline_agg is not None:
            payload["pipelineAgg"] = self.pipeline_agg
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pipelineAgg" in data:
            args["pipeline_agg"] = data["pipelineAgg"]
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchSerialDiffSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class CumulativeSum:
    pipeline_agg: typing.Optional[str]
    field: typing.Optional[str]
    type_val: typing.Literal["cumulative_sum"]
    id_val: str
    settings: typing.Optional['ElasticsearchCumulativeSumSettings']
    hide: typing.Optional[bool]

    def __init__(self, pipeline_agg: typing.Optional[str] = None, field: typing.Optional[str] = None, id_val: str = "", settings: typing.Optional['ElasticsearchCumulativeSumSettings'] = None, hide: typing.Optional[bool] = None):
        self.pipeline_agg = pipeline_agg
        self.field = field
        self.type_val = "cumulative_sum"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.pipeline_agg is not None:
            payload["pipelineAgg"] = self.pipeline_agg
        if self.field is not None:
            payload["field"] = self.field
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pipelineAgg" in data:
            args["pipeline_agg"] = data["pipelineAgg"]
        if "field" in data:
            args["field"] = data["field"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchCumulativeSumSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class BucketScript:
    type_val: typing.Literal["bucket_script"]
    pipeline_variables: typing.Optional[list['PipelineVariable']]
    id_val: str
    settings: typing.Optional['ElasticsearchBucketScriptSettings']
    hide: typing.Optional[bool]

    def __init__(self, pipeline_variables: typing.Optional[list['PipelineVariable']] = None, id_val: str = "", settings: typing.Optional['ElasticsearchBucketScriptSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "bucket_script"
        self.pipeline_variables = pipeline_variables
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.pipeline_variables is not None:
            payload["pipelineVariables"] = self.pipeline_variables
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "pipelineVariables" in data:
            args["pipeline_variables"] = data["pipelineVariables"]
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchBucketScriptSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


class TopMetrics:
    type_val: typing.Literal["top_metrics"]
    id_val: str
    settings: typing.Optional['ElasticsearchTopMetricsSettings']
    hide: typing.Optional[bool]

    def __init__(self, id_val: str = "", settings: typing.Optional['ElasticsearchTopMetricsSettings'] = None, hide: typing.Optional[bool] = None):
        self.type_val = "top_metrics"
        self.id_val = id_val
        self.settings = settings
        self.hide = hide

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "type": self.type_val,
            "id": self.id_val,
        }
        if self.settings is not None:
            payload["settings"] = self.settings
        if self.hide is not None:
            payload["hide"] = self.hide
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "id" in data:
            args["id_val"] = data["id"]
        if "settings" in data:
            args["settings"] = ElasticsearchTopMetricsSettings.from_json(data["settings"])
        if "hide" in data:
            args["hide"] = data["hide"]        

        return cls(**args)


PipelineMetricAggregation: typing.TypeAlias = typing.Union['MovingAverage', 'Derivative', 'CumulativeSum', 'BucketScript']


MetricAggregationWithSettings: typing.TypeAlias = typing.Union['BucketScript', 'CumulativeSum', 'Derivative', 'SerialDiff', 'RawData', 'RawDocument', 'UniqueCount', 'Percentiles', 'ExtendedStats', 'Min', 'Max', 'Sum', 'Average', 'MovingAverage', 'MovingFunction', 'Logs', 'Rate', 'TopMetrics']


class Dataquery(cogvariants.Dataquery):
    # Alias pattern
    alias: typing.Optional[str]
    # Lucene query
    query: typing.Optional[str]
    # Name of time field
    time_field: typing.Optional[str]
    # List of bucket aggregations
    bucket_aggs: typing.Optional[list['BucketAggregation']]
    # List of metric aggregations
    metrics: typing.Optional[list['MetricAggregation']]
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
    datasource: typing.Optional[dashboard.DataSourceRef]

    def __init__(self, alias: typing.Optional[str] = None, query: typing.Optional[str] = None, time_field: typing.Optional[str] = None, bucket_aggs: typing.Optional[list['BucketAggregation']] = None, metrics: typing.Optional[list['MetricAggregation']] = None, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, datasource: typing.Optional[dashboard.DataSourceRef] = None):
        self.alias = alias
        self.query = query
        self.time_field = time_field
        self.bucket_aggs = bucket_aggs
        self.metrics = metrics
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
        if self.query is not None:
            payload["query"] = self.query
        if self.time_field is not None:
            payload["timeField"] = self.time_field
        if self.bucket_aggs is not None:
            payload["bucketAggs"] = self.bucket_aggs
        if self.metrics is not None:
            payload["metrics"] = self.metrics
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
        if "query" in data:
            args["query"] = data["query"]
        if "timeField" in data:
            args["time_field"] = data["timeField"]
        if "bucketAggs" in data:
            args["bucket_aggs"] = data["bucketAggs"]
        if "metrics" in data:
            args["metrics"] = data["metrics"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "datasource" in data:
            args["datasource"] = dashboard.DataSourceRef.from_json(data["datasource"])        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="elasticsearch",
        from_json_hook=Dataquery.from_json,
    )


class ElasticsearchDateHistogramSettings:
    interval: typing.Optional[str]
    min_doc_count: typing.Optional[str]
    trim_edges: typing.Optional[str]
    offset: typing.Optional[str]
    time_zone: typing.Optional[str]

    def __init__(self, interval: typing.Optional[str] = None, min_doc_count: typing.Optional[str] = None, trim_edges: typing.Optional[str] = None, offset: typing.Optional[str] = None, time_zone: typing.Optional[str] = None):
        self.interval = interval
        self.min_doc_count = min_doc_count
        self.trim_edges = trim_edges
        self.offset = offset
        self.time_zone = time_zone

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.interval is not None:
            payload["interval"] = self.interval
        if self.min_doc_count is not None:
            payload["min_doc_count"] = self.min_doc_count
        if self.trim_edges is not None:
            payload["trimEdges"] = self.trim_edges
        if self.offset is not None:
            payload["offset"] = self.offset
        if self.time_zone is not None:
            payload["timeZone"] = self.time_zone
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "interval" in data:
            args["interval"] = data["interval"]
        if "min_doc_count" in data:
            args["min_doc_count"] = data["min_doc_count"]
        if "trimEdges" in data:
            args["trim_edges"] = data["trimEdges"]
        if "offset" in data:
            args["offset"] = data["offset"]
        if "timeZone" in data:
            args["time_zone"] = data["timeZone"]        

        return cls(**args)


class ElasticsearchHistogramSettings:
    interval: typing.Optional[str]
    min_doc_count: typing.Optional[str]

    def __init__(self, interval: typing.Optional[str] = None, min_doc_count: typing.Optional[str] = None):
        self.interval = interval
        self.min_doc_count = min_doc_count

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.interval is not None:
            payload["interval"] = self.interval
        if self.min_doc_count is not None:
            payload["min_doc_count"] = self.min_doc_count
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "interval" in data:
            args["interval"] = data["interval"]
        if "min_doc_count" in data:
            args["min_doc_count"] = data["min_doc_count"]        

        return cls(**args)


class ElasticsearchTermsSettings:
    order: typing.Optional['TermsOrder']
    size: typing.Optional[str]
    min_doc_count: typing.Optional[str]
    order_by: typing.Optional[str]
    missing: typing.Optional[str]

    def __init__(self, order: typing.Optional['TermsOrder'] = None, size: typing.Optional[str] = None, min_doc_count: typing.Optional[str] = None, order_by: typing.Optional[str] = None, missing: typing.Optional[str] = None):
        self.order = order
        self.size = size
        self.min_doc_count = min_doc_count
        self.order_by = order_by
        self.missing = missing

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.order is not None:
            payload["order"] = self.order
        if self.size is not None:
            payload["size"] = self.size
        if self.min_doc_count is not None:
            payload["min_doc_count"] = self.min_doc_count
        if self.order_by is not None:
            payload["orderBy"] = self.order_by
        if self.missing is not None:
            payload["missing"] = self.missing
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "order" in data:
            args["order"] = data["order"]
        if "size" in data:
            args["size"] = data["size"]
        if "min_doc_count" in data:
            args["min_doc_count"] = data["min_doc_count"]
        if "orderBy" in data:
            args["order_by"] = data["orderBy"]
        if "missing" in data:
            args["missing"] = data["missing"]        

        return cls(**args)


class ElasticsearchFiltersSettings:
    filters: typing.Optional[list['Filter']]

    def __init__(self, filters: typing.Optional[list['Filter']] = None):
        self.filters = filters

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.filters is not None:
            payload["filters"] = self.filters
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "filters" in data:
            args["filters"] = data["filters"]        

        return cls(**args)


class ElasticsearchGeoHashGridSettings:
    precision: typing.Optional[str]

    def __init__(self, precision: typing.Optional[str] = None):
        self.precision = precision

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.precision is not None:
            payload["precision"] = self.precision
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "precision" in data:
            args["precision"] = data["precision"]        

        return cls(**args)


class ElasticsearchMetricAggregationWithMissingSupportSettings:
    missing: typing.Optional[str]

    def __init__(self, missing: typing.Optional[str] = None):
        self.missing = missing

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.missing is not None:
            payload["missing"] = self.missing
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "missing" in data:
            args["missing"] = data["missing"]        

        return cls(**args)


class ElasticsearchInlineScript:
    inline: typing.Optional[str]

    def __init__(self, inline: typing.Optional[str] = None):
        self.inline = inline

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.inline is not None:
            payload["inline"] = self.inline
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "inline" in data:
            args["inline"] = data["inline"]        

        return cls(**args)


class ElasticsearchMetricAggregationWithInlineScriptSettings:
    script: typing.Optional['InlineScript']

    def __init__(self, script: typing.Optional['InlineScript'] = None):
        self.script = script

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.script is not None:
            payload["script"] = self.script
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "script" in data:
            args["script"] = data["script"]        

        return cls(**args)


class ElasticsearchAverageSettings:
    script: typing.Optional['InlineScript']
    missing: typing.Optional[str]

    def __init__(self, script: typing.Optional['InlineScript'] = None, missing: typing.Optional[str] = None):
        self.script = script
        self.missing = missing

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.script is not None:
            payload["script"] = self.script
        if self.missing is not None:
            payload["missing"] = self.missing
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "script" in data:
            args["script"] = data["script"]
        if "missing" in data:
            args["missing"] = data["missing"]        

        return cls(**args)


class ElasticsearchSumSettings:
    script: typing.Optional['InlineScript']
    missing: typing.Optional[str]

    def __init__(self, script: typing.Optional['InlineScript'] = None, missing: typing.Optional[str] = None):
        self.script = script
        self.missing = missing

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.script is not None:
            payload["script"] = self.script
        if self.missing is not None:
            payload["missing"] = self.missing
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "script" in data:
            args["script"] = data["script"]
        if "missing" in data:
            args["missing"] = data["missing"]        

        return cls(**args)


class ElasticsearchMaxSettings:
    script: typing.Optional['InlineScript']
    missing: typing.Optional[str]

    def __init__(self, script: typing.Optional['InlineScript'] = None, missing: typing.Optional[str] = None):
        self.script = script
        self.missing = missing

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.script is not None:
            payload["script"] = self.script
        if self.missing is not None:
            payload["missing"] = self.missing
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "script" in data:
            args["script"] = data["script"]
        if "missing" in data:
            args["missing"] = data["missing"]        

        return cls(**args)


class ElasticsearchMinSettings:
    script: typing.Optional['InlineScript']
    missing: typing.Optional[str]

    def __init__(self, script: typing.Optional['InlineScript'] = None, missing: typing.Optional[str] = None):
        self.script = script
        self.missing = missing

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.script is not None:
            payload["script"] = self.script
        if self.missing is not None:
            payload["missing"] = self.missing
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "script" in data:
            args["script"] = data["script"]
        if "missing" in data:
            args["missing"] = data["missing"]        

        return cls(**args)


class ElasticsearchExtendedStatsSettings:
    script: typing.Optional['InlineScript']
    missing: typing.Optional[str]
    sigma: typing.Optional[str]

    def __init__(self, script: typing.Optional['InlineScript'] = None, missing: typing.Optional[str] = None, sigma: typing.Optional[str] = None):
        self.script = script
        self.missing = missing
        self.sigma = sigma

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.script is not None:
            payload["script"] = self.script
        if self.missing is not None:
            payload["missing"] = self.missing
        if self.sigma is not None:
            payload["sigma"] = self.sigma
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "script" in data:
            args["script"] = data["script"]
        if "missing" in data:
            args["missing"] = data["missing"]
        if "sigma" in data:
            args["sigma"] = data["sigma"]        

        return cls(**args)


class ElasticsearchPercentilesSettings:
    script: typing.Optional['InlineScript']
    missing: typing.Optional[str]
    percents: typing.Optional[list[str]]

    def __init__(self, script: typing.Optional['InlineScript'] = None, missing: typing.Optional[str] = None, percents: typing.Optional[list[str]] = None):
        self.script = script
        self.missing = missing
        self.percents = percents

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.script is not None:
            payload["script"] = self.script
        if self.missing is not None:
            payload["missing"] = self.missing
        if self.percents is not None:
            payload["percents"] = self.percents
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "script" in data:
            args["script"] = data["script"]
        if "missing" in data:
            args["missing"] = data["missing"]
        if "percents" in data:
            args["percents"] = data["percents"]        

        return cls(**args)


class ElasticsearchUniqueCountSettings:
    precision_threshold: typing.Optional[str]
    missing: typing.Optional[str]

    def __init__(self, precision_threshold: typing.Optional[str] = None, missing: typing.Optional[str] = None):
        self.precision_threshold = precision_threshold
        self.missing = missing

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.precision_threshold is not None:
            payload["precision_threshold"] = self.precision_threshold
        if self.missing is not None:
            payload["missing"] = self.missing
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "precision_threshold" in data:
            args["precision_threshold"] = data["precision_threshold"]
        if "missing" in data:
            args["missing"] = data["missing"]        

        return cls(**args)


class ElasticsearchRawDocumentSettings:
    size: typing.Optional[str]

    def __init__(self, size: typing.Optional[str] = None):
        self.size = size

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.size is not None:
            payload["size"] = self.size
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "size" in data:
            args["size"] = data["size"]        

        return cls(**args)


class ElasticsearchRawDataSettings:
    size: typing.Optional[str]

    def __init__(self, size: typing.Optional[str] = None):
        self.size = size

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.size is not None:
            payload["size"] = self.size
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "size" in data:
            args["size"] = data["size"]        

        return cls(**args)


class ElasticsearchLogsSettings:
    limit: typing.Optional[str]

    def __init__(self, limit: typing.Optional[str] = None):
        self.limit = limit

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.limit is not None:
            payload["limit"] = self.limit
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "limit" in data:
            args["limit"] = data["limit"]        

        return cls(**args)


class ElasticsearchRateSettings:
    unit: typing.Optional[str]
    mode: typing.Optional[str]

    def __init__(self, unit: typing.Optional[str] = None, mode: typing.Optional[str] = None):
        self.unit = unit
        self.mode = mode

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.unit is not None:
            payload["unit"] = self.unit
        if self.mode is not None:
            payload["mode"] = self.mode
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "unit" in data:
            args["unit"] = data["unit"]
        if "mode" in data:
            args["mode"] = data["mode"]        

        return cls(**args)


class ElasticsearchMovingAverageEWMAModelSettingsSettings:
    alpha: typing.Optional[str]

    def __init__(self, alpha: typing.Optional[str] = None):
        self.alpha = alpha

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.alpha is not None:
            payload["alpha"] = self.alpha
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "alpha" in data:
            args["alpha"] = data["alpha"]        

        return cls(**args)


class ElasticsearchMovingAverageHoltModelSettingsSettings:
    alpha: typing.Optional[str]
    beta: typing.Optional[str]

    def __init__(self, alpha: typing.Optional[str] = None, beta: typing.Optional[str] = None):
        self.alpha = alpha
        self.beta = beta

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.alpha is not None:
            payload["alpha"] = self.alpha
        if self.beta is not None:
            payload["beta"] = self.beta
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "alpha" in data:
            args["alpha"] = data["alpha"]
        if "beta" in data:
            args["beta"] = data["beta"]        

        return cls(**args)


class ElasticsearchMovingAverageHoltWintersModelSettingsSettings:
    alpha: typing.Optional[str]
    beta: typing.Optional[str]
    gamma: typing.Optional[str]
    period: typing.Optional[str]
    pad: typing.Optional[bool]

    def __init__(self, alpha: typing.Optional[str] = None, beta: typing.Optional[str] = None, gamma: typing.Optional[str] = None, period: typing.Optional[str] = None, pad: typing.Optional[bool] = None):
        self.alpha = alpha
        self.beta = beta
        self.gamma = gamma
        self.period = period
        self.pad = pad

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.alpha is not None:
            payload["alpha"] = self.alpha
        if self.beta is not None:
            payload["beta"] = self.beta
        if self.gamma is not None:
            payload["gamma"] = self.gamma
        if self.period is not None:
            payload["period"] = self.period
        if self.pad is not None:
            payload["pad"] = self.pad
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "alpha" in data:
            args["alpha"] = data["alpha"]
        if "beta" in data:
            args["beta"] = data["beta"]
        if "gamma" in data:
            args["gamma"] = data["gamma"]
        if "period" in data:
            args["period"] = data["period"]
        if "pad" in data:
            args["pad"] = data["pad"]        

        return cls(**args)


class ElasticsearchMovingFunctionSettings:
    window: typing.Optional[str]
    script: typing.Optional['InlineScript']
    shift: typing.Optional[str]

    def __init__(self, window: typing.Optional[str] = None, script: typing.Optional['InlineScript'] = None, shift: typing.Optional[str] = None):
        self.window = window
        self.script = script
        self.shift = shift

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.window is not None:
            payload["window"] = self.window
        if self.script is not None:
            payload["script"] = self.script
        if self.shift is not None:
            payload["shift"] = self.shift
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "window" in data:
            args["window"] = data["window"]
        if "script" in data:
            args["script"] = data["script"]
        if "shift" in data:
            args["shift"] = data["shift"]        

        return cls(**args)


class ElasticsearchDerivativeSettings:
    unit: typing.Optional[str]

    def __init__(self, unit: typing.Optional[str] = None):
        self.unit = unit

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.unit is not None:
            payload["unit"] = self.unit
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "unit" in data:
            args["unit"] = data["unit"]        

        return cls(**args)


class ElasticsearchSerialDiffSettings:
    lag: typing.Optional[str]

    def __init__(self, lag: typing.Optional[str] = None):
        self.lag = lag

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.lag is not None:
            payload["lag"] = self.lag
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "lag" in data:
            args["lag"] = data["lag"]        

        return cls(**args)


class ElasticsearchCumulativeSumSettings:
    format_val: typing.Optional[str]

    def __init__(self, format_val: typing.Optional[str] = None):
        self.format_val = format_val

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.format_val is not None:
            payload["format"] = self.format_val
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "format" in data:
            args["format_val"] = data["format"]        

        return cls(**args)


class ElasticsearchBucketScriptSettings:
    script: typing.Optional['InlineScript']

    def __init__(self, script: typing.Optional['InlineScript'] = None):
        self.script = script

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.script is not None:
            payload["script"] = self.script
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "script" in data:
            args["script"] = data["script"]        

        return cls(**args)


class ElasticsearchTopMetricsSettings:
    order: typing.Optional[str]
    order_by: typing.Optional[str]
    metrics: typing.Optional[list[str]]

    def __init__(self, order: typing.Optional[str] = None, order_by: typing.Optional[str] = None, metrics: typing.Optional[list[str]] = None):
        self.order = order
        self.order_by = order_by
        self.metrics = metrics

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
        }
        if self.order is not None:
            payload["order"] = self.order
        if self.order_by is not None:
            payload["orderBy"] = self.order_by
        if self.metrics is not None:
            payload["metrics"] = self.metrics
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "order" in data:
            args["order"] = data["order"]
        if "orderBy" in data:
            args["order_by"] = data["orderBy"]
        if "metrics" in data:
            args["metrics"] = data["metrics"]        

        return cls(**args)




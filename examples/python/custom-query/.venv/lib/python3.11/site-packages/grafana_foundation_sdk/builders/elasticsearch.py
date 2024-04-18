# Code generated - EDITING IS FUTILE. DO NOT EDIT.

import typing
from ..cog import builder as cogbuilder
from ..models import elasticsearch


class BaseBucketAggregation(cogbuilder.Builder[elasticsearch.BaseBucketAggregation]):    
    __internal: elasticsearch.BaseBucketAggregation

    def __init__(self):
        self.__internal = elasticsearch.BaseBucketAggregation()

    def build(self) -> elasticsearch.BaseBucketAggregation:
        return self.__internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def type_val(self, type_val: elasticsearch.BucketAggregationType) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def settings(self, settings: object) -> typing.Self:        
        self.__internal.settings = settings
    
        return self
    

class BucketAggregationWithField(cogbuilder.Builder[elasticsearch.BucketAggregationWithField]):    
    __internal: elasticsearch.BucketAggregationWithField

    def __init__(self):
        self.__internal = elasticsearch.BucketAggregationWithField()

    def build(self) -> elasticsearch.BucketAggregationWithField:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def type_val(self, type_val: elasticsearch.BucketAggregationType) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def settings(self, settings: object) -> typing.Self:        
        self.__internal.settings = settings
    
        return self
    

class DateHistogram(cogbuilder.Builder[elasticsearch.DateHistogram]):    
    __internal: elasticsearch.DateHistogram

    def __init__(self):
        self.__internal = elasticsearch.DateHistogram()        
        self.__internal.type_val = "date_histogram"

    def build(self) -> elasticsearch.DateHistogram:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchDateHistogramSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    

class DateHistogramSettings(cogbuilder.Builder[elasticsearch.DateHistogramSettings]):    
    __internal: elasticsearch.DateHistogramSettings

    def __init__(self):
        self.__internal = elasticsearch.DateHistogramSettings()

    def build(self) -> elasticsearch.DateHistogramSettings:
        return self.__internal    
    
    def interval(self, interval: str) -> typing.Self:        
        self.__internal.interval = interval
    
        return self
    
    def min_doc_count(self, min_doc_count: str) -> typing.Self:        
        self.__internal.min_doc_count = min_doc_count
    
        return self
    
    def trim_edges(self, trim_edges: str) -> typing.Self:        
        self.__internal.trim_edges = trim_edges
    
        return self
    
    def offset(self, offset: str) -> typing.Self:        
        self.__internal.offset = offset
    
        return self
    
    def time_zone(self, time_zone: str) -> typing.Self:        
        self.__internal.time_zone = time_zone
    
        return self
    

class Histogram(cogbuilder.Builder[elasticsearch.Histogram]):    
    __internal: elasticsearch.Histogram

    def __init__(self):
        self.__internal = elasticsearch.Histogram()        
        self.__internal.type_val = "histogram"

    def build(self) -> elasticsearch.Histogram:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchHistogramSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    

class HistogramSettings(cogbuilder.Builder[elasticsearch.HistogramSettings]):    
    __internal: elasticsearch.HistogramSettings

    def __init__(self):
        self.__internal = elasticsearch.HistogramSettings()

    def build(self) -> elasticsearch.HistogramSettings:
        return self.__internal    
    
    def interval(self, interval: str) -> typing.Self:        
        self.__internal.interval = interval
    
        return self
    
    def min_doc_count(self, min_doc_count: str) -> typing.Self:        
        self.__internal.min_doc_count = min_doc_count
    
        return self
    

class Nested(cogbuilder.Builder[elasticsearch.Nested]):    
    __internal: elasticsearch.Nested

    def __init__(self):
        self.__internal = elasticsearch.Nested()        
        self.__internal.type_val = "nested"

    def build(self) -> elasticsearch.Nested:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: object) -> typing.Self:        
        self.__internal.settings = settings
    
        return self
    

class Terms(cogbuilder.Builder[elasticsearch.Terms]):    
    __internal: elasticsearch.Terms

    def __init__(self):
        self.__internal = elasticsearch.Terms()        
        self.__internal.type_val = "terms"

    def build(self) -> elasticsearch.Terms:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchTermsSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    

class TermsSettings(cogbuilder.Builder[elasticsearch.TermsSettings]):    
    __internal: elasticsearch.TermsSettings

    def __init__(self):
        self.__internal = elasticsearch.TermsSettings()

    def build(self) -> elasticsearch.TermsSettings:
        return self.__internal    
    
    def order(self, order: elasticsearch.TermsOrder) -> typing.Self:        
        self.__internal.order = order
    
        return self
    
    def size(self, size: str) -> typing.Self:        
        self.__internal.size = size
    
        return self
    
    def min_doc_count(self, min_doc_count: str) -> typing.Self:        
        self.__internal.min_doc_count = min_doc_count
    
        return self
    
    def order_by(self, order_by: str) -> typing.Self:        
        self.__internal.order_by = order_by
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    

class Filters(cogbuilder.Builder[elasticsearch.Filters]):    
    __internal: elasticsearch.Filters

    def __init__(self):
        self.__internal = elasticsearch.Filters()        
        self.__internal.type_val = "filters"

    def build(self) -> elasticsearch.Filters:
        return self.__internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchFiltersSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    

class Filter(cogbuilder.Builder[elasticsearch.Filter]):    
    __internal: elasticsearch.Filter

    def __init__(self):
        self.__internal = elasticsearch.Filter()

    def build(self) -> elasticsearch.Filter:
        return self.__internal    
    
    def query(self, query: str) -> typing.Self:        
        self.__internal.query = query
    
        return self
    
    def label(self, label: str) -> typing.Self:        
        self.__internal.label = label
    
        return self
    

class FiltersSettings(cogbuilder.Builder[elasticsearch.FiltersSettings]):    
    __internal: elasticsearch.FiltersSettings

    def __init__(self):
        self.__internal = elasticsearch.FiltersSettings()

    def build(self) -> elasticsearch.FiltersSettings:
        return self.__internal    
    
    def filters(self, filters: list[cogbuilder.Builder[elasticsearch.Filter]]) -> typing.Self:        
        filters_resources = [r1.build() for r1 in filters]
        self.__internal.filters = filters_resources
    
        return self
    

class GeoHashGrid(cogbuilder.Builder[elasticsearch.GeoHashGrid]):    
    __internal: elasticsearch.GeoHashGrid

    def __init__(self):
        self.__internal = elasticsearch.GeoHashGrid()        
        self.__internal.type_val = "geohash_grid"

    def build(self) -> elasticsearch.GeoHashGrid:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchGeoHashGridSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    

class GeoHashGridSettings(cogbuilder.Builder[elasticsearch.GeoHashGridSettings]):    
    __internal: elasticsearch.GeoHashGridSettings

    def __init__(self):
        self.__internal = elasticsearch.GeoHashGridSettings()

    def build(self) -> elasticsearch.GeoHashGridSettings:
        return self.__internal    
    
    def precision(self, precision: str) -> typing.Self:        
        self.__internal.precision = precision
    
        return self
    

class BaseMetricAggregation(cogbuilder.Builder[elasticsearch.BaseMetricAggregation]):    
    __internal: elasticsearch.BaseMetricAggregation

    def __init__(self):
        self.__internal = elasticsearch.BaseMetricAggregation()

    def build(self) -> elasticsearch.BaseMetricAggregation:
        return self.__internal    
    
    def type_val(self, type_val: elasticsearch.MetricAggregationType) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class PipelineVariable(cogbuilder.Builder[elasticsearch.PipelineVariable]):    
    __internal: elasticsearch.PipelineVariable

    def __init__(self):
        self.__internal = elasticsearch.PipelineVariable()

    def build(self) -> elasticsearch.PipelineVariable:
        return self.__internal    
    
    def name(self, name: str) -> typing.Self:        
        self.__internal.name = name
    
        return self
    
    def pipeline_agg(self, pipeline_agg: str) -> typing.Self:        
        self.__internal.pipeline_agg = pipeline_agg
    
        return self
    

class MetricAggregationWithField(cogbuilder.Builder[elasticsearch.MetricAggregationWithField]):    
    __internal: elasticsearch.MetricAggregationWithField

    def __init__(self):
        self.__internal = elasticsearch.MetricAggregationWithField()

    def build(self) -> elasticsearch.MetricAggregationWithField:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def type_val(self, type_val: elasticsearch.MetricAggregationType) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class MetricAggregationWithMissingSupport(cogbuilder.Builder[elasticsearch.MetricAggregationWithMissingSupport]):    
    __internal: elasticsearch.MetricAggregationWithMissingSupport

    def __init__(self):
        self.__internal = elasticsearch.MetricAggregationWithMissingSupport()

    def build(self) -> elasticsearch.MetricAggregationWithMissingSupport:
        return self.__internal    
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchMetricAggregationWithMissingSupportSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def type_val(self, type_val: elasticsearch.MetricAggregationType) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class MetricAggregationWithInlineScript(cogbuilder.Builder[elasticsearch.MetricAggregationWithInlineScript]):    
    __internal: elasticsearch.MetricAggregationWithInlineScript

    def __init__(self):
        self.__internal = elasticsearch.MetricAggregationWithInlineScript()

    def build(self) -> elasticsearch.MetricAggregationWithInlineScript:
        return self.__internal    
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchMetricAggregationWithInlineScriptSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def type_val(self, type_val: elasticsearch.MetricAggregationType) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Count(cogbuilder.Builder[elasticsearch.Count]):    
    __internal: elasticsearch.Count

    def __init__(self):
        self.__internal = elasticsearch.Count()        
        self.__internal.type_val = "count"

    def build(self) -> elasticsearch.Count:
        return self.__internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Average(cogbuilder.Builder[elasticsearch.Average]):    
    __internal: elasticsearch.Average

    def __init__(self):
        self.__internal = elasticsearch.Average()        
        self.__internal.type_val = "avg"

    def build(self) -> elasticsearch.Average:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchAverageSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Sum(cogbuilder.Builder[elasticsearch.Sum]):    
    __internal: elasticsearch.Sum

    def __init__(self):
        self.__internal = elasticsearch.Sum()        
        self.__internal.type_val = "sum"

    def build(self) -> elasticsearch.Sum:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchSumSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Max(cogbuilder.Builder[elasticsearch.Max]):    
    __internal: elasticsearch.Max

    def __init__(self):
        self.__internal = elasticsearch.Max()        
        self.__internal.type_val = "max"

    def build(self) -> elasticsearch.Max:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchMaxSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Min(cogbuilder.Builder[elasticsearch.Min]):    
    __internal: elasticsearch.Min

    def __init__(self):
        self.__internal = elasticsearch.Min()        
        self.__internal.type_val = "min"

    def build(self) -> elasticsearch.Min:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchMinSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class ExtendedStat(cogbuilder.Builder[elasticsearch.ExtendedStat]):    
    __internal: elasticsearch.ExtendedStat

    def __init__(self):
        self.__internal = elasticsearch.ExtendedStat()

    def build(self) -> elasticsearch.ExtendedStat:
        return self.__internal    
    
    def label(self, label: str) -> typing.Self:        
        self.__internal.label = label
    
        return self
    
    def value(self, value: elasticsearch.ExtendedStatMetaType) -> typing.Self:        
        self.__internal.value = value
    
        return self
    

class ExtendedStats(cogbuilder.Builder[elasticsearch.ExtendedStats]):    
    __internal: elasticsearch.ExtendedStats

    def __init__(self):
        self.__internal = elasticsearch.ExtendedStats()        
        self.__internal.type_val = "extended_stats"

    def build(self) -> elasticsearch.ExtendedStats:
        return self.__internal    
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchExtendedStatsSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def meta(self, meta: object) -> typing.Self:        
        self.__internal.meta = meta
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Percentiles(cogbuilder.Builder[elasticsearch.Percentiles]):    
    __internal: elasticsearch.Percentiles

    def __init__(self):
        self.__internal = elasticsearch.Percentiles()        
        self.__internal.type_val = "percentiles"

    def build(self) -> elasticsearch.Percentiles:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchPercentilesSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class UniqueCount(cogbuilder.Builder[elasticsearch.UniqueCount]):    
    __internal: elasticsearch.UniqueCount

    def __init__(self):
        self.__internal = elasticsearch.UniqueCount()        
        self.__internal.type_val = "cardinality"

    def build(self) -> elasticsearch.UniqueCount:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchUniqueCountSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class RawDocument(cogbuilder.Builder[elasticsearch.RawDocument]):    
    __internal: elasticsearch.RawDocument

    def __init__(self):
        self.__internal = elasticsearch.RawDocument()        
        self.__internal.type_val = "raw_document"

    def build(self) -> elasticsearch.RawDocument:
        return self.__internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchRawDocumentSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class RawData(cogbuilder.Builder[elasticsearch.RawData]):    
    __internal: elasticsearch.RawData

    def __init__(self):
        self.__internal = elasticsearch.RawData()        
        self.__internal.type_val = "raw_data"

    def build(self) -> elasticsearch.RawData:
        return self.__internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchRawDataSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Logs(cogbuilder.Builder[elasticsearch.Logs]):    
    __internal: elasticsearch.Logs

    def __init__(self):
        self.__internal = elasticsearch.Logs()        
        self.__internal.type_val = "logs"

    def build(self) -> elasticsearch.Logs:
        return self.__internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchLogsSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Rate(cogbuilder.Builder[elasticsearch.Rate]):    
    __internal: elasticsearch.Rate

    def __init__(self):
        self.__internal = elasticsearch.Rate()        
        self.__internal.type_val = "rate"

    def build(self) -> elasticsearch.Rate:
        return self.__internal    
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchRateSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class BasePipelineMetricAggregation(cogbuilder.Builder[elasticsearch.BasePipelineMetricAggregation]):    
    __internal: elasticsearch.BasePipelineMetricAggregation

    def __init__(self):
        self.__internal = elasticsearch.BasePipelineMetricAggregation()

    def build(self) -> elasticsearch.BasePipelineMetricAggregation:
        return self.__internal    
    
    def pipeline_agg(self, pipeline_agg: str) -> typing.Self:        
        self.__internal.pipeline_agg = pipeline_agg
    
        return self
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def type_val(self, type_val: str) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class PipelineMetricAggregationWithMultipleBucketPaths(cogbuilder.Builder[elasticsearch.PipelineMetricAggregationWithMultipleBucketPaths]):    
    __internal: elasticsearch.PipelineMetricAggregationWithMultipleBucketPaths

    def __init__(self):
        self.__internal = elasticsearch.PipelineMetricAggregationWithMultipleBucketPaths()

    def build(self) -> elasticsearch.PipelineMetricAggregationWithMultipleBucketPaths:
        return self.__internal    
    
    def pipeline_variables(self, pipeline_variables: list[cogbuilder.Builder[elasticsearch.PipelineVariable]]) -> typing.Self:        
        pipeline_variables_resources = [r1.build() for r1 in pipeline_variables]
        self.__internal.pipeline_variables = pipeline_variables_resources
    
        return self
    
    def type_val(self, type_val: elasticsearch.MetricAggregationType) -> typing.Self:        
        self.__internal.type_val = type_val
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class MovingAverageModelOption(cogbuilder.Builder[elasticsearch.MovingAverageModelOption]):    
    __internal: elasticsearch.MovingAverageModelOption

    def __init__(self):
        self.__internal = elasticsearch.MovingAverageModelOption()

    def build(self) -> elasticsearch.MovingAverageModelOption:
        return self.__internal    
    
    def label(self, label: str) -> typing.Self:        
        self.__internal.label = label
    
        return self
    
    def value(self, value: elasticsearch.MovingAverageModel) -> typing.Self:        
        self.__internal.value = value
    
        return self
    

class BaseMovingAverageModelSettings(cogbuilder.Builder[elasticsearch.BaseMovingAverageModelSettings]):    
    __internal: elasticsearch.BaseMovingAverageModelSettings

    def __init__(self):
        self.__internal = elasticsearch.BaseMovingAverageModelSettings()

    def build(self) -> elasticsearch.BaseMovingAverageModelSettings:
        return self.__internal    
    
    def model(self, model: elasticsearch.MovingAverageModel) -> typing.Self:        
        self.__internal.model = model
    
        return self
    
    def window(self, window: str) -> typing.Self:        
        self.__internal.window = window
    
        return self
    
    def predict(self, predict: str) -> typing.Self:        
        self.__internal.predict = predict
    
        return self
    

class MovingAverageSimpleModelSettings(cogbuilder.Builder[elasticsearch.MovingAverageSimpleModelSettings]):    
    __internal: elasticsearch.MovingAverageSimpleModelSettings

    def __init__(self):
        self.__internal = elasticsearch.MovingAverageSimpleModelSettings()        
        self.__internal.model = "simple"

    def build(self) -> elasticsearch.MovingAverageSimpleModelSettings:
        return self.__internal    
    
    def window(self, window: str) -> typing.Self:        
        self.__internal.window = window
    
        return self
    
    def predict(self, predict: str) -> typing.Self:        
        self.__internal.predict = predict
    
        return self
    

class MovingAverageLinearModelSettings(cogbuilder.Builder[elasticsearch.MovingAverageLinearModelSettings]):    
    __internal: elasticsearch.MovingAverageLinearModelSettings

    def __init__(self):
        self.__internal = elasticsearch.MovingAverageLinearModelSettings()        
        self.__internal.model = "linear"

    def build(self) -> elasticsearch.MovingAverageLinearModelSettings:
        return self.__internal    
    
    def window(self, window: str) -> typing.Self:        
        self.__internal.window = window
    
        return self
    
    def predict(self, predict: str) -> typing.Self:        
        self.__internal.predict = predict
    
        return self
    

class MovingAverageEWMAModelSettings(cogbuilder.Builder[elasticsearch.MovingAverageEWMAModelSettings]):    
    __internal: elasticsearch.MovingAverageEWMAModelSettings

    def __init__(self):
        self.__internal = elasticsearch.MovingAverageEWMAModelSettings()        
        self.__internal.model = "ewma"

    def build(self) -> elasticsearch.MovingAverageEWMAModelSettings:
        return self.__internal    
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchMovingAverageEWMAModelSettingsSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def window(self, window: str) -> typing.Self:        
        self.__internal.window = window
    
        return self
    
    def minimize(self, minimize: bool) -> typing.Self:        
        self.__internal.minimize = minimize
    
        return self
    
    def predict(self, predict: str) -> typing.Self:        
        self.__internal.predict = predict
    
        return self
    

class MovingAverageHoltModelSettings(cogbuilder.Builder[elasticsearch.MovingAverageHoltModelSettings]):    
    __internal: elasticsearch.MovingAverageHoltModelSettings

    def __init__(self):
        self.__internal = elasticsearch.MovingAverageHoltModelSettings()        
        self.__internal.model = "holt"

    def build(self) -> elasticsearch.MovingAverageHoltModelSettings:
        return self.__internal    
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchMovingAverageHoltModelSettingsSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def window(self, window: str) -> typing.Self:        
        self.__internal.window = window
    
        return self
    
    def minimize(self, minimize: bool) -> typing.Self:        
        self.__internal.minimize = minimize
    
        return self
    
    def predict(self, predict: str) -> typing.Self:        
        self.__internal.predict = predict
    
        return self
    

class MovingAverageHoltWintersModelSettings(cogbuilder.Builder[elasticsearch.MovingAverageHoltWintersModelSettings]):    
    __internal: elasticsearch.MovingAverageHoltWintersModelSettings

    def __init__(self):
        self.__internal = elasticsearch.MovingAverageHoltWintersModelSettings()        
        self.__internal.model = "holt_winters"

    def build(self) -> elasticsearch.MovingAverageHoltWintersModelSettings:
        return self.__internal    
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchMovingAverageHoltWintersModelSettingsSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def window(self, window: str) -> typing.Self:        
        self.__internal.window = window
    
        return self
    
    def minimize(self, minimize: bool) -> typing.Self:        
        self.__internal.minimize = minimize
    
        return self
    
    def predict(self, predict: str) -> typing.Self:        
        self.__internal.predict = predict
    
        return self
    

class MovingAverage(cogbuilder.Builder[elasticsearch.MovingAverage]):    
    """
    #MovingAverage's settings are overridden in types.ts
    """
    
    __internal: elasticsearch.MovingAverage

    def __init__(self):
        self.__internal = elasticsearch.MovingAverage()        
        self.__internal.type_val = "moving_avg"

    def build(self) -> elasticsearch.MovingAverage:
        return self.__internal    
    
    def pipeline_agg(self, pipeline_agg: str) -> typing.Self:        
        self.__internal.pipeline_agg = pipeline_agg
    
        return self
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: object) -> typing.Self:        
        self.__internal.settings = settings
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class MovingFunction(cogbuilder.Builder[elasticsearch.MovingFunction]):    
    __internal: elasticsearch.MovingFunction

    def __init__(self):
        self.__internal = elasticsearch.MovingFunction()        
        self.__internal.type_val = "moving_fn"

    def build(self) -> elasticsearch.MovingFunction:
        return self.__internal    
    
    def pipeline_agg(self, pipeline_agg: str) -> typing.Self:        
        self.__internal.pipeline_agg = pipeline_agg
    
        return self
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchMovingFunctionSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Derivative(cogbuilder.Builder[elasticsearch.Derivative]):    
    __internal: elasticsearch.Derivative

    def __init__(self):
        self.__internal = elasticsearch.Derivative()        
        self.__internal.type_val = "derivative"

    def build(self) -> elasticsearch.Derivative:
        return self.__internal    
    
    def pipeline_agg(self, pipeline_agg: str) -> typing.Self:        
        self.__internal.pipeline_agg = pipeline_agg
    
        return self
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchDerivativeSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class SerialDiff(cogbuilder.Builder[elasticsearch.SerialDiff]):    
    __internal: elasticsearch.SerialDiff

    def __init__(self):
        self.__internal = elasticsearch.SerialDiff()        
        self.__internal.type_val = "serial_diff"

    def build(self) -> elasticsearch.SerialDiff:
        return self.__internal    
    
    def pipeline_agg(self, pipeline_agg: str) -> typing.Self:        
        self.__internal.pipeline_agg = pipeline_agg
    
        return self
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchSerialDiffSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class CumulativeSum(cogbuilder.Builder[elasticsearch.CumulativeSum]):    
    __internal: elasticsearch.CumulativeSum

    def __init__(self):
        self.__internal = elasticsearch.CumulativeSum()        
        self.__internal.type_val = "cumulative_sum"

    def build(self) -> elasticsearch.CumulativeSum:
        return self.__internal    
    
    def pipeline_agg(self, pipeline_agg: str) -> typing.Self:        
        self.__internal.pipeline_agg = pipeline_agg
    
        return self
    
    def field(self, field: str) -> typing.Self:        
        self.__internal.field = field
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchCumulativeSumSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class BucketScript(cogbuilder.Builder[elasticsearch.BucketScript]):    
    __internal: elasticsearch.BucketScript

    def __init__(self):
        self.__internal = elasticsearch.BucketScript()        
        self.__internal.type_val = "bucket_script"

    def build(self) -> elasticsearch.BucketScript:
        return self.__internal    
    
    def pipeline_variables(self, pipeline_variables: list[cogbuilder.Builder[elasticsearch.PipelineVariable]]) -> typing.Self:        
        pipeline_variables_resources = [r1.build() for r1 in pipeline_variables]
        self.__internal.pipeline_variables = pipeline_variables_resources
    
        return self
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchBucketScriptSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class TopMetrics(cogbuilder.Builder[elasticsearch.TopMetrics]):    
    __internal: elasticsearch.TopMetrics

    def __init__(self):
        self.__internal = elasticsearch.TopMetrics()        
        self.__internal.type_val = "top_metrics"

    def build(self) -> elasticsearch.TopMetrics:
        return self.__internal    
    
    def id_val(self, id_val: str) -> typing.Self:        
        self.__internal.id_val = id_val
    
        return self
    
    def settings(self, settings: cogbuilder.Builder[elasticsearch.ElasticsearchTopMetricsSettings]) -> typing.Self:        
        settings_resource = settings.build()
        self.__internal.settings = settings_resource
    
        return self
    
    def hide(self, hide: bool) -> typing.Self:        
        self.__internal.hide = hide
    
        return self
    

class Dataquery(cogbuilder.Builder[elasticsearch.Dataquery]):    
    __internal: elasticsearch.Dataquery

    def __init__(self):
        self.__internal = elasticsearch.Dataquery()

    def build(self) -> elasticsearch.Dataquery:
        return self.__internal    
    
    def alias(self, alias: str) -> typing.Self:        
        self.__internal.alias = alias
    
        return self
    
    def query(self, query: str) -> typing.Self:        
        self.__internal.query = query
    
        return self
    
    def time_field(self, time_field: str) -> typing.Self:        
        self.__internal.time_field = time_field
    
        return self
    
    def bucket_aggs(self, bucket_aggs: list[elasticsearch.BucketAggregation]) -> typing.Self:        
        self.__internal.bucket_aggs = bucket_aggs
    
        return self
    
    def metrics(self, metrics: list[elasticsearch.MetricAggregation]) -> typing.Self:        
        self.__internal.metrics = metrics
    
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
    

class ElasticsearchDateHistogramSettings(cogbuilder.Builder[elasticsearch.ElasticsearchDateHistogramSettings]):    
    __internal: elasticsearch.ElasticsearchDateHistogramSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchDateHistogramSettings()

    def build(self) -> elasticsearch.ElasticsearchDateHistogramSettings:
        return self.__internal    
    
    def interval(self, interval: str) -> typing.Self:        
        self.__internal.interval = interval
    
        return self
    
    def min_doc_count(self, min_doc_count: str) -> typing.Self:        
        self.__internal.min_doc_count = min_doc_count
    
        return self
    
    def trim_edges(self, trim_edges: str) -> typing.Self:        
        self.__internal.trim_edges = trim_edges
    
        return self
    
    def offset(self, offset: str) -> typing.Self:        
        self.__internal.offset = offset
    
        return self
    
    def time_zone(self, time_zone: str) -> typing.Self:        
        self.__internal.time_zone = time_zone
    
        return self
    

class ElasticsearchHistogramSettings(cogbuilder.Builder[elasticsearch.ElasticsearchHistogramSettings]):    
    __internal: elasticsearch.ElasticsearchHistogramSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchHistogramSettings()

    def build(self) -> elasticsearch.ElasticsearchHistogramSettings:
        return self.__internal    
    
    def interval(self, interval: str) -> typing.Self:        
        self.__internal.interval = interval
    
        return self
    
    def min_doc_count(self, min_doc_count: str) -> typing.Self:        
        self.__internal.min_doc_count = min_doc_count
    
        return self
    

class ElasticsearchTermsSettings(cogbuilder.Builder[elasticsearch.ElasticsearchTermsSettings]):    
    __internal: elasticsearch.ElasticsearchTermsSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchTermsSettings()

    def build(self) -> elasticsearch.ElasticsearchTermsSettings:
        return self.__internal    
    
    def order(self, order: elasticsearch.TermsOrder) -> typing.Self:        
        self.__internal.order = order
    
        return self
    
    def size(self, size: str) -> typing.Self:        
        self.__internal.size = size
    
        return self
    
    def min_doc_count(self, min_doc_count: str) -> typing.Self:        
        self.__internal.min_doc_count = min_doc_count
    
        return self
    
    def order_by(self, order_by: str) -> typing.Self:        
        self.__internal.order_by = order_by
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    

class ElasticsearchFiltersSettings(cogbuilder.Builder[elasticsearch.ElasticsearchFiltersSettings]):    
    __internal: elasticsearch.ElasticsearchFiltersSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchFiltersSettings()

    def build(self) -> elasticsearch.ElasticsearchFiltersSettings:
        return self.__internal    
    
    def filters(self, filters: list[cogbuilder.Builder[elasticsearch.Filter]]) -> typing.Self:        
        filters_resources = [r1.build() for r1 in filters]
        self.__internal.filters = filters_resources
    
        return self
    

class ElasticsearchGeoHashGridSettings(cogbuilder.Builder[elasticsearch.ElasticsearchGeoHashGridSettings]):    
    __internal: elasticsearch.ElasticsearchGeoHashGridSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchGeoHashGridSettings()

    def build(self) -> elasticsearch.ElasticsearchGeoHashGridSettings:
        return self.__internal    
    
    def precision(self, precision: str) -> typing.Self:        
        self.__internal.precision = precision
    
        return self
    

class ElasticsearchMetricAggregationWithMissingSupportSettings(cogbuilder.Builder[elasticsearch.ElasticsearchMetricAggregationWithMissingSupportSettings]):    
    __internal: elasticsearch.ElasticsearchMetricAggregationWithMissingSupportSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchMetricAggregationWithMissingSupportSettings()

    def build(self) -> elasticsearch.ElasticsearchMetricAggregationWithMissingSupportSettings:
        return self.__internal    
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    

class ElasticsearchInlineScript(cogbuilder.Builder[elasticsearch.ElasticsearchInlineScript]):    
    __internal: elasticsearch.ElasticsearchInlineScript

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchInlineScript()

    def build(self) -> elasticsearch.ElasticsearchInlineScript:
        return self.__internal    
    
    def inline(self, inline: str) -> typing.Self:        
        self.__internal.inline = inline
    
        return self
    

class ElasticsearchMetricAggregationWithInlineScriptSettings(cogbuilder.Builder[elasticsearch.ElasticsearchMetricAggregationWithInlineScriptSettings]):    
    __internal: elasticsearch.ElasticsearchMetricAggregationWithInlineScriptSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchMetricAggregationWithInlineScriptSettings()

    def build(self) -> elasticsearch.ElasticsearchMetricAggregationWithInlineScriptSettings:
        return self.__internal    
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    

class ElasticsearchAverageSettings(cogbuilder.Builder[elasticsearch.ElasticsearchAverageSettings]):    
    __internal: elasticsearch.ElasticsearchAverageSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchAverageSettings()

    def build(self) -> elasticsearch.ElasticsearchAverageSettings:
        return self.__internal    
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    

class ElasticsearchSumSettings(cogbuilder.Builder[elasticsearch.ElasticsearchSumSettings]):    
    __internal: elasticsearch.ElasticsearchSumSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchSumSettings()

    def build(self) -> elasticsearch.ElasticsearchSumSettings:
        return self.__internal    
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    

class ElasticsearchMaxSettings(cogbuilder.Builder[elasticsearch.ElasticsearchMaxSettings]):    
    __internal: elasticsearch.ElasticsearchMaxSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchMaxSettings()

    def build(self) -> elasticsearch.ElasticsearchMaxSettings:
        return self.__internal    
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    

class ElasticsearchMinSettings(cogbuilder.Builder[elasticsearch.ElasticsearchMinSettings]):    
    __internal: elasticsearch.ElasticsearchMinSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchMinSettings()

    def build(self) -> elasticsearch.ElasticsearchMinSettings:
        return self.__internal    
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    

class ElasticsearchExtendedStatsSettings(cogbuilder.Builder[elasticsearch.ElasticsearchExtendedStatsSettings]):    
    __internal: elasticsearch.ElasticsearchExtendedStatsSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchExtendedStatsSettings()

    def build(self) -> elasticsearch.ElasticsearchExtendedStatsSettings:
        return self.__internal    
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    
    def sigma(self, sigma: str) -> typing.Self:        
        self.__internal.sigma = sigma
    
        return self
    

class ElasticsearchPercentilesSettings(cogbuilder.Builder[elasticsearch.ElasticsearchPercentilesSettings]):    
    __internal: elasticsearch.ElasticsearchPercentilesSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchPercentilesSettings()

    def build(self) -> elasticsearch.ElasticsearchPercentilesSettings:
        return self.__internal    
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    
    def percents(self, percents: list[str]) -> typing.Self:        
        self.__internal.percents = percents
    
        return self
    

class ElasticsearchUniqueCountSettings(cogbuilder.Builder[elasticsearch.ElasticsearchUniqueCountSettings]):    
    __internal: elasticsearch.ElasticsearchUniqueCountSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchUniqueCountSettings()

    def build(self) -> elasticsearch.ElasticsearchUniqueCountSettings:
        return self.__internal    
    
    def precision_threshold(self, precision_threshold: str) -> typing.Self:        
        self.__internal.precision_threshold = precision_threshold
    
        return self
    
    def missing(self, missing: str) -> typing.Self:        
        self.__internal.missing = missing
    
        return self
    

class ElasticsearchRawDocumentSettings(cogbuilder.Builder[elasticsearch.ElasticsearchRawDocumentSettings]):    
    __internal: elasticsearch.ElasticsearchRawDocumentSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchRawDocumentSettings()

    def build(self) -> elasticsearch.ElasticsearchRawDocumentSettings:
        return self.__internal    
    
    def size(self, size: str) -> typing.Self:        
        self.__internal.size = size
    
        return self
    

class ElasticsearchRawDataSettings(cogbuilder.Builder[elasticsearch.ElasticsearchRawDataSettings]):    
    __internal: elasticsearch.ElasticsearchRawDataSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchRawDataSettings()

    def build(self) -> elasticsearch.ElasticsearchRawDataSettings:
        return self.__internal    
    
    def size(self, size: str) -> typing.Self:        
        self.__internal.size = size
    
        return self
    

class ElasticsearchLogsSettings(cogbuilder.Builder[elasticsearch.ElasticsearchLogsSettings]):    
    __internal: elasticsearch.ElasticsearchLogsSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchLogsSettings()

    def build(self) -> elasticsearch.ElasticsearchLogsSettings:
        return self.__internal    
    
    def limit(self, limit: str) -> typing.Self:        
        self.__internal.limit = limit
    
        return self
    

class ElasticsearchRateSettings(cogbuilder.Builder[elasticsearch.ElasticsearchRateSettings]):    
    __internal: elasticsearch.ElasticsearchRateSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchRateSettings()

    def build(self) -> elasticsearch.ElasticsearchRateSettings:
        return self.__internal    
    
    def unit(self, unit: str) -> typing.Self:        
        self.__internal.unit = unit
    
        return self
    
    def mode(self, mode: str) -> typing.Self:        
        self.__internal.mode = mode
    
        return self
    

class ElasticsearchMovingAverageEWMAModelSettingsSettings(cogbuilder.Builder[elasticsearch.ElasticsearchMovingAverageEWMAModelSettingsSettings]):    
    __internal: elasticsearch.ElasticsearchMovingAverageEWMAModelSettingsSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchMovingAverageEWMAModelSettingsSettings()

    def build(self) -> elasticsearch.ElasticsearchMovingAverageEWMAModelSettingsSettings:
        return self.__internal    
    
    def alpha(self, alpha: str) -> typing.Self:        
        self.__internal.alpha = alpha
    
        return self
    

class ElasticsearchMovingAverageHoltModelSettingsSettings(cogbuilder.Builder[elasticsearch.ElasticsearchMovingAverageHoltModelSettingsSettings]):    
    __internal: elasticsearch.ElasticsearchMovingAverageHoltModelSettingsSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchMovingAverageHoltModelSettingsSettings()

    def build(self) -> elasticsearch.ElasticsearchMovingAverageHoltModelSettingsSettings:
        return self.__internal    
    
    def alpha(self, alpha: str) -> typing.Self:        
        self.__internal.alpha = alpha
    
        return self
    
    def beta(self, beta: str) -> typing.Self:        
        self.__internal.beta = beta
    
        return self
    

class ElasticsearchMovingAverageHoltWintersModelSettingsSettings(cogbuilder.Builder[elasticsearch.ElasticsearchMovingAverageHoltWintersModelSettingsSettings]):    
    __internal: elasticsearch.ElasticsearchMovingAverageHoltWintersModelSettingsSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchMovingAverageHoltWintersModelSettingsSettings()

    def build(self) -> elasticsearch.ElasticsearchMovingAverageHoltWintersModelSettingsSettings:
        return self.__internal    
    
    def alpha(self, alpha: str) -> typing.Self:        
        self.__internal.alpha = alpha
    
        return self
    
    def beta(self, beta: str) -> typing.Self:        
        self.__internal.beta = beta
    
        return self
    
    def gamma(self, gamma: str) -> typing.Self:        
        self.__internal.gamma = gamma
    
        return self
    
    def period(self, period: str) -> typing.Self:        
        self.__internal.period = period
    
        return self
    
    def pad(self, pad: bool) -> typing.Self:        
        self.__internal.pad = pad
    
        return self
    

class ElasticsearchMovingFunctionSettings(cogbuilder.Builder[elasticsearch.ElasticsearchMovingFunctionSettings]):    
    __internal: elasticsearch.ElasticsearchMovingFunctionSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchMovingFunctionSettings()

    def build(self) -> elasticsearch.ElasticsearchMovingFunctionSettings:
        return self.__internal    
    
    def window(self, window: str) -> typing.Self:        
        self.__internal.window = window
    
        return self
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    
    def shift(self, shift: str) -> typing.Self:        
        self.__internal.shift = shift
    
        return self
    

class ElasticsearchDerivativeSettings(cogbuilder.Builder[elasticsearch.ElasticsearchDerivativeSettings]):    
    __internal: elasticsearch.ElasticsearchDerivativeSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchDerivativeSettings()

    def build(self) -> elasticsearch.ElasticsearchDerivativeSettings:
        return self.__internal    
    
    def unit(self, unit: str) -> typing.Self:        
        self.__internal.unit = unit
    
        return self
    

class ElasticsearchSerialDiffSettings(cogbuilder.Builder[elasticsearch.ElasticsearchSerialDiffSettings]):    
    __internal: elasticsearch.ElasticsearchSerialDiffSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchSerialDiffSettings()

    def build(self) -> elasticsearch.ElasticsearchSerialDiffSettings:
        return self.__internal    
    
    def lag(self, lag: str) -> typing.Self:        
        self.__internal.lag = lag
    
        return self
    

class ElasticsearchCumulativeSumSettings(cogbuilder.Builder[elasticsearch.ElasticsearchCumulativeSumSettings]):    
    __internal: elasticsearch.ElasticsearchCumulativeSumSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchCumulativeSumSettings()

    def build(self) -> elasticsearch.ElasticsearchCumulativeSumSettings:
        return self.__internal    
    
    def format_val(self, format_val: str) -> typing.Self:        
        self.__internal.format_val = format_val
    
        return self
    

class ElasticsearchBucketScriptSettings(cogbuilder.Builder[elasticsearch.ElasticsearchBucketScriptSettings]):    
    __internal: elasticsearch.ElasticsearchBucketScriptSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchBucketScriptSettings()

    def build(self) -> elasticsearch.ElasticsearchBucketScriptSettings:
        return self.__internal    
    
    def script(self, script: elasticsearch.InlineScript) -> typing.Self:        
        self.__internal.script = script
    
        return self
    

class ElasticsearchTopMetricsSettings(cogbuilder.Builder[elasticsearch.ElasticsearchTopMetricsSettings]):    
    __internal: elasticsearch.ElasticsearchTopMetricsSettings

    def __init__(self):
        self.__internal = elasticsearch.ElasticsearchTopMetricsSettings()

    def build(self) -> elasticsearch.ElasticsearchTopMetricsSettings:
        return self.__internal    
    
    def order(self, order: str) -> typing.Self:        
        self.__internal.order = order
    
        return self
    
    def order_by(self, order_by: str) -> typing.Self:        
        self.__internal.order_by = order_by
    
        return self
    
    def metrics(self, metrics: list[str]) -> typing.Self:        
        self.__internal.metrics = metrics
    
        return self
    
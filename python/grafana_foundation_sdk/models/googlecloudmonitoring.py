# Code generated - EDITING IS FUTILE. DO NOT EDIT.

from ..cog import variants as cogvariants
import typing
from ..cog import runtime as cogruntime
import enum


class CloudMonitoringQuery(cogvariants.Dataquery):
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
    # Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    alias_by: typing.Optional[str]
    # GCM query type.
    # queryType: #QueryType
    # Time Series List sub-query properties.
    time_series_list: typing.Optional['TimeSeriesList']
    # Time Series sub-query properties.
    time_series_query: typing.Optional['TimeSeriesQuery']
    # SLO sub-query properties.
    slo_query: typing.Optional['SLOQuery']
    # PromQL sub-query properties.
    prom_ql_query: typing.Optional['PromQLQuery']
    # For mixed data sources the selected datasource is on the query level.
    # For non mixed scenarios this is undefined.
    # TODO find a better way to do this ^ that's friendly to schema
    # TODO this shouldn't be unknown but DataSourceRef | null
    datasource: typing.Optional[object]
    # Time interval in milliseconds.
    interval_ms: typing.Optional[float]

    def __init__(self, ref_id: str = "", hide: typing.Optional[bool] = None, query_type: typing.Optional[str] = None, alias_by: typing.Optional[str] = None, time_series_list: typing.Optional['TimeSeriesList'] = None, time_series_query: typing.Optional['TimeSeriesQuery'] = None, slo_query: typing.Optional['SLOQuery'] = None, prom_ql_query: typing.Optional['PromQLQuery'] = None, datasource: typing.Optional[object] = None, interval_ms: typing.Optional[float] = None):
        self.ref_id = ref_id
        self.hide = hide
        self.query_type = query_type
        self.alias_by = alias_by
        self.time_series_list = time_series_list
        self.time_series_query = time_series_query
        self.slo_query = slo_query
        self.prom_ql_query = prom_ql_query
        self.datasource = datasource
        self.interval_ms = interval_ms

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "refId": self.ref_id,
        }
        if self.hide is not None:
            payload["hide"] = self.hide
        if self.query_type is not None:
            payload["queryType"] = self.query_type
        if self.alias_by is not None:
            payload["aliasBy"] = self.alias_by
        if self.time_series_list is not None:
            payload["timeSeriesList"] = self.time_series_list
        if self.time_series_query is not None:
            payload["timeSeriesQuery"] = self.time_series_query
        if self.slo_query is not None:
            payload["sloQuery"] = self.slo_query
        if self.prom_ql_query is not None:
            payload["promQLQuery"] = self.prom_ql_query
        if self.datasource is not None:
            payload["datasource"] = self.datasource
        if self.interval_ms is not None:
            payload["intervalMs"] = self.interval_ms
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "hide" in data:
            args["hide"] = data["hide"]
        if "queryType" in data:
            args["query_type"] = data["queryType"]
        if "aliasBy" in data:
            args["alias_by"] = data["aliasBy"]
        if "timeSeriesList" in data:
            args["time_series_list"] = TimeSeriesList.from_json(data["timeSeriesList"])
        if "timeSeriesQuery" in data:
            args["time_series_query"] = TimeSeriesQuery.from_json(data["timeSeriesQuery"])
        if "sloQuery" in data:
            args["slo_query"] = SLOQuery.from_json(data["sloQuery"])
        if "promQLQuery" in data:
            args["prom_ql_query"] = PromQLQuery.from_json(data["promQLQuery"])
        if "datasource" in data:
            args["datasource"] = data["datasource"]
        if "intervalMs" in data:
            args["interval_ms"] = data["intervalMs"]        

        return cls(**args)


def variant_config() -> cogruntime.DataqueryConfig:
    return cogruntime.DataqueryConfig(
        identifier="cloud-monitoring",
        from_json_hook=CloudMonitoringQuery.from_json,
    )


class QueryType(enum.StrEnum):
    """
    Defines the supported queryTypes.
    """

    TIME_SERIES_LIST = "timeSeriesList"
    TIME_SERIES_QUERY = "timeSeriesQuery"
    SLO = "slo"
    ANNOTATION = "annotation"
    PROMQL = "promQL"


class TimeSeriesList:
    """
    Time Series List sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    cross_series_reducer: str
    # Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignment_period: typing.Optional[str]
    # Alignment function to be used. Defaults to ALIGN_MEAN.
    per_series_aligner: typing.Optional[str]
    # Array of labels to group data by.
    group_bys: typing.Optional[list[str]]
    # Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    filters: typing.Optional[list[str]]
    # Data view, defaults to FULL.
    view: typing.Optional[str]
    # Annotation title.
    title: typing.Optional[str]
    # Annotation text.
    text: typing.Optional[str]
    # Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    secondary_cross_series_reducer: typing.Optional[str]
    # Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    secondary_alignment_period: typing.Optional[str]
    # Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
    secondary_per_series_aligner: typing.Optional[str]
    # Only present if a preprocessor is selected. Array of labels to group data by.
    secondary_group_bys: typing.Optional[list[str]]
    # Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    preprocessor: typing.Optional['PreprocessorType']

    def __init__(self, project_name: str = "", cross_series_reducer: str = "", alignment_period: typing.Optional[str] = None, per_series_aligner: typing.Optional[str] = None, group_bys: typing.Optional[list[str]] = None, filters: typing.Optional[list[str]] = None, view: typing.Optional[str] = None, title: typing.Optional[str] = None, text: typing.Optional[str] = None, secondary_cross_series_reducer: typing.Optional[str] = None, secondary_alignment_period: typing.Optional[str] = None, secondary_per_series_aligner: typing.Optional[str] = None, secondary_group_bys: typing.Optional[list[str]] = None, preprocessor: typing.Optional['PreprocessorType'] = None):
        self.project_name = project_name
        self.cross_series_reducer = cross_series_reducer
        self.alignment_period = alignment_period
        self.per_series_aligner = per_series_aligner
        self.group_bys = group_bys
        self.filters = filters
        self.view = view
        self.title = title
        self.text = text
        self.secondary_cross_series_reducer = secondary_cross_series_reducer
        self.secondary_alignment_period = secondary_alignment_period
        self.secondary_per_series_aligner = secondary_per_series_aligner
        self.secondary_group_bys = secondary_group_bys
        self.preprocessor = preprocessor

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "projectName": self.project_name,
            "crossSeriesReducer": self.cross_series_reducer,
        }
        if self.alignment_period is not None:
            payload["alignmentPeriod"] = self.alignment_period
        if self.per_series_aligner is not None:
            payload["perSeriesAligner"] = self.per_series_aligner
        if self.group_bys is not None:
            payload["groupBys"] = self.group_bys
        if self.filters is not None:
            payload["filters"] = self.filters
        if self.view is not None:
            payload["view"] = self.view
        if self.title is not None:
            payload["title"] = self.title
        if self.text is not None:
            payload["text"] = self.text
        if self.secondary_cross_series_reducer is not None:
            payload["secondaryCrossSeriesReducer"] = self.secondary_cross_series_reducer
        if self.secondary_alignment_period is not None:
            payload["secondaryAlignmentPeriod"] = self.secondary_alignment_period
        if self.secondary_per_series_aligner is not None:
            payload["secondaryPerSeriesAligner"] = self.secondary_per_series_aligner
        if self.secondary_group_bys is not None:
            payload["secondaryGroupBys"] = self.secondary_group_bys
        if self.preprocessor is not None:
            payload["preprocessor"] = self.preprocessor
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "projectName" in data:
            args["project_name"] = data["projectName"]
        if "crossSeriesReducer" in data:
            args["cross_series_reducer"] = data["crossSeriesReducer"]
        if "alignmentPeriod" in data:
            args["alignment_period"] = data["alignmentPeriod"]
        if "perSeriesAligner" in data:
            args["per_series_aligner"] = data["perSeriesAligner"]
        if "groupBys" in data:
            args["group_bys"] = data["groupBys"]
        if "filters" in data:
            args["filters"] = data["filters"]
        if "view" in data:
            args["view"] = data["view"]
        if "title" in data:
            args["title"] = data["title"]
        if "text" in data:
            args["text"] = data["text"]
        if "secondaryCrossSeriesReducer" in data:
            args["secondary_cross_series_reducer"] = data["secondaryCrossSeriesReducer"]
        if "secondaryAlignmentPeriod" in data:
            args["secondary_alignment_period"] = data["secondaryAlignmentPeriod"]
        if "secondaryPerSeriesAligner" in data:
            args["secondary_per_series_aligner"] = data["secondaryPerSeriesAligner"]
        if "secondaryGroupBys" in data:
            args["secondary_group_bys"] = data["secondaryGroupBys"]
        if "preprocessor" in data:
            args["preprocessor"] = data["preprocessor"]        

        return cls(**args)


class PreprocessorType(enum.StrEnum):
    """
    Types of pre-processor available. Defined by the metric.
    """

    NONE = "none"
    RATE = "rate"
    DELTA = "delta"


class TimeSeriesQuery:
    """
    Time Series sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # MQL query to be executed.
    query: str
    # To disable the graphPeriod, it should explictly be set to 'disabled'.
    graph_period: typing.Optional[str]

    def __init__(self, project_name: str = "", query: str = "", graph_period: typing.Optional[str] = "disabled"):
        self.project_name = project_name
        self.query = query
        self.graph_period = graph_period

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "projectName": self.project_name,
            "query": self.query,
        }
        if self.graph_period is not None:
            payload["graphPeriod"] = self.graph_period
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "projectName" in data:
            args["project_name"] = data["projectName"]
        if "query" in data:
            args["query"] = data["query"]
        if "graphPeriod" in data:
            args["graph_period"] = data["graphPeriod"]        

        return cls(**args)


class SLOQuery:
    """
    SLO sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # Alignment function to be used. Defaults to ALIGN_MEAN.
    per_series_aligner: typing.Optional[str]
    # Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignment_period: typing.Optional[str]
    # SLO selector.
    selector_name: str
    # ID for the service the SLO is in.
    service_id: str
    # Name for the service the SLO is in.
    service_name: str
    # ID for the SLO.
    slo_id: str
    # Name of the SLO.
    slo_name: str
    # SLO goal value.
    goal: typing.Optional[float]
    # Specific lookback period for the SLO.
    lookback_period: typing.Optional[str]

    def __init__(self, project_name: str = "", per_series_aligner: typing.Optional[str] = None, alignment_period: typing.Optional[str] = None, selector_name: str = "", service_id: str = "", service_name: str = "", slo_id: str = "", slo_name: str = "", goal: typing.Optional[float] = None, lookback_period: typing.Optional[str] = None):
        self.project_name = project_name
        self.per_series_aligner = per_series_aligner
        self.alignment_period = alignment_period
        self.selector_name = selector_name
        self.service_id = service_id
        self.service_name = service_name
        self.slo_id = slo_id
        self.slo_name = slo_name
        self.goal = goal
        self.lookback_period = lookback_period

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "projectName": self.project_name,
            "selectorName": self.selector_name,
            "serviceId": self.service_id,
            "serviceName": self.service_name,
            "sloId": self.slo_id,
            "sloName": self.slo_name,
        }
        if self.per_series_aligner is not None:
            payload["perSeriesAligner"] = self.per_series_aligner
        if self.alignment_period is not None:
            payload["alignmentPeriod"] = self.alignment_period
        if self.goal is not None:
            payload["goal"] = self.goal
        if self.lookback_period is not None:
            payload["lookbackPeriod"] = self.lookback_period
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "projectName" in data:
            args["project_name"] = data["projectName"]
        if "perSeriesAligner" in data:
            args["per_series_aligner"] = data["perSeriesAligner"]
        if "alignmentPeriod" in data:
            args["alignment_period"] = data["alignmentPeriod"]
        if "selectorName" in data:
            args["selector_name"] = data["selectorName"]
        if "serviceId" in data:
            args["service_id"] = data["serviceId"]
        if "serviceName" in data:
            args["service_name"] = data["serviceName"]
        if "sloId" in data:
            args["slo_id"] = data["sloId"]
        if "sloName" in data:
            args["slo_name"] = data["sloName"]
        if "goal" in data:
            args["goal"] = data["goal"]
        if "lookbackPeriod" in data:
            args["lookback_period"] = data["lookbackPeriod"]        

        return cls(**args)


class PromQLQuery:
    """
    PromQL sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # PromQL expression/query to be executed.
    expr: str
    # PromQL min step
    step: str

    def __init__(self, project_name: str = "", expr: str = "", step: str = ""):
        self.project_name = project_name
        self.expr = expr
        self.step = step

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "projectName": self.project_name,
            "expr": self.expr,
            "step": self.step,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "projectName" in data:
            args["project_name"] = data["projectName"]
        if "expr" in data:
            args["expr"] = data["expr"]
        if "step" in data:
            args["step"] = data["step"]        

        return cls(**args)


class MetricQuery:
    """
    @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
    """

    # GCP project to execute the query against.
    project_name: str
    # Alignment function to be used. Defaults to ALIGN_MEAN.
    per_series_aligner: typing.Optional[str]
    # Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
    alignment_period: typing.Optional[str]
    # Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
    alias_by: typing.Optional[str]
    editor_mode: str
    metric_type: str
    # Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
    cross_series_reducer: str
    # Array of labels to group data by.
    group_bys: typing.Optional[list[str]]
    # Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    filters: typing.Optional[list[str]]
    metric_kind: typing.Optional['MetricKind']
    value_type: typing.Optional[str]
    view: typing.Optional[str]
    # MQL query to be executed.
    query: str
    # Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
    preprocessor: typing.Optional['PreprocessorType']
    # To disable the graphPeriod, it should explictly be set to 'disabled'.
    graph_period: typing.Optional[str]

    def __init__(self, project_name: str = "", per_series_aligner: typing.Optional[str] = None, alignment_period: typing.Optional[str] = None, alias_by: typing.Optional[str] = None, editor_mode: str = "", metric_type: str = "", cross_series_reducer: str = "", group_bys: typing.Optional[list[str]] = None, filters: typing.Optional[list[str]] = None, metric_kind: typing.Optional['MetricKind'] = None, value_type: typing.Optional[str] = None, view: typing.Optional[str] = None, query: str = "", preprocessor: typing.Optional['PreprocessorType'] = None, graph_period: typing.Optional[str] = "disabled"):
        self.project_name = project_name
        self.per_series_aligner = per_series_aligner
        self.alignment_period = alignment_period
        self.alias_by = alias_by
        self.editor_mode = editor_mode
        self.metric_type = metric_type
        self.cross_series_reducer = cross_series_reducer
        self.group_bys = group_bys
        self.filters = filters
        self.metric_kind = metric_kind
        self.value_type = value_type
        self.view = view
        self.query = query
        self.preprocessor = preprocessor
        self.graph_period = graph_period

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "projectName": self.project_name,
            "editorMode": self.editor_mode,
            "metricType": self.metric_type,
            "crossSeriesReducer": self.cross_series_reducer,
            "query": self.query,
        }
        if self.per_series_aligner is not None:
            payload["perSeriesAligner"] = self.per_series_aligner
        if self.alignment_period is not None:
            payload["alignmentPeriod"] = self.alignment_period
        if self.alias_by is not None:
            payload["aliasBy"] = self.alias_by
        if self.group_bys is not None:
            payload["groupBys"] = self.group_bys
        if self.filters is not None:
            payload["filters"] = self.filters
        if self.metric_kind is not None:
            payload["metricKind"] = self.metric_kind
        if self.value_type is not None:
            payload["valueType"] = self.value_type
        if self.view is not None:
            payload["view"] = self.view
        if self.preprocessor is not None:
            payload["preprocessor"] = self.preprocessor
        if self.graph_period is not None:
            payload["graphPeriod"] = self.graph_period
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "projectName" in data:
            args["project_name"] = data["projectName"]
        if "perSeriesAligner" in data:
            args["per_series_aligner"] = data["perSeriesAligner"]
        if "alignmentPeriod" in data:
            args["alignment_period"] = data["alignmentPeriod"]
        if "aliasBy" in data:
            args["alias_by"] = data["aliasBy"]
        if "editorMode" in data:
            args["editor_mode"] = data["editorMode"]
        if "metricType" in data:
            args["metric_type"] = data["metricType"]
        if "crossSeriesReducer" in data:
            args["cross_series_reducer"] = data["crossSeriesReducer"]
        if "groupBys" in data:
            args["group_bys"] = data["groupBys"]
        if "filters" in data:
            args["filters"] = data["filters"]
        if "metricKind" in data:
            args["metric_kind"] = data["metricKind"]
        if "valueType" in data:
            args["value_type"] = data["valueType"]
        if "view" in data:
            args["view"] = data["view"]
        if "query" in data:
            args["query"] = data["query"]
        if "preprocessor" in data:
            args["preprocessor"] = data["preprocessor"]
        if "graphPeriod" in data:
            args["graph_period"] = data["graphPeriod"]        

        return cls(**args)


class MetricKind(enum.StrEnum):
    METRIC_KIND_UNSPECIFIED = "METRIC_KIND_UNSPECIFIED"
    GAUGE = "GAUGE"
    DELTA = "DELTA"
    CUMULATIVE = "CUMULATIVE"


class ValueTypes(enum.StrEnum):
    VALUE_TYPE_UNSPECIFIED = "VALUE_TYPE_UNSPECIFIED"
    BOOL = "BOOL"
    INT64 = "INT64"
    DOUBLE = "DOUBLE"
    STRING = "STRING"
    DISTRIBUTION = "DISTRIBUTION"
    MONEY = "MONEY"


class AlignmentTypes(enum.StrEnum):
    ALIGN_DELTA = "ALIGN_DELTA"
    ALIGN_RATE = "ALIGN_RATE"
    ALIGN_INTERPOLATE = "ALIGN_INTERPOLATE"
    ALIGN_NEXT_OLDER = "ALIGN_NEXT_OLDER"
    ALIGN_MIN = "ALIGN_MIN"
    ALIGN_MAX = "ALIGN_MAX"
    ALIGN_MEAN = "ALIGN_MEAN"
    ALIGN_COUNT = "ALIGN_COUNT"
    ALIGN_SUM = "ALIGN_SUM"
    ALIGN_STDDEV = "ALIGN_STDDEV"
    ALIGN_COUNT_TRUE = "ALIGN_COUNT_TRUE"
    ALIGN_COUNT_FALSE = "ALIGN_COUNT_FALSE"
    ALIGN_FRACTION_TRUE = "ALIGN_FRACTION_TRUE"
    ALIGN_PERCENTILE_99 = "ALIGN_PERCENTILE_99"
    ALIGN_PERCENTILE_95 = "ALIGN_PERCENTILE_95"
    ALIGN_PERCENTILE_50 = "ALIGN_PERCENTILE_50"
    ALIGN_PERCENTILE_05 = "ALIGN_PERCENTILE_05"
    ALIGN_PERCENT_CHANGE = "ALIGN_PERCENT_CHANGE"
    ALIGN_NONE = "ALIGN_NONE"


class LegacyCloudMonitoringAnnotationQuery:
    """
    @deprecated Use TimeSeriesList instead. Legacy annotation query properties for migration purposes.
    """

    # GCP project to execute the query against.
    project_name: str
    metric_type: str
    # Query refId.
    ref_id: str
    # Array of filters to query data by. Labels that can be filtered on are defined by the metric.
    filters: list[str]
    metric_kind: 'MetricKind'
    value_type: str
    # Annotation title.
    title: str
    # Annotation text.
    text: str

    def __init__(self, project_name: str = "", metric_type: str = "", ref_id: str = "", filters: typing.Optional[list[str]] = None, metric_kind: typing.Optional['MetricKind'] = None, value_type: str = "", title: str = "", text: str = ""):
        self.project_name = project_name
        self.metric_type = metric_type
        self.ref_id = ref_id
        self.filters = filters if filters is not None else []
        self.metric_kind = metric_kind if metric_kind is not None else MetricKind.METRIC_KIND_UNSPECIFIED
        self.value_type = value_type
        self.title = title
        self.text = text

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "projectName": self.project_name,
            "metricType": self.metric_type,
            "refId": self.ref_id,
            "filters": self.filters,
            "metricKind": self.metric_kind,
            "valueType": self.value_type,
            "title": self.title,
            "text": self.text,
        }
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "projectName" in data:
            args["project_name"] = data["projectName"]
        if "metricType" in data:
            args["metric_type"] = data["metricType"]
        if "refId" in data:
            args["ref_id"] = data["refId"]
        if "filters" in data:
            args["filters"] = data["filters"]
        if "metricKind" in data:
            args["metric_kind"] = data["metricKind"]
        if "valueType" in data:
            args["value_type"] = data["valueType"]
        if "title" in data:
            args["title"] = data["title"]
        if "text" in data:
            args["text"] = data["text"]        

        return cls(**args)


class Filter:
    """
    Query filter representation.
    """

    # Filter key.
    key: str
    # Filter operator.
    operator: str
    # Filter value.
    value: str
    # Filter condition.
    condition: typing.Optional[str]

    def __init__(self, key: str = "", operator: str = "", value: str = "", condition: typing.Optional[str] = None):
        self.key = key
        self.operator = operator
        self.value = value
        self.condition = condition

    def to_json(self) -> dict[str, object]:
        payload: dict[str, object] = {
            "key": self.key,
            "operator": self.operator,
            "value": self.value,
        }
        if self.condition is not None:
            payload["condition"] = self.condition
        return payload

    @classmethod
    def from_json(cls, data: dict[str, typing.Any]) -> typing.Self:
        args: dict[str, typing.Any] = {}
        
        if "key" in data:
            args["key"] = data["key"]
        if "operator" in data:
            args["operator"] = data["operator"]
        if "value" in data:
            args["value"] = data["value"]
        if "condition" in data:
            args["condition"] = data["condition"]        

        return cls(**args)


class MetricFindQueryTypes(enum.StrEnum):
    PROJECTS = "projects"
    SERVICES = "services"
    DEFAULT_PROJECT = "defaultProject"
    METRIC_TYPES = "metricTypes"
    LABEL_KEYS = "labelKeys"
    LABEL_VALUES = "labelValues"
    RESOURCE_TYPES = "resourceTypes"
    AGGREGATIONS = "aggregations"
    ALIGNERS = "aligners"
    ALIGNMENT_PERIODS = "alignmentPeriods"
    SELECTORS = "selectors"
    SLO_SERVICES = "sloServices"
    SLO = "slo"




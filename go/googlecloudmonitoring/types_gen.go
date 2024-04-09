// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

type CloudMonitoringQuery struct {
	// A unique identifier for the query within the list of targets.
	// In server side expressions, the refId is used as a variable name to identify results.
	// By default, the UI will assign A->Z; however setting meaningful names may be useful.
	RefId string `json:"refId"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// Note this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Specify the query flavor
	// TODO make this required and give it a default
	QueryType *string `json:"queryType,omitempty"`
	// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
	AliasBy *string `json:"aliasBy,omitempty"`
	// GCM query type.
	// queryType: #QueryType
	// Time Series List sub-query properties.
	TimeSeriesList *TimeSeriesList `json:"timeSeriesList,omitempty"`
	// Time Series sub-query properties.
	TimeSeriesQuery *TimeSeriesQuery `json:"timeSeriesQuery,omitempty"`
	// SLO sub-query properties.
	SloQuery *SLOQuery `json:"sloQuery,omitempty"`
	// PromQL sub-query properties.
	PromQLQuery *PromQLQuery `json:"promQLQuery,omitempty"`
	// For mixed data sources the selected datasource is on the query level.
	// For non mixed scenarios this is undefined.
	// TODO find a better way to do this ^ that's friendly to schema
	// TODO this shouldn't be unknown but DataSourceRef | null
	Datasource any `json:"datasource,omitempty"`
	// Time interval in milliseconds.
	IntervalMs *float64 `json:"intervalMs,omitempty"`
}

func (resource CloudMonitoringQuery) ImplementsDataqueryVariant() {}

// Defines the supported queryTypes.
type QueryType string

const (
	QueryTypeTIMESERIESLIST  QueryType = "timeSeriesList"
	QueryTypeTIMESERIESQUERY QueryType = "timeSeriesQuery"
	QueryTypeSLO             QueryType = "slo"
	QueryTypeANNOTATION      QueryType = "annotation"
	QueryTypePROMQL          QueryType = "promQL"
)

// Time Series List sub-query properties.
type TimeSeriesList struct {
	// GCP project to execute the query against.
	ProjectName string `json:"projectName"`
	// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	CrossSeriesReducer string `json:"crossSeriesReducer"`
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
	// Array of labels to group data by.
	GroupBys []string `json:"groupBys,omitempty"`
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	Filters []string `json:"filters,omitempty"`
	// Data view, defaults to FULL.
	View *string `json:"view,omitempty"`
	// Annotation title.
	Title *string `json:"title,omitempty"`
	// Annotation text.
	Text *string `json:"text,omitempty"`
	// Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	SecondaryCrossSeriesReducer *string `json:"secondaryCrossSeriesReducer,omitempty"`
	// Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	SecondaryAlignmentPeriod *string `json:"secondaryAlignmentPeriod,omitempty"`
	// Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
	SecondaryPerSeriesAligner *string `json:"secondaryPerSeriesAligner,omitempty"`
	// Only present if a preprocessor is selected. Array of labels to group data by.
	SecondaryGroupBys []string `json:"secondaryGroupBys,omitempty"`
	// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
	Preprocessor *PreprocessorType `json:"preprocessor,omitempty"`
}

// Types of pre-processor available. Defined by the metric.
type PreprocessorType string

const (
	PreprocessorTypeNone  PreprocessorType = "none"
	PreprocessorTypeRate  PreprocessorType = "rate"
	PreprocessorTypeDelta PreprocessorType = "delta"
)

// Time Series sub-query properties.
type TimeSeriesQuery struct {
	// GCP project to execute the query against.
	ProjectName string `json:"projectName"`
	// MQL query to be executed.
	Query string `json:"query"`
	// To disable the graphPeriod, it should explictly be set to 'disabled'.
	GraphPeriod string `json:"graphPeriod,omitempty"`
}

// SLO sub-query properties.
type SLOQuery struct {
	// GCP project to execute the query against.
	ProjectName string `json:"projectName"`
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`
	// SLO selector.
	SelectorName string `json:"selectorName"`
	// ID for the service the SLO is in.
	ServiceId string `json:"serviceId"`
	// Name for the service the SLO is in.
	ServiceName string `json:"serviceName"`
	// ID for the SLO.
	SloId string `json:"sloId"`
	// Name of the SLO.
	SloName string `json:"sloName"`
	// SLO goal value.
	Goal *float64 `json:"goal,omitempty"`
	// Specific lookback period for the SLO.
	LookbackPeriod *string `json:"lookbackPeriod,omitempty"`
}

// PromQL sub-query properties.
type PromQLQuery struct {
	// GCP project to execute the query against.
	ProjectName string `json:"projectName"`
	// PromQL expression/query to be executed.
	Expr string `json:"expr"`
	// PromQL min step
	Step string `json:"step"`
}

// @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
type MetricQuery struct {
	// GCP project to execute the query against.
	ProjectName string `json:"projectName"`
	// Alignment function to be used. Defaults to ALIGN_MEAN.
	PerSeriesAligner *string `json:"perSeriesAligner,omitempty"`
	// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
	AlignmentPeriod *string `json:"alignmentPeriod,omitempty"`
	// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
	AliasBy    *string `json:"aliasBy,omitempty"`
	EditorMode string  `json:"editorMode"`
	MetricType string  `json:"metricType"`
	// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
	CrossSeriesReducer string `json:"crossSeriesReducer"`
	// Array of labels to group data by.
	GroupBys []string `json:"groupBys,omitempty"`
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	Filters    []string    `json:"filters,omitempty"`
	MetricKind *MetricKind `json:"metricKind,omitempty"`
	ValueType  *string     `json:"valueType,omitempty"`
	View       *string     `json:"view,omitempty"`
	// MQL query to be executed.
	Query string `json:"query"`
	// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
	Preprocessor *PreprocessorType `json:"preprocessor,omitempty"`
	// To disable the graphPeriod, it should explictly be set to 'disabled'.
	GraphPeriod string `json:"graphPeriod,omitempty"`
}

type MetricKind string

const (
	MetricKindMETRICKINDUNSPECIFIED MetricKind = "METRIC_KIND_UNSPECIFIED"
	MetricKindGAUGE                 MetricKind = "GAUGE"
	MetricKindDELTA                 MetricKind = "DELTA"
	MetricKindCUMULATIVE            MetricKind = "CUMULATIVE"
)

type ValueTypes string

const (
	ValueTypesVALUETYPEUNSPECIFIED ValueTypes = "VALUE_TYPE_UNSPECIFIED"
	ValueTypesBOOL                 ValueTypes = "BOOL"
	ValueTypesINT64                ValueTypes = "INT64"
	ValueTypesDOUBLE               ValueTypes = "DOUBLE"
	ValueTypesSTRING               ValueTypes = "STRING"
	ValueTypesDISTRIBUTION         ValueTypes = "DISTRIBUTION"
	ValueTypesMONEY                ValueTypes = "MONEY"
)

type AlignmentTypes string

const (
	AlignmentTypesALIGNDELTA         AlignmentTypes = "ALIGN_DELTA"
	AlignmentTypesALIGNRATE          AlignmentTypes = "ALIGN_RATE"
	AlignmentTypesALIGNINTERPOLATE   AlignmentTypes = "ALIGN_INTERPOLATE"
	AlignmentTypesALIGNNEXTOLDER     AlignmentTypes = "ALIGN_NEXT_OLDER"
	AlignmentTypesALIGNMIN           AlignmentTypes = "ALIGN_MIN"
	AlignmentTypesALIGNMAX           AlignmentTypes = "ALIGN_MAX"
	AlignmentTypesALIGNMEAN          AlignmentTypes = "ALIGN_MEAN"
	AlignmentTypesALIGNCOUNT         AlignmentTypes = "ALIGN_COUNT"
	AlignmentTypesALIGNSUM           AlignmentTypes = "ALIGN_SUM"
	AlignmentTypesALIGNSTDDEV        AlignmentTypes = "ALIGN_STDDEV"
	AlignmentTypesALIGNCOUNTTRUE     AlignmentTypes = "ALIGN_COUNT_TRUE"
	AlignmentTypesALIGNCOUNTFALSE    AlignmentTypes = "ALIGN_COUNT_FALSE"
	AlignmentTypesALIGNFRACTIONTRUE  AlignmentTypes = "ALIGN_FRACTION_TRUE"
	AlignmentTypesALIGNPERCENTILE99  AlignmentTypes = "ALIGN_PERCENTILE_99"
	AlignmentTypesALIGNPERCENTILE95  AlignmentTypes = "ALIGN_PERCENTILE_95"
	AlignmentTypesALIGNPERCENTILE50  AlignmentTypes = "ALIGN_PERCENTILE_50"
	AlignmentTypesALIGNPERCENTILE05  AlignmentTypes = "ALIGN_PERCENTILE_05"
	AlignmentTypesALIGNPERCENTCHANGE AlignmentTypes = "ALIGN_PERCENT_CHANGE"
	AlignmentTypesALIGNNONE          AlignmentTypes = "ALIGN_NONE"
)

// @deprecated Use TimeSeriesList instead. Legacy annotation query properties for migration purposes.
type LegacyCloudMonitoringAnnotationQuery struct {
	// GCP project to execute the query against.
	ProjectName string `json:"projectName"`
	MetricType  string `json:"metricType"`
	// Query refId.
	RefId string `json:"refId"`
	// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
	Filters    []string   `json:"filters"`
	MetricKind MetricKind `json:"metricKind"`
	ValueType  string     `json:"valueType"`
	// Annotation title.
	Title string `json:"title"`
	// Annotation text.
	Text string `json:"text"`
}

// Query filter representation.
type Filter struct {
	// Filter key.
	Key string `json:"key"`
	// Filter operator.
	Operator string `json:"operator"`
	// Filter value.
	Value string `json:"value"`
	// Filter condition.
	Condition *string `json:"condition,omitempty"`
}

type MetricFindQueryTypes string

const (
	MetricFindQueryTypesProjects         MetricFindQueryTypes = "projects"
	MetricFindQueryTypesServices         MetricFindQueryTypes = "services"
	MetricFindQueryTypesDefaultProject   MetricFindQueryTypes = "defaultProject"
	MetricFindQueryTypesMetricTypes      MetricFindQueryTypes = "metricTypes"
	MetricFindQueryTypesLabelKeys        MetricFindQueryTypes = "labelKeys"
	MetricFindQueryTypesLabelValues      MetricFindQueryTypes = "labelValues"
	MetricFindQueryTypesResourceTypes    MetricFindQueryTypes = "resourceTypes"
	MetricFindQueryTypesAggregations     MetricFindQueryTypes = "aggregations"
	MetricFindQueryTypesAligners         MetricFindQueryTypes = "aligners"
	MetricFindQueryTypesAlignmentPeriods MetricFindQueryTypes = "alignmentPeriods"
	MetricFindQueryTypesSelectors        MetricFindQueryTypes = "selectors"
	MetricFindQueryTypesSLOServices      MetricFindQueryTypes = "sloServices"
	MetricFindQueryTypesSLO              MetricFindQueryTypes = "slo"
)
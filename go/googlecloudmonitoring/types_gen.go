// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"encoding/json"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

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
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// Time interval in milliseconds.
	IntervalMs *float64 `json:"intervalMs,omitempty"`
}

func (resource CloudMonitoringQuery) ImplementsDataqueryVariant() {}

func (resource CloudMonitoringQuery) DataqueryType() string {
	return "cloud-monitoring"
}

func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "cloud-monitoring",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &CloudMonitoringQuery{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery CloudMonitoringQuery
			if cast, ok := input.(*CloudMonitoringQuery); ok {
				dataquery = *cast
			} else {
				dataquery = input.(CloudMonitoringQuery)
			}
			return CloudMonitoringQueryConverter(dataquery)
		},
	}
}

func (resource CloudMonitoringQuery) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(CloudMonitoringQuery)
	if !ok {
		return false
	}
	if resource.RefId != other.RefId {
		return false
	}
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
			return false
		}
	}
	if resource.QueryType == nil && other.QueryType != nil || resource.QueryType != nil && other.QueryType == nil {
		return false
	}

	if resource.QueryType != nil {
		if *resource.QueryType != *other.QueryType {
			return false
		}
	}
	if resource.AliasBy == nil && other.AliasBy != nil || resource.AliasBy != nil && other.AliasBy == nil {
		return false
	}

	if resource.AliasBy != nil {
		if *resource.AliasBy != *other.AliasBy {
			return false
		}
	}
	if resource.TimeSeriesList == nil && other.TimeSeriesList != nil || resource.TimeSeriesList != nil && other.TimeSeriesList == nil {
		return false
	}

	if resource.TimeSeriesList != nil {
		if !resource.TimeSeriesList.Equals(*other.TimeSeriesList) {
			return false
		}
	}
	if resource.TimeSeriesQuery == nil && other.TimeSeriesQuery != nil || resource.TimeSeriesQuery != nil && other.TimeSeriesQuery == nil {
		return false
	}

	if resource.TimeSeriesQuery != nil {
		if !resource.TimeSeriesQuery.Equals(*other.TimeSeriesQuery) {
			return false
		}
	}
	if resource.SloQuery == nil && other.SloQuery != nil || resource.SloQuery != nil && other.SloQuery == nil {
		return false
	}

	if resource.SloQuery != nil {
		if !resource.SloQuery.Equals(*other.SloQuery) {
			return false
		}
	}
	if resource.PromQLQuery == nil && other.PromQLQuery != nil || resource.PromQLQuery != nil && other.PromQLQuery == nil {
		return false
	}

	if resource.PromQLQuery != nil {
		if !resource.PromQLQuery.Equals(*other.PromQLQuery) {
			return false
		}
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	if resource.IntervalMs == nil && other.IntervalMs != nil || resource.IntervalMs != nil && other.IntervalMs == nil {
		return false
	}

	if resource.IntervalMs != nil {
		if *resource.IntervalMs != *other.IntervalMs {
			return false
		}
	}

	return true
}

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

func (resource TimeSeriesList) Equals(other TimeSeriesList) bool {
	if resource.ProjectName != other.ProjectName {
		return false
	}
	if resource.CrossSeriesReducer != other.CrossSeriesReducer {
		return false
	}
	if resource.AlignmentPeriod == nil && other.AlignmentPeriod != nil || resource.AlignmentPeriod != nil && other.AlignmentPeriod == nil {
		return false
	}

	if resource.AlignmentPeriod != nil {
		if *resource.AlignmentPeriod != *other.AlignmentPeriod {
			return false
		}
	}
	if resource.PerSeriesAligner == nil && other.PerSeriesAligner != nil || resource.PerSeriesAligner != nil && other.PerSeriesAligner == nil {
		return false
	}

	if resource.PerSeriesAligner != nil {
		if *resource.PerSeriesAligner != *other.PerSeriesAligner {
			return false
		}
	}

	if len(resource.GroupBys) != len(other.GroupBys) {
		return false
	}

	for i1 := range resource.GroupBys {
		if resource.GroupBys[i1] != other.GroupBys[i1] {
			return false
		}
	}

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if resource.Filters[i1] != other.Filters[i1] {
			return false
		}
	}
	if resource.View == nil && other.View != nil || resource.View != nil && other.View == nil {
		return false
	}

	if resource.View != nil {
		if *resource.View != *other.View {
			return false
		}
	}
	if resource.Title == nil && other.Title != nil || resource.Title != nil && other.Title == nil {
		return false
	}

	if resource.Title != nil {
		if *resource.Title != *other.Title {
			return false
		}
	}
	if resource.Text == nil && other.Text != nil || resource.Text != nil && other.Text == nil {
		return false
	}

	if resource.Text != nil {
		if *resource.Text != *other.Text {
			return false
		}
	}
	if resource.SecondaryCrossSeriesReducer == nil && other.SecondaryCrossSeriesReducer != nil || resource.SecondaryCrossSeriesReducer != nil && other.SecondaryCrossSeriesReducer == nil {
		return false
	}

	if resource.SecondaryCrossSeriesReducer != nil {
		if *resource.SecondaryCrossSeriesReducer != *other.SecondaryCrossSeriesReducer {
			return false
		}
	}
	if resource.SecondaryAlignmentPeriod == nil && other.SecondaryAlignmentPeriod != nil || resource.SecondaryAlignmentPeriod != nil && other.SecondaryAlignmentPeriod == nil {
		return false
	}

	if resource.SecondaryAlignmentPeriod != nil {
		if *resource.SecondaryAlignmentPeriod != *other.SecondaryAlignmentPeriod {
			return false
		}
	}
	if resource.SecondaryPerSeriesAligner == nil && other.SecondaryPerSeriesAligner != nil || resource.SecondaryPerSeriesAligner != nil && other.SecondaryPerSeriesAligner == nil {
		return false
	}

	if resource.SecondaryPerSeriesAligner != nil {
		if *resource.SecondaryPerSeriesAligner != *other.SecondaryPerSeriesAligner {
			return false
		}
	}

	if len(resource.SecondaryGroupBys) != len(other.SecondaryGroupBys) {
		return false
	}

	for i1 := range resource.SecondaryGroupBys {
		if resource.SecondaryGroupBys[i1] != other.SecondaryGroupBys[i1] {
			return false
		}
	}
	if resource.Preprocessor == nil && other.Preprocessor != nil || resource.Preprocessor != nil && other.Preprocessor == nil {
		return false
	}

	if resource.Preprocessor != nil {
		if *resource.Preprocessor != *other.Preprocessor {
			return false
		}
	}

	return true
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
	GraphPeriod *string `json:"graphPeriod,omitempty"`
}

func (resource TimeSeriesQuery) Equals(other TimeSeriesQuery) bool {
	if resource.ProjectName != other.ProjectName {
		return false
	}
	if resource.Query != other.Query {
		return false
	}
	if resource.GraphPeriod == nil && other.GraphPeriod != nil || resource.GraphPeriod != nil && other.GraphPeriod == nil {
		return false
	}

	if resource.GraphPeriod != nil {
		if *resource.GraphPeriod != *other.GraphPeriod {
			return false
		}
	}

	return true
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

func (resource SLOQuery) Equals(other SLOQuery) bool {
	if resource.ProjectName != other.ProjectName {
		return false
	}
	if resource.PerSeriesAligner == nil && other.PerSeriesAligner != nil || resource.PerSeriesAligner != nil && other.PerSeriesAligner == nil {
		return false
	}

	if resource.PerSeriesAligner != nil {
		if *resource.PerSeriesAligner != *other.PerSeriesAligner {
			return false
		}
	}
	if resource.AlignmentPeriod == nil && other.AlignmentPeriod != nil || resource.AlignmentPeriod != nil && other.AlignmentPeriod == nil {
		return false
	}

	if resource.AlignmentPeriod != nil {
		if *resource.AlignmentPeriod != *other.AlignmentPeriod {
			return false
		}
	}
	if resource.SelectorName != other.SelectorName {
		return false
	}
	if resource.ServiceId != other.ServiceId {
		return false
	}
	if resource.ServiceName != other.ServiceName {
		return false
	}
	if resource.SloId != other.SloId {
		return false
	}
	if resource.SloName != other.SloName {
		return false
	}
	if resource.Goal == nil && other.Goal != nil || resource.Goal != nil && other.Goal == nil {
		return false
	}

	if resource.Goal != nil {
		if *resource.Goal != *other.Goal {
			return false
		}
	}
	if resource.LookbackPeriod == nil && other.LookbackPeriod != nil || resource.LookbackPeriod != nil && other.LookbackPeriod == nil {
		return false
	}

	if resource.LookbackPeriod != nil {
		if *resource.LookbackPeriod != *other.LookbackPeriod {
			return false
		}
	}

	return true
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

func (resource PromQLQuery) Equals(other PromQLQuery) bool {
	if resource.ProjectName != other.ProjectName {
		return false
	}
	if resource.Expr != other.Expr {
		return false
	}
	if resource.Step != other.Step {
		return false
	}

	return true
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
	GraphPeriod *string `json:"graphPeriod,omitempty"`
}

func (resource MetricQuery) Equals(other MetricQuery) bool {
	if resource.ProjectName != other.ProjectName {
		return false
	}
	if resource.PerSeriesAligner == nil && other.PerSeriesAligner != nil || resource.PerSeriesAligner != nil && other.PerSeriesAligner == nil {
		return false
	}

	if resource.PerSeriesAligner != nil {
		if *resource.PerSeriesAligner != *other.PerSeriesAligner {
			return false
		}
	}
	if resource.AlignmentPeriod == nil && other.AlignmentPeriod != nil || resource.AlignmentPeriod != nil && other.AlignmentPeriod == nil {
		return false
	}

	if resource.AlignmentPeriod != nil {
		if *resource.AlignmentPeriod != *other.AlignmentPeriod {
			return false
		}
	}
	if resource.AliasBy == nil && other.AliasBy != nil || resource.AliasBy != nil && other.AliasBy == nil {
		return false
	}

	if resource.AliasBy != nil {
		if *resource.AliasBy != *other.AliasBy {
			return false
		}
	}
	if resource.EditorMode != other.EditorMode {
		return false
	}
	if resource.MetricType != other.MetricType {
		return false
	}
	if resource.CrossSeriesReducer != other.CrossSeriesReducer {
		return false
	}

	if len(resource.GroupBys) != len(other.GroupBys) {
		return false
	}

	for i1 := range resource.GroupBys {
		if resource.GroupBys[i1] != other.GroupBys[i1] {
			return false
		}
	}

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if resource.Filters[i1] != other.Filters[i1] {
			return false
		}
	}
	if resource.MetricKind == nil && other.MetricKind != nil || resource.MetricKind != nil && other.MetricKind == nil {
		return false
	}

	if resource.MetricKind != nil {
		if *resource.MetricKind != *other.MetricKind {
			return false
		}
	}
	if resource.ValueType == nil && other.ValueType != nil || resource.ValueType != nil && other.ValueType == nil {
		return false
	}

	if resource.ValueType != nil {
		if *resource.ValueType != *other.ValueType {
			return false
		}
	}
	if resource.View == nil && other.View != nil || resource.View != nil && other.View == nil {
		return false
	}

	if resource.View != nil {
		if *resource.View != *other.View {
			return false
		}
	}
	if resource.Query != other.Query {
		return false
	}
	if resource.Preprocessor == nil && other.Preprocessor != nil || resource.Preprocessor != nil && other.Preprocessor == nil {
		return false
	}

	if resource.Preprocessor != nil {
		if *resource.Preprocessor != *other.Preprocessor {
			return false
		}
	}
	if resource.GraphPeriod == nil && other.GraphPeriod != nil || resource.GraphPeriod != nil && other.GraphPeriod == nil {
		return false
	}

	if resource.GraphPeriod != nil {
		if *resource.GraphPeriod != *other.GraphPeriod {
			return false
		}
	}

	return true
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

func (resource LegacyCloudMonitoringAnnotationQuery) Equals(other LegacyCloudMonitoringAnnotationQuery) bool {
	if resource.ProjectName != other.ProjectName {
		return false
	}
	if resource.MetricType != other.MetricType {
		return false
	}
	if resource.RefId != other.RefId {
		return false
	}

	if len(resource.Filters) != len(other.Filters) {
		return false
	}

	for i1 := range resource.Filters {
		if resource.Filters[i1] != other.Filters[i1] {
			return false
		}
	}
	if resource.MetricKind != other.MetricKind {
		return false
	}
	if resource.ValueType != other.ValueType {
		return false
	}
	if resource.Title != other.Title {
		return false
	}
	if resource.Text != other.Text {
		return false
	}

	return true
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

func (resource Filter) Equals(other Filter) bool {
	if resource.Key != other.Key {
		return false
	}
	if resource.Operator != other.Operator {
		return false
	}
	if resource.Value != other.Value {
		return false
	}
	if resource.Condition == nil && other.Condition != nil || resource.Condition != nil && other.Condition == nil {
		return false
	}

	if resource.Condition != nil {
		if *resource.Condition != *other.Condition {
			return false
		}
	}

	return true
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

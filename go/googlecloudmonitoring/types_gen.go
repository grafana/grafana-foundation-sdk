// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	"encoding/json"
	"errors"
	"fmt"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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

// VariantConfig returns the configuration related to cloud-monitoring dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
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
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &CloudMonitoringQuery{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CloudMonitoringQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *CloudMonitoringQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

	}
	// Field "queryType"
	if fields["queryType"] != nil {
		if string(fields["queryType"]) != "null" {
			if err := json.Unmarshal(fields["queryType"], &resource.QueryType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("queryType", err)...)
			}

		}
		delete(fields, "queryType")

	}
	// Field "aliasBy"
	if fields["aliasBy"] != nil {
		if string(fields["aliasBy"]) != "null" {
			if err := json.Unmarshal(fields["aliasBy"], &resource.AliasBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("aliasBy", err)...)
			}

		}
		delete(fields, "aliasBy")

	}
	// Field "timeSeriesList"
	if fields["timeSeriesList"] != nil {
		if string(fields["timeSeriesList"]) != "null" {

			resource.TimeSeriesList = &TimeSeriesList{}
			if err := resource.TimeSeriesList.UnmarshalJSONStrict(fields["timeSeriesList"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeSeriesList", err)...)
			}

		}
		delete(fields, "timeSeriesList")

	}
	// Field "timeSeriesQuery"
	if fields["timeSeriesQuery"] != nil {
		if string(fields["timeSeriesQuery"]) != "null" {

			resource.TimeSeriesQuery = &TimeSeriesQuery{}
			if err := resource.TimeSeriesQuery.UnmarshalJSONStrict(fields["timeSeriesQuery"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeSeriesQuery", err)...)
			}

		}
		delete(fields, "timeSeriesQuery")

	}
	// Field "sloQuery"
	if fields["sloQuery"] != nil {
		if string(fields["sloQuery"]) != "null" {

			resource.SloQuery = &SLOQuery{}
			if err := resource.SloQuery.UnmarshalJSONStrict(fields["sloQuery"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sloQuery", err)...)
			}

		}
		delete(fields, "sloQuery")

	}
	// Field "promQLQuery"
	if fields["promQLQuery"] != nil {
		if string(fields["promQLQuery"]) != "null" {

			resource.PromQLQuery = &PromQLQuery{}
			if err := resource.PromQLQuery.UnmarshalJSONStrict(fields["promQLQuery"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("promQLQuery", err)...)
			}

		}
		delete(fields, "promQLQuery")

	}
	// Field "datasource"
	if fields["datasource"] != nil {
		if string(fields["datasource"]) != "null" {

			resource.Datasource = &dashboard.DataSourceRef{}
			if err := resource.Datasource.UnmarshalJSONStrict(fields["datasource"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
			}

		}
		delete(fields, "datasource")

	}
	// Field "intervalMs"
	if fields["intervalMs"] != nil {
		if string(fields["intervalMs"]) != "null" {
			if err := json.Unmarshal(fields["intervalMs"], &resource.IntervalMs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intervalMs", err)...)
			}

		}
		delete(fields, "intervalMs")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("CloudMonitoringQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `CloudMonitoringQuery` fields for violations and returns them.
func (resource CloudMonitoringQuery) Validate() error {
	var errs cog.BuildErrors
	if resource.TimeSeriesList != nil {
		if err := resource.TimeSeriesList.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeSeriesList", err)...)
		}
	}
	if resource.TimeSeriesQuery != nil {
		if err := resource.TimeSeriesQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeSeriesQuery", err)...)
		}
	}
	if resource.SloQuery != nil {
		if err := resource.SloQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("sloQuery", err)...)
		}
	}
	if resource.PromQLQuery != nil {
		if err := resource.PromQLQuery.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("promQLQuery", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeSeriesList` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TimeSeriesList) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "projectName"
	if fields["projectName"] != nil {
		if string(fields["projectName"]) != "null" {
			if err := json.Unmarshal(fields["projectName"], &resource.ProjectName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("projectName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is null"))...)

		}
		delete(fields, "projectName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is missing from input"))...)
	}
	// Field "crossSeriesReducer"
	if fields["crossSeriesReducer"] != nil {
		if string(fields["crossSeriesReducer"]) != "null" {
			if err := json.Unmarshal(fields["crossSeriesReducer"], &resource.CrossSeriesReducer); err != nil {
				errs = append(errs, cog.MakeBuildErrors("crossSeriesReducer", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("crossSeriesReducer", errors.New("required field is null"))...)

		}
		delete(fields, "crossSeriesReducer")
	} else {
		errs = append(errs, cog.MakeBuildErrors("crossSeriesReducer", errors.New("required field is missing from input"))...)
	}
	// Field "alignmentPeriod"
	if fields["alignmentPeriod"] != nil {
		if string(fields["alignmentPeriod"]) != "null" {
			if err := json.Unmarshal(fields["alignmentPeriod"], &resource.AlignmentPeriod); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alignmentPeriod", err)...)
			}

		}
		delete(fields, "alignmentPeriod")

	}
	// Field "perSeriesAligner"
	if fields["perSeriesAligner"] != nil {
		if string(fields["perSeriesAligner"]) != "null" {
			if err := json.Unmarshal(fields["perSeriesAligner"], &resource.PerSeriesAligner); err != nil {
				errs = append(errs, cog.MakeBuildErrors("perSeriesAligner", err)...)
			}

		}
		delete(fields, "perSeriesAligner")

	}
	// Field "groupBys"
	if fields["groupBys"] != nil {
		if string(fields["groupBys"]) != "null" {

			if err := json.Unmarshal(fields["groupBys"], &resource.GroupBys); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupBys", err)...)
			}

		}
		delete(fields, "groupBys")

	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			if err := json.Unmarshal(fields["filters"], &resource.Filters); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filters", err)...)
			}

		}
		delete(fields, "filters")

	}
	// Field "view"
	if fields["view"] != nil {
		if string(fields["view"]) != "null" {
			if err := json.Unmarshal(fields["view"], &resource.View); err != nil {
				errs = append(errs, cog.MakeBuildErrors("view", err)...)
			}

		}
		delete(fields, "view")

	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}

		}
		delete(fields, "title")

	}
	// Field "text"
	if fields["text"] != nil {
		if string(fields["text"]) != "null" {
			if err := json.Unmarshal(fields["text"], &resource.Text); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}

		}
		delete(fields, "text")

	}
	// Field "secondaryCrossSeriesReducer"
	if fields["secondaryCrossSeriesReducer"] != nil {
		if string(fields["secondaryCrossSeriesReducer"]) != "null" {
			if err := json.Unmarshal(fields["secondaryCrossSeriesReducer"], &resource.SecondaryCrossSeriesReducer); err != nil {
				errs = append(errs, cog.MakeBuildErrors("secondaryCrossSeriesReducer", err)...)
			}

		}
		delete(fields, "secondaryCrossSeriesReducer")

	}
	// Field "secondaryAlignmentPeriod"
	if fields["secondaryAlignmentPeriod"] != nil {
		if string(fields["secondaryAlignmentPeriod"]) != "null" {
			if err := json.Unmarshal(fields["secondaryAlignmentPeriod"], &resource.SecondaryAlignmentPeriod); err != nil {
				errs = append(errs, cog.MakeBuildErrors("secondaryAlignmentPeriod", err)...)
			}

		}
		delete(fields, "secondaryAlignmentPeriod")

	}
	// Field "secondaryPerSeriesAligner"
	if fields["secondaryPerSeriesAligner"] != nil {
		if string(fields["secondaryPerSeriesAligner"]) != "null" {
			if err := json.Unmarshal(fields["secondaryPerSeriesAligner"], &resource.SecondaryPerSeriesAligner); err != nil {
				errs = append(errs, cog.MakeBuildErrors("secondaryPerSeriesAligner", err)...)
			}

		}
		delete(fields, "secondaryPerSeriesAligner")

	}
	// Field "secondaryGroupBys"
	if fields["secondaryGroupBys"] != nil {
		if string(fields["secondaryGroupBys"]) != "null" {

			if err := json.Unmarshal(fields["secondaryGroupBys"], &resource.SecondaryGroupBys); err != nil {
				errs = append(errs, cog.MakeBuildErrors("secondaryGroupBys", err)...)
			}

		}
		delete(fields, "secondaryGroupBys")

	}
	// Field "preprocessor"
	if fields["preprocessor"] != nil {
		if string(fields["preprocessor"]) != "null" {
			if err := json.Unmarshal(fields["preprocessor"], &resource.Preprocessor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("preprocessor", err)...)
			}

		}
		delete(fields, "preprocessor")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TimeSeriesList", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TimeSeriesList` objects.
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

// Validate checks all the validation constraints that may be defined on `TimeSeriesList` fields for violations and returns them.
func (resource TimeSeriesList) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TimeSeriesQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TimeSeriesQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "projectName"
	if fields["projectName"] != nil {
		if string(fields["projectName"]) != "null" {
			if err := json.Unmarshal(fields["projectName"], &resource.ProjectName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("projectName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is null"))...)

		}
		delete(fields, "projectName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is missing from input"))...)
	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")
	} else {
		errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is missing from input"))...)
	}
	// Field "graphPeriod"
	if fields["graphPeriod"] != nil {
		if string(fields["graphPeriod"]) != "null" {
			if err := json.Unmarshal(fields["graphPeriod"], &resource.GraphPeriod); err != nil {
				errs = append(errs, cog.MakeBuildErrors("graphPeriod", err)...)
			}

		}
		delete(fields, "graphPeriod")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TimeSeriesQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `TimeSeriesQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `TimeSeriesQuery` fields for violations and returns them.
func (resource TimeSeriesQuery) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SLOQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *SLOQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "projectName"
	if fields["projectName"] != nil {
		if string(fields["projectName"]) != "null" {
			if err := json.Unmarshal(fields["projectName"], &resource.ProjectName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("projectName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is null"))...)

		}
		delete(fields, "projectName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is missing from input"))...)
	}
	// Field "perSeriesAligner"
	if fields["perSeriesAligner"] != nil {
		if string(fields["perSeriesAligner"]) != "null" {
			if err := json.Unmarshal(fields["perSeriesAligner"], &resource.PerSeriesAligner); err != nil {
				errs = append(errs, cog.MakeBuildErrors("perSeriesAligner", err)...)
			}

		}
		delete(fields, "perSeriesAligner")

	}
	// Field "alignmentPeriod"
	if fields["alignmentPeriod"] != nil {
		if string(fields["alignmentPeriod"]) != "null" {
			if err := json.Unmarshal(fields["alignmentPeriod"], &resource.AlignmentPeriod); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alignmentPeriod", err)...)
			}

		}
		delete(fields, "alignmentPeriod")

	}
	// Field "selectorName"
	if fields["selectorName"] != nil {
		if string(fields["selectorName"]) != "null" {
			if err := json.Unmarshal(fields["selectorName"], &resource.SelectorName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("selectorName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("selectorName", errors.New("required field is null"))...)

		}
		delete(fields, "selectorName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("selectorName", errors.New("required field is missing from input"))...)
	}
	// Field "serviceId"
	if fields["serviceId"] != nil {
		if string(fields["serviceId"]) != "null" {
			if err := json.Unmarshal(fields["serviceId"], &resource.ServiceId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("serviceId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("serviceId", errors.New("required field is null"))...)

		}
		delete(fields, "serviceId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("serviceId", errors.New("required field is missing from input"))...)
	}
	// Field "serviceName"
	if fields["serviceName"] != nil {
		if string(fields["serviceName"]) != "null" {
			if err := json.Unmarshal(fields["serviceName"], &resource.ServiceName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("serviceName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("serviceName", errors.New("required field is null"))...)

		}
		delete(fields, "serviceName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("serviceName", errors.New("required field is missing from input"))...)
	}
	// Field "sloId"
	if fields["sloId"] != nil {
		if string(fields["sloId"]) != "null" {
			if err := json.Unmarshal(fields["sloId"], &resource.SloId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sloId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("sloId", errors.New("required field is null"))...)

		}
		delete(fields, "sloId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("sloId", errors.New("required field is missing from input"))...)
	}
	// Field "sloName"
	if fields["sloName"] != nil {
		if string(fields["sloName"]) != "null" {
			if err := json.Unmarshal(fields["sloName"], &resource.SloName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("sloName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("sloName", errors.New("required field is null"))...)

		}
		delete(fields, "sloName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("sloName", errors.New("required field is missing from input"))...)
	}
	// Field "goal"
	if fields["goal"] != nil {
		if string(fields["goal"]) != "null" {
			if err := json.Unmarshal(fields["goal"], &resource.Goal); err != nil {
				errs = append(errs, cog.MakeBuildErrors("goal", err)...)
			}

		}
		delete(fields, "goal")

	}
	// Field "lookbackPeriod"
	if fields["lookbackPeriod"] != nil {
		if string(fields["lookbackPeriod"]) != "null" {
			if err := json.Unmarshal(fields["lookbackPeriod"], &resource.LookbackPeriod); err != nil {
				errs = append(errs, cog.MakeBuildErrors("lookbackPeriod", err)...)
			}

		}
		delete(fields, "lookbackPeriod")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("SLOQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `SLOQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `SLOQuery` fields for violations and returns them.
func (resource SLOQuery) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PromQLQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *PromQLQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "projectName"
	if fields["projectName"] != nil {
		if string(fields["projectName"]) != "null" {
			if err := json.Unmarshal(fields["projectName"], &resource.ProjectName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("projectName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is null"))...)

		}
		delete(fields, "projectName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is missing from input"))...)
	}
	// Field "expr"
	if fields["expr"] != nil {
		if string(fields["expr"]) != "null" {
			if err := json.Unmarshal(fields["expr"], &resource.Expr); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expr", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expr", errors.New("required field is null"))...)

		}
		delete(fields, "expr")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expr", errors.New("required field is missing from input"))...)
	}
	// Field "step"
	if fields["step"] != nil {
		if string(fields["step"]) != "null" {
			if err := json.Unmarshal(fields["step"], &resource.Step); err != nil {
				errs = append(errs, cog.MakeBuildErrors("step", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("step", errors.New("required field is null"))...)

		}
		delete(fields, "step")
	} else {
		errs = append(errs, cog.MakeBuildErrors("step", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("PromQLQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `PromQLQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `PromQLQuery` fields for violations and returns them.
func (resource PromQLQuery) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *MetricQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "projectName"
	if fields["projectName"] != nil {
		if string(fields["projectName"]) != "null" {
			if err := json.Unmarshal(fields["projectName"], &resource.ProjectName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("projectName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is null"))...)

		}
		delete(fields, "projectName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is missing from input"))...)
	}
	// Field "perSeriesAligner"
	if fields["perSeriesAligner"] != nil {
		if string(fields["perSeriesAligner"]) != "null" {
			if err := json.Unmarshal(fields["perSeriesAligner"], &resource.PerSeriesAligner); err != nil {
				errs = append(errs, cog.MakeBuildErrors("perSeriesAligner", err)...)
			}

		}
		delete(fields, "perSeriesAligner")

	}
	// Field "alignmentPeriod"
	if fields["alignmentPeriod"] != nil {
		if string(fields["alignmentPeriod"]) != "null" {
			if err := json.Unmarshal(fields["alignmentPeriod"], &resource.AlignmentPeriod); err != nil {
				errs = append(errs, cog.MakeBuildErrors("alignmentPeriod", err)...)
			}

		}
		delete(fields, "alignmentPeriod")

	}
	// Field "aliasBy"
	if fields["aliasBy"] != nil {
		if string(fields["aliasBy"]) != "null" {
			if err := json.Unmarshal(fields["aliasBy"], &resource.AliasBy); err != nil {
				errs = append(errs, cog.MakeBuildErrors("aliasBy", err)...)
			}

		}
		delete(fields, "aliasBy")

	}
	// Field "editorMode"
	if fields["editorMode"] != nil {
		if string(fields["editorMode"]) != "null" {
			if err := json.Unmarshal(fields["editorMode"], &resource.EditorMode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("editorMode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("editorMode", errors.New("required field is null"))...)

		}
		delete(fields, "editorMode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("editorMode", errors.New("required field is missing from input"))...)
	}
	// Field "metricType"
	if fields["metricType"] != nil {
		if string(fields["metricType"]) != "null" {
			if err := json.Unmarshal(fields["metricType"], &resource.MetricType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricType", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metricType", errors.New("required field is null"))...)

		}
		delete(fields, "metricType")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metricType", errors.New("required field is missing from input"))...)
	}
	// Field "crossSeriesReducer"
	if fields["crossSeriesReducer"] != nil {
		if string(fields["crossSeriesReducer"]) != "null" {
			if err := json.Unmarshal(fields["crossSeriesReducer"], &resource.CrossSeriesReducer); err != nil {
				errs = append(errs, cog.MakeBuildErrors("crossSeriesReducer", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("crossSeriesReducer", errors.New("required field is null"))...)

		}
		delete(fields, "crossSeriesReducer")
	} else {
		errs = append(errs, cog.MakeBuildErrors("crossSeriesReducer", errors.New("required field is missing from input"))...)
	}
	// Field "groupBys"
	if fields["groupBys"] != nil {
		if string(fields["groupBys"]) != "null" {

			if err := json.Unmarshal(fields["groupBys"], &resource.GroupBys); err != nil {
				errs = append(errs, cog.MakeBuildErrors("groupBys", err)...)
			}

		}
		delete(fields, "groupBys")

	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			if err := json.Unmarshal(fields["filters"], &resource.Filters); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filters", err)...)
			}

		}
		delete(fields, "filters")

	}
	// Field "metricKind"
	if fields["metricKind"] != nil {
		if string(fields["metricKind"]) != "null" {
			if err := json.Unmarshal(fields["metricKind"], &resource.MetricKind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricKind", err)...)
			}

		}
		delete(fields, "metricKind")

	}
	// Field "valueType"
	if fields["valueType"] != nil {
		if string(fields["valueType"]) != "null" {
			if err := json.Unmarshal(fields["valueType"], &resource.ValueType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valueType", err)...)
			}

		}
		delete(fields, "valueType")

	}
	// Field "view"
	if fields["view"] != nil {
		if string(fields["view"]) != "null" {
			if err := json.Unmarshal(fields["view"], &resource.View); err != nil {
				errs = append(errs, cog.MakeBuildErrors("view", err)...)
			}

		}
		delete(fields, "view")

	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {
			if err := json.Unmarshal(fields["query"], &resource.Query); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")
	} else {
		errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is missing from input"))...)
	}
	// Field "preprocessor"
	if fields["preprocessor"] != nil {
		if string(fields["preprocessor"]) != "null" {
			if err := json.Unmarshal(fields["preprocessor"], &resource.Preprocessor); err != nil {
				errs = append(errs, cog.MakeBuildErrors("preprocessor", err)...)
			}

		}
		delete(fields, "preprocessor")

	}
	// Field "graphPeriod"
	if fields["graphPeriod"] != nil {
		if string(fields["graphPeriod"]) != "null" {
			if err := json.Unmarshal(fields["graphPeriod"], &resource.GraphPeriod); err != nil {
				errs = append(errs, cog.MakeBuildErrors("graphPeriod", err)...)
			}

		}
		delete(fields, "graphPeriod")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("MetricQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `MetricQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `MetricQuery` fields for violations and returns them.
func (resource MetricQuery) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `LegacyCloudMonitoringAnnotationQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *LegacyCloudMonitoringAnnotationQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "projectName"
	if fields["projectName"] != nil {
		if string(fields["projectName"]) != "null" {
			if err := json.Unmarshal(fields["projectName"], &resource.ProjectName); err != nil {
				errs = append(errs, cog.MakeBuildErrors("projectName", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is null"))...)

		}
		delete(fields, "projectName")
	} else {
		errs = append(errs, cog.MakeBuildErrors("projectName", errors.New("required field is missing from input"))...)
	}
	// Field "metricType"
	if fields["metricType"] != nil {
		if string(fields["metricType"]) != "null" {
			if err := json.Unmarshal(fields["metricType"], &resource.MetricType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricType", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metricType", errors.New("required field is null"))...)

		}
		delete(fields, "metricType")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metricType", errors.New("required field is missing from input"))...)
	}
	// Field "refId"
	if fields["refId"] != nil {
		if string(fields["refId"]) != "null" {
			if err := json.Unmarshal(fields["refId"], &resource.RefId); err != nil {
				errs = append(errs, cog.MakeBuildErrors("refId", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is null"))...)

		}
		delete(fields, "refId")
	} else {
		errs = append(errs, cog.MakeBuildErrors("refId", errors.New("required field is missing from input"))...)
	}
	// Field "filters"
	if fields["filters"] != nil {
		if string(fields["filters"]) != "null" {

			if err := json.Unmarshal(fields["filters"], &resource.Filters); err != nil {
				errs = append(errs, cog.MakeBuildErrors("filters", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("filters", errors.New("required field is null"))...)

		}
		delete(fields, "filters")
	} else {
		errs = append(errs, cog.MakeBuildErrors("filters", errors.New("required field is missing from input"))...)
	}
	// Field "metricKind"
	if fields["metricKind"] != nil {
		if string(fields["metricKind"]) != "null" {
			if err := json.Unmarshal(fields["metricKind"], &resource.MetricKind); err != nil {
				errs = append(errs, cog.MakeBuildErrors("metricKind", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("metricKind", errors.New("required field is null"))...)

		}
		delete(fields, "metricKind")
	} else {
		errs = append(errs, cog.MakeBuildErrors("metricKind", errors.New("required field is missing from input"))...)
	}
	// Field "valueType"
	if fields["valueType"] != nil {
		if string(fields["valueType"]) != "null" {
			if err := json.Unmarshal(fields["valueType"], &resource.ValueType); err != nil {
				errs = append(errs, cog.MakeBuildErrors("valueType", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("valueType", errors.New("required field is null"))...)

		}
		delete(fields, "valueType")
	} else {
		errs = append(errs, cog.MakeBuildErrors("valueType", errors.New("required field is missing from input"))...)
	}
	// Field "title"
	if fields["title"] != nil {
		if string(fields["title"]) != "null" {
			if err := json.Unmarshal(fields["title"], &resource.Title); err != nil {
				errs = append(errs, cog.MakeBuildErrors("title", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is null"))...)

		}
		delete(fields, "title")
	} else {
		errs = append(errs, cog.MakeBuildErrors("title", errors.New("required field is missing from input"))...)
	}
	// Field "text"
	if fields["text"] != nil {
		if string(fields["text"]) != "null" {
			if err := json.Unmarshal(fields["text"], &resource.Text); err != nil {
				errs = append(errs, cog.MakeBuildErrors("text", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("text", errors.New("required field is null"))...)

		}
		delete(fields, "text")
	} else {
		errs = append(errs, cog.MakeBuildErrors("text", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("LegacyCloudMonitoringAnnotationQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `LegacyCloudMonitoringAnnotationQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `LegacyCloudMonitoringAnnotationQuery` fields for violations and returns them.
func (resource LegacyCloudMonitoringAnnotationQuery) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Filter` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *Filter) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "key"
	if fields["key"] != nil {
		if string(fields["key"]) != "null" {
			if err := json.Unmarshal(fields["key"], &resource.Key); err != nil {
				errs = append(errs, cog.MakeBuildErrors("key", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is null"))...)

		}
		delete(fields, "key")
	} else {
		errs = append(errs, cog.MakeBuildErrors("key", errors.New("required field is missing from input"))...)
	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {
			if err := json.Unmarshal(fields["operator"], &resource.Operator); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is null"))...)

		}
		delete(fields, "operator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is missing from input"))...)
	}
	// Field "value"
	if fields["value"] != nil {
		if string(fields["value"]) != "null" {
			if err := json.Unmarshal(fields["value"], &resource.Value); err != nil {
				errs = append(errs, cog.MakeBuildErrors("value", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is null"))...)

		}
		delete(fields, "value")
	} else {
		errs = append(errs, cog.MakeBuildErrors("value", errors.New("required field is missing from input"))...)
	}
	// Field "condition"
	if fields["condition"] != nil {
		if string(fields["condition"]) != "null" {
			if err := json.Unmarshal(fields["condition"], &resource.Condition); err != nil {
				errs = append(errs, cog.MakeBuildErrors("condition", err)...)
			}

		}
		delete(fields, "condition")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("Filter", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `Filter` objects.
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

// Validate checks all the validation constraints that may be defined on `Filter` fields for violations and returns them.
func (resource Filter) Validate() error {
	return nil
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

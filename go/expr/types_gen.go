// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"

	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Expr = TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql

func VariantConfig() variants.DataqueryConfig {
	return variants.DataqueryConfig{
		Identifier: "__expr__",
		DataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Expr{}

			if err := json.Unmarshal(raw, dataquery); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
	}
}

type TypeMath struct {
	// The datasource
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// General math expression
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *ExprTypeMathResultAssertions `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *ExprTypeMathTimeRange `json:"timeRange,omitempty"`
	Type      string                 `json:"type"`
}

func (resource TypeMath) ImplementsDataqueryVariant() {}

func (resource TypeMath) DataqueryType() string {
	return "__expr__"
}

func (resource TypeMath) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(TypeMath)
	if !ok {
		return false
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	if resource.Expression != other.Expression {
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
	if resource.IntervalMs == nil && other.IntervalMs != nil || resource.IntervalMs != nil && other.IntervalMs == nil {
		return false
	}

	if resource.IntervalMs != nil {
		if *resource.IntervalMs != *other.IntervalMs {
			return false
		}
	}
	if resource.MaxDataPoints == nil && other.MaxDataPoints != nil || resource.MaxDataPoints != nil && other.MaxDataPoints == nil {
		return false
	}

	if resource.MaxDataPoints != nil {
		if *resource.MaxDataPoints != *other.MaxDataPoints {
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
	if resource.RefId != other.RefId {
		return false
	}
	if resource.ResultAssertions == nil && other.ResultAssertions != nil || resource.ResultAssertions != nil && other.ResultAssertions == nil {
		return false
	}

	if resource.ResultAssertions != nil {
		if !resource.ResultAssertions.Equals(*other.ResultAssertions) {
			return false
		}
	}
	if resource.TimeRange == nil && other.TimeRange != nil || resource.TimeRange != nil && other.TimeRange == nil {
		return false
	}

	if resource.TimeRange != nil {
		if !resource.TimeRange.Equals(*other.TimeRange) {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

type TypeReduce struct {
	// The datasource
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// Reference to single query result
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// The reducer
	// Possible enum values:
	//  - `"sum"`
	//  - `"mean"`
	//  - `"min"`
	//  - `"max"`
	//  - `"count"`
	//  - `"last"`
	//  - `"median"`
	Reducer TypeReduceReducer `json:"reducer"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *ExprTypeReduceResultAssertions `json:"resultAssertions,omitempty"`
	// Reducer Options
	Settings *ExprTypeReduceSettings `json:"settings,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *ExprTypeReduceTimeRange `json:"timeRange,omitempty"`
	Type      string                   `json:"type"`
}

func (resource TypeReduce) ImplementsDataqueryVariant() {}

func (resource TypeReduce) DataqueryType() string {
	return "__expr__"
}

func (resource TypeReduce) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(TypeReduce)
	if !ok {
		return false
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	if resource.Expression != other.Expression {
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
	if resource.IntervalMs == nil && other.IntervalMs != nil || resource.IntervalMs != nil && other.IntervalMs == nil {
		return false
	}

	if resource.IntervalMs != nil {
		if *resource.IntervalMs != *other.IntervalMs {
			return false
		}
	}
	if resource.MaxDataPoints == nil && other.MaxDataPoints != nil || resource.MaxDataPoints != nil && other.MaxDataPoints == nil {
		return false
	}

	if resource.MaxDataPoints != nil {
		if *resource.MaxDataPoints != *other.MaxDataPoints {
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
	if resource.Reducer != other.Reducer {
		return false
	}
	if resource.RefId != other.RefId {
		return false
	}
	if resource.ResultAssertions == nil && other.ResultAssertions != nil || resource.ResultAssertions != nil && other.ResultAssertions == nil {
		return false
	}

	if resource.ResultAssertions != nil {
		if !resource.ResultAssertions.Equals(*other.ResultAssertions) {
			return false
		}
	}
	if resource.Settings == nil && other.Settings != nil || resource.Settings != nil && other.Settings == nil {
		return false
	}

	if resource.Settings != nil {
		if !resource.Settings.Equals(*other.Settings) {
			return false
		}
	}
	if resource.TimeRange == nil && other.TimeRange != nil || resource.TimeRange != nil && other.TimeRange == nil {
		return false
	}

	if resource.TimeRange != nil {
		if !resource.TimeRange.Equals(*other.TimeRange) {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

type TypeResample struct {
	// The datasource
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// The downsample function
	// Possible enum values:
	//  - `"sum"`
	//  - `"mean"`
	//  - `"min"`
	//  - `"max"`
	//  - `"count"`
	//  - `"last"`
	//  - `"median"`
	Downsampler TypeResampleDownsampler `json:"downsampler"`
	// The math expression
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *ExprTypeResampleResultAssertions `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *ExprTypeResampleTimeRange `json:"timeRange,omitempty"`
	Type      string                     `json:"type"`
	// The upsample function
	// Possible enum values:
	//  - `"pad"` Use the last seen value
	//  - `"backfilling"` backfill
	//  - `"fillna"` Do not fill values (nill)
	Upsampler TypeResampleUpsampler `json:"upsampler"`
	// The time duration
	Window string `json:"window"`
}

func (resource TypeResample) ImplementsDataqueryVariant() {}

func (resource TypeResample) DataqueryType() string {
	return "__expr__"
}

func (resource TypeResample) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(TypeResample)
	if !ok {
		return false
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	if resource.Downsampler != other.Downsampler {
		return false
	}
	if resource.Expression != other.Expression {
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
	if resource.IntervalMs == nil && other.IntervalMs != nil || resource.IntervalMs != nil && other.IntervalMs == nil {
		return false
	}

	if resource.IntervalMs != nil {
		if *resource.IntervalMs != *other.IntervalMs {
			return false
		}
	}
	if resource.MaxDataPoints == nil && other.MaxDataPoints != nil || resource.MaxDataPoints != nil && other.MaxDataPoints == nil {
		return false
	}

	if resource.MaxDataPoints != nil {
		if *resource.MaxDataPoints != *other.MaxDataPoints {
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
	if resource.RefId != other.RefId {
		return false
	}
	if resource.ResultAssertions == nil && other.ResultAssertions != nil || resource.ResultAssertions != nil && other.ResultAssertions == nil {
		return false
	}

	if resource.ResultAssertions != nil {
		if !resource.ResultAssertions.Equals(*other.ResultAssertions) {
			return false
		}
	}
	if resource.TimeRange == nil && other.TimeRange != nil || resource.TimeRange != nil && other.TimeRange == nil {
		return false
	}

	if resource.TimeRange != nil {
		if !resource.TimeRange.Equals(*other.TimeRange) {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}
	if resource.Upsampler != other.Upsampler {
		return false
	}
	if resource.Window != other.Window {
		return false
	}

	return true
}

type TypeClassicConditions struct {
	Conditions []ExprTypeClassicConditionsConditions `json:"conditions"`
	// The datasource
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *ExprTypeClassicConditionsResultAssertions `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *ExprTypeClassicConditionsTimeRange `json:"timeRange,omitempty"`
	Type      string                              `json:"type"`
}

func (resource TypeClassicConditions) ImplementsDataqueryVariant() {}

func (resource TypeClassicConditions) DataqueryType() string {
	return "__expr__"
}

func (resource TypeClassicConditions) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(TypeClassicConditions)
	if !ok {
		return false
	}

	if len(resource.Conditions) != len(other.Conditions) {
		return false
	}

	for i1 := range resource.Conditions {
		if !resource.Conditions[i1].Equals(other.Conditions[i1]) {
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
	if resource.Hide == nil && other.Hide != nil || resource.Hide != nil && other.Hide == nil {
		return false
	}

	if resource.Hide != nil {
		if *resource.Hide != *other.Hide {
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
	if resource.MaxDataPoints == nil && other.MaxDataPoints != nil || resource.MaxDataPoints != nil && other.MaxDataPoints == nil {
		return false
	}

	if resource.MaxDataPoints != nil {
		if *resource.MaxDataPoints != *other.MaxDataPoints {
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
	if resource.RefId != other.RefId {
		return false
	}
	if resource.ResultAssertions == nil && other.ResultAssertions != nil || resource.ResultAssertions != nil && other.ResultAssertions == nil {
		return false
	}

	if resource.ResultAssertions != nil {
		if !resource.ResultAssertions.Equals(*other.ResultAssertions) {
			return false
		}
	}
	if resource.TimeRange == nil && other.TimeRange != nil || resource.TimeRange != nil && other.TimeRange == nil {
		return false
	}

	if resource.TimeRange != nil {
		if !resource.TimeRange.Equals(*other.TimeRange) {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

type TypeThreshold struct {
	// Threshold Conditions
	Conditions []ExprTypeThresholdConditions `json:"conditions"`
	// The datasource
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	// Reference to single query result
	Expression string `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *ExprTypeThresholdResultAssertions `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *ExprTypeThresholdTimeRange `json:"timeRange,omitempty"`
	Type      string                      `json:"type"`
}

func (resource TypeThreshold) ImplementsDataqueryVariant() {}

func (resource TypeThreshold) DataqueryType() string {
	return "__expr__"
}

func (resource TypeThreshold) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(TypeThreshold)
	if !ok {
		return false
	}

	if len(resource.Conditions) != len(other.Conditions) {
		return false
	}

	for i1 := range resource.Conditions {
		if !resource.Conditions[i1].Equals(other.Conditions[i1]) {
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
	if resource.Expression != other.Expression {
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
	if resource.IntervalMs == nil && other.IntervalMs != nil || resource.IntervalMs != nil && other.IntervalMs == nil {
		return false
	}

	if resource.IntervalMs != nil {
		if *resource.IntervalMs != *other.IntervalMs {
			return false
		}
	}
	if resource.MaxDataPoints == nil && other.MaxDataPoints != nil || resource.MaxDataPoints != nil && other.MaxDataPoints == nil {
		return false
	}

	if resource.MaxDataPoints != nil {
		if *resource.MaxDataPoints != *other.MaxDataPoints {
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
	if resource.RefId != other.RefId {
		return false
	}
	if resource.ResultAssertions == nil && other.ResultAssertions != nil || resource.ResultAssertions != nil && other.ResultAssertions == nil {
		return false
	}

	if resource.ResultAssertions != nil {
		if !resource.ResultAssertions.Equals(*other.ResultAssertions) {
			return false
		}
	}
	if resource.TimeRange == nil && other.TimeRange != nil || resource.TimeRange != nil && other.TimeRange == nil {
		return false
	}

	if resource.TimeRange != nil {
		if !resource.TimeRange.Equals(*other.TimeRange) {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

type TypeSql struct {
	// The datasource
	Datasource *dashboard.DataSourceRef `json:"datasource,omitempty"`
	Expression string                   `json:"expression"`
	// true if query is disabled (ie should not be returned to the dashboard)
	// NOTE: this does not always imply that the query should not be executed since
	// the results from a hidden query may be used as the input to other queries (SSE etc)
	Hide *bool `json:"hide,omitempty"`
	// Interval is the suggested duration between time points in a time series query.
	// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
	// from the interval required to fill a pixels in the visualization
	IntervalMs *float64 `json:"intervalMs,omitempty"`
	// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
	// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
	// from the number of pixels visible in a visualization
	MaxDataPoints *int64 `json:"maxDataPoints,omitempty"`
	// QueryType is an optional identifier for the type of query.
	// It can be used to distinguish different types of queries.
	QueryType *string `json:"queryType,omitempty"`
	// RefID is the unique identifier of the query, set by the frontend call.
	RefId string `json:"refId"`
	// Optionally define expected query result behavior
	ResultAssertions *ExprTypeSqlResultAssertions `json:"resultAssertions,omitempty"`
	// TimeRange represents the query range
	// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
	// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
	TimeRange *ExprTypeSqlTimeRange `json:"timeRange,omitempty"`
	Type      string                `json:"type"`
}

func (resource TypeSql) ImplementsDataqueryVariant() {}

func (resource TypeSql) DataqueryType() string {
	return "__expr__"
}

func (resource TypeSql) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(TypeSql)
	if !ok {
		return false
	}
	if resource.Datasource == nil && other.Datasource != nil || resource.Datasource != nil && other.Datasource == nil {
		return false
	}

	if resource.Datasource != nil {
		if !resource.Datasource.Equals(*other.Datasource) {
			return false
		}
	}
	if resource.Expression != other.Expression {
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
	if resource.IntervalMs == nil && other.IntervalMs != nil || resource.IntervalMs != nil && other.IntervalMs == nil {
		return false
	}

	if resource.IntervalMs != nil {
		if *resource.IntervalMs != *other.IntervalMs {
			return false
		}
	}
	if resource.MaxDataPoints == nil && other.MaxDataPoints != nil || resource.MaxDataPoints != nil && other.MaxDataPoints == nil {
		return false
	}

	if resource.MaxDataPoints != nil {
		if *resource.MaxDataPoints != *other.MaxDataPoints {
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
	if resource.RefId != other.RefId {
		return false
	}
	if resource.ResultAssertions == nil && other.ResultAssertions != nil || resource.ResultAssertions != nil && other.ResultAssertions == nil {
		return false
	}

	if resource.ResultAssertions != nil {
		if !resource.ResultAssertions.Equals(*other.ResultAssertions) {
			return false
		}
	}
	if resource.TimeRange == nil && other.TimeRange != nil || resource.TimeRange != nil && other.TimeRange == nil {
		return false
	}

	if resource.TimeRange != nil {
		if !resource.TimeRange.Equals(*other.TimeRange) {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

type TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql struct {
	TypeMath              *TypeMath              `json:"TypeMath,omitempty"`
	TypeReduce            *TypeReduce            `json:"TypeReduce,omitempty"`
	TypeResample          *TypeResample          `json:"TypeResample,omitempty"`
	TypeClassicConditions *TypeClassicConditions `json:"TypeClassicConditions,omitempty"`
	TypeThreshold         *TypeThreshold         `json:"TypeThreshold,omitempty"`
	TypeSql               *TypeSql               `json:"TypeSql,omitempty"`
}

func (resource TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) ImplementsDataqueryVariant() {
}

func (resource TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) DataqueryType() string {
	return "__expr__"
}

func (resource TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) MarshalJSON() ([]byte, error) {
	if resource.TypeMath != nil {
		return json.Marshal(resource.TypeMath)
	}
	if resource.TypeReduce != nil {
		return json.Marshal(resource.TypeReduce)
	}
	if resource.TypeResample != nil {
		return json.Marshal(resource.TypeResample)
	}
	if resource.TypeClassicConditions != nil {
		return json.Marshal(resource.TypeClassicConditions)
	}
	if resource.TypeThreshold != nil {
		return json.Marshal(resource.TypeThreshold)
	}
	if resource.TypeSql != nil {
		return json.Marshal(resource.TypeSql)
	}

	return nil, fmt.Errorf("no value for disjunction of refs")
}

func (resource *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) UnmarshalJSON(raw []byte) error {
	if raw == nil {
		return nil
	}

	// FIXME: this is wasteful, we need to find a more efficient way to unmarshal this.
	parsedAsMap := make(map[string]any)
	if err := json.Unmarshal(raw, &parsedAsMap); err != nil {
		return err
	}

	discriminator, found := parsedAsMap["type"]
	if !found {
		return errors.New("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "classic_conditions":
		var typeClassicConditions TypeClassicConditions
		if err := json.Unmarshal(raw, &typeClassicConditions); err != nil {
			return err
		}

		resource.TypeClassicConditions = &typeClassicConditions
		return nil
	case "math":
		var typeMath TypeMath
		if err := json.Unmarshal(raw, &typeMath); err != nil {
			return err
		}

		resource.TypeMath = &typeMath
		return nil
	case "reduce":
		var typeReduce TypeReduce
		if err := json.Unmarshal(raw, &typeReduce); err != nil {
			return err
		}

		resource.TypeReduce = &typeReduce
		return nil
	case "resample":
		var typeResample TypeResample
		if err := json.Unmarshal(raw, &typeResample); err != nil {
			return err
		}

		resource.TypeResample = &typeResample
		return nil
	case "sql":
		var typeSql TypeSql
		if err := json.Unmarshal(raw, &typeSql); err != nil {
			return err
		}

		resource.TypeSql = &typeSql
		return nil
	case "threshold":
		var typeThreshold TypeThreshold
		if err := json.Unmarshal(raw, &typeThreshold); err != nil {
			return err
		}

		resource.TypeThreshold = &typeThreshold
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

func (resource TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) Equals(otherCandidate variants.Dataquery) bool {
	if otherCandidate == nil {
		return false
	}

	other, ok := otherCandidate.(TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql)
	if !ok {
		return false
	}
	if resource.TypeMath == nil && other.TypeMath != nil || resource.TypeMath != nil && other.TypeMath == nil {
		return false
	}

	if resource.TypeMath != nil {
		if !resource.TypeMath.Equals(*other.TypeMath) {
			return false
		}
	}
	if resource.TypeReduce == nil && other.TypeReduce != nil || resource.TypeReduce != nil && other.TypeReduce == nil {
		return false
	}

	if resource.TypeReduce != nil {
		if !resource.TypeReduce.Equals(*other.TypeReduce) {
			return false
		}
	}
	if resource.TypeResample == nil && other.TypeResample != nil || resource.TypeResample != nil && other.TypeResample == nil {
		return false
	}

	if resource.TypeResample != nil {
		if !resource.TypeResample.Equals(*other.TypeResample) {
			return false
		}
	}
	if resource.TypeClassicConditions == nil && other.TypeClassicConditions != nil || resource.TypeClassicConditions != nil && other.TypeClassicConditions == nil {
		return false
	}

	if resource.TypeClassicConditions != nil {
		if !resource.TypeClassicConditions.Equals(*other.TypeClassicConditions) {
			return false
		}
	}
	if resource.TypeThreshold == nil && other.TypeThreshold != nil || resource.TypeThreshold != nil && other.TypeThreshold == nil {
		return false
	}

	if resource.TypeThreshold != nil {
		if !resource.TypeThreshold.Equals(*other.TypeThreshold) {
			return false
		}
	}
	if resource.TypeSql == nil && other.TypeSql != nil || resource.TypeSql != nil && other.TypeSql == nil {
		return false
	}

	if resource.TypeSql != nil {
		if !resource.TypeSql.Equals(*other.TypeSql) {
			return false
		}
	}

	return true
}

type TypeMathType string

const (
	TypeMathTypeNone             TypeMathType = ""
	TypeMathTypeTimeseriesWide   TypeMathType = "timeseries-wide"
	TypeMathTypeTimeseriesLong   TypeMathType = "timeseries-long"
	TypeMathTypeTimeseriesMany   TypeMathType = "timeseries-many"
	TypeMathTypeTimeseriesMulti  TypeMathType = "timeseries-multi"
	TypeMathTypeDirectoryListing TypeMathType = "directory-listing"
	TypeMathTypeTable            TypeMathType = "table"
	TypeMathTypeNumericWide      TypeMathType = "numeric-wide"
	TypeMathTypeNumericMulti     TypeMathType = "numeric-multi"
	TypeMathTypeNumericLong      TypeMathType = "numeric-long"
	TypeMathTypeLogLines         TypeMathType = "log-lines"
)

type TypeReduceReducer string

const (
	TypeReduceReducerSum    TypeReduceReducer = "sum"
	TypeReduceReducerMean   TypeReduceReducer = "mean"
	TypeReduceReducerMin    TypeReduceReducer = "min"
	TypeReduceReducerMax    TypeReduceReducer = "max"
	TypeReduceReducerCount  TypeReduceReducer = "count"
	TypeReduceReducerLast   TypeReduceReducer = "last"
	TypeReduceReducerMedian TypeReduceReducer = "median"
)

type TypeReduceType string

const (
	TypeReduceTypeNone             TypeReduceType = ""
	TypeReduceTypeTimeseriesWide   TypeReduceType = "timeseries-wide"
	TypeReduceTypeTimeseriesLong   TypeReduceType = "timeseries-long"
	TypeReduceTypeTimeseriesMany   TypeReduceType = "timeseries-many"
	TypeReduceTypeTimeseriesMulti  TypeReduceType = "timeseries-multi"
	TypeReduceTypeDirectoryListing TypeReduceType = "directory-listing"
	TypeReduceTypeTable            TypeReduceType = "table"
	TypeReduceTypeNumericWide      TypeReduceType = "numeric-wide"
	TypeReduceTypeNumericMulti     TypeReduceType = "numeric-multi"
	TypeReduceTypeNumericLong      TypeReduceType = "numeric-long"
	TypeReduceTypeLogLines         TypeReduceType = "log-lines"
)

type TypeReduceMode string

const (
	TypeReduceModeDropNN    TypeReduceMode = "dropNN"
	TypeReduceModeReplaceNN TypeReduceMode = "replaceNN"
)

type TypeResampleDownsampler string

const (
	TypeResampleDownsamplerSum    TypeResampleDownsampler = "sum"
	TypeResampleDownsamplerMean   TypeResampleDownsampler = "mean"
	TypeResampleDownsamplerMin    TypeResampleDownsampler = "min"
	TypeResampleDownsamplerMax    TypeResampleDownsampler = "max"
	TypeResampleDownsamplerCount  TypeResampleDownsampler = "count"
	TypeResampleDownsamplerLast   TypeResampleDownsampler = "last"
	TypeResampleDownsamplerMedian TypeResampleDownsampler = "median"
)

type TypeResampleType string

const (
	TypeResampleTypeNone             TypeResampleType = ""
	TypeResampleTypeTimeseriesWide   TypeResampleType = "timeseries-wide"
	TypeResampleTypeTimeseriesLong   TypeResampleType = "timeseries-long"
	TypeResampleTypeTimeseriesMany   TypeResampleType = "timeseries-many"
	TypeResampleTypeTimeseriesMulti  TypeResampleType = "timeseries-multi"
	TypeResampleTypeDirectoryListing TypeResampleType = "directory-listing"
	TypeResampleTypeTable            TypeResampleType = "table"
	TypeResampleTypeNumericWide      TypeResampleType = "numeric-wide"
	TypeResampleTypeNumericMulti     TypeResampleType = "numeric-multi"
	TypeResampleTypeNumericLong      TypeResampleType = "numeric-long"
	TypeResampleTypeLogLines         TypeResampleType = "log-lines"
)

type TypeResampleUpsampler string

const (
	TypeResampleUpsamplerPad         TypeResampleUpsampler = "pad"
	TypeResampleUpsamplerBackfilling TypeResampleUpsampler = "backfilling"
	TypeResampleUpsamplerFillna      TypeResampleUpsampler = "fillna"
)

type TypeClassicConditionsType string

const (
	TypeClassicConditionsTypeNone             TypeClassicConditionsType = ""
	TypeClassicConditionsTypeTimeseriesWide   TypeClassicConditionsType = "timeseries-wide"
	TypeClassicConditionsTypeTimeseriesLong   TypeClassicConditionsType = "timeseries-long"
	TypeClassicConditionsTypeTimeseriesMany   TypeClassicConditionsType = "timeseries-many"
	TypeClassicConditionsTypeTimeseriesMulti  TypeClassicConditionsType = "timeseries-multi"
	TypeClassicConditionsTypeDirectoryListing TypeClassicConditionsType = "directory-listing"
	TypeClassicConditionsTypeTable            TypeClassicConditionsType = "table"
	TypeClassicConditionsTypeNumericWide      TypeClassicConditionsType = "numeric-wide"
	TypeClassicConditionsTypeNumericMulti     TypeClassicConditionsType = "numeric-multi"
	TypeClassicConditionsTypeNumericLong      TypeClassicConditionsType = "numeric-long"
	TypeClassicConditionsTypeLogLines         TypeClassicConditionsType = "log-lines"
)

type TypeThresholdType string

const (
	TypeThresholdTypeNone             TypeThresholdType = ""
	TypeThresholdTypeTimeseriesWide   TypeThresholdType = "timeseries-wide"
	TypeThresholdTypeTimeseriesLong   TypeThresholdType = "timeseries-long"
	TypeThresholdTypeTimeseriesMany   TypeThresholdType = "timeseries-many"
	TypeThresholdTypeTimeseriesMulti  TypeThresholdType = "timeseries-multi"
	TypeThresholdTypeDirectoryListing TypeThresholdType = "directory-listing"
	TypeThresholdTypeTable            TypeThresholdType = "table"
	TypeThresholdTypeNumericWide      TypeThresholdType = "numeric-wide"
	TypeThresholdTypeNumericMulti     TypeThresholdType = "numeric-multi"
	TypeThresholdTypeNumericLong      TypeThresholdType = "numeric-long"
	TypeThresholdTypeLogLines         TypeThresholdType = "log-lines"
)

type TypeSqlType string

const (
	TypeSqlTypeNone             TypeSqlType = ""
	TypeSqlTypeTimeseriesWide   TypeSqlType = "timeseries-wide"
	TypeSqlTypeTimeseriesLong   TypeSqlType = "timeseries-long"
	TypeSqlTypeTimeseriesMany   TypeSqlType = "timeseries-many"
	TypeSqlTypeTimeseriesMulti  TypeSqlType = "timeseries-multi"
	TypeSqlTypeDirectoryListing TypeSqlType = "directory-listing"
	TypeSqlTypeTable            TypeSqlType = "table"
	TypeSqlTypeNumericWide      TypeSqlType = "numeric-wide"
	TypeSqlTypeNumericMulti     TypeSqlType = "numeric-multi"
	TypeSqlTypeNumericLong      TypeSqlType = "numeric-long"
	TypeSqlTypeLogLines         TypeSqlType = "log-lines"
)

type ExprTypeMathResultAssertions struct {
	// Maximum frame count
	MaxFrames *int64 `json:"maxFrames,omitempty"`
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""`
	//  - `"timeseries-wide"`
	//  - `"timeseries-long"`
	//  - `"timeseries-many"`
	//  - `"timeseries-multi"`
	//  - `"directory-listing"`
	//  - `"table"`
	//  - `"numeric-wide"`
	//  - `"numeric-multi"`
	//  - `"numeric-long"`
	//  - `"log-lines"`
	Type *TypeMathType `json:"type,omitempty"`
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	TypeVersion []int64 `json:"typeVersion"`
}

func (resource ExprTypeMathResultAssertions) Equals(other ExprTypeMathResultAssertions) bool {
	if resource.MaxFrames == nil && other.MaxFrames != nil || resource.MaxFrames != nil && other.MaxFrames == nil {
		return false
	}

	if resource.MaxFrames != nil {
		if *resource.MaxFrames != *other.MaxFrames {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}

	if len(resource.TypeVersion) != len(other.TypeVersion) {
		return false
	}

	for i1 := range resource.TypeVersion {
		if resource.TypeVersion[i1] != other.TypeVersion[i1] {
			return false
		}
	}

	return true
}

type ExprTypeMathTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

func (resource ExprTypeMathTimeRange) Equals(other ExprTypeMathTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

type ExprTypeReduceResultAssertions struct {
	// Maximum frame count
	MaxFrames *int64 `json:"maxFrames,omitempty"`
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""`
	//  - `"timeseries-wide"`
	//  - `"timeseries-long"`
	//  - `"timeseries-many"`
	//  - `"timeseries-multi"`
	//  - `"directory-listing"`
	//  - `"table"`
	//  - `"numeric-wide"`
	//  - `"numeric-multi"`
	//  - `"numeric-long"`
	//  - `"log-lines"`
	Type *TypeReduceType `json:"type,omitempty"`
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	TypeVersion []int64 `json:"typeVersion"`
}

func (resource ExprTypeReduceResultAssertions) Equals(other ExprTypeReduceResultAssertions) bool {
	if resource.MaxFrames == nil && other.MaxFrames != nil || resource.MaxFrames != nil && other.MaxFrames == nil {
		return false
	}

	if resource.MaxFrames != nil {
		if *resource.MaxFrames != *other.MaxFrames {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}

	if len(resource.TypeVersion) != len(other.TypeVersion) {
		return false
	}

	for i1 := range resource.TypeVersion {
		if resource.TypeVersion[i1] != other.TypeVersion[i1] {
			return false
		}
	}

	return true
}

type ExprTypeReduceSettings struct {
	// Non-number reduce behavior
	// Possible enum values:
	//  - `"dropNN"` Drop non-numbers
	//  - `"replaceNN"` Replace non-numbers
	Mode TypeReduceMode `json:"mode"`
	// Only valid when mode is replace
	ReplaceWithValue *float64 `json:"replaceWithValue,omitempty"`
}

func (resource ExprTypeReduceSettings) Equals(other ExprTypeReduceSettings) bool {
	if resource.Mode != other.Mode {
		return false
	}
	if resource.ReplaceWithValue == nil && other.ReplaceWithValue != nil || resource.ReplaceWithValue != nil && other.ReplaceWithValue == nil {
		return false
	}

	if resource.ReplaceWithValue != nil {
		if *resource.ReplaceWithValue != *other.ReplaceWithValue {
			return false
		}
	}

	return true
}

type ExprTypeReduceTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

func (resource ExprTypeReduceTimeRange) Equals(other ExprTypeReduceTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

type ExprTypeResampleResultAssertions struct {
	// Maximum frame count
	MaxFrames *int64 `json:"maxFrames,omitempty"`
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""`
	//  - `"timeseries-wide"`
	//  - `"timeseries-long"`
	//  - `"timeseries-many"`
	//  - `"timeseries-multi"`
	//  - `"directory-listing"`
	//  - `"table"`
	//  - `"numeric-wide"`
	//  - `"numeric-multi"`
	//  - `"numeric-long"`
	//  - `"log-lines"`
	Type *TypeResampleType `json:"type,omitempty"`
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	TypeVersion []int64 `json:"typeVersion"`
}

func (resource ExprTypeResampleResultAssertions) Equals(other ExprTypeResampleResultAssertions) bool {
	if resource.MaxFrames == nil && other.MaxFrames != nil || resource.MaxFrames != nil && other.MaxFrames == nil {
		return false
	}

	if resource.MaxFrames != nil {
		if *resource.MaxFrames != *other.MaxFrames {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}

	if len(resource.TypeVersion) != len(other.TypeVersion) {
		return false
	}

	for i1 := range resource.TypeVersion {
		if resource.TypeVersion[i1] != other.TypeVersion[i1] {
			return false
		}
	}

	return true
}

type ExprTypeResampleTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

func (resource ExprTypeResampleTimeRange) Equals(other ExprTypeResampleTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

type ExprTypeClassicConditionsConditionsEvaluator struct {
	Params []float64 `json:"params"`
	// e.g. "gt"
	Type string `json:"type"`
}

func (resource ExprTypeClassicConditionsConditionsEvaluator) Equals(other ExprTypeClassicConditionsConditionsEvaluator) bool {

	if len(resource.Params) != len(other.Params) {
		return false
	}

	for i1 := range resource.Params {
		if resource.Params[i1] != other.Params[i1] {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

type ExprTypeClassicConditionsConditionsOperator struct {
	Type TypeClassicConditionsType `json:"type"`
}

func (resource ExprTypeClassicConditionsConditionsOperator) Equals(other ExprTypeClassicConditionsConditionsOperator) bool {
	if resource.Type != other.Type {
		return false
	}

	return true
}

type ExprTypeClassicConditionsConditionsQuery struct {
	Params []string `json:"params"`
}

func (resource ExprTypeClassicConditionsConditionsQuery) Equals(other ExprTypeClassicConditionsConditionsQuery) bool {

	if len(resource.Params) != len(other.Params) {
		return false
	}

	for i1 := range resource.Params {
		if resource.Params[i1] != other.Params[i1] {
			return false
		}
	}

	return true
}

type ExprTypeClassicConditionsConditionsReducer struct {
	Type string `json:"type"`
}

func (resource ExprTypeClassicConditionsConditionsReducer) Equals(other ExprTypeClassicConditionsConditionsReducer) bool {
	if resource.Type != other.Type {
		return false
	}

	return true
}

type ExprTypeClassicConditionsConditions struct {
	Evaluator ExprTypeClassicConditionsConditionsEvaluator `json:"evaluator"`
	Operator  ExprTypeClassicConditionsConditionsOperator  `json:"operator"`
	Query     ExprTypeClassicConditionsConditionsQuery     `json:"query"`
	Reducer   ExprTypeClassicConditionsConditionsReducer   `json:"reducer"`
}

func (resource ExprTypeClassicConditionsConditions) Equals(other ExprTypeClassicConditionsConditions) bool {
	if !resource.Evaluator.Equals(other.Evaluator) {
		return false
	}
	if !resource.Operator.Equals(other.Operator) {
		return false
	}
	if !resource.Query.Equals(other.Query) {
		return false
	}
	if !resource.Reducer.Equals(other.Reducer) {
		return false
	}

	return true
}

type ExprTypeClassicConditionsResultAssertions struct {
	// Maximum frame count
	MaxFrames *int64 `json:"maxFrames,omitempty"`
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""`
	//  - `"timeseries-wide"`
	//  - `"timeseries-long"`
	//  - `"timeseries-many"`
	//  - `"timeseries-multi"`
	//  - `"directory-listing"`
	//  - `"table"`
	//  - `"numeric-wide"`
	//  - `"numeric-multi"`
	//  - `"numeric-long"`
	//  - `"log-lines"`
	Type *TypeClassicConditionsType `json:"type,omitempty"`
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	TypeVersion []int64 `json:"typeVersion"`
}

func (resource ExprTypeClassicConditionsResultAssertions) Equals(other ExprTypeClassicConditionsResultAssertions) bool {
	if resource.MaxFrames == nil && other.MaxFrames != nil || resource.MaxFrames != nil && other.MaxFrames == nil {
		return false
	}

	if resource.MaxFrames != nil {
		if *resource.MaxFrames != *other.MaxFrames {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}

	if len(resource.TypeVersion) != len(other.TypeVersion) {
		return false
	}

	for i1 := range resource.TypeVersion {
		if resource.TypeVersion[i1] != other.TypeVersion[i1] {
			return false
		}
	}

	return true
}

type ExprTypeClassicConditionsTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

func (resource ExprTypeClassicConditionsTimeRange) Equals(other ExprTypeClassicConditionsTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

type ExprTypeThresholdConditionsEvaluator struct {
	Params []float64 `json:"params"`
	// e.g. "gt"
	Type TypeThresholdType `json:"type"`
}

func (resource ExprTypeThresholdConditionsEvaluator) Equals(other ExprTypeThresholdConditionsEvaluator) bool {

	if len(resource.Params) != len(other.Params) {
		return false
	}

	for i1 := range resource.Params {
		if resource.Params[i1] != other.Params[i1] {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

type ExprTypeThresholdConditionsUnloadEvaluator struct {
	Params []float64 `json:"params"`
	// e.g. "gt"
	Type TypeThresholdType `json:"type"`
}

func (resource ExprTypeThresholdConditionsUnloadEvaluator) Equals(other ExprTypeThresholdConditionsUnloadEvaluator) bool {

	if len(resource.Params) != len(other.Params) {
		return false
	}

	for i1 := range resource.Params {
		if resource.Params[i1] != other.Params[i1] {
			return false
		}
	}
	if resource.Type != other.Type {
		return false
	}

	return true
}

type ExprTypeThresholdConditions struct {
	Evaluator        ExprTypeThresholdConditionsEvaluator        `json:"evaluator"`
	LoadedDimensions any                                         `json:"loadedDimensions,omitempty"`
	UnloadEvaluator  *ExprTypeThresholdConditionsUnloadEvaluator `json:"unloadEvaluator,omitempty"`
}

func (resource ExprTypeThresholdConditions) Equals(other ExprTypeThresholdConditions) bool {
	if !resource.Evaluator.Equals(other.Evaluator) {
		return false
	}
	// is DeepEqual good enough here?
	if !reflect.DeepEqual(resource.LoadedDimensions, other.LoadedDimensions) {
		return false
	}
	if resource.UnloadEvaluator == nil && other.UnloadEvaluator != nil || resource.UnloadEvaluator != nil && other.UnloadEvaluator == nil {
		return false
	}

	if resource.UnloadEvaluator != nil {
		if !resource.UnloadEvaluator.Equals(*other.UnloadEvaluator) {
			return false
		}
	}

	return true
}

type ExprTypeThresholdResultAssertions struct {
	// Maximum frame count
	MaxFrames *int64 `json:"maxFrames,omitempty"`
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""`
	//  - `"timeseries-wide"`
	//  - `"timeseries-long"`
	//  - `"timeseries-many"`
	//  - `"timeseries-multi"`
	//  - `"directory-listing"`
	//  - `"table"`
	//  - `"numeric-wide"`
	//  - `"numeric-multi"`
	//  - `"numeric-long"`
	//  - `"log-lines"`
	Type *TypeThresholdType `json:"type,omitempty"`
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	TypeVersion []int64 `json:"typeVersion"`
}

func (resource ExprTypeThresholdResultAssertions) Equals(other ExprTypeThresholdResultAssertions) bool {
	if resource.MaxFrames == nil && other.MaxFrames != nil || resource.MaxFrames != nil && other.MaxFrames == nil {
		return false
	}

	if resource.MaxFrames != nil {
		if *resource.MaxFrames != *other.MaxFrames {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}

	if len(resource.TypeVersion) != len(other.TypeVersion) {
		return false
	}

	for i1 := range resource.TypeVersion {
		if resource.TypeVersion[i1] != other.TypeVersion[i1] {
			return false
		}
	}

	return true
}

type ExprTypeThresholdTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

func (resource ExprTypeThresholdTimeRange) Equals(other ExprTypeThresholdTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

type ExprTypeSqlResultAssertions struct {
	// Maximum frame count
	MaxFrames *int64 `json:"maxFrames,omitempty"`
	// Type asserts that the frame matches a known type structure.
	// Possible enum values:
	//  - `""`
	//  - `"timeseries-wide"`
	//  - `"timeseries-long"`
	//  - `"timeseries-many"`
	//  - `"timeseries-multi"`
	//  - `"directory-listing"`
	//  - `"table"`
	//  - `"numeric-wide"`
	//  - `"numeric-multi"`
	//  - `"numeric-long"`
	//  - `"log-lines"`
	Type *TypeSqlType `json:"type,omitempty"`
	// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
	// contract documentation https://grafana.github.io/dataplane/contract/.
	TypeVersion []int64 `json:"typeVersion"`
}

func (resource ExprTypeSqlResultAssertions) Equals(other ExprTypeSqlResultAssertions) bool {
	if resource.MaxFrames == nil && other.MaxFrames != nil || resource.MaxFrames != nil && other.MaxFrames == nil {
		return false
	}

	if resource.MaxFrames != nil {
		if *resource.MaxFrames != *other.MaxFrames {
			return false
		}
	}
	if resource.Type == nil && other.Type != nil || resource.Type != nil && other.Type == nil {
		return false
	}

	if resource.Type != nil {
		if *resource.Type != *other.Type {
			return false
		}
	}

	if len(resource.TypeVersion) != len(other.TypeVersion) {
		return false
	}

	for i1 := range resource.TypeVersion {
		if resource.TypeVersion[i1] != other.TypeVersion[i1] {
			return false
		}
	}

	return true
}

type ExprTypeSqlTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

func (resource ExprTypeSqlTimeRange) Equals(other ExprTypeSqlTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

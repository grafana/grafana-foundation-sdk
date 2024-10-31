// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

type Expr = TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql

// VariantConfig returns the configuration related to __expr__ dataqueries.
// This configuration describes how to unmarshal it, convert it to code, …
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
		StrictDataqueryUnmarshaler: func(raw []byte) (variants.Dataquery, error) {
			dataquery := &Expr{}

			if err := dataquery.UnmarshalJSONStrict(raw); err != nil {
				return nil, err
			}

			return dataquery, nil
		},
		GoConverter: func(input any) string {
			var dataquery Expr
			if cast, ok := input.(*Expr); ok {
				dataquery = *cast
			} else {
				dataquery = input.(Expr)
			}

			if dataquery.TypeMath != nil {
				return TypeMathConverter(*dataquery.TypeMath)
			}
			if dataquery.TypeReduce != nil {
				return TypeReduceConverter(*dataquery.TypeReduce)
			}
			if dataquery.TypeResample != nil {
				return TypeResampleConverter(*dataquery.TypeResample)
			}
			if dataquery.TypeClassicConditions != nil {
				return TypeClassicConditionsConverter(*dataquery.TypeClassicConditions)
			}
			if dataquery.TypeThreshold != nil {
				return TypeThresholdConverter(*dataquery.TypeThreshold)
			}
			if dataquery.TypeSql != nil {
				return TypeSqlConverter(*dataquery.TypeSql)
			}

			return ""
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeMath` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TypeMath) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "expression"
	if fields["expression"] != nil {
		if string(fields["expression"]) != "null" {
			if err := json.Unmarshal(fields["expression"], &resource.Expression); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expression", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is null"))...)

		}
		delete(fields, "expression")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is missing from input"))...)
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
	// Field "intervalMs"
	if fields["intervalMs"] != nil {
		if string(fields["intervalMs"]) != "null" {
			if err := json.Unmarshal(fields["intervalMs"], &resource.IntervalMs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intervalMs", err)...)
			}

		}
		delete(fields, "intervalMs")

	}
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

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
	// Field "resultAssertions"
	if fields["resultAssertions"] != nil {
		if string(fields["resultAssertions"]) != "null" {

			resource.ResultAssertions = &ExprTypeMathResultAssertions{}
			if err := resource.ResultAssertions.UnmarshalJSONStrict(fields["resultAssertions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
			}

		}
		delete(fields, "resultAssertions")

	}
	// Field "timeRange"
	if fields["timeRange"] != nil {
		if string(fields["timeRange"]) != "null" {

			resource.TimeRange = &ExprTypeMathTimeRange{}
			if err := resource.TimeRange.UnmarshalJSONStrict(fields["timeRange"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
			}

		}
		delete(fields, "timeRange")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TypeMath", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `TypeMath` fields for violations and returns them.
func (resource TypeMath) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if !(len([]rune(resource.Expression)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"expression",
			errors.New("must be >= 1"),
		)...)
	}
	if resource.ResultAssertions != nil {
		if err := resource.ResultAssertions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
		}
	}
	if resource.TimeRange != nil {
		if err := resource.TimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeReduce` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TypeReduce) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "expression"
	if fields["expression"] != nil {
		if string(fields["expression"]) != "null" {
			if err := json.Unmarshal(fields["expression"], &resource.Expression); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expression", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is null"))...)

		}
		delete(fields, "expression")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is missing from input"))...)
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
	// Field "intervalMs"
	if fields["intervalMs"] != nil {
		if string(fields["intervalMs"]) != "null" {
			if err := json.Unmarshal(fields["intervalMs"], &resource.IntervalMs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intervalMs", err)...)
			}

		}
		delete(fields, "intervalMs")

	}
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

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
	// Field "reducer"
	if fields["reducer"] != nil {
		if string(fields["reducer"]) != "null" {
			if err := json.Unmarshal(fields["reducer"], &resource.Reducer); err != nil {
				errs = append(errs, cog.MakeBuildErrors("reducer", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("reducer", errors.New("required field is null"))...)

		}
		delete(fields, "reducer")
	} else {
		errs = append(errs, cog.MakeBuildErrors("reducer", errors.New("required field is missing from input"))...)
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
	// Field "resultAssertions"
	if fields["resultAssertions"] != nil {
		if string(fields["resultAssertions"]) != "null" {

			resource.ResultAssertions = &ExprTypeReduceResultAssertions{}
			if err := resource.ResultAssertions.UnmarshalJSONStrict(fields["resultAssertions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
			}

		}
		delete(fields, "resultAssertions")

	}
	// Field "settings"
	if fields["settings"] != nil {
		if string(fields["settings"]) != "null" {

			resource.Settings = &ExprTypeReduceSettings{}
			if err := resource.Settings.UnmarshalJSONStrict(fields["settings"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("settings", err)...)
			}

		}
		delete(fields, "settings")

	}
	// Field "timeRange"
	if fields["timeRange"] != nil {
		if string(fields["timeRange"]) != "null" {

			resource.TimeRange = &ExprTypeReduceTimeRange{}
			if err := resource.TimeRange.UnmarshalJSONStrict(fields["timeRange"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
			}

		}
		delete(fields, "timeRange")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TypeReduce", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `TypeReduce` fields for violations and returns them.
func (resource TypeReduce) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if !(len([]rune(resource.Expression)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"expression",
			errors.New("must be >= 1"),
		)...)
	}
	if resource.ResultAssertions != nil {
		if err := resource.ResultAssertions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
		}
	}
	if resource.Settings != nil {
		if err := resource.Settings.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("settings", err)...)
		}
	}
	if resource.TimeRange != nil {
		if err := resource.TimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeResample` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TypeResample) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "downsampler"
	if fields["downsampler"] != nil {
		if string(fields["downsampler"]) != "null" {
			if err := json.Unmarshal(fields["downsampler"], &resource.Downsampler); err != nil {
				errs = append(errs, cog.MakeBuildErrors("downsampler", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("downsampler", errors.New("required field is null"))...)

		}
		delete(fields, "downsampler")
	} else {
		errs = append(errs, cog.MakeBuildErrors("downsampler", errors.New("required field is missing from input"))...)
	}
	// Field "expression"
	if fields["expression"] != nil {
		if string(fields["expression"]) != "null" {
			if err := json.Unmarshal(fields["expression"], &resource.Expression); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expression", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is null"))...)

		}
		delete(fields, "expression")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is missing from input"))...)
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
	// Field "intervalMs"
	if fields["intervalMs"] != nil {
		if string(fields["intervalMs"]) != "null" {
			if err := json.Unmarshal(fields["intervalMs"], &resource.IntervalMs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intervalMs", err)...)
			}

		}
		delete(fields, "intervalMs")

	}
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

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
	// Field "resultAssertions"
	if fields["resultAssertions"] != nil {
		if string(fields["resultAssertions"]) != "null" {

			resource.ResultAssertions = &ExprTypeResampleResultAssertions{}
			if err := resource.ResultAssertions.UnmarshalJSONStrict(fields["resultAssertions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
			}

		}
		delete(fields, "resultAssertions")

	}
	// Field "timeRange"
	if fields["timeRange"] != nil {
		if string(fields["timeRange"]) != "null" {

			resource.TimeRange = &ExprTypeResampleTimeRange{}
			if err := resource.TimeRange.UnmarshalJSONStrict(fields["timeRange"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
			}

		}
		delete(fields, "timeRange")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}
	// Field "upsampler"
	if fields["upsampler"] != nil {
		if string(fields["upsampler"]) != "null" {
			if err := json.Unmarshal(fields["upsampler"], &resource.Upsampler); err != nil {
				errs = append(errs, cog.MakeBuildErrors("upsampler", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("upsampler", errors.New("required field is null"))...)

		}
		delete(fields, "upsampler")
	} else {
		errs = append(errs, cog.MakeBuildErrors("upsampler", errors.New("required field is missing from input"))...)
	}
	// Field "window"
	if fields["window"] != nil {
		if string(fields["window"]) != "null" {
			if err := json.Unmarshal(fields["window"], &resource.Window); err != nil {
				errs = append(errs, cog.MakeBuildErrors("window", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is null"))...)

		}
		delete(fields, "window")
	} else {
		errs = append(errs, cog.MakeBuildErrors("window", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TypeResample", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `TypeResample` fields for violations and returns them.
func (resource TypeResample) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if !(len([]rune(resource.Expression)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"expression",
			errors.New("must be >= 1"),
		)...)
	}
	if resource.ResultAssertions != nil {
		if err := resource.ResultAssertions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
		}
	}
	if resource.TimeRange != nil {
		if err := resource.TimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
		}
	}
	if !(len([]rune(resource.Window)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"window",
			errors.New("must be >= 1"),
		)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeClassicConditions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TypeClassicConditions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "conditions"
	if fields["conditions"] != nil {
		if string(fields["conditions"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["conditions"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 ExprTypeClassicConditionsConditions

				result1 = ExprTypeClassicConditionsConditions{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("conditions["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Conditions = append(resource.Conditions, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("conditions", errors.New("required field is null"))...)

		}
		delete(fields, "conditions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("conditions", errors.New("required field is missing from input"))...)
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
	// Field "hide"
	if fields["hide"] != nil {
		if string(fields["hide"]) != "null" {
			if err := json.Unmarshal(fields["hide"], &resource.Hide); err != nil {
				errs = append(errs, cog.MakeBuildErrors("hide", err)...)
			}

		}
		delete(fields, "hide")

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
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

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
	// Field "resultAssertions"
	if fields["resultAssertions"] != nil {
		if string(fields["resultAssertions"]) != "null" {

			resource.ResultAssertions = &ExprTypeClassicConditionsResultAssertions{}
			if err := resource.ResultAssertions.UnmarshalJSONStrict(fields["resultAssertions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
			}

		}
		delete(fields, "resultAssertions")

	}
	// Field "timeRange"
	if fields["timeRange"] != nil {
		if string(fields["timeRange"]) != "null" {

			resource.TimeRange = &ExprTypeClassicConditionsTimeRange{}
			if err := resource.TimeRange.UnmarshalJSONStrict(fields["timeRange"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
			}

		}
		delete(fields, "timeRange")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TypeClassicConditions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `TypeClassicConditions` fields for violations and returns them.
func (resource TypeClassicConditions) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Conditions {
		if err := resource.Conditions[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("conditions["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if resource.ResultAssertions != nil {
		if err := resource.ResultAssertions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
		}
	}
	if resource.TimeRange != nil {
		if err := resource.TimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeThreshold` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TypeThreshold) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "conditions"
	if fields["conditions"] != nil {
		if string(fields["conditions"]) != "null" {

			partialArray := []json.RawMessage{}
			if err := json.Unmarshal(fields["conditions"], &partialArray); err != nil {
				return err
			}

			for i1 := range partialArray {
				var result1 ExprTypeThresholdConditions

				result1 = ExprTypeThresholdConditions{}
				if err := result1.UnmarshalJSONStrict(partialArray[i1]); err != nil {
					errs = append(errs, cog.MakeBuildErrors("conditions["+strconv.Itoa(i1)+"]", err)...)
				}
				resource.Conditions = append(resource.Conditions, result1)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("conditions", errors.New("required field is null"))...)

		}
		delete(fields, "conditions")
	} else {
		errs = append(errs, cog.MakeBuildErrors("conditions", errors.New("required field is missing from input"))...)
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
	// Field "expression"
	if fields["expression"] != nil {
		if string(fields["expression"]) != "null" {
			if err := json.Unmarshal(fields["expression"], &resource.Expression); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expression", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is null"))...)

		}
		delete(fields, "expression")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is missing from input"))...)
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
	// Field "intervalMs"
	if fields["intervalMs"] != nil {
		if string(fields["intervalMs"]) != "null" {
			if err := json.Unmarshal(fields["intervalMs"], &resource.IntervalMs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intervalMs", err)...)
			}

		}
		delete(fields, "intervalMs")

	}
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

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
	// Field "resultAssertions"
	if fields["resultAssertions"] != nil {
		if string(fields["resultAssertions"]) != "null" {

			resource.ResultAssertions = &ExprTypeThresholdResultAssertions{}
			if err := resource.ResultAssertions.UnmarshalJSONStrict(fields["resultAssertions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
			}

		}
		delete(fields, "resultAssertions")

	}
	// Field "timeRange"
	if fields["timeRange"] != nil {
		if string(fields["timeRange"]) != "null" {

			resource.TimeRange = &ExprTypeThresholdTimeRange{}
			if err := resource.TimeRange.UnmarshalJSONStrict(fields["timeRange"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
			}

		}
		delete(fields, "timeRange")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TypeThreshold", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `TypeThreshold` fields for violations and returns them.
func (resource TypeThreshold) Validate() error {
	var errs cog.BuildErrors

	for i1 := range resource.Conditions {
		if err := resource.Conditions[i1].Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("conditions["+strconv.Itoa(i1)+"]", err)...)
		}
	}
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if !(len([]rune(resource.Expression)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"expression",
			errors.New("must be >= 1"),
		)...)
	}
	if resource.ResultAssertions != nil {
		if err := resource.ResultAssertions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
		}
	}
	if resource.TimeRange != nil {
		if err := resource.TimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeSql` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TypeSql) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
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
	// Field "expression"
	if fields["expression"] != nil {
		if string(fields["expression"]) != "null" {
			if err := json.Unmarshal(fields["expression"], &resource.Expression); err != nil {
				errs = append(errs, cog.MakeBuildErrors("expression", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is null"))...)

		}
		delete(fields, "expression")
	} else {
		errs = append(errs, cog.MakeBuildErrors("expression", errors.New("required field is missing from input"))...)
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
	// Field "intervalMs"
	if fields["intervalMs"] != nil {
		if string(fields["intervalMs"]) != "null" {
			if err := json.Unmarshal(fields["intervalMs"], &resource.IntervalMs); err != nil {
				errs = append(errs, cog.MakeBuildErrors("intervalMs", err)...)
			}

		}
		delete(fields, "intervalMs")

	}
	// Field "maxDataPoints"
	if fields["maxDataPoints"] != nil {
		if string(fields["maxDataPoints"]) != "null" {
			if err := json.Unmarshal(fields["maxDataPoints"], &resource.MaxDataPoints); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxDataPoints", err)...)
			}

		}
		delete(fields, "maxDataPoints")

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
	// Field "resultAssertions"
	if fields["resultAssertions"] != nil {
		if string(fields["resultAssertions"]) != "null" {

			resource.ResultAssertions = &ExprTypeSqlResultAssertions{}
			if err := resource.ResultAssertions.UnmarshalJSONStrict(fields["resultAssertions"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
			}

		}
		delete(fields, "resultAssertions")

	}
	// Field "timeRange"
	if fields["timeRange"] != nil {
		if string(fields["timeRange"]) != "null" {

			resource.TimeRange = &ExprTypeSqlTimeRange{}
			if err := resource.TimeRange.UnmarshalJSONStrict(fields["timeRange"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
			}

		}
		delete(fields, "timeRange")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("TypeSql", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `TypeSql` fields for violations and returns them.
func (resource TypeSql) Validate() error {
	var errs cog.BuildErrors
	if resource.Datasource != nil {
		if err := resource.Datasource.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("datasource", err)...)
		}
	}
	if !(len([]rune(resource.Expression)) >= 1) {
		errs = append(errs, cog.MakeBuildErrors(
			"expression",
			errors.New("must be >= 1"),
		)...)
	}
	if resource.ResultAssertions != nil {
		if err := resource.ResultAssertions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("resultAssertions", err)...)
		}
	}
	if resource.TimeRange != nil {
		if err := resource.TimeRange.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("timeRange", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// MarshalJSON implements a custom JSON marshalling logic to encode `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` as JSON.
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

// UnmarshalJSON implements a custom JSON unmarshalling logic to decode `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` from JSON.
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) UnmarshalJSONStrict(raw []byte) error {
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
		return fmt.Errorf("discriminator field 'type' not found in payload")
	}

	switch discriminator {
	case "classic_conditions":
		typeClassicConditions := &TypeClassicConditions{}
		if err := typeClassicConditions.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TypeClassicConditions = typeClassicConditions
		return nil
	case "math":
		typeMath := &TypeMath{}
		if err := typeMath.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TypeMath = typeMath
		return nil
	case "reduce":
		typeReduce := &TypeReduce{}
		if err := typeReduce.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TypeReduce = typeReduce
		return nil
	case "resample":
		typeResample := &TypeResample{}
		if err := typeResample.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TypeResample = typeResample
		return nil
	case "sql":
		typeSql := &TypeSql{}
		if err := typeSql.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TypeSql = typeSql
		return nil
	case "threshold":
		typeThreshold := &TypeThreshold{}
		if err := typeThreshold.UnmarshalJSONStrict(raw); err != nil {
			return err
		}

		resource.TypeThreshold = typeThreshold
		return nil
	}

	return fmt.Errorf("could not unmarshal resource with `type = %v`", discriminator)
}

// Equals tests the equality of two dataqueries.
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

// Validate checks all the validation constraints that may be defined on `TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql` fields for violations and returns them.
func (resource TypeMathOrTypeReduceOrTypeResampleOrTypeClassicConditionsOrTypeThresholdOrTypeSql) Validate() error {
	var errs cog.BuildErrors
	if resource.TypeMath != nil {
		if err := resource.TypeMath.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TypeMath", err)...)
		}
	}
	if resource.TypeReduce != nil {
		if err := resource.TypeReduce.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TypeReduce", err)...)
		}
	}
	if resource.TypeResample != nil {
		if err := resource.TypeResample.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TypeResample", err)...)
		}
	}
	if resource.TypeClassicConditions != nil {
		if err := resource.TypeClassicConditions.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TypeClassicConditions", err)...)
		}
	}
	if resource.TypeThreshold != nil {
		if err := resource.TypeThreshold.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TypeThreshold", err)...)
		}
	}
	if resource.TypeSql != nil {
		if err := resource.TypeSql.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("TypeSql", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeMathResultAssertions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeMathResultAssertions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "maxFrames"
	if fields["maxFrames"] != nil {
		if string(fields["maxFrames"]) != "null" {
			if err := json.Unmarshal(fields["maxFrames"], &resource.MaxFrames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxFrames", err)...)
			}

		}
		delete(fields, "maxFrames")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}

		}
		delete(fields, "type")

	}
	// Field "typeVersion"
	if fields["typeVersion"] != nil {
		if string(fields["typeVersion"]) != "null" {

			if err := json.Unmarshal(fields["typeVersion"], &resource.TypeVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("typeVersion", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is null"))...)

		}
		delete(fields, "typeVersion")
	} else {
		errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeMathResultAssertions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeMathResultAssertions` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeMathResultAssertions` fields for violations and returns them.
func (resource ExprTypeMathResultAssertions) Validate() error {
	return nil
}

type ExprTypeMathTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeMathTimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeMathTimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeMathTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeMathTimeRange` objects.
func (resource ExprTypeMathTimeRange) Equals(other ExprTypeMathTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExprTypeMathTimeRange` fields for violations and returns them.
func (resource ExprTypeMathTimeRange) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeReduceResultAssertions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeReduceResultAssertions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "maxFrames"
	if fields["maxFrames"] != nil {
		if string(fields["maxFrames"]) != "null" {
			if err := json.Unmarshal(fields["maxFrames"], &resource.MaxFrames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxFrames", err)...)
			}

		}
		delete(fields, "maxFrames")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}

		}
		delete(fields, "type")

	}
	// Field "typeVersion"
	if fields["typeVersion"] != nil {
		if string(fields["typeVersion"]) != "null" {

			if err := json.Unmarshal(fields["typeVersion"], &resource.TypeVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("typeVersion", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is null"))...)

		}
		delete(fields, "typeVersion")
	} else {
		errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeReduceResultAssertions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeReduceResultAssertions` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeReduceResultAssertions` fields for violations and returns them.
func (resource ExprTypeReduceResultAssertions) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeReduceSettings` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeReduceSettings) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "mode"
	if fields["mode"] != nil {
		if string(fields["mode"]) != "null" {
			if err := json.Unmarshal(fields["mode"], &resource.Mode); err != nil {
				errs = append(errs, cog.MakeBuildErrors("mode", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is null"))...)

		}
		delete(fields, "mode")
	} else {
		errs = append(errs, cog.MakeBuildErrors("mode", errors.New("required field is missing from input"))...)
	}
	// Field "replaceWithValue"
	if fields["replaceWithValue"] != nil {
		if string(fields["replaceWithValue"]) != "null" {
			if err := json.Unmarshal(fields["replaceWithValue"], &resource.ReplaceWithValue); err != nil {
				errs = append(errs, cog.MakeBuildErrors("replaceWithValue", err)...)
			}

		}
		delete(fields, "replaceWithValue")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeReduceSettings", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeReduceSettings` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeReduceSettings` fields for violations and returns them.
func (resource ExprTypeReduceSettings) Validate() error {
	return nil
}

type ExprTypeReduceTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeReduceTimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeReduceTimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeReduceTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeReduceTimeRange` objects.
func (resource ExprTypeReduceTimeRange) Equals(other ExprTypeReduceTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExprTypeReduceTimeRange` fields for violations and returns them.
func (resource ExprTypeReduceTimeRange) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeResampleResultAssertions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeResampleResultAssertions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "maxFrames"
	if fields["maxFrames"] != nil {
		if string(fields["maxFrames"]) != "null" {
			if err := json.Unmarshal(fields["maxFrames"], &resource.MaxFrames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxFrames", err)...)
			}

		}
		delete(fields, "maxFrames")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}

		}
		delete(fields, "type")

	}
	// Field "typeVersion"
	if fields["typeVersion"] != nil {
		if string(fields["typeVersion"]) != "null" {

			if err := json.Unmarshal(fields["typeVersion"], &resource.TypeVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("typeVersion", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is null"))...)

		}
		delete(fields, "typeVersion")
	} else {
		errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeResampleResultAssertions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeResampleResultAssertions` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeResampleResultAssertions` fields for violations and returns them.
func (resource ExprTypeResampleResultAssertions) Validate() error {
	return nil
}

type ExprTypeResampleTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeResampleTimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeResampleTimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeResampleTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeResampleTimeRange` objects.
func (resource ExprTypeResampleTimeRange) Equals(other ExprTypeResampleTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExprTypeResampleTimeRange` fields for violations and returns them.
func (resource ExprTypeResampleTimeRange) Validate() error {
	return nil
}

type ExprTypeClassicConditionsConditionsEvaluator struct {
	Params []float64 `json:"params"`
	// e.g. "gt"
	Type string `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditionsEvaluator` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeClassicConditionsConditionsEvaluator) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "params"
	if fields["params"] != nil {
		if string(fields["params"]) != "null" {

			if err := json.Unmarshal(fields["params"], &resource.Params); err != nil {
				errs = append(errs, cog.MakeBuildErrors("params", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("params", errors.New("required field is null"))...)

		}
		delete(fields, "params")
	} else {
		errs = append(errs, cog.MakeBuildErrors("params", errors.New("required field is missing from input"))...)
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditionsEvaluator", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeClassicConditionsConditionsEvaluator` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditionsEvaluator` fields for violations and returns them.
func (resource ExprTypeClassicConditionsConditionsEvaluator) Validate() error {
	return nil
}

type ExprTypeClassicConditionsConditionsOperator struct {
	Type TypeClassicConditionsType `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditionsOperator` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeClassicConditionsConditionsOperator) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditionsOperator", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeClassicConditionsConditionsOperator` objects.
func (resource ExprTypeClassicConditionsConditionsOperator) Equals(other ExprTypeClassicConditionsConditionsOperator) bool {
	if resource.Type != other.Type {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditionsOperator` fields for violations and returns them.
func (resource ExprTypeClassicConditionsConditionsOperator) Validate() error {
	return nil
}

type ExprTypeClassicConditionsConditionsQuery struct {
	Params []string `json:"params"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditionsQuery` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeClassicConditionsConditionsQuery) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "params"
	if fields["params"] != nil {
		if string(fields["params"]) != "null" {

			if err := json.Unmarshal(fields["params"], &resource.Params); err != nil {
				errs = append(errs, cog.MakeBuildErrors("params", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("params", errors.New("required field is null"))...)

		}
		delete(fields, "params")
	} else {
		errs = append(errs, cog.MakeBuildErrors("params", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditionsQuery", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeClassicConditionsConditionsQuery` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditionsQuery` fields for violations and returns them.
func (resource ExprTypeClassicConditionsConditionsQuery) Validate() error {
	return nil
}

type ExprTypeClassicConditionsConditionsReducer struct {
	Type string `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditionsReducer` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeClassicConditionsConditionsReducer) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditionsReducer", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeClassicConditionsConditionsReducer` objects.
func (resource ExprTypeClassicConditionsConditionsReducer) Equals(other ExprTypeClassicConditionsConditionsReducer) bool {
	if resource.Type != other.Type {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditionsReducer` fields for violations and returns them.
func (resource ExprTypeClassicConditionsConditionsReducer) Validate() error {
	return nil
}

type ExprTypeClassicConditionsConditions struct {
	Evaluator ExprTypeClassicConditionsConditionsEvaluator `json:"evaluator"`
	Operator  ExprTypeClassicConditionsConditionsOperator  `json:"operator"`
	Query     ExprTypeClassicConditionsConditionsQuery     `json:"query"`
	Reducer   ExprTypeClassicConditionsConditionsReducer   `json:"reducer"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsConditions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeClassicConditionsConditions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "evaluator"
	if fields["evaluator"] != nil {
		if string(fields["evaluator"]) != "null" {

			resource.Evaluator = ExprTypeClassicConditionsConditionsEvaluator{}
			if err := resource.Evaluator.UnmarshalJSONStrict(fields["evaluator"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("evaluator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("evaluator", errors.New("required field is null"))...)

		}
		delete(fields, "evaluator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("evaluator", errors.New("required field is missing from input"))...)
	}
	// Field "operator"
	if fields["operator"] != nil {
		if string(fields["operator"]) != "null" {

			resource.Operator = ExprTypeClassicConditionsConditionsOperator{}
			if err := resource.Operator.UnmarshalJSONStrict(fields["operator"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("operator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is null"))...)

		}
		delete(fields, "operator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("operator", errors.New("required field is missing from input"))...)
	}
	// Field "query"
	if fields["query"] != nil {
		if string(fields["query"]) != "null" {

			resource.Query = ExprTypeClassicConditionsConditionsQuery{}
			if err := resource.Query.UnmarshalJSONStrict(fields["query"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("query", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is null"))...)

		}
		delete(fields, "query")
	} else {
		errs = append(errs, cog.MakeBuildErrors("query", errors.New("required field is missing from input"))...)
	}
	// Field "reducer"
	if fields["reducer"] != nil {
		if string(fields["reducer"]) != "null" {

			resource.Reducer = ExprTypeClassicConditionsConditionsReducer{}
			if err := resource.Reducer.UnmarshalJSONStrict(fields["reducer"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("reducer", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("reducer", errors.New("required field is null"))...)

		}
		delete(fields, "reducer")
	} else {
		errs = append(errs, cog.MakeBuildErrors("reducer", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsConditions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeClassicConditionsConditions` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsConditions` fields for violations and returns them.
func (resource ExprTypeClassicConditionsConditions) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Evaluator.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("evaluator", err)...)
	}
	if err := resource.Operator.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("operator", err)...)
	}
	if err := resource.Query.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("query", err)...)
	}
	if err := resource.Reducer.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("reducer", err)...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsResultAssertions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeClassicConditionsResultAssertions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "maxFrames"
	if fields["maxFrames"] != nil {
		if string(fields["maxFrames"]) != "null" {
			if err := json.Unmarshal(fields["maxFrames"], &resource.MaxFrames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxFrames", err)...)
			}

		}
		delete(fields, "maxFrames")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}

		}
		delete(fields, "type")

	}
	// Field "typeVersion"
	if fields["typeVersion"] != nil {
		if string(fields["typeVersion"]) != "null" {

			if err := json.Unmarshal(fields["typeVersion"], &resource.TypeVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("typeVersion", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is null"))...)

		}
		delete(fields, "typeVersion")
	} else {
		errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsResultAssertions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeClassicConditionsResultAssertions` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsResultAssertions` fields for violations and returns them.
func (resource ExprTypeClassicConditionsResultAssertions) Validate() error {
	return nil
}

type ExprTypeClassicConditionsTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeClassicConditionsTimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeClassicConditionsTimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeClassicConditionsTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeClassicConditionsTimeRange` objects.
func (resource ExprTypeClassicConditionsTimeRange) Equals(other ExprTypeClassicConditionsTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExprTypeClassicConditionsTimeRange` fields for violations and returns them.
func (resource ExprTypeClassicConditionsTimeRange) Validate() error {
	return nil
}

type ExprTypeThresholdConditionsEvaluator struct {
	Params []float64 `json:"params"`
	// e.g. "gt"
	Type TypeThresholdType `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdConditionsEvaluator` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeThresholdConditionsEvaluator) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "params"
	if fields["params"] != nil {
		if string(fields["params"]) != "null" {

			if err := json.Unmarshal(fields["params"], &resource.Params); err != nil {
				errs = append(errs, cog.MakeBuildErrors("params", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("params", errors.New("required field is null"))...)

		}
		delete(fields, "params")
	} else {
		errs = append(errs, cog.MakeBuildErrors("params", errors.New("required field is missing from input"))...)
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeThresholdConditionsEvaluator", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeThresholdConditionsEvaluator` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeThresholdConditionsEvaluator` fields for violations and returns them.
func (resource ExprTypeThresholdConditionsEvaluator) Validate() error {
	return nil
}

type ExprTypeThresholdConditionsUnloadEvaluator struct {
	Params []float64 `json:"params"`
	// e.g. "gt"
	Type TypeThresholdType `json:"type"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdConditionsUnloadEvaluator` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeThresholdConditionsUnloadEvaluator) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "params"
	if fields["params"] != nil {
		if string(fields["params"]) != "null" {

			if err := json.Unmarshal(fields["params"], &resource.Params); err != nil {
				errs = append(errs, cog.MakeBuildErrors("params", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("params", errors.New("required field is null"))...)

		}
		delete(fields, "params")
	} else {
		errs = append(errs, cog.MakeBuildErrors("params", errors.New("required field is missing from input"))...)
	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is null"))...)

		}
		delete(fields, "type")
	} else {
		errs = append(errs, cog.MakeBuildErrors("type", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeThresholdConditionsUnloadEvaluator", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeThresholdConditionsUnloadEvaluator` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeThresholdConditionsUnloadEvaluator` fields for violations and returns them.
func (resource ExprTypeThresholdConditionsUnloadEvaluator) Validate() error {
	return nil
}

type ExprTypeThresholdConditions struct {
	Evaluator        ExprTypeThresholdConditionsEvaluator        `json:"evaluator"`
	LoadedDimensions any                                         `json:"loadedDimensions,omitempty"`
	UnloadEvaluator  *ExprTypeThresholdConditionsUnloadEvaluator `json:"unloadEvaluator,omitempty"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdConditions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeThresholdConditions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "evaluator"
	if fields["evaluator"] != nil {
		if string(fields["evaluator"]) != "null" {

			resource.Evaluator = ExprTypeThresholdConditionsEvaluator{}
			if err := resource.Evaluator.UnmarshalJSONStrict(fields["evaluator"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("evaluator", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("evaluator", errors.New("required field is null"))...)

		}
		delete(fields, "evaluator")
	} else {
		errs = append(errs, cog.MakeBuildErrors("evaluator", errors.New("required field is missing from input"))...)
	}
	// Field "loadedDimensions"
	if fields["loadedDimensions"] != nil {
		if string(fields["loadedDimensions"]) != "null" {
			if err := json.Unmarshal(fields["loadedDimensions"], &resource.LoadedDimensions); err != nil {
				errs = append(errs, cog.MakeBuildErrors("loadedDimensions", err)...)
			}

		}
		delete(fields, "loadedDimensions")

	}
	// Field "unloadEvaluator"
	if fields["unloadEvaluator"] != nil {
		if string(fields["unloadEvaluator"]) != "null" {

			resource.UnloadEvaluator = &ExprTypeThresholdConditionsUnloadEvaluator{}
			if err := resource.UnloadEvaluator.UnmarshalJSONStrict(fields["unloadEvaluator"]); err != nil {
				errs = append(errs, cog.MakeBuildErrors("unloadEvaluator", err)...)
			}

		}
		delete(fields, "unloadEvaluator")

	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeThresholdConditions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeThresholdConditions` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeThresholdConditions` fields for violations and returns them.
func (resource ExprTypeThresholdConditions) Validate() error {
	var errs cog.BuildErrors
	if err := resource.Evaluator.Validate(); err != nil {
		errs = append(errs, cog.MakeBuildErrors("evaluator", err)...)
	}
	if resource.UnloadEvaluator != nil {
		if err := resource.UnloadEvaluator.Validate(); err != nil {
			errs = append(errs, cog.MakeBuildErrors("unloadEvaluator", err)...)
		}
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdResultAssertions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeThresholdResultAssertions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "maxFrames"
	if fields["maxFrames"] != nil {
		if string(fields["maxFrames"]) != "null" {
			if err := json.Unmarshal(fields["maxFrames"], &resource.MaxFrames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxFrames", err)...)
			}

		}
		delete(fields, "maxFrames")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}

		}
		delete(fields, "type")

	}
	// Field "typeVersion"
	if fields["typeVersion"] != nil {
		if string(fields["typeVersion"]) != "null" {

			if err := json.Unmarshal(fields["typeVersion"], &resource.TypeVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("typeVersion", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is null"))...)

		}
		delete(fields, "typeVersion")
	} else {
		errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeThresholdResultAssertions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeThresholdResultAssertions` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeThresholdResultAssertions` fields for violations and returns them.
func (resource ExprTypeThresholdResultAssertions) Validate() error {
	return nil
}

type ExprTypeThresholdTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeThresholdTimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeThresholdTimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeThresholdTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeThresholdTimeRange` objects.
func (resource ExprTypeThresholdTimeRange) Equals(other ExprTypeThresholdTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExprTypeThresholdTimeRange` fields for violations and returns them.
func (resource ExprTypeThresholdTimeRange) Validate() error {
	return nil
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

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeSqlResultAssertions` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeSqlResultAssertions) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "maxFrames"
	if fields["maxFrames"] != nil {
		if string(fields["maxFrames"]) != "null" {
			if err := json.Unmarshal(fields["maxFrames"], &resource.MaxFrames); err != nil {
				errs = append(errs, cog.MakeBuildErrors("maxFrames", err)...)
			}

		}
		delete(fields, "maxFrames")

	}
	// Field "type"
	if fields["type"] != nil {
		if string(fields["type"]) != "null" {
			if err := json.Unmarshal(fields["type"], &resource.Type); err != nil {
				errs = append(errs, cog.MakeBuildErrors("type", err)...)
			}

		}
		delete(fields, "type")

	}
	// Field "typeVersion"
	if fields["typeVersion"] != nil {
		if string(fields["typeVersion"]) != "null" {

			if err := json.Unmarshal(fields["typeVersion"], &resource.TypeVersion); err != nil {
				errs = append(errs, cog.MakeBuildErrors("typeVersion", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is null"))...)

		}
		delete(fields, "typeVersion")
	} else {
		errs = append(errs, cog.MakeBuildErrors("typeVersion", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeSqlResultAssertions", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeSqlResultAssertions` objects.
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

// Validate checks all the validation constraints that may be defined on `ExprTypeSqlResultAssertions` fields for violations and returns them.
func (resource ExprTypeSqlResultAssertions) Validate() error {
	return nil
}

type ExprTypeSqlTimeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}

// UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ExprTypeSqlTimeRange` from JSON.
// Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …
func (resource *ExprTypeSqlTimeRange) UnmarshalJSONStrict(raw []byte) error {
	if raw == nil {
		return nil
	}
	var errs cog.BuildErrors

	fields := make(map[string]json.RawMessage)
	if err := json.Unmarshal(raw, &fields); err != nil {
		return err
	}
	// Field "from"
	if fields["from"] != nil {
		if string(fields["from"]) != "null" {
			if err := json.Unmarshal(fields["from"], &resource.From); err != nil {
				errs = append(errs, cog.MakeBuildErrors("from", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is null"))...)

		}
		delete(fields, "from")
	} else {
		errs = append(errs, cog.MakeBuildErrors("from", errors.New("required field is missing from input"))...)
	}
	// Field "to"
	if fields["to"] != nil {
		if string(fields["to"]) != "null" {
			if err := json.Unmarshal(fields["to"], &resource.To); err != nil {
				errs = append(errs, cog.MakeBuildErrors("to", err)...)
			}
		} else {
			errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is null"))...)

		}
		delete(fields, "to")
	} else {
		errs = append(errs, cog.MakeBuildErrors("to", errors.New("required field is missing from input"))...)
	}

	for field := range fields {
		errs = append(errs, cog.MakeBuildErrors("ExprTypeSqlTimeRange", fmt.Errorf("unexpected field '%s'", field))...)
	}

	if len(errs) == 0 {
		return nil
	}

	return errs
}

// Equals tests the equality of two `ExprTypeSqlTimeRange` objects.
func (resource ExprTypeSqlTimeRange) Equals(other ExprTypeSqlTimeRange) bool {
	if resource.From != other.From {
		return false
	}
	if resource.To != other.To {
		return false
	}

	return true
}

// Validate checks all the validation constraints that may be defined on `ExprTypeSqlTimeRange` fields for violations and returns them.
func (resource ExprTypeSqlTimeRange) Validate() error {
	return nil
}

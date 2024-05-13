// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[cogvariants.Dataquery] = (*TypeClassicConditionsBuilder)(nil)

type TypeClassicConditionsBuilder struct {
	internal *TypeClassicConditions
	errors   map[string]cog.BuildErrors
}

func NewTypeClassicConditionsBuilder() *TypeClassicConditionsBuilder {
	resource := &TypeClassicConditions{}
	builder := &TypeClassicConditionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "classic_conditions"

	return builder
}

func (builder *TypeClassicConditionsBuilder) Build() (cogvariants.Dataquery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("TypeClassicConditions", err)...)
	}

	if len(errs) != 0 {
		return TypeClassicConditions{}, errs
	}

	return *builder.internal, nil
}

func (builder *TypeClassicConditionsBuilder) Conditions(conditions []struct {
	Evaluator struct {
		Params []float64 `json:"params"`
		// e.g. "gt"
		Type string `json:"type"`
	} `json:"evaluator"`
	Operator struct {
		Type TypeClassicConditionsType `json:"type"`
	} `json:"operator"`
	Query struct {
		Params []string `json:"params"`
	} `json:"query"`
	Reducer struct {
		Type string `json:"type"`
	} `json:"reducer"`
}) *TypeClassicConditionsBuilder {
	builder.internal.Conditions = conditions

	return builder
}

// The datasource
func (builder *TypeClassicConditionsBuilder) Datasource(datasource struct {
	// The datasource plugin type
	Type string `json:"type"`
	// Datasource UID
	Uid *string `json:"uid,omitempty"`
}) *TypeClassicConditionsBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *TypeClassicConditionsBuilder) Hide(hide bool) *TypeClassicConditionsBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *TypeClassicConditionsBuilder) IntervalMs(intervalMs float64) *TypeClassicConditionsBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *TypeClassicConditionsBuilder) MaxDataPoints(maxDataPoints int64) *TypeClassicConditionsBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *TypeClassicConditionsBuilder) QueryType(queryType string) *TypeClassicConditionsBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *TypeClassicConditionsBuilder) RefId(refId string) *TypeClassicConditionsBuilder {
	builder.internal.RefId = refId

	return builder
}

// Optionally define expected query result behavior
func (builder *TypeClassicConditionsBuilder) ResultAssertions(resultAssertions struct {
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
}) *TypeClassicConditionsBuilder {
	builder.internal.ResultAssertions = &resultAssertions

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *TypeClassicConditionsBuilder) TimeRange(timeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}) *TypeClassicConditionsBuilder {
	builder.internal.TimeRange = &timeRange

	return builder
}

func (builder *TypeClassicConditionsBuilder) applyDefaults() {
}

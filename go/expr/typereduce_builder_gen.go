// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[cogvariants.Dataquery] = (*TypeReduceBuilder)(nil)

type TypeReduceBuilder struct {
	internal *TypeReduce
	errors   map[string]cog.BuildErrors
}

func NewTypeReduceBuilder() *TypeReduceBuilder {
	resource := &TypeReduce{}
	builder := &TypeReduceBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "reduce"

	return builder
}

func (builder *TypeReduceBuilder) Build() (cogvariants.Dataquery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("TypeReduce", err)...)
	}

	if len(errs) != 0 {
		return TypeReduce{}, errs
	}

	return *builder.internal, nil
}

// The datasource
func (builder *TypeReduceBuilder) Datasource(datasource struct {
	// The apiserver version
	ApiVersion *string `json:"apiVersion,omitempty"`
	// The datasource plugin type
	Type string `json:"type"`
	// Datasource UID (NOTE: name in k8s)
	Uid *string `json:"uid,omitempty"`
}) *TypeReduceBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Reference to single query result
func (builder *TypeReduceBuilder) Expression(expression string) *TypeReduceBuilder {
	if !(len([]rune(expression)) >= 1) {
		builder.errors["expression"] = cog.MakeBuildErrors("expression", errors.New("len([]rune(expression)) must be >= 1"))
		return builder
	}
	builder.internal.Expression = expression

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// NOTE: this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *TypeReduceBuilder) Hide(hide bool) *TypeReduceBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *TypeReduceBuilder) IntervalMs(intervalMs float64) *TypeReduceBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *TypeReduceBuilder) MaxDataPoints(maxDataPoints int64) *TypeReduceBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *TypeReduceBuilder) QueryType(queryType string) *TypeReduceBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// The reducer
// Possible enum values:
//   - `"sum"`
//   - `"mean"`
//   - `"min"`
//   - `"max"`
//   - `"count"`
//   - `"last"`
func (builder *TypeReduceBuilder) Reducer(reducer TypeReduceReducer) *TypeReduceBuilder {
	builder.internal.Reducer = reducer

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *TypeReduceBuilder) RefId(refId string) *TypeReduceBuilder {
	builder.internal.RefId = refId

	return builder
}

// Optionally define expected query result behavior
func (builder *TypeReduceBuilder) ResultAssertions(resultAssertions struct {
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
}) *TypeReduceBuilder {
	builder.internal.ResultAssertions = &resultAssertions

	return builder
}

// Reducer Options
func (builder *TypeReduceBuilder) Settings(settings struct {
	// Non-number reduce behavior
	// Possible enum values:
	//  - `"dropNN"` Drop non-numbers
	//  - `"replaceNN"` Replace non-numbers
	Mode TypeReduceMode `json:"mode"`
	// Only valid when mode is replace
	ReplaceWithValue *float64 `json:"replaceWithValue,omitempty"`
}) *TypeReduceBuilder {
	builder.internal.Settings = &settings

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *TypeReduceBuilder) TimeRange(timeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}) *TypeReduceBuilder {
	builder.internal.TimeRange = &timeRange

	return builder
}

func (builder *TypeReduceBuilder) applyDefaults() {
}

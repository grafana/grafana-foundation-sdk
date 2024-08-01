// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[variants.Dataquery] = (*TypeMathBuilder)(nil)

type TypeMathBuilder struct {
	internal *TypeMath
	errors   map[string]cog.BuildErrors
}

func NewTypeMathBuilder() *TypeMathBuilder {
	resource := &TypeMath{}
	builder := &TypeMathBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "math"

	return builder
}

func (builder *TypeMathBuilder) Build() (variants.Dataquery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("TypeMath", err)...)
	}

	if len(errs) != 0 {
		return TypeMath{}, errs
	}

	return *builder.internal, nil
}

// The datasource
func (builder *TypeMathBuilder) Datasource(datasource struct {
	// The apiserver version
	ApiVersion *string `json:"apiVersion,omitempty"`
	// The datasource plugin type
	Type string `json:"type"`
	// Datasource UID (NOTE: name in k8s)
	Uid *string `json:"uid,omitempty"`
}) *TypeMathBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// General math expression
func (builder *TypeMathBuilder) Expression(expression string) *TypeMathBuilder {
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
func (builder *TypeMathBuilder) Hide(hide bool) *TypeMathBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *TypeMathBuilder) IntervalMs(intervalMs float64) *TypeMathBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *TypeMathBuilder) MaxDataPoints(maxDataPoints int64) *TypeMathBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *TypeMathBuilder) QueryType(queryType string) *TypeMathBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *TypeMathBuilder) RefId(refId string) *TypeMathBuilder {
	builder.internal.RefId = refId

	return builder
}

// Optionally define expected query result behavior
func (builder *TypeMathBuilder) ResultAssertions(resultAssertions struct {
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
}) *TypeMathBuilder {
	builder.internal.ResultAssertions = &resultAssertions

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *TypeMathBuilder) TimeRange(timeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}) *TypeMathBuilder {
	builder.internal.TimeRange = &timeRange

	return builder
}

func (builder *TypeMathBuilder) applyDefaults() {
}

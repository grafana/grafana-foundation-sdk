// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[variants.Dataquery] = (*TypeThresholdBuilder)(nil)

type TypeThresholdBuilder struct {
	internal *TypeThreshold
	errors   map[string]cog.BuildErrors
}

func NewTypeThresholdBuilder() *TypeThresholdBuilder {
	resource := &TypeThreshold{}
	builder := &TypeThresholdBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "threshold"

	return builder
}

func (builder *TypeThresholdBuilder) Build() (variants.Dataquery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("TypeThreshold", err)...)
	}

	if len(errs) != 0 {
		return TypeThreshold{}, errs
	}

	return *builder.internal, nil
}

// Threshold Conditions
func (builder *TypeThresholdBuilder) Conditions(conditions []struct {
	Evaluator struct {
		Params []float64 `json:"params"`
		// e.g. "gt"
		Type TypeThresholdType `json:"type"`
	} `json:"evaluator"`
	LoadedDimensions any `json:"loadedDimensions,omitempty"`
	UnloadEvaluator  *struct {
		Params []float64 `json:"params"`
		// e.g. "gt"
		Type TypeThresholdType `json:"type"`
	} `json:"unloadEvaluator,omitempty"`
}) *TypeThresholdBuilder {
	builder.internal.Conditions = conditions

	return builder
}

// The datasource
func (builder *TypeThresholdBuilder) Datasource(datasource struct {
	// The apiserver version
	ApiVersion *string `json:"apiVersion,omitempty"`
	// The datasource plugin type
	Type string `json:"type"`
	// Datasource UID (NOTE: name in k8s)
	Uid *string `json:"uid,omitempty"`
}) *TypeThresholdBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Reference to single query result
func (builder *TypeThresholdBuilder) Expression(expression string) *TypeThresholdBuilder {
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
func (builder *TypeThresholdBuilder) Hide(hide bool) *TypeThresholdBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *TypeThresholdBuilder) IntervalMs(intervalMs float64) *TypeThresholdBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *TypeThresholdBuilder) MaxDataPoints(maxDataPoints int64) *TypeThresholdBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *TypeThresholdBuilder) QueryType(queryType string) *TypeThresholdBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *TypeThresholdBuilder) RefId(refId string) *TypeThresholdBuilder {
	builder.internal.RefId = refId

	return builder
}

// Optionally define expected query result behavior
func (builder *TypeThresholdBuilder) ResultAssertions(resultAssertions struct {
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
}) *TypeThresholdBuilder {
	builder.internal.ResultAssertions = &resultAssertions

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *TypeThresholdBuilder) TimeRange(timeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}) *TypeThresholdBuilder {
	builder.internal.TimeRange = &timeRange

	return builder
}

func (builder *TypeThresholdBuilder) applyDefaults() {
}

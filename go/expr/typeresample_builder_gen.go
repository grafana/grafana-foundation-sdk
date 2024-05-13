// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	"errors"

	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[cogvariants.Dataquery] = (*TypeResampleBuilder)(nil)

type TypeResampleBuilder struct {
	internal *TypeResample
	errors   map[string]cog.BuildErrors
}

func NewTypeResampleBuilder() *TypeResampleBuilder {
	resource := &TypeResample{}
	builder := &TypeResampleBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "resample"

	return builder
}

func (builder *TypeResampleBuilder) Build() (cogvariants.Dataquery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("TypeResample", err)...)
	}

	if len(errs) != 0 {
		return TypeResample{}, errs
	}

	return *builder.internal, nil
}

// The datasource
func (builder *TypeResampleBuilder) Datasource(datasource struct {
	// The datasource plugin type
	Type string `json:"type"`
	// Datasource UID
	Uid *string `json:"uid,omitempty"`
}) *TypeResampleBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// The downsample function
// Possible enum values:
//   - `"sum"`
//   - `"mean"`
//   - `"min"`
//   - `"max"`
//   - `"count"`
//   - `"last"`
func (builder *TypeResampleBuilder) Downsampler(downsampler TypeResampleDownsampler) *TypeResampleBuilder {
	builder.internal.Downsampler = downsampler

	return builder
}

// The math expression
func (builder *TypeResampleBuilder) Expression(expression string) *TypeResampleBuilder {
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
func (builder *TypeResampleBuilder) Hide(hide bool) *TypeResampleBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Interval is the suggested duration between time points in a time series query.
// NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
// from the interval required to fill a pixels in the visualization
func (builder *TypeResampleBuilder) IntervalMs(intervalMs float64) *TypeResampleBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

// MaxDataPoints is the maximum number of data points that should be returned from a time series query.
// NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
// from the number of pixels visible in a visualization
func (builder *TypeResampleBuilder) MaxDataPoints(maxDataPoints int64) *TypeResampleBuilder {
	builder.internal.MaxDataPoints = &maxDataPoints

	return builder
}

// QueryType is an optional identifier for the type of query.
// It can be used to distinguish different types of queries.
func (builder *TypeResampleBuilder) QueryType(queryType string) *TypeResampleBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// RefID is the unique identifier of the query, set by the frontend call.
func (builder *TypeResampleBuilder) RefId(refId string) *TypeResampleBuilder {
	builder.internal.RefId = refId

	return builder
}

// Optionally define expected query result behavior
func (builder *TypeResampleBuilder) ResultAssertions(resultAssertions struct {
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
}) *TypeResampleBuilder {
	builder.internal.ResultAssertions = &resultAssertions

	return builder
}

// TimeRange represents the query range
// NOTE: unlike generic /ds/query, we can now send explicit time values in each query
// NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
func (builder *TypeResampleBuilder) TimeRange(timeRange struct {
	// From is the start time of the query.
	From string `json:"from"`
	// To is the end time of the query.
	To string `json:"to"`
}) *TypeResampleBuilder {
	builder.internal.TimeRange = &timeRange

	return builder
}

// The upsample function
// Possible enum values:
//   - `"pad"` Use the last seen value
//   - `"backfilling"` backfill
//   - `"fillna"` Do not fill values (nill)
func (builder *TypeResampleBuilder) Upsampler(upsampler TypeResampleUpsampler) *TypeResampleBuilder {
	builder.internal.Upsampler = upsampler

	return builder
}

// The time duration
func (builder *TypeResampleBuilder) Window(window string) *TypeResampleBuilder {
	if !(len([]rune(window)) >= 1) {
		builder.errors["window"] = cog.MakeBuildErrors("window", errors.New("len([]rune(window)) must be >= 1"))
		return builder
	}
	builder.internal.Window = window

	return builder
}

func (builder *TypeResampleBuilder) applyDefaults() {
}

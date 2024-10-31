// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeThresholdResultAssertions] = (*ExprTypeThresholdResultAssertionsBuilder)(nil)

type ExprTypeThresholdResultAssertionsBuilder struct {
	internal *ExprTypeThresholdResultAssertions
	errors   map[string]cog.BuildErrors
}

func NewExprTypeThresholdResultAssertionsBuilder() *ExprTypeThresholdResultAssertionsBuilder {
	resource := &ExprTypeThresholdResultAssertions{}
	builder := &ExprTypeThresholdResultAssertionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeThresholdResultAssertionsBuilder) Build() (ExprTypeThresholdResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeThresholdResultAssertions{}, err
	}

	return *builder.internal, nil
}

// Maximum frame count
func (builder *ExprTypeThresholdResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeThresholdResultAssertionsBuilder {
	builder.internal.MaxFrames = &maxFrames

	return builder
}

// Type asserts that the frame matches a known type structure.
// Possible enum values:
//   - `""`
//   - `"timeseries-wide"`
//   - `"timeseries-long"`
//   - `"timeseries-many"`
//   - `"timeseries-multi"`
//   - `"directory-listing"`
//   - `"table"`
//   - `"numeric-wide"`
//   - `"numeric-multi"`
//   - `"numeric-long"`
//   - `"log-lines"`
func (builder *ExprTypeThresholdResultAssertionsBuilder) Type(typeArg TypeThresholdType) *ExprTypeThresholdResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ExprTypeThresholdResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeThresholdResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

func (builder *ExprTypeThresholdResultAssertionsBuilder) applyDefaults() {
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeReduceResultAssertions] = (*ExprTypeReduceResultAssertionsBuilder)(nil)

type ExprTypeReduceResultAssertionsBuilder struct {
	internal *ExprTypeReduceResultAssertions
	errors   map[string]cog.BuildErrors
}

func NewExprTypeReduceResultAssertionsBuilder() *ExprTypeReduceResultAssertionsBuilder {
	resource := &ExprTypeReduceResultAssertions{}
	builder := &ExprTypeReduceResultAssertionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeReduceResultAssertionsBuilder) Build() (ExprTypeReduceResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeReduceResultAssertions{}, err
	}

	return *builder.internal, nil
}

// Maximum frame count
func (builder *ExprTypeReduceResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeReduceResultAssertionsBuilder {
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
func (builder *ExprTypeReduceResultAssertionsBuilder) Type(typeArg TypeReduceType) *ExprTypeReduceResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ExprTypeReduceResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeReduceResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

func (builder *ExprTypeReduceResultAssertionsBuilder) applyDefaults() {
}

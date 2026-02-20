// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeReduceResultAssertions] = (*ExprTypeReduceResultAssertionsBuilder)(nil)

type ExprTypeReduceResultAssertionsBuilder struct {
	internal *ExprTypeReduceResultAssertions
	errors   cog.BuildErrors
}

func NewExprTypeReduceResultAssertionsBuilder() *ExprTypeReduceResultAssertionsBuilder {
	resource := NewExprTypeReduceResultAssertions()
	builder := &ExprTypeReduceResultAssertionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeReduceResultAssertionsBuilder) Build() (ExprTypeReduceResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeReduceResultAssertions{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeReduceResultAssertions{}, cog.MakeBuildErrors("expr.exprTypeReduceResultAssertions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeReduceResultAssertionsBuilder) RecordError(path string, err error) *ExprTypeReduceResultAssertionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
func (builder *ExprTypeReduceResultAssertionsBuilder) Type(typeArg ExprTypeReduceResultAssertionsType) *ExprTypeReduceResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ExprTypeReduceResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeReduceResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

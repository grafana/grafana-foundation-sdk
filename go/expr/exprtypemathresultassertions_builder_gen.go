// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeMathResultAssertions] = (*ExprTypeMathResultAssertionsBuilder)(nil)

type ExprTypeMathResultAssertionsBuilder struct {
	internal *ExprTypeMathResultAssertions
	errors   cog.BuildErrors
}

func NewExprTypeMathResultAssertionsBuilder() *ExprTypeMathResultAssertionsBuilder {
	resource := NewExprTypeMathResultAssertions()
	builder := &ExprTypeMathResultAssertionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprTypeMathResultAssertionsBuilder) Build() (ExprTypeMathResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeMathResultAssertions{}, err
	}

	if len(builder.errors) > 0 {
		return ExprTypeMathResultAssertions{}, cog.MakeBuildErrors("expr.exprTypeMathResultAssertions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprTypeMathResultAssertionsBuilder) RecordError(path string, err error) *ExprTypeMathResultAssertionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Maximum frame count
func (builder *ExprTypeMathResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeMathResultAssertionsBuilder {
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
func (builder *ExprTypeMathResultAssertionsBuilder) Type(typeArg ExprTypeMathResultAssertionsType) *ExprTypeMathResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ExprTypeMathResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeMathResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

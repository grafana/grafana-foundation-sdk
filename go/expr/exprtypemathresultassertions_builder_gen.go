// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeMathResultAssertions] = (*ExprTypeMathResultAssertionsBuilder)(nil)

type ExprTypeMathResultAssertionsBuilder struct {
	internal *ExprTypeMathResultAssertions
	errors   map[string]cog.BuildErrors
}

func NewExprTypeMathResultAssertionsBuilder() *ExprTypeMathResultAssertionsBuilder {
	resource := &ExprTypeMathResultAssertions{}
	builder := &ExprTypeMathResultAssertionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeMathResultAssertionsBuilder) Build() (ExprTypeMathResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeMathResultAssertions{}, err
	}

	return *builder.internal, nil
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
func (builder *ExprTypeMathResultAssertionsBuilder) Type(typeArg TypeMathType) *ExprTypeMathResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ExprTypeMathResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeMathResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

func (builder *ExprTypeMathResultAssertionsBuilder) applyDefaults() {
}

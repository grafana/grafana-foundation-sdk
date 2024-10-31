// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeClassicConditionsResultAssertions] = (*ExprTypeClassicConditionsResultAssertionsBuilder)(nil)

type ExprTypeClassicConditionsResultAssertionsBuilder struct {
	internal *ExprTypeClassicConditionsResultAssertions
	errors   map[string]cog.BuildErrors
}

func NewExprTypeClassicConditionsResultAssertionsBuilder() *ExprTypeClassicConditionsResultAssertionsBuilder {
	resource := &ExprTypeClassicConditionsResultAssertions{}
	builder := &ExprTypeClassicConditionsResultAssertionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) Build() (ExprTypeClassicConditionsResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeClassicConditionsResultAssertions{}, err
	}

	return *builder.internal, nil
}

// Maximum frame count
func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeClassicConditionsResultAssertionsBuilder {
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
func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) Type(typeArg TypeClassicConditionsType) *ExprTypeClassicConditionsResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeClassicConditionsResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

func (builder *ExprTypeClassicConditionsResultAssertionsBuilder) applyDefaults() {
}

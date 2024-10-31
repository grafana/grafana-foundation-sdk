// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExprTypeSqlResultAssertions] = (*ExprTypeSqlResultAssertionsBuilder)(nil)

type ExprTypeSqlResultAssertionsBuilder struct {
	internal *ExprTypeSqlResultAssertions
	errors   map[string]cog.BuildErrors
}

func NewExprTypeSqlResultAssertionsBuilder() *ExprTypeSqlResultAssertionsBuilder {
	resource := &ExprTypeSqlResultAssertions{}
	builder := &ExprTypeSqlResultAssertionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExprTypeSqlResultAssertionsBuilder) Build() (ExprTypeSqlResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExprTypeSqlResultAssertions{}, err
	}

	return *builder.internal, nil
}

// Maximum frame count
func (builder *ExprTypeSqlResultAssertionsBuilder) MaxFrames(maxFrames int64) *ExprTypeSqlResultAssertionsBuilder {
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
func (builder *ExprTypeSqlResultAssertionsBuilder) Type(typeArg TypeSqlType) *ExprTypeSqlResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ExprTypeSqlResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ExprTypeSqlResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

func (builder *ExprTypeSqlResultAssertionsBuilder) applyDefaults() {
}

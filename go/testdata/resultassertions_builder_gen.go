// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package testdata

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ResultAssertions] = (*ResultAssertionsBuilder)(nil)

type ResultAssertionsBuilder struct {
	internal *ResultAssertions
	errors   cog.BuildErrors
}

func NewResultAssertionsBuilder() *ResultAssertionsBuilder {
	resource := NewResultAssertions()
	builder := &ResultAssertionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ResultAssertionsBuilder) Build() (ResultAssertions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ResultAssertions{}, err
	}

	if len(builder.errors) > 0 {
		return ResultAssertions{}, cog.MakeBuildErrors("testdata.resultAssertions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ResultAssertionsBuilder) RecordError(path string, err error) *ResultAssertionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Maximum frame count
func (builder *ResultAssertionsBuilder) MaxFrames(maxFrames int64) *ResultAssertionsBuilder {
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
func (builder *ResultAssertionsBuilder) Type(typeArg ResultAssertionsType) *ResultAssertionsBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// TypeVersion is the version of the Type property. Versions greater than 0.0 correspond to the dataplane
// contract documentation https://grafana.github.io/dataplane/contract/.
func (builder *ResultAssertionsBuilder) TypeVersion(typeVersion []int64) *ResultAssertionsBuilder {
	builder.internal.TypeVersion = typeVersion

	return builder
}

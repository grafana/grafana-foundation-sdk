// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ReduceDataOptions] = (*ReduceDataOptionsBuilder)(nil)

// TODO docs
type ReduceDataOptionsBuilder struct {
	internal *ReduceDataOptions
	errors   map[string]cog.BuildErrors
}

func NewReduceDataOptionsBuilder() *ReduceDataOptionsBuilder {
	resource := NewReduceDataOptions()
	builder := &ReduceDataOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ReduceDataOptionsBuilder) Build() (ReduceDataOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return ReduceDataOptions{}, err
	}

	return *builder.internal, nil
}

// If true show each row value
func (builder *ReduceDataOptionsBuilder) Values(values bool) *ReduceDataOptionsBuilder {
	builder.internal.Values = &values

	return builder
}

// if showing all values limit
func (builder *ReduceDataOptionsBuilder) Limit(limit float64) *ReduceDataOptionsBuilder {
	builder.internal.Limit = &limit

	return builder
}

// When !values, pick one value for the whole field
func (builder *ReduceDataOptionsBuilder) Calcs(calcs []string) *ReduceDataOptionsBuilder {
	builder.internal.Calcs = calcs

	return builder
}

// Which fields to show.  By default this is only numeric fields
func (builder *ReduceDataOptionsBuilder) Fields(fields string) *ReduceDataOptionsBuilder {
	builder.internal.Fields = &fields

	return builder
}

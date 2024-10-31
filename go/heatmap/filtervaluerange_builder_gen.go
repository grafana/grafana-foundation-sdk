// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[FilterValueRange] = (*FilterValueRangeBuilder)(nil)

// Controls the value filter range
type FilterValueRangeBuilder struct {
	internal *FilterValueRange
	errors   map[string]cog.BuildErrors
}

func NewFilterValueRangeBuilder() *FilterValueRangeBuilder {
	resource := &FilterValueRange{}
	builder := &FilterValueRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *FilterValueRangeBuilder) Build() (FilterValueRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return FilterValueRange{}, err
	}

	return *builder.internal, nil
}

// Sets the filter range to values less than or equal to the given value
func (builder *FilterValueRangeBuilder) Le(le float32) *FilterValueRangeBuilder {
	builder.internal.Le = &le

	return builder
}

// Sets the filter range to values greater than or equal to the given value
func (builder *FilterValueRangeBuilder) Ge(ge float32) *FilterValueRangeBuilder {
	builder.internal.Ge = &ge

	return builder
}

func (builder *FilterValueRangeBuilder) applyDefaults() {
}

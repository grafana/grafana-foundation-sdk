// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[FacetValue] = (*FacetValueBuilder)(nil)

// #FacetValue is a single value/count pair in a facet breakdown.
type FacetValueBuilder struct {
	internal *FacetValue
	errors   cog.BuildErrors
}

func NewFacetValueBuilder() *FacetValueBuilder {
	resource := NewFacetValue()
	builder := &FacetValueBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FacetValueBuilder) Build() (FacetValue, error) {
	if err := builder.internal.Validate(); err != nil {
		return FacetValue{}, err
	}

	if len(builder.errors) > 0 {
		return FacetValue{}, cog.MakeBuildErrors("alertingv0alpha1.facetValue", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FacetValueBuilder) RecordError(path string, err error) *FacetValueBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *FacetValueBuilder) Value(value string) *FacetValueBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *FacetValueBuilder) Count(count int64) *FacetValueBuilder {
	builder.internal.Count = count

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[OptionsWithLegend] = (*OptionsWithLegendBuilder)(nil)

// TODO docs
type OptionsWithLegendBuilder struct {
	internal *OptionsWithLegend
	errors   cog.BuildErrors
}

func NewOptionsWithLegendBuilder() *OptionsWithLegendBuilder {
	resource := NewOptionsWithLegend()
	builder := &OptionsWithLegendBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsWithLegendBuilder) Build() (OptionsWithLegend, error) {
	if err := builder.internal.Validate(); err != nil {
		return OptionsWithLegend{}, err
	}

	if len(builder.errors) > 0 {
		return OptionsWithLegend{}, cog.MakeBuildErrors("common.optionsWithLegend", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsWithLegendBuilder) Legend(legend cog.Builder[VizLegendOptions]) *OptionsWithLegendBuilder {
	legendResource, err := legend.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Legend = legendResource

	return builder
}

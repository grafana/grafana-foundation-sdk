// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[OptionsWithTooltip] = (*OptionsWithTooltipBuilder)(nil)

// TODO docs
type OptionsWithTooltipBuilder struct {
	internal *OptionsWithTooltip
	errors   cog.BuildErrors
}

func NewOptionsWithTooltipBuilder() *OptionsWithTooltipBuilder {
	resource := NewOptionsWithTooltip()
	builder := &OptionsWithTooltipBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsWithTooltipBuilder) Build() (OptionsWithTooltip, error) {
	if err := builder.internal.Validate(); err != nil {
		return OptionsWithTooltip{}, err
	}

	if len(builder.errors) > 0 {
		return OptionsWithTooltip{}, cog.MakeBuildErrors("common.optionsWithTooltip", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsWithTooltipBuilder) RecordError(path string, err error) *OptionsWithTooltipBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *OptionsWithTooltipBuilder) Tooltip(tooltip cog.Builder[VizTooltipOptions]) *OptionsWithTooltipBuilder {
	tooltipResource, err := tooltip.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Tooltip = tooltipResource

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VizTooltipOptions] = (*VizTooltipOptionsBuilder)(nil)

// TODO docs
type VizTooltipOptionsBuilder struct {
	internal *VizTooltipOptions
	errors   cog.BuildErrors
}

func NewVizTooltipOptionsBuilder() *VizTooltipOptionsBuilder {
	resource := NewVizTooltipOptions()
	builder := &VizTooltipOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *VizTooltipOptionsBuilder) Build() (VizTooltipOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return VizTooltipOptions{}, err
	}

	if len(builder.errors) > 0 {
		return VizTooltipOptions{}, cog.MakeBuildErrors("common.vizTooltipOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *VizTooltipOptionsBuilder) RecordError(path string, err error) *VizTooltipOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *VizTooltipOptionsBuilder) Mode(mode TooltipDisplayMode) *VizTooltipOptionsBuilder {
	builder.internal.Mode = mode

	return builder
}

func (builder *VizTooltipOptionsBuilder) Sort(sort SortOrder) *VizTooltipOptionsBuilder {
	builder.internal.Sort = sort

	return builder
}

func (builder *VizTooltipOptionsBuilder) MaxWidth(maxWidth float64) *VizTooltipOptionsBuilder {
	builder.internal.MaxWidth = &maxWidth

	return builder
}

func (builder *VizTooltipOptionsBuilder) MaxHeight(maxHeight float64) *VizTooltipOptionsBuilder {
	builder.internal.MaxHeight = &maxHeight

	return builder
}

func (builder *VizTooltipOptionsBuilder) HideZeros(hideZeros bool) *VizTooltipOptionsBuilder {
	builder.internal.HideZeros = &hideZeros

	return builder
}

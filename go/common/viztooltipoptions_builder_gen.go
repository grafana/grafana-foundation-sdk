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

func (builder *VizTooltipOptionsBuilder) Mode(mode TooltipDisplayMode) *VizTooltipOptionsBuilder {
	builder.internal.Mode = mode

	return builder
}

func (builder *VizTooltipOptionsBuilder) Sort(sort SortOrder) *VizTooltipOptionsBuilder {
	builder.internal.Sort = sort

	return builder
}

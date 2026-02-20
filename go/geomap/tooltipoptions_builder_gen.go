// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TooltipOptions] = (*TooltipOptionsBuilder)(nil)

type TooltipOptionsBuilder struct {
	internal *TooltipOptions
	errors   cog.BuildErrors
}

func NewTooltipOptionsBuilder() *TooltipOptionsBuilder {
	resource := NewTooltipOptions()
	builder := &TooltipOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TooltipOptionsBuilder) Build() (TooltipOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return TooltipOptions{}, err
	}

	if len(builder.errors) > 0 {
		return TooltipOptions{}, cog.MakeBuildErrors("geomap.tooltipOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TooltipOptionsBuilder) RecordError(path string, err error) *TooltipOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *TooltipOptionsBuilder) Mode(mode TooltipMode) *TooltipOptionsBuilder {
	builder.internal.Mode = mode

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package geomap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TooltipOptions] = (*TooltipOptionsBuilder)(nil)

type TooltipOptionsBuilder struct {
	internal *TooltipOptions
	errors   map[string]cog.BuildErrors
}

func NewTooltipOptionsBuilder() *TooltipOptionsBuilder {
	resource := &TooltipOptions{}
	builder := &TooltipOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TooltipOptionsBuilder) Build() (TooltipOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return TooltipOptions{}, err
	}

	return *builder.internal, nil
}

func (builder *TooltipOptionsBuilder) Mode(mode TooltipMode) *TooltipOptionsBuilder {
	builder.internal.Mode = mode

	return builder
}

func (builder *TooltipOptionsBuilder) applyDefaults() {
}

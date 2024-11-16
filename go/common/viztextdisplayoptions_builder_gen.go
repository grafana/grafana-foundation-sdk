// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VizTextDisplayOptions] = (*VizTextDisplayOptionsBuilder)(nil)

// TODO docs
type VizTextDisplayOptionsBuilder struct {
	internal *VizTextDisplayOptions
	errors   map[string]cog.BuildErrors
}

func NewVizTextDisplayOptionsBuilder() *VizTextDisplayOptionsBuilder {
	resource := NewVizTextDisplayOptions()
	builder := &VizTextDisplayOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *VizTextDisplayOptionsBuilder) Build() (VizTextDisplayOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return VizTextDisplayOptions{}, err
	}

	return *builder.internal, nil
}

// Explicit title text size
func (builder *VizTextDisplayOptionsBuilder) TitleSize(titleSize float64) *VizTextDisplayOptionsBuilder {
	builder.internal.TitleSize = &titleSize

	return builder
}

// Explicit value text size
func (builder *VizTextDisplayOptionsBuilder) ValueSize(valueSize float64) *VizTextDisplayOptionsBuilder {
	builder.internal.ValueSize = &valueSize

	return builder
}

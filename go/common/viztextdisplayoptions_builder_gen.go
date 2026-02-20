// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VizTextDisplayOptions] = (*VizTextDisplayOptionsBuilder)(nil)

// TODO docs
type VizTextDisplayOptionsBuilder struct {
	internal *VizTextDisplayOptions
	errors   cog.BuildErrors
}

func NewVizTextDisplayOptionsBuilder() *VizTextDisplayOptionsBuilder {
	resource := NewVizTextDisplayOptions()
	builder := &VizTextDisplayOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *VizTextDisplayOptionsBuilder) Build() (VizTextDisplayOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return VizTextDisplayOptions{}, err
	}

	if len(builder.errors) > 0 {
		return VizTextDisplayOptions{}, cog.MakeBuildErrors("common.vizTextDisplayOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *VizTextDisplayOptionsBuilder) RecordError(path string, err error) *VizTextDisplayOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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

// Explicit percent text size
func (builder *VizTextDisplayOptionsBuilder) PercentSize(percentSize float64) *VizTextDisplayOptionsBuilder {
	builder.internal.PercentSize = &percentSize

	return builder
}

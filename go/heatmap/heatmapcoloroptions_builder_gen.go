// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package heatmap

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HeatmapColorOptions] = (*HeatmapColorOptionsBuilder)(nil)

// Controls various color options
type HeatmapColorOptionsBuilder struct {
	internal *HeatmapColorOptions
	errors   map[string]cog.BuildErrors
}

func NewHeatmapColorOptionsBuilder() *HeatmapColorOptionsBuilder {
	resource := &HeatmapColorOptions{}
	builder := &HeatmapColorOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *HeatmapColorOptionsBuilder) Build() (HeatmapColorOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return HeatmapColorOptions{}, err
	}

	return *builder.internal, nil
}

// Sets the color mode
func (builder *HeatmapColorOptionsBuilder) Mode(mode HeatmapColorMode) *HeatmapColorOptionsBuilder {
	builder.internal.Mode = &mode

	return builder
}

// Controls the color scheme used
func (builder *HeatmapColorOptionsBuilder) Scheme(scheme string) *HeatmapColorOptionsBuilder {
	builder.internal.Scheme = scheme

	return builder
}

// Controls the color fill when in opacity mode
func (builder *HeatmapColorOptionsBuilder) Fill(fill string) *HeatmapColorOptionsBuilder {
	builder.internal.Fill = fill

	return builder
}

// Controls the color scale
func (builder *HeatmapColorOptionsBuilder) Scale(scale HeatmapColorScale) *HeatmapColorOptionsBuilder {
	builder.internal.Scale = &scale

	return builder
}

// Controls the exponent when scale is set to exponential
func (builder *HeatmapColorOptionsBuilder) Exponent(exponent float32) *HeatmapColorOptionsBuilder {
	builder.internal.Exponent = exponent

	return builder
}

// Controls the number of color steps
func (builder *HeatmapColorOptionsBuilder) Steps(steps uint64) *HeatmapColorOptionsBuilder {
	builder.internal.Steps = steps

	return builder
}

// Reverses the color scheme
func (builder *HeatmapColorOptionsBuilder) Reverse(reverse bool) *HeatmapColorOptionsBuilder {
	builder.internal.Reverse = reverse

	return builder
}

// Sets the minimum value for the color scale
func (builder *HeatmapColorOptionsBuilder) Min(min float32) *HeatmapColorOptionsBuilder {
	builder.internal.Min = &min

	return builder
}

// Sets the maximum value for the color scale
func (builder *HeatmapColorOptionsBuilder) Max(max float32) *HeatmapColorOptionsBuilder {
	builder.internal.Max = &max

	return builder
}

func (builder *HeatmapColorOptionsBuilder) applyDefaults() {
}

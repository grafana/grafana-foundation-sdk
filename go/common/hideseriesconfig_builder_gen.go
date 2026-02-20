// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HideSeriesConfig] = (*HideSeriesConfigBuilder)(nil)

// TODO docs
type HideSeriesConfigBuilder struct {
	internal *HideSeriesConfig
	errors   cog.BuildErrors
}

func NewHideSeriesConfigBuilder() *HideSeriesConfigBuilder {
	resource := NewHideSeriesConfig()
	builder := &HideSeriesConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *HideSeriesConfigBuilder) Build() (HideSeriesConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return HideSeriesConfig{}, err
	}

	if len(builder.errors) > 0 {
		return HideSeriesConfig{}, cog.MakeBuildErrors("common.hideSeriesConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *HideSeriesConfigBuilder) RecordError(path string, err error) *HideSeriesConfigBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *HideSeriesConfigBuilder) Tooltip(tooltip bool) *HideSeriesConfigBuilder {
	builder.internal.Tooltip = tooltip

	return builder
}

func (builder *HideSeriesConfigBuilder) Legend(legend bool) *HideSeriesConfigBuilder {
	builder.internal.Legend = legend

	return builder
}

func (builder *HideSeriesConfigBuilder) Viz(viz bool) *HideSeriesConfigBuilder {
	builder.internal.Viz = viz

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[FillConfig] = (*FillConfigBuilder)(nil)

// TODO docs
type FillConfigBuilder struct {
	internal *FillConfig
	errors   cog.BuildErrors
}

func NewFillConfigBuilder() *FillConfigBuilder {
	resource := NewFillConfig()
	builder := &FillConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FillConfigBuilder) Build() (FillConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return FillConfig{}, err
	}

	if len(builder.errors) > 0 {
		return FillConfig{}, cog.MakeBuildErrors("common.fillConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FillConfigBuilder) FillColor(fillColor string) *FillConfigBuilder {
	builder.internal.FillColor = &fillColor

	return builder
}

func (builder *FillConfigBuilder) FillOpacity(fillOpacity float64) *FillConfigBuilder {
	builder.internal.FillOpacity = &fillOpacity

	return builder
}

func (builder *FillConfigBuilder) FillBelowTo(fillBelowTo string) *FillConfigBuilder {
	builder.internal.FillBelowTo = &fillBelowTo

	return builder
}

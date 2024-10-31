// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Rate] = (*RateBuilder)(nil)

type RateBuilder struct {
	internal *Rate
	errors   map[string]cog.BuildErrors
}

func NewRateBuilder() *RateBuilder {
	resource := &Rate{}
	builder := &RateBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "rate"

	return builder
}

func (builder *RateBuilder) Build() (Rate, error) {
	if err := builder.internal.Validate(); err != nil {
		return Rate{}, err
	}

	return *builder.internal, nil
}

func (builder *RateBuilder) Field(field string) *RateBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *RateBuilder) Id(id string) *RateBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *RateBuilder) Settings(settings cog.Builder[ElasticsearchRateSettings]) *RateBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *RateBuilder) Hide(hide bool) *RateBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *RateBuilder) applyDefaults() {
}

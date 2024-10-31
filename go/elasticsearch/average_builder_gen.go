// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Average] = (*AverageBuilder)(nil)

type AverageBuilder struct {
	internal *Average
	errors   map[string]cog.BuildErrors
}

func NewAverageBuilder() *AverageBuilder {
	resource := &Average{}
	builder := &AverageBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "avg"

	return builder
}

func (builder *AverageBuilder) Build() (Average, error) {
	if err := builder.internal.Validate(); err != nil {
		return Average{}, err
	}

	return *builder.internal, nil
}

func (builder *AverageBuilder) Field(field string) *AverageBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *AverageBuilder) Id(id string) *AverageBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *AverageBuilder) Settings(settings cog.Builder[ElasticsearchAverageSettings]) *AverageBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *AverageBuilder) Hide(hide bool) *AverageBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *AverageBuilder) applyDefaults() {
}

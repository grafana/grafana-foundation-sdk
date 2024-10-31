// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Max] = (*MaxBuilder)(nil)

type MaxBuilder struct {
	internal *Max
	errors   map[string]cog.BuildErrors
}

func NewMaxBuilder() *MaxBuilder {
	resource := &Max{}
	builder := &MaxBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "max"

	return builder
}

func (builder *MaxBuilder) Build() (Max, error) {
	if err := builder.internal.Validate(); err != nil {
		return Max{}, err
	}

	return *builder.internal, nil
}

func (builder *MaxBuilder) Field(field string) *MaxBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *MaxBuilder) Id(id string) *MaxBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MaxBuilder) Settings(settings cog.Builder[ElasticsearchMaxSettings]) *MaxBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *MaxBuilder) Hide(hide bool) *MaxBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *MaxBuilder) applyDefaults() {
}

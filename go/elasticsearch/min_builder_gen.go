// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Min] = (*MinBuilder)(nil)

type MinBuilder struct {
	internal *Min
	errors   map[string]cog.BuildErrors
}

func NewMinBuilder() *MinBuilder {
	resource := &Min{}
	builder := &MinBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "min"

	return builder
}

func (builder *MinBuilder) Build() (Min, error) {
	if err := builder.internal.Validate(); err != nil {
		return Min{}, err
	}

	return *builder.internal, nil
}

func (builder *MinBuilder) Field(field string) *MinBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *MinBuilder) Id(id string) *MinBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MinBuilder) Settings(settings cog.Builder[ElasticsearchMinSettings]) *MinBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *MinBuilder) Hide(hide bool) *MinBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *MinBuilder) applyDefaults() {
}

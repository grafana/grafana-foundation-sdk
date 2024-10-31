// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Percentiles] = (*PercentilesBuilder)(nil)

type PercentilesBuilder struct {
	internal *Percentiles
	errors   map[string]cog.BuildErrors
}

func NewPercentilesBuilder() *PercentilesBuilder {
	resource := &Percentiles{}
	builder := &PercentilesBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "percentiles"

	return builder
}

func (builder *PercentilesBuilder) Build() (Percentiles, error) {
	if err := builder.internal.Validate(); err != nil {
		return Percentiles{}, err
	}

	return *builder.internal, nil
}

func (builder *PercentilesBuilder) Field(field string) *PercentilesBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *PercentilesBuilder) Id(id string) *PercentilesBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *PercentilesBuilder) Settings(settings cog.Builder[ElasticsearchPercentilesSettings]) *PercentilesBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *PercentilesBuilder) Hide(hide bool) *PercentilesBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *PercentilesBuilder) applyDefaults() {
}

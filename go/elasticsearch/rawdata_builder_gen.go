// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RawData] = (*RawDataBuilder)(nil)

type RawDataBuilder struct {
	internal *RawData
	errors   map[string]cog.BuildErrors
}

func NewRawDataBuilder() *RawDataBuilder {
	resource := &RawData{}
	builder := &RawDataBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "raw_data"

	return builder
}

func (builder *RawDataBuilder) Build() (RawData, error) {
	if err := builder.internal.Validate(); err != nil {
		return RawData{}, err
	}

	return *builder.internal, nil
}

func (builder *RawDataBuilder) Id(id string) *RawDataBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *RawDataBuilder) Settings(settings cog.Builder[ElasticsearchRawDataSettings]) *RawDataBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *RawDataBuilder) Hide(hide bool) *RawDataBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *RawDataBuilder) applyDefaults() {
}

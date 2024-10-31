// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SerialDiff] = (*SerialDiffBuilder)(nil)

type SerialDiffBuilder struct {
	internal *SerialDiff
	errors   map[string]cog.BuildErrors
}

func NewSerialDiffBuilder() *SerialDiffBuilder {
	resource := &SerialDiff{}
	builder := &SerialDiffBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "serial_diff"

	return builder
}

func (builder *SerialDiffBuilder) Build() (SerialDiff, error) {
	if err := builder.internal.Validate(); err != nil {
		return SerialDiff{}, err
	}

	return *builder.internal, nil
}

func (builder *SerialDiffBuilder) PipelineAgg(pipelineAgg string) *SerialDiffBuilder {
	builder.internal.PipelineAgg = &pipelineAgg

	return builder
}

func (builder *SerialDiffBuilder) Field(field string) *SerialDiffBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *SerialDiffBuilder) Id(id string) *SerialDiffBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *SerialDiffBuilder) Settings(settings cog.Builder[ElasticsearchSerialDiffSettings]) *SerialDiffBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *SerialDiffBuilder) Hide(hide bool) *SerialDiffBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *SerialDiffBuilder) applyDefaults() {
}

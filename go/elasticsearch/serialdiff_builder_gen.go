// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SerialDiff] = (*SerialDiffBuilder)(nil)

type SerialDiffBuilder struct {
	internal *SerialDiff
	errors   cog.BuildErrors
}

func NewSerialDiffBuilder() *SerialDiffBuilder {
	resource := NewSerialDiff()
	builder := &SerialDiffBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SerialDiffBuilder) Build() (SerialDiff, error) {
	if err := builder.internal.Validate(); err != nil {
		return SerialDiff{}, err
	}

	if len(builder.errors) > 0 {
		return SerialDiff{}, cog.MakeBuildErrors("elasticsearch.serialDiff", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SerialDiffBuilder) RecordError(path string, err error) *SerialDiffBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *SerialDiffBuilder) Hide(hide bool) *SerialDiffBuilder {
	builder.internal.Hide = &hide

	return builder
}

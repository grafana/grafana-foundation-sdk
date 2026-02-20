// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Average] = (*AverageBuilder)(nil)

type AverageBuilder struct {
	internal *Average
	errors   cog.BuildErrors
}

func NewAverageBuilder() *AverageBuilder {
	resource := NewAverage()
	builder := &AverageBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AverageBuilder) Build() (Average, error) {
	if err := builder.internal.Validate(); err != nil {
		return Average{}, err
	}

	if len(builder.errors) > 0 {
		return Average{}, cog.MakeBuildErrors("elasticsearch.average", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AverageBuilder) RecordError(path string, err error) *AverageBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *AverageBuilder) Hide(hide bool) *AverageBuilder {
	builder.internal.Hide = &hide

	return builder
}

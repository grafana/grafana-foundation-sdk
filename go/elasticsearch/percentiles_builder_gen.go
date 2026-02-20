// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Percentiles] = (*PercentilesBuilder)(nil)

type PercentilesBuilder struct {
	internal *Percentiles
	errors   cog.BuildErrors
}

func NewPercentilesBuilder() *PercentilesBuilder {
	resource := NewPercentiles()
	builder := &PercentilesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PercentilesBuilder) Build() (Percentiles, error) {
	if err := builder.internal.Validate(); err != nil {
		return Percentiles{}, err
	}

	if len(builder.errors) > 0 {
		return Percentiles{}, cog.MakeBuildErrors("elasticsearch.percentiles", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PercentilesBuilder) RecordError(path string, err error) *PercentilesBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *PercentilesBuilder) Hide(hide bool) *PercentilesBuilder {
	builder.internal.Hide = &hide

	return builder
}

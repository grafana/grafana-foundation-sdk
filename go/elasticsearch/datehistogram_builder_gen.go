// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DateHistogram] = (*DateHistogramBuilder)(nil)

type DateHistogramBuilder struct {
	internal *DateHistogram
	errors   cog.BuildErrors
}

func NewDateHistogramBuilder() *DateHistogramBuilder {
	resource := NewDateHistogram()
	builder := &DateHistogramBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DateHistogramBuilder) Build() (DateHistogram, error) {
	if err := builder.internal.Validate(); err != nil {
		return DateHistogram{}, err
	}

	if len(builder.errors) > 0 {
		return DateHistogram{}, cog.MakeBuildErrors("elasticsearch.dateHistogram", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DateHistogramBuilder) RecordError(path string, err error) *DateHistogramBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *DateHistogramBuilder) Field(field string) *DateHistogramBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *DateHistogramBuilder) Id(id string) *DateHistogramBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *DateHistogramBuilder) Settings(settings cog.Builder[ElasticsearchDateHistogramSettings]) *DateHistogramBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

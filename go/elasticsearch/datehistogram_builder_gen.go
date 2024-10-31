// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DateHistogram] = (*DateHistogramBuilder)(nil)

type DateHistogramBuilder struct {
	internal *DateHistogram
	errors   map[string]cog.BuildErrors
}

func NewDateHistogramBuilder() *DateHistogramBuilder {
	resource := &DateHistogram{}
	builder := &DateHistogramBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "date_histogram"

	return builder
}

func (builder *DateHistogramBuilder) Build() (DateHistogram, error) {
	if err := builder.internal.Validate(); err != nil {
		return DateHistogram{}, err
	}

	return *builder.internal, nil
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
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *DateHistogramBuilder) applyDefaults() {
}

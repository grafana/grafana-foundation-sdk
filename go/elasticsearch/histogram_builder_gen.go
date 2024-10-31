// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Histogram] = (*HistogramBuilder)(nil)

type HistogramBuilder struct {
	internal *Histogram
	errors   map[string]cog.BuildErrors
}

func NewHistogramBuilder() *HistogramBuilder {
	resource := &Histogram{}
	builder := &HistogramBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "histogram"

	return builder
}

func (builder *HistogramBuilder) Build() (Histogram, error) {
	if err := builder.internal.Validate(); err != nil {
		return Histogram{}, err
	}

	return *builder.internal, nil
}

func (builder *HistogramBuilder) Field(field string) *HistogramBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *HistogramBuilder) Id(id string) *HistogramBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *HistogramBuilder) Settings(settings cog.Builder[ElasticsearchHistogramSettings]) *HistogramBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *HistogramBuilder) applyDefaults() {
}

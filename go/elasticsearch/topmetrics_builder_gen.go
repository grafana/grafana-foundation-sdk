// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TopMetrics] = (*TopMetricsBuilder)(nil)

type TopMetricsBuilder struct {
	internal *TopMetrics
	errors   map[string]cog.BuildErrors
}

func NewTopMetricsBuilder() *TopMetricsBuilder {
	resource := &TopMetrics{}
	builder := &TopMetricsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "top_metrics"

	return builder
}

func (builder *TopMetricsBuilder) Build() (TopMetrics, error) {
	if err := builder.internal.Validate(); err != nil {
		return TopMetrics{}, err
	}

	return *builder.internal, nil
}

func (builder *TopMetricsBuilder) Id(id string) *TopMetricsBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *TopMetricsBuilder) Settings(settings cog.Builder[ElasticsearchTopMetricsSettings]) *TopMetricsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *TopMetricsBuilder) Hide(hide bool) *TopMetricsBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *TopMetricsBuilder) applyDefaults() {
}

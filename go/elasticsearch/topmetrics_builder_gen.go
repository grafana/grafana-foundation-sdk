// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TopMetrics] = (*TopMetricsBuilder)(nil)

type TopMetricsBuilder struct {
	internal *TopMetrics
	errors   cog.BuildErrors
}

func NewTopMetricsBuilder() *TopMetricsBuilder {
	resource := NewTopMetrics()
	builder := &TopMetricsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TopMetricsBuilder) Build() (TopMetrics, error) {
	if err := builder.internal.Validate(); err != nil {
		return TopMetrics{}, err
	}

	if len(builder.errors) > 0 {
		return TopMetrics{}, cog.MakeBuildErrors("elasticsearch.topMetrics", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TopMetricsBuilder) RecordError(path string, err error) *TopMetricsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *TopMetricsBuilder) Id(id string) *TopMetricsBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *TopMetricsBuilder) Settings(settings cog.Builder[ElasticsearchTopMetricsSettings]) *TopMetricsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *TopMetricsBuilder) Hide(hide bool) *TopMetricsBuilder {
	builder.internal.Hide = &hide

	return builder
}

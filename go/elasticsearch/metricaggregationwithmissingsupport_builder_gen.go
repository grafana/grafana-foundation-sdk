// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricAggregationWithMissingSupport] = (*MetricAggregationWithMissingSupportBuilder)(nil)

type MetricAggregationWithMissingSupportBuilder struct {
	internal *MetricAggregationWithMissingSupport
	errors   cog.BuildErrors
}

func NewMetricAggregationWithMissingSupportBuilder() *MetricAggregationWithMissingSupportBuilder {
	resource := NewMetricAggregationWithMissingSupport()
	builder := &MetricAggregationWithMissingSupportBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) Build() (MetricAggregationWithMissingSupport, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricAggregationWithMissingSupport{}, err
	}

	if len(builder.errors) > 0 {
		return MetricAggregationWithMissingSupport{}, cog.MakeBuildErrors("elasticsearch.metricAggregationWithMissingSupport", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetricAggregationWithMissingSupportBuilder) Settings(settings cog.Builder[ElasticsearchMetricAggregationWithMissingSupportSettings]) *MetricAggregationWithMissingSupportBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) Type(typeArg MetricAggregationType) *MetricAggregationWithMissingSupportBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) Id(id string) *MetricAggregationWithMissingSupportBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MetricAggregationWithMissingSupportBuilder) Hide(hide bool) *MetricAggregationWithMissingSupportBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricAggregationWithField] = (*MetricAggregationWithFieldBuilder)(nil)

type MetricAggregationWithFieldBuilder struct {
	internal *MetricAggregationWithField
	errors   cog.BuildErrors
}

func NewMetricAggregationWithFieldBuilder() *MetricAggregationWithFieldBuilder {
	resource := NewMetricAggregationWithField()
	builder := &MetricAggregationWithFieldBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MetricAggregationWithFieldBuilder) Build() (MetricAggregationWithField, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricAggregationWithField{}, err
	}

	if len(builder.errors) > 0 {
		return MetricAggregationWithField{}, cog.MakeBuildErrors("elasticsearch.metricAggregationWithField", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetricAggregationWithFieldBuilder) Field(field string) *MetricAggregationWithFieldBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *MetricAggregationWithFieldBuilder) Type(typeArg MetricAggregationType) *MetricAggregationWithFieldBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *MetricAggregationWithFieldBuilder) Id(id string) *MetricAggregationWithFieldBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MetricAggregationWithFieldBuilder) Hide(hide bool) *MetricAggregationWithFieldBuilder {
	builder.internal.Hide = &hide

	return builder
}

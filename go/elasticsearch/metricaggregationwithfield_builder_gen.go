// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricAggregationWithField] = (*MetricAggregationWithFieldBuilder)(nil)

type MetricAggregationWithFieldBuilder struct {
	internal *MetricAggregationWithField
	errors   map[string]cog.BuildErrors
}

func NewMetricAggregationWithFieldBuilder() *MetricAggregationWithFieldBuilder {
	resource := &MetricAggregationWithField{}
	builder := &MetricAggregationWithFieldBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MetricAggregationWithFieldBuilder) Build() (MetricAggregationWithField, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricAggregationWithField{}, err
	}

	return *builder.internal, nil
}

func (builder *MetricAggregationWithFieldBuilder) Field(field string) *MetricAggregationWithFieldBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *MetricAggregationWithFieldBuilder) Type(typeArg cog.Builder[MetricAggregationType]) *MetricAggregationWithFieldBuilder {
	typeArgResource, err := typeArg.Build()
	if err != nil {
		builder.errors["type"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Type = typeArgResource

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

func (builder *MetricAggregationWithFieldBuilder) applyDefaults() {
}

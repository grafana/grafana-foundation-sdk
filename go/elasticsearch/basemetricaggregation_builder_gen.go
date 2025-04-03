// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseMetricAggregation] = (*BaseMetricAggregationBuilder)(nil)

type BaseMetricAggregationBuilder struct {
	internal *BaseMetricAggregation
	errors   map[string]cog.BuildErrors
}

func NewBaseMetricAggregationBuilder() *BaseMetricAggregationBuilder {
	resource := NewBaseMetricAggregation()
	builder := &BaseMetricAggregationBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *BaseMetricAggregationBuilder) Build() (BaseMetricAggregation, error) {
	if err := builder.internal.Validate(); err != nil {
		return BaseMetricAggregation{}, err
	}

	return *builder.internal, nil
}

func (builder *BaseMetricAggregationBuilder) Type(typeArg MetricAggregationType) *BaseMetricAggregationBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *BaseMetricAggregationBuilder) Id(id string) *BaseMetricAggregationBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *BaseMetricAggregationBuilder) Hide(hide bool) *BaseMetricAggregationBuilder {
	builder.internal.Hide = &hide

	return builder
}

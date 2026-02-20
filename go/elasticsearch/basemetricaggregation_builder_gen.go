// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BaseMetricAggregation] = (*BaseMetricAggregationBuilder)(nil)

type BaseMetricAggregationBuilder struct {
	internal *BaseMetricAggregation
	errors   cog.BuildErrors
}

func NewBaseMetricAggregationBuilder() *BaseMetricAggregationBuilder {
	resource := NewBaseMetricAggregation()
	builder := &BaseMetricAggregationBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BaseMetricAggregationBuilder) Build() (BaseMetricAggregation, error) {
	if err := builder.internal.Validate(); err != nil {
		return BaseMetricAggregation{}, err
	}

	if len(builder.errors) > 0 {
		return BaseMetricAggregation{}, cog.MakeBuildErrors("elasticsearch.baseMetricAggregation", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BaseMetricAggregationBuilder) RecordError(path string, err error) *BaseMetricAggregationBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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

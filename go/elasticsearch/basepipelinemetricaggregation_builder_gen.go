// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BasePipelineMetricAggregation] = (*BasePipelineMetricAggregationBuilder)(nil)

type BasePipelineMetricAggregationBuilder struct {
	internal *BasePipelineMetricAggregation
	errors   map[string]cog.BuildErrors
}

func NewBasePipelineMetricAggregationBuilder() *BasePipelineMetricAggregationBuilder {
	resource := &BasePipelineMetricAggregation{}
	builder := &BasePipelineMetricAggregationBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *BasePipelineMetricAggregationBuilder) Build() (BasePipelineMetricAggregation, error) {
	if err := builder.internal.Validate(); err != nil {
		return BasePipelineMetricAggregation{}, err
	}

	return *builder.internal, nil
}

func (builder *BasePipelineMetricAggregationBuilder) PipelineAgg(pipelineAgg string) *BasePipelineMetricAggregationBuilder {
	builder.internal.PipelineAgg = &pipelineAgg

	return builder
}

func (builder *BasePipelineMetricAggregationBuilder) Field(field string) *BasePipelineMetricAggregationBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *BasePipelineMetricAggregationBuilder) Type(typeArg string) *BasePipelineMetricAggregationBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *BasePipelineMetricAggregationBuilder) Id(id string) *BasePipelineMetricAggregationBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *BasePipelineMetricAggregationBuilder) Hide(hide bool) *BasePipelineMetricAggregationBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *BasePipelineMetricAggregationBuilder) applyDefaults() {
}

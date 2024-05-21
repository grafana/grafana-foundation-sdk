// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StringOrPipelineMetricAggregationType] = (*StringOrPipelineMetricAggregationTypeBuilder)(nil)

type StringOrPipelineMetricAggregationTypeBuilder struct {
	internal *StringOrPipelineMetricAggregationType
	errors   map[string]cog.BuildErrors
}

func NewStringOrPipelineMetricAggregationTypeBuilder() *StringOrPipelineMetricAggregationTypeBuilder {
	resource := &StringOrPipelineMetricAggregationType{}
	builder := &StringOrPipelineMetricAggregationTypeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	valString := "count"
	builder.internal.String = &valString

	return builder
}

func (builder *StringOrPipelineMetricAggregationTypeBuilder) Build() (StringOrPipelineMetricAggregationType, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("StringOrPipelineMetricAggregationType", err)...)
	}

	if len(errs) != 0 {
		return StringOrPipelineMetricAggregationType{}, errs
	}

	return *builder.internal, nil
}

func (builder *StringOrPipelineMetricAggregationTypeBuilder) PipelineMetricAggregationType(pipelineMetricAggregationType PipelineMetricAggregationType) *StringOrPipelineMetricAggregationTypeBuilder {
	builder.internal.PipelineMetricAggregationType = &pipelineMetricAggregationType

	return builder
}

func (builder *StringOrPipelineMetricAggregationTypeBuilder) applyDefaults() {
}

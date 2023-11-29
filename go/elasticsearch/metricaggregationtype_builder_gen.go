// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricAggregationType] = (*MetricAggregationTypeBuilder)(nil)

type MetricAggregationTypeBuilder struct {
	internal *MetricAggregationType
	errors   map[string]cog.BuildErrors
}

func NewMetricAggregationTypeBuilder() *MetricAggregationTypeBuilder {
	resource := &MetricAggregationType{}
	builder := &MetricAggregationTypeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	valString := "count"
	builder.internal.String = &valString

	return builder
}

func (builder *MetricAggregationTypeBuilder) Build() (MetricAggregationType, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("MetricAggregationType", err)...)
	}

	if len(errs) != 0 {
		return MetricAggregationType{}, errs
	}

	return *builder.internal, nil
}

func (builder *MetricAggregationTypeBuilder) PipelineMetricAggregationType(pipelineMetricAggregationType PipelineMetricAggregationType) *MetricAggregationTypeBuilder {
	builder.internal.PipelineMetricAggregationType = &pipelineMetricAggregationType

	return builder
}

func (builder *MetricAggregationTypeBuilder) applyDefaults() {
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PipelineMetricAggregationWithMultipleBucketPaths] = (*PipelineMetricAggregationWithMultipleBucketPathsBuilder)(nil)

type PipelineMetricAggregationWithMultipleBucketPathsBuilder struct {
	internal *PipelineMetricAggregationWithMultipleBucketPaths
	errors   cog.BuildErrors
}

func NewPipelineMetricAggregationWithMultipleBucketPathsBuilder() *PipelineMetricAggregationWithMultipleBucketPathsBuilder {
	resource := NewPipelineMetricAggregationWithMultipleBucketPaths()
	builder := &PipelineMetricAggregationWithMultipleBucketPathsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Build() (PipelineMetricAggregationWithMultipleBucketPaths, error) {
	if err := builder.internal.Validate(); err != nil {
		return PipelineMetricAggregationWithMultipleBucketPaths{}, err
	}

	if len(builder.errors) > 0 {
		return PipelineMetricAggregationWithMultipleBucketPaths{}, cog.MakeBuildErrors("elasticsearch.pipelineMetricAggregationWithMultipleBucketPaths", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) PipelineVariables(pipelineVariables []cog.Builder[PipelineVariable]) *PipelineMetricAggregationWithMultipleBucketPathsBuilder {
	pipelineVariablesResources := make([]PipelineVariable, 0, len(pipelineVariables))
	for _, r1 := range pipelineVariables {
		pipelineVariablesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		pipelineVariablesResources = append(pipelineVariablesResources, pipelineVariablesDepth1)
	}
	builder.internal.PipelineVariables = pipelineVariablesResources

	return builder
}

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Type(typeArg MetricAggregationType) *PipelineMetricAggregationWithMultipleBucketPathsBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Id(id string) *PipelineMetricAggregationWithMultipleBucketPathsBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Hide(hide bool) *PipelineMetricAggregationWithMultipleBucketPathsBuilder {
	builder.internal.Hide = &hide

	return builder
}

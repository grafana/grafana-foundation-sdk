// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PipelineMetricAggregationWithMultipleBucketPaths] = (*PipelineMetricAggregationWithMultipleBucketPathsBuilder)(nil)

type PipelineMetricAggregationWithMultipleBucketPathsBuilder struct {
	internal *PipelineMetricAggregationWithMultipleBucketPaths
	errors   map[string]cog.BuildErrors
}

func NewPipelineMetricAggregationWithMultipleBucketPathsBuilder() *PipelineMetricAggregationWithMultipleBucketPathsBuilder {
	resource := &PipelineMetricAggregationWithMultipleBucketPaths{}
	builder := &PipelineMetricAggregationWithMultipleBucketPathsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Build() (PipelineMetricAggregationWithMultipleBucketPaths, error) {
	if err := builder.internal.Validate(); err != nil {
		return PipelineMetricAggregationWithMultipleBucketPaths{}, err
	}

	return *builder.internal, nil
}

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) PipelineVariables(pipelineVariables []cog.Builder[PipelineVariable]) *PipelineMetricAggregationWithMultipleBucketPathsBuilder {
	pipelineVariablesResources := make([]PipelineVariable, 0, len(pipelineVariables))
	for _, r1 := range pipelineVariables {
		pipelineVariablesDepth1, err := r1.Build()
		if err != nil {
			builder.errors["pipelineVariables"] = err.(cog.BuildErrors)
			return builder
		}
		pipelineVariablesResources = append(pipelineVariablesResources, pipelineVariablesDepth1)
	}
	builder.internal.PipelineVariables = pipelineVariablesResources

	return builder
}

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) Type(typeArg cog.Builder[MetricAggregationType]) *PipelineMetricAggregationWithMultipleBucketPathsBuilder {
	typeArgResource, err := typeArg.Build()
	if err != nil {
		builder.errors["type"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Type = typeArgResource

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

func (builder *PipelineMetricAggregationWithMultipleBucketPathsBuilder) applyDefaults() {
}

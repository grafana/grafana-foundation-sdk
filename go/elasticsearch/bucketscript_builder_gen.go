// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BucketScript] = (*BucketScriptBuilder)(nil)

type BucketScriptBuilder struct {
	internal *BucketScript
	errors   cog.BuildErrors
}

func NewBucketScriptBuilder() *BucketScriptBuilder {
	resource := NewBucketScript()
	builder := &BucketScriptBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BucketScriptBuilder) Build() (BucketScript, error) {
	if err := builder.internal.Validate(); err != nil {
		return BucketScript{}, err
	}

	if len(builder.errors) > 0 {
		return BucketScript{}, cog.MakeBuildErrors("elasticsearch.bucketScript", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BucketScriptBuilder) RecordError(path string, err error) *BucketScriptBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *BucketScriptBuilder) PipelineVariables(pipelineVariables []cog.Builder[PipelineVariable]) *BucketScriptBuilder {
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

func (builder *BucketScriptBuilder) Id(id string) *BucketScriptBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *BucketScriptBuilder) Settings(settings cog.Builder[ElasticsearchBucketScriptSettings]) *BucketScriptBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *BucketScriptBuilder) Hide(hide bool) *BucketScriptBuilder {
	builder.internal.Hide = &hide

	return builder
}

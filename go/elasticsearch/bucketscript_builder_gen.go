// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BucketScript] = (*BucketScriptBuilder)(nil)

type BucketScriptBuilder struct {
	internal *BucketScript
	errors   map[string]cog.BuildErrors
}

func NewBucketScriptBuilder() *BucketScriptBuilder {
	resource := &BucketScript{}
	builder := &BucketScriptBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "bucket_script"

	return builder
}

func (builder *BucketScriptBuilder) Build() (BucketScript, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("BucketScript", err)...)
	}

	if len(errs) != 0 {
		return BucketScript{}, errs
	}

	return *builder.internal, nil
}

func (builder *BucketScriptBuilder) PipelineVariables(pipelineVariables []cog.Builder[PipelineVariable]) *BucketScriptBuilder {
	pipelineVariablesResources := make([]PipelineVariable, 0, len(pipelineVariables))
	for _, r := range pipelineVariables {
		pipelineVariablesResource, err := r.Build()
		if err != nil {
			builder.errors["pipelineVariables"] = err.(cog.BuildErrors)
			return builder
		}
		pipelineVariablesResources = append(pipelineVariablesResources, pipelineVariablesResource)
	}
	builder.internal.PipelineVariables = pipelineVariablesResources

	return builder
}

func (builder *BucketScriptBuilder) Id(id string) *BucketScriptBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *BucketScriptBuilder) Settings(settings struct {
	Script *InlineScript `json:"script,omitempty"`
}) *BucketScriptBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *BucketScriptBuilder) Hide(hide bool) *BucketScriptBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *BucketScriptBuilder) applyDefaults() {
}

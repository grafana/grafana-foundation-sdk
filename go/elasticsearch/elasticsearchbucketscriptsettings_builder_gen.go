// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchBucketScriptSettings] = (*ElasticsearchBucketScriptSettingsBuilder)(nil)

type ElasticsearchBucketScriptSettingsBuilder struct {
	internal *ElasticsearchBucketScriptSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchBucketScriptSettingsBuilder() *ElasticsearchBucketScriptSettingsBuilder {
	resource := &ElasticsearchBucketScriptSettings{}
	builder := &ElasticsearchBucketScriptSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchBucketScriptSettingsBuilder) Build() (ElasticsearchBucketScriptSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchBucketScriptSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchBucketScriptSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchBucketScriptSettingsBuilder) Script(script cog.Builder[InlineScript]) *ElasticsearchBucketScriptSettingsBuilder {
	scriptResource, err := script.Build()
	if err != nil {
		builder.errors["script"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Script = &scriptResource

	return builder
}

func (builder *ElasticsearchBucketScriptSettingsBuilder) applyDefaults() {
}

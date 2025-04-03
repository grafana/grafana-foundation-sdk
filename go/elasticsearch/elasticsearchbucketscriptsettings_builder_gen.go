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
	resource := NewElasticsearchBucketScriptSettings()
	builder := &ElasticsearchBucketScriptSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ElasticsearchBucketScriptSettingsBuilder) Build() (ElasticsearchBucketScriptSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchBucketScriptSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchBucketScriptSettingsBuilder) Script(script InlineScript) *ElasticsearchBucketScriptSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

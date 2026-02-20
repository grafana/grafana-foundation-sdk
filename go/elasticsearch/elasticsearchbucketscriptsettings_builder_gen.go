// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchBucketScriptSettings] = (*ElasticsearchBucketScriptSettingsBuilder)(nil)

type ElasticsearchBucketScriptSettingsBuilder struct {
	internal *ElasticsearchBucketScriptSettings
	errors   cog.BuildErrors
}

func NewElasticsearchBucketScriptSettingsBuilder() *ElasticsearchBucketScriptSettingsBuilder {
	resource := NewElasticsearchBucketScriptSettings()
	builder := &ElasticsearchBucketScriptSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchBucketScriptSettingsBuilder) Build() (ElasticsearchBucketScriptSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchBucketScriptSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchBucketScriptSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchBucketScriptSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchBucketScriptSettingsBuilder) RecordError(path string, err error) *ElasticsearchBucketScriptSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchBucketScriptSettingsBuilder) Script(script InlineScript) *ElasticsearchBucketScriptSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

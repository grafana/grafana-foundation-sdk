// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMetricAggregationWithInlineScriptSettings] = (*ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder)(nil)

type ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder struct {
	internal *ElasticsearchMetricAggregationWithInlineScriptSettings
	errors   cog.BuildErrors
}

func NewElasticsearchMetricAggregationWithInlineScriptSettingsBuilder() *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder {
	resource := NewElasticsearchMetricAggregationWithInlineScriptSettings()
	builder := &ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) Build() (ElasticsearchMetricAggregationWithInlineScriptSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMetricAggregationWithInlineScriptSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchMetricAggregationWithInlineScriptSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchMetricAggregationWithInlineScriptSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) RecordError(path string, err error) *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) Script(script InlineScript) *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

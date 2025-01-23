// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMetricAggregationWithInlineScriptSettings] = (*ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder)(nil)

type ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder struct {
	internal *ElasticsearchMetricAggregationWithInlineScriptSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchMetricAggregationWithInlineScriptSettingsBuilder() *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder {
	resource := NewElasticsearchMetricAggregationWithInlineScriptSettings()
	builder := &ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) Build() (ElasticsearchMetricAggregationWithInlineScriptSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMetricAggregationWithInlineScriptSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) Script(script InlineScript) *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder {
	builder.internal.Script = &script

	return builder
}

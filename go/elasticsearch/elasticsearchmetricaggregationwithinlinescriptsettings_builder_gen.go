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
	resource := &ElasticsearchMetricAggregationWithInlineScriptSettings{}
	builder := &ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) Build() (ElasticsearchMetricAggregationWithInlineScriptSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchMetricAggregationWithInlineScriptSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchMetricAggregationWithInlineScriptSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) Script(script cog.Builder[InlineScript]) *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder {
	scriptResource, err := script.Build()
	if err != nil {
		builder.errors["script"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Script = &scriptResource

	return builder
}

func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) applyDefaults() {
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricAggregationWithInlineScript] = (*MetricAggregationWithInlineScriptBuilder)(nil)

type MetricAggregationWithInlineScriptBuilder struct {
	internal *MetricAggregationWithInlineScript
	errors   map[string]cog.BuildErrors
}

func NewMetricAggregationWithInlineScriptBuilder() *MetricAggregationWithInlineScriptBuilder {
	resource := NewMetricAggregationWithInlineScript()
	builder := &MetricAggregationWithInlineScriptBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *MetricAggregationWithInlineScriptBuilder) Build() (MetricAggregationWithInlineScript, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricAggregationWithInlineScript{}, err
	}

	return *builder.internal, nil
}

func (builder *MetricAggregationWithInlineScriptBuilder) Settings(settings cog.Builder[ElasticsearchMetricAggregationWithInlineScriptSettings]) *MetricAggregationWithInlineScriptBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *MetricAggregationWithInlineScriptBuilder) Type(typeArg MetricAggregationType) *MetricAggregationWithInlineScriptBuilder {
	builder.internal.Type = typeArg

	return builder
}

func (builder *MetricAggregationWithInlineScriptBuilder) Id(id string) *MetricAggregationWithInlineScriptBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *MetricAggregationWithInlineScriptBuilder) Hide(hide bool) *MetricAggregationWithInlineScriptBuilder {
	builder.internal.Hide = &hide

	return builder
}

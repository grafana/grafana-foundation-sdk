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
	resource := &MetricAggregationWithInlineScript{}
	builder := &MetricAggregationWithInlineScriptBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MetricAggregationWithInlineScriptBuilder) Build() (MetricAggregationWithInlineScript, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("MetricAggregationWithInlineScript", err)...)
	}

	if len(errs) != 0 {
		return MetricAggregationWithInlineScript{}, errs
	}

	return *builder.internal, nil
}

func (builder *MetricAggregationWithInlineScriptBuilder) Settings(settings struct {
	Script *InlineScript `json:"script,omitempty"`
}) *MetricAggregationWithInlineScriptBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *MetricAggregationWithInlineScriptBuilder) Type(typeArg cog.Builder[MetricAggregationType]) *MetricAggregationWithInlineScriptBuilder {
	typeArgResource, err := typeArg.Build()
	if err != nil {
		builder.errors["type"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Type = typeArgResource

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

func (builder *MetricAggregationWithInlineScriptBuilder) applyDefaults() {
}

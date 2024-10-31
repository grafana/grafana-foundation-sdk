// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMetricAggregationWithMissingSupportSettings] = (*ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder)(nil)

type ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder struct {
	internal *ElasticsearchMetricAggregationWithMissingSupportSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchMetricAggregationWithMissingSupportSettingsBuilder() *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder {
	resource := &ElasticsearchMetricAggregationWithMissingSupportSettings{}
	builder := &ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder) Build() (ElasticsearchMetricAggregationWithMissingSupportSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMetricAggregationWithMissingSupportSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder) Missing(missing string) *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder) applyDefaults() {
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMetricAggregationWithMissingSupportSettings] = (*ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder)(nil)

type ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder struct {
	internal *ElasticsearchMetricAggregationWithMissingSupportSettings
	errors   cog.BuildErrors
}

func NewElasticsearchMetricAggregationWithMissingSupportSettingsBuilder() *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder {
	resource := NewElasticsearchMetricAggregationWithMissingSupportSettings()
	builder := &ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder) Build() (ElasticsearchMetricAggregationWithMissingSupportSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMetricAggregationWithMissingSupportSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchMetricAggregationWithMissingSupportSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchMetricAggregationWithMissingSupportSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder) RecordError(path string, err error) *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder) Missing(missing string) *ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

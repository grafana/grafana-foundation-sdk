// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchTopMetricsSettings] = (*ElasticsearchTopMetricsSettingsBuilder)(nil)

type ElasticsearchTopMetricsSettingsBuilder struct {
	internal *ElasticsearchTopMetricsSettings
	errors   cog.BuildErrors
}

func NewElasticsearchTopMetricsSettingsBuilder() *ElasticsearchTopMetricsSettingsBuilder {
	resource := NewElasticsearchTopMetricsSettings()
	builder := &ElasticsearchTopMetricsSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchTopMetricsSettingsBuilder) Build() (ElasticsearchTopMetricsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchTopMetricsSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchTopMetricsSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchTopMetricsSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchTopMetricsSettingsBuilder) Order(order string) *ElasticsearchTopMetricsSettingsBuilder {
	builder.internal.Order = &order

	return builder
}

func (builder *ElasticsearchTopMetricsSettingsBuilder) OrderBy(orderBy string) *ElasticsearchTopMetricsSettingsBuilder {
	builder.internal.OrderBy = &orderBy

	return builder
}

func (builder *ElasticsearchTopMetricsSettingsBuilder) Metrics(metrics []string) *ElasticsearchTopMetricsSettingsBuilder {
	builder.internal.Metrics = metrics

	return builder
}

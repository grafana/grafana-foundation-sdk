// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchTopMetricsSettings] = (*ElasticsearchTopMetricsSettingsBuilder)(nil)

type ElasticsearchTopMetricsSettingsBuilder struct {
	internal *ElasticsearchTopMetricsSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchTopMetricsSettingsBuilder() *ElasticsearchTopMetricsSettingsBuilder {
	resource := &ElasticsearchTopMetricsSettings{}
	builder := &ElasticsearchTopMetricsSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchTopMetricsSettingsBuilder) Build() (ElasticsearchTopMetricsSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchTopMetricsSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchTopMetricsSettings{}, errs
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

func (builder *ElasticsearchTopMetricsSettingsBuilder) applyDefaults() {
}

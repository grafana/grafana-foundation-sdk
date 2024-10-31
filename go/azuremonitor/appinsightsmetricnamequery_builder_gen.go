// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AppInsightsMetricNameQuery] = (*AppInsightsMetricNameQueryBuilder)(nil)

type AppInsightsMetricNameQueryBuilder struct {
	internal *AppInsightsMetricNameQuery
	errors   map[string]cog.BuildErrors
}

func NewAppInsightsMetricNameQueryBuilder() *AppInsightsMetricNameQueryBuilder {
	resource := &AppInsightsMetricNameQuery{}
	builder := &AppInsightsMetricNameQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "AppInsightsMetricNameQuery"

	return builder
}

func (builder *AppInsightsMetricNameQueryBuilder) Build() (AppInsightsMetricNameQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AppInsightsMetricNameQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *AppInsightsMetricNameQueryBuilder) RawQuery(rawQuery string) *AppInsightsMetricNameQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *AppInsightsMetricNameQueryBuilder) applyDefaults() {
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AppInsightsGroupByQuery] = (*AppInsightsGroupByQueryBuilder)(nil)

type AppInsightsGroupByQueryBuilder struct {
	internal *AppInsightsGroupByQuery
	errors   map[string]cog.BuildErrors
}

func NewAppInsightsGroupByQueryBuilder() *AppInsightsGroupByQueryBuilder {
	resource := &AppInsightsGroupByQuery{}
	builder := &AppInsightsGroupByQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Kind = "AppInsightsGroupByQuery"

	return builder
}

func (builder *AppInsightsGroupByQueryBuilder) Build() (AppInsightsGroupByQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AppInsightsGroupByQuery{}, err
	}

	return *builder.internal, nil
}

func (builder *AppInsightsGroupByQueryBuilder) RawQuery(rawQuery string) *AppInsightsGroupByQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *AppInsightsGroupByQueryBuilder) MetricName(metricName string) *AppInsightsGroupByQueryBuilder {
	builder.internal.MetricName = metricName

	return builder
}

func (builder *AppInsightsGroupByQueryBuilder) applyDefaults() {
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AppInsightsMetricNameQuery] = (*AppInsightsMetricNameQueryBuilder)(nil)

type AppInsightsMetricNameQueryBuilder struct {
	internal *AppInsightsMetricNameQuery
	errors   cog.BuildErrors
}

func NewAppInsightsMetricNameQueryBuilder() *AppInsightsMetricNameQueryBuilder {
	resource := NewAppInsightsMetricNameQuery()
	builder := &AppInsightsMetricNameQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "AppInsightsMetricNameQuery"

	return builder
}

func (builder *AppInsightsMetricNameQueryBuilder) Build() (AppInsightsMetricNameQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AppInsightsMetricNameQuery{}, err
	}

	if len(builder.errors) > 0 {
		return AppInsightsMetricNameQuery{}, cog.MakeBuildErrors("azuremonitor.appInsightsMetricNameQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AppInsightsMetricNameQueryBuilder) RecordError(path string, err error) *AppInsightsMetricNameQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AppInsightsMetricNameQueryBuilder) RawQuery(rawQuery string) *AppInsightsMetricNameQueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

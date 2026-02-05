// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchMovingAverageHoltWintersModelSettingsSettings] = (*ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder)(nil)

type ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder struct {
	internal *ElasticsearchMovingAverageHoltWintersModelSettingsSettings
	errors   cog.BuildErrors
}

func NewElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder() *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder {
	resource := NewElasticsearchMovingAverageHoltWintersModelSettingsSettings()
	builder := &ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder) Build() (ElasticsearchMovingAverageHoltWintersModelSettingsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchMovingAverageHoltWintersModelSettingsSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchMovingAverageHoltWintersModelSettingsSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchMovingAverageHoltWintersModelSettingsSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder) Alpha(alpha string) *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder {
	builder.internal.Alpha = &alpha

	return builder
}

func (builder *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder) Beta(beta string) *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder {
	builder.internal.Beta = &beta

	return builder
}

func (builder *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder) Gamma(gamma string) *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder {
	builder.internal.Gamma = &gamma

	return builder
}

func (builder *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder) Period(period string) *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder {
	builder.internal.Period = &period

	return builder
}

func (builder *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder) Pad(pad bool) *ElasticsearchMovingAverageHoltWintersModelSettingsSettingsBuilder {
	builder.internal.Pad = &pad

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchHistogramSettings] = (*ElasticsearchHistogramSettingsBuilder)(nil)

type ElasticsearchHistogramSettingsBuilder struct {
	internal *ElasticsearchHistogramSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchHistogramSettingsBuilder() *ElasticsearchHistogramSettingsBuilder {
	resource := &ElasticsearchHistogramSettings{}
	builder := &ElasticsearchHistogramSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchHistogramSettingsBuilder) Build() (ElasticsearchHistogramSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchHistogramSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchHistogramSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchHistogramSettingsBuilder) Interval(interval string) *ElasticsearchHistogramSettingsBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *ElasticsearchHistogramSettingsBuilder) MinDocCount(minDocCount string) *ElasticsearchHistogramSettingsBuilder {
	builder.internal.MinDocCount = &minDocCount

	return builder
}

func (builder *ElasticsearchHistogramSettingsBuilder) applyDefaults() {
}

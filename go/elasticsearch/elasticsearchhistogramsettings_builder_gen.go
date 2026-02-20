// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchHistogramSettings] = (*ElasticsearchHistogramSettingsBuilder)(nil)

type ElasticsearchHistogramSettingsBuilder struct {
	internal *ElasticsearchHistogramSettings
	errors   cog.BuildErrors
}

func NewElasticsearchHistogramSettingsBuilder() *ElasticsearchHistogramSettingsBuilder {
	resource := NewElasticsearchHistogramSettings()
	builder := &ElasticsearchHistogramSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchHistogramSettingsBuilder) Build() (ElasticsearchHistogramSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchHistogramSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchHistogramSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchHistogramSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchHistogramSettingsBuilder) RecordError(path string, err error) *ElasticsearchHistogramSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchHistogramSettingsBuilder) Interval(interval string) *ElasticsearchHistogramSettingsBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *ElasticsearchHistogramSettingsBuilder) MinDocCount(minDocCount string) *ElasticsearchHistogramSettingsBuilder {
	builder.internal.MinDocCount = &minDocCount

	return builder
}

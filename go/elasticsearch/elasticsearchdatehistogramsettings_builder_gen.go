// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchDateHistogramSettings] = (*ElasticsearchDateHistogramSettingsBuilder)(nil)

type ElasticsearchDateHistogramSettingsBuilder struct {
	internal *ElasticsearchDateHistogramSettings
	errors   cog.BuildErrors
}

func NewElasticsearchDateHistogramSettingsBuilder() *ElasticsearchDateHistogramSettingsBuilder {
	resource := NewElasticsearchDateHistogramSettings()
	builder := &ElasticsearchDateHistogramSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchDateHistogramSettingsBuilder) Build() (ElasticsearchDateHistogramSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchDateHistogramSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchDateHistogramSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchDateHistogramSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchDateHistogramSettingsBuilder) RecordError(path string, err error) *ElasticsearchDateHistogramSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchDateHistogramSettingsBuilder) Interval(interval string) *ElasticsearchDateHistogramSettingsBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *ElasticsearchDateHistogramSettingsBuilder) MinDocCount(minDocCount string) *ElasticsearchDateHistogramSettingsBuilder {
	builder.internal.MinDocCount = &minDocCount

	return builder
}

func (builder *ElasticsearchDateHistogramSettingsBuilder) TrimEdges(trimEdges string) *ElasticsearchDateHistogramSettingsBuilder {
	builder.internal.TrimEdges = &trimEdges

	return builder
}

func (builder *ElasticsearchDateHistogramSettingsBuilder) Offset(offset string) *ElasticsearchDateHistogramSettingsBuilder {
	builder.internal.Offset = &offset

	return builder
}

func (builder *ElasticsearchDateHistogramSettingsBuilder) TimeZone(timeZone string) *ElasticsearchDateHistogramSettingsBuilder {
	builder.internal.TimeZone = &timeZone

	return builder
}

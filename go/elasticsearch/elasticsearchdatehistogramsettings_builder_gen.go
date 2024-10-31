// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchDateHistogramSettings] = (*ElasticsearchDateHistogramSettingsBuilder)(nil)

type ElasticsearchDateHistogramSettingsBuilder struct {
	internal *ElasticsearchDateHistogramSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchDateHistogramSettingsBuilder() *ElasticsearchDateHistogramSettingsBuilder {
	resource := &ElasticsearchDateHistogramSettings{}
	builder := &ElasticsearchDateHistogramSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchDateHistogramSettingsBuilder) Build() (ElasticsearchDateHistogramSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchDateHistogramSettings{}, err
	}

	return *builder.internal, nil
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

func (builder *ElasticsearchDateHistogramSettingsBuilder) applyDefaults() {
}

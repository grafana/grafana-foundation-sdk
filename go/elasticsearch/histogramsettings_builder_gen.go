// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HistogramSettings] = (*HistogramSettingsBuilder)(nil)

type HistogramSettingsBuilder struct {
	internal *HistogramSettings
	errors   cog.BuildErrors
}

func NewHistogramSettingsBuilder() *HistogramSettingsBuilder {
	resource := NewHistogramSettings()
	builder := &HistogramSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *HistogramSettingsBuilder) Build() (HistogramSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return HistogramSettings{}, err
	}

	if len(builder.errors) > 0 {
		return HistogramSettings{}, cog.MakeBuildErrors("elasticsearch.histogramSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *HistogramSettingsBuilder) RecordError(path string, err error) *HistogramSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *HistogramSettingsBuilder) Interval(interval string) *HistogramSettingsBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *HistogramSettingsBuilder) MinDocCount(minDocCount string) *HistogramSettingsBuilder {
	builder.internal.MinDocCount = &minDocCount

	return builder
}

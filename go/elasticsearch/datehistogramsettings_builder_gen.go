// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DateHistogramSettings] = (*DateHistogramSettingsBuilder)(nil)

type DateHistogramSettingsBuilder struct {
	internal *DateHistogramSettings
	errors   cog.BuildErrors
}

func NewDateHistogramSettingsBuilder() *DateHistogramSettingsBuilder {
	resource := NewDateHistogramSettings()
	builder := &DateHistogramSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DateHistogramSettingsBuilder) Build() (DateHistogramSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return DateHistogramSettings{}, err
	}

	if len(builder.errors) > 0 {
		return DateHistogramSettings{}, cog.MakeBuildErrors("elasticsearch.dateHistogramSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DateHistogramSettingsBuilder) RecordError(path string, err error) *DateHistogramSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *DateHistogramSettingsBuilder) Interval(interval string) *DateHistogramSettingsBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *DateHistogramSettingsBuilder) MinDocCount(minDocCount string) *DateHistogramSettingsBuilder {
	builder.internal.MinDocCount = &minDocCount

	return builder
}

func (builder *DateHistogramSettingsBuilder) TrimEdges(trimEdges string) *DateHistogramSettingsBuilder {
	builder.internal.TrimEdges = &trimEdges

	return builder
}

func (builder *DateHistogramSettingsBuilder) Offset(offset string) *DateHistogramSettingsBuilder {
	builder.internal.Offset = &offset

	return builder
}

func (builder *DateHistogramSettingsBuilder) TimeZone(timeZone string) *DateHistogramSettingsBuilder {
	builder.internal.TimeZone = &timeZone

	return builder
}

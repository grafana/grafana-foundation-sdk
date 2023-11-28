// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[HistogramSettings] = (*HistogramSettingsBuilder)(nil)

type HistogramSettingsBuilder struct {
	internal *HistogramSettings
	errors   map[string]cog.BuildErrors
}

func NewHistogramSettingsBuilder() *HistogramSettingsBuilder {
	resource := &HistogramSettings{}
	builder := &HistogramSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *HistogramSettingsBuilder) Build() (HistogramSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("HistogramSettings", err)...)
	}

	if len(errs) != 0 {
		return HistogramSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *HistogramSettingsBuilder) Interval(interval string) *HistogramSettingsBuilder {
	builder.internal.Interval = &interval

	return builder
}

func (builder *HistogramSettingsBuilder) MinDocCount(minDocCount string) *HistogramSettingsBuilder {
	builder.internal.MinDocCount = &minDocCount

	return builder
}

func (builder *HistogramSettingsBuilder) applyDefaults() {
}

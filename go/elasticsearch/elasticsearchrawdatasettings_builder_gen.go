// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchRawDataSettings] = (*ElasticsearchRawDataSettingsBuilder)(nil)

type ElasticsearchRawDataSettingsBuilder struct {
	internal *ElasticsearchRawDataSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchRawDataSettingsBuilder() *ElasticsearchRawDataSettingsBuilder {
	resource := &ElasticsearchRawDataSettings{}
	builder := &ElasticsearchRawDataSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchRawDataSettingsBuilder) Build() (ElasticsearchRawDataSettings, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("ElasticsearchRawDataSettings", err)...)
	}

	if len(errs) != 0 {
		return ElasticsearchRawDataSettings{}, errs
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchRawDataSettingsBuilder) Size(size string) *ElasticsearchRawDataSettingsBuilder {
	builder.internal.Size = &size

	return builder
}

func (builder *ElasticsearchRawDataSettingsBuilder) applyDefaults() {
}

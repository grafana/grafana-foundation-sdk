// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchCumulativeSumSettings] = (*ElasticsearchCumulativeSumSettingsBuilder)(nil)

type ElasticsearchCumulativeSumSettingsBuilder struct {
	internal *ElasticsearchCumulativeSumSettings
	errors   cog.BuildErrors
}

func NewElasticsearchCumulativeSumSettingsBuilder() *ElasticsearchCumulativeSumSettingsBuilder {
	resource := NewElasticsearchCumulativeSumSettings()
	builder := &ElasticsearchCumulativeSumSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchCumulativeSumSettingsBuilder) Build() (ElasticsearchCumulativeSumSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchCumulativeSumSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchCumulativeSumSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchCumulativeSumSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchCumulativeSumSettingsBuilder) RecordError(path string, err error) *ElasticsearchCumulativeSumSettingsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ElasticsearchCumulativeSumSettingsBuilder) Format(format string) *ElasticsearchCumulativeSumSettingsBuilder {
	builder.internal.Format = &format

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchDerivativeSettings] = (*ElasticsearchDerivativeSettingsBuilder)(nil)

type ElasticsearchDerivativeSettingsBuilder struct {
	internal *ElasticsearchDerivativeSettings
	errors   cog.BuildErrors
}

func NewElasticsearchDerivativeSettingsBuilder() *ElasticsearchDerivativeSettingsBuilder {
	resource := NewElasticsearchDerivativeSettings()
	builder := &ElasticsearchDerivativeSettingsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchDerivativeSettingsBuilder) Build() (ElasticsearchDerivativeSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchDerivativeSettings{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchDerivativeSettings{}, cog.MakeBuildErrors("elasticsearch.elasticsearchDerivativeSettings", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchDerivativeSettingsBuilder) Unit(unit string) *ElasticsearchDerivativeSettingsBuilder {
	builder.internal.Unit = &unit

	return builder
}

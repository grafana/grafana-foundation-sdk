// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchTermsSettings] = (*ElasticsearchTermsSettingsBuilder)(nil)

type ElasticsearchTermsSettingsBuilder struct {
	internal *ElasticsearchTermsSettings
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchTermsSettingsBuilder() *ElasticsearchTermsSettingsBuilder {
	resource := &ElasticsearchTermsSettings{}
	builder := &ElasticsearchTermsSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchTermsSettingsBuilder) Build() (ElasticsearchTermsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchTermsSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchTermsSettingsBuilder) Order(order TermsOrder) *ElasticsearchTermsSettingsBuilder {
	builder.internal.Order = &order

	return builder
}

func (builder *ElasticsearchTermsSettingsBuilder) Size(size string) *ElasticsearchTermsSettingsBuilder {
	builder.internal.Size = &size

	return builder
}

func (builder *ElasticsearchTermsSettingsBuilder) MinDocCount(minDocCount string) *ElasticsearchTermsSettingsBuilder {
	builder.internal.MinDocCount = &minDocCount

	return builder
}

func (builder *ElasticsearchTermsSettingsBuilder) OrderBy(orderBy string) *ElasticsearchTermsSettingsBuilder {
	builder.internal.OrderBy = &orderBy

	return builder
}

func (builder *ElasticsearchTermsSettingsBuilder) Missing(missing string) *ElasticsearchTermsSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *ElasticsearchTermsSettingsBuilder) applyDefaults() {
}

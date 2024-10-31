// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TermsSettings] = (*TermsSettingsBuilder)(nil)

type TermsSettingsBuilder struct {
	internal *TermsSettings
	errors   map[string]cog.BuildErrors
}

func NewTermsSettingsBuilder() *TermsSettingsBuilder {
	resource := &TermsSettings{}
	builder := &TermsSettingsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TermsSettingsBuilder) Build() (TermsSettings, error) {
	if err := builder.internal.Validate(); err != nil {
		return TermsSettings{}, err
	}

	return *builder.internal, nil
}

func (builder *TermsSettingsBuilder) Order(order TermsOrder) *TermsSettingsBuilder {
	builder.internal.Order = &order

	return builder
}

func (builder *TermsSettingsBuilder) Size(size string) *TermsSettingsBuilder {
	builder.internal.Size = &size

	return builder
}

func (builder *TermsSettingsBuilder) MinDocCount(minDocCount string) *TermsSettingsBuilder {
	builder.internal.MinDocCount = &minDocCount

	return builder
}

func (builder *TermsSettingsBuilder) OrderBy(orderBy string) *TermsSettingsBuilder {
	builder.internal.OrderBy = &orderBy

	return builder
}

func (builder *TermsSettingsBuilder) Missing(missing string) *TermsSettingsBuilder {
	builder.internal.Missing = &missing

	return builder
}

func (builder *TermsSettingsBuilder) applyDefaults() {
}

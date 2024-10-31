// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Terms] = (*TermsBuilder)(nil)

type TermsBuilder struct {
	internal *Terms
	errors   map[string]cog.BuildErrors
}

func NewTermsBuilder() *TermsBuilder {
	resource := &Terms{}
	builder := &TermsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "terms"

	return builder
}

func (builder *TermsBuilder) Build() (Terms, error) {
	if err := builder.internal.Validate(); err != nil {
		return Terms{}, err
	}

	return *builder.internal, nil
}

func (builder *TermsBuilder) Field(field string) *TermsBuilder {
	builder.internal.Field = &field

	return builder
}

func (builder *TermsBuilder) Id(id string) *TermsBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *TermsBuilder) Settings(settings cog.Builder[ElasticsearchTermsSettings]) *TermsBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *TermsBuilder) applyDefaults() {
}

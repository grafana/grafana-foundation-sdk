// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Terms] = (*TermsBuilder)(nil)

type TermsBuilder struct {
	internal *Terms
	errors   cog.BuildErrors
}

func NewTermsBuilder() *TermsBuilder {
	resource := NewTerms()
	builder := &TermsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TermsBuilder) Build() (Terms, error) {
	if err := builder.internal.Validate(); err != nil {
		return Terms{}, err
	}

	if len(builder.errors) > 0 {
		return Terms{}, cog.MakeBuildErrors("elasticsearch.terms", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TermsBuilder) RecordError(path string, err error) *TermsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

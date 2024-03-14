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
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Terms", err)...)
	}

	if len(errs) != 0 {
		return Terms{}, errs
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

func (builder *TermsBuilder) Settings(settings struct {
	Order       *TermsOrder `json:"order,omitempty"`
	Size        *string     `json:"size,omitempty"`
	MinDocCount *string     `json:"min_doc_count,omitempty"`
	OrderBy     *string     `json:"orderBy,omitempty"`
	Missing     *string     `json:"missing,omitempty"`
}) *TermsBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *TermsBuilder) applyDefaults() {
}

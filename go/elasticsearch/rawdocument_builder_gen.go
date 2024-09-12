// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RawDocument] = (*RawDocumentBuilder)(nil)

type RawDocumentBuilder struct {
	internal *RawDocument
	errors   map[string]cog.BuildErrors
}

func NewRawDocumentBuilder() *RawDocumentBuilder {
	resource := &RawDocument{}
	builder := &RawDocumentBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "raw_document"

	return builder
}

func (builder *RawDocumentBuilder) Build() (RawDocument, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("RawDocument", err)...)
	}

	if len(errs) != 0 {
		return RawDocument{}, errs
	}

	return *builder.internal, nil
}

func (builder *RawDocumentBuilder) Id(id string) *RawDocumentBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *RawDocumentBuilder) Settings(settings struct {
	Size *string `json:"size,omitempty"`
}) *RawDocumentBuilder {
	builder.internal.Settings = &settings

	return builder
}

func (builder *RawDocumentBuilder) Hide(hide bool) *RawDocumentBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *RawDocumentBuilder) applyDefaults() {
}

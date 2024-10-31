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
	if err := builder.internal.Validate(); err != nil {
		return RawDocument{}, err
	}

	return *builder.internal, nil
}

func (builder *RawDocumentBuilder) Id(id string) *RawDocumentBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *RawDocumentBuilder) Settings(settings cog.Builder[ElasticsearchRawDocumentSettings]) *RawDocumentBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors["settings"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *RawDocumentBuilder) Hide(hide bool) *RawDocumentBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *RawDocumentBuilder) applyDefaults() {
}

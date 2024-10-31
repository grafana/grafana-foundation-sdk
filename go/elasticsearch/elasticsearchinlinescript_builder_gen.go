// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchInlineScript] = (*ElasticsearchInlineScriptBuilder)(nil)

type ElasticsearchInlineScriptBuilder struct {
	internal *ElasticsearchInlineScript
	errors   map[string]cog.BuildErrors
}

func NewElasticsearchInlineScriptBuilder() *ElasticsearchInlineScriptBuilder {
	resource := &ElasticsearchInlineScript{}
	builder := &ElasticsearchInlineScriptBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ElasticsearchInlineScriptBuilder) Build() (ElasticsearchInlineScript, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchInlineScript{}, err
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchInlineScriptBuilder) Inline(inline string) *ElasticsearchInlineScriptBuilder {
	builder.internal.Inline = &inline

	return builder
}

func (builder *ElasticsearchInlineScriptBuilder) applyDefaults() {
}

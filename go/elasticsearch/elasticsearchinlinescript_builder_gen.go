// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ElasticsearchInlineScript] = (*ElasticsearchInlineScriptBuilder)(nil)

type ElasticsearchInlineScriptBuilder struct {
	internal *ElasticsearchInlineScript
	errors   cog.BuildErrors
}

func NewElasticsearchInlineScriptBuilder() *ElasticsearchInlineScriptBuilder {
	resource := NewElasticsearchInlineScript()
	builder := &ElasticsearchInlineScriptBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ElasticsearchInlineScriptBuilder) Build() (ElasticsearchInlineScript, error) {
	if err := builder.internal.Validate(); err != nil {
		return ElasticsearchInlineScript{}, err
	}

	if len(builder.errors) > 0 {
		return ElasticsearchInlineScript{}, cog.MakeBuildErrors("elasticsearch.elasticsearchInlineScript", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ElasticsearchInlineScriptBuilder) Inline(inline string) *ElasticsearchInlineScriptBuilder {
	builder.internal.Inline = &inline

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[InlineScript] = (*InlineScriptBuilder)(nil)

type InlineScriptBuilder struct {
	internal *InlineScript
	errors   map[string]cog.BuildErrors
}

func NewInlineScriptBuilder() *InlineScriptBuilder {
	resource := &InlineScript{}
	builder := &InlineScriptBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *InlineScriptBuilder) Build() (InlineScript, error) {
	if err := builder.internal.Validate(); err != nil {
		return InlineScript{}, err
	}

	return *builder.internal, nil
}

func (builder *InlineScriptBuilder) String(stringArg string) *InlineScriptBuilder {
	builder.internal.String = &stringArg

	return builder
}

func (builder *InlineScriptBuilder) ElasticsearchInlineScript(elasticsearchInlineScript cog.Builder[ElasticsearchInlineScript]) *InlineScriptBuilder {
	elasticsearchInlineScriptResource, err := elasticsearchInlineScript.Build()
	if err != nil {
		builder.errors["ElasticsearchInlineScript"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.ElasticsearchInlineScript = &elasticsearchInlineScriptResource

	return builder
}

func (builder *InlineScriptBuilder) applyDefaults() {
}

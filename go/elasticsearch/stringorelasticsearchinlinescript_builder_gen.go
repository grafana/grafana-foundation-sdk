// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StringOrElasticsearchInlineScript] = (*StringOrElasticsearchInlineScriptBuilder)(nil)

type StringOrElasticsearchInlineScriptBuilder struct {
	internal *StringOrElasticsearchInlineScript
	errors   map[string]cog.BuildErrors
}

func NewStringOrElasticsearchInlineScriptBuilder() *StringOrElasticsearchInlineScriptBuilder {
	resource := &StringOrElasticsearchInlineScript{}
	builder := &StringOrElasticsearchInlineScriptBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *StringOrElasticsearchInlineScriptBuilder) Build() (StringOrElasticsearchInlineScript, error) {
	if err := builder.internal.Validate(); err != nil {
		return StringOrElasticsearchInlineScript{}, err
	}

	return *builder.internal, nil
}

func (builder *StringOrElasticsearchInlineScriptBuilder) String(stringArg string) *StringOrElasticsearchInlineScriptBuilder {
	builder.internal.String = &stringArg

	return builder
}

func (builder *StringOrElasticsearchInlineScriptBuilder) ElasticsearchInlineScript(elasticsearchInlineScript cog.Builder[ElasticsearchInlineScript]) *StringOrElasticsearchInlineScriptBuilder {
	elasticsearchInlineScriptResource, err := elasticsearchInlineScript.Build()
	if err != nil {
		builder.errors["ElasticsearchInlineScript"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.ElasticsearchInlineScript = &elasticsearchInlineScriptResource

	return builder
}

func (builder *StringOrElasticsearchInlineScriptBuilder) applyDefaults() {
}

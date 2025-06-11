// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[StringOrElasticsearchInlineScript] = (*StringOrElasticsearchInlineScriptBuilder)(nil)

type StringOrElasticsearchInlineScriptBuilder struct {
	internal *StringOrElasticsearchInlineScript
	errors   cog.BuildErrors
}

func NewStringOrElasticsearchInlineScriptBuilder() *StringOrElasticsearchInlineScriptBuilder {
	resource := NewStringOrElasticsearchInlineScript()
	builder := &StringOrElasticsearchInlineScriptBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *StringOrElasticsearchInlineScriptBuilder) Build() (StringOrElasticsearchInlineScript, error) {
	if err := builder.internal.Validate(); err != nil {
		return StringOrElasticsearchInlineScript{}, err
	}

	if len(builder.errors) > 0 {
		return StringOrElasticsearchInlineScript{}, cog.MakeBuildErrors("elasticsearch.stringOrElasticsearchInlineScript", builder.errors)
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
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ElasticsearchInlineScript = &elasticsearchInlineScriptResource

	return builder
}

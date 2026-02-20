// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package text

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CodeOptions] = (*CodeOptionsBuilder)(nil)

type CodeOptionsBuilder struct {
	internal *CodeOptions
	errors   cog.BuildErrors
}

func NewCodeOptionsBuilder() *CodeOptionsBuilder {
	resource := NewCodeOptions()
	builder := &CodeOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CodeOptionsBuilder) Build() (CodeOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return CodeOptions{}, err
	}

	if len(builder.errors) > 0 {
		return CodeOptions{}, cog.MakeBuildErrors("text.codeOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CodeOptionsBuilder) RecordError(path string, err error) *CodeOptionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The language passed to monaco code editor
func (builder *CodeOptionsBuilder) Language(language CodeLanguage) *CodeOptionsBuilder {
	builder.internal.Language = language

	return builder
}

func (builder *CodeOptionsBuilder) ShowLineNumbers(showLineNumbers bool) *CodeOptionsBuilder {
	builder.internal.ShowLineNumbers = showLineNumbers

	return builder
}

func (builder *CodeOptionsBuilder) ShowMiniMap(showMiniMap bool) *CodeOptionsBuilder {
	builder.internal.ShowMiniMap = showMiniMap

	return builder
}

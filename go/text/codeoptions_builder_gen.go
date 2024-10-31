// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package text

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CodeOptions] = (*CodeOptionsBuilder)(nil)

type CodeOptionsBuilder struct {
	internal *CodeOptions
	errors   map[string]cog.BuildErrors
}

func NewCodeOptionsBuilder() *CodeOptionsBuilder {
	resource := &CodeOptions{}
	builder := &CodeOptionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CodeOptionsBuilder) Build() (CodeOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return CodeOptions{}, err
	}

	return *builder.internal, nil
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

func (builder *CodeOptionsBuilder) applyDefaults() {
	builder.Language("plaintext")
	builder.ShowLineNumbers(false)
	builder.ShowMiniMap(false)
}

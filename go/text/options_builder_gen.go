// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package text

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Options] = (*OptionsBuilder)(nil)

type OptionsBuilder struct {
	internal *Options
	errors   cog.BuildErrors
}

func NewOptionsBuilder() *OptionsBuilder {
	resource := NewOptions()
	builder := &OptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsBuilder) Build() (Options, error) {
	if err := builder.internal.Validate(); err != nil {
		return Options{}, err
	}

	if len(builder.errors) > 0 {
		return Options{}, cog.MakeBuildErrors("text.options", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsBuilder) Mode(mode TextMode) *OptionsBuilder {
	builder.internal.Mode = mode

	return builder
}

func (builder *OptionsBuilder) Code(code cog.Builder[CodeOptions]) *OptionsBuilder {
	codeResource, err := code.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Code = &codeResource

	return builder
}

func (builder *OptionsBuilder) Content(content string) *OptionsBuilder {
	builder.internal.Content = content

	return builder
}

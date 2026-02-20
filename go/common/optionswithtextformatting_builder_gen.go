// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[OptionsWithTextFormatting] = (*OptionsWithTextFormattingBuilder)(nil)

// TODO docs
type OptionsWithTextFormattingBuilder struct {
	internal *OptionsWithTextFormatting
	errors   cog.BuildErrors
}

func NewOptionsWithTextFormattingBuilder() *OptionsWithTextFormattingBuilder {
	resource := NewOptionsWithTextFormatting()
	builder := &OptionsWithTextFormattingBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *OptionsWithTextFormattingBuilder) Build() (OptionsWithTextFormatting, error) {
	if err := builder.internal.Validate(); err != nil {
		return OptionsWithTextFormatting{}, err
	}

	if len(builder.errors) > 0 {
		return OptionsWithTextFormatting{}, cog.MakeBuildErrors("common.optionsWithTextFormatting", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *OptionsWithTextFormattingBuilder) RecordError(path string, err error) *OptionsWithTextFormattingBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *OptionsWithTextFormattingBuilder) Text(text cog.Builder[VizTextDisplayOptions]) *OptionsWithTextFormattingBuilder {
	textResource, err := text.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Text = &textResource

	return builder
}

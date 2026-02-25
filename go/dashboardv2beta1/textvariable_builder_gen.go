// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TextVariableKind] = (*TextVariableBuilder)(nil)

// Text variable kind
type TextVariableBuilder struct {
	internal *TextVariableKind
	errors   cog.BuildErrors
}

func NewTextVariableBuilder(name string) *TextVariableBuilder {
	resource := NewTextVariableKind()
	builder := &TextVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "TextVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *TextVariableBuilder) Build() (TextVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return TextVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return TextVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.textVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TextVariableBuilder) RecordError(path string, err error) *TextVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *TextVariableBuilder) Name(name string) *TextVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *TextVariableBuilder) Current(current VariableOption) *TextVariableBuilder {
	builder.internal.Spec.Current = current

	return builder
}

func (builder *TextVariableBuilder) Query(query string) *TextVariableBuilder {
	builder.internal.Spec.Query = query

	return builder
}

func (builder *TextVariableBuilder) Label(label string) *TextVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *TextVariableBuilder) Hide(hide VariableHide) *TextVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *TextVariableBuilder) SkipUrlSync(skipUrlSync bool) *TextVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *TextVariableBuilder) Description(description string) *TextVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}

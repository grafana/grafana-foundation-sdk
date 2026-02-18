// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableModel] = (*TextBoxVariableBuilder)(nil)

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
type TextBoxVariableBuilder struct {
	internal *VariableModel
	errors   cog.BuildErrors
}

func NewTextBoxVariableBuilder(name string) *TextBoxVariableBuilder {
	resource := NewVariableModel()
	builder := &TextBoxVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Name = name
	builder.internal.Type = "textbox"

	return builder
}

func (builder *TextBoxVariableBuilder) Build() (VariableModel, error) {
	if err := builder.internal.Validate(); err != nil {
		return VariableModel{}, err
	}

	if len(builder.errors) > 0 {
		return VariableModel{}, cog.MakeBuildErrors("dashboard.textBoxVariable", builder.errors)
	}

	return *builder.internal, nil
}

// Name of variable
func (builder *TextBoxVariableBuilder) Name(name string) *TextBoxVariableBuilder {
	builder.internal.Name = name

	return builder
}

// Optional display name
func (builder *TextBoxVariableBuilder) Label(label string) *TextBoxVariableBuilder {
	builder.internal.Label = &label

	return builder
}

// Visibility configuration for the variable
func (builder *TextBoxVariableBuilder) Hide(hide VariableHide) *TextBoxVariableBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Description of variable. It can be defined but `null`.
func (builder *TextBoxVariableBuilder) Description(description string) *TextBoxVariableBuilder {
	builder.internal.Description = &description

	return builder
}

// Query used to fetch values for a variable
func (builder *TextBoxVariableBuilder) DefaultValue(query StringOrMap) *TextBoxVariableBuilder {
	builder.internal.Query = &query

	return builder
}

// Shows current selected variable text/value on the dashboard
func (builder *TextBoxVariableBuilder) Current(current VariableOption) *TextBoxVariableBuilder {
	builder.internal.Current = &current

	return builder
}

// Allow custom values to be entered in the variable
func (builder *TextBoxVariableBuilder) AllowCustomValue(allowCustomValue bool) *TextBoxVariableBuilder {
	builder.internal.AllowCustomValue = &allowCustomValue

	return builder
}

// Options that can be selected for a variable.
func (builder *TextBoxVariableBuilder) Options(options []VariableOption) *TextBoxVariableBuilder {
	builder.internal.Options = options

	return builder
}

func (builder *TextBoxVariableBuilder) Definition(definition string) *TextBoxVariableBuilder {
	builder.internal.Definition = &definition

	return builder
}

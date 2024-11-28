// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableModel] = (*ConstantVariableBuilder)(nil)

// A variable is a placeholder for a value. You can use variables in metric queries and in panel titles.
type ConstantVariableBuilder struct {
	internal *VariableModel
	errors   map[string]cog.BuildErrors
}

func NewConstantVariableBuilder(name string) *ConstantVariableBuilder {
	resource := NewVariableModel()
	builder := &ConstantVariableBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	builder.internal.Name = name
	builder.internal.Type = "constant"
	valHide := VariableHideHideVariable
	builder.internal.Hide = &valHide

	return builder
}

func (builder *ConstantVariableBuilder) Build() (VariableModel, error) {
	if err := builder.internal.Validate(); err != nil {
		return VariableModel{}, err
	}

	return *builder.internal, nil
}

// Name of variable
func (builder *ConstantVariableBuilder) Name(name string) *ConstantVariableBuilder {
	builder.internal.Name = name

	return builder
}

// Optional display name
func (builder *ConstantVariableBuilder) Label(label string) *ConstantVariableBuilder {
	builder.internal.Label = &label

	return builder
}

// Description of variable. It can be defined but `null`.
func (builder *ConstantVariableBuilder) Description(description string) *ConstantVariableBuilder {
	builder.internal.Description = &description

	return builder
}

// Query used to fetch values for a variable
func (builder *ConstantVariableBuilder) Value(query StringOrMap) *ConstantVariableBuilder {
	builder.internal.Query = &query

	return builder
}

// Allow custom values to be entered in the variable
func (builder *ConstantVariableBuilder) AllowCustomValue(allowCustomValue bool) *ConstantVariableBuilder {
	builder.internal.AllowCustomValue = &allowCustomValue

	return builder
}

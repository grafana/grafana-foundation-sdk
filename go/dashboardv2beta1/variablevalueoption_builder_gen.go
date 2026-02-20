// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[VariableValueOption] = (*VariableValueOptionBuilder)(nil)

// FIXME: should we introduce this? --- Variable value option
type VariableValueOptionBuilder struct {
	internal *VariableValueOption
	errors   cog.BuildErrors
}

func NewVariableValueOptionBuilder() *VariableValueOptionBuilder {
	resource := NewVariableValueOption()
	builder := &VariableValueOptionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *VariableValueOptionBuilder) Build() (VariableValueOption, error) {
	if err := builder.internal.Validate(); err != nil {
		return VariableValueOption{}, err
	}

	if len(builder.errors) > 0 {
		return VariableValueOption{}, cog.MakeBuildErrors("dashboardv2beta1.variableValueOption", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *VariableValueOptionBuilder) RecordError(path string, err error) *VariableValueOptionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *VariableValueOptionBuilder) Label(label string) *VariableValueOptionBuilder {
	builder.internal.Label = label

	return builder
}

func (builder *VariableValueOptionBuilder) Value(value cog.Builder[VariableValueSingle]) *VariableValueOptionBuilder {
	valueResource, err := value.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Value = valueResource

	return builder
}

func (builder *VariableValueOptionBuilder) Group(group string) *VariableValueOptionBuilder {
	builder.internal.Group = &group

	return builder
}

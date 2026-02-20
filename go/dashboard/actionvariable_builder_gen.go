// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ActionVariable] = (*ActionVariableBuilder)(nil)

type ActionVariableBuilder struct {
	internal *ActionVariable
	errors   cog.BuildErrors
}

func NewActionVariableBuilder() *ActionVariableBuilder {
	resource := NewActionVariable()
	builder := &ActionVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ActionVariableBuilder) Build() (ActionVariable, error) {
	if err := builder.internal.Validate(); err != nil {
		return ActionVariable{}, err
	}

	if len(builder.errors) > 0 {
		return ActionVariable{}, cog.MakeBuildErrors("dashboard.actionVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ActionVariableBuilder) RecordError(path string, err error) *ActionVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ActionVariableBuilder) Key(key string) *ActionVariableBuilder {
	builder.internal.Key = key

	return builder
}

func (builder *ActionVariableBuilder) Name(name string) *ActionVariableBuilder {
	builder.internal.Name = name

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[GroupByVariableKind] = (*GroupByVariableBuilder)(nil)

// Group variable kind
type GroupByVariableBuilder struct {
	internal *GroupByVariableKind
	errors   cog.BuildErrors
}

func NewGroupByVariableBuilder(name string) *GroupByVariableBuilder {
	resource := NewGroupByVariableKind()
	builder := &GroupByVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "GroupByVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *GroupByVariableBuilder) Build() (GroupByVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return GroupByVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return GroupByVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.groupByVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *GroupByVariableBuilder) RecordError(path string, err error) *GroupByVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *GroupByVariableBuilder) Group(group string) *GroupByVariableBuilder {
	builder.internal.Group = group

	return builder
}

func (builder *GroupByVariableBuilder) Datasource(datasource cog.Builder[Dashboardv2beta1GroupByVariableKindDatasource]) *GroupByVariableBuilder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

func (builder *GroupByVariableBuilder) Spec(spec GroupByVariableSpec) *GroupByVariableBuilder {
	builder.internal.Spec = spec

	return builder
}

func (builder *GroupByVariableBuilder) Name(name string) *GroupByVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *GroupByVariableBuilder) DefaultValue(defaultValue VariableOption) *GroupByVariableBuilder {
	builder.internal.Spec.DefaultValue = &defaultValue

	return builder
}

func (builder *GroupByVariableBuilder) Current(current VariableOption) *GroupByVariableBuilder {
	builder.internal.Spec.Current = current

	return builder
}

func (builder *GroupByVariableBuilder) Options(options []VariableOption) *GroupByVariableBuilder {
	builder.internal.Spec.Options = options

	return builder
}

func (builder *GroupByVariableBuilder) Multi(multi bool) *GroupByVariableBuilder {
	builder.internal.Spec.Multi = multi

	return builder
}

func (builder *GroupByVariableBuilder) Label(label string) *GroupByVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *GroupByVariableBuilder) Hide(hide VariableHide) *GroupByVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *GroupByVariableBuilder) SkipUrlSync(skipUrlSync bool) *GroupByVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *GroupByVariableBuilder) Description(description string) *GroupByVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}

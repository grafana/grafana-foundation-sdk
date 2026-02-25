// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryVariableKind] = (*QueryVariableBuilder)(nil)

// Query variable kind
type QueryVariableBuilder struct {
	internal *QueryVariableKind
	errors   cog.BuildErrors
}

func NewQueryVariableBuilder(name string) *QueryVariableBuilder {
	resource := NewQueryVariableKind()
	builder := &QueryVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "QueryVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *QueryVariableBuilder) Build() (QueryVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return QueryVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.queryVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryVariableBuilder) RecordError(path string, err error) *QueryVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryVariableBuilder) Name(name string) *QueryVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *QueryVariableBuilder) Current(current VariableOption) *QueryVariableBuilder {
	builder.internal.Spec.Current = current

	return builder
}

func (builder *QueryVariableBuilder) Label(label string) *QueryVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *QueryVariableBuilder) Hide(hide VariableHide) *QueryVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *QueryVariableBuilder) Refresh(refresh VariableRefresh) *QueryVariableBuilder {
	builder.internal.Spec.Refresh = refresh

	return builder
}

func (builder *QueryVariableBuilder) SkipUrlSync(skipUrlSync bool) *QueryVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *QueryVariableBuilder) Description(description string) *QueryVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}

func (builder *QueryVariableBuilder) Query(query cog.Builder[DataQueryKind]) *QueryVariableBuilder {
	queryResource, err := query.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Query = queryResource

	return builder
}

func (builder *QueryVariableBuilder) Regex(regex string) *QueryVariableBuilder {
	builder.internal.Spec.Regex = regex

	return builder
}

func (builder *QueryVariableBuilder) Sort(sort VariableSort) *QueryVariableBuilder {
	builder.internal.Spec.Sort = sort

	return builder
}

func (builder *QueryVariableBuilder) Definition(definition string) *QueryVariableBuilder {
	builder.internal.Spec.Definition = &definition

	return builder
}

func (builder *QueryVariableBuilder) Options(options []VariableOption) *QueryVariableBuilder {
	builder.internal.Spec.Options = options

	return builder
}

func (builder *QueryVariableBuilder) Multi(multi bool) *QueryVariableBuilder {
	builder.internal.Spec.Multi = multi

	return builder
}

func (builder *QueryVariableBuilder) IncludeAll(includeAll bool) *QueryVariableBuilder {
	builder.internal.Spec.IncludeAll = includeAll

	return builder
}

func (builder *QueryVariableBuilder) AllValue(allValue string) *QueryVariableBuilder {
	builder.internal.Spec.AllValue = &allValue

	return builder
}

func (builder *QueryVariableBuilder) Placeholder(placeholder string) *QueryVariableBuilder {
	builder.internal.Spec.Placeholder = &placeholder

	return builder
}

func (builder *QueryVariableBuilder) AllowCustomValue(allowCustomValue bool) *QueryVariableBuilder {
	builder.internal.Spec.AllowCustomValue = allowCustomValue

	return builder
}

func (builder *QueryVariableBuilder) StaticOptions(staticOptions []VariableOption) *QueryVariableBuilder {
	builder.internal.Spec.StaticOptions = staticOptions

	return builder
}

func (builder *QueryVariableBuilder) StaticOptionsOrder(staticOptionsOrder QueryVariableSpecStaticOptionsOrder) *QueryVariableBuilder {
	builder.internal.Spec.StaticOptionsOrder = &staticOptionsOrder

	return builder
}

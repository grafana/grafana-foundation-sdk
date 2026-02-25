// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[IntervalVariableKind] = (*IntervalVariableBuilder)(nil)

// Interval variable kind
type IntervalVariableBuilder struct {
	internal *IntervalVariableKind
	errors   cog.BuildErrors
}

func NewIntervalVariableBuilder(name string) *IntervalVariableBuilder {
	resource := NewIntervalVariableKind()
	builder := &IntervalVariableBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "IntervalVariable"
	builder.internal.Spec.Name = name

	return builder
}

func (builder *IntervalVariableBuilder) Build() (IntervalVariableKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return IntervalVariableKind{}, err
	}

	if len(builder.errors) > 0 {
		return IntervalVariableKind{}, cog.MakeBuildErrors("dashboardv2beta1.intervalVariable", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *IntervalVariableBuilder) RecordError(path string, err error) *IntervalVariableBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *IntervalVariableBuilder) Name(name string) *IntervalVariableBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *IntervalVariableBuilder) Query(query string) *IntervalVariableBuilder {
	builder.internal.Spec.Query = query

	return builder
}

func (builder *IntervalVariableBuilder) Current(current VariableOption) *IntervalVariableBuilder {
	builder.internal.Spec.Current = current

	return builder
}

func (builder *IntervalVariableBuilder) Options(options []VariableOption) *IntervalVariableBuilder {
	builder.internal.Spec.Options = options

	return builder
}

func (builder *IntervalVariableBuilder) Auto(auto bool) *IntervalVariableBuilder {
	builder.internal.Spec.Auto = auto

	return builder
}

func (builder *IntervalVariableBuilder) AutoMin(autoMin string) *IntervalVariableBuilder {
	builder.internal.Spec.AutoMin = autoMin

	return builder
}

func (builder *IntervalVariableBuilder) AutoCount(autoCount int64) *IntervalVariableBuilder {
	builder.internal.Spec.AutoCount = autoCount

	return builder
}

func (builder *IntervalVariableBuilder) Refresh(refresh VariableRefresh) *IntervalVariableBuilder {
	builder.internal.Spec.Refresh = refresh

	return builder
}

func (builder *IntervalVariableBuilder) Label(label string) *IntervalVariableBuilder {
	builder.internal.Spec.Label = &label

	return builder
}

func (builder *IntervalVariableBuilder) Hide(hide VariableHide) *IntervalVariableBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *IntervalVariableBuilder) SkipUrlSync(skipUrlSync bool) *IntervalVariableBuilder {
	builder.internal.Spec.SkipUrlSync = skipUrlSync

	return builder
}

func (builder *IntervalVariableBuilder) Description(description string) *IntervalVariableBuilder {
	builder.internal.Spec.Description = &description

	return builder
}

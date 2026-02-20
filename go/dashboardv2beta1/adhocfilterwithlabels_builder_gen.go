// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AdHocFilterWithLabels] = (*AdHocFilterWithLabelsBuilder)(nil)

// Define the AdHocFilterWithLabels type
type AdHocFilterWithLabelsBuilder struct {
	internal *AdHocFilterWithLabels
	errors   cog.BuildErrors
}

func NewAdHocFilterWithLabelsBuilder() *AdHocFilterWithLabelsBuilder {
	resource := NewAdHocFilterWithLabels()
	builder := &AdHocFilterWithLabelsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) Build() (AdHocFilterWithLabels, error) {
	if err := builder.internal.Validate(); err != nil {
		return AdHocFilterWithLabels{}, err
	}

	if len(builder.errors) > 0 {
		return AdHocFilterWithLabels{}, cog.MakeBuildErrors("dashboardv2beta1.adHocFilterWithLabels", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AdHocFilterWithLabelsBuilder) RecordError(path string, err error) *AdHocFilterWithLabelsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) Key(key string) *AdHocFilterWithLabelsBuilder {
	builder.internal.Key = key

	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) Operator(operator string) *AdHocFilterWithLabelsBuilder {
	builder.internal.Operator = operator

	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) Value(value string) *AdHocFilterWithLabelsBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) Values(values []string) *AdHocFilterWithLabelsBuilder {
	builder.internal.Values = values

	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) KeyLabel(keyLabel string) *AdHocFilterWithLabelsBuilder {
	builder.internal.KeyLabel = &keyLabel

	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) ValueLabels(valueLabels []string) *AdHocFilterWithLabelsBuilder {
	builder.internal.ValueLabels = valueLabels

	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) ForceEdit(forceEdit bool) *AdHocFilterWithLabelsBuilder {
	builder.internal.ForceEdit = &forceEdit

	return builder
}

func (builder *AdHocFilterWithLabelsBuilder) Origin(origin string) *AdHocFilterWithLabelsBuilder {
	builder.internal.Origin = &origin

	return builder
}

// @deprecated
func (builder *AdHocFilterWithLabelsBuilder) Condition(condition string) *AdHocFilterWithLabelsBuilder {
	builder.internal.Condition = &condition

	return builder
}

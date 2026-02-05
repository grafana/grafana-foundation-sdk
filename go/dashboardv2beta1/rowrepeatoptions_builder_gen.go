// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RowRepeatOptions] = (*RowRepeatOptionsBuilder)(nil)

type RowRepeatOptionsBuilder struct {
	internal *RowRepeatOptions
	errors   cog.BuildErrors
}

func NewRowRepeatOptionsBuilder() *RowRepeatOptionsBuilder {
	resource := NewRowRepeatOptions()
	builder := &RowRepeatOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RowRepeatOptionsBuilder) Build() (RowRepeatOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return RowRepeatOptions{}, err
	}

	if len(builder.errors) > 0 {
		return RowRepeatOptions{}, cog.MakeBuildErrors("dashboardv2beta1.rowRepeatOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RowRepeatOptionsBuilder) Value(value string) *RowRepeatOptionsBuilder {
	builder.internal.Value = value

	return builder
}

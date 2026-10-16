// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CellSpec] = (*CellSpecBuilder)(nil)

type CellSpecBuilder struct {
	internal *CellSpec
	errors   cog.BuildErrors
}

func NewCellSpecBuilder() *CellSpecBuilder {
	resource := NewCellSpec()
	builder := &CellSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CellSpecBuilder) Build() (CellSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return CellSpec{}, err
	}

	if len(builder.errors) > 0 {
		return CellSpec{}, cog.MakeBuildErrors("dashboardv2beta1.cellSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CellSpecBuilder) RecordError(path string, err error) *CellSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *CellSpecBuilder) Content(content CellContentKind) *CellSpecBuilder {
	builder.internal.Content = content

	return builder
}

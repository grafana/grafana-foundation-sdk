// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RowsLayoutSpec] = (*RowsLayoutSpecBuilder)(nil)

type RowsLayoutSpecBuilder struct {
	internal *RowsLayoutSpec
	errors   cog.BuildErrors
}

func NewRowsLayoutSpecBuilder() *RowsLayoutSpecBuilder {
	resource := NewRowsLayoutSpec()
	builder := &RowsLayoutSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RowsLayoutSpecBuilder) Build() (RowsLayoutSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return RowsLayoutSpec{}, err
	}

	if len(builder.errors) > 0 {
		return RowsLayoutSpec{}, cog.MakeBuildErrors("dashboardv2beta1.rowsLayoutSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RowsLayoutSpecBuilder) Rows(rows []cog.Builder[RowsLayoutRowKind]) *RowsLayoutSpecBuilder {
	rowsResources := make([]RowsLayoutRowKind, 0, len(rows))
	for _, r1 := range rows {
		rowsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		rowsResources = append(rowsResources, rowsDepth1)
	}
	builder.internal.Rows = rowsResources

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RowsLayoutKind] = (*RowsLayoutBuilder)(nil)

type RowsLayoutBuilder struct {
	internal *RowsLayoutKind
	errors   cog.BuildErrors
}

func NewRowsLayoutBuilder() *RowsLayoutBuilder {
	resource := NewRowsLayoutKind()
	builder := &RowsLayoutBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "RowsLayout"

	return builder
}

func (builder *RowsLayoutBuilder) Build() (RowsLayoutKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return RowsLayoutKind{}, err
	}

	if len(builder.errors) > 0 {
		return RowsLayoutKind{}, cog.MakeBuildErrors("dashboardv2beta1.rowsLayout", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RowsLayoutBuilder) RecordError(path string, err error) *RowsLayoutBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RowsLayoutBuilder) Rows(rows []cog.Builder[RowsLayoutRowKind]) *RowsLayoutBuilder {
	rowsResources := make([]RowsLayoutRowKind, 0, len(rows))
	for _, r1 := range rows {
		rowsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		rowsResources = append(rowsResources, rowsDepth1)
	}
	builder.internal.Spec.Rows = rowsResources

	return builder
}

func (builder *RowsLayoutBuilder) Row(row cog.Builder[RowsLayoutRowKind]) *RowsLayoutBuilder {
	rowResource, err := row.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Rows = append(builder.internal.Spec.Rows, rowResource)

	return builder
}

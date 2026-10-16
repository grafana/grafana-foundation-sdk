// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotebookLayoutSpec] = (*NotebookLayoutSpecBuilder)(nil)

type NotebookLayoutSpecBuilder struct {
	internal *NotebookLayoutSpec
	errors   cog.BuildErrors
}

func NewNotebookLayoutSpecBuilder() *NotebookLayoutSpecBuilder {
	resource := NewNotebookLayoutSpec()
	builder := &NotebookLayoutSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *NotebookLayoutSpecBuilder) Build() (NotebookLayoutSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotebookLayoutSpec{}, err
	}

	if len(builder.errors) > 0 {
		return NotebookLayoutSpec{}, cog.MakeBuildErrors("dashboardv2beta1.notebookLayoutSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NotebookLayoutSpecBuilder) RecordError(path string, err error) *NotebookLayoutSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *NotebookLayoutSpecBuilder) Cells(cells []cog.Builder[NotebookLayoutItemKind]) *NotebookLayoutSpecBuilder {
	cellsResources := make([]NotebookLayoutItemKind, 0, len(cells))
	for _, r1 := range cells {
		cellsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		cellsResources = append(cellsResources, cellsDepth1)
	}
	builder.internal.Cells = cellsResources

	return builder
}

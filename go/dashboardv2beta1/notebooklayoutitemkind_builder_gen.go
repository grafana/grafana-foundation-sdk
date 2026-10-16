// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotebookLayoutItemKind] = (*NotebookLayoutItemKindBuilder)(nil)

type NotebookLayoutItemKindBuilder struct {
	internal *NotebookLayoutItemKind
	errors   cog.BuildErrors
}

func NewNotebookLayoutItemKindBuilder() *NotebookLayoutItemKindBuilder {
	resource := NewNotebookLayoutItemKind()
	builder := &NotebookLayoutItemKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "NotebookLayoutItem"

	return builder
}

func (builder *NotebookLayoutItemKindBuilder) Build() (NotebookLayoutItemKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotebookLayoutItemKind{}, err
	}

	if len(builder.errors) > 0 {
		return NotebookLayoutItemKind{}, cog.MakeBuildErrors("dashboardv2beta1.notebookLayoutItemKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NotebookLayoutItemKindBuilder) RecordError(path string, err error) *NotebookLayoutItemKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *NotebookLayoutItemKindBuilder) Spec(spec cog.Builder[NotebookLayoutItemSpec]) *NotebookLayoutItemKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}

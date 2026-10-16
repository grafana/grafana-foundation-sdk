// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotebookLayoutKind] = (*NotebookLayoutKindBuilder)(nil)

type NotebookLayoutKindBuilder struct {
	internal *NotebookLayoutKind
	errors   cog.BuildErrors
}

func NewNotebookLayoutKindBuilder() *NotebookLayoutKindBuilder {
	resource := NewNotebookLayoutKind()
	builder := &NotebookLayoutKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "NotebookLayout"

	return builder
}

func (builder *NotebookLayoutKindBuilder) Build() (NotebookLayoutKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotebookLayoutKind{}, err
	}

	if len(builder.errors) > 0 {
		return NotebookLayoutKind{}, cog.MakeBuildErrors("dashboardv2beta1.notebookLayoutKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NotebookLayoutKindBuilder) RecordError(path string, err error) *NotebookLayoutKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *NotebookLayoutKindBuilder) Spec(spec cog.Builder[NotebookLayoutSpec]) *NotebookLayoutKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}

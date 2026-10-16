// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotebookLayoutItemSpec] = (*NotebookLayoutItemSpecBuilder)(nil)

// One ordered item in a notebook layout. `element` references either a CellKind
// (markdown/code content) or a V2PanelKind in the notebook's elements map. `source`
// records who authored the cell; `collapsed` hides the body in the UI.
type NotebookLayoutItemSpecBuilder struct {
	internal *NotebookLayoutItemSpec
	errors   cog.BuildErrors
}

func NewNotebookLayoutItemSpecBuilder() *NotebookLayoutItemSpecBuilder {
	resource := NewNotebookLayoutItemSpec()
	builder := &NotebookLayoutItemSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *NotebookLayoutItemSpecBuilder) Build() (NotebookLayoutItemSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotebookLayoutItemSpec{}, err
	}

	if len(builder.errors) > 0 {
		return NotebookLayoutItemSpec{}, cog.MakeBuildErrors("dashboardv2beta1.notebookLayoutItemSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NotebookLayoutItemSpecBuilder) RecordError(path string, err error) *NotebookLayoutItemSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *NotebookLayoutItemSpecBuilder) Element(element cog.Builder[ElementReference]) *NotebookLayoutItemSpecBuilder {
	elementResource, err := element.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Element = elementResource

	return builder
}

func (builder *NotebookLayoutItemSpecBuilder) Source(source NotebookLayoutItemSpecSource) *NotebookLayoutItemSpecBuilder {
	builder.internal.Source = source

	return builder
}

func (builder *NotebookLayoutItemSpecBuilder) Collapsed(collapsed bool) *NotebookLayoutItemSpecBuilder {
	builder.internal.Collapsed = &collapsed

	return builder
}

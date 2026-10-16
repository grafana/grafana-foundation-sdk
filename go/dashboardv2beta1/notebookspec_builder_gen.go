// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[NotebookSpec] = (*NotebookSpecBuilder)(nil)

// A notebook spec is a dashboard spec with narrative content. It has a title, optional description, tags, time
// settings, and a map of elements (panels and cells) referenced by the layout.
type NotebookSpecBuilder struct {
	internal *NotebookSpec
	errors   cog.BuildErrors
}

func NewNotebookSpecBuilder() *NotebookSpecBuilder {
	resource := NewNotebookSpec()
	builder := &NotebookSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *NotebookSpecBuilder) Build() (NotebookSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return NotebookSpec{}, err
	}

	if len(builder.errors) > 0 {
		return NotebookSpec{}, cog.MakeBuildErrors("dashboardv2beta1.notebookSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *NotebookSpecBuilder) RecordError(path string, err error) *NotebookSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *NotebookSpecBuilder) Title(title string) *NotebookSpecBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *NotebookSpecBuilder) Description(description string) *NotebookSpecBuilder {
	builder.internal.Description = &description

	return builder
}

func (builder *NotebookSpecBuilder) Tags(tags []string) *NotebookSpecBuilder {
	builder.internal.Tags = tags

	return builder
}

func (builder *NotebookSpecBuilder) TimeSettings(timeSettings cog.Builder[TimeSettingsSpec]) *NotebookSpecBuilder {
	timeSettingsResource, err := timeSettings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TimeSettings = timeSettingsResource

	return builder
}

func (builder *NotebookSpecBuilder) Elements(elements map[string]NotebookElement) *NotebookSpecBuilder {
	builder.internal.Elements = elements

	return builder
}

func (builder *NotebookSpecBuilder) Layout(layout cog.Builder[NotebookLayoutKind]) *NotebookSpecBuilder {
	layoutResource, err := layout.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Layout = layoutResource

	return builder
}

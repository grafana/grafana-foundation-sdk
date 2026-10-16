// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MarkdownCellContentSpec] = (*MarkdownCellContentSpecBuilder)(nil)

type MarkdownCellContentSpecBuilder struct {
	internal *MarkdownCellContentSpec
	errors   cog.BuildErrors
}

func NewMarkdownCellContentSpecBuilder() *MarkdownCellContentSpecBuilder {
	resource := NewMarkdownCellContentSpec()
	builder := &MarkdownCellContentSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MarkdownCellContentSpecBuilder) Build() (MarkdownCellContentSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return MarkdownCellContentSpec{}, err
	}

	if len(builder.errors) > 0 {
		return MarkdownCellContentSpec{}, cog.MakeBuildErrors("dashboardv2beta1.markdownCellContentSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MarkdownCellContentSpecBuilder) RecordError(path string, err error) *MarkdownCellContentSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MarkdownCellContentSpecBuilder) Text(text string) *MarkdownCellContentSpecBuilder {
	builder.internal.Text = text

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MarkdownCellContentKind] = (*MarkdownCellContentKindBuilder)(nil)

type MarkdownCellContentKindBuilder struct {
	internal *MarkdownCellContentKind
	errors   cog.BuildErrors
}

func NewMarkdownCellContentKindBuilder() *MarkdownCellContentKindBuilder {
	resource := NewMarkdownCellContentKind()
	builder := &MarkdownCellContentKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "Markdown"

	return builder
}

func (builder *MarkdownCellContentKindBuilder) Build() (MarkdownCellContentKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return MarkdownCellContentKind{}, err
	}

	if len(builder.errors) > 0 {
		return MarkdownCellContentKind{}, cog.MakeBuildErrors("dashboardv2beta1.markdownCellContentKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MarkdownCellContentKindBuilder) RecordError(path string, err error) *MarkdownCellContentKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MarkdownCellContentKindBuilder) Spec(spec cog.Builder[MarkdownCellContentSpec]) *MarkdownCellContentKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CodeCellContentKind] = (*CodeCellContentKindBuilder)(nil)

type CodeCellContentKindBuilder struct {
	internal *CodeCellContentKind
	errors   cog.BuildErrors
}

func NewCodeCellContentKindBuilder() *CodeCellContentKindBuilder {
	resource := NewCodeCellContentKind()
	builder := &CodeCellContentKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "Code"

	return builder
}

func (builder *CodeCellContentKindBuilder) Build() (CodeCellContentKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return CodeCellContentKind{}, err
	}

	if len(builder.errors) > 0 {
		return CodeCellContentKind{}, cog.MakeBuildErrors("dashboardv2beta1.codeCellContentKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CodeCellContentKindBuilder) RecordError(path string, err error) *CodeCellContentKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *CodeCellContentKindBuilder) Spec(spec cog.Builder[CodeCellContentSpec]) *CodeCellContentKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}

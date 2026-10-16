// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CodeCellContentSpec] = (*CodeCellContentSpecBuilder)(nil)

type CodeCellContentSpecBuilder struct {
	internal *CodeCellContentSpec
	errors   cog.BuildErrors
}

func NewCodeCellContentSpecBuilder() *CodeCellContentSpecBuilder {
	resource := NewCodeCellContentSpec()
	builder := &CodeCellContentSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *CodeCellContentSpecBuilder) Build() (CodeCellContentSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return CodeCellContentSpec{}, err
	}

	if len(builder.errors) > 0 {
		return CodeCellContentSpec{}, cog.MakeBuildErrors("dashboardv2beta1.codeCellContentSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CodeCellContentSpecBuilder) RecordError(path string, err error) *CodeCellContentSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *CodeCellContentSpecBuilder) Language(language string) *CodeCellContentSpecBuilder {
	builder.internal.Language = language

	return builder
}

func (builder *CodeCellContentSpecBuilder) Code(code string) *CodeCellContentSpecBuilder {
	builder.internal.Code = code

	return builder
}

func (builder *CodeCellContentSpecBuilder) Highlight(highlight []int64) *CodeCellContentSpecBuilder {
	builder.internal.Highlight = highlight

	return builder
}

func (builder *CodeCellContentSpecBuilder) Annotation(annotation string) *CodeCellContentSpecBuilder {
	builder.internal.Annotation = &annotation

	return builder
}

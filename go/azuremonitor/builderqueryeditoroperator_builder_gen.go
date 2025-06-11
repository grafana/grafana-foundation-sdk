// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorOperator] = (*BuilderQueryEditorOperatorBuilder)(nil)

type BuilderQueryEditorOperatorBuilder struct {
	internal *BuilderQueryEditorOperator
	errors   cog.BuildErrors
}

func NewBuilderQueryEditorOperatorBuilder() *BuilderQueryEditorOperatorBuilder {
	resource := NewBuilderQueryEditorOperator()
	builder := &BuilderQueryEditorOperatorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BuilderQueryEditorOperatorBuilder) Build() (BuilderQueryEditorOperator, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorOperator{}, err
	}

	if len(builder.errors) > 0 {
		return BuilderQueryEditorOperator{}, cog.MakeBuildErrors("azuremonitor.builderQueryEditorOperator", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorOperatorBuilder) Name(name string) *BuilderQueryEditorOperatorBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *BuilderQueryEditorOperatorBuilder) Value(value string) *BuilderQueryEditorOperatorBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *BuilderQueryEditorOperatorBuilder) LabelValue(labelValue string) *BuilderQueryEditorOperatorBuilder {
	builder.internal.LabelValue = &labelValue

	return builder
}

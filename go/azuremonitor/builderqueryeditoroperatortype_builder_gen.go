// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryEditorOperatorType] = (*BuilderQueryEditorOperatorTypeBuilder)(nil)

type BuilderQueryEditorOperatorTypeBuilder struct {
	internal *BuilderQueryEditorOperatorType
	errors   cog.BuildErrors
}

func NewBuilderQueryEditorOperatorTypeBuilder() *BuilderQueryEditorOperatorTypeBuilder {
	resource := NewBuilderQueryEditorOperatorType()
	builder := &BuilderQueryEditorOperatorTypeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BuilderQueryEditorOperatorTypeBuilder) Build() (BuilderQueryEditorOperatorType, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryEditorOperatorType{}, err
	}

	if len(builder.errors) > 0 {
		return BuilderQueryEditorOperatorType{}, cog.MakeBuildErrors("azuremonitor.builderQueryEditorOperatorType", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryEditorOperatorTypeBuilder) String(stringArg string) *BuilderQueryEditorOperatorTypeBuilder {
	builder.internal.String = &stringArg

	return builder
}

func (builder *BuilderQueryEditorOperatorTypeBuilder) Bool(boolArg bool) *BuilderQueryEditorOperatorTypeBuilder {
	builder.internal.Bool = &boolArg

	return builder
}

func (builder *BuilderQueryEditorOperatorTypeBuilder) Float64(float64Arg float64) *BuilderQueryEditorOperatorTypeBuilder {
	builder.internal.Float64 = &float64Arg

	return builder
}

func (builder *BuilderQueryEditorOperatorTypeBuilder) SelectableValue(selectableValue cog.Builder[SelectableValue]) *BuilderQueryEditorOperatorTypeBuilder {
	selectableValueResource, err := selectableValue.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.SelectableValue = &selectableValueResource

	return builder
}

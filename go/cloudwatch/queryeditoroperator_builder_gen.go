// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorOperator] = (*QueryEditorOperatorBuilder)(nil)

// TS type is QueryEditorOperator<T extends QueryEditorOperatorValueType>, extended in veneer
type QueryEditorOperatorBuilder struct {
	internal *QueryEditorOperator
	errors   cog.BuildErrors
}

func NewQueryEditorOperatorBuilder() *QueryEditorOperatorBuilder {
	resource := NewQueryEditorOperator()
	builder := &QueryEditorOperatorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryEditorOperatorBuilder) Build() (QueryEditorOperator, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorOperator{}, err
	}

	if len(builder.errors) > 0 {
		return QueryEditorOperator{}, cog.MakeBuildErrors("cloudwatch.queryEditorOperator", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryEditorOperatorBuilder) Name(name string) *QueryEditorOperatorBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *QueryEditorOperatorBuilder) Value(operatorTypes []QueryEditorOperatorType) *QueryEditorOperatorBuilder {
	if builder.internal.Value == nil {
		builder.internal.Value = NewStringOrBoolOrInt64OrArrayOfQueryEditorOperatorType()
	}
	builder.internal.Value.ArrayOfQueryEditorOperatorType = operatorTypes

	return builder
}

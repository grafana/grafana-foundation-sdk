// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorGroupByExpression] = (*QueryEditorGroupByExpressionBuilder)(nil)

type QueryEditorGroupByExpressionBuilder struct {
	internal *QueryEditorGroupByExpression
	errors   cog.BuildErrors
}

func NewQueryEditorGroupByExpressionBuilder() *QueryEditorGroupByExpressionBuilder {
	resource := NewQueryEditorGroupByExpression()
	builder := &QueryEditorGroupByExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryEditorGroupByExpressionBuilder) Build() (QueryEditorGroupByExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorGroupByExpression{}, err
	}

	if len(builder.errors) > 0 {
		return QueryEditorGroupByExpression{}, cog.MakeBuildErrors("cloudwatch.queryEditorGroupByExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryEditorGroupByExpressionBuilder) RecordError(path string, err error) *QueryEditorGroupByExpressionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryEditorGroupByExpressionBuilder) Property(property cog.Builder[QueryEditorProperty]) *QueryEditorGroupByExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

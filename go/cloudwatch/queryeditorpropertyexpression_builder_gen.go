// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorPropertyExpression] = (*QueryEditorPropertyExpressionBuilder)(nil)

type QueryEditorPropertyExpressionBuilder struct {
	internal *QueryEditorPropertyExpression
	errors   cog.BuildErrors
}

func NewQueryEditorPropertyExpressionBuilder() *QueryEditorPropertyExpressionBuilder {
	resource := NewQueryEditorPropertyExpression()
	builder := &QueryEditorPropertyExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryEditorPropertyExpressionBuilder) Build() (QueryEditorPropertyExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorPropertyExpression{}, err
	}

	if len(builder.errors) > 0 {
		return QueryEditorPropertyExpression{}, cog.MakeBuildErrors("cloudwatch.queryEditorPropertyExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryEditorPropertyExpressionBuilder) RecordError(path string, err error) *QueryEditorPropertyExpressionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryEditorPropertyExpressionBuilder) Property(property cog.Builder[QueryEditorProperty]) *QueryEditorPropertyExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorOperatorExpression] = (*QueryEditorOperatorExpressionBuilder)(nil)

type QueryEditorOperatorExpressionBuilder struct {
	internal *QueryEditorOperatorExpression
	errors   cog.BuildErrors
}

func NewQueryEditorOperatorExpressionBuilder() *QueryEditorOperatorExpressionBuilder {
	resource := NewQueryEditorOperatorExpression()
	builder := &QueryEditorOperatorExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *QueryEditorOperatorExpressionBuilder) Build() (QueryEditorOperatorExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorOperatorExpression{}, err
	}

	if len(builder.errors) > 0 {
		return QueryEditorOperatorExpression{}, cog.MakeBuildErrors("cloudwatch.queryEditorOperatorExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryEditorOperatorExpressionBuilder) Property(property cog.Builder[QueryEditorProperty]) *QueryEditorOperatorExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

// TS type is operator: QueryEditorOperator<QueryEditorOperatorValueType>, extended in veneer
func (builder *QueryEditorOperatorExpressionBuilder) Operator(operator cog.Builder[QueryEditorOperator]) *QueryEditorOperatorExpressionBuilder {
	operatorResource, err := operator.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Operator = operatorResource

	return builder
}

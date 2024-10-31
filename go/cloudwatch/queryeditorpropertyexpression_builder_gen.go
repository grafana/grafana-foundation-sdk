// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorPropertyExpression] = (*QueryEditorPropertyExpressionBuilder)(nil)

type QueryEditorPropertyExpressionBuilder struct {
	internal *QueryEditorPropertyExpression
	errors   map[string]cog.BuildErrors
}

func NewQueryEditorPropertyExpressionBuilder() *QueryEditorPropertyExpressionBuilder {
	resource := &QueryEditorPropertyExpression{}
	builder := &QueryEditorPropertyExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "property"

	return builder
}

func (builder *QueryEditorPropertyExpressionBuilder) Build() (QueryEditorPropertyExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryEditorPropertyExpression{}, err
	}

	return *builder.internal, nil
}

func (builder *QueryEditorPropertyExpressionBuilder) Property(property cog.Builder[QueryEditorProperty]) *QueryEditorPropertyExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors["property"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

func (builder *QueryEditorPropertyExpressionBuilder) applyDefaults() {
}

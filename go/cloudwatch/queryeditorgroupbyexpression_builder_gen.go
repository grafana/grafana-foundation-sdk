// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryEditorGroupByExpression] = (*QueryEditorGroupByExpressionBuilder)(nil)

type QueryEditorGroupByExpressionBuilder struct {
	internal *QueryEditorGroupByExpression
	errors   map[string]cog.BuildErrors
}

func NewQueryEditorGroupByExpressionBuilder() *QueryEditorGroupByExpressionBuilder {
	resource := &QueryEditorGroupByExpression{}
	builder := &QueryEditorGroupByExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "groupBy"

	return builder
}

func (builder *QueryEditorGroupByExpressionBuilder) Build() (QueryEditorGroupByExpression, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("QueryEditorGroupByExpression", err)...)
	}

	if len(errs) != 0 {
		return QueryEditorGroupByExpression{}, errs
	}

	return *builder.internal, nil
}

func (builder *QueryEditorGroupByExpressionBuilder) Property(property cog.Builder[QueryEditorProperty]) *QueryEditorGroupByExpressionBuilder {
	propertyResource, err := property.Build()
	if err != nil {
		builder.errors["property"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Property = propertyResource

	return builder
}

func (builder *QueryEditorGroupByExpressionBuilder) applyDefaults() {
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[Expression] = (*ExpressionBuilder)(nil)

type ExpressionBuilder struct {
	internal *Expression
	errors   cog.BuildErrors
}

func NewExpressionBuilder() *ExpressionBuilder {
	resource := NewExpression()
	builder := &ExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExpressionBuilder) Build() (Expression, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expression{}, err
	}

	if len(builder.errors) > 0 {
		return Expression{}, cog.MakeBuildErrors("alertingv1beta1.expression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExpressionBuilder) RecordError(path string, err error) *ExpressionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The type of query if this is a query expression
func (builder *ExpressionBuilder) QueryType(queryType string) *ExpressionBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

func (builder *ExpressionBuilder) RelativeTimeRange(relativeTimeRange cog.Builder[RelativeTimeRange]) *ExpressionBuilder {
	relativeTimeRangeResource, err := relativeTimeRange.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.RelativeTimeRange = &relativeTimeRangeResource

	return builder
}

// The UID of the datasource to run this expression against. If omitted, the expression will be run against the `__expr__` datasource
func (builder *ExpressionBuilder) DatasourceUID(datasourceUID DatasourceUID) *ExpressionBuilder {
	builder.internal.DatasourceUID = &datasourceUID

	return builder
}

func (builder *ExpressionBuilder) Model(model any) *ExpressionBuilder {
	builder.internal.Model = model

	return builder
}

// Used to mark the expression to be used as the final source for the rule evaluation
// Only one expression in a rule can be marked as the source
// For AlertRules, this is the expression that will be evaluated against the alerting condition
// For RecordingRules, this is the expression that will be recorded
func (builder *ExpressionBuilder) Source(source bool) *ExpressionBuilder {
	builder.internal.Source = &source

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SQLExpression] = (*SQLExpressionBuilder)(nil)

type SQLExpressionBuilder struct {
	internal *SQLExpression
	errors   map[string]cog.BuildErrors
}

func NewSQLExpressionBuilder() *SQLExpressionBuilder {
	resource := &SQLExpression{}
	builder := &SQLExpressionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *SQLExpressionBuilder) Build() (SQLExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return SQLExpression{}, err
	}

	return *builder.internal, nil
}

// SELECT part of the SQL expression
func (builder *SQLExpressionBuilder) Select(selectArg cog.Builder[QueryEditorFunctionExpression]) *SQLExpressionBuilder {
	selectArgResource, err := selectArg.Build()
	if err != nil {
		builder.errors["select"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Select = &selectArgResource

	return builder
}

// FROM part of the SQL expression
func (builder *SQLExpressionBuilder) From(from QueryEditorPropertyExpressionOrQueryEditorFunctionExpression) *SQLExpressionBuilder {
	builder.internal.From = &from

	return builder
}

// WHERE part of the SQL expression
func (builder *SQLExpressionBuilder) Where(where cog.Builder[QueryEditorArrayExpression]) *SQLExpressionBuilder {
	whereResource, err := where.Build()
	if err != nil {
		builder.errors["where"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Where = &whereResource

	return builder
}

// GROUP BY part of the SQL expression
func (builder *SQLExpressionBuilder) GroupBy(groupBy cog.Builder[QueryEditorArrayExpression]) *SQLExpressionBuilder {
	groupByResource, err := groupBy.Build()
	if err != nil {
		builder.errors["groupBy"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.GroupBy = &groupByResource

	return builder
}

// ORDER BY part of the SQL expression
func (builder *SQLExpressionBuilder) OrderBy(orderBy cog.Builder[QueryEditorFunctionExpression]) *SQLExpressionBuilder {
	orderByResource, err := orderBy.Build()
	if err != nil {
		builder.errors["orderBy"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.OrderBy = &orderByResource

	return builder
}

// The sort order of the SQL expression, `ASC` or `DESC`
func (builder *SQLExpressionBuilder) OrderByDirection(orderByDirection string) *SQLExpressionBuilder {
	builder.internal.OrderByDirection = &orderByDirection

	return builder
}

// LIMIT part of the SQL expression
func (builder *SQLExpressionBuilder) Limit(limit int64) *SQLExpressionBuilder {
	builder.internal.Limit = &limit

	return builder
}

func (builder *SQLExpressionBuilder) applyDefaults() {
}

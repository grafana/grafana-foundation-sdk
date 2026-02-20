// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SQLExpression] = (*SQLExpressionBuilder)(nil)

type SQLExpressionBuilder struct {
	internal *SQLExpression
	errors   cog.BuildErrors
}

func NewSQLExpressionBuilder() *SQLExpressionBuilder {
	resource := NewSQLExpression()
	builder := &SQLExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SQLExpressionBuilder) Build() (SQLExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return SQLExpression{}, err
	}

	if len(builder.errors) > 0 {
		return SQLExpression{}, cog.MakeBuildErrors("bigquery.sQLExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SQLExpressionBuilder) RecordError(path string, err error) *SQLExpressionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SQLExpressionBuilder) Columns(columns []cog.Builder[QueryEditorFunctionExpression]) *SQLExpressionBuilder {
	columnsResources := make([]QueryEditorFunctionExpression, 0, len(columns))
	for _, r1 := range columns {
		columnsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		columnsResources = append(columnsResources, columnsDepth1)
	}
	builder.internal.Columns = columnsResources

	return builder
}

func (builder *SQLExpressionBuilder) From(from string) *SQLExpressionBuilder {
	builder.internal.From = &from

	return builder
}

// whereJsonTree?: _
func (builder *SQLExpressionBuilder) WhereString(whereString string) *SQLExpressionBuilder {
	builder.internal.WhereString = &whereString

	return builder
}

func (builder *SQLExpressionBuilder) GroupBy(groupBy []cog.Builder[QueryEditorGroupByExpression]) *SQLExpressionBuilder {
	groupByResources := make([]QueryEditorGroupByExpression, 0, len(groupBy))
	for _, r1 := range groupBy {
		groupByDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		groupByResources = append(groupByResources, groupByDepth1)
	}
	builder.internal.GroupBy = groupByResources

	return builder
}

func (builder *SQLExpressionBuilder) OrderBy(orderBy cog.Builder[QueryEditorPropertyExpression]) *SQLExpressionBuilder {
	orderByResource, err := orderBy.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.OrderBy = &orderByResource

	return builder
}

func (builder *SQLExpressionBuilder) OrderByDirection(orderByDirection OrderByDirection) *SQLExpressionBuilder {
	builder.internal.OrderByDirection = &orderByDirection

	return builder
}

func (builder *SQLExpressionBuilder) Limit(limit int64) *SQLExpressionBuilder {
	builder.internal.Limit = &limit

	return builder
}

func (builder *SQLExpressionBuilder) Offset(offset int64) *SQLExpressionBuilder {
	builder.internal.Offset = &offset

	return builder
}

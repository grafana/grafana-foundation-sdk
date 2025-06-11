// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[BuilderQueryExpression] = (*BuilderQueryExpressionBuilder)(nil)

type BuilderQueryExpressionBuilder struct {
	internal *BuilderQueryExpression
	errors   cog.BuildErrors
}

func NewBuilderQueryExpressionBuilder() *BuilderQueryExpressionBuilder {
	resource := NewBuilderQueryExpression()
	builder := &BuilderQueryExpressionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *BuilderQueryExpressionBuilder) Build() (BuilderQueryExpression, error) {
	if err := builder.internal.Validate(); err != nil {
		return BuilderQueryExpression{}, err
	}

	if len(builder.errors) > 0 {
		return BuilderQueryExpression{}, cog.MakeBuildErrors("azuremonitor.builderQueryExpression", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *BuilderQueryExpressionBuilder) From(from cog.Builder[BuilderQueryEditorPropertyExpression]) *BuilderQueryExpressionBuilder {
	fromResource, err := from.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.From = &fromResource

	return builder
}

func (builder *BuilderQueryExpressionBuilder) Columns(columns cog.Builder[BuilderQueryEditorColumnsExpression]) *BuilderQueryExpressionBuilder {
	columnsResource, err := columns.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Columns = &columnsResource

	return builder
}

func (builder *BuilderQueryExpressionBuilder) Where(where cog.Builder[BuilderQueryEditorWhereExpressionArray]) *BuilderQueryExpressionBuilder {
	whereResource, err := where.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Where = &whereResource

	return builder
}

func (builder *BuilderQueryExpressionBuilder) Reduce(reduce cog.Builder[BuilderQueryEditorReduceExpressionArray]) *BuilderQueryExpressionBuilder {
	reduceResource, err := reduce.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Reduce = &reduceResource

	return builder
}

func (builder *BuilderQueryExpressionBuilder) GroupBy(groupBy cog.Builder[BuilderQueryEditorGroupByExpressionArray]) *BuilderQueryExpressionBuilder {
	groupByResource, err := groupBy.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.GroupBy = &groupByResource

	return builder
}

func (builder *BuilderQueryExpressionBuilder) Limit(limit int64) *BuilderQueryExpressionBuilder {
	builder.internal.Limit = &limit

	return builder
}

func (builder *BuilderQueryExpressionBuilder) OrderBy(orderBy cog.Builder[BuilderQueryEditorOrderByExpressionArray]) *BuilderQueryExpressionBuilder {
	orderByResource, err := orderBy.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.OrderBy = &orderByResource

	return builder
}

func (builder *BuilderQueryExpressionBuilder) FuzzySearch(fuzzySearch cog.Builder[BuilderQueryEditorWhereExpressionArray]) *BuilderQueryExpressionBuilder {
	fuzzySearchResource, err := fuzzySearch.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.FuzzySearch = &fuzzySearchResource

	return builder
}

func (builder *BuilderQueryExpressionBuilder) TimeFilter(timeFilter cog.Builder[BuilderQueryEditorWhereExpressionArray]) *BuilderQueryExpressionBuilder {
	timeFilterResource, err := timeFilter.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TimeFilter = &timeFilterResource

	return builder
}

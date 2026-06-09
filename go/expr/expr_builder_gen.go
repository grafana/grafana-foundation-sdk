// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[variants.Dataquery] = (*ExprBuilder)(nil)

type ExprBuilder struct {
	internal *Expr
	errors   cog.BuildErrors
}

func NewExprBuilder() *ExprBuilder {
	resource := NewExpr()
	builder := &ExprBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExprBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Expr{}, err
	}

	if len(builder.errors) > 0 {
		return Expr{}, cog.MakeBuildErrors("expr.expr", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExprBuilder) RecordError(path string, err error) *ExprBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ExprBuilder) TypeMath(typeMath cog.Builder[TypeMath]) *ExprBuilder {
	typeMathResource, err := typeMath.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TypeMath = &typeMathResource

	return builder
}

func (builder *ExprBuilder) TypeReduce(typeReduce cog.Builder[TypeReduce]) *ExprBuilder {
	typeReduceResource, err := typeReduce.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TypeReduce = &typeReduceResource

	return builder
}

func (builder *ExprBuilder) TypeResample(typeResample cog.Builder[TypeResample]) *ExprBuilder {
	typeResampleResource, err := typeResample.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TypeResample = &typeResampleResource

	return builder
}

func (builder *ExprBuilder) TypeClassicConditions(typeClassicConditions cog.Builder[TypeClassicConditions]) *ExprBuilder {
	typeClassicConditionsResource, err := typeClassicConditions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TypeClassicConditions = &typeClassicConditionsResource

	return builder
}

func (builder *ExprBuilder) TypeThreshold(typeThreshold cog.Builder[TypeThreshold]) *ExprBuilder {
	typeThresholdResource, err := typeThreshold.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TypeThreshold = &typeThresholdResource

	return builder
}

func (builder *ExprBuilder) TypeSql(typeSql cog.Builder[TypeSql]) *ExprBuilder {
	typeSqlResource, err := typeSql.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.TypeSql = &typeSqlResource

	return builder
}

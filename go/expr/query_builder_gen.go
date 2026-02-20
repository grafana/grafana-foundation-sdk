// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2beta1 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2beta1"
)

var _ cog.Builder[dashboardv2beta1.DataQueryKind] = (*QueryBuilder)(nil)

type QueryBuilder struct {
	internal *dashboardv2beta1.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryBuilder() *QueryBuilder {
	resource := dashboardv2beta1.NewDataQueryKind()
	builder := &QueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "__expr__"

	return builder
}

func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.DataQueryKind{}, cog.MakeBuildErrors("expr.query", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryBuilder) RecordError(path string, err error) *QueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryBuilder) Version(version string) *QueryBuilder {
	builder.internal.Version = version

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryBuilder) Datasource(datasource cog.Builder[dashboardv2beta1.Dashboardv2beta1DataQueryKindDatasource]) *QueryBuilder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

func (builder *QueryBuilder) TypeMath(typeMath cog.Builder[TypeMath]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewExpr()
	}
	typeMathResource, err := typeMath.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Expr).TypeMath = &typeMathResource

	return builder
}

func (builder *QueryBuilder) TypeReduce(typeReduce cog.Builder[TypeReduce]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewExpr()
	}
	typeReduceResource, err := typeReduce.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Expr).TypeReduce = &typeReduceResource

	return builder
}

func (builder *QueryBuilder) TypeResample(typeResample cog.Builder[TypeResample]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewExpr()
	}
	typeResampleResource, err := typeResample.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Expr).TypeResample = &typeResampleResource

	return builder
}

func (builder *QueryBuilder) TypeClassicConditions(typeClassicConditions cog.Builder[TypeClassicConditions]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewExpr()
	}
	typeClassicConditionsResource, err := typeClassicConditions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Expr).TypeClassicConditions = &typeClassicConditionsResource

	return builder
}

func (builder *QueryBuilder) TypeThreshold(typeThreshold cog.Builder[TypeThreshold]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewExpr()
	}
	typeThresholdResource, err := typeThreshold.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Expr).TypeThreshold = &typeThresholdResource

	return builder
}

func (builder *QueryBuilder) TypeSql(typeSql cog.Builder[TypeSql]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewExpr()
	}
	typeSqlResource, err := typeSql.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Expr).TypeSql = &typeSqlResource

	return builder
}

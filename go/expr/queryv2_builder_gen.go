// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package expr

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	dashboardv2 "github.com/grafana/grafana-foundation-sdk/go/dashboardv2"
)

var _ cog.Builder[dashboardv2.DataQueryKind] = (*QueryV2Builder)(nil)

type QueryV2Builder struct {
	internal *dashboardv2.DataQueryKind
	errors   cog.BuildErrors
}

func NewQueryV2Builder() *QueryV2Builder {
	resource := dashboardv2.NewDataQueryKind()
	builder := &QueryV2Builder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"
	builder.internal.Group = "__expr__"

	return builder
}

func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2.DataQueryKind{}, cog.MakeBuildErrors("expr.queryV2", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryV2Builder) RecordError(path string, err error) *QueryV2Builder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *QueryV2Builder) Version(version string) *QueryV2Builder {
	builder.internal.Version = version

	return builder
}

func (builder *QueryV2Builder) Labels(labels map[string]string) *QueryV2Builder {
	builder.internal.Labels = labels

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *QueryV2Builder) Datasource(datasource cog.Builder[dashboardv2.Dashboardv2DataQueryKindDatasource]) *QueryV2Builder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

func (builder *QueryV2Builder) TypeMath(typeMath cog.Builder[TypeMath]) *QueryV2Builder {
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

func (builder *QueryV2Builder) TypeReduce(typeReduce cog.Builder[TypeReduce]) *QueryV2Builder {
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

func (builder *QueryV2Builder) TypeResample(typeResample cog.Builder[TypeResample]) *QueryV2Builder {
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

func (builder *QueryV2Builder) TypeClassicConditions(typeClassicConditions cog.Builder[TypeClassicConditions]) *QueryV2Builder {
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

func (builder *QueryV2Builder) TypeThreshold(typeThreshold cog.Builder[TypeThreshold]) *QueryV2Builder {
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

func (builder *QueryV2Builder) TypeSql(typeSql cog.Builder[TypeSql]) *QueryV2Builder {
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

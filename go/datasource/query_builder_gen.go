// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package datasource

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
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
	builder.internal.Group = "datasource"

	return builder
}

func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.DataQueryKind{}, cog.MakeBuildErrors("datasource.query", builder.errors)
	}

	return *builder.internal, nil
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

// Panel ID from wich the queries will be reused.
func (builder *QueryBuilder) PanelId(panelId uint32) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).PanelId = panelId

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *QueryBuilder) RefId(refId string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RefId = &refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *QueryBuilder) Hide(hide bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *QueryBuilder) QueryType(queryType string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).QueryType = &queryType

	return builder
}

func (builder *QueryBuilder) WithTransforms(withTransforms bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).WithTransforms = withTransforms

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *QueryBuilder) OldDatasource(datasource common.DataSourceRef) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Datasource = &datasource

	return builder
}

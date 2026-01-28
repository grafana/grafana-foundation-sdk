// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

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
	builder.internal.Group = "grafana-bigquery-datasource"

	return builder
}

func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.DataQueryKind{}, cog.MakeBuildErrors("bigquery.query", builder.errors)
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

func (builder *QueryBuilder) Dataset(dataset string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Dataset = &dataset

	return builder
}

func (builder *QueryBuilder) Table(table string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Table = &table

	return builder
}

func (builder *QueryBuilder) Project(project string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Project = &project

	return builder
}

func (builder *QueryBuilder) Format(format QueryFormat) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Format = format

	return builder
}

func (builder *QueryBuilder) RawQuery(rawQuery bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RawQuery = &rawQuery

	return builder
}

func (builder *QueryBuilder) RawSql(rawSql string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).RawSql = rawSql

	return builder
}

func (builder *QueryBuilder) Location(location string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Location = &location

	return builder
}

func (builder *QueryBuilder) Partitioned(partitioned bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Partitioned = &partitioned

	return builder
}

func (builder *QueryBuilder) PartitionedField(partitionedField string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).PartitionedField = &partitionedField

	return builder
}

func (builder *QueryBuilder) ConvertToUTC(convertToUTC bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).ConvertToUTC = &convertToUTC

	return builder
}

func (builder *QueryBuilder) Sharded(sharded bool) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).Sharded = &sharded

	return builder
}

func (builder *QueryBuilder) QueryPriority(queryPriority QueryPriority) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).QueryPriority = &queryPriority

	return builder
}

func (builder *QueryBuilder) TimeShift(timeShift string) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).TimeShift = &timeShift

	return builder
}

func (builder *QueryBuilder) EditorMode(editorMode EditorMode) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	builder.internal.Spec.(*Dataquery).EditorMode = &editorMode

	return builder
}

func (builder *QueryBuilder) Sql(sql cog.Builder[SQLExpression]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewDataquery()
	}
	sqlResource, err := sql.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*Dataquery).Sql = &sqlResource

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

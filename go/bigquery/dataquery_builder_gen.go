// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package bigquery

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*DataqueryBuilder)(nil)

// Manually converted from https://github.com/grafana/google-bigquery-datasource/blob/18680e42ba557791d109c7c540c2c3f2647592f0/src/types.ts#L75
type DataqueryBuilder struct {
	internal *Dataquery
	errors   map[string]cog.BuildErrors
}

func NewDataqueryBuilder() *DataqueryBuilder {
	resource := NewDataquery()
	builder := &DataqueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *DataqueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dataquery{}, err
	}

	return *builder.internal, nil
}

func (builder *DataqueryBuilder) Dataset(dataset string) *DataqueryBuilder {
	builder.internal.Dataset = &dataset

	return builder
}

func (builder *DataqueryBuilder) Table(table string) *DataqueryBuilder {
	builder.internal.Table = &table

	return builder
}

func (builder *DataqueryBuilder) Project(project string) *DataqueryBuilder {
	builder.internal.Project = &project

	return builder
}

func (builder *DataqueryBuilder) Format(format QueryFormat) *DataqueryBuilder {
	builder.internal.Format = format

	return builder
}

func (builder *DataqueryBuilder) RawQuery(rawQuery bool) *DataqueryBuilder {
	builder.internal.RawQuery = &rawQuery

	return builder
}

func (builder *DataqueryBuilder) RawSql(rawSql string) *DataqueryBuilder {
	builder.internal.RawSql = rawSql

	return builder
}

func (builder *DataqueryBuilder) Location(location string) *DataqueryBuilder {
	builder.internal.Location = &location

	return builder
}

func (builder *DataqueryBuilder) Partitioned(partitioned bool) *DataqueryBuilder {
	builder.internal.Partitioned = &partitioned

	return builder
}

func (builder *DataqueryBuilder) PartitionedField(partitionedField string) *DataqueryBuilder {
	builder.internal.PartitionedField = &partitionedField

	return builder
}

func (builder *DataqueryBuilder) ConvertToUTC(convertToUTC bool) *DataqueryBuilder {
	builder.internal.ConvertToUTC = &convertToUTC

	return builder
}

func (builder *DataqueryBuilder) Sharded(sharded bool) *DataqueryBuilder {
	builder.internal.Sharded = &sharded

	return builder
}

func (builder *DataqueryBuilder) QueryPriority(queryPriority QueryPriority) *DataqueryBuilder {
	builder.internal.QueryPriority = &queryPriority

	return builder
}

func (builder *DataqueryBuilder) TimeShift(timeShift string) *DataqueryBuilder {
	builder.internal.TimeShift = &timeShift

	return builder
}

func (builder *DataqueryBuilder) EditorMode(editorMode EditorMode) *DataqueryBuilder {
	builder.internal.EditorMode = &editorMode

	return builder
}

func (builder *DataqueryBuilder) Sql(sql cog.Builder[SQLExpression]) *DataqueryBuilder {
	sqlResource, err := sql.Build()
	if err != nil {
		builder.errors["sql"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Sql = &sqlResource

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *DataqueryBuilder) Datasource(datasource dashboard.DataSourceRef) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

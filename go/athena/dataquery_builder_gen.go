// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package athena

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[variants.Dataquery] = (*DataqueryBuilder)(nil)

// Manually converted from https://github.com/grafana/athena-datasource/blob/57ad707147b7a11e9a521a836d6bf9799877e0e3/src/types.ts
type DataqueryBuilder struct {
	internal *Dataquery
	errors   cog.BuildErrors
}

func NewDataqueryBuilder() *DataqueryBuilder {
	resource := NewDataquery()
	builder := &DataqueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DataqueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dataquery{}, err
	}

	if len(builder.errors) > 0 {
		return Dataquery{}, cog.MakeBuildErrors("athena.dataquery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DataqueryBuilder) RecordError(path string, err error) *DataqueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *DataqueryBuilder) Format(format FormatOptions) *DataqueryBuilder {
	builder.internal.Format = format

	return builder
}

func (builder *DataqueryBuilder) ConnectionArgs(connectionArgs cog.Builder[ConnectionArgs]) *DataqueryBuilder {
	connectionArgsResource, err := connectionArgs.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.ConnectionArgs = connectionArgsResource

	return builder
}

func (builder *DataqueryBuilder) Table(table string) *DataqueryBuilder {
	builder.internal.Table = &table

	return builder
}

func (builder *DataqueryBuilder) Column(column string) *DataqueryBuilder {
	builder.internal.Column = &column

	return builder
}

func (builder *DataqueryBuilder) QueryID(queryID string) *DataqueryBuilder {
	builder.internal.QueryID = &queryID

	return builder
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = &refId

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

func (builder *DataqueryBuilder) RawSQL(rawSQL string) *DataqueryBuilder {
	builder.internal.RawSQL = rawSQL

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *DataqueryBuilder) Datasource(datasource common.DataSourceRef) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

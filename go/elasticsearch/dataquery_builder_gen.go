// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*DataqueryBuilder)(nil)

type DataqueryBuilder struct {
	internal *Dataquery
	errors   map[string]cog.BuildErrors
}

func NewDataqueryBuilder() *DataqueryBuilder {
	resource := &Dataquery{}
	builder := &DataqueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *DataqueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return Dataquery{}, err
	}

	return *builder.internal, nil
}

// Alias pattern
func (builder *DataqueryBuilder) Alias(alias string) *DataqueryBuilder {
	builder.internal.Alias = &alias

	return builder
}

// Lucene query
func (builder *DataqueryBuilder) Query(query string) *DataqueryBuilder {
	builder.internal.Query = &query

	return builder
}

// Name of time field
func (builder *DataqueryBuilder) TimeField(timeField string) *DataqueryBuilder {
	builder.internal.TimeField = &timeField

	return builder
}

// List of bucket aggregations
func (builder *DataqueryBuilder) BucketAggs(bucketAggs []BucketAggregation) *DataqueryBuilder {
	builder.internal.BucketAggs = bucketAggs

	return builder
}

// List of metric aggregations
func (builder *DataqueryBuilder) Metrics(metrics []MetricAggregation) *DataqueryBuilder {
	builder.internal.Metrics = metrics

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

func (builder *DataqueryBuilder) applyDefaults() {
}

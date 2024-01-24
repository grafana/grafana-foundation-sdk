// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	cogvariants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[cogvariants.Dataquery] = (*DataqueryBuilder)(nil)

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

func (builder *DataqueryBuilder) Build() (cogvariants.Dataquery, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("Dataquery", err)...)
	}

	if len(errs) != 0 {
		return Dataquery{}, errs
	}

	return *builder.internal, nil
}

func (builder *DataqueryBuilder) Alias(alias string) *DataqueryBuilder {
	builder.internal.Alias = &alias

	return builder
}

func (builder *DataqueryBuilder) Query(query string) *DataqueryBuilder {
	builder.internal.Query = &query

	return builder
}

func (builder *DataqueryBuilder) TimeField(timeField string) *DataqueryBuilder {
	builder.internal.TimeField = &timeField

	return builder
}

func (builder *DataqueryBuilder) BucketAggs(bucketAggs []BucketAggregation) *DataqueryBuilder {
	builder.internal.BucketAggs = bucketAggs

	return builder
}

func (builder *DataqueryBuilder) Metrics(metrics []MetricAggregation) *DataqueryBuilder {
	builder.internal.Metrics = metrics

	return builder
}

func (builder *DataqueryBuilder) RefId(refId string) *DataqueryBuilder {
	builder.internal.RefId = &refId

	return builder
}

func (builder *DataqueryBuilder) Hide(hide bool) *DataqueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

func (builder *DataqueryBuilder) QueryType(queryType string) *DataqueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

func (builder *DataqueryBuilder) Datasource(datasource any) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

func (builder *DataqueryBuilder) applyDefaults() {
}

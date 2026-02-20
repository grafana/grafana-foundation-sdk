// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package grafanapyroscope

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	common "github.com/grafana/grafana-foundation-sdk/go/common"
)

var _ cog.Builder[variants.Dataquery] = (*DataqueryBuilder)(nil)

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
		return Dataquery{}, cog.MakeBuildErrors("grafanapyroscope.dataquery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DataqueryBuilder) RecordError(path string, err error) *DataqueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Specifies the query label selectors.
func (builder *DataqueryBuilder) LabelSelector(labelSelector string) *DataqueryBuilder {
	builder.internal.LabelSelector = labelSelector

	return builder
}

// Specifies the query span selectors.
func (builder *DataqueryBuilder) SpanSelector(spanSelector []string) *DataqueryBuilder {
	builder.internal.SpanSelector = spanSelector

	return builder
}

// Specifies the type of profile to query.
func (builder *DataqueryBuilder) ProfileTypeId(profileTypeId string) *DataqueryBuilder {
	builder.internal.ProfileTypeId = profileTypeId

	return builder
}

// Allows to group the results.
func (builder *DataqueryBuilder) GroupBy(groupBy []string) *DataqueryBuilder {
	builder.internal.GroupBy = groupBy

	return builder
}

// Sets the maximum number of time series.
func (builder *DataqueryBuilder) Limit(limit int64) *DataqueryBuilder {
	builder.internal.Limit = &limit

	return builder
}

// Sets the maximum number of nodes in the flamegraph.
func (builder *DataqueryBuilder) MaxNodes(maxNodes int64) *DataqueryBuilder {
	builder.internal.MaxNodes = &maxNodes

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

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *DataqueryBuilder) Datasource(datasource common.DataSourceRef) *DataqueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

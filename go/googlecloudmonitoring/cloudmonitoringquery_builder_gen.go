// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
	dashboard "github.com/grafana/grafana-foundation-sdk/go/dashboard"
)

var _ cog.Builder[variants.Dataquery] = (*CloudMonitoringQueryBuilder)(nil)

type CloudMonitoringQueryBuilder struct {
	internal *CloudMonitoringQuery
	errors   map[string]cog.BuildErrors
}

func NewCloudMonitoringQueryBuilder() *CloudMonitoringQueryBuilder {
	resource := &CloudMonitoringQuery{}
	builder := &CloudMonitoringQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *CloudMonitoringQueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return CloudMonitoringQuery{}, err
	}

	return *builder.internal, nil
}

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *CloudMonitoringQueryBuilder) RefId(refId string) *CloudMonitoringQueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// true if query is disabled (ie should not be returned to the dashboard)
// Note this does not always imply that the query should not be executed since
// the results from a hidden query may be used as the input to other queries (SSE etc)
func (builder *CloudMonitoringQueryBuilder) Hide(hide bool) *CloudMonitoringQueryBuilder {
	builder.internal.Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *CloudMonitoringQueryBuilder) QueryType(queryType string) *CloudMonitoringQueryBuilder {
	builder.internal.QueryType = &queryType

	return builder
}

// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
func (builder *CloudMonitoringQueryBuilder) AliasBy(aliasBy string) *CloudMonitoringQueryBuilder {
	builder.internal.AliasBy = &aliasBy

	return builder
}

// GCM query type.
// queryType: #QueryType
// Time Series List sub-query properties.
func (builder *CloudMonitoringQueryBuilder) TimeSeriesList(timeSeriesList cog.Builder[TimeSeriesList]) *CloudMonitoringQueryBuilder {
	timeSeriesListResource, err := timeSeriesList.Build()
	if err != nil {
		builder.errors["timeSeriesList"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.TimeSeriesList = &timeSeriesListResource

	return builder
}

// Time Series sub-query properties.
func (builder *CloudMonitoringQueryBuilder) TimeSeriesQuery(timeSeriesQuery cog.Builder[TimeSeriesQuery]) *CloudMonitoringQueryBuilder {
	timeSeriesQueryResource, err := timeSeriesQuery.Build()
	if err != nil {
		builder.errors["timeSeriesQuery"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.TimeSeriesQuery = &timeSeriesQueryResource

	return builder
}

// SLO sub-query properties.
func (builder *CloudMonitoringQueryBuilder) SloQuery(sloQuery cog.Builder[SLOQuery]) *CloudMonitoringQueryBuilder {
	sloQueryResource, err := sloQuery.Build()
	if err != nil {
		builder.errors["sloQuery"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.SloQuery = &sloQueryResource

	return builder
}

// For mixed data sources the selected datasource is on the query level.
// For non mixed scenarios this is undefined.
// TODO find a better way to do this ^ that's friendly to schema
// TODO this shouldn't be unknown but DataSourceRef | null
func (builder *CloudMonitoringQueryBuilder) Datasource(datasource dashboard.DataSourceRef) *CloudMonitoringQueryBuilder {
	builder.internal.Datasource = &datasource

	return builder
}

// Time interval in milliseconds.
func (builder *CloudMonitoringQueryBuilder) IntervalMs(intervalMs float64) *CloudMonitoringQueryBuilder {
	builder.internal.IntervalMs = &intervalMs

	return builder
}

func (builder *CloudMonitoringQueryBuilder) applyDefaults() {
}

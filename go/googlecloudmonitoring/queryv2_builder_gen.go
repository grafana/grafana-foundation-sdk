// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

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
	builder.internal.Group = "cloud-monitoring"

	return builder
}

func (builder *QueryV2Builder) Build() (dashboardv2.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2.DataQueryKind{}, cog.MakeBuildErrors("googlecloudmonitoring.queryV2", builder.errors)
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

// A unique identifier for the query within the list of targets.
// In server side expressions, the refId is used as a variable name to identify results.
// By default, the UI will assign A->Z; however setting meaningful names may be useful.
func (builder *QueryV2Builder) RefId(refId string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	builder.internal.Spec.(*CloudMonitoringQuery).RefId = &refId

	return builder
}

// If hide is set to true, Grafana will filter out the response(s) associated with this query before returning it to the panel.
func (builder *QueryV2Builder) Hide(hide bool) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	builder.internal.Spec.(*CloudMonitoringQuery).Hide = &hide

	return builder
}

// Specify the query flavor
// TODO make this required and give it a default
func (builder *QueryV2Builder) QueryType(queryType string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	builder.internal.Spec.(*CloudMonitoringQuery).QueryType = &queryType

	return builder
}

// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
func (builder *QueryV2Builder) AliasBy(aliasBy string) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	builder.internal.Spec.(*CloudMonitoringQuery).AliasBy = &aliasBy

	return builder
}

// GCM query type.
// queryType: #QueryType
// Time Series List sub-query properties.
func (builder *QueryV2Builder) TimeSeriesList(timeSeriesList cog.Builder[TimeSeriesList]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	timeSeriesListResource, err := timeSeriesList.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*CloudMonitoringQuery).TimeSeriesList = &timeSeriesListResource

	return builder
}

// Time Series sub-query properties.
func (builder *QueryV2Builder) TimeSeriesQuery(timeSeriesQuery cog.Builder[TimeSeriesQuery]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	timeSeriesQueryResource, err := timeSeriesQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*CloudMonitoringQuery).TimeSeriesQuery = &timeSeriesQueryResource

	return builder
}

// SLO sub-query properties.
func (builder *QueryV2Builder) SloQuery(sloQuery cog.Builder[SLOQuery]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	sloQueryResource, err := sloQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*CloudMonitoringQuery).SloQuery = &sloQueryResource

	return builder
}

// PromQL sub-query properties.
func (builder *QueryV2Builder) PromQLQuery(promQLQuery cog.Builder[PromQLQuery]) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	promQLQueryResource, err := promQLQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*CloudMonitoringQuery).PromQLQuery = &promQLQueryResource

	return builder
}

// Time interval in milliseconds.
func (builder *QueryV2Builder) IntervalMs(intervalMs float64) *QueryV2Builder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudMonitoringQuery()
	}
	builder.internal.Spec.(*CloudMonitoringQuery).IntervalMs = &intervalMs

	return builder
}

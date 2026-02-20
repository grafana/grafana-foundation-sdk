// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
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
	builder.internal.Group = "cloudwatch"

	return builder
}

func (builder *QueryBuilder) Build() (dashboardv2beta1.DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return dashboardv2beta1.DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return dashboardv2beta1.DataQueryKind{}, cog.MakeBuildErrors("cloudwatch.query", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryBuilder) RecordError(path string, err error) *QueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
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

func (builder *QueryBuilder) CloudWatchMetricsQuery(cloudWatchMetricsQuery cog.Builder[CloudWatchMetricsQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudWatchQuery()
	}
	cloudWatchMetricsQueryResource, err := cloudWatchMetricsQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*CloudWatchQuery).CloudWatchMetricsQuery = &cloudWatchMetricsQueryResource

	return builder
}

func (builder *QueryBuilder) CloudWatchLogsQuery(cloudWatchLogsQuery cog.Builder[CloudWatchLogsQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudWatchQuery()
	}
	cloudWatchLogsQueryResource, err := cloudWatchLogsQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*CloudWatchQuery).CloudWatchLogsQuery = &cloudWatchLogsQueryResource

	return builder
}

func (builder *QueryBuilder) CloudWatchAnnotationQuery(cloudWatchAnnotationQuery cog.Builder[CloudWatchAnnotationQuery]) *QueryBuilder {
	if builder.internal.Spec == nil {
		builder.internal.Spec = NewCloudWatchQuery()
	}
	cloudWatchAnnotationQueryResource, err := cloudWatchAnnotationQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.(*CloudWatchQuery).CloudWatchAnnotationQuery = &cloudWatchAnnotationQueryResource

	return builder
}

// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package cloudwatch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
	variants "github.com/grafana/grafana-foundation-sdk/go/cog/variants"
)

var _ cog.Builder[variants.Dataquery] = (*MetricsQueryOrLogsQueryOrAnnotationQueryBuilder)(nil)

type MetricsQueryOrLogsQueryOrAnnotationQueryBuilder struct {
	internal *MetricsQueryOrLogsQueryOrAnnotationQuery
	errors   cog.BuildErrors
}

func NewMetricsQueryOrLogsQueryOrAnnotationQueryBuilder() *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder {
	resource := NewMetricsQueryOrLogsQueryOrAnnotationQuery()
	builder := &MetricsQueryOrLogsQueryOrAnnotationQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) Build() (variants.Dataquery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricsQueryOrLogsQueryOrAnnotationQuery{}, err
	}

	if len(builder.errors) > 0 {
		return MetricsQueryOrLogsQueryOrAnnotationQuery{}, cog.MakeBuildErrors("cloudwatch.metricsQueryOrLogsQueryOrAnnotationQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) RecordError(path string, err error) *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) MetricsQuery(metricsQuery cog.Builder[MetricsQuery]) *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder {
	metricsQueryResource, err := metricsQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.MetricsQuery = &metricsQueryResource

	return builder
}

func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) LogsQuery(logsQuery cog.Builder[LogsQuery]) *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder {
	logsQueryResource, err := logsQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.LogsQuery = &logsQueryResource

	return builder
}

func (builder *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder) AnnotationQuery(annotationQuery cog.Builder[AnnotationQuery]) *MetricsQueryOrLogsQueryOrAnnotationQueryBuilder {
	annotationQueryResource, err := annotationQuery.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.AnnotationQuery = &annotationQueryResource

	return builder
}

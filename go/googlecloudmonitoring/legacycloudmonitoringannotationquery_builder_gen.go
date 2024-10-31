// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LegacyCloudMonitoringAnnotationQuery] = (*LegacyCloudMonitoringAnnotationQueryBuilder)(nil)

// @deprecated Use AnnotationQuery instead. Legacy annotation query properties for migration purposes.
type LegacyCloudMonitoringAnnotationQueryBuilder struct {
	internal *LegacyCloudMonitoringAnnotationQuery
	errors   map[string]cog.BuildErrors
}

func NewLegacyCloudMonitoringAnnotationQueryBuilder() *LegacyCloudMonitoringAnnotationQueryBuilder {
	resource := &LegacyCloudMonitoringAnnotationQuery{}
	builder := &LegacyCloudMonitoringAnnotationQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) Build() (LegacyCloudMonitoringAnnotationQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return LegacyCloudMonitoringAnnotationQuery{}, err
	}

	return *builder.internal, nil
}

// GCP project to execute the query against.
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) ProjectName(projectName string) *LegacyCloudMonitoringAnnotationQueryBuilder {
	builder.internal.ProjectName = projectName

	return builder
}

func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) MetricType(metricType string) *LegacyCloudMonitoringAnnotationQueryBuilder {
	builder.internal.MetricType = metricType

	return builder
}

// Query refId.
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) RefId(refId string) *LegacyCloudMonitoringAnnotationQueryBuilder {
	builder.internal.RefId = refId

	return builder
}

// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) Filters(filters []string) *LegacyCloudMonitoringAnnotationQueryBuilder {
	builder.internal.Filters = filters

	return builder
}

func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) MetricKind(metricKind MetricKind) *LegacyCloudMonitoringAnnotationQueryBuilder {
	builder.internal.MetricKind = metricKind

	return builder
}

func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) ValueType(valueType string) *LegacyCloudMonitoringAnnotationQueryBuilder {
	builder.internal.ValueType = valueType

	return builder
}

// Annotation title.
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) Title(title string) *LegacyCloudMonitoringAnnotationQueryBuilder {
	builder.internal.Title = title

	return builder
}

// Annotation text.
func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) Text(text string) *LegacyCloudMonitoringAnnotationQueryBuilder {
	builder.internal.Text = text

	return builder
}

func (builder *LegacyCloudMonitoringAnnotationQueryBuilder) applyDefaults() {
}

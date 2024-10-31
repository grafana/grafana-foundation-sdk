// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricQuery] = (*MetricQueryBuilder)(nil)

// @deprecated This type is for migration purposes only. Replaced by TimeSeriesList Metric sub-query properties.
type MetricQueryBuilder struct {
	internal *MetricQuery
	errors   map[string]cog.BuildErrors
}

func NewMetricQueryBuilder() *MetricQueryBuilder {
	resource := &MetricQuery{}
	builder := &MetricQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MetricQueryBuilder) Build() (MetricQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricQuery{}, err
	}

	return *builder.internal, nil
}

// GCP project to execute the query against.
func (builder *MetricQueryBuilder) ProjectName(projectName string) *MetricQueryBuilder {
	builder.internal.ProjectName = projectName

	return builder
}

// Alignment function to be used. Defaults to ALIGN_MEAN.
func (builder *MetricQueryBuilder) PerSeriesAligner(perSeriesAligner string) *MetricQueryBuilder {
	builder.internal.PerSeriesAligner = &perSeriesAligner

	return builder
}

// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
func (builder *MetricQueryBuilder) AlignmentPeriod(alignmentPeriod string) *MetricQueryBuilder {
	builder.internal.AlignmentPeriod = &alignmentPeriod

	return builder
}

// Aliases can be set to modify the legend labels. e.g. {{metric.label.xxx}}. See docs for more detail.
func (builder *MetricQueryBuilder) AliasBy(aliasBy string) *MetricQueryBuilder {
	builder.internal.AliasBy = &aliasBy

	return builder
}

func (builder *MetricQueryBuilder) EditorMode(editorMode string) *MetricQueryBuilder {
	builder.internal.EditorMode = editorMode

	return builder
}

func (builder *MetricQueryBuilder) MetricType(metricType string) *MetricQueryBuilder {
	builder.internal.MetricType = metricType

	return builder
}

// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
func (builder *MetricQueryBuilder) CrossSeriesReducer(crossSeriesReducer string) *MetricQueryBuilder {
	builder.internal.CrossSeriesReducer = crossSeriesReducer

	return builder
}

// Array of labels to group data by.
func (builder *MetricQueryBuilder) GroupBys(groupBys []string) *MetricQueryBuilder {
	builder.internal.GroupBys = groupBys

	return builder
}

// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
func (builder *MetricQueryBuilder) Filters(filters []string) *MetricQueryBuilder {
	builder.internal.Filters = filters

	return builder
}

func (builder *MetricQueryBuilder) MetricKind(metricKind MetricKind) *MetricQueryBuilder {
	builder.internal.MetricKind = &metricKind

	return builder
}

func (builder *MetricQueryBuilder) ValueType(valueType string) *MetricQueryBuilder {
	builder.internal.ValueType = &valueType

	return builder
}

func (builder *MetricQueryBuilder) View(view string) *MetricQueryBuilder {
	builder.internal.View = &view

	return builder
}

// MQL query to be executed.
func (builder *MetricQueryBuilder) Query(query string) *MetricQueryBuilder {
	builder.internal.Query = query

	return builder
}

// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
func (builder *MetricQueryBuilder) Preprocessor(preprocessor PreprocessorType) *MetricQueryBuilder {
	builder.internal.Preprocessor = &preprocessor

	return builder
}

// To disable the graphPeriod, it should explictly be set to 'disabled'.
func (builder *MetricQueryBuilder) GraphPeriod(graphPeriod string) *MetricQueryBuilder {
	builder.internal.GraphPeriod = &graphPeriod

	return builder
}

func (builder *MetricQueryBuilder) applyDefaults() {
	builder.GraphPeriod("disabled")
}

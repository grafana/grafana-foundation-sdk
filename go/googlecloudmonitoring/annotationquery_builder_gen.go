// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationQuery] = (*AnnotationQueryBuilder)(nil)

// Annotation sub-query properties.
type AnnotationQueryBuilder struct {
	internal *AnnotationQuery
	errors   map[string]cog.BuildErrors
}

func NewAnnotationQueryBuilder() *AnnotationQueryBuilder {
	resource := &AnnotationQuery{}
	builder := &AnnotationQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AnnotationQueryBuilder) Build() (AnnotationQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationQuery{}, err
	}

	return *builder.internal, nil
}

// GCP project to execute the query against.
func (builder *AnnotationQueryBuilder) ProjectName(projectName string) *AnnotationQueryBuilder {
	builder.internal.ProjectName = projectName

	return builder
}

// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
func (builder *AnnotationQueryBuilder) CrossSeriesReducer(crossSeriesReducer string) *AnnotationQueryBuilder {
	builder.internal.CrossSeriesReducer = crossSeriesReducer

	return builder
}

// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
func (builder *AnnotationQueryBuilder) AlignmentPeriod(alignmentPeriod string) *AnnotationQueryBuilder {
	builder.internal.AlignmentPeriod = &alignmentPeriod

	return builder
}

// Alignment function to be used. Defaults to ALIGN_MEAN.
func (builder *AnnotationQueryBuilder) PerSeriesAligner(perSeriesAligner string) *AnnotationQueryBuilder {
	builder.internal.PerSeriesAligner = &perSeriesAligner

	return builder
}

// Array of labels to group data by.
func (builder *AnnotationQueryBuilder) GroupBys(groupBys []string) *AnnotationQueryBuilder {
	builder.internal.GroupBys = groupBys

	return builder
}

// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
func (builder *AnnotationQueryBuilder) Filters(filters []string) *AnnotationQueryBuilder {
	builder.internal.Filters = filters

	return builder
}

// Data view, defaults to FULL.
func (builder *AnnotationQueryBuilder) View(view string) *AnnotationQueryBuilder {
	builder.internal.View = &view

	return builder
}

// Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
func (builder *AnnotationQueryBuilder) SecondaryCrossSeriesReducer(secondaryCrossSeriesReducer string) *AnnotationQueryBuilder {
	builder.internal.SecondaryCrossSeriesReducer = &secondaryCrossSeriesReducer

	return builder
}

// Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
func (builder *AnnotationQueryBuilder) SecondaryAlignmentPeriod(secondaryAlignmentPeriod string) *AnnotationQueryBuilder {
	builder.internal.SecondaryAlignmentPeriod = &secondaryAlignmentPeriod

	return builder
}

// Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
func (builder *AnnotationQueryBuilder) SecondaryPerSeriesAligner(secondaryPerSeriesAligner string) *AnnotationQueryBuilder {
	builder.internal.SecondaryPerSeriesAligner = &secondaryPerSeriesAligner

	return builder
}

// Only present if a preprocessor is selected. Array of labels to group data by.
func (builder *AnnotationQueryBuilder) SecondaryGroupBys(secondaryGroupBys []string) *AnnotationQueryBuilder {
	builder.internal.SecondaryGroupBys = secondaryGroupBys

	return builder
}

// Annotation title.
func (builder *AnnotationQueryBuilder) Title(title string) *AnnotationQueryBuilder {
	builder.internal.Title = &title

	return builder
}

// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
func (builder *AnnotationQueryBuilder) Preprocessor(preprocessor PreprocessorType) *AnnotationQueryBuilder {
	builder.internal.Preprocessor = &preprocessor

	return builder
}

// Annotation text.
func (builder *AnnotationQueryBuilder) Text(text string) *AnnotationQueryBuilder {
	builder.internal.Text = &text

	return builder
}

func (builder *AnnotationQueryBuilder) applyDefaults() {
}

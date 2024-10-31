// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeSeriesList] = (*TimeSeriesListBuilder)(nil)

// Time Series List sub-query properties.
type TimeSeriesListBuilder struct {
	internal *TimeSeriesList
	errors   map[string]cog.BuildErrors
}

func NewTimeSeriesListBuilder() *TimeSeriesListBuilder {
	resource := &TimeSeriesList{}
	builder := &TimeSeriesListBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TimeSeriesListBuilder) Build() (TimeSeriesList, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeSeriesList{}, err
	}

	return *builder.internal, nil
}

// GCP project to execute the query against.
func (builder *TimeSeriesListBuilder) ProjectName(projectName string) *TimeSeriesListBuilder {
	builder.internal.ProjectName = projectName

	return builder
}

// Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
func (builder *TimeSeriesListBuilder) CrossSeriesReducer(crossSeriesReducer string) *TimeSeriesListBuilder {
	builder.internal.CrossSeriesReducer = crossSeriesReducer

	return builder
}

// Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
func (builder *TimeSeriesListBuilder) AlignmentPeriod(alignmentPeriod string) *TimeSeriesListBuilder {
	builder.internal.AlignmentPeriod = &alignmentPeriod

	return builder
}

// Alignment function to be used. Defaults to ALIGN_MEAN.
func (builder *TimeSeriesListBuilder) PerSeriesAligner(perSeriesAligner string) *TimeSeriesListBuilder {
	builder.internal.PerSeriesAligner = &perSeriesAligner

	return builder
}

// Array of labels to group data by.
func (builder *TimeSeriesListBuilder) GroupBys(groupBys []string) *TimeSeriesListBuilder {
	builder.internal.GroupBys = groupBys

	return builder
}

// Array of filters to query data by. Labels that can be filtered on are defined by the metric.
func (builder *TimeSeriesListBuilder) Filters(filters []string) *TimeSeriesListBuilder {
	builder.internal.Filters = filters

	return builder
}

// Data view, defaults to FULL.
func (builder *TimeSeriesListBuilder) View(view string) *TimeSeriesListBuilder {
	builder.internal.View = &view

	return builder
}

// Only present if a preprocessor is selected. Reducer applied across a set of time-series values. Defaults to REDUCE_NONE.
func (builder *TimeSeriesListBuilder) SecondaryCrossSeriesReducer(secondaryCrossSeriesReducer string) *TimeSeriesListBuilder {
	builder.internal.SecondaryCrossSeriesReducer = &secondaryCrossSeriesReducer

	return builder
}

// Only present if a preprocessor is selected. Alignment period to use when regularizing data. Defaults to cloud-monitoring-auto.
func (builder *TimeSeriesListBuilder) SecondaryAlignmentPeriod(secondaryAlignmentPeriod string) *TimeSeriesListBuilder {
	builder.internal.SecondaryAlignmentPeriod = &secondaryAlignmentPeriod

	return builder
}

// Only present if a preprocessor is selected. Alignment function to be used. Defaults to ALIGN_MEAN.
func (builder *TimeSeriesListBuilder) SecondaryPerSeriesAligner(secondaryPerSeriesAligner string) *TimeSeriesListBuilder {
	builder.internal.SecondaryPerSeriesAligner = &secondaryPerSeriesAligner

	return builder
}

// Only present if a preprocessor is selected. Array of labels to group data by.
func (builder *TimeSeriesListBuilder) SecondaryGroupBys(secondaryGroupBys []string) *TimeSeriesListBuilder {
	builder.internal.SecondaryGroupBys = secondaryGroupBys

	return builder
}

// Preprocessor is not part of the API, but is used to store the preprocessor and not affect the UI for the rest of parameters
func (builder *TimeSeriesListBuilder) Preprocessor(preprocessor PreprocessorType) *TimeSeriesListBuilder {
	builder.internal.Preprocessor = &preprocessor

	return builder
}

// Annotation title.
func (builder *TimeSeriesListBuilder) Title(title string) *TimeSeriesListBuilder {
	builder.internal.Title = &title

	return builder
}

// Annotation text.
func (builder *TimeSeriesListBuilder) Text(text string) *TimeSeriesListBuilder {
	builder.internal.Text = &text

	return builder
}

func (builder *TimeSeriesListBuilder) applyDefaults() {
}

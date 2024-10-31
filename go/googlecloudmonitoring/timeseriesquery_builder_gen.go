// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package googlecloudmonitoring

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeSeriesQuery] = (*TimeSeriesQueryBuilder)(nil)

// Time Series sub-query properties.
type TimeSeriesQueryBuilder struct {
	internal *TimeSeriesQuery
	errors   map[string]cog.BuildErrors
}

func NewTimeSeriesQueryBuilder() *TimeSeriesQueryBuilder {
	resource := &TimeSeriesQuery{}
	builder := &TimeSeriesQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TimeSeriesQueryBuilder) Build() (TimeSeriesQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeSeriesQuery{}, err
	}

	return *builder.internal, nil
}

// GCP project to execute the query against.
func (builder *TimeSeriesQueryBuilder) ProjectName(projectName string) *TimeSeriesQueryBuilder {
	builder.internal.ProjectName = projectName

	return builder
}

// MQL query to be executed.
func (builder *TimeSeriesQueryBuilder) Query(query string) *TimeSeriesQueryBuilder {
	builder.internal.Query = query

	return builder
}

// To disable the graphPeriod, it should explictly be set to 'disabled'.
func (builder *TimeSeriesQueryBuilder) GraphPeriod(graphPeriod string) *TimeSeriesQueryBuilder {
	builder.internal.GraphPeriod = &graphPeriod

	return builder
}

func (builder *TimeSeriesQueryBuilder) applyDefaults() {
	builder.GraphPeriod("disabled")
}

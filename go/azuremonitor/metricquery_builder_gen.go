// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MetricQuery] = (*MetricQueryBuilder)(nil)

type MetricQueryBuilder struct {
	internal *MetricQuery
	errors   cog.BuildErrors
}

func NewMetricQueryBuilder() *MetricQueryBuilder {
	resource := NewMetricQuery()
	builder := &MetricQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MetricQueryBuilder) Build() (MetricQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return MetricQuery{}, err
	}

	if len(builder.errors) > 0 {
		return MetricQuery{}, cog.MakeBuildErrors("azuremonitor.metricQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MetricQueryBuilder) RecordError(path string, err error) *MetricQueryBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Array of resource URIs to be queried.
func (builder *MetricQueryBuilder) Resources(resources []cog.Builder[MonitorResource]) *MetricQueryBuilder {
	resourcesResources := make([]MonitorResource, 0, len(resources))
	for _, r1 := range resources {
		resourcesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		resourcesResources = append(resourcesResources, resourcesDepth1)
	}
	builder.internal.Resources = resourcesResources

	return builder
}

// metricNamespace is used as the resource type (or resource namespace).
// It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
// Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
func (builder *MetricQueryBuilder) MetricNamespace(metricNamespace string) *MetricQueryBuilder {
	builder.internal.MetricNamespace = &metricNamespace

	return builder
}

// Used as the value for the metricNamespace property when it's different from the resource namespace.
func (builder *MetricQueryBuilder) CustomNamespace(customNamespace string) *MetricQueryBuilder {
	builder.internal.CustomNamespace = &customNamespace

	return builder
}

// The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
func (builder *MetricQueryBuilder) MetricName(metricName string) *MetricQueryBuilder {
	builder.internal.MetricName = &metricName

	return builder
}

// The Azure region containing the resource(s).
func (builder *MetricQueryBuilder) Region(region string) *MetricQueryBuilder {
	builder.internal.Region = &region

	return builder
}

// The granularity of data points to be queried. Defaults to auto.
func (builder *MetricQueryBuilder) TimeGrain(timeGrain string) *MetricQueryBuilder {
	builder.internal.TimeGrain = &timeGrain

	return builder
}

// The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
func (builder *MetricQueryBuilder) Aggregation(aggregation string) *MetricQueryBuilder {
	builder.internal.Aggregation = &aggregation

	return builder
}

// Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
func (builder *MetricQueryBuilder) DimensionFilters(dimensionFilters []cog.Builder[MetricDimension]) *MetricQueryBuilder {
	dimensionFiltersResources := make([]MetricDimension, 0, len(dimensionFilters))
	for _, r1 := range dimensionFilters {
		dimensionFiltersDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		dimensionFiltersResources = append(dimensionFiltersResources, dimensionFiltersDepth1)
	}
	builder.internal.DimensionFilters = dimensionFiltersResources

	return builder
}

// Maximum number of records to return. Defaults to 10.
func (builder *MetricQueryBuilder) Top(top string) *MetricQueryBuilder {
	builder.internal.Top = &top

	return builder
}

// Time grains that are supported by the metric.
func (builder *MetricQueryBuilder) AllowedTimeGrainsMs(allowedTimeGrainsMs []int64) *MetricQueryBuilder {
	builder.internal.AllowedTimeGrainsMs = allowedTimeGrainsMs

	return builder
}

// Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
func (builder *MetricQueryBuilder) Alias(alias string) *MetricQueryBuilder {
	builder.internal.Alias = &alias

	return builder
}

// @deprecated
func (builder *MetricQueryBuilder) TimeGrainUnit(timeGrainUnit string) *MetricQueryBuilder {
	builder.internal.TimeGrainUnit = &timeGrainUnit

	return builder
}

// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
func (builder *MetricQueryBuilder) Dimension(dimension string) *MetricQueryBuilder {
	builder.internal.Dimension = &dimension

	return builder
}

// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
func (builder *MetricQueryBuilder) DimensionFilter(dimensionFilter string) *MetricQueryBuilder {
	builder.internal.DimensionFilter = &dimensionFilter

	return builder
}

// @deprecated Use metricNamespace instead
func (builder *MetricQueryBuilder) MetricDefinition(metricDefinition string) *MetricQueryBuilder {
	builder.internal.MetricDefinition = &metricDefinition

	return builder
}

// @deprecated Use resourceGroup, resourceName and metricNamespace instead
func (builder *MetricQueryBuilder) ResourceUri(resourceUri string) *MetricQueryBuilder {
	builder.internal.ResourceUri = &resourceUri

	return builder
}

// @deprecated Use resources instead
func (builder *MetricQueryBuilder) ResourceGroup(resourceGroup string) *MetricQueryBuilder {
	builder.internal.ResourceGroup = &resourceGroup

	return builder
}

// @deprecated Use resources instead
func (builder *MetricQueryBuilder) ResourceName(resourceName string) *MetricQueryBuilder {
	builder.internal.ResourceName = &resourceName

	return builder
}

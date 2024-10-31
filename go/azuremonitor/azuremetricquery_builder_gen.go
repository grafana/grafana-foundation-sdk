// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package azuremonitor

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AzureMetricQuery] = (*AzureMetricQueryBuilder)(nil)

type AzureMetricQueryBuilder struct {
	internal *AzureMetricQuery
	errors   map[string]cog.BuildErrors
}

func NewAzureMetricQueryBuilder() *AzureMetricQueryBuilder {
	resource := &AzureMetricQuery{}
	builder := &AzureMetricQueryBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AzureMetricQueryBuilder) Build() (AzureMetricQuery, error) {
	if err := builder.internal.Validate(); err != nil {
		return AzureMetricQuery{}, err
	}

	return *builder.internal, nil
}

// Array of resource URIs to be queried.
func (builder *AzureMetricQueryBuilder) Resources(resources []cog.Builder[AzureMonitorResource]) *AzureMetricQueryBuilder {
	resourcesResources := make([]AzureMonitorResource, 0, len(resources))
	for _, r1 := range resources {
		resourcesDepth1, err := r1.Build()
		if err != nil {
			builder.errors["resources"] = err.(cog.BuildErrors)
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
func (builder *AzureMetricQueryBuilder) MetricNamespace(metricNamespace string) *AzureMetricQueryBuilder {
	builder.internal.MetricNamespace = &metricNamespace

	return builder
}

// Used as the value for the metricNamespace property when it's different from the resource namespace.
func (builder *AzureMetricQueryBuilder) CustomNamespace(customNamespace string) *AzureMetricQueryBuilder {
	builder.internal.CustomNamespace = &customNamespace

	return builder
}

// The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
func (builder *AzureMetricQueryBuilder) MetricName(metricName string) *AzureMetricQueryBuilder {
	builder.internal.MetricName = &metricName

	return builder
}

// The Azure region containing the resource(s).
func (builder *AzureMetricQueryBuilder) Region(region string) *AzureMetricQueryBuilder {
	builder.internal.Region = &region

	return builder
}

// The granularity of data points to be queried. Defaults to auto.
func (builder *AzureMetricQueryBuilder) TimeGrain(timeGrain string) *AzureMetricQueryBuilder {
	builder.internal.TimeGrain = &timeGrain

	return builder
}

// The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
func (builder *AzureMetricQueryBuilder) Aggregation(aggregation string) *AzureMetricQueryBuilder {
	builder.internal.Aggregation = &aggregation

	return builder
}

// Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
func (builder *AzureMetricQueryBuilder) DimensionFilters(dimensionFilters []cog.Builder[AzureMetricDimension]) *AzureMetricQueryBuilder {
	dimensionFiltersResources := make([]AzureMetricDimension, 0, len(dimensionFilters))
	for _, r1 := range dimensionFilters {
		dimensionFiltersDepth1, err := r1.Build()
		if err != nil {
			builder.errors["dimensionFilters"] = err.(cog.BuildErrors)
			return builder
		}
		dimensionFiltersResources = append(dimensionFiltersResources, dimensionFiltersDepth1)
	}
	builder.internal.DimensionFilters = dimensionFiltersResources

	return builder
}

// Maximum number of records to return. Defaults to 10.
func (builder *AzureMetricQueryBuilder) Top(top string) *AzureMetricQueryBuilder {
	builder.internal.Top = &top

	return builder
}

// Time grains that are supported by the metric.
func (builder *AzureMetricQueryBuilder) AllowedTimeGrainsMs(allowedTimeGrainsMs []int64) *AzureMetricQueryBuilder {
	builder.internal.AllowedTimeGrainsMs = allowedTimeGrainsMs

	return builder
}

// Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
func (builder *AzureMetricQueryBuilder) Alias(alias string) *AzureMetricQueryBuilder {
	builder.internal.Alias = &alias

	return builder
}

// @deprecated
func (builder *AzureMetricQueryBuilder) TimeGrainUnit(timeGrainUnit string) *AzureMetricQueryBuilder {
	builder.internal.TimeGrainUnit = &timeGrainUnit

	return builder
}

// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
func (builder *AzureMetricQueryBuilder) Dimension(dimension string) *AzureMetricQueryBuilder {
	builder.internal.Dimension = &dimension

	return builder
}

// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
func (builder *AzureMetricQueryBuilder) DimensionFilter(dimensionFilter string) *AzureMetricQueryBuilder {
	builder.internal.DimensionFilter = &dimensionFilter

	return builder
}

// @deprecated Use metricNamespace instead
func (builder *AzureMetricQueryBuilder) MetricDefinition(metricDefinition string) *AzureMetricQueryBuilder {
	builder.internal.MetricDefinition = &metricDefinition

	return builder
}

// @deprecated Use resourceGroup, resourceName and metricNamespace instead
func (builder *AzureMetricQueryBuilder) ResourceUri(resourceUri string) *AzureMetricQueryBuilder {
	builder.internal.ResourceUri = &resourceUri

	return builder
}

// @deprecated Use resources instead
func (builder *AzureMetricQueryBuilder) ResourceGroup(resourceGroup string) *AzureMetricQueryBuilder {
	builder.internal.ResourceGroup = &resourceGroup

	return builder
}

// @deprecated Use resources instead
func (builder *AzureMetricQueryBuilder) ResourceName(resourceName string) *AzureMetricQueryBuilder {
	builder.internal.ResourceName = &resourceName

	return builder
}

func (builder *AzureMetricQueryBuilder) applyDefaults() {
}

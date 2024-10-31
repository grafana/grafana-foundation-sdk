---
title: <span class="badge object-type-interface"></span> AzureMetricQuery
---
# <span class="badge object-type-interface"></span> AzureMetricQuery

## Definition

```typescript
export interface AzureMetricQuery {
	// Array of resource URIs to be queried.
	resources?: azuremonitor.AzureMonitorResource[];
	// metricNamespace is used as the resource type (or resource namespace).
	// It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
	// Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
	metricNamespace?: string;
	// Used as the value for the metricNamespace property when it's different from the resource namespace.
	customNamespace?: string;
	// The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
	metricName?: string;
	// The Azure region containing the resource(s).
	region?: string;
	// The granularity of data points to be queried. Defaults to auto.
	timeGrain?: string;
	// The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
	aggregation?: string;
	// Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
	dimensionFilters?: azuremonitor.AzureMetricDimension[];
	// Maximum number of records to return. Defaults to 10.
	top?: string;
	// Time grains that are supported by the metric.
	allowedTimeGrainsMs?: number[];
	// Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
	alias?: string;
	// @deprecated
	timeGrainUnit?: string;
	// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
	dimension?: string;
	// @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
	dimensionFilter?: string;
	// @deprecated Use metricNamespace instead
	metricDefinition?: string;
	// @deprecated Use resourceGroup, resourceName and metricNamespace instead
	resourceUri?: string;
	// @deprecated Use resources instead
	resourceGroup?: string;
	// @deprecated Use resources instead
	resourceName?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AzureMetricQueryBuilder](./builder-AzureMetricQueryBuilder.md)

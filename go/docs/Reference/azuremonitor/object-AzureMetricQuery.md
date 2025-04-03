---
title: <span class="badge object-type-struct"></span> AzureMetricQuery
---
# <span class="badge object-type-struct"></span> AzureMetricQuery

## Definition

```go
type AzureMetricQuery struct {
    // Array of resource URIs to be queried.
    Resources []azuremonitor.AzureMonitorResource `json:"resources,omitempty"`
    // metricNamespace is used as the resource type (or resource namespace).
    // It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
    // Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
    MetricNamespace *string `json:"metricNamespace,omitempty"`
    // Used as the value for the metricNamespace property when it's different from the resource namespace.
    CustomNamespace *string `json:"customNamespace,omitempty"`
    // The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
    MetricName *string `json:"metricName,omitempty"`
    // The Azure region containing the resource(s).
    Region *string `json:"region,omitempty"`
    // The granularity of data points to be queried. Defaults to auto.
    TimeGrain *string `json:"timeGrain,omitempty"`
    // The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
    Aggregation *string `json:"aggregation,omitempty"`
    // Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
    DimensionFilters []azuremonitor.AzureMetricDimension `json:"dimensionFilters,omitempty"`
    // Maximum number of records to return. Defaults to 10.
    Top *string `json:"top,omitempty"`
    // Time grains that are supported by the metric.
    AllowedTimeGrainsMs []int64 `json:"allowedTimeGrainsMs,omitempty"`
    // Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
    Alias *string `json:"alias,omitempty"`
    // @deprecated
    TimeGrainUnit *string `json:"timeGrainUnit,omitempty"`
    // @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    Dimension *string `json:"dimension,omitempty"`
    // @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    DimensionFilter *string `json:"dimensionFilter,omitempty"`
    // @deprecated Use metricNamespace instead
    MetricDefinition *string `json:"metricDefinition,omitempty"`
    // @deprecated Use resourceGroup, resourceName and metricNamespace instead
    ResourceUri *string `json:"resourceUri,omitempty"`
    // @deprecated Use resources instead
    ResourceGroup *string `json:"resourceGroup,omitempty"`
    // @deprecated Use resources instead
    ResourceName *string `json:"resourceName,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AzureMetricQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (azureMetricQuery *AzureMetricQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AzureMetricQuery` objects.

```go
func (azureMetricQuery *AzureMetricQuery) Equals(other AzureMetricQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AzureMetricQuery` fields for violations and returns them.

```go
func (azureMetricQuery *AzureMetricQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [AzureMetricQueryBuilder](./builder-AzureMetricQueryBuilder.md)

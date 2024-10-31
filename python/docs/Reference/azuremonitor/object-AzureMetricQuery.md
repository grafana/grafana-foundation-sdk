---
title: <span class="badge object-type-class"></span> AzureMetricQuery
---
# <span class="badge object-type-class"></span> AzureMetricQuery

## Definition

```python
class AzureMetricQuery:
    # Array of resource URIs to be queried.
    resources: typing.Optional[list[azuremonitor.AzureMonitorResource]]
    # metricNamespace is used as the resource type (or resource namespace).
    # It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
    # Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
    metric_namespace: typing.Optional[str]
    # Used as the value for the metricNamespace property when it's different from the resource namespace.
    custom_namespace: typing.Optional[str]
    # The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
    metric_name: typing.Optional[str]
    # The Azure region containing the resource(s).
    region: typing.Optional[str]
    # The granularity of data points to be queried. Defaults to auto.
    time_grain: typing.Optional[str]
    # The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
    aggregation: typing.Optional[str]
    # Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
    dimension_filters: typing.Optional[list[azuremonitor.AzureMetricDimension]]
    # Maximum number of records to return. Defaults to 10.
    top: typing.Optional[str]
    # Time grains that are supported by the metric.
    allowed_time_grains_ms: typing.Optional[list[int]]
    # Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
    alias: typing.Optional[str]
    # @deprecated
    time_grain_unit: typing.Optional[str]
    # @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    dimension: typing.Optional[str]
    # @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
    dimension_filter: typing.Optional[str]
    # @deprecated Use metricNamespace instead
    metric_definition: typing.Optional[str]
    # @deprecated Use resourceGroup, resourceName and metricNamespace instead
    resource_uri: typing.Optional[str]
    # @deprecated Use resources instead
    resource_group: typing.Optional[str]
    # @deprecated Use resources instead
    resource_name: typing.Optional[str]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [AzureMetricQuery](./builder-AzureMetricQuery.md)

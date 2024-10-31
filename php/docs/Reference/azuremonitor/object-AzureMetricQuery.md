---
title: <span class="badge object-type-class"></span> AzureMetricQuery
---
# <span class="badge object-type-class"></span> AzureMetricQuery

## Definition

```php
class AzureMetricQuery implements \JsonSerializable
{
    /**
     * Array of resource URIs to be queried.
     * @var array<\Grafana\Foundation\Azuremonitor\AzureMonitorResource>|null
     */
    public ?array $resources;

    /**
     * metricNamespace is used as the resource type (or resource namespace).
     * It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts
     * Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.
     */
    public ?string $metricNamespace;

    /**
     * Used as the value for the metricNamespace property when it's different from the resource namespace.
     */
    public ?string $customNamespace;

    /**
     * The metric to query data for within the specified metricNamespace. e.g. UsedCapacity
     */
    public ?string $metricName;

    /**
     * The Azure region containing the resource(s).
     */
    public ?string $region;

    /**
     * The granularity of data points to be queried. Defaults to auto.
     */
    public ?string $timeGrain;

    /**
     * The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.
     */
    public ?string $aggregation;

    /**
     * Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.
     * @var array<\Grafana\Foundation\Azuremonitor\AzureMetricDimension>|null
     */
    public ?array $dimensionFilters;

    /**
     * Maximum number of records to return. Defaults to 10.
     */
    public ?string $top;

    /**
     * Time grains that are supported by the metric.
     * @var array<int>|null
     */
    public ?array $allowedTimeGrainsMs;

    /**
     * Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.
     */
    public ?string $alias;

    /**
     * @deprecated
     */
    public ?string $timeGrainUnit;

    /**
     * @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
     */
    public ?string $dimension;

    /**
     * @deprecated This property was migrated to dimensionFilters and should only be accessed in the migration
     */
    public ?string $dimensionFilter;

    /**
     * @deprecated Use metricNamespace instead
     */
    public ?string $metricDefinition;

    /**
     * @deprecated Use resourceGroup, resourceName and metricNamespace instead
     */
    public ?string $resourceUri;

    /**
     * @deprecated Use resources instead
     */
    public ?string $resourceGroup;

    /**
     * @deprecated Use resources instead
     */
    public ?string $resourceName;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [AzureMetricQueryBuilder](./builder-AzureMetricQueryBuilder.md)

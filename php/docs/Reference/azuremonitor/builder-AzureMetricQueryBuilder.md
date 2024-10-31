---
title: <span class="badge builder"></span> AzureMetricQueryBuilder
---
# <span class="badge builder"></span> AzureMetricQueryBuilder

## Constructor

```php
new AzureMetricQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> aggregation

The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.

```php
aggregation(string $aggregation)
```

### <span class="badge object-method"></span> alias

Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.

```php
alias(string $alias)
```

### <span class="badge object-method"></span> allowedTimeGrainsMs

Time grains that are supported by the metric.

@param array<int> $allowedTimeGrainsMs

```php
allowedTimeGrainsMs(array $allowedTimeGrainsMs)
```

### <span class="badge object-method"></span> customNamespace

Used as the value for the metricNamespace property when it's different from the resource namespace.

```php
customNamespace(string $customNamespace)
```

### <span class="badge object-method"></span> dimension

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```php
dimension(string $dimension)
```

### <span class="badge object-method"></span> dimensionFilter

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```php
dimensionFilter(string $dimensionFilter)
```

### <span class="badge object-method"></span> dimensionFilters

Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMetricDimension>> $dimensionFilters

```php
dimensionFilters(array $dimensionFilters)
```

### <span class="badge object-method"></span> metricDefinition

@deprecated Use metricNamespace instead

```php
metricDefinition(string $metricDefinition)
```

### <span class="badge object-method"></span> metricName

The metric to query data for within the specified metricNamespace. e.g. UsedCapacity

```php
metricName(string $metricName)
```

### <span class="badge object-method"></span> metricNamespace

metricNamespace is used as the resource type (or resource namespace).

It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts

Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.

```php
metricNamespace(string $metricNamespace)
```

### <span class="badge object-method"></span> region

The Azure region containing the resource(s).

```php
region(string $region)
```

### <span class="badge object-method"></span> resourceGroup

@deprecated Use resources instead

```php
resourceGroup(string $resourceGroup)
```

### <span class="badge object-method"></span> resourceName

@deprecated Use resources instead

```php
resourceName(string $resourceName)
```

### <span class="badge object-method"></span> resourceUri

@deprecated Use resourceGroup, resourceName and metricNamespace instead

```php
resourceUri(string $resourceUri)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Azuremonitor\AzureMonitorResource>> $resources

```php
resources(array $resources)
```

### <span class="badge object-method"></span> timeGrain

The granularity of data points to be queried. Defaults to auto.

```php
timeGrain(string $timeGrain)
```

### <span class="badge object-method"></span> timeGrainUnit

@deprecated

```php
timeGrainUnit(string $timeGrainUnit)
```

### <span class="badge object-method"></span> top

Maximum number of records to return. Defaults to 10.

```php
top(string $top)
```

## See also

 * <span class="badge object-type-class"></span> [AzureMetricQuery](./object-AzureMetricQuery.md)

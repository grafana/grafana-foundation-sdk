---
title: <span class="badge builder"></span> AzureMetricQueryBuilder
---
# <span class="badge builder"></span> AzureMetricQueryBuilder

## Constructor

```typescript
new AzureMetricQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> aggregation

The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.

```typescript
aggregation(aggregation: string)
```

### <span class="badge object-method"></span> alias

Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.

```typescript
alias(alias: string)
```

### <span class="badge object-method"></span> allowedTimeGrainsMs

Time grains that are supported by the metric.

```typescript
allowedTimeGrainsMs(allowedTimeGrainsMs: number[])
```

### <span class="badge object-method"></span> customNamespace

Used as the value for the metricNamespace property when it's different from the resource namespace.

```typescript
customNamespace(customNamespace: string)
```

### <span class="badge object-method"></span> dimension

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```typescript
dimension(dimension: string)
```

### <span class="badge object-method"></span> dimensionFilter

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```typescript
dimensionFilter(dimensionFilter: string)
```

### <span class="badge object-method"></span> dimensionFilters

Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.

```typescript
dimensionFilters(dimensionFilters: cog.Builder<azuremonitor.AzureMetricDimension>[])
```

### <span class="badge object-method"></span> metricDefinition

@deprecated Use metricNamespace instead

```typescript
metricDefinition(metricDefinition: string)
```

### <span class="badge object-method"></span> metricName

The metric to query data for within the specified metricNamespace. e.g. UsedCapacity

```typescript
metricName(metricName: string)
```

### <span class="badge object-method"></span> metricNamespace

metricNamespace is used as the resource type (or resource namespace).

It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts

Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.

```typescript
metricNamespace(metricNamespace: string)
```

### <span class="badge object-method"></span> region

The Azure region containing the resource(s).

```typescript
region(region: string)
```

### <span class="badge object-method"></span> resourceGroup

@deprecated Use resources instead

```typescript
resourceGroup(resourceGroup: string)
```

### <span class="badge object-method"></span> resourceName

@deprecated Use resources instead

```typescript
resourceName(resourceName: string)
```

### <span class="badge object-method"></span> resourceUri

@deprecated Use resourceGroup, resourceName and metricNamespace instead

```typescript
resourceUri(resourceUri: string)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```typescript
resources(resources: cog.Builder<azuremonitor.AzureMonitorResource>[])
```

### <span class="badge object-method"></span> timeGrain

The granularity of data points to be queried. Defaults to auto.

```typescript
timeGrain(timeGrain: string)
```

### <span class="badge object-method"></span> timeGrainUnit

@deprecated

```typescript
timeGrainUnit(timeGrainUnit: string)
```

### <span class="badge object-method"></span> top

Maximum number of records to return. Defaults to 10.

```typescript
top(top: string)
```

## See also

 * <span class="badge object-type-interface"></span> [AzureMetricQuery](./object-AzureMetricQuery.md)

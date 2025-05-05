---
title: <span class="badge builder"></span> AzureMetricQueryBuilder
---
# <span class="badge builder"></span> AzureMetricQueryBuilder

## Constructor

```java
new AzureMetricQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AzureMetricQuery build()
```

### <span class="badge object-method"></span> aggregation

The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.

```java
public AzureMetricQueryBuilder aggregation(String aggregation)
```

### <span class="badge object-method"></span> alias

Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.

```java
public AzureMetricQueryBuilder alias(String alias)
```

### <span class="badge object-method"></span> allowedTimeGrainsMs

Time grains that are supported by the metric.

```java
public AzureMetricQueryBuilder allowedTimeGrainsMs(List<Long> allowedTimeGrainsMs)
```

### <span class="badge object-method"></span> customNamespace

Used as the value for the metricNamespace property when it's different from the resource namespace.

```java
public AzureMetricQueryBuilder customNamespace(String customNamespace)
```

### <span class="badge object-method"></span> dimension

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```java
public AzureMetricQueryBuilder dimension(String dimension)
```

### <span class="badge object-method"></span> dimensionFilter

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```java
public AzureMetricQueryBuilder dimensionFilter(String dimensionFilter)
```

### <span class="badge object-method"></span> dimensionFilters

Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.

```java
public AzureMetricQueryBuilder dimensionFilters(List<com.grafana.foundation.cog.Builder<AzureMetricDimension>> dimensionFilters)
```

### <span class="badge object-method"></span> metricDefinition

@deprecated Use metricNamespace instead

```java
public AzureMetricQueryBuilder metricDefinition(String metricDefinition)
```

### <span class="badge object-method"></span> metricName

The metric to query data for within the specified metricNamespace. e.g. UsedCapacity

```java
public AzureMetricQueryBuilder metricName(String metricName)
```

### <span class="badge object-method"></span> metricNamespace

metricNamespace is used as the resource type (or resource namespace).

It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts

Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.

```java
public AzureMetricQueryBuilder metricNamespace(String metricNamespace)
```

### <span class="badge object-method"></span> region

The Azure region containing the resource(s).

```java
public AzureMetricQueryBuilder region(String region)
```

### <span class="badge object-method"></span> resourceGroup

@deprecated Use resources instead

```java
public AzureMetricQueryBuilder resourceGroup(String resourceGroup)
```

### <span class="badge object-method"></span> resourceName

@deprecated Use resources instead

```java
public AzureMetricQueryBuilder resourceName(String resourceName)
```

### <span class="badge object-method"></span> resourceUri

@deprecated Use resourceGroup, resourceName and metricNamespace instead

```java
public AzureMetricQueryBuilder resourceUri(String resourceUri)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```java
public AzureMetricQueryBuilder resources(List<com.grafana.foundation.cog.Builder<AzureMonitorResource>> resources)
```

### <span class="badge object-method"></span> timeGrain

The granularity of data points to be queried. Defaults to auto.

```java
public AzureMetricQueryBuilder timeGrain(String timeGrain)
```

### <span class="badge object-method"></span> timeGrainUnit

@deprecated

```java
public AzureMetricQueryBuilder timeGrainUnit(String timeGrainUnit)
```

### <span class="badge object-method"></span> top

Maximum number of records to return. Defaults to 10.

```java
public AzureMetricQueryBuilder top(String top)
```

## See also

 * <span class="badge object-type-class"></span> [AzureMetricQuery](./object-AzureMetricQuery.md)

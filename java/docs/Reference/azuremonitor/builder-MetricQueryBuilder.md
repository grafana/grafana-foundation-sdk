---
title: <span class="badge builder"></span> MetricQueryBuilder
---
# <span class="badge builder"></span> MetricQueryBuilder

## Constructor

```java
new MetricQueryBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public MetricQuery build()
```

### <span class="badge object-method"></span> aggregation

The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.

```java
public MetricQueryBuilder aggregation(String aggregation)
```

### <span class="badge object-method"></span> alias

Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.

```java
public MetricQueryBuilder alias(String alias)
```

### <span class="badge object-method"></span> allowedTimeGrainsMs

Time grains that are supported by the metric.

```java
public MetricQueryBuilder allowedTimeGrainsMs(List<Long> allowedTimeGrainsMs)
```

### <span class="badge object-method"></span> customNamespace

Used as the value for the metricNamespace property when it's different from the resource namespace.

```java
public MetricQueryBuilder customNamespace(String customNamespace)
```

### <span class="badge object-method"></span> dimension

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```java
public MetricQueryBuilder dimension(String dimension)
```

### <span class="badge object-method"></span> dimensionFilter

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```java
public MetricQueryBuilder dimensionFilter(String dimensionFilter)
```

### <span class="badge object-method"></span> dimensionFilters

Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.

```java
public MetricQueryBuilder dimensionFilters(List<com.grafana.foundation.cog.Builder<MetricDimension>> dimensionFilters)
```

### <span class="badge object-method"></span> metricDefinition

@deprecated Use metricNamespace instead

```java
public MetricQueryBuilder metricDefinition(String metricDefinition)
```

### <span class="badge object-method"></span> metricName

The metric to query data for within the specified metricNamespace. e.g. UsedCapacity

```java
public MetricQueryBuilder metricName(String metricName)
```

### <span class="badge object-method"></span> metricNamespace

metricNamespace is used as the resource type (or resource namespace).

It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts

Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.

```java
public MetricQueryBuilder metricNamespace(String metricNamespace)
```

### <span class="badge object-method"></span> region

The Azure region containing the resource(s).

```java
public MetricQueryBuilder region(String region)
```

### <span class="badge object-method"></span> resourceGroup

@deprecated Use resources instead

```java
public MetricQueryBuilder resourceGroup(String resourceGroup)
```

### <span class="badge object-method"></span> resourceName

@deprecated Use resources instead

```java
public MetricQueryBuilder resourceName(String resourceName)
```

### <span class="badge object-method"></span> resourceUri

@deprecated Use resourceGroup, resourceName and metricNamespace instead

```java
public MetricQueryBuilder resourceUri(String resourceUri)
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```java
public MetricQueryBuilder resources(List<com.grafana.foundation.cog.Builder<MonitorResource>> resources)
```

### <span class="badge object-method"></span> timeGrain

The granularity of data points to be queried. Defaults to auto.

```java
public MetricQueryBuilder timeGrain(String timeGrain)
```

### <span class="badge object-method"></span> timeGrainUnit

@deprecated

```java
public MetricQueryBuilder timeGrainUnit(String timeGrainUnit)
```

### <span class="badge object-method"></span> top

Maximum number of records to return. Defaults to 10.

```java
public MetricQueryBuilder top(String top)
```

## See also

 * <span class="badge object-type-class"></span> [MetricQuery](./object-MetricQuery.md)

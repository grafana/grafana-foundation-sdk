---
title: <span class="badge builder"></span> AzureMetricQuery
---
# <span class="badge builder"></span> AzureMetricQuery

## Constructor

```python
AzureMetricQuery()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> azuremonitor.AzureMetricQuery
```

### <span class="badge object-method"></span> aggregation

The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.

```python
def aggregation(aggregation: str) -> typing.Self
```

### <span class="badge object-method"></span> alias

Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.

```python
def alias(alias: str) -> typing.Self
```

### <span class="badge object-method"></span> allowed_time_grains_ms

Time grains that are supported by the metric.

```python
def allowed_time_grains_ms(allowed_time_grains_ms: list[int]) -> typing.Self
```

### <span class="badge object-method"></span> custom_namespace

Used as the value for the metricNamespace property when it's different from the resource namespace.

```python
def custom_namespace(custom_namespace: str) -> typing.Self
```

### <span class="badge object-method"></span> dimension

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```python
def dimension(dimension: str) -> typing.Self
```

### <span class="badge object-method"></span> dimension_filter

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```python
def dimension_filter(dimension_filter: str) -> typing.Self
```

### <span class="badge object-method"></span> dimension_filters

Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.

```python
def dimension_filters(dimension_filters: list[cogbuilder.Builder[azuremonitor.AzureMetricDimension]]) -> typing.Self
```

### <span class="badge object-method"></span> metric_definition

@deprecated Use metricNamespace instead

```python
def metric_definition(metric_definition: str) -> typing.Self
```

### <span class="badge object-method"></span> metric_name

The metric to query data for within the specified metricNamespace. e.g. UsedCapacity

```python
def metric_name(metric_name: str) -> typing.Self
```

### <span class="badge object-method"></span> metric_namespace

metricNamespace is used as the resource type (or resource namespace).

It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts

Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.

```python
def metric_namespace(metric_namespace: str) -> typing.Self
```

### <span class="badge object-method"></span> region

The Azure region containing the resource(s).

```python
def region(region: str) -> typing.Self
```

### <span class="badge object-method"></span> resource_group

@deprecated Use resources instead

```python
def resource_group(resource_group: str) -> typing.Self
```

### <span class="badge object-method"></span> resource_name

@deprecated Use resources instead

```python
def resource_name(resource_name: str) -> typing.Self
```

### <span class="badge object-method"></span> resource_uri

@deprecated Use resourceGroup, resourceName and metricNamespace instead

```python
def resource_uri(resource_uri: str) -> typing.Self
```

### <span class="badge object-method"></span> resources

Array of resource URIs to be queried.

```python
def resources(resources: list[cogbuilder.Builder[azuremonitor.AzureMonitorResource]]) -> typing.Self
```

### <span class="badge object-method"></span> time_grain

The granularity of data points to be queried. Defaults to auto.

```python
def time_grain(time_grain: str) -> typing.Self
```

### <span class="badge object-method"></span> time_grain_unit

@deprecated

```python
def time_grain_unit(time_grain_unit: str) -> typing.Self
```

### <span class="badge object-method"></span> top

Maximum number of records to return. Defaults to 10.

```python
def top(top: str) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [AzureMetricQuery](./object-AzureMetricQuery.md)

---
title: <span class="badge builder"></span> MetricQueryBuilder
---
# <span class="badge builder"></span> MetricQueryBuilder

## Constructor

```go
func NewMetricQueryBuilder() *MetricQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricQueryBuilder) Build() (MetricQuery, error)
```

### <span class="badge object-method"></span> Aggregation

The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.

```go
func (builder *MetricQueryBuilder) Aggregation(aggregation string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> Alias

Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.

```go
func (builder *MetricQueryBuilder) Alias(alias string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> AllowedTimeGrainsMs

Time grains that are supported by the metric.

```go
func (builder *MetricQueryBuilder) AllowedTimeGrainsMs(allowedTimeGrainsMs []int64) *MetricQueryBuilder
```

### <span class="badge object-method"></span> CustomNamespace

Used as the value for the metricNamespace property when it's different from the resource namespace.

```go
func (builder *MetricQueryBuilder) CustomNamespace(customNamespace string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> Dimension

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```go
func (builder *MetricQueryBuilder) Dimension(dimension string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> DimensionFilter

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```go
func (builder *MetricQueryBuilder) DimensionFilter(dimensionFilter string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> DimensionFilters

Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.

```go
func (builder *MetricQueryBuilder) DimensionFilters(dimensionFilters []cog.Builder[azuremonitor.MetricDimension]) *MetricQueryBuilder
```

### <span class="badge object-method"></span> MetricDefinition

@deprecated Use metricNamespace instead

```go
func (builder *MetricQueryBuilder) MetricDefinition(metricDefinition string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> MetricName

The metric to query data for within the specified metricNamespace. e.g. UsedCapacity

```go
func (builder *MetricQueryBuilder) MetricName(metricName string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> MetricNamespace

metricNamespace is used as the resource type (or resource namespace).

It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts

Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.

```go
func (builder *MetricQueryBuilder) MetricNamespace(metricNamespace string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> Region

The Azure region containing the resource(s).

```go
func (builder *MetricQueryBuilder) Region(region string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

@deprecated Use resources instead

```go
func (builder *MetricQueryBuilder) ResourceGroup(resourceGroup string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> ResourceName

@deprecated Use resources instead

```go
func (builder *MetricQueryBuilder) ResourceName(resourceName string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> ResourceUri

@deprecated Use resourceGroup, resourceName and metricNamespace instead

```go
func (builder *MetricQueryBuilder) ResourceUri(resourceUri string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> Resources

Array of resource URIs to be queried.

```go
func (builder *MetricQueryBuilder) Resources(resources []cog.Builder[azuremonitor.MonitorResource]) *MetricQueryBuilder
```

### <span class="badge object-method"></span> TimeGrain

The granularity of data points to be queried. Defaults to auto.

```go
func (builder *MetricQueryBuilder) TimeGrain(timeGrain string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> TimeGrainUnit

@deprecated

```go
func (builder *MetricQueryBuilder) TimeGrainUnit(timeGrainUnit string) *MetricQueryBuilder
```

### <span class="badge object-method"></span> Top

Maximum number of records to return. Defaults to 10.

```go
func (builder *MetricQueryBuilder) Top(top string) *MetricQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricQuery](./object-MetricQuery.md)

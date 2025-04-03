---
title: <span class="badge builder"></span> AzureMetricQueryBuilder
---
# <span class="badge builder"></span> AzureMetricQueryBuilder

## Constructor

```go
func NewAzureMetricQueryBuilder() *AzureMetricQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AzureMetricQueryBuilder) Build() (AzureMetricQuery, error)
```

### <span class="badge object-method"></span> Aggregation

The aggregation to be used within the query. Defaults to the primaryAggregationType defined by the metric.

```go
func (builder *AzureMetricQueryBuilder) Aggregation(aggregation string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> Alias

Aliases can be set to modify the legend labels. e.g. {{ resourceGroup }}. See docs for more detail.

```go
func (builder *AzureMetricQueryBuilder) Alias(alias string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> AllowedTimeGrainsMs

Time grains that are supported by the metric.

```go
func (builder *AzureMetricQueryBuilder) AllowedTimeGrainsMs(allowedTimeGrainsMs []int64) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> CustomNamespace

Used as the value for the metricNamespace property when it's different from the resource namespace.

```go
func (builder *AzureMetricQueryBuilder) CustomNamespace(customNamespace string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> Dimension

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```go
func (builder *AzureMetricQueryBuilder) Dimension(dimension string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> DimensionFilter

@deprecated This property was migrated to dimensionFilters and should only be accessed in the migration

```go
func (builder *AzureMetricQueryBuilder) DimensionFilter(dimensionFilter string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> DimensionFilters

Filters to reduce the set of data returned. Dimensions that can be filtered on are defined by the metric.

```go
func (builder *AzureMetricQueryBuilder) DimensionFilters(dimensionFilters []cog.Builder[azuremonitor.AzureMetricDimension]) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> MetricDefinition

@deprecated Use metricNamespace instead

```go
func (builder *AzureMetricQueryBuilder) MetricDefinition(metricDefinition string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> MetricName

The metric to query data for within the specified metricNamespace. e.g. UsedCapacity

```go
func (builder *AzureMetricQueryBuilder) MetricName(metricName string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> MetricNamespace

metricNamespace is used as the resource type (or resource namespace).

It's usually equal to the target metric namespace. e.g. microsoft.storage/storageaccounts

Kept the name of the variable as metricNamespace to avoid backward incompatibility issues.

```go
func (builder *AzureMetricQueryBuilder) MetricNamespace(metricNamespace string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> Region

The Azure region containing the resource(s).

```go
func (builder *AzureMetricQueryBuilder) Region(region string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

@deprecated Use resources instead

```go
func (builder *AzureMetricQueryBuilder) ResourceGroup(resourceGroup string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> ResourceName

@deprecated Use resources instead

```go
func (builder *AzureMetricQueryBuilder) ResourceName(resourceName string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> ResourceUri

@deprecated Use resourceGroup, resourceName and metricNamespace instead

```go
func (builder *AzureMetricQueryBuilder) ResourceUri(resourceUri string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> Resources

Array of resource URIs to be queried.

```go
func (builder *AzureMetricQueryBuilder) Resources(resources []cog.Builder[azuremonitor.AzureMonitorResource]) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> TimeGrain

The granularity of data points to be queried. Defaults to auto.

```go
func (builder *AzureMetricQueryBuilder) TimeGrain(timeGrain string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> TimeGrainUnit

@deprecated

```go
func (builder *AzureMetricQueryBuilder) TimeGrainUnit(timeGrainUnit string) *AzureMetricQueryBuilder
```

### <span class="badge object-method"></span> Top

Maximum number of records to return. Defaults to 10.

```go
func (builder *AzureMetricQueryBuilder) Top(top string) *AzureMetricQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AzureMetricQuery](./object-AzureMetricQuery.md)

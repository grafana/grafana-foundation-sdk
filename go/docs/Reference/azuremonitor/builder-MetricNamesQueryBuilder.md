---
title: <span class="badge builder"></span> MetricNamesQueryBuilder
---
# <span class="badge builder"></span> MetricNamesQueryBuilder

## Constructor

```go
func NewMetricNamesQueryBuilder() *MetricNamesQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricNamesQueryBuilder) Build() (MetricNamesQuery, error)
```

### <span class="badge object-method"></span> MetricNamespace

```go
func (builder *MetricNamesQueryBuilder) MetricNamespace(metricNamespace string) *MetricNamesQueryBuilder
```

### <span class="badge object-method"></span> RawQuery

```go
func (builder *MetricNamesQueryBuilder) RawQuery(rawQuery string) *MetricNamesQueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

```go
func (builder *MetricNamesQueryBuilder) ResourceGroup(resourceGroup string) *MetricNamesQueryBuilder
```

### <span class="badge object-method"></span> ResourceName

```go
func (builder *MetricNamesQueryBuilder) ResourceName(resourceName string) *MetricNamesQueryBuilder
```

### <span class="badge object-method"></span> Subscription

```go
func (builder *MetricNamesQueryBuilder) Subscription(subscription string) *MetricNamesQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricNamesQuery](./object-MetricNamesQuery.md)

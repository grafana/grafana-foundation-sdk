---
title: <span class="badge builder"></span> MetricNamespaceQueryBuilder
---
# <span class="badge builder"></span> MetricNamespaceQueryBuilder

## Constructor

```go
func NewMetricNamespaceQueryBuilder() *MetricNamespaceQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricNamespaceQueryBuilder) Build() (MetricNamespaceQuery, error)
```

### <span class="badge object-method"></span> MetricNamespace

```go
func (builder *MetricNamespaceQueryBuilder) MetricNamespace(metricNamespace string) *MetricNamespaceQueryBuilder
```

### <span class="badge object-method"></span> RawQuery

```go
func (builder *MetricNamespaceQueryBuilder) RawQuery(rawQuery string) *MetricNamespaceQueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

```go
func (builder *MetricNamespaceQueryBuilder) ResourceGroup(resourceGroup string) *MetricNamespaceQueryBuilder
```

### <span class="badge object-method"></span> ResourceName

```go
func (builder *MetricNamespaceQueryBuilder) ResourceName(resourceName string) *MetricNamespaceQueryBuilder
```

### <span class="badge object-method"></span> Subscription

```go
func (builder *MetricNamespaceQueryBuilder) Subscription(subscription string) *MetricNamespaceQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricNamespaceQuery](./object-MetricNamespaceQuery.md)

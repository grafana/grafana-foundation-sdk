---
title: <span class="badge builder"></span> MetricDefinitionsQueryBuilder
---
# <span class="badge builder"></span> MetricDefinitionsQueryBuilder

## Constructor

```go
func NewMetricDefinitionsQueryBuilder() *MetricDefinitionsQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricDefinitionsQueryBuilder) Build() (MetricDefinitionsQuery, error)
```

### <span class="badge object-method"></span> MetricNamespace

```go
func (builder *MetricDefinitionsQueryBuilder) MetricNamespace(metricNamespace string) *MetricDefinitionsQueryBuilder
```

### <span class="badge object-method"></span> RawQuery

```go
func (builder *MetricDefinitionsQueryBuilder) RawQuery(rawQuery string) *MetricDefinitionsQueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

```go
func (builder *MetricDefinitionsQueryBuilder) ResourceGroup(resourceGroup string) *MetricDefinitionsQueryBuilder
```

### <span class="badge object-method"></span> ResourceName

```go
func (builder *MetricDefinitionsQueryBuilder) ResourceName(resourceName string) *MetricDefinitionsQueryBuilder
```

### <span class="badge object-method"></span> Subscription

```go
func (builder *MetricDefinitionsQueryBuilder) Subscription(subscription string) *MetricDefinitionsQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricDefinitionsQuery](./object-MetricDefinitionsQuery.md)

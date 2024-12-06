---
title: <span class="badge builder"></span> ResourceNamesQueryBuilder
---
# <span class="badge builder"></span> ResourceNamesQueryBuilder

## Constructor

```go
func NewResourceNamesQueryBuilder() *ResourceNamesQueryBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ResourceNamesQueryBuilder) Build() (ResourceNamesQuery, error)
```

### <span class="badge object-method"></span> MetricNamespace

```go
func (builder *ResourceNamesQueryBuilder) MetricNamespace(metricNamespace string) *ResourceNamesQueryBuilder
```

### <span class="badge object-method"></span> RawQuery

```go
func (builder *ResourceNamesQueryBuilder) RawQuery(rawQuery string) *ResourceNamesQueryBuilder
```

### <span class="badge object-method"></span> ResourceGroup

```go
func (builder *ResourceNamesQueryBuilder) ResourceGroup(resourceGroup string) *ResourceNamesQueryBuilder
```

### <span class="badge object-method"></span> Subscription

```go
func (builder *ResourceNamesQueryBuilder) Subscription(subscription string) *ResourceNamesQueryBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ResourceNamesQuery](./object-ResourceNamesQuery.md)

---
title: <span class="badge builder"></span> MetricFindValueBuilder
---
# <span class="badge builder"></span> MetricFindValueBuilder

## Constructor

```go
func NewMetricFindValueBuilder() *MetricFindValueBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricFindValueBuilder) Build() (MetricFindValue, error)
```

### <span class="badge object-method"></span> Expandable

```go
func (builder *MetricFindValueBuilder) Expandable(expandable bool) *MetricFindValueBuilder
```

### <span class="badge object-method"></span> Group

```go
func (builder *MetricFindValueBuilder) Group(group string) *MetricFindValueBuilder
```

### <span class="badge object-method"></span> Text

```go
func (builder *MetricFindValueBuilder) Text(text string) *MetricFindValueBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *MetricFindValueBuilder) Value(value dashboardv2beta1.StringOrFloat64) *MetricFindValueBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricFindValue](./object-MetricFindValue.md)

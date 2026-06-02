---
title: <span class="badge builder"></span> MetricDimensionBuilder
---
# <span class="badge builder"></span> MetricDimensionBuilder

## Constructor

```go
func NewMetricDimensionBuilder() *MetricDimensionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricDimensionBuilder) Build() (MetricDimension, error)
```

### <span class="badge object-method"></span> Dimension

Name of Dimension to be filtered on.

```go
func (builder *MetricDimensionBuilder) Dimension(dimension string) *MetricDimensionBuilder
```

### <span class="badge object-method"></span> Filter

@deprecated filter is deprecated in favour of filters to support multiselect.

```go
func (builder *MetricDimensionBuilder) Filter(filter string) *MetricDimensionBuilder
```

### <span class="badge object-method"></span> Filters

Values to match with the filter.

```go
func (builder *MetricDimensionBuilder) Filters(filters []string) *MetricDimensionBuilder
```

### <span class="badge object-method"></span> Operator

String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.

```go
func (builder *MetricDimensionBuilder) Operator(operator string) *MetricDimensionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricDimension](./object-MetricDimension.md)

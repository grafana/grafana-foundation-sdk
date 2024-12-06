---
title: <span class="badge builder"></span> AzureMetricDimensionBuilder
---
# <span class="badge builder"></span> AzureMetricDimensionBuilder

## Constructor

```go
func NewAzureMetricDimensionBuilder() *AzureMetricDimensionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AzureMetricDimensionBuilder) Build() (AzureMetricDimension, error)
```

### <span class="badge object-method"></span> Dimension

Name of Dimension to be filtered on.

```go
func (builder *AzureMetricDimensionBuilder) Dimension(dimension string) *AzureMetricDimensionBuilder
```

### <span class="badge object-method"></span> Filter

@deprecated filter is deprecated in favour of filters to support multiselect.

```go
func (builder *AzureMetricDimensionBuilder) Filter(filter string) *AzureMetricDimensionBuilder
```

### <span class="badge object-method"></span> Filters

Values to match with the filter.

```go
func (builder *AzureMetricDimensionBuilder) Filters(filters []string) *AzureMetricDimensionBuilder
```

### <span class="badge object-method"></span> Operator

String denoting the filter operation. Supports 'eq' - equals,'ne' - not equals, 'sw' - starts with. Note that some dimensions may not support all operators.

```go
func (builder *AzureMetricDimensionBuilder) Operator(operator string) *AzureMetricDimensionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AzureMetricDimension](./object-AzureMetricDimension.md)

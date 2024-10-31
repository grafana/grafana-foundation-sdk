---
title: <span class="badge builder"></span> MetricAggregationWithFieldBuilder
---
# <span class="badge builder"></span> MetricAggregationWithFieldBuilder

## Constructor

```go
func NewMetricAggregationWithFieldBuilder() *MetricAggregationWithFieldBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricAggregationWithFieldBuilder) Build() (MetricAggregationWithField, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *MetricAggregationWithFieldBuilder) Field(field string) *MetricAggregationWithFieldBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *MetricAggregationWithFieldBuilder) Hide(hide bool) *MetricAggregationWithFieldBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *MetricAggregationWithFieldBuilder) Id(id string) *MetricAggregationWithFieldBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *MetricAggregationWithFieldBuilder) Type(typeArg cog.Builder[elasticsearch.MetricAggregationType]) *MetricAggregationWithFieldBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricAggregationWithField](./object-MetricAggregationWithField.md)

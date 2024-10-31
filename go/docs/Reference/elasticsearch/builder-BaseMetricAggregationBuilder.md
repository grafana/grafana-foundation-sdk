---
title: <span class="badge builder"></span> BaseMetricAggregationBuilder
---
# <span class="badge builder"></span> BaseMetricAggregationBuilder

## Constructor

```go
func NewBaseMetricAggregationBuilder() *BaseMetricAggregationBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BaseMetricAggregationBuilder) Build() (BaseMetricAggregation, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *BaseMetricAggregationBuilder) Hide(hide bool) *BaseMetricAggregationBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *BaseMetricAggregationBuilder) Id(id string) *BaseMetricAggregationBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *BaseMetricAggregationBuilder) Type(typeArg cog.Builder[elasticsearch.MetricAggregationType]) *BaseMetricAggregationBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BaseMetricAggregation](./object-BaseMetricAggregation.md)

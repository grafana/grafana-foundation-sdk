---
title: <span class="badge builder"></span> MetricAggregationWithMissingSupportBuilder
---
# <span class="badge builder"></span> MetricAggregationWithMissingSupportBuilder

## Constructor

```go
func NewMetricAggregationWithMissingSupportBuilder() *MetricAggregationWithMissingSupportBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricAggregationWithMissingSupportBuilder) Build() (MetricAggregationWithMissingSupport, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *MetricAggregationWithMissingSupportBuilder) Hide(hide bool) *MetricAggregationWithMissingSupportBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *MetricAggregationWithMissingSupportBuilder) Id(id string) *MetricAggregationWithMissingSupportBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MetricAggregationWithMissingSupportBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchMetricAggregationWithMissingSupportSettings]) *MetricAggregationWithMissingSupportBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *MetricAggregationWithMissingSupportBuilder) Type(typeArg cog.Builder[elasticsearch.MetricAggregationType]) *MetricAggregationWithMissingSupportBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricAggregationWithMissingSupport](./object-MetricAggregationWithMissingSupport.md)

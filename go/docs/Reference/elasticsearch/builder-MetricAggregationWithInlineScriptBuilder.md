---
title: <span class="badge builder"></span> MetricAggregationWithInlineScriptBuilder
---
# <span class="badge builder"></span> MetricAggregationWithInlineScriptBuilder

## Constructor

```go
func NewMetricAggregationWithInlineScriptBuilder() *MetricAggregationWithInlineScriptBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MetricAggregationWithInlineScriptBuilder) Build() (MetricAggregationWithInlineScript, error)
```

### <span class="badge object-method"></span> Hide

```go
func (builder *MetricAggregationWithInlineScriptBuilder) Hide(hide bool) *MetricAggregationWithInlineScriptBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *MetricAggregationWithInlineScriptBuilder) Id(id string) *MetricAggregationWithInlineScriptBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MetricAggregationWithInlineScriptBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchMetricAggregationWithInlineScriptSettings]) *MetricAggregationWithInlineScriptBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *MetricAggregationWithInlineScriptBuilder) Type(typeArg cog.Builder[elasticsearch.MetricAggregationType]) *MetricAggregationWithInlineScriptBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MetricAggregationWithInlineScript](./object-MetricAggregationWithInlineScript.md)

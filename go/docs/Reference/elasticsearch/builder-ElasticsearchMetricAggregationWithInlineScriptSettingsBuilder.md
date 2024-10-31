---
title: <span class="badge builder"></span> ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder
---
# <span class="badge builder"></span> ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder

## Constructor

```go
func NewElasticsearchMetricAggregationWithInlineScriptSettingsBuilder() *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) Build() (ElasticsearchMetricAggregationWithInlineScriptSettings, error)
```

### <span class="badge object-method"></span> Script

```go
func (builder *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder) Script(script cog.Builder[elasticsearch.InlineScript]) *ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ElasticsearchMetricAggregationWithInlineScriptSettings](./object-ElasticsearchMetricAggregationWithInlineScriptSettings.md)

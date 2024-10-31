---
title: <span class="badge object-type-struct"></span> ElasticsearchMetricAggregationWithInlineScriptSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchMetricAggregationWithInlineScriptSettings

## Definition

```go
type ElasticsearchMetricAggregationWithInlineScriptSettings struct {
    Script *elasticsearch.InlineScript `json:"script,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMetricAggregationWithInlineScriptSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchMetricAggregationWithInlineScriptSettings *ElasticsearchMetricAggregationWithInlineScriptSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchMetricAggregationWithInlineScriptSettings` objects.

```go
func (elasticsearchMetricAggregationWithInlineScriptSettings *ElasticsearchMetricAggregationWithInlineScriptSettings) Equals(other ElasticsearchMetricAggregationWithInlineScriptSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchMetricAggregationWithInlineScriptSettings` fields for violations and returns them.

```go
func (elasticsearchMetricAggregationWithInlineScriptSettings *ElasticsearchMetricAggregationWithInlineScriptSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder](./builder-ElasticsearchMetricAggregationWithInlineScriptSettingsBuilder.md)

---
title: <span class="badge object-type-struct"></span> MetricAggregationWithInlineScript
---
# <span class="badge object-type-struct"></span> MetricAggregationWithInlineScript

## Definition

```go
type MetricAggregationWithInlineScript struct {
    Settings *elasticsearch.ElasticsearchMetricAggregationWithInlineScriptSettings `json:"settings,omitempty"`
    Type elasticsearch.MetricAggregationType `json:"type"`
    Id string `json:"id"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricAggregationWithInlineScript` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricAggregationWithInlineScript *MetricAggregationWithInlineScript) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricAggregationWithInlineScript` objects.

```go
func (metricAggregationWithInlineScript *MetricAggregationWithInlineScript) Equals(other MetricAggregationWithInlineScript) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricAggregationWithInlineScript` fields for violations and returns them.

```go
func (metricAggregationWithInlineScript *MetricAggregationWithInlineScript) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricAggregationWithInlineScriptBuilder](./builder-MetricAggregationWithInlineScriptBuilder.md)

---
title: <span class="badge object-type-struct"></span> MetricAggregationWithMissingSupport
---
# <span class="badge object-type-struct"></span> MetricAggregationWithMissingSupport

## Definition

```go
type MetricAggregationWithMissingSupport struct {
    Settings *elasticsearch.ElasticsearchMetricAggregationWithMissingSupportSettings `json:"settings,omitempty"`
    Type elasticsearch.MetricAggregationType `json:"type"`
    Id string `json:"id"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MetricAggregationWithMissingSupport` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (metricAggregationWithMissingSupport *MetricAggregationWithMissingSupport) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MetricAggregationWithMissingSupport` objects.

```go
func (metricAggregationWithMissingSupport *MetricAggregationWithMissingSupport) Equals(other MetricAggregationWithMissingSupport) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MetricAggregationWithMissingSupport` fields for violations and returns them.

```go
func (metricAggregationWithMissingSupport *MetricAggregationWithMissingSupport) Validate() error
```

## See also

 * <span class="badge builder"></span> [MetricAggregationWithMissingSupportBuilder](./builder-MetricAggregationWithMissingSupportBuilder.md)

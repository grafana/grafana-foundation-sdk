---
title: <span class="badge object-type-struct"></span> ElasticsearchMetricAggregationWithMissingSupportSettings
---
# <span class="badge object-type-struct"></span> ElasticsearchMetricAggregationWithMissingSupportSettings

## Definition

```go
type ElasticsearchMetricAggregationWithMissingSupportSettings struct {
    Missing *string `json:"missing,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ElasticsearchMetricAggregationWithMissingSupportSettings` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (elasticsearchMetricAggregationWithMissingSupportSettings *ElasticsearchMetricAggregationWithMissingSupportSettings) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ElasticsearchMetricAggregationWithMissingSupportSettings` objects.

```go
func (elasticsearchMetricAggregationWithMissingSupportSettings *ElasticsearchMetricAggregationWithMissingSupportSettings) Equals(other ElasticsearchMetricAggregationWithMissingSupportSettings) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ElasticsearchMetricAggregationWithMissingSupportSettings` fields for violations and returns them.

```go
func (elasticsearchMetricAggregationWithMissingSupportSettings *ElasticsearchMetricAggregationWithMissingSupportSettings) Validate() error
```

## See also

 * <span class="badge builder"></span> [ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder](./builder-ElasticsearchMetricAggregationWithMissingSupportSettingsBuilder.md)

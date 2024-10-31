---
title: <span class="badge object-type-struct"></span> TopMetrics
---
# <span class="badge object-type-struct"></span> TopMetrics

## Definition

```go
type TopMetrics struct {
    Type string `json:"type"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchTopMetricsSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `TopMetrics` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (topMetrics *TopMetrics) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `TopMetrics` objects.

```go
func (topMetrics *TopMetrics) Equals(other TopMetrics) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `TopMetrics` fields for violations and returns them.

```go
func (topMetrics *TopMetrics) Validate() error
```

## See also

 * <span class="badge builder"></span> [TopMetricsBuilder](./builder-TopMetricsBuilder.md)

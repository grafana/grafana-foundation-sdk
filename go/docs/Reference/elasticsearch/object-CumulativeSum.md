---
title: <span class="badge object-type-struct"></span> CumulativeSum
---
# <span class="badge object-type-struct"></span> CumulativeSum

## Definition

```go
type CumulativeSum struct {
    PipelineAgg *string `json:"pipelineAgg,omitempty"`
    Field *string `json:"field,omitempty"`
    Type string `json:"type"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchCumulativeSumSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CumulativeSum` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cumulativeSum *CumulativeSum) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CumulativeSum` objects.

```go
func (cumulativeSum *CumulativeSum) Equals(other CumulativeSum) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CumulativeSum` fields for violations and returns them.

```go
func (cumulativeSum *CumulativeSum) Validate() error
```

## See also

 * <span class="badge builder"></span> [CumulativeSumBuilder](./builder-CumulativeSumBuilder.md)

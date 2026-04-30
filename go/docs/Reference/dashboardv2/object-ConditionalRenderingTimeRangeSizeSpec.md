---
title: <span class="badge object-type-struct"></span> ConditionalRenderingTimeRangeSizeSpec
---
# <span class="badge object-type-struct"></span> ConditionalRenderingTimeRangeSizeSpec

## Definition

```go
type ConditionalRenderingTimeRangeSizeSpec struct {
    Value string `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingTimeRangeSizeSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (conditionalRenderingTimeRangeSizeSpec *ConditionalRenderingTimeRangeSizeSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConditionalRenderingTimeRangeSizeSpec` objects.

```go
func (conditionalRenderingTimeRangeSizeSpec *ConditionalRenderingTimeRangeSizeSpec) Equals(other ConditionalRenderingTimeRangeSizeSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConditionalRenderingTimeRangeSizeSpec` fields for violations and returns them.

```go
func (conditionalRenderingTimeRangeSizeSpec *ConditionalRenderingTimeRangeSizeSpec) Validate() error
```


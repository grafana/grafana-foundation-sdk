---
title: <span class="badge object-type-struct"></span> ConditionalRenderingDataSpec
---
# <span class="badge object-type-struct"></span> ConditionalRenderingDataSpec

## Definition

```go
type ConditionalRenderingDataSpec struct {
    Value bool `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingDataSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (conditionalRenderingDataSpec *ConditionalRenderingDataSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConditionalRenderingDataSpec` objects.

```go
func (conditionalRenderingDataSpec *ConditionalRenderingDataSpec) Equals(other ConditionalRenderingDataSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConditionalRenderingDataSpec` fields for violations and returns them.

```go
func (conditionalRenderingDataSpec *ConditionalRenderingDataSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [ConditionalRenderingDataSpecBuilder](./builder-ConditionalRenderingDataSpecBuilder.md)

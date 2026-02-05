---
title: <span class="badge object-type-struct"></span> ConditionalRenderingGroupKind
---
# <span class="badge object-type-struct"></span> ConditionalRenderingGroupKind

## Definition

```go
type ConditionalRenderingGroupKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.ConditionalRenderingGroupSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingGroupKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (conditionalRenderingGroupKind *ConditionalRenderingGroupKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConditionalRenderingGroupKind` objects.

```go
func (conditionalRenderingGroupKind *ConditionalRenderingGroupKind) Equals(other ConditionalRenderingGroupKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConditionalRenderingGroupKind` fields for violations and returns them.

```go
func (conditionalRenderingGroupKind *ConditionalRenderingGroupKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [ConditionalRenderingGroupBuilder](./builder-ConditionalRenderingGroupBuilder.md)

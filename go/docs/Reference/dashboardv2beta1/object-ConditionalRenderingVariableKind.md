---
title: <span class="badge object-type-struct"></span> ConditionalRenderingVariableKind
---
# <span class="badge object-type-struct"></span> ConditionalRenderingVariableKind

## Definition

```go
type ConditionalRenderingVariableKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.ConditionalRenderingVariableSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingVariableKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (conditionalRenderingVariableKind *ConditionalRenderingVariableKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConditionalRenderingVariableKind` objects.

```go
func (conditionalRenderingVariableKind *ConditionalRenderingVariableKind) Equals(other ConditionalRenderingVariableKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConditionalRenderingVariableKind` fields for violations and returns them.

```go
func (conditionalRenderingVariableKind *ConditionalRenderingVariableKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [ConditionalRenderingVariableBuilder](./builder-ConditionalRenderingVariableBuilder.md)

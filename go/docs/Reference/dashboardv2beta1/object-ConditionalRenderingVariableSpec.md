---
title: <span class="badge object-type-struct"></span> ConditionalRenderingVariableSpec
---
# <span class="badge object-type-struct"></span> ConditionalRenderingVariableSpec

## Definition

```go
type ConditionalRenderingVariableSpec struct {
    Variable string `json:"variable"`
    Operator dashboardv2beta1.ConditionalRenderingVariableSpecOperator `json:"operator"`
    Value string `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingVariableSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (conditionalRenderingVariableSpec *ConditionalRenderingVariableSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConditionalRenderingVariableSpec` objects.

```go
func (conditionalRenderingVariableSpec *ConditionalRenderingVariableSpec) Equals(other ConditionalRenderingVariableSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConditionalRenderingVariableSpec` fields for violations and returns them.

```go
func (conditionalRenderingVariableSpec *ConditionalRenderingVariableSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [ConditionalRenderingVariableSpecBuilder](./builder-ConditionalRenderingVariableSpecBuilder.md)

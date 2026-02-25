---
title: <span class="badge object-type-struct"></span> ConditionalRenderingGroupSpec
---
# <span class="badge object-type-struct"></span> ConditionalRenderingGroupSpec

## Definition

```go
type ConditionalRenderingGroupSpec struct {
    Visibility dashboardv2beta1.ConditionalRenderingGroupSpecVisibility `json:"visibility"`
    Condition dashboardv2beta1.ConditionalRenderingGroupSpecCondition `json:"condition"`
    Items []dashboardv2beta1.ConditionalRenderingVariableKindOrConditionalRenderingDataKindOrConditionalRenderingTimeRangeSizeKind `json:"items"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `ConditionalRenderingGroupSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (conditionalRenderingGroupSpec *ConditionalRenderingGroupSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `ConditionalRenderingGroupSpec` objects.

```go
func (conditionalRenderingGroupSpec *ConditionalRenderingGroupSpec) Equals(other ConditionalRenderingGroupSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `ConditionalRenderingGroupSpec` fields for violations and returns them.

```go
func (conditionalRenderingGroupSpec *ConditionalRenderingGroupSpec) Validate() error
```


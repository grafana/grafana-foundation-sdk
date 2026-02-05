---
title: <span class="badge object-type-struct"></span> AutoGridLayoutItemSpec
---
# <span class="badge object-type-struct"></span> AutoGridLayoutItemSpec

## Definition

```go
type AutoGridLayoutItemSpec struct {
    Element dashboardv2beta1.ElementReference `json:"element"`
    Repeat *dashboardv2beta1.AutoGridRepeatOptions `json:"repeat,omitempty"`
    ConditionalRendering *dashboardv2beta1.ConditionalRenderingGroupKind `json:"conditionalRendering,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutItemSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (autoGridLayoutItemSpec *AutoGridLayoutItemSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AutoGridLayoutItemSpec` objects.

```go
func (autoGridLayoutItemSpec *AutoGridLayoutItemSpec) Equals(other AutoGridLayoutItemSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AutoGridLayoutItemSpec` fields for violations and returns them.

```go
func (autoGridLayoutItemSpec *AutoGridLayoutItemSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [AutoGridLayoutItemSpecBuilder](./builder-AutoGridLayoutItemSpecBuilder.md)

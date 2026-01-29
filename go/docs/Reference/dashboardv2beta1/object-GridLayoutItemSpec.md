---
title: <span class="badge object-type-struct"></span> GridLayoutItemSpec
---
# <span class="badge object-type-struct"></span> GridLayoutItemSpec

## Definition

```go
type GridLayoutItemSpec struct {
    X int64 `json:"x"`
    Y int64 `json:"y"`
    Width int64 `json:"width"`
    Height int64 `json:"height"`
    // reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
    Element dashboardv2beta1.ElementReference `json:"element"`
    Repeat *dashboardv2beta1.RepeatOptions `json:"repeat,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutItemSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (gridLayoutItemSpec *GridLayoutItemSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GridLayoutItemSpec` objects.

```go
func (gridLayoutItemSpec *GridLayoutItemSpec) Equals(other GridLayoutItemSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GridLayoutItemSpec` fields for violations and returns them.

```go
func (gridLayoutItemSpec *GridLayoutItemSpec) Validate() error
```


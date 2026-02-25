---
title: <span class="badge object-type-struct"></span> GridLayoutItemKind
---
# <span class="badge object-type-struct"></span> GridLayoutItemKind

## Definition

```go
type GridLayoutItemKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.GridLayoutItemSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `GridLayoutItemKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (gridLayoutItemKind *GridLayoutItemKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `GridLayoutItemKind` objects.

```go
func (gridLayoutItemKind *GridLayoutItemKind) Equals(other GridLayoutItemKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `GridLayoutItemKind` fields for violations and returns them.

```go
func (gridLayoutItemKind *GridLayoutItemKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [GridItemBuilder](./builder-GridItemBuilder.md)

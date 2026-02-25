---
title: <span class="badge object-type-struct"></span> AutoGridLayoutItemKind
---
# <span class="badge object-type-struct"></span> AutoGridLayoutItemKind

## Definition

```go
type AutoGridLayoutItemKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.AutoGridLayoutItemSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AutoGridLayoutItemKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (autoGridLayoutItemKind *AutoGridLayoutItemKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AutoGridLayoutItemKind` objects.

```go
func (autoGridLayoutItemKind *AutoGridLayoutItemKind) Equals(other AutoGridLayoutItemKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AutoGridLayoutItemKind` fields for violations and returns them.

```go
func (autoGridLayoutItemKind *AutoGridLayoutItemKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [AutoGridItemBuilder](./builder-AutoGridItemBuilder.md)

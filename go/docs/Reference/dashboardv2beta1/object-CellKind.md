---
title: <span class="badge object-type-struct"></span> CellKind
---
# <span class="badge object-type-struct"></span> CellKind

A cell holds non-panel narrative content (markdown text, code) in a notebook layout.

Panel cells are not represented here — they reuse V2PanelKind.

## Definition

```go
type CellKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.CellSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `CellKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (cellKind *CellKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `CellKind` objects.

```go
func (cellKind *CellKind) Equals(other CellKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `CellKind` fields for violations and returns them.

```go
func (cellKind *CellKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [CellKindBuilder](./builder-CellKindBuilder.md)

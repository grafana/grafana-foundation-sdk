---
title: <span class="badge object-type-struct"></span> PanelQueryKind
---
# <span class="badge object-type-struct"></span> PanelQueryKind

## Definition

```go
type PanelQueryKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.PanelQuerySpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelQueryKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (panelQueryKind *PanelQueryKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PanelQueryKind` objects.

```go
func (panelQueryKind *PanelQueryKind) Equals(other PanelQueryKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PanelQueryKind` fields for violations and returns them.

```go
func (panelQueryKind *PanelQueryKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [TargetBuilder](./builder-TargetBuilder.md)

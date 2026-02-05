---
title: <span class="badge object-type-struct"></span> PanelKind
---
# <span class="badge object-type-struct"></span> PanelKind

## Definition

```go
type PanelKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.PanelSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (panelKind *PanelKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PanelKind` objects.

```go
func (panelKind *PanelKind) Equals(other PanelKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PanelKind` fields for violations and returns them.

```go
func (panelKind *PanelKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [PanelBuilder](./builder-PanelBuilder.md)

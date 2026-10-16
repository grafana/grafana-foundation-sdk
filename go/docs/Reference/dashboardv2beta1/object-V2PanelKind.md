---
title: <span class="badge object-type-struct"></span> V2PanelKind
---
# <span class="badge object-type-struct"></span> V2PanelKind

The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,

which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has

to be forked rather than shared because PanelKind reaches TransformationKind through

QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not

change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.

## Definition

```go
type V2PanelKind struct {
    Kind string `json:"kind"`
    Spec dashboardv2beta1.V2PanelSpec `json:"spec"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `V2PanelKind` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (v2PanelKind *V2PanelKind) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `V2PanelKind` objects.

```go
func (v2PanelKind *V2PanelKind) Equals(other V2PanelKind) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `V2PanelKind` fields for violations and returns them.

```go
func (v2PanelKind *V2PanelKind) Validate() error
```

## See also

 * <span class="badge builder"></span> [V2PanelKindBuilder](./builder-V2PanelKindBuilder.md)

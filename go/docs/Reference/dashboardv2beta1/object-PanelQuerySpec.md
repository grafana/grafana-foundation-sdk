---
title: <span class="badge object-type-struct"></span> PanelQuerySpec
---
# <span class="badge object-type-struct"></span> PanelQuerySpec

## Definition

```go
type PanelQuerySpec struct {
    Query dashboardv2beta1.DataQueryKind `json:"query"`
    RefId string `json:"refId"`
    Hidden bool `json:"hidden"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelQuerySpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (panelQuerySpec *PanelQuerySpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PanelQuerySpec` objects.

```go
func (panelQuerySpec *PanelQuerySpec) Equals(other PanelQuerySpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PanelQuerySpec` fields for violations and returns them.

```go
func (panelQuerySpec *PanelQuerySpec) Validate() error
```


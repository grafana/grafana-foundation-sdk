---
title: <span class="badge object-type-struct"></span> PanelSpec
---
# <span class="badge object-type-struct"></span> PanelSpec

## Definition

```go
type PanelSpec struct {
    Id float64 `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    Links []dashboardv2beta1.DataLink `json:"links"`
    Data dashboardv2beta1.QueryGroupKind `json:"data"`
    VizConfig dashboardv2beta1.VizConfigKind `json:"vizConfig"`
    Transparent *bool `json:"transparent,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (panelSpec *PanelSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PanelSpec` objects.

```go
func (panelSpec *PanelSpec) Equals(other PanelSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PanelSpec` fields for violations and returns them.

```go
func (panelSpec *PanelSpec) Validate() error
```


---
title: <span class="badge object-type-struct"></span> V2PanelSpec
---
# <span class="badge object-type-struct"></span> V2PanelSpec

## Definition

```go
type V2PanelSpec struct {
    Id float64 `json:"id"`
    Title string `json:"title"`
    // Shown in a info icon tooltip next to panel title
    Description *string `json:"description,omitempty"`
    // Shown in a sub header below the title.
    Subtitle *string `json:"subtitle,omitempty"`
    Links []dashboardv2beta1.DataLink `json:"links"`
    Data dashboardv2beta1.V2QueryGroupKind `json:"data"`
    VizConfig dashboardv2beta1.VizConfigKind `json:"vizConfig"`
    Transparent *bool `json:"transparent,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `V2PanelSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (v2PanelSpec *V2PanelSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `V2PanelSpec` objects.

```go
func (v2PanelSpec *V2PanelSpec) Equals(other V2PanelSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `V2PanelSpec` fields for violations and returns them.

```go
func (v2PanelSpec *V2PanelSpec) Validate() error
```

## See also

 * <span class="badge builder"></span> [V2PanelSpecBuilder](./builder-V2PanelSpecBuilder.md)

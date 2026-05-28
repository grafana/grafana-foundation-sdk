---
title: <span class="badge object-type-struct"></span> RowsLayoutRowSpec
---
# <span class="badge object-type-struct"></span> RowsLayoutRowSpec

## Definition

```go
type RowsLayoutRowSpec struct {
    Title *string `json:"title,omitempty"`
    Collapse *bool `json:"collapse,omitempty"`
    HideHeader *bool `json:"hideHeader,omitempty"`
    FillScreen *bool `json:"fillScreen,omitempty"`
    ConditionalRendering *dashboardv2.ConditionalRenderingGroupKind `json:"conditionalRendering,omitempty"`
    Repeat *dashboardv2.RowRepeatOptions `json:"repeat,omitempty"`
    Layout dashboardv2.GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind `json:"layout"`
    Variables []dashboardv2.VariableKind `json:"variables,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `RowsLayoutRowSpec` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (rowsLayoutRowSpec *RowsLayoutRowSpec) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `RowsLayoutRowSpec` objects.

```go
func (rowsLayoutRowSpec *RowsLayoutRowSpec) Equals(other RowsLayoutRowSpec) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `RowsLayoutRowSpec` fields for violations and returns them.

```go
func (rowsLayoutRowSpec *RowsLayoutRowSpec) Validate() error
```


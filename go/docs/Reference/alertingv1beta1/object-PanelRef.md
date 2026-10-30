---
title: <span class="badge object-type-struct"></span> PanelRef
---
# <span class="badge object-type-struct"></span> PanelRef

## Definition

```go
type PanelRef struct {
    DashboardUID string `json:"dashboardUID"`
    PanelID int64 `json:"panelID"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PanelRef` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (panelRef *PanelRef) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PanelRef` objects.

```go
func (panelRef *PanelRef) Equals(other PanelRef) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PanelRef` fields for violations and returns them.

```go
func (panelRef *PanelRef) Validate() error
```

## See also

 * <span class="badge builder"></span> [PanelRefBuilder](./builder-PanelRefBuilder.md)

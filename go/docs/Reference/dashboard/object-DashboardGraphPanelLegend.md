---
title: <span class="badge object-type-struct"></span> DashboardGraphPanelLegend
---
# <span class="badge object-type-struct"></span> DashboardGraphPanelLegend

## Definition

```go
type DashboardGraphPanelLegend struct {
    Show bool `json:"show"`
    Sort *string `json:"sort,omitempty"`
    SortDesc *bool `json:"sortDesc,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardGraphPanelLegend` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardGraphPanelLegend *DashboardGraphPanelLegend) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardGraphPanelLegend` objects.

```go
func (dashboardGraphPanelLegend *DashboardGraphPanelLegend) Equals(other DashboardGraphPanelLegend) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardGraphPanelLegend` fields for violations and returns them.

```go
func (dashboardGraphPanelLegend *DashboardGraphPanelLegend) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardGraphPanelLegendBuilder](./builder-DashboardGraphPanelLegendBuilder.md)

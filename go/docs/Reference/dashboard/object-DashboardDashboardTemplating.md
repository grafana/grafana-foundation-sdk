---
title: <span class="badge object-type-struct"></span> DashboardDashboardTemplating
---
# <span class="badge object-type-struct"></span> DashboardDashboardTemplating

## Definition

```go
type DashboardDashboardTemplating struct {
    // List of configured template variables with their saved values along with some other metadata
    List []dashboard.VariableModel `json:"list,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardDashboardTemplating` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardDashboardTemplating *DashboardDashboardTemplating) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardDashboardTemplating` objects.

```go
func (dashboardDashboardTemplating *DashboardDashboardTemplating) Equals(other DashboardDashboardTemplating) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardDashboardTemplating` fields for violations and returns them.

```go
func (dashboardDashboardTemplating *DashboardDashboardTemplating) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardDashboardTemplatingBuilder](./builder-DashboardDashboardTemplatingBuilder.md)

---
title: <span class="badge object-type-struct"></span> DashboardDashboardTime
---
# <span class="badge object-type-struct"></span> DashboardDashboardTime

## Definition

```go
type DashboardDashboardTime struct {
    From string `json:"from"`
    To string `json:"to"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardDashboardTime` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardDashboardTime *DashboardDashboardTime) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardDashboardTime` objects.

```go
func (dashboardDashboardTime *DashboardDashboardTime) Equals(other DashboardDashboardTime) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardDashboardTime` fields for violations and returns them.

```go
func (dashboardDashboardTime *DashboardDashboardTime) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardDashboardTimeBuilder](./builder-DashboardDashboardTimeBuilder.md)

---
title: <span class="badge object-type-struct"></span> DashboardRangeMapOptions
---
# <span class="badge object-type-struct"></span> DashboardRangeMapOptions

## Definition

```go
type DashboardRangeMapOptions struct {
    // Min value of the range. It can be null which means -Infinity
    From *float64 `json:"from"`
    // Max value of the range. It can be null which means +Infinity
    To *float64 `json:"to"`
    // Config to apply when the value is within the range
    Result dashboard.ValueMappingResult `json:"result"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardRangeMapOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardRangeMapOptions *DashboardRangeMapOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardRangeMapOptions` objects.

```go
func (dashboardRangeMapOptions *DashboardRangeMapOptions) Equals(other DashboardRangeMapOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardRangeMapOptions` fields for violations and returns them.

```go
func (dashboardRangeMapOptions *DashboardRangeMapOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardRangeMapOptionsBuilder](./builder-DashboardRangeMapOptionsBuilder.md)

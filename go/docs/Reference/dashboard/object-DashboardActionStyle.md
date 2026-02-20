---
title: <span class="badge object-type-struct"></span> DashboardActionStyle
---
# <span class="badge object-type-struct"></span> DashboardActionStyle

## Definition

```go
type DashboardActionStyle struct {
    BackgroundColor *string `json:"backgroundColor,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardActionStyle` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardActionStyle *DashboardActionStyle) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardActionStyle` objects.

```go
func (dashboardActionStyle *DashboardActionStyle) Equals(other DashboardActionStyle) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardActionStyle` fields for violations and returns them.

```go
func (dashboardActionStyle *DashboardActionStyle) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardActionStyleBuilder](./builder-DashboardActionStyleBuilder.md)

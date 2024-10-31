---
title: <span class="badge object-type-struct"></span> DashboardSpecialValueMapOptions
---
# <span class="badge object-type-struct"></span> DashboardSpecialValueMapOptions

## Definition

```go
type DashboardSpecialValueMapOptions struct {
    // Special value to match against
    Match dashboard.SpecialValueMatch `json:"match"`
    // Config to apply when the value matches the special value
    Result dashboard.ValueMappingResult `json:"result"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardSpecialValueMapOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardSpecialValueMapOptions *DashboardSpecialValueMapOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardSpecialValueMapOptions` objects.

```go
func (dashboardSpecialValueMapOptions *DashboardSpecialValueMapOptions) Equals(other DashboardSpecialValueMapOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardSpecialValueMapOptions` fields for violations and returns them.

```go
func (dashboardSpecialValueMapOptions *DashboardSpecialValueMapOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardSpecialValueMapOptionsBuilder](./builder-DashboardSpecialValueMapOptionsBuilder.md)

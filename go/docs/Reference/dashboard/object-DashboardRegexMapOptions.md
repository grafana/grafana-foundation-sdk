---
title: <span class="badge object-type-struct"></span> DashboardRegexMapOptions
---
# <span class="badge object-type-struct"></span> DashboardRegexMapOptions

## Definition

```go
type DashboardRegexMapOptions struct {
    // Regular expression to match against
    Pattern string `json:"pattern"`
    // Config to apply when the value matches the regex
    Result dashboard.ValueMappingResult `json:"result"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardRegexMapOptions` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardRegexMapOptions *DashboardRegexMapOptions) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardRegexMapOptions` objects.

```go
func (dashboardRegexMapOptions *DashboardRegexMapOptions) Equals(other DashboardRegexMapOptions) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardRegexMapOptions` fields for violations and returns them.

```go
func (dashboardRegexMapOptions *DashboardRegexMapOptions) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardRegexMapOptionsBuilder](./builder-DashboardRegexMapOptionsBuilder.md)

---
title: <span class="badge object-type-struct"></span> DashboardFieldConfigSourceOverrides
---
# <span class="badge object-type-struct"></span> DashboardFieldConfigSourceOverrides

## Definition

```go
type DashboardFieldConfigSourceOverrides struct {
    Matcher dashboard.MatcherConfig `json:"matcher"`
    Properties []dashboard.DynamicConfigValue `json:"properties"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardFieldConfigSourceOverrides` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardFieldConfigSourceOverrides *DashboardFieldConfigSourceOverrides) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardFieldConfigSourceOverrides` objects.

```go
func (dashboardFieldConfigSourceOverrides *DashboardFieldConfigSourceOverrides) Equals(other DashboardFieldConfigSourceOverrides) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardFieldConfigSourceOverrides` fields for violations and returns them.

```go
func (dashboardFieldConfigSourceOverrides *DashboardFieldConfigSourceOverrides) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardFieldConfigSourceOverridesBuilder](./builder-DashboardFieldConfigSourceOverridesBuilder.md)

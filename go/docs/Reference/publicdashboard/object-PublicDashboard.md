---
title: <span class="badge object-type-struct"></span> PublicDashboard
---
# <span class="badge object-type-struct"></span> PublicDashboard

## Definition

```go
type PublicDashboard struct {
    // Unique public dashboard identifier
    Uid string `json:"uid"`
    // Dashboard unique identifier referenced by this public dashboard
    DashboardUid string `json:"dashboardUid"`
    // Unique public access token
    AccessToken *string `json:"accessToken,omitempty"`
    // Flag that indicates if the public dashboard is enabled
    IsEnabled bool `json:"isEnabled"`
    // Flag that indicates if annotations are enabled
    AnnotationsEnabled bool `json:"annotationsEnabled"`
    // Flag that indicates if the time range picker is enabled
    TimeSelectionEnabled bool `json:"timeSelectionEnabled"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PublicDashboard` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (publicDashboard *PublicDashboard) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PublicDashboard` objects.

```go
func (publicDashboard *PublicDashboard) Equals(other PublicDashboard) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PublicDashboard` fields for violations and returns them.

```go
func (publicDashboard *PublicDashboard) Validate() error
```

## See also

 * <span class="badge builder"></span> [PublicDashboardBuilder](./builder-PublicDashboardBuilder.md)

---
title: <span class="badge builder"></span> PublicDashboardBuilder
---
# <span class="badge builder"></span> PublicDashboardBuilder

## Constructor

```go
func NewPublicDashboardBuilder() *PublicDashboardBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PublicDashboardBuilder) Build() (PublicDashboard, error)
```

### <span class="badge object-method"></span> AccessToken

Unique public access token

```go
func (builder *PublicDashboardBuilder) AccessToken(accessToken string) *PublicDashboardBuilder
```

### <span class="badge object-method"></span> AnnotationsEnabled

Flag that indicates if annotations are enabled

```go
func (builder *PublicDashboardBuilder) AnnotationsEnabled(annotationsEnabled bool) *PublicDashboardBuilder
```

### <span class="badge object-method"></span> DashboardUid

Dashboard unique identifier referenced by this public dashboard

```go
func (builder *PublicDashboardBuilder) DashboardUid(dashboardUid string) *PublicDashboardBuilder
```

### <span class="badge object-method"></span> IsEnabled

Flag that indicates if the public dashboard is enabled

```go
func (builder *PublicDashboardBuilder) IsEnabled(isEnabled bool) *PublicDashboardBuilder
```

### <span class="badge object-method"></span> TimeSelectionEnabled

Flag that indicates if the time range picker is enabled

```go
func (builder *PublicDashboardBuilder) TimeSelectionEnabled(timeSelectionEnabled bool) *PublicDashboardBuilder
```

### <span class="badge object-method"></span> Uid

Unique public dashboard identifier

```go
func (builder *PublicDashboardBuilder) Uid(uid string) *PublicDashboardBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PublicDashboard](./object-PublicDashboard.md)

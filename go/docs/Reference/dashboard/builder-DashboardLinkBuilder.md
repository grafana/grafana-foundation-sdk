---
title: <span class="badge builder"></span> DashboardLinkBuilder
---
# <span class="badge builder"></span> DashboardLinkBuilder

## Constructor

```go
func NewDashboardLinkBuilder(title string) *DashboardLinkBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DashboardLinkBuilder) Build() (DashboardLink, error)
```

### <span class="badge object-method"></span> AsDropdown

If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards

```go
func (builder *DashboardLinkBuilder) AsDropdown(asDropdown bool) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> Icon

Icon name to be displayed with the link

```go
func (builder *DashboardLinkBuilder) Icon(icon string) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> IncludeVars

If true, includes current template variables values in the link as query params

```go
func (builder *DashboardLinkBuilder) IncludeVars(includeVars bool) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> KeepTime

If true, includes current time range in the link as query params

```go
func (builder *DashboardLinkBuilder) KeepTime(keepTime bool) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> Tags

List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards

```go
func (builder *DashboardLinkBuilder) Tags(tags []string) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> TargetBlank

If true, the link will be opened in a new tab

```go
func (builder *DashboardLinkBuilder) TargetBlank(targetBlank bool) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> Title

Title to display with the link

```go
func (builder *DashboardLinkBuilder) Title(title string) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> Tooltip

Tooltip to display when the user hovers their mouse over it

```go
func (builder *DashboardLinkBuilder) Tooltip(tooltip string) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> Type

Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)

```go
func (builder *DashboardLinkBuilder) Type(typeArg dashboard.DashboardLinkType) *DashboardLinkBuilder
```

### <span class="badge object-method"></span> Url

Link URL. Only required/valid if the type is link

```go
func (builder *DashboardLinkBuilder) Url(url string) *DashboardLinkBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [DashboardLink](./object-DashboardLink.md)

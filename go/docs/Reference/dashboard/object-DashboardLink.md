---
title: <span class="badge object-type-struct"></span> DashboardLink
---
# <span class="badge object-type-struct"></span> DashboardLink

Links with references to other dashboards or external resources

## Definition

```go
type DashboardLink struct {
    // Title to display with the link
    Title string `json:"title"`
    // Link type. Accepted values are dashboards (to refer to another dashboard) and link (to refer to an external resource)
    Type dashboard.DashboardLinkType `json:"type"`
    // Icon name to be displayed with the link
    Icon string `json:"icon"`
    // Tooltip to display when the user hovers their mouse over it
    Tooltip string `json:"tooltip"`
    // Link URL. Only required/valid if the type is link
    Url string `json:"url"`
    // List of tags to limit the linked dashboards. If empty, all dashboards will be displayed. Only valid if the type is dashboards
    Tags []string `json:"tags"`
    // If true, all dashboards links will be displayed in a dropdown. If false, all dashboards links will be displayed side by side. Only valid if the type is dashboards
    AsDropdown bool `json:"asDropdown"`
    // If true, the link will be opened in a new tab
    TargetBlank bool `json:"targetBlank"`
    // If true, includes current template variables values in the link as query params
    IncludeVars bool `json:"includeVars"`
    // If true, includes current time range in the link as query params
    KeepTime bool `json:"keepTime"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `DashboardLink` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboardLink *DashboardLink) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `DashboardLink` objects.

```go
func (dashboardLink *DashboardLink) Equals(other DashboardLink) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `DashboardLink` fields for violations and returns them.

```go
func (dashboardLink *DashboardLink) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardLinkBuilder](./builder-DashboardLinkBuilder.md)

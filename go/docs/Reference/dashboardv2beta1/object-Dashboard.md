---
title: <span class="badge object-type-struct"></span> Dashboard
---
# <span class="badge object-type-struct"></span> Dashboard

## Definition

```go
type Dashboard struct {
    Annotations []dashboardv2beta1.AnnotationQueryKind `json:"annotations"`
    // Configuration of dashboard cursor sync behavior.
    // "Off" for no shared crosshair or tooltip (default).
    // "Crosshair" for shared crosshair.
    // "Tooltip" for shared crosshair AND shared tooltip.
    CursorSync dashboardv2beta1.DashboardCursorSync `json:"cursorSync"`
    // Description of dashboard.
    Description *string `json:"description,omitempty"`
    // Whether a dashboard is editable or not.
    Editable *bool `json:"editable,omitempty"`
    Elements map[string]dashboardv2beta1.Element `json:"elements"`
    Layout dashboardv2beta1.GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind `json:"layout"`
    // Links with references to other dashboards or external websites.
    Links []dashboardv2beta1.DashboardLink `json:"links"`
    // When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    // This will keep data "moving left" regardless of the query refresh rate. This setting helps
    // avoid dashboards presenting stale live data.
    LiveNow *bool `json:"liveNow,omitempty"`
    // When set to true, the dashboard will load all panels in the dashboard when it's loaded.
    Preload bool `json:"preload"`
    // Plugins only. The version of the dashboard installed together with the plugin.
    // This is used to determine if the dashboard should be updated when the plugin is updated.
    Revision *uint16 `json:"revision,omitempty"`
    // Tags associated with dashboard.
    Tags []string `json:"tags"`
    TimeSettings dashboardv2beta1.TimeSettingsSpec `json:"timeSettings"`
    // Title of dashboard.
    Title string `json:"title"`
    // Configured template variables.
    Variables []dashboardv2beta1.VariableKind `json:"variables"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Dashboard` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (dashboard *Dashboard) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Dashboard` objects.

```go
func (dashboard *Dashboard) Equals(other Dashboard) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Dashboard` fields for violations and returns them.

```go
func (dashboard *Dashboard) Validate() error
```

## See also

 * <span class="badge builder"></span> [DashboardBuilder](./builder-DashboardBuilder.md)

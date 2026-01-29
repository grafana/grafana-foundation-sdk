---
title: <span class="badge object-type-class"></span> Dashboard
---
# <span class="badge object-type-class"></span> Dashboard

## Definition

```python
class Dashboard:
    annotations: list[dashboardv2beta1.AnnotationQueryKind]
    # Configuration of dashboard cursor sync behavior.
    # "Off" for no shared crosshair or tooltip (default).
    # "Crosshair" for shared crosshair.
    # "Tooltip" for shared crosshair AND shared tooltip.
    cursor_sync: dashboardv2beta1.DashboardCursorSync
    # Description of dashboard.
    description: typing.Optional[str]
    # Whether a dashboard is editable or not.
    editable: typing.Optional[bool]
    elements: dict[str, dashboardv2beta1.Element]
    layout: typing.Union[dashboardv2beta1.GridLayoutKind, dashboardv2beta1.RowsLayoutKind, dashboardv2beta1.AutoGridLayoutKind, dashboardv2beta1.TabsLayoutKind]
    # Links with references to other dashboards or external websites.
    links: list[dashboardv2beta1.DashboardLink]
    # When set to true, the dashboard will redraw panels at an interval matching the pixel width.
    # This will keep data "moving left" regardless of the query refresh rate. This setting helps
    # avoid dashboards presenting stale live data.
    live_now: typing.Optional[bool]
    # When set to true, the dashboard will load all panels in the dashboard when it's loaded.
    preload: bool
    # Plugins only. The version of the dashboard installed together with the plugin.
    # This is used to determine if the dashboard should be updated when the plugin is updated.
    revision: typing.Optional[int]
    # Tags associated with dashboard.
    tags: list[str]
    time_settings: dashboardv2beta1.TimeSettingsSpec
    # Title of dashboard.
    title: str
    # Configured template variables.
    variables: list[dashboardv2beta1.VariableKind]
```
## Methods

### <span class="badge object-method"></span> to_json

Converts this object into a representation that can easily be encoded to JSON.

```python
def to_json() -> dict[str, object]
```

### <span class="badge object-method"></span> from_json

Builds this object from a JSON-decoded dict.

```python
@classmethod
def from_json(data: dict[str, typing.Any]) -> typing.Self
```

## See also

 * <span class="badge builder"></span> [Dashboard](./builder-Dashboard.md)

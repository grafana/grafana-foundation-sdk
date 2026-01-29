---
title: <span class="badge builder"></span> Dashboard
---
# <span class="badge builder"></span> Dashboard

## Constructor

```python
Dashboard(title: str)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```python
def build() -> dashboardv2beta1.Dashboard
```

### <span class="badge object-method"></span> annotations

```python
def annotations(annotations: list[cogbuilder.Builder[dashboardv2beta1.AnnotationQueryKind]]) -> typing.Self
```

### <span class="badge object-method"></span> auto_grid_layout

```python
def auto_grid_layout(auto_grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.AutoGridLayoutKind]) -> typing.Self
```

### <span class="badge object-method"></span> cursor_sync

Configuration of dashboard cursor sync behavior.

"Off" for no shared crosshair or tooltip (default).

"Crosshair" for shared crosshair.

"Tooltip" for shared crosshair AND shared tooltip.

```python
def cursor_sync(cursor_sync: dashboardv2beta1.DashboardCursorSync) -> typing.Self
```

### <span class="badge object-method"></span> description

Description of dashboard.

```python
def description(description: str) -> typing.Self
```

### <span class="badge object-method"></span> editable

Whether a dashboard is editable or not.

```python
def editable(editable: bool) -> typing.Self
```

### <span class="badge object-method"></span> element

```python
def element(key: str, element: cogbuilder.Builder[dashboardv2beta1.Element]) -> typing.Self
```

### <span class="badge object-method"></span> elements

```python
def elements(elements: dict[str, cogbuilder.Builder[dashboardv2beta1.Element]]) -> typing.Self
```

### <span class="badge object-method"></span> grid_layout

```python
def grid_layout(grid_layout_kind: cogbuilder.Builder[dashboardv2beta1.GridLayoutKind]) -> typing.Self
```

### <span class="badge object-method"></span> links

Links with references to other dashboards or external websites.

```python
def links(links: list[cogbuilder.Builder[dashboardv2beta1.DashboardLink]]) -> typing.Self
```

### <span class="badge object-method"></span> live_now

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data.

```python
def live_now(live_now: bool) -> typing.Self
```

### <span class="badge object-method"></span> preload

When set to true, the dashboard will load all panels in the dashboard when it's loaded.

```python
def preload(preload: bool) -> typing.Self
```

### <span class="badge object-method"></span> revision

Plugins only. The version of the dashboard installed together with the plugin.

This is used to determine if the dashboard should be updated when the plugin is updated.

```python
def revision(revision: int) -> typing.Self
```

### <span class="badge object-method"></span> rows_layout

```python
def rows_layout(rows_layout_kind: cogbuilder.Builder[dashboardv2beta1.RowsLayoutKind]) -> typing.Self
```

### <span class="badge object-method"></span> tabs_layout

```python
def tabs_layout(tabs_layout_kind: cogbuilder.Builder[dashboardv2beta1.TabsLayoutKind]) -> typing.Self
```

### <span class="badge object-method"></span> tags

Tags associated with dashboard.

```python
def tags(tags: list[str]) -> typing.Self
```

### <span class="badge object-method"></span> time_settings

```python
def time_settings(time_settings: cogbuilder.Builder[dashboardv2beta1.TimeSettingsSpec]) -> typing.Self
```

### <span class="badge object-method"></span> title

Title of dashboard.

```python
def title(title: str) -> typing.Self
```

### <span class="badge object-method"></span> variable

Configured template variables.

```python
def variable(variable: cogbuilder.Builder[dashboardv2beta1.VariableKind]) -> typing.Self
```

### <span class="badge object-method"></span> variables

Configured template variables.

```python
def variables(variables: list[cogbuilder.Builder[dashboardv2beta1.VariableKind]]) -> typing.Self
```

## See also

 * <span class="badge object-type-class"></span> [Dashboard](./object-Dashboard.md)

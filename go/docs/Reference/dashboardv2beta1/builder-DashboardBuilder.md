---
title: <span class="badge builder"></span> DashboardBuilder
---
# <span class="badge builder"></span> DashboardBuilder

## Constructor

```go
func NewDashboardBuilder(title string) *DashboardBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DashboardBuilder) Build() (Dashboard, error)
```

### <span class="badge object-method"></span> SwitchVariableKind

```go
func (builder *DashboardBuilder) SwitchVariableKind(switchVariableKind cog.Builder[dashboardv2beta1.SwitchVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> AdhocVariable

```go
func (builder *DashboardBuilder) AdhocVariable(adhocVariableKind cog.Builder[dashboardv2beta1.AdhocVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> Annotations

```go
func (builder *DashboardBuilder) Annotations(annotations []cog.Builder[dashboardv2beta1.AnnotationQueryKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> AutoGridLayout

```go
func (builder *DashboardBuilder) AutoGridLayout(autoGridLayoutKind cog.Builder[dashboardv2beta1.AutoGridLayoutKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> ConstantVariable

```go
func (builder *DashboardBuilder) ConstantVariable(constantVariableKind cog.Builder[dashboardv2beta1.ConstantVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> CursorSync

Configuration of dashboard cursor sync behavior.

"Off" for no shared crosshair or tooltip (default).

"Crosshair" for shared crosshair.

"Tooltip" for shared crosshair AND shared tooltip.

```go
func (builder *DashboardBuilder) CursorSync(cursorSync dashboardv2beta1.DashboardCursorSync) *DashboardBuilder
```

### <span class="badge object-method"></span> CustomVariable

```go
func (builder *DashboardBuilder) CustomVariable(customVariableKind cog.Builder[dashboardv2beta1.CustomVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> DatasourceVariable

```go
func (builder *DashboardBuilder) DatasourceVariable(datasourceVariableKind cog.Builder[dashboardv2beta1.DatasourceVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> Description

Description of dashboard.

```go
func (builder *DashboardBuilder) Description(description string) *DashboardBuilder
```

### <span class="badge object-method"></span> Editable

Whether a dashboard is editable or not.

```go
func (builder *DashboardBuilder) Editable(editable bool) *DashboardBuilder
```

### <span class="badge object-method"></span> Elements

```go
func (builder *DashboardBuilder) Elements(elements map[string]dashboardv2beta1.Element) *DashboardBuilder
```

### <span class="badge object-method"></span> GridLayout

```go
func (builder *DashboardBuilder) GridLayout(gridLayoutKind cog.Builder[dashboardv2beta1.GridLayoutKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> GroupByVariable

```go
func (builder *DashboardBuilder) GroupByVariable(groupByVariableKind cog.Builder[dashboardv2beta1.GroupByVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> IntervalVariable

```go
func (builder *DashboardBuilder) IntervalVariable(intervalVariableKind cog.Builder[dashboardv2beta1.IntervalVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> LibraryPanel

```go
func (builder *DashboardBuilder) LibraryPanel(key string, libraryPanelKind cog.Builder[dashboardv2beta1.LibraryPanelKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> Links

Links with references to other dashboards or external websites.

```go
func (builder *DashboardBuilder) Links(links []cog.Builder[dashboardv2beta1.DashboardLink]) *DashboardBuilder
```

### <span class="badge object-method"></span> LiveNow

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data.

```go
func (builder *DashboardBuilder) LiveNow(liveNow bool) *DashboardBuilder
```

### <span class="badge object-method"></span> Panel

```go
func (builder *DashboardBuilder) Panel(key string, panelKind cog.Builder[dashboardv2beta1.PanelKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> Preload

When set to true, the dashboard will load all panels in the dashboard when it's loaded.

```go
func (builder *DashboardBuilder) Preload(preload bool) *DashboardBuilder
```

### <span class="badge object-method"></span> QueryVariable

```go
func (builder *DashboardBuilder) QueryVariable(queryVariableKind cog.Builder[dashboardv2beta1.QueryVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> Revision

Plugins only. The version of the dashboard installed together with the plugin.

This is used to determine if the dashboard should be updated when the plugin is updated.

```go
func (builder *DashboardBuilder) Revision(revision uint16) *DashboardBuilder
```

### <span class="badge object-method"></span> RowsLayout

```go
func (builder *DashboardBuilder) RowsLayout(rowsLayoutKind cog.Builder[dashboardv2beta1.RowsLayoutKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> TabsLayout

```go
func (builder *DashboardBuilder) TabsLayout(tabsLayoutKind cog.Builder[dashboardv2beta1.TabsLayoutKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> Tags

Tags associated with dashboard.

```go
func (builder *DashboardBuilder) Tags(tags []string) *DashboardBuilder
```

### <span class="badge object-method"></span> TextVariable

```go
func (builder *DashboardBuilder) TextVariable(textVariableKind cog.Builder[dashboardv2beta1.TextVariableKind]) *DashboardBuilder
```

### <span class="badge object-method"></span> TimeSettings

```go
func (builder *DashboardBuilder) TimeSettings(timeSettings cog.Builder[dashboardv2beta1.TimeSettingsSpec]) *DashboardBuilder
```

### <span class="badge object-method"></span> Title

Title of dashboard.

```go
func (builder *DashboardBuilder) Title(title string) *DashboardBuilder
```

### <span class="badge object-method"></span> Variables

Configured template variables.

```go
func (builder *DashboardBuilder) Variables(variables []dashboardv2beta1.VariableKind) *DashboardBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Dashboard](./object-Dashboard.md)

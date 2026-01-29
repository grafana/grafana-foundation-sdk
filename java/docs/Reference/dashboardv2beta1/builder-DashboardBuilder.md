---
title: <span class="badge builder"></span> DashboardBuilder
---
# <span class="badge builder"></span> DashboardBuilder

## Constructor

```java
new DashboardBuilder(String title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public Dashboard build()
```

### <span class="badge object-method"></span> switchVariableKind

```java
public DashboardBuilder switchVariableKind(com.grafana.foundation.cog.Builder<SwitchVariableKind> switchVariableKind)
```

### <span class="badge object-method"></span> adhocVariable

```java
public DashboardBuilder adhocVariable(com.grafana.foundation.cog.Builder<AdhocVariableKind> adhocVariableKind)
```

### <span class="badge object-method"></span> annotations

```java
public DashboardBuilder annotations(List<com.grafana.foundation.cog.Builder<AnnotationQueryKind>> annotations)
```

### <span class="badge object-method"></span> autoGridLayout

```java
public DashboardBuilder autoGridLayout(com.grafana.foundation.cog.Builder<AutoGridLayoutKind> autoGridLayoutKind)
```

### <span class="badge object-method"></span> constantVariable

```java
public DashboardBuilder constantVariable(com.grafana.foundation.cog.Builder<ConstantVariableKind> constantVariableKind)
```

### <span class="badge object-method"></span> cursorSync

Configuration of dashboard cursor sync behavior.

"Off" for no shared crosshair or tooltip (default).

"Crosshair" for shared crosshair.

"Tooltip" for shared crosshair AND shared tooltip.

```java
public DashboardBuilder cursorSync(DashboardCursorSync cursorSync)
```

### <span class="badge object-method"></span> customVariable

```java
public DashboardBuilder customVariable(com.grafana.foundation.cog.Builder<CustomVariableKind> customVariableKind)
```

### <span class="badge object-method"></span> datasourceVariable

```java
public DashboardBuilder datasourceVariable(com.grafana.foundation.cog.Builder<DatasourceVariableKind> datasourceVariableKind)
```

### <span class="badge object-method"></span> description

Description of dashboard.

```java
public DashboardBuilder description(String description)
```

### <span class="badge object-method"></span> editable

Whether a dashboard is editable or not.

```java
public DashboardBuilder editable(Boolean editable)
```

### <span class="badge object-method"></span> elements

```java
public DashboardBuilder elements(Map<String, Element> elements)
```

### <span class="badge object-method"></span> gridLayout

```java
public DashboardBuilder gridLayout(com.grafana.foundation.cog.Builder<GridLayoutKind> gridLayoutKind)
```

### <span class="badge object-method"></span> groupByVariable

```java
public DashboardBuilder groupByVariable(com.grafana.foundation.cog.Builder<GroupByVariableKind> groupByVariableKind)
```

### <span class="badge object-method"></span> intervalVariable

```java
public DashboardBuilder intervalVariable(com.grafana.foundation.cog.Builder<IntervalVariableKind> intervalVariableKind)
```

### <span class="badge object-method"></span> libraryPanel

```java
public DashboardBuilder libraryPanel(String key, com.grafana.foundation.cog.Builder<LibraryPanelKind> libraryPanelKind)
```

### <span class="badge object-method"></span> links

Links with references to other dashboards or external websites.

```java
public DashboardBuilder links(List<com.grafana.foundation.cog.Builder<DashboardLink>> links)
```

### <span class="badge object-method"></span> liveNow

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data.

```java
public DashboardBuilder liveNow(Boolean liveNow)
```

### <span class="badge object-method"></span> panel

```java
public DashboardBuilder panel(String key, com.grafana.foundation.cog.Builder<PanelKind> panelKind)
```

### <span class="badge object-method"></span> preload

When set to true, the dashboard will load all panels in the dashboard when it's loaded.

```java
public DashboardBuilder preload(Boolean preload)
```

### <span class="badge object-method"></span> queryVariable

```java
public DashboardBuilder queryVariable(com.grafana.foundation.cog.Builder<QueryVariableKind> queryVariableKind)
```

### <span class="badge object-method"></span> revision

Plugins only. The version of the dashboard installed together with the plugin.

This is used to determine if the dashboard should be updated when the plugin is updated.

```java
public DashboardBuilder revision(Short revision)
```

### <span class="badge object-method"></span> rowsLayout

```java
public DashboardBuilder rowsLayout(com.grafana.foundation.cog.Builder<RowsLayoutKind> rowsLayoutKind)
```

### <span class="badge object-method"></span> tabsLayout

```java
public DashboardBuilder tabsLayout(com.grafana.foundation.cog.Builder<TabsLayoutKind> tabsLayoutKind)
```

### <span class="badge object-method"></span> tags

Tags associated with dashboard.

```java
public DashboardBuilder tags(List<String> tags)
```

### <span class="badge object-method"></span> textVariable

```java
public DashboardBuilder textVariable(com.grafana.foundation.cog.Builder<TextVariableKind> textVariableKind)
```

### <span class="badge object-method"></span> timeSettings

```java
public DashboardBuilder timeSettings(com.grafana.foundation.cog.Builder<TimeSettingsSpec> timeSettings)
```

### <span class="badge object-method"></span> title

Title of dashboard.

```java
public DashboardBuilder title(String title)
```

### <span class="badge object-method"></span> variables

Configured template variables.

```java
public DashboardBuilder variables(List<VariableKind> variables)
```

## See also

 * <span class="badge object-type-class"></span> [Dashboard](./object-Dashboard.md)

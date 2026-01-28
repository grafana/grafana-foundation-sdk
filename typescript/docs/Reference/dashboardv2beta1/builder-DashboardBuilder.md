---
title: <span class="badge builder"></span> DashboardBuilder
---
# <span class="badge builder"></span> DashboardBuilder

## Constructor

```typescript
new DashboardBuilder(title: string)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> annotations

```typescript
annotations(annotations: cog.Builder<dashboardv2beta1.AnnotationQueryKind>[])
```

### <span class="badge object-method"></span> autoGridLayout

```typescript
autoGridLayout(autoGridLayoutKind: cog.Builder<dashboardv2beta1.AutoGridLayoutKind>)
```

### <span class="badge object-method"></span> cursorSync

Configuration of dashboard cursor sync behavior.

"Off" for no shared crosshair or tooltip (default).

"Crosshair" for shared crosshair.

"Tooltip" for shared crosshair AND shared tooltip.

```typescript
cursorSync(cursorSync: dashboardv2beta1.DashboardCursorSync)
```

### <span class="badge object-method"></span> description

Description of dashboard.

```typescript
description(description: string)
```

### <span class="badge object-method"></span> editable

Whether a dashboard is editable or not.

```typescript
editable(editable: boolean)
```

### <span class="badge object-method"></span> element

```typescript
element(key: string, element: cog.Builder<dashboardv2beta1.Element>)
```

### <span class="badge object-method"></span> elements

```typescript
elements(elements: Record<string, cog.Builder<dashboardv2beta1.Element>>)
```

### <span class="badge object-method"></span> gridLayout

```typescript
gridLayout(gridLayoutKind: cog.Builder<dashboardv2beta1.GridLayoutKind>)
```

### <span class="badge object-method"></span> links

Links with references to other dashboards or external websites.

```typescript
links(links: cog.Builder<dashboardv2beta1.DashboardLink>[])
```

### <span class="badge object-method"></span> liveNow

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data.

```typescript
liveNow(liveNow: boolean)
```

### <span class="badge object-method"></span> preload

When set to true, the dashboard will load all panels in the dashboard when it's loaded.

```typescript
preload(preload: boolean)
```

### <span class="badge object-method"></span> revision

Plugins only. The version of the dashboard installed together with the plugin.

This is used to determine if the dashboard should be updated when the plugin is updated.

```typescript
revision(revision: number)
```

### <span class="badge object-method"></span> rowsLayout

```typescript
rowsLayout(rowsLayoutKind: cog.Builder<dashboardv2beta1.RowsLayoutKind>)
```

### <span class="badge object-method"></span> tabsLayout

```typescript
tabsLayout(tabsLayoutKind: cog.Builder<dashboardv2beta1.TabsLayoutKind>)
```

### <span class="badge object-method"></span> tags

Tags associated with dashboard.

```typescript
tags(tags: string[])
```

### <span class="badge object-method"></span> timeSettings

```typescript
timeSettings(timeSettings: cog.Builder<dashboardv2beta1.TimeSettingsSpec>)
```

### <span class="badge object-method"></span> title

Title of dashboard.

```typescript
title(title: string)
```

### <span class="badge object-method"></span> variable

Configured template variables.

```typescript
variable(variable: cog.Builder<dashboardv2beta1.VariableKind>)
```

### <span class="badge object-method"></span> variables

Configured template variables.

```typescript
variables(variables: cog.Builder<dashboardv2beta1.VariableKind>[])
```

## See also

 * <span class="badge object-type-interface"></span> [Dashboard](./object-Dashboard.md)

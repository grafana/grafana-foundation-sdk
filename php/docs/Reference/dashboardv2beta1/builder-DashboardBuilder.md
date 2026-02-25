---
title: <span class="badge builder"></span> DashboardBuilder
---
# <span class="badge builder"></span> DashboardBuilder

## Constructor

```php
new DashboardBuilder(string $title)
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> annotations

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AnnotationQueryKind>> $annotations

```php
annotations(array $annotations)
```

### <span class="badge object-method"></span> cursorSync

Configuration of dashboard cursor sync behavior.

"Off" for no shared crosshair or tooltip (default).

"Crosshair" for shared crosshair.

"Tooltip" for shared crosshair AND shared tooltip.

```php
cursorSync(\Grafana\Foundation\Dashboardv2beta1\DashboardCursorSync $cursorSync)
```

### <span class="badge object-method"></span> description

Description of dashboard.

```php
description(string $description)
```

### <span class="badge object-method"></span> editable

Whether a dashboard is editable or not.

```php
editable(bool $editable)
```

### <span class="badge object-method"></span> element

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind> $element

```php
element(string $key, $element)
```

### <span class="badge object-method"></span> elements

@param array<string, \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\PanelKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\LibraryPanelKind>> $elements

```php
elements(array $elements)
```

### <span class="badge object-method"></span> layout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind> $layout

```php
layout($layout)
```

### <span class="badge object-method"></span> links

Links with references to other dashboards or external websites.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DashboardLink>> $links

```php
links(array $links)
```

### <span class="badge object-method"></span> liveNow

When set to true, the dashboard will redraw panels at an interval matching the pixel width.

This will keep data "moving left" regardless of the query refresh rate. This setting helps

avoid dashboards presenting stale live data.

```php
liveNow(bool $liveNow)
```

### <span class="badge object-method"></span> preload

When set to true, the dashboard will load all panels in the dashboard when it's loaded.

```php
preload(bool $preload)
```

### <span class="badge object-method"></span> revision

Plugins only. The version of the dashboard installed together with the plugin.

This is used to determine if the dashboard should be updated when the plugin is updated.

```php
revision(int $revision)
```

### <span class="badge object-method"></span> tags

Tags associated with dashboard.

@param array<string> $tags

```php
tags(array $tags)
```

### <span class="badge object-method"></span> timeSettings

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TimeSettingsSpec> $timeSettings

```php
timeSettings(\Grafana\Foundation\Cog\Builder $timeSettings)
```

### <span class="badge object-method"></span> title

Title of dashboard.

```php
title(string $title)
```

### <span class="badge object-method"></span> variable

Configured template variables.

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TextVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind> $variable

```php
variable($variable)
```

### <span class="badge object-method"></span> variables

Configured template variables.

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TextVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind>> $variables

```php
variables(array $variables)
```

## See also

 * <span class="badge object-type-class"></span> [Dashboard](./object-Dashboard.md)

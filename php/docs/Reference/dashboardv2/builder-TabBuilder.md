---
title: <span class="badge builder"></span> TabBuilder
---
# <span class="badge builder"></span> TabBuilder

## Constructor

```php
new TabBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> conditionalRendering

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConditionalRenderingGroupKind> $conditionalRendering

```php
conditionalRendering(\Grafana\Foundation\Cog\Builder $conditionalRendering)
```

### <span class="badge object-method"></span> layout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\RowsLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\TabsLayoutKind> $layout

```php
layout($layout)
```

### <span class="badge object-method"></span> repeat

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\TabRepeatOptions> $repeat

```php
repeat(\Grafana\Foundation\Cog\Builder $repeat)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

### <span class="badge object-method"></span> variables

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\QueryVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\TextVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\ConstantVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\DatasourceVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\IntervalVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\CustomVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\GroupByVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\AdhocVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2\SwitchVariableKind>> $variables

```php
variables(array $variables)
```

## See also

 * <span class="badge object-type-class"></span> [TabsLayoutTabKind](./object-TabsLayoutTabKind.md)

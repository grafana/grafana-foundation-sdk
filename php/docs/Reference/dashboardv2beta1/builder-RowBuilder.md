---
title: <span class="badge builder"></span> RowBuilder
---
# <span class="badge builder"></span> RowBuilder

## Constructor

```php
new RowBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> collapse

```php
collapse(bool $collapse)
```

### <span class="badge object-method"></span> conditionalRendering

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind> $conditionalRendering

```php
conditionalRendering(\Grafana\Foundation\Cog\Builder $conditionalRendering)
```

### <span class="badge object-method"></span> fillScreen

```php
fillScreen(bool $fillScreen)
```

### <span class="badge object-method"></span> hideHeader

```php
hideHeader(bool $hideHeader)
```

### <span class="badge object-method"></span> layout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind> $layout

```php
layout($layout)
```

### <span class="badge object-method"></span> repeat

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowRepeatOptions> $repeat

```php
repeat(\Grafana\Foundation\Cog\Builder $repeat)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

### <span class="badge object-method"></span> variables

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\QueryVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TextVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConstantVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\DatasourceVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\IntervalVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\CustomVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GroupByVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AdhocVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\SwitchVariableKind>> $variables

```php
variables(array $variables)
```

## See also

 * <span class="badge object-type-class"></span> [RowsLayoutRowKind](./object-RowsLayoutRowKind.md)

---
title: <span class="badge builder"></span> TabsLayoutTabSpecBuilder
---
# <span class="badge builder"></span> TabsLayoutTabSpecBuilder

## Constructor

```php
new TabsLayoutTabSpecBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> conditionalRendering

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind> $conditionalRendering

```php
conditionalRendering(\Grafana\Foundation\Cog\Builder $conditionalRendering)
```

### <span class="badge object-method"></span> layout

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\GridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind> $layout

```php
layout($layout)
```

### <span class="badge object-method"></span> repeat

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\TabRepeatOptions> $repeat

```php
repeat(\Grafana\Foundation\Cog\Builder $repeat)
```

### <span class="badge object-method"></span> title

```php
title(string $title)
```

## See also

 * <span class="badge object-type-class"></span> [TabsLayoutTabSpec](./object-TabsLayoutTabSpec.md)

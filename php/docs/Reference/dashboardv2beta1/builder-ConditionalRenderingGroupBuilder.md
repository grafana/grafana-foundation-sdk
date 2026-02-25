---
title: <span class="badge builder"></span> ConditionalRenderingGroupBuilder
---
# <span class="badge builder"></span> ConditionalRenderingGroupBuilder

## Constructor

```php
new ConditionalRenderingGroupBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> condition

```php
condition(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition $condition)
```

### <span class="badge object-method"></span> item

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind> $item

```php
item($item)
```

### <span class="badge object-method"></span> items

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind>> $items

```php
items(array $items)
```

### <span class="badge object-method"></span> visibility

```php
visibility(\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility $visibility)
```

## See also

 * <span class="badge object-type-class"></span> [ConditionalRenderingGroupKind](./object-ConditionalRenderingGroupKind.md)

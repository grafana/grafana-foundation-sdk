---
title: <span class="badge builder"></span> ConditionalRenderingGroupSpecBuilder
---
# <span class="badge builder"></span> ConditionalRenderingGroupSpecBuilder

## Constructor

```php
new ConditionalRenderingGroupSpecBuilder()
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

 * <span class="badge object-type-class"></span> [ConditionalRenderingGroupSpec](./object-ConditionalRenderingGroupSpec.md)

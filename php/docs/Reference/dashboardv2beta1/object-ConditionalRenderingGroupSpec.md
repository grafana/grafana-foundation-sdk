---
title: <span class="badge object-type-class"></span> ConditionalRenderingGroupSpec
---
# <span class="badge object-type-class"></span> ConditionalRenderingGroupSpec

## Definition

```php
class ConditionalRenderingGroupSpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecVisibility $visibility;

    public \Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupSpecCondition $condition;

    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingVariableKind|\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingDataKind|\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingTimeRangeSizeKind>
     */
    public array $items;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [ConditionalRenderingGroupSpecBuilder](./builder-ConditionalRenderingGroupSpecBuilder.md)

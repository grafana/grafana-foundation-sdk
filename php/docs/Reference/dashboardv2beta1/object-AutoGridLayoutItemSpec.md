---
title: <span class="badge object-type-class"></span> AutoGridLayoutItemSpec
---
# <span class="badge object-type-class"></span> AutoGridLayoutItemSpec

## Definition

```php
class AutoGridLayoutItemSpec implements \JsonSerializable
{
    public \Grafana\Foundation\Dashboardv2beta1\ElementReference $element;

    public ?\Grafana\Foundation\Dashboardv2beta1\AutoGridRepeatOptions $repeat;

    public ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering;

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

 * <span class="badge builder"></span> [AutoGridLayoutItemSpecBuilder](./builder-AutoGridLayoutItemSpecBuilder.md)

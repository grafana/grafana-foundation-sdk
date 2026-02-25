---
title: <span class="badge object-type-class"></span> TabsLayoutTabSpec
---
# <span class="badge object-type-class"></span> TabsLayoutTabSpec

## Definition

```php
class TabsLayoutTabSpec implements \JsonSerializable
{
    public ?string $title;

    /**
     * @var \Grafana\Foundation\Dashboardv2beta1\GridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\RowsLayoutKind|\Grafana\Foundation\Dashboardv2beta1\AutoGridLayoutKind|\Grafana\Foundation\Dashboardv2beta1\TabsLayoutKind
     */
    public $layout;

    public ?\Grafana\Foundation\Dashboardv2beta1\ConditionalRenderingGroupKind $conditionalRendering;

    public ?\Grafana\Foundation\Dashboardv2beta1\TabRepeatOptions $repeat;

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


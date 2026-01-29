---
title: <span class="badge object-type-class"></span> GridLayoutItemSpec
---
# <span class="badge object-type-class"></span> GridLayoutItemSpec

## Definition

```php
class GridLayoutItemSpec implements \JsonSerializable
{
    public int $x;

    public int $y;

    public int $width;

    public int $height;

    /**
     * reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
     */
    public \Grafana\Foundation\Dashboardv2beta1\ElementReference $element;

    public ?\Grafana\Foundation\Dashboardv2beta1\RepeatOptions $repeat;

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


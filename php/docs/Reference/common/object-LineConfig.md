---
title: <span class="badge object-type-class"></span> LineConfig
---
# <span class="badge object-type-class"></span> LineConfig

TODO docs

## Definition

```php
class LineConfig implements \JsonSerializable
{
    public ?string $lineColor;

    public ?float $lineWidth;

    public ?\Grafana\Foundation\Common\LineInterpolation $lineInterpolation;

    public ?\Grafana\Foundation\Common\LineStyle $lineStyle;

    /**
     * Indicate if null values should be treated as gaps or connected.
     * When the value is a number, it represents the maximum delta in the
     * X axis that should be considered connected.  For timeseries, this is milliseconds
     * @var bool|float|null
     */
    public $spanNulls;

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

 * <span class="badge builder"></span> [LineConfigBuilder](./builder-LineConfigBuilder.md)

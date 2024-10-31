---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

## Definition

```php
class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Common\BarGaugeDisplayMode $displayMode;

    public \Grafana\Foundation\Common\BarGaugeValueMode $valueMode;

    public bool $showUnfilled;

    public int $minVizWidth;

    public \Grafana\Foundation\Common\ReduceDataOptions $reduceOptions;

    public ?\Grafana\Foundation\Common\VizTextDisplayOptions $text;

    public int $minVizHeight;

    public \Grafana\Foundation\Common\VizOrientation $orientation;

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


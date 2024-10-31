---
title: <span class="badge object-type-class"></span> Options
---
# <span class="badge object-type-class"></span> Options

Identical to timeseries... except it does not have timezone settings

## Definition

```php
class Options implements \JsonSerializable
{
    public \Grafana\Foundation\Common\VizLegendOptions $legend;

    public \Grafana\Foundation\Common\VizTooltipOptions $tooltip;

    /**
     * Name of the x field to use (defaults to first number)
     */
    public ?string $xField;

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


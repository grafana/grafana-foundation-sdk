---
title: <span class="badge object-type-class"></span> XYSeriesConfig
---
# <span class="badge object-type-class"></span> XYSeriesConfig

## Definition

```php
class XYSeriesConfig implements \JsonSerializable
{
    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigName $name;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigFrame $frame;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigX $x;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigY $y;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigColor $color;

    public ?\Grafana\Foundation\Xychart\XychartXYSeriesConfigSize $size;

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


---
title: <span class="badge object-type-class"></span> VizTooltipOptions
---
# <span class="badge object-type-class"></span> VizTooltipOptions

TODO docs

## Definition

```php
class VizTooltipOptions implements \JsonSerializable
{
    public \Grafana\Foundation\Common\TooltipDisplayMode $mode;

    public \Grafana\Foundation\Common\SortOrder $sort;

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

 * <span class="badge builder"></span> [VizTooltipOptionsBuilder](./builder-VizTooltipOptionsBuilder.md)

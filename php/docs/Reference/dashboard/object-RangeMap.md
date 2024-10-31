---
title: <span class="badge object-type-class"></span> RangeMap
---
# <span class="badge object-type-class"></span> RangeMap

Maps numerical ranges to a display text and color.

For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.

## Definition

```php
class RangeMap implements \JsonSerializable
{
    public string $type;

    /**
     * Range to match against and the result to apply when the value is within the range
     */
    public \Grafana\Foundation\Dashboard\DashboardRangeMapOptions $options;

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

 * <span class="badge builder"></span> [RangeMapBuilder](./builder-RangeMapBuilder.md)

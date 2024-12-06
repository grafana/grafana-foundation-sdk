---
title: <span class="badge object-type-class"></span> DashboardRangeMapOptions
---
# <span class="badge object-type-class"></span> DashboardRangeMapOptions

## Definition

```php
class DashboardRangeMapOptions implements \JsonSerializable
{
    /**
     * Min value of the range. It can be null which means -Infinity
     */
    public ?float $from;

    /**
     * Max value of the range. It can be null which means +Infinity
     */
    public ?float $to;

    /**
     * Config to apply when the value is within the range
     */
    public \Grafana\Foundation\Dashboard\ValueMappingResult $result;

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

 * <span class="badge builder"></span> [DashboardRangeMapOptionsBuilder](./builder-DashboardRangeMapOptionsBuilder.md)

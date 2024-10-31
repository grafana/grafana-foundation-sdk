---
title: <span class="badge object-type-class"></span> TimeRange
---
# <span class="badge object-type-class"></span> TimeRange

## Definition

```php
class TimeRange implements \JsonSerializable
{
    /**
     * From is the start time of the query.
     */
    public string $from;

    /**
     * To is the end time of the query.
     */
    public string $to;

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

 * <span class="badge builder"></span> [TimeRangeBuilder](./builder-TimeRangeBuilder.md)

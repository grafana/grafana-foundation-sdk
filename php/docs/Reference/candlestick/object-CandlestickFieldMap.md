---
title: <span class="badge object-type-class"></span> CandlestickFieldMap
---
# <span class="badge object-type-class"></span> CandlestickFieldMap

## Definition

```php
class CandlestickFieldMap implements \JsonSerializable
{
    /**
     * Corresponds to the starting value of the given period
     */
    public ?string $open;

    /**
     * Corresponds to the highest value of the given period
     */
    public ?string $high;

    /**
     * Corresponds to the lowest value of the given period
     */
    public ?string $low;

    /**
     * Corresponds to the final (end) value of the given period
     */
    public ?string $close;

    /**
     * Corresponds to the sample count in the given period. (e.g. number of trades)
     */
    public ?string $volume;

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

 * <span class="badge builder"></span> [CandlestickFieldMapBuilder](./builder-CandlestickFieldMapBuilder.md)

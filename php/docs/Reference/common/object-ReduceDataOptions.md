---
title: <span class="badge object-type-class"></span> ReduceDataOptions
---
# <span class="badge object-type-class"></span> ReduceDataOptions

TODO docs

## Definition

```php
class ReduceDataOptions implements \JsonSerializable
{
    /**
     * If true show each row value
     */
    public ?bool $values;

    /**
     * if showing all values limit
     */
    public ?float $limit;

    /**
     * When !values, pick one value for the whole field
     * @var array<string>
     */
    public array $calcs;

    /**
     * Which fields to show.  By default this is only numeric fields
     */
    public ?string $fields;

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

 * <span class="badge builder"></span> [ReduceDataOptionsBuilder](./builder-ReduceDataOptionsBuilder.md)

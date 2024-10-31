---
title: <span class="badge object-type-class"></span> EdgeOptions
---
# <span class="badge object-type-class"></span> EdgeOptions

## Definition

```php
class EdgeOptions implements \JsonSerializable
{
    /**
     * Unit for the main stat to override what ever is set in the data frame.
     */
    public ?string $mainStatUnit;

    /**
     * Unit for the secondary stat to override what ever is set in the data frame.
     */
    public ?string $secondaryStatUnit;

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

 * <span class="badge builder"></span> [EdgeOptionsBuilder](./builder-EdgeOptionsBuilder.md)

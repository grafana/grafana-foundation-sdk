---
title: <span class="badge object-type-class"></span> ArcOption
---
# <span class="badge object-type-class"></span> ArcOption

## Definition

```php
class ArcOption implements \JsonSerializable
{
    /**
     * Field from which to get the value. Values should be less than 1, representing fraction of a circle.
     */
    public ?string $field;

    /**
     * The color of the arc.
     */
    public ?string $color;

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

 * <span class="badge builder"></span> [ArcOptionBuilder](./builder-ArcOptionBuilder.md)

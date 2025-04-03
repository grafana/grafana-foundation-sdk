---
title: <span class="badge object-type-class"></span> Placement
---
# <span class="badge object-type-class"></span> Placement

## Definition

```php
class Placement implements \JsonSerializable
{
    public ?float $top;

    public ?float $left;

    public ?float $right;

    public ?float $bottom;

    public ?float $width;

    public ?float $height;

    public ?float $rotation;

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

 * <span class="badge builder"></span> [PlacementBuilder](./builder-PlacementBuilder.md)

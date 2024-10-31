---
title: <span class="badge object-type-class"></span> Key
---
# <span class="badge object-type-class"></span> Key

## Definition

```php
class Key implements \JsonSerializable
{
    public float $tick;

    public string $type;

    public ?string $uid;

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

 * <span class="badge builder"></span> [KeyBuilder](./builder-KeyBuilder.md)

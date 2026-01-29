---
title: <span class="badge object-type-class"></span> Threshold
---
# <span class="badge object-type-class"></span> Threshold

## Definition

```php
class Threshold implements \JsonSerializable
{
    /**
     * Value null means -Infinity
     */
    public ?float $value;

    public string $color;

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


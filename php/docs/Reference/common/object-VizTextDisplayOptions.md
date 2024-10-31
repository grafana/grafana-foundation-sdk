---
title: <span class="badge object-type-class"></span> VizTextDisplayOptions
---
# <span class="badge object-type-class"></span> VizTextDisplayOptions

TODO docs

## Definition

```php
class VizTextDisplayOptions implements \JsonSerializable
{
    /**
     * Explicit title text size
     */
    public ?float $titleSize;

    /**
     * Explicit value text size
     */
    public ?float $valueSize;

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

 * <span class="badge builder"></span> [VizTextDisplayOptionsBuilder](./builder-VizTextDisplayOptionsBuilder.md)

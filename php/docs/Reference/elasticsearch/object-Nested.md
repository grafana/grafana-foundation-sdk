---
title: <span class="badge object-type-class"></span> Nested
---
# <span class="badge object-type-class"></span> Nested

## Definition

```php
class Nested implements \JsonSerializable
{
    public ?string $field;

    public string $id;

    public string $type;

    /**
     * @var mixed|null
     */
    public $settings;

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

 * <span class="badge builder"></span> [NestedBuilder](./builder-NestedBuilder.md)

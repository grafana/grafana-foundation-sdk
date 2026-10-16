---
title: <span class="badge object-type-class"></span> CodeCellContentSpec
---
# <span class="badge object-type-class"></span> CodeCellContentSpec

## Definition

```php
class CodeCellContentSpec implements \JsonSerializable
{
    public string $language;

    public string $code;

    /**
     * @var array<int>|null
     */
    public ?array $highlight;

    public ?string $annotation;

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

 * <span class="badge builder"></span> [CodeCellContentSpecBuilder](./builder-CodeCellContentSpecBuilder.md)

---
title: <span class="badge object-type-class"></span> ValueMappingResult
---
# <span class="badge object-type-class"></span> ValueMappingResult

Result used as replacement with text and color when the value matches

## Definition

```php
class ValueMappingResult implements \JsonSerializable
{
    /**
     * Text to display when the value matches
     */
    public ?string $text;

    /**
     * Text to use when the value matches
     */
    public ?string $color;

    /**
     * Icon to display when the value matches. Only specific visualizations.
     */
    public ?string $icon;

    /**
     * Position in the mapping array. Only used internally.
     */
    public ?int $index;

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

 * <span class="badge builder"></span> [ValueMappingResultBuilder](./builder-ValueMappingResultBuilder.md)

---
title: <span class="badge object-type-class"></span> ValueMap
---
# <span class="badge object-type-class"></span> ValueMap

Maps text values to a color or different display text and color.

For example, you can configure a value mapping so that all instances of the value 10 appear as Perfection! rather than the number.

## Definition

```php
class ValueMap implements \JsonSerializable
{
    public string $type;

    /**
     * Map with <value_to_match>: ValueMappingResult. For example: { "10": { text: "Perfection!", color: "green" } }
     * @var array<string, \Grafana\Foundation\Dashboard\ValueMappingResult>
     */
    public array $options;

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

 * <span class="badge builder"></span> [ValueMapBuilder](./builder-ValueMapBuilder.md)

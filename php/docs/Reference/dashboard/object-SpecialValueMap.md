---
title: <span class="badge object-type-class"></span> SpecialValueMap
---
# <span class="badge object-type-class"></span> SpecialValueMap

Maps special values like Null, NaN (not a number), and boolean values like true and false to a display text and color.

See SpecialValueMatch to see the list of special values.

For example, you can configure a special value mapping so that null values appear as N/A.

## Definition

```php
class SpecialValueMap implements \JsonSerializable
{
    public string $type;

    public \Grafana\Foundation\Dashboard\DashboardSpecialValueMapOptions $options;

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

 * <span class="badge builder"></span> [SpecialValueMapBuilder](./builder-SpecialValueMapBuilder.md)

---
title: <span class="badge object-type-class"></span> TextVariableKind
---
# <span class="badge object-type-class"></span> TextVariableKind

Text variable kind

## Definition

```php
class TextVariableKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\TextVariableSpec $spec;

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

 * <span class="badge builder"></span> [TextVariableBuilder](./builder-TextVariableBuilder.md)

---
title: <span class="badge object-type-class"></span> V2QueryGroupKind
---
# <span class="badge object-type-class"></span> V2QueryGroupKind

## Definition

```php
class V2QueryGroupKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\V2QueryGroupSpec $spec;

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

 * <span class="badge builder"></span> [V2QueryGroupKindBuilder](./builder-V2QueryGroupKindBuilder.md)

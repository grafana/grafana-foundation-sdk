---
title: <span class="badge object-type-class"></span> RowsLayoutRowKind
---
# <span class="badge object-type-class"></span> RowsLayoutRowKind

## Definition

```php
class RowsLayoutRowKind implements \JsonSerializable
{
    public string $kind;

    public \Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowSpec $spec;

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

 * <span class="badge builder"></span> [RowBuilder](./builder-RowBuilder.md)

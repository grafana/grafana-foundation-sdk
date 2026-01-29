---
title: <span class="badge object-type-class"></span> RowsLayoutSpec
---
# <span class="badge object-type-class"></span> RowsLayoutSpec

## Definition

```php
class RowsLayoutSpec implements \JsonSerializable
{
    /**
     * @var array<\Grafana\Foundation\Dashboardv2beta1\RowsLayoutRowKind>
     */
    public array $rows;

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

 * <span class="badge builder"></span> [RowsLayoutSpecBuilder](./builder-RowsLayoutSpecBuilder.md)

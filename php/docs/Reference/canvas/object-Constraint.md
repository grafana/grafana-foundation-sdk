---
title: <span class="badge object-type-class"></span> Constraint
---
# <span class="badge object-type-class"></span> Constraint

## Definition

```php
class Constraint implements \JsonSerializable
{
    public ?\Grafana\Foundation\Canvas\HorizontalConstraint $horizontal;

    public ?\Grafana\Foundation\Canvas\VerticalConstraint $vertical;

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

 * <span class="badge builder"></span> [ConstraintBuilder](./builder-ConstraintBuilder.md)
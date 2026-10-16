---
title: <span class="badge object-type-class"></span> V2TransformationKind
---
# <span class="badge object-type-class"></span> V2TransformationKind

Dashboard v2 shape: the transformation ID moved from `kind` to `group`.

## Definition

```php
class V2TransformationKind implements \JsonSerializable
{
    public string $kind;

    /**
     * The group is the transformation ID
     */
    public string $group;

    public \Grafana\Foundation\Dashboardv2beta1\V2TransformationSpec $spec;

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

 * <span class="badge builder"></span> [V2TransformationKindBuilder](./builder-V2TransformationKindBuilder.md)

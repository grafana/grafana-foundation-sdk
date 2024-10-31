---
title: <span class="badge object-type-class"></span> RoleRef
---
# <span class="badge object-type-class"></span> RoleRef

## Definition

```php
class RoleRef implements \JsonSerializable
{
    /**
     * Policies can apply to roles, teams, or users
     * Applying policies to individual users is supported, but discouraged
     */
    public \Grafana\Foundation\Accesspolicy\RoleRefKind $kind;

    public string $name;

    public string $xname;

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

 * <span class="badge builder"></span> [RoleRefBuilder](./builder-RoleRefBuilder.md)

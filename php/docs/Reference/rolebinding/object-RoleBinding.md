---
title: <span class="badge object-type-class"></span> RoleBinding
---
# <span class="badge object-type-class"></span> RoleBinding

## Definition

```php
class RoleBinding implements \JsonSerializable
{
    /**
     * The role we are discussing
     * @var \Grafana\Foundation\Rolebinding\BuiltinRoleRef|\Grafana\Foundation\Rolebinding\CustomRoleRef
     */
    public $role;

    /**
     * The team or user that has the specified role
     */
    public \Grafana\Foundation\Rolebinding\RoleBindingSubject $subject;

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

 * <span class="badge builder"></span> [RoleBindingBuilder](./builder-RoleBindingBuilder.md)

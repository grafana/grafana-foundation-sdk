---
title: <span class="badge object-type-class"></span> AccessPolicy
---
# <span class="badge object-type-class"></span> AccessPolicy

## Definition

```php
class AccessPolicy implements \JsonSerializable
{
    /**
     * The scope where these policies should apply
     */
    public \Grafana\Foundation\Accesspolicy\ResourceRef $scope;

    /**
     * The role that must apply this policy
     */
    public \Grafana\Foundation\Accesspolicy\RoleRef $role;

    /**
     * The set of rules to apply.  Note that * is required to modify
     * access policy rules, and that "none" will reject all actions
     * @var array<\Grafana\Foundation\Accesspolicy\AccessRule>
     */
    public array $rules;

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

 * <span class="badge builder"></span> [AccessPolicyBuilder](./builder-AccessPolicyBuilder.md)

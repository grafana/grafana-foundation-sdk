---
title: <span class="badge builder"></span> RoleBindingBuilder
---
# <span class="badge builder"></span> RoleBindingBuilder

## Constructor

```php
new RoleBindingBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> role

The role we are discussing

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\BuiltinRoleRef>|\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\CustomRoleRef> $role

```php
role($role)
```

### <span class="badge object-method"></span> subject

The team or user that has the specified role

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Rolebinding\RoleBindingSubject> $subject

```php
subject(\Grafana\Foundation\Cog\Builder $subject)
```

## See also

 * <span class="badge object-type-class"></span> [RoleBinding](./object-RoleBinding.md)

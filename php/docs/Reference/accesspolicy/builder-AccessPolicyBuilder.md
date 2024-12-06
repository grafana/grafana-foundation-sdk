---
title: <span class="badge builder"></span> AccessPolicyBuilder
---
# <span class="badge builder"></span> AccessPolicyBuilder

## Constructor

```php
new AccessPolicyBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> role

The role that must apply this policy

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\RoleRef> $role

```php
role(\Grafana\Foundation\Cog\Builder $role)
```

### <span class="badge object-method"></span> rules

The set of rules to apply.  Note that * is required to modify

access policy rules, and that "none" will reject all actions

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\AccessRule> $rule

```php
rules(\Grafana\Foundation\Cog\Builder $rule)
```

### <span class="badge object-method"></span> scope

The scope where these policies should apply

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Accesspolicy\ResourceRef> $scope

```php
scope(\Grafana\Foundation\Cog\Builder $scope)
```

## See also

 * <span class="badge object-type-class"></span> [AccessPolicy](./object-AccessPolicy.md)

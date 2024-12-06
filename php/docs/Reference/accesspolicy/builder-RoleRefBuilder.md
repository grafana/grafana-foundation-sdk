---
title: <span class="badge builder"></span> RoleRefBuilder
---
# <span class="badge builder"></span> RoleRefBuilder

## Constructor

```php
new RoleRefBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> kind

Policies can apply to roles, teams, or users

Applying policies to individual users is supported, but discouraged

```php
kind(\Grafana\Foundation\Accesspolicy\RoleRefKind $kind)
```

### <span class="badge object-method"></span> name

```php
name(string $name)
```

### <span class="badge object-method"></span> xname

```php
xname(string $xname)
```

## See also

 * <span class="badge object-type-class"></span> [RoleRef](./object-RoleRef.md)

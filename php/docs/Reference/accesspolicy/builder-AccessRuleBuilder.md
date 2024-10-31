---
title: <span class="badge builder"></span> AccessRuleBuilder
---
# <span class="badge builder"></span> AccessRuleBuilder

## Constructor

```php
new AccessRuleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> kind

The kind this rule applies to (dashboards, alert, etc)

```php
kind(string $kind)
```

### <span class="badge object-method"></span> target

Specific sub-elements like "alert.rules" or "dashboard.permissions"????

```php
target(string $target)
```

### <span class="badge object-method"></span> verb

READ, WRITE, CREATE, DELETE, ...

should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"

@param string $verb

```php
verb($verb)
```

## See also

 * <span class="badge object-type-class"></span> [AccessRule](./object-AccessRule.md)

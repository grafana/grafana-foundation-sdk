---
title: <span class="badge builder"></span> RoleBindingBuilder
---
# <span class="badge builder"></span> RoleBindingBuilder

## Constructor

```java
new RoleBindingBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public RoleBinding build()
```

### <span class="badge object-method"></span> role

The role we are discussing

```java
public RoleBindingBuilder role(BuiltinRoleRefOrCustomRoleRef role)
```

### <span class="badge object-method"></span> subject

The team or user that has the specified role

```java
public RoleBindingBuilder subject(com.grafana.foundation.cog.Builder<RoleBindingSubject> subject)
```

## See also

 * <span class="badge object-type-class"></span> [RoleBinding](./object-RoleBinding.md)

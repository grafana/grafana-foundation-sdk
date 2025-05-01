---
title: <span class="badge builder"></span> RoleRefBuilder
---
# <span class="badge builder"></span> RoleRefBuilder

## Constructor

```java
new RoleRefBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public RoleRef build()
```

### <span class="badge object-method"></span> kind

Policies can apply to roles, teams, or users

Applying policies to individual users is supported, but discouraged

```java
public RoleRefBuilder kind(RoleRefKind kind)
```

### <span class="badge object-method"></span> name

```java
public RoleRefBuilder name(String name)
```

### <span class="badge object-method"></span> xname

```java
public RoleRefBuilder xname(String xname)
```

## See also

 * <span class="badge object-type-class"></span> [RoleRef](./object-RoleRef.md)

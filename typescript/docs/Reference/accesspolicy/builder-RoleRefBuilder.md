---
title: <span class="badge builder"></span> RoleRefBuilder
---
# <span class="badge builder"></span> RoleRefBuilder

## Constructor

```typescript
new RoleRefBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> kind

Policies can apply to roles, teams, or users

Applying policies to individual users is supported, but discouraged

```typescript
kind(kind: "Role" | "BuiltinRole" | "Team" | "User")
```

### <span class="badge object-method"></span> name

```typescript
name(name: string)
```

### <span class="badge object-method"></span> xname

```typescript
xname(xname: string)
```

## See also

 * <span class="badge object-type-interface"></span> [RoleRef](./object-RoleRef.md)

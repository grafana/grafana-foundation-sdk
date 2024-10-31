---
title: <span class="badge builder"></span> RoleBindingBuilder
---
# <span class="badge builder"></span> RoleBindingBuilder

## Constructor

```typescript
new RoleBindingBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> role

The role we are discussing

```typescript
role(role: cog.Builder<rolebinding.BuiltinRoleRef> | cog.Builder<rolebinding.CustomRoleRef>)
```

### <span class="badge object-method"></span> subject

The team or user that has the specified role

```typescript
subject(subject: cog.Builder<rolebinding.RoleBindingSubject>)
```

## See also

 * <span class="badge object-type-interface"></span> [RoleBinding](./object-RoleBinding.md)

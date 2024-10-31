---
title: <span class="badge builder"></span> AccessPolicyBuilder
---
# <span class="badge builder"></span> AccessPolicyBuilder

## Constructor

```typescript
new AccessPolicyBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> role

The role that must apply this policy

```typescript
role(role: cog.Builder<accesspolicy.RoleRef>)
```

### <span class="badge object-method"></span> rules

The set of rules to apply.  Note that * is required to modify

access policy rules, and that "none" will reject all actions

```typescript
rules(rule: cog.Builder<accesspolicy.AccessRule>)
```

### <span class="badge object-method"></span> scope

The scope where these policies should apply

```typescript
scope(scope: cog.Builder<accesspolicy.ResourceRef>)
```

## See also

 * <span class="badge object-type-interface"></span> [AccessPolicy](./object-AccessPolicy.md)

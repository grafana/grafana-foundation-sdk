---
title: <span class="badge builder"></span> AccessRuleBuilder
---
# <span class="badge builder"></span> AccessRuleBuilder

## Constructor

```typescript
new AccessRuleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> kind

The kind this rule applies to (dashboards, alert, etc)

```typescript
kind(kind: string)
```

### <span class="badge object-method"></span> target

Specific sub-elements like "alert.rules" or "dashboard.permissions"????

```typescript
target(target: string)
```

### <span class="badge object-method"></span> verb

READ, WRITE, CREATE, DELETE, ...

should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"

```typescript
verb(verb: "*" | "none" | string)
```

## See also

 * <span class="badge object-type-interface"></span> [AccessRule](./object-AccessRule.md)

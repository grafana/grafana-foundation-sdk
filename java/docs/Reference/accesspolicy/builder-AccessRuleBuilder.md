---
title: <span class="badge builder"></span> AccessRuleBuilder
---
# <span class="badge builder"></span> AccessRuleBuilder

## Constructor

```java
new AccessRuleBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AccessRule build()
```

### <span class="badge object-method"></span> kind

The kind this rule applies to (dashboards, alert, etc)

```java
public AccessRuleBuilder kind(String kind)
```

### <span class="badge object-method"></span> target

Specific sub-elements like "alert.rules" or "dashboard.permissions"????

```java
public AccessRuleBuilder target(String target)
```

### <span class="badge object-method"></span> verb

READ, WRITE, CREATE, DELETE, ...

should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"

```java
public AccessRuleBuilder verb(String verb)
```

## See also

 * <span class="badge object-type-class"></span> [AccessRule](./object-AccessRule.md)

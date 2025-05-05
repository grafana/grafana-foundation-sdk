---
title: <span class="badge builder"></span> AccessPolicyBuilder
---
# <span class="badge builder"></span> AccessPolicyBuilder

## Constructor

```java
new AccessPolicyBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public AccessPolicy build()
```

### <span class="badge object-method"></span> role

The role that must apply this policy

```java
public AccessPolicyBuilder role(com.grafana.foundation.cog.Builder<RoleRef> role)
```

### <span class="badge object-method"></span> rules

The set of rules to apply.  Note that * is required to modify

access policy rules, and that "none" will reject all actions

```java
public AccessPolicyBuilder rules(com.grafana.foundation.cog.Builder<AccessRule> rule)
```

### <span class="badge object-method"></span> scope

The scope where these policies should apply

```java
public AccessPolicyBuilder scope(com.grafana.foundation.cog.Builder<ResourceRef> scope)
```

## See also

 * <span class="badge object-type-class"></span> [AccessPolicy](./object-AccessPolicy.md)

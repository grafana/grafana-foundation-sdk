---
title: <span class="badge builder"></span> AccessRuleBuilder
---
# <span class="badge builder"></span> AccessRuleBuilder

## Constructor

```go
func NewAccessRuleBuilder() *AccessRuleBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AccessRuleBuilder) Build() (AccessRule, error)
```

### <span class="badge object-method"></span> Kind

The kind this rule applies to (dashboards, alert, etc)

```go
func (builder *AccessRuleBuilder) Kind(kind string) *AccessRuleBuilder
```

### <span class="badge object-method"></span> Target

Specific sub-elements like "alert.rules" or "dashboard.permissions"????

```go
func (builder *AccessRuleBuilder) Target(target string) *AccessRuleBuilder
```

### <span class="badge object-method"></span> Verb

READ, WRITE, CREATE, DELETE, ...

should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"

```go
func (builder *AccessRuleBuilder) Verb(verb string) *AccessRuleBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AccessRule](./object-AccessRule.md)

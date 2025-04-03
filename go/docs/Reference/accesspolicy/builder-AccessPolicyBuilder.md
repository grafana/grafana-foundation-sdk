---
title: <span class="badge builder"></span> AccessPolicyBuilder
---
# <span class="badge builder"></span> AccessPolicyBuilder

## Constructor

```go
func NewAccessPolicyBuilder() *AccessPolicyBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AccessPolicyBuilder) Build() (AccessPolicy, error)
```

### <span class="badge object-method"></span> Role

The role that must apply this policy

```go
func (builder *AccessPolicyBuilder) Role(role cog.Builder[accesspolicy.RoleRef]) *AccessPolicyBuilder
```

### <span class="badge object-method"></span> Rules

The set of rules to apply.  Note that * is required to modify

access policy rules, and that "none" will reject all actions

```go
func (builder *AccessPolicyBuilder) Rules(rule cog.Builder[accesspolicy.AccessRule]) *AccessPolicyBuilder
```

### <span class="badge object-method"></span> Scope

The scope where these policies should apply

```go
func (builder *AccessPolicyBuilder) Scope(scope cog.Builder[accesspolicy.ResourceRef]) *AccessPolicyBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AccessPolicy](./object-AccessPolicy.md)

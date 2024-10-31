# accesspolicy

## Objects

 * <span class="badge object-type-struct"></span> [AccessPolicy](./object-AccessPolicy.md)
 * <span class="badge object-type-struct"></span> [AccessRule](./object-AccessRule.md)
 * <span class="badge object-type-struct"></span> [ResourceRef](./object-ResourceRef.md)
 * <span class="badge object-type-struct"></span> [RoleRef](./object-RoleRef.md)
 * <span class="badge object-type-enum"></span> [RoleRefKind](./object-RoleRefKind.md)
## Builders

 * <span class="badge builder"></span> [AccessPolicyBuilder](./builder-AccessPolicyBuilder.md)
 * <span class="badge builder"></span> [AccessRuleBuilder](./builder-AccessRuleBuilder.md)
 * <span class="badge builder"></span> [ResourceRefBuilder](./builder-ResourceRefBuilder.md)
 * <span class="badge builder"></span> [RoleRefBuilder](./builder-RoleRefBuilder.md)
## Functions

### <span class="badge function"></span> AccessPolicyConverter

AccessPolicyConverter accepts a `AccessPolicy` object and generates the Go code to build this object using builders.

```go
func AccessPolicyConverter(input AccessPolicy) string
```

### <span class="badge function"></span> RoleRefConverter

RoleRefConverter accepts a `RoleRef` object and generates the Go code to build this object using builders.

```go
func RoleRefConverter(input RoleRef) string
```

### <span class="badge function"></span> ResourceRefConverter

ResourceRefConverter accepts a `ResourceRef` object and generates the Go code to build this object using builders.

```go
func ResourceRefConverter(input ResourceRef) string
```

### <span class="badge function"></span> AccessRuleConverter

AccessRuleConverter accepts a `AccessRule` object and generates the Go code to build this object using builders.

```go
func AccessRuleConverter(input AccessRule) string
```


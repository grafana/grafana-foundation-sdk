---
title: <span class="badge object-type-struct"></span> AccessPolicy
---
# <span class="badge object-type-struct"></span> AccessPolicy

## Definition

```go
type AccessPolicy struct {
    // The scope where these policies should apply
    Scope accesspolicy.ResourceRef `json:"scope"`
    // The role that must apply this policy
    Role accesspolicy.RoleRef `json:"role"`
    // The set of rules to apply.  Note that * is required to modify
    // access policy rules, and that "none" will reject all actions
    Rules []accesspolicy.AccessRule `json:"rules"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AccessPolicy` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (accessPolicy *AccessPolicy) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AccessPolicy` objects.

```go
func (accessPolicy *AccessPolicy) Equals(other AccessPolicy) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AccessPolicy` fields for violations and returns them.

```go
func (accessPolicy *AccessPolicy) Validate() error
```

## See also

 * <span class="badge builder"></span> [AccessPolicyBuilder](./builder-AccessPolicyBuilder.md)

---
title: <span class="badge object-type-struct"></span> AccessRule
---
# <span class="badge object-type-struct"></span> AccessRule

## Definition

```go
type AccessRule struct {
    // The kind this rule applies to (dashboards, alert, etc)
    Kind string `json:"kind"`
    // READ, WRITE, CREATE, DELETE, ...
    // should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
    Verb string `json:"verb"`
    // Specific sub-elements like "alert.rules" or "dashboard.permissions"????
    Target *string `json:"target,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `AccessRule` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (accessRule *AccessRule) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `AccessRule` objects.

```go
func (accessRule *AccessRule) Equals(other AccessRule) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `AccessRule` fields for violations and returns them.

```go
func (accessRule *AccessRule) Validate() error
```

## See also

 * <span class="badge builder"></span> [AccessRuleBuilder](./builder-AccessRuleBuilder.md)

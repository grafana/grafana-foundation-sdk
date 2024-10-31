---
title: <span class="badge object-type-interface"></span> AccessPolicy
---
# <span class="badge object-type-interface"></span> AccessPolicy

## Definition

```typescript
export interface AccessPolicy {
	// The scope where these policies should apply
	scope: accesspolicy.ResourceRef;
	// The role that must apply this policy
	role: accesspolicy.RoleRef;
	// The set of rules to apply.  Note that * is required to modify
	// access policy rules, and that "none" will reject all actions
	rules: accesspolicy.AccessRule[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AccessPolicyBuilder](./builder-AccessPolicyBuilder.md)

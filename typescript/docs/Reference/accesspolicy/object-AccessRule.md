---
title: <span class="badge object-type-interface"></span> AccessRule
---
# <span class="badge object-type-interface"></span> AccessRule

## Definition

```typescript
export interface AccessRule {
	// The kind this rule applies to (dashboards, alert, etc)
	kind: string;
	// READ, WRITE, CREATE, DELETE, ...
	// should move to k8s style verbs like: "get", "list", "watch", "create", "update", "patch", "delete"
	verb: "*" | "none" | string;
	// Specific sub-elements like "alert.rules" or "dashboard.permissions"????
	target?: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AccessRuleBuilder](./builder-AccessRuleBuilder.md)

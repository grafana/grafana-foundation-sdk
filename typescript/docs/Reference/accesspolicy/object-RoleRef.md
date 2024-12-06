---
title: <span class="badge object-type-interface"></span> RoleRef
---
# <span class="badge object-type-interface"></span> RoleRef

## Definition

```typescript
export interface RoleRef {
	// Policies can apply to roles, teams, or users
	// Applying policies to individual users is supported, but discouraged
	kind: "Role" | "BuiltinRole" | "Team" | "User";
	name: string;
	xname: string;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [RoleRefBuilder](./builder-RoleRefBuilder.md)

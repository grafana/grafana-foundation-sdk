---
title: <span class="badge object-type-interface"></span> Role
---
# <span class="badge object-type-interface"></span> Role

## Definition

```typescript
export interface Role {
	// The role identifier `managed:builtins:editor:permissions`
	name: string;
	// Optional display
	displayName?: string;
	// Name of the team.
	groupName?: string;
	// Role description
	description?: string;
	// Do not show this role
	hidden: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [RoleBuilder](./builder-RoleBuilder.md)

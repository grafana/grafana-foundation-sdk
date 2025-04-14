---
title: <span class="badge object-type-interface"></span> RoleBindingSubject
---
# <span class="badge object-type-interface"></span> RoleBindingSubject

## Definition

```typescript
export interface RoleBindingSubject {
	kind: "Team" | "User";
	// The team/user identifier name
	name: string;
}

```
## See also

 * <span class="badge builder"></span> [RoleBindingSubjectBuilder](./builder-RoleBindingSubjectBuilder.md)

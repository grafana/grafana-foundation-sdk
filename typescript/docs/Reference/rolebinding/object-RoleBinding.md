---
title: <span class="badge object-type-interface"></span> RoleBinding
---
# <span class="badge object-type-interface"></span> RoleBinding

## Definition

```typescript
export interface RoleBinding {
	// The role we are discussing
	role: rolebinding.BuiltinRoleRef | rolebinding.CustomRoleRef;
	// The team or user that has the specified role
	subject: rolebinding.RoleBindingSubject;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [RoleBindingBuilder](./builder-RoleBindingBuilder.md)

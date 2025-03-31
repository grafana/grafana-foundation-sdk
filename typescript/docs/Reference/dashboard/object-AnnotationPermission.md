---
title: <span class="badge object-type-interface"></span> AnnotationPermission
---
# <span class="badge object-type-interface"></span> AnnotationPermission

+k8s:deepcopy-gen=true

## Definition

```typescript
export interface AnnotationPermission {
	// +k8s:deepcopy-gen=true
	dashboard?: dashboard.AnnotationActions;
	// +k8s:deepcopy-gen=true
	organization?: dashboard.AnnotationActions;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [AnnotationPermissionBuilder](./builder-AnnotationPermissionBuilder.md)

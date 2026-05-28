---
title: <span class="badge object-type-interface"></span> TransformationKind
---
# <span class="badge object-type-interface"></span> TransformationKind

## Definition

```typescript
export interface TransformationKind {
	kind: "Transformation";
	// The group is the transformation ID
	group: string;
	spec: dashboardv2.TransformationSpec;
}

```
## See also

 * <span class="badge builder"></span> [TransformationBuilder](./builder-TransformationBuilder.md)

---
title: <span class="badge object-type-interface"></span> V2TransformationKind
---
# <span class="badge object-type-interface"></span> V2TransformationKind

Dashboard v2 shape: the transformation ID moved from `kind` to `group`.

## Definition

```typescript
export interface V2TransformationKind {
	kind: "Transformation";
	// The group is the transformation ID
	group: string;
	spec: dashboardv2beta1.V2TransformationSpec;
}

```
## See also

 * <span class="badge builder"></span> [V2TransformationKindBuilder](./builder-V2TransformationKindBuilder.md)

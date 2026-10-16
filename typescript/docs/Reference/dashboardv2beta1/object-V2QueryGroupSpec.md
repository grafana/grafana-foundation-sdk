---
title: <span class="badge object-type-interface"></span> V2QueryGroupSpec
---
# <span class="badge object-type-interface"></span> V2QueryGroupSpec

## Definition

```typescript
export interface V2QueryGroupSpec {
	queries: dashboardv2beta1.PanelQueryKind[];
	transformations: dashboardv2beta1.V2TransformationKind[];
	queryOptions: dashboardv2beta1.QueryOptionsSpec;
}

```
## See also

 * <span class="badge builder"></span> [V2QueryGroupSpecBuilder](./builder-V2QueryGroupSpecBuilder.md)

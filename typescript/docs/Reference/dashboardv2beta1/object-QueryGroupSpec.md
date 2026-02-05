---
title: <span class="badge object-type-interface"></span> QueryGroupSpec
---
# <span class="badge object-type-interface"></span> QueryGroupSpec

## Definition

```typescript
export interface QueryGroupSpec {
	queries: dashboardv2beta1.PanelQueryKind[];
	transformations: dashboardv2beta1.TransformationKind[];
	queryOptions: dashboardv2beta1.QueryOptionsSpec;
}

```

---
title: <span class="badge object-type-interface"></span> ConditionalRenderingGroupSpec
---
# <span class="badge object-type-interface"></span> ConditionalRenderingGroupSpec

## Definition

```typescript
export interface ConditionalRenderingGroupSpec {
	visibility: "show" | "hide";
	condition: "and" | "or";
	items: (dashboardv2beta1.ConditionalRenderingVariableKind | dashboardv2beta1.ConditionalRenderingDataKind | dashboardv2beta1.ConditionalRenderingTimeRangeSizeKind)[];
}

```

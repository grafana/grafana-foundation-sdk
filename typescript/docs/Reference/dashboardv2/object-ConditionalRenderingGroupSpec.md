---
title: <span class="badge object-type-interface"></span> ConditionalRenderingGroupSpec
---
# <span class="badge object-type-interface"></span> ConditionalRenderingGroupSpec

## Definition

```typescript
export interface ConditionalRenderingGroupSpec {
	visibility: "show" | "hide";
	condition: "and" | "or";
	items: (dashboardv2.ConditionalRenderingVariableKind | dashboardv2.ConditionalRenderingDataKind | dashboardv2.ConditionalRenderingTimeRangeSizeKind)[];
}

```

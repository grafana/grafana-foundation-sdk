---
title: <span class="badge object-type-interface"></span> TabsLayoutTabSpec
---
# <span class="badge object-type-interface"></span> TabsLayoutTabSpec

## Definition

```typescript
export interface TabsLayoutTabSpec {
	title?: string;
	layout: dashboardv2.GridLayoutKind | dashboardv2.RowsLayoutKind | dashboardv2.AutoGridLayoutKind | dashboardv2.TabsLayoutKind;
	conditionalRendering?: dashboardv2.ConditionalRenderingGroupKind;
	repeat?: dashboardv2.TabRepeatOptions;
	variables?: dashboardv2.VariableKind[];
}

```

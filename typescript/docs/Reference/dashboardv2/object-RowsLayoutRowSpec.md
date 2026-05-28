---
title: <span class="badge object-type-interface"></span> RowsLayoutRowSpec
---
# <span class="badge object-type-interface"></span> RowsLayoutRowSpec

## Definition

```typescript
export interface RowsLayoutRowSpec {
	title?: string;
	collapse?: boolean;
	hideHeader?: boolean;
	fillScreen?: boolean;
	conditionalRendering?: dashboardv2.ConditionalRenderingGroupKind;
	repeat?: dashboardv2.RowRepeatOptions;
	layout: dashboardv2.GridLayoutKind | dashboardv2.AutoGridLayoutKind | dashboardv2.TabsLayoutKind | dashboardv2.RowsLayoutKind;
	variables?: dashboardv2.VariableKind[];
}

```

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
	conditionalRendering?: dashboardv2beta1.ConditionalRenderingGroupKind;
	repeat?: dashboardv2beta1.RowRepeatOptions;
	layout: dashboardv2beta1.GridLayoutKind | dashboardv2beta1.AutoGridLayoutKind | dashboardv2beta1.TabsLayoutKind | dashboardv2beta1.RowsLayoutKind;
}

```

---
title: <span class="badge object-type-interface"></span> TabsLayoutTabSpec
---
# <span class="badge object-type-interface"></span> TabsLayoutTabSpec

## Definition

```typescript
export interface TabsLayoutTabSpec {
	title?: string;
	layout: dashboardv2beta1.GridLayoutKind | dashboardv2beta1.RowsLayoutKind | dashboardv2beta1.AutoGridLayoutKind | dashboardv2beta1.TabsLayoutKind;
	conditionalRendering?: dashboardv2beta1.ConditionalRenderingGroupKind;
	repeat?: dashboardv2beta1.TabRepeatOptions;
}

```
## See also

 * <span class="badge builder"></span> [TabsLayoutTabSpecBuilder](./builder-TabsLayoutTabSpecBuilder.md)

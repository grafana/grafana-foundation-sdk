---
title: <span class="badge object-type-interface"></span> AutoGridLayoutSpec
---
# <span class="badge object-type-interface"></span> AutoGridLayoutSpec

## Definition

```typescript
export interface AutoGridLayoutSpec {
	maxColumnCount?: number;
	columnWidthMode: "narrow" | "standard" | "wide" | "custom";
	columnWidth?: number;
	rowHeightMode: "short" | "standard" | "tall" | "custom";
	rowHeight?: number;
	fillScreen?: boolean;
	items: dashboardv2beta1.AutoGridLayoutItemKind[];
}

```
## See also

 * <span class="badge builder"></span> [AutoGridLayoutSpecBuilder](./builder-AutoGridLayoutSpecBuilder.md)

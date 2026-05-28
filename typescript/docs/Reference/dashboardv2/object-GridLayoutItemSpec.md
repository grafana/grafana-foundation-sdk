---
title: <span class="badge object-type-interface"></span> GridLayoutItemSpec
---
# <span class="badge object-type-interface"></span> GridLayoutItemSpec

## Definition

```typescript
export interface GridLayoutItemSpec {
	x: number;
	y: number;
	width: number;
	height: number;
	// reference to a PanelKind from dashboard.spec.elements Expressed as JSON Schema reference
	element: dashboardv2.ElementReference;
	repeat?: dashboardv2.RepeatOptions;
}

```

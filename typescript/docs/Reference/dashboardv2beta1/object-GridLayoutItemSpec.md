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
	element: dashboardv2beta1.ElementReference;
	repeat?: dashboardv2beta1.RepeatOptions;
}

```

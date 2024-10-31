---
title: <span class="badge object-type-interface"></span> CanvasElementOptions
---
# <span class="badge object-type-interface"></span> CanvasElementOptions

## Definition

```typescript
export interface CanvasElementOptions {
	name: string;
	type: string;
	// TODO: figure out how to define this (element config(s))
	config?: any;
	constraint?: canvas.Constraint;
	placement?: canvas.Placement;
	background?: canvas.BackgroundConfig;
	border?: canvas.LineConfig;
	connections?: canvas.CanvasConnection[];
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [CanvasElementOptionsBuilder](./builder-CanvasElementOptionsBuilder.md)

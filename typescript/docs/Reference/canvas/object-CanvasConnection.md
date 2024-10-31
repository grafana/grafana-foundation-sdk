---
title: <span class="badge object-type-interface"></span> CanvasConnection
---
# <span class="badge object-type-interface"></span> CanvasConnection

## Definition

```typescript
export interface CanvasConnection {
	source: canvas.ConnectionCoordinates;
	target: canvas.ConnectionCoordinates;
	targetName?: string;
	path: canvas.ConnectionPath;
	color?: common.ColorDimensionConfig;
	size?: common.ScaleDimensionConfig;
	vertices?: canvas.ConnectionCoordinates[];
	sourceOriginal?: canvas.ConnectionCoordinates;
	targetOriginal?: canvas.ConnectionCoordinates;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [CanvasConnectionBuilder](./builder-CanvasConnectionBuilder.md)

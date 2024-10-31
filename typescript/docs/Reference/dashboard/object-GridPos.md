---
title: <span class="badge object-type-interface"></span> GridPos
---
# <span class="badge object-type-interface"></span> GridPos

Position and dimensions of a panel in the grid

## Definition

```typescript
export interface GridPos {
	// Panel height. The height is the number of rows from the top edge of the panel.
	h: number;
	// Panel width. The width is the number of columns from the left edge of the panel.
	w: number;
	// Panel x. The x coordinate is the number of columns from the left edge of the grid
	x: number;
	// Panel y. The y coordinate is the number of rows from the top edge of the grid
	y: number;
	// Whether the panel is fixed within the grid. If true, the panel will not be affected by other panels' interactions
	static?: boolean;
}

```
## Methods

No methods.

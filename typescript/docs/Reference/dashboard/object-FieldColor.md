---
title: <span class="badge object-type-interface"></span> FieldColor
---
# <span class="badge object-type-interface"></span> FieldColor

Map a field to a color.

## Definition

```typescript
export interface FieldColor {
	// The main color scheme mode.
	mode: dashboard.FieldColorModeId;
	// The fixed color value for fixed or shades color modes.
	fixedColor?: string;
	// Some visualizations need to know how to assign a series color from by value color schemes.
	seriesBy?: dashboard.FieldColorSeriesByMode;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [FieldColorBuilder](./builder-FieldColorBuilder.md)

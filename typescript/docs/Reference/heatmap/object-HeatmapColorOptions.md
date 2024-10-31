---
title: <span class="badge object-type-interface"></span> HeatmapColorOptions
---
# <span class="badge object-type-interface"></span> HeatmapColorOptions

Controls various color options

## Definition

```typescript
export interface HeatmapColorOptions {
	// Sets the color mode
	mode?: heatmap.HeatmapColorMode;
	// Controls the color scheme used
	scheme: string;
	// Controls the color fill when in opacity mode
	fill: string;
	// Controls the color scale
	scale?: heatmap.HeatmapColorScale;
	// Controls the exponent when scale is set to exponential
	exponent: number;
	// Controls the number of color steps
	steps: number;
	// Reverses the color scheme
	reverse: boolean;
	// Sets the minimum value for the color scale
	min?: number;
	// Sets the maximum value for the color scale
	max?: number;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [HeatmapColorOptionsBuilder](./builder-HeatmapColorOptionsBuilder.md)

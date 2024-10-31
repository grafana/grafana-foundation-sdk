---
title: <span class="badge object-type-interface"></span> HeatmapTooltip
---
# <span class="badge object-type-interface"></span> HeatmapTooltip

Controls tooltip options

## Definition

```typescript
export interface HeatmapTooltip {
	// Controls how the tooltip is shown
	mode: common.TooltipDisplayMode;
	maxHeight?: number;
	maxWidth?: number;
	// Controls if the tooltip shows a histogram of the y-axis values
	yHistogram?: boolean;
	// Controls if the tooltip shows a color scale in header
	showColorScale?: boolean;
}

```
## Methods

No methods.
## See also

 * <span class="badge builder"></span> [HeatmapTooltipBuilder](./builder-HeatmapTooltipBuilder.md)

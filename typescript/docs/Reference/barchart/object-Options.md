---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Manually select which field from the dataset to represent the x field.
	xField?: string;
	// Use the color value for a sibling field to color each bar value.
	colorByField?: string;
	// Controls the orientation of the bar chart, either vertical or horizontal.
	orientation: common.VizOrientation;
	// Controls the radius of each bar.
	barRadius?: number;
	// Controls the rotation of the x axis labels.
	xTickLabelRotation: number;
	// Sets the max length that a label can have before it is truncated.
	xTickLabelMaxLength: number;
	// Controls the spacing between x axis labels.
	// negative values indicate backwards skipping behavior
	xTickLabelSpacing?: number;
	// Controls whether bars are stacked or not, either normally or in percent mode.
	stacking: common.StackingMode;
	// This controls whether values are shown on top or to the left of bars.
	showValue: common.VisibilityMode;
	// Controls the width of bars. 1 = Max width, 0 = Min width.
	barWidth: number;
	// Controls the width of groups. 1 = max with, 0 = min width.
	groupWidth: number;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	text?: common.VizTextDisplayOptions;
	// Enables mode which highlights the entire bar area and shows tooltip when cursor
	// hovers over highlighted area
	fullHighlight: boolean;
}

```
## Methods

No methods.

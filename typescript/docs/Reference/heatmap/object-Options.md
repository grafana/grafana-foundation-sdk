---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Controls if the heatmap should be calculated from data
	calculate?: boolean;
	// Calculation options for the heatmap
	calculation?: common.HeatmapCalculationOptions;
	// Controls the color options
	color: heatmap.HeatmapColorOptions;
	// Filters values between a given range
	filterValues?: heatmap.FilterValueRange;
	// Controls tick alignment and value name when not calculating from data
	rowsFrame?: heatmap.RowsHeatmapOptions;
	// | *{
	// 	layout: ui.HeatmapCellLayout & "auto" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
	// }
	// Controls the display of the value in the cell
	showValue: common.VisibilityMode;
	// Controls gap between cells
	cellGap?: number;
	// Controls cell radius
	cellRadius?: number;
	// Controls cell value unit
	cellValues?: heatmap.CellValues;
	// Controls yAxis placement
	yAxis: heatmap.YAxisConfig;
	// | *{
	// 	axisPlacement: ui.AxisPlacement & "left" // TODO: fix after remove when https://github.com/grafana/cuetsy/issues/74 is fixed
	// }
	// Controls legend options
	legend: heatmap.HeatmapLegend;
	// Controls tooltip options
	tooltip: heatmap.HeatmapTooltip;
	// Controls exemplar options
	exemplars: heatmap.ExemplarConfig;
	// Controls which axis to allow selection on
	selectionMode?: heatmap.HeatmapSelectionMode;
}

```
## Methods

No methods.

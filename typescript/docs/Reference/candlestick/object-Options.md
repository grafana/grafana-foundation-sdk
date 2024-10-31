---
title: <span class="badge object-type-interface"></span> Options
---
# <span class="badge object-type-interface"></span> Options

## Definition

```typescript
export interface Options {
	// Sets which dimensions are used for the visualization
	mode: candlestick.VizDisplayMode;
	// Sets the style of the candlesticks
	candleStyle: candlestick.CandleStyle;
	// Sets the color strategy for the candlesticks
	colorStrategy: candlestick.ColorStrategy;
	// Map fields to appropriate dimension
	fields: candlestick.CandlestickFieldMap;
	// Set which colors are used when the price movement is up or down
	colors: candlestick.CandlestickColors;
	legend: common.VizLegendOptions;
	tooltip: common.VizTooltipOptions;
	// When enabled, all fields will be sent to the graph
	includeAllFields?: boolean;
}

```
## Methods

No methods.

---
title: <span class="badge object-type-enum"></span> TableCellDisplayMode
---
# <span class="badge object-type-enum"></span> TableCellDisplayMode

Internally, this is the "type" of cell that's being displayed

in the table such as colored text, JSON, gauge, etc.

The color-background-solid, gradient-gauge, and lcd-gauge

modes are deprecated in favor of new cell subOptions

## Definition

```typescript
export enum TableCellDisplayMode {
	Auto = "auto",
	ColorText = "color-text",
	ColorBackground = "color-background",
	ColorBackgroundSolid = "color-background-solid",
	GradientGauge = "gradient-gauge",
	LcdGauge = "lcd-gauge",
	JSONView = "json-view",
	BasicGauge = "basic",
	Image = "image",
	Gauge = "gauge",
	Sparkline = "sparkline",
	DataLinks = "data-links",
	Custom = "custom",
	Actions = "actions",
}

```

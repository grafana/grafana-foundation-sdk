---
title: <span class="badge object-type-enum"></span> FieldColorModeId
---
# <span class="badge object-type-enum"></span> FieldColorModeId

Color mode for a field. You can specify a single color, or select a continuous (gradient) color schemes, based on a value.

Continuous color interpolates a color using the percentage of a value relative to min and max.

Accepted values are:

`thresholds`: From thresholds. Informs Grafana to take the color from the matching threshold

`palette-classic`: Classic palette. Grafana will assign color by looking up a color in a palette by series index. Useful for Graphs and pie charts and other categorical data visualizations

`palette-classic-by-name`: Classic palette (by name). Grafana will assign color by looking up a color in a palette by series name. Useful for Graphs and pie charts and other categorical data visualizations

`continuous-GrYlRd`: ontinuous Green-Yellow-Red palette mode

`continuous-RdYlGr`: Continuous Red-Yellow-Green palette mode

`continuous-BlYlRd`: Continuous Blue-Yellow-Red palette mode

`continuous-YlRd`: Continuous Yellow-Red palette mode

`continuous-BlPu`: Continuous Blue-Purple palette mode

`continuous-YlBl`: Continuous Yellow-Blue palette mode

`continuous-blues`: Continuous Blue palette mode

`continuous-reds`: Continuous Red palette mode

`continuous-greens`: Continuous Green palette mode

`continuous-purples`: Continuous Purple palette mode

`shades`: Shades of a single color. Specify a single color, useful in an override rule.

`fixed`: Fixed color mode. Specify a single color, useful in an override rule.

## Definition

```typescript
export enum FieldColorModeId {
	Thresholds = "thresholds",
	PaletteClassic = "palette-classic",
	PaletteClassicByName = "palette-classic-by-name",
	ContinuousGrYlRd = "continuous-GrYlRd",
	ContinuousRdYlGr = "continuous-RdYlGr",
	ContinuousBlYlRd = "continuous-BlYlRd",
	ContinuousYlRd = "continuous-YlRd",
	ContinuousBlPu = "continuous-BlPu",
	ContinuousYlBl = "continuous-YlBl",
	ContinuousBlues = "continuous-blues",
	ContinuousReds = "continuous-reds",
	ContinuousGreens = "continuous-greens",
	ContinuousPurples = "continuous-purples",
	Fixed = "fixed",
	Shades = "shades",
}

```

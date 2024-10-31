---
title: <span class="badge object-type-enum"></span> TableCellDisplayMode
---
# <span class="badge object-type-enum"></span> TableCellDisplayMode

Internally, this is the "type" of cell that's being displayed

in the table such as colored text, JSON, gauge, etc.

The color-background-solid, gradient-gauge, and lcd-gauge

modes are deprecated in favor of new cell subOptions

## Definition

```go
type TableCellDisplayMode string
const (
	TableCellDisplayModeAuto TableCellDisplayMode = "auto"
	TableCellDisplayModeColorText TableCellDisplayMode = "color-text"
	TableCellDisplayModeColorBackground TableCellDisplayMode = "color-background"
	TableCellDisplayModeColorBackgroundSolid TableCellDisplayMode = "color-background-solid"
	TableCellDisplayModeGradientGauge TableCellDisplayMode = "gradient-gauge"
	TableCellDisplayModeLcdGauge TableCellDisplayMode = "lcd-gauge"
	TableCellDisplayModeJSONView TableCellDisplayMode = "json-view"
	TableCellDisplayModeBasicGauge TableCellDisplayMode = "basic"
	TableCellDisplayModeImage TableCellDisplayMode = "image"
	TableCellDisplayModeGauge TableCellDisplayMode = "gauge"
	TableCellDisplayModeSparkline TableCellDisplayMode = "sparkline"
	TableCellDisplayModeDataLinks TableCellDisplayMode = "data-links"
	TableCellDisplayModeCustom TableCellDisplayMode = "custom"
)

```

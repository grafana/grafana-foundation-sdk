---
title: <span class="badge builder"></span> HeatmapTooltipBuilder
---
# <span class="badge builder"></span> HeatmapTooltipBuilder

## Constructor

```go
func NewHeatmapTooltipBuilder() *HeatmapTooltipBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *HeatmapTooltipBuilder) Build() (HeatmapTooltip, error)
```

### <span class="badge object-method"></span> Mode

Controls how the tooltip is shown

```go
func (builder *HeatmapTooltipBuilder) Mode(mode common.TooltipDisplayMode) *HeatmapTooltipBuilder
```

### <span class="badge object-method"></span> ShowColorScale

Controls if the tooltip shows a color scale in header

```go
func (builder *HeatmapTooltipBuilder) ShowColorScale(showColorScale bool) *HeatmapTooltipBuilder
```

### <span class="badge object-method"></span> YHistogram

Controls if the tooltip shows a histogram of the y-axis values

```go
func (builder *HeatmapTooltipBuilder) YHistogram(yHistogram bool) *HeatmapTooltipBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [HeatmapTooltip](./object-HeatmapTooltip.md)

---
title: <span class="badge builder"></span> ScatterSeriesConfigBuilder
---
# <span class="badge builder"></span> ScatterSeriesConfigBuilder

## Constructor

```go
func NewScatterSeriesConfigBuilder() *ScatterSeriesConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ScatterSeriesConfigBuilder) Build() (ScatterSeriesConfig, error)
```

### <span class="badge object-method"></span> AxisBorderShow

```go
func (builder *ScatterSeriesConfigBuilder) AxisBorderShow(axisBorderShow bool) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> AxisCenteredZero

```go
func (builder *ScatterSeriesConfigBuilder) AxisCenteredZero(axisCenteredZero bool) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> AxisColorMode

```go
func (builder *ScatterSeriesConfigBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> AxisGridShow

```go
func (builder *ScatterSeriesConfigBuilder) AxisGridShow(axisGridShow bool) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> AxisLabel

```go
func (builder *ScatterSeriesConfigBuilder) AxisLabel(axisLabel string) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> AxisPlacement

```go
func (builder *ScatterSeriesConfigBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMax

```go
func (builder *ScatterSeriesConfigBuilder) AxisSoftMax(axisSoftMax float64) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMin

```go
func (builder *ScatterSeriesConfigBuilder) AxisSoftMin(axisSoftMin float64) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> AxisWidth

```go
func (builder *ScatterSeriesConfigBuilder) AxisWidth(axisWidth float64) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> Frame

```go
func (builder *ScatterSeriesConfigBuilder) Frame(frame float64) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *ScatterSeriesConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *ScatterSeriesConfigBuilder) Label(label common.VisibilityMode) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> LabelValue

```go
func (builder *ScatterSeriesConfigBuilder) LabelValue(labelValue cog.Builder[common.TextDimensionConfig]) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> LineColor

```go
func (builder *ScatterSeriesConfigBuilder) LineColor(lineColor cog.Builder[common.ColorDimensionConfig]) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *ScatterSeriesConfigBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *ScatterSeriesConfigBuilder) LineWidth(lineWidth int32) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *ScatterSeriesConfigBuilder) Name(name string) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> PointColor

```go
func (builder *ScatterSeriesConfigBuilder) PointColor(pointColor cog.Builder[common.ColorDimensionConfig]) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *ScatterSeriesConfigBuilder) PointSize(pointSize cog.Builder[common.ScaleDimensionConfig]) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *ScatterSeriesConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> Show

```go
func (builder *ScatterSeriesConfigBuilder) Show(show xychart.ScatterShow) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> X

```go
func (builder *ScatterSeriesConfigBuilder) X(x string) *ScatterSeriesConfigBuilder
```

### <span class="badge object-method"></span> Y

```go
func (builder *ScatterSeriesConfigBuilder) Y(y string) *ScatterSeriesConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ScatterSeriesConfig](./object-ScatterSeriesConfig.md)

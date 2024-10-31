---
title: <span class="badge builder"></span> TableSparklineCellOptionsBuilder
---
# <span class="badge builder"></span> TableSparklineCellOptionsBuilder

## Constructor

```go
func NewTableSparklineCellOptionsBuilder() *TableSparklineCellOptionsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TableSparklineCellOptionsBuilder) Build() (TableSparklineCellOptions, error)
```

### <span class="badge object-method"></span> AxisCenteredZero

```go
func (builder *TableSparklineCellOptionsBuilder) AxisCenteredZero(axisCenteredZero bool) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> AxisColorMode

```go
func (builder *TableSparklineCellOptionsBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> AxisGridShow

```go
func (builder *TableSparklineCellOptionsBuilder) AxisGridShow(axisGridShow bool) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> AxisLabel

```go
func (builder *TableSparklineCellOptionsBuilder) AxisLabel(axisLabel string) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> AxisPlacement

```go
func (builder *TableSparklineCellOptionsBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> AxisSoftMax

```go
func (builder *TableSparklineCellOptionsBuilder) AxisSoftMax(axisSoftMax float64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> AxisSoftMin

```go
func (builder *TableSparklineCellOptionsBuilder) AxisSoftMin(axisSoftMin float64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> AxisWidth

```go
func (builder *TableSparklineCellOptionsBuilder) AxisWidth(axisWidth float64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> BarAlignment

```go
func (builder *TableSparklineCellOptionsBuilder) BarAlignment(barAlignment common.BarAlignment) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> BarMaxWidth

```go
func (builder *TableSparklineCellOptionsBuilder) BarMaxWidth(barMaxWidth float64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> BarWidthFactor

```go
func (builder *TableSparklineCellOptionsBuilder) BarWidthFactor(barWidthFactor float64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> DrawStyle

```go
func (builder *TableSparklineCellOptionsBuilder) DrawStyle(drawStyle common.GraphDrawStyle) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> FillBelowTo

```go
func (builder *TableSparklineCellOptionsBuilder) FillBelowTo(fillBelowTo string) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> FillColor

```go
func (builder *TableSparklineCellOptionsBuilder) FillColor(fillColor string) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> FillOpacity

```go
func (builder *TableSparklineCellOptionsBuilder) FillOpacity(fillOpacity float64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> GradientMode

```go
func (builder *TableSparklineCellOptionsBuilder) GradientMode(gradientMode common.GraphGradientMode) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *TableSparklineCellOptionsBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> LineColor

```go
func (builder *TableSparklineCellOptionsBuilder) LineColor(lineColor string) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> LineInterpolation

```go
func (builder *TableSparklineCellOptionsBuilder) LineInterpolation(lineInterpolation common.LineInterpolation) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *TableSparklineCellOptionsBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *TableSparklineCellOptionsBuilder) LineWidth(lineWidth float64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> PointColor

```go
func (builder *TableSparklineCellOptionsBuilder) PointColor(pointColor string) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *TableSparklineCellOptionsBuilder) PointSize(pointSize float64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> PointSymbol

```go
func (builder *TableSparklineCellOptionsBuilder) PointSymbol(pointSymbol string) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *TableSparklineCellOptionsBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> ShowPoints

```go
func (builder *TableSparklineCellOptionsBuilder) ShowPoints(showPoints common.VisibilityMode) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> SpanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```go
func (builder *TableSparklineCellOptionsBuilder) SpanNulls(spanNulls common.BoolOrFloat64) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> Stacking

```go
func (builder *TableSparklineCellOptionsBuilder) Stacking(stacking cog.Builder[common.StackingConfig]) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> ThresholdsStyle

```go
func (builder *TableSparklineCellOptionsBuilder) ThresholdsStyle(thresholdsStyle cog.Builder[common.GraphThresholdsStyleConfig]) *TableSparklineCellOptionsBuilder
```

### <span class="badge object-method"></span> Transform

```go
func (builder *TableSparklineCellOptionsBuilder) Transform(transform common.GraphTransform) *TableSparklineCellOptionsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TableSparklineCellOptions](./object-TableSparklineCellOptions.md)

---
title: <span class="badge builder"></span> GraphFieldConfigBuilder
---
# <span class="badge builder"></span> GraphFieldConfigBuilder

## Constructor

```go
func NewGraphFieldConfigBuilder() *GraphFieldConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *GraphFieldConfigBuilder) Build() (GraphFieldConfig, error)
```

### <span class="badge object-method"></span> AxisBorderShow

```go
func (builder *GraphFieldConfigBuilder) AxisBorderShow(axisBorderShow bool) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> AxisCenteredZero

```go
func (builder *GraphFieldConfigBuilder) AxisCenteredZero(axisCenteredZero bool) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> AxisColorMode

```go
func (builder *GraphFieldConfigBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> AxisGridShow

```go
func (builder *GraphFieldConfigBuilder) AxisGridShow(axisGridShow bool) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> AxisLabel

```go
func (builder *GraphFieldConfigBuilder) AxisLabel(axisLabel string) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> AxisPlacement

```go
func (builder *GraphFieldConfigBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMax

```go
func (builder *GraphFieldConfigBuilder) AxisSoftMax(axisSoftMax float64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMin

```go
func (builder *GraphFieldConfigBuilder) AxisSoftMin(axisSoftMin float64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> AxisWidth

```go
func (builder *GraphFieldConfigBuilder) AxisWidth(axisWidth float64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> BarAlignment

```go
func (builder *GraphFieldConfigBuilder) BarAlignment(barAlignment common.BarAlignment) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> BarMaxWidth

```go
func (builder *GraphFieldConfigBuilder) BarMaxWidth(barMaxWidth float64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> BarWidthFactor

```go
func (builder *GraphFieldConfigBuilder) BarWidthFactor(barWidthFactor float64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> DrawStyle

```go
func (builder *GraphFieldConfigBuilder) DrawStyle(drawStyle common.GraphDrawStyle) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> FillBelowTo

```go
func (builder *GraphFieldConfigBuilder) FillBelowTo(fillBelowTo string) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> FillColor

```go
func (builder *GraphFieldConfigBuilder) FillColor(fillColor string) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> FillOpacity

```go
func (builder *GraphFieldConfigBuilder) FillOpacity(fillOpacity float64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> GradientMode

```go
func (builder *GraphFieldConfigBuilder) GradientMode(gradientMode common.GraphGradientMode) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *GraphFieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> InsertNulls

```go
func (builder *GraphFieldConfigBuilder) InsertNulls(insertNulls common.BoolOrFloat64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> LineColor

```go
func (builder *GraphFieldConfigBuilder) LineColor(lineColor string) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> LineInterpolation

```go
func (builder *GraphFieldConfigBuilder) LineInterpolation(lineInterpolation common.LineInterpolation) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *GraphFieldConfigBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *GraphFieldConfigBuilder) LineWidth(lineWidth float64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> PointColor

```go
func (builder *GraphFieldConfigBuilder) PointColor(pointColor string) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *GraphFieldConfigBuilder) PointSize(pointSize float64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> PointSymbol

```go
func (builder *GraphFieldConfigBuilder) PointSymbol(pointSymbol string) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *GraphFieldConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> ShowPoints

```go
func (builder *GraphFieldConfigBuilder) ShowPoints(showPoints common.VisibilityMode) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> SpanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```go
func (builder *GraphFieldConfigBuilder) SpanNulls(spanNulls common.BoolOrFloat64) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> Stacking

```go
func (builder *GraphFieldConfigBuilder) Stacking(stacking cog.Builder[common.StackingConfig]) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> ThresholdsStyle

```go
func (builder *GraphFieldConfigBuilder) ThresholdsStyle(thresholdsStyle cog.Builder[common.GraphThresholdsStyleConfig]) *GraphFieldConfigBuilder
```

### <span class="badge object-method"></span> Transform

```go
func (builder *GraphFieldConfigBuilder) Transform(transform common.GraphTransform) *GraphFieldConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [GraphFieldConfig](./object-GraphFieldConfig.md)

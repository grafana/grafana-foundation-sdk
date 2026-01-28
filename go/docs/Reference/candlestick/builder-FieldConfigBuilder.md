---
title: <span class="badge builder"></span> FieldConfigBuilder
---
# <span class="badge builder"></span> FieldConfigBuilder

## Constructor

```go
func NewFieldConfigBuilder() *FieldConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FieldConfigBuilder) Build() (FieldConfig, error)
```

### <span class="badge object-method"></span> AxisBorderShow

```go
func (builder *FieldConfigBuilder) AxisBorderShow(axisBorderShow bool) *FieldConfigBuilder
```

### <span class="badge object-method"></span> AxisCenteredZero

```go
func (builder *FieldConfigBuilder) AxisCenteredZero(axisCenteredZero bool) *FieldConfigBuilder
```

### <span class="badge object-method"></span> AxisColorMode

```go
func (builder *FieldConfigBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *FieldConfigBuilder
```

### <span class="badge object-method"></span> AxisGridShow

```go
func (builder *FieldConfigBuilder) AxisGridShow(axisGridShow bool) *FieldConfigBuilder
```

### <span class="badge object-method"></span> AxisLabel

```go
func (builder *FieldConfigBuilder) AxisLabel(axisLabel string) *FieldConfigBuilder
```

### <span class="badge object-method"></span> AxisPlacement

```go
func (builder *FieldConfigBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *FieldConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMax

```go
func (builder *FieldConfigBuilder) AxisSoftMax(axisSoftMax float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMin

```go
func (builder *FieldConfigBuilder) AxisSoftMin(axisSoftMin float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> AxisWidth

```go
func (builder *FieldConfigBuilder) AxisWidth(axisWidth float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> BarAlignment

```go
func (builder *FieldConfigBuilder) BarAlignment(barAlignment common.BarAlignment) *FieldConfigBuilder
```

### <span class="badge object-method"></span> BarMaxWidth

```go
func (builder *FieldConfigBuilder) BarMaxWidth(barMaxWidth float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> BarWidthFactor

```go
func (builder *FieldConfigBuilder) BarWidthFactor(barWidthFactor float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> DrawStyle

```go
func (builder *FieldConfigBuilder) DrawStyle(drawStyle common.GraphDrawStyle) *FieldConfigBuilder
```

### <span class="badge object-method"></span> FillBelowTo

```go
func (builder *FieldConfigBuilder) FillBelowTo(fillBelowTo string) *FieldConfigBuilder
```

### <span class="badge object-method"></span> FillColor

```go
func (builder *FieldConfigBuilder) FillColor(fillColor string) *FieldConfigBuilder
```

### <span class="badge object-method"></span> FillOpacity

```go
func (builder *FieldConfigBuilder) FillOpacity(fillOpacity float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> GradientMode

```go
func (builder *FieldConfigBuilder) GradientMode(gradientMode common.GraphGradientMode) *FieldConfigBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *FieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> InsertNulls

```go
func (builder *FieldConfigBuilder) InsertNulls(insertNulls common.BoolOrFloat64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineColor

```go
func (builder *FieldConfigBuilder) LineColor(lineColor string) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineInterpolation

```go
func (builder *FieldConfigBuilder) LineInterpolation(lineInterpolation common.LineInterpolation) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *FieldConfigBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *FieldConfigBuilder) LineWidth(lineWidth float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> PointColor

```go
func (builder *FieldConfigBuilder) PointColor(pointColor string) *FieldConfigBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *FieldConfigBuilder) PointSize(pointSize float64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> PointSymbol

```go
func (builder *FieldConfigBuilder) PointSymbol(pointSymbol string) *FieldConfigBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *FieldConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> ShowPoints

```go
func (builder *FieldConfigBuilder) ShowPoints(showPoints common.VisibilityMode) *FieldConfigBuilder
```

### <span class="badge object-method"></span> SpanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```go
func (builder *FieldConfigBuilder) SpanNulls(spanNulls common.BoolOrFloat64) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Stacking

```go
func (builder *FieldConfigBuilder) Stacking(stacking cog.Builder[common.StackingConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> ThresholdsStyle

```go
func (builder *FieldConfigBuilder) ThresholdsStyle(thresholdsStyle cog.Builder[common.GraphThresholdsStyleConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Transform

```go
func (builder *FieldConfigBuilder) Transform(transform common.GraphTransform) *FieldConfigBuilder
```

## See also

 * <span class="badge object-type-ref"></span> [FieldConfig](./object-FieldConfig.md)

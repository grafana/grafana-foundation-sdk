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

### <span class="badge object-method"></span> HideFrom

```go
func (builder *FieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Label

```go
func (builder *FieldConfigBuilder) Label(label common.VisibilityMode) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LabelValue

```go
func (builder *FieldConfigBuilder) LabelValue(labelValue cog.Builder[common.TextDimensionConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineColor

```go
func (builder *FieldConfigBuilder) LineColor(lineColor cog.Builder[common.ColorDimensionConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *FieldConfigBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *FieldConfigBuilder) LineWidth(lineWidth int32) *FieldConfigBuilder
```

### <span class="badge object-method"></span> PointColor

```go
func (builder *FieldConfigBuilder) PointColor(pointColor cog.Builder[common.ColorDimensionConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *FieldConfigBuilder) PointSize(pointSize cog.Builder[common.ScaleDimensionConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *FieldConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Show

```go
func (builder *FieldConfigBuilder) Show(show xychart.ScatterShow) *FieldConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)

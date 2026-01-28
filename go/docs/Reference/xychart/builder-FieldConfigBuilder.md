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

### <span class="badge object-method"></span> FillOpacity

```go
func (builder *FieldConfigBuilder) FillOpacity(fillOpacity uint32) *FieldConfigBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *FieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *FieldConfigBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *FieldConfigBuilder) LineWidth(lineWidth int32) *FieldConfigBuilder
```

### <span class="badge object-method"></span> PointShape

```go
func (builder *FieldConfigBuilder) PointShape(pointShape xychart.PointShape) *FieldConfigBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *FieldConfigBuilder) PointSize(pointSize cog.Builder[xychart.XychartFieldConfigPointSize]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> PointStrokeWidth

```go
func (builder *FieldConfigBuilder) PointStrokeWidth(pointStrokeWidth int32) *FieldConfigBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *FieldConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Show

```go
func (builder *FieldConfigBuilder) Show(show xychart.XYShowMode) *FieldConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)

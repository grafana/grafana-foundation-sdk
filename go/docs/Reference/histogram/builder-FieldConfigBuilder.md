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

Controls the fill opacity of the bars.

```go
func (builder *FieldConfigBuilder) FillOpacity(fillOpacity uint32) *FieldConfigBuilder
```

### <span class="badge object-method"></span> GradientMode

Set the mode of the gradient fill. Fill gradient is based on the line color. To change the color, use the standard color scheme field option.

Gradient appearance is influenced by the Fill opacity setting.

```go
func (builder *FieldConfigBuilder) GradientMode(gradientMode common.GraphGradientMode) *FieldConfigBuilder
```

### <span class="badge object-method"></span> HideFrom

```go
func (builder *FieldConfigBuilder) HideFrom(hideFrom cog.Builder[common.HideSeriesConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> LineWidth

Controls line width of the bars.

```go
func (builder *FieldConfigBuilder) LineWidth(lineWidth uint32) *FieldConfigBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *FieldConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *FieldConfigBuilder
```

### <span class="badge object-method"></span> Stacking

```go
func (builder *FieldConfigBuilder) Stacking(stacking cog.Builder[common.StackingConfig]) *FieldConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FieldConfig](./object-FieldConfig.md)

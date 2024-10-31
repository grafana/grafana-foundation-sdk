---
title: <span class="badge builder"></span> YAxisConfigBuilder
---
# <span class="badge builder"></span> YAxisConfigBuilder

## Constructor

```go
func NewYAxisConfigBuilder() *YAxisConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *YAxisConfigBuilder) Build() (YAxisConfig, error)
```

### <span class="badge object-method"></span> AxisCenteredZero

```go
func (builder *YAxisConfigBuilder) AxisCenteredZero(axisCenteredZero bool) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> AxisColorMode

```go
func (builder *YAxisConfigBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> AxisGridShow

```go
func (builder *YAxisConfigBuilder) AxisGridShow(axisGridShow bool) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> AxisLabel

```go
func (builder *YAxisConfigBuilder) AxisLabel(axisLabel string) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> AxisPlacement

```go
func (builder *YAxisConfigBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMax

```go
func (builder *YAxisConfigBuilder) AxisSoftMax(axisSoftMax float64) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMin

```go
func (builder *YAxisConfigBuilder) AxisSoftMin(axisSoftMin float64) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> AxisWidth

```go
func (builder *YAxisConfigBuilder) AxisWidth(axisWidth float64) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> Decimals

Controls the number of decimals for yAxis values

```go
func (builder *YAxisConfigBuilder) Decimals(decimals float32) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> Max

Sets the maximum value for the yAxis

```go
func (builder *YAxisConfigBuilder) Max(max float32) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> Min

Sets the minimum value for the yAxis

```go
func (builder *YAxisConfigBuilder) Min(min float32) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> Reverse

Reverses the yAxis

```go
func (builder *YAxisConfigBuilder) Reverse(reverse bool) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *YAxisConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *YAxisConfigBuilder
```

### <span class="badge object-method"></span> Unit

Sets the yAxis unit

```go
func (builder *YAxisConfigBuilder) Unit(unit string) *YAxisConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [YAxisConfig](./object-YAxisConfig.md)

---
title: <span class="badge builder"></span> AxisConfigBuilder
---
# <span class="badge builder"></span> AxisConfigBuilder

## Constructor

```go
func NewAxisConfigBuilder() *AxisConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *AxisConfigBuilder) Build() (AxisConfig, error)
```

### <span class="badge object-method"></span> AxisCenteredZero

```go
func (builder *AxisConfigBuilder) AxisCenteredZero(axisCenteredZero bool) *AxisConfigBuilder
```

### <span class="badge object-method"></span> AxisColorMode

```go
func (builder *AxisConfigBuilder) AxisColorMode(axisColorMode common.AxisColorMode) *AxisConfigBuilder
```

### <span class="badge object-method"></span> AxisGridShow

```go
func (builder *AxisConfigBuilder) AxisGridShow(axisGridShow bool) *AxisConfigBuilder
```

### <span class="badge object-method"></span> AxisLabel

```go
func (builder *AxisConfigBuilder) AxisLabel(axisLabel string) *AxisConfigBuilder
```

### <span class="badge object-method"></span> AxisPlacement

```go
func (builder *AxisConfigBuilder) AxisPlacement(axisPlacement common.AxisPlacement) *AxisConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMax

```go
func (builder *AxisConfigBuilder) AxisSoftMax(axisSoftMax float64) *AxisConfigBuilder
```

### <span class="badge object-method"></span> AxisSoftMin

```go
func (builder *AxisConfigBuilder) AxisSoftMin(axisSoftMin float64) *AxisConfigBuilder
```

### <span class="badge object-method"></span> AxisWidth

```go
func (builder *AxisConfigBuilder) AxisWidth(axisWidth float64) *AxisConfigBuilder
```

### <span class="badge object-method"></span> ScaleDistribution

```go
func (builder *AxisConfigBuilder) ScaleDistribution(scaleDistribution cog.Builder[common.ScaleDistributionConfig]) *AxisConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [AxisConfig](./object-AxisConfig.md)

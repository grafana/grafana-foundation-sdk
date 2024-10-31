---
title: <span class="badge builder"></span> PointsConfigBuilder
---
# <span class="badge builder"></span> PointsConfigBuilder

## Constructor

```go
func NewPointsConfigBuilder() *PointsConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PointsConfigBuilder) Build() (PointsConfig, error)
```

### <span class="badge object-method"></span> PointColor

```go
func (builder *PointsConfigBuilder) PointColor(pointColor string) *PointsConfigBuilder
```

### <span class="badge object-method"></span> PointSize

```go
func (builder *PointsConfigBuilder) PointSize(pointSize float64) *PointsConfigBuilder
```

### <span class="badge object-method"></span> PointSymbol

```go
func (builder *PointsConfigBuilder) PointSymbol(pointSymbol string) *PointsConfigBuilder
```

### <span class="badge object-method"></span> ShowPoints

```go
func (builder *PointsConfigBuilder) ShowPoints(showPoints common.VisibilityMode) *PointsConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PointsConfig](./object-PointsConfig.md)

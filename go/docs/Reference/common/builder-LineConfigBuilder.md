---
title: <span class="badge builder"></span> LineConfigBuilder
---
# <span class="badge builder"></span> LineConfigBuilder

## Constructor

```go
func NewLineConfigBuilder() *LineConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *LineConfigBuilder) Build() (LineConfig, error)
```

### <span class="badge object-method"></span> LineColor

```go
func (builder *LineConfigBuilder) LineColor(lineColor string) *LineConfigBuilder
```

### <span class="badge object-method"></span> LineInterpolation

```go
func (builder *LineConfigBuilder) LineInterpolation(lineInterpolation common.LineInterpolation) *LineConfigBuilder
```

### <span class="badge object-method"></span> LineStyle

```go
func (builder *LineConfigBuilder) LineStyle(lineStyle cog.Builder[common.LineStyle]) *LineConfigBuilder
```

### <span class="badge object-method"></span> LineWidth

```go
func (builder *LineConfigBuilder) LineWidth(lineWidth float64) *LineConfigBuilder
```

### <span class="badge object-method"></span> SpanNulls

Indicate if null values should be treated as gaps or connected.

When the value is a number, it represents the maximum delta in the

X axis that should be considered connected.  For timeseries, this is milliseconds

```go
func (builder *LineConfigBuilder) SpanNulls(spanNulls common.BoolOrFloat64) *LineConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LineConfig](./object-LineConfig.md)

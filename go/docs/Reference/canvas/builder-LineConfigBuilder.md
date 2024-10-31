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

### <span class="badge object-method"></span> Color

```go
func (builder *LineConfigBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *LineConfigBuilder
```

### <span class="badge object-method"></span> Width

```go
func (builder *LineConfigBuilder) Width(width float64) *LineConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [LineConfig](./object-LineConfig.md)

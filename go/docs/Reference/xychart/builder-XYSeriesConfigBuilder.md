---
title: <span class="badge builder"></span> XYSeriesConfigBuilder
---
# <span class="badge builder"></span> XYSeriesConfigBuilder

## Constructor

```go
func NewXYSeriesConfigBuilder() *XYSeriesConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *XYSeriesConfigBuilder) Build() (XYSeriesConfig, error)
```

### <span class="badge object-method"></span> Color

```go
func (builder *XYSeriesConfigBuilder) Color(color cog.Builder[xychart.XychartXYSeriesConfigColor]) *XYSeriesConfigBuilder
```

### <span class="badge object-method"></span> Frame

```go
func (builder *XYSeriesConfigBuilder) Frame(frame cog.Builder[xychart.XychartXYSeriesConfigFrame]) *XYSeriesConfigBuilder
```

### <span class="badge object-method"></span> Name

```go
func (builder *XYSeriesConfigBuilder) Name(name cog.Builder[xychart.XychartXYSeriesConfigName]) *XYSeriesConfigBuilder
```

### <span class="badge object-method"></span> Size

```go
func (builder *XYSeriesConfigBuilder) Size(size cog.Builder[xychart.XychartXYSeriesConfigSize]) *XYSeriesConfigBuilder
```

### <span class="badge object-method"></span> X

```go
func (builder *XYSeriesConfigBuilder) X(x cog.Builder[xychart.XychartXYSeriesConfigX]) *XYSeriesConfigBuilder
```

### <span class="badge object-method"></span> Y

```go
func (builder *XYSeriesConfigBuilder) Y(y cog.Builder[xychart.XychartXYSeriesConfigY]) *XYSeriesConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [XYSeriesConfig](./object-XYSeriesConfig.md)

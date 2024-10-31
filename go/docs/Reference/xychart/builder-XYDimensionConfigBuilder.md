---
title: <span class="badge builder"></span> XYDimensionConfigBuilder
---
# <span class="badge builder"></span> XYDimensionConfigBuilder

## Constructor

```go
func NewXYDimensionConfigBuilder() *XYDimensionConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *XYDimensionConfigBuilder) Build() (XYDimensionConfig, error)
```

### <span class="badge object-method"></span> Exclude

```go
func (builder *XYDimensionConfigBuilder) Exclude(exclude []string) *XYDimensionConfigBuilder
```

### <span class="badge object-method"></span> Frame

```go
func (builder *XYDimensionConfigBuilder) Frame(frame int32) *XYDimensionConfigBuilder
```

### <span class="badge object-method"></span> X

```go
func (builder *XYDimensionConfigBuilder) X(x string) *XYDimensionConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [XYDimensionConfig](./object-XYDimensionConfig.md)

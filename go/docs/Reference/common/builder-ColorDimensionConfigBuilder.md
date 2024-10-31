---
title: <span class="badge builder"></span> ColorDimensionConfigBuilder
---
# <span class="badge builder"></span> ColorDimensionConfigBuilder

## Constructor

```go
func NewColorDimensionConfigBuilder() *ColorDimensionConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ColorDimensionConfigBuilder) Build() (ColorDimensionConfig, error)
```

### <span class="badge object-method"></span> Field

fixed: T -- will be added by each element

```go
func (builder *ColorDimensionConfigBuilder) Field(field string) *ColorDimensionConfigBuilder
```

### <span class="badge object-method"></span> Fixed

color value

```go
func (builder *ColorDimensionConfigBuilder) Fixed(fixed string) *ColorDimensionConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ColorDimensionConfig](./object-ColorDimensionConfig.md)

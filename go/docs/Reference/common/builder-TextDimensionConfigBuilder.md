---
title: <span class="badge builder"></span> TextDimensionConfigBuilder
---
# <span class="badge builder"></span> TextDimensionConfigBuilder

## Constructor

```go
func NewTextDimensionConfigBuilder() *TextDimensionConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *TextDimensionConfigBuilder) Build() (TextDimensionConfig, error)
```

### <span class="badge object-method"></span> Field

fixed: T -- will be added by each element

```go
func (builder *TextDimensionConfigBuilder) Field(field string) *TextDimensionConfigBuilder
```

### <span class="badge object-method"></span> Fixed

```go
func (builder *TextDimensionConfigBuilder) Fixed(fixed string) *TextDimensionConfigBuilder
```

### <span class="badge object-method"></span> Mode

```go
func (builder *TextDimensionConfigBuilder) Mode(mode common.TextDimensionMode) *TextDimensionConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [TextDimensionConfig](./object-TextDimensionConfig.md)

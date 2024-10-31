---
title: <span class="badge builder"></span> BackgroundConfigBuilder
---
# <span class="badge builder"></span> BackgroundConfigBuilder

## Constructor

```go
func NewBackgroundConfigBuilder() *BackgroundConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BackgroundConfigBuilder) Build() (BackgroundConfig, error)
```

### <span class="badge object-method"></span> Color

```go
func (builder *BackgroundConfigBuilder) Color(color cog.Builder[common.ColorDimensionConfig]) *BackgroundConfigBuilder
```

### <span class="badge object-method"></span> Image

```go
func (builder *BackgroundConfigBuilder) Image(image cog.Builder[common.ResourceDimensionConfig]) *BackgroundConfigBuilder
```

### <span class="badge object-method"></span> Size

```go
func (builder *BackgroundConfigBuilder) Size(size canvas.BackgroundImageSize) *BackgroundConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BackgroundConfig](./object-BackgroundConfig.md)

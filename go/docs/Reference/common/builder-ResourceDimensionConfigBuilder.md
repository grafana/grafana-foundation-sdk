---
title: <span class="badge builder"></span> ResourceDimensionConfigBuilder
---
# <span class="badge builder"></span> ResourceDimensionConfigBuilder

## Constructor

```go
func NewResourceDimensionConfigBuilder() *ResourceDimensionConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ResourceDimensionConfigBuilder) Build() (ResourceDimensionConfig, error)
```

### <span class="badge object-method"></span> Field

fixed: T -- will be added by each element

```go
func (builder *ResourceDimensionConfigBuilder) Field(field string) *ResourceDimensionConfigBuilder
```

### <span class="badge object-method"></span> Fixed

```go
func (builder *ResourceDimensionConfigBuilder) Fixed(fixed string) *ResourceDimensionConfigBuilder
```

### <span class="badge object-method"></span> Mode

```go
func (builder *ResourceDimensionConfigBuilder) Mode(mode common.ResourceDimensionMode) *ResourceDimensionConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ResourceDimensionConfig](./object-ResourceDimensionConfig.md)

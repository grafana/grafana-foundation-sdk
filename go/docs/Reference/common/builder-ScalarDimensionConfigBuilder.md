---
title: <span class="badge builder"></span> ScalarDimensionConfigBuilder
---
# <span class="badge builder"></span> ScalarDimensionConfigBuilder

## Constructor

```go
func NewScalarDimensionConfigBuilder() *ScalarDimensionConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ScalarDimensionConfigBuilder) Build() (ScalarDimensionConfig, error)
```

### <span class="badge object-method"></span> Field

fixed: T -- will be added by each element

```go
func (builder *ScalarDimensionConfigBuilder) Field(field string) *ScalarDimensionConfigBuilder
```

### <span class="badge object-method"></span> Fixed

```go
func (builder *ScalarDimensionConfigBuilder) Fixed(fixed float64) *ScalarDimensionConfigBuilder
```

### <span class="badge object-method"></span> Max

```go
func (builder *ScalarDimensionConfigBuilder) Max(max float64) *ScalarDimensionConfigBuilder
```

### <span class="badge object-method"></span> Min

```go
func (builder *ScalarDimensionConfigBuilder) Min(min float64) *ScalarDimensionConfigBuilder
```

### <span class="badge object-method"></span> Mode

```go
func (builder *ScalarDimensionConfigBuilder) Mode(mode common.ScalarDimensionMode) *ScalarDimensionConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ScalarDimensionConfig](./object-ScalarDimensionConfig.md)

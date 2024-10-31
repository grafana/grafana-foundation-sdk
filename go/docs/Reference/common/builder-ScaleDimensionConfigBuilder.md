---
title: <span class="badge builder"></span> ScaleDimensionConfigBuilder
---
# <span class="badge builder"></span> ScaleDimensionConfigBuilder

## Constructor

```go
func NewScaleDimensionConfigBuilder() *ScaleDimensionConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ScaleDimensionConfigBuilder) Build() (ScaleDimensionConfig, error)
```

### <span class="badge object-method"></span> Field

fixed: T -- will be added by each element

```go
func (builder *ScaleDimensionConfigBuilder) Field(field string) *ScaleDimensionConfigBuilder
```

### <span class="badge object-method"></span> Fixed

```go
func (builder *ScaleDimensionConfigBuilder) Fixed(fixed float64) *ScaleDimensionConfigBuilder
```

### <span class="badge object-method"></span> Max

```go
func (builder *ScaleDimensionConfigBuilder) Max(max float64) *ScaleDimensionConfigBuilder
```

### <span class="badge object-method"></span> Min

```go
func (builder *ScaleDimensionConfigBuilder) Min(min float64) *ScaleDimensionConfigBuilder
```

### <span class="badge object-method"></span> Mode

| *"linear"

```go
func (builder *ScaleDimensionConfigBuilder) Mode(mode common.ScaleDimensionMode) *ScaleDimensionConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ScaleDimensionConfig](./object-ScaleDimensionConfig.md)

---
title: <span class="badge builder"></span> ScaleDistributionConfigBuilder
---
# <span class="badge builder"></span> ScaleDistributionConfigBuilder

## Constructor

```go
func NewScaleDistributionConfigBuilder() *ScaleDistributionConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ScaleDistributionConfigBuilder) Build() (ScaleDistributionConfig, error)
```

### <span class="badge object-method"></span> LinearThreshold

```go
func (builder *ScaleDistributionConfigBuilder) LinearThreshold(linearThreshold float64) *ScaleDistributionConfigBuilder
```

### <span class="badge object-method"></span> Log

```go
func (builder *ScaleDistributionConfigBuilder) Log(log float64) *ScaleDistributionConfigBuilder
```

### <span class="badge object-method"></span> Type

```go
func (builder *ScaleDistributionConfigBuilder) Type(typeArg common.ScaleDistribution) *ScaleDistributionConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ScaleDistributionConfig](./object-ScaleDistributionConfig.md)

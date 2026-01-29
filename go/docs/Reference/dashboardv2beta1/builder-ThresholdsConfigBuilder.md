---
title: <span class="badge builder"></span> ThresholdsConfigBuilder
---
# <span class="badge builder"></span> ThresholdsConfigBuilder

## Constructor

```go
func NewThresholdsConfigBuilder() *ThresholdsConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *ThresholdsConfigBuilder) Build() (ThresholdsConfig, error)
```

### <span class="badge object-method"></span> Mode

```go
func (builder *ThresholdsConfigBuilder) Mode(mode dashboardv2beta1.ThresholdsMode) *ThresholdsConfigBuilder
```

### <span class="badge object-method"></span> Steps

```go
func (builder *ThresholdsConfigBuilder) Steps(steps []dashboardv2beta1.Threshold) *ThresholdsConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ThresholdsConfig](./object-ThresholdsConfig.md)

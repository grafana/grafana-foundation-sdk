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

Thresholds mode.

```go
func (builder *ThresholdsConfigBuilder) Mode(mode dashboard.ThresholdsMode) *ThresholdsConfigBuilder
```

### <span class="badge object-method"></span> Steps

Must be sorted by 'value', first value is always -Infinity

```go
func (builder *ThresholdsConfigBuilder) Steps(steps []dashboard.Threshold) *ThresholdsConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [ThresholdsConfig](./object-ThresholdsConfig.md)

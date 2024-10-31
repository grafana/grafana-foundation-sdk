---
title: <span class="badge builder"></span> BarConfigBuilder
---
# <span class="badge builder"></span> BarConfigBuilder

## Constructor

```go
func NewBarConfigBuilder() *BarConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *BarConfigBuilder) Build() (BarConfig, error)
```

### <span class="badge object-method"></span> BarAlignment

```go
func (builder *BarConfigBuilder) BarAlignment(barAlignment common.BarAlignment) *BarConfigBuilder
```

### <span class="badge object-method"></span> BarMaxWidth

```go
func (builder *BarConfigBuilder) BarMaxWidth(barMaxWidth float64) *BarConfigBuilder
```

### <span class="badge object-method"></span> BarWidthFactor

```go
func (builder *BarConfigBuilder) BarWidthFactor(barWidthFactor float64) *BarConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [BarConfig](./object-BarConfig.md)

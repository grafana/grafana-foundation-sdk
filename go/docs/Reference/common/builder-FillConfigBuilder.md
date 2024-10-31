---
title: <span class="badge builder"></span> FillConfigBuilder
---
# <span class="badge builder"></span> FillConfigBuilder

## Constructor

```go
func NewFillConfigBuilder() *FillConfigBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *FillConfigBuilder) Build() (FillConfig, error)
```

### <span class="badge object-method"></span> FillBelowTo

```go
func (builder *FillConfigBuilder) FillBelowTo(fillBelowTo string) *FillConfigBuilder
```

### <span class="badge object-method"></span> FillColor

```go
func (builder *FillConfigBuilder) FillColor(fillColor string) *FillConfigBuilder
```

### <span class="badge object-method"></span> FillOpacity

```go
func (builder *FillConfigBuilder) FillOpacity(fillOpacity float64) *FillConfigBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [FillConfig](./object-FillConfig.md)

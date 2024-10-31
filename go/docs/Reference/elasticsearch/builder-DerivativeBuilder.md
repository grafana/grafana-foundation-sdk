---
title: <span class="badge builder"></span> DerivativeBuilder
---
# <span class="badge builder"></span> DerivativeBuilder

## Constructor

```go
func NewDerivativeBuilder() *DerivativeBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *DerivativeBuilder) Build() (Derivative, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *DerivativeBuilder) Field(field string) *DerivativeBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *DerivativeBuilder) Hide(hide bool) *DerivativeBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *DerivativeBuilder) Id(id string) *DerivativeBuilder
```

### <span class="badge object-method"></span> PipelineAgg

```go
func (builder *DerivativeBuilder) PipelineAgg(pipelineAgg string) *DerivativeBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *DerivativeBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchDerivativeSettings]) *DerivativeBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [Derivative](./object-Derivative.md)

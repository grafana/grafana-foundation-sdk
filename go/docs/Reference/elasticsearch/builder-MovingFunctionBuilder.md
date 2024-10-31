---
title: <span class="badge builder"></span> MovingFunctionBuilder
---
# <span class="badge builder"></span> MovingFunctionBuilder

## Constructor

```go
func NewMovingFunctionBuilder() *MovingFunctionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MovingFunctionBuilder) Build() (MovingFunction, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *MovingFunctionBuilder) Field(field string) *MovingFunctionBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *MovingFunctionBuilder) Hide(hide bool) *MovingFunctionBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *MovingFunctionBuilder) Id(id string) *MovingFunctionBuilder
```

### <span class="badge object-method"></span> PipelineAgg

```go
func (builder *MovingFunctionBuilder) PipelineAgg(pipelineAgg string) *MovingFunctionBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MovingFunctionBuilder) Settings(settings cog.Builder[elasticsearch.ElasticsearchMovingFunctionSettings]) *MovingFunctionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MovingFunction](./object-MovingFunction.md)

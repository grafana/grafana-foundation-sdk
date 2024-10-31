---
title: <span class="badge builder"></span> MovingAverageBuilder
---
# <span class="badge builder"></span> MovingAverageBuilder

## Constructor

```go
func NewMovingAverageBuilder() *MovingAverageBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MovingAverageBuilder) Build() (MovingAverage, error)
```

### <span class="badge object-method"></span> Field

```go
func (builder *MovingAverageBuilder) Field(field string) *MovingAverageBuilder
```

### <span class="badge object-method"></span> Hide

```go
func (builder *MovingAverageBuilder) Hide(hide bool) *MovingAverageBuilder
```

### <span class="badge object-method"></span> Id

```go
func (builder *MovingAverageBuilder) Id(id string) *MovingAverageBuilder
```

### <span class="badge object-method"></span> PipelineAgg

```go
func (builder *MovingAverageBuilder) PipelineAgg(pipelineAgg string) *MovingAverageBuilder
```

### <span class="badge object-method"></span> Settings

```go
func (builder *MovingAverageBuilder) Settings(settings map[string]any) *MovingAverageBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MovingAverage](./object-MovingAverage.md)

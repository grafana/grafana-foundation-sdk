---
title: <span class="badge builder"></span> MovingAverageModelOptionBuilder
---
# <span class="badge builder"></span> MovingAverageModelOptionBuilder

## Constructor

```go
func NewMovingAverageModelOptionBuilder() *MovingAverageModelOptionBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *MovingAverageModelOptionBuilder) Build() (MovingAverageModelOption, error)
```

### <span class="badge object-method"></span> Label

```go
func (builder *MovingAverageModelOptionBuilder) Label(label string) *MovingAverageModelOptionBuilder
```

### <span class="badge object-method"></span> Value

```go
func (builder *MovingAverageModelOptionBuilder) Value(value elasticsearch.MovingAverageModel) *MovingAverageModelOptionBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [MovingAverageModelOption](./object-MovingAverageModelOption.md)

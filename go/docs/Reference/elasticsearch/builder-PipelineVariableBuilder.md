---
title: <span class="badge builder"></span> PipelineVariableBuilder
---
# <span class="badge builder"></span> PipelineVariableBuilder

## Constructor

```go
func NewPipelineVariableBuilder() *PipelineVariableBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *PipelineVariableBuilder) Build() (PipelineVariable, error)
```

### <span class="badge object-method"></span> Name

```go
func (builder *PipelineVariableBuilder) Name(name string) *PipelineVariableBuilder
```

### <span class="badge object-method"></span> PipelineAgg

```go
func (builder *PipelineVariableBuilder) PipelineAgg(pipelineAgg string) *PipelineVariableBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [PipelineVariable](./object-PipelineVariable.md)

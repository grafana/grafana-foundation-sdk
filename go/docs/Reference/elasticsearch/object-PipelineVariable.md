---
title: <span class="badge object-type-struct"></span> PipelineVariable
---
# <span class="badge object-type-struct"></span> PipelineVariable

## Definition

```go
type PipelineVariable struct {
    Name string `json:"name"`
    PipelineAgg string `json:"pipelineAgg"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `PipelineVariable` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (pipelineVariable *PipelineVariable) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `PipelineVariable` objects.

```go
func (pipelineVariable *PipelineVariable) Equals(other PipelineVariable) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `PipelineVariable` fields for violations and returns them.

```go
func (pipelineVariable *PipelineVariable) Validate() error
```

## See also

 * <span class="badge builder"></span> [PipelineVariableBuilder](./builder-PipelineVariableBuilder.md)

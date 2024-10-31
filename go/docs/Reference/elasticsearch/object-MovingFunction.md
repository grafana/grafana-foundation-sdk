---
title: <span class="badge object-type-struct"></span> MovingFunction
---
# <span class="badge object-type-struct"></span> MovingFunction

## Definition

```go
type MovingFunction struct {
    PipelineAgg *string `json:"pipelineAgg,omitempty"`
    Field *string `json:"field,omitempty"`
    Type string `json:"type"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchMovingFunctionSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingFunction` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingFunction *MovingFunction) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingFunction` objects.

```go
func (movingFunction *MovingFunction) Equals(other MovingFunction) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingFunction` fields for violations and returns them.

```go
func (movingFunction *MovingFunction) Validate() error
```

## See also

 * <span class="badge builder"></span> [MovingFunctionBuilder](./builder-MovingFunctionBuilder.md)

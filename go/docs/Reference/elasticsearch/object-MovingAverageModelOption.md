---
title: <span class="badge object-type-struct"></span> MovingAverageModelOption
---
# <span class="badge object-type-struct"></span> MovingAverageModelOption

## Definition

```go
type MovingAverageModelOption struct {
    Label string `json:"label"`
    Value elasticsearch.MovingAverageModel `json:"value"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverageModelOption` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingAverageModelOption *MovingAverageModelOption) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingAverageModelOption` objects.

```go
func (movingAverageModelOption *MovingAverageModelOption) Equals(other MovingAverageModelOption) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingAverageModelOption` fields for violations and returns them.

```go
func (movingAverageModelOption *MovingAverageModelOption) Validate() error
```

## See also

 * <span class="badge builder"></span> [MovingAverageModelOptionBuilder](./builder-MovingAverageModelOptionBuilder.md)

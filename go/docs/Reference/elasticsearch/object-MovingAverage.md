---
title: <span class="badge object-type-struct"></span> MovingAverage
---
# <span class="badge object-type-struct"></span> MovingAverage

#MovingAverage's settings are overridden in types.ts

## Definition

```go
type MovingAverage struct {
    PipelineAgg *string `json:"pipelineAgg,omitempty"`
    Field *string `json:"field,omitempty"`
    Type string `json:"type"`
    Id string `json:"id"`
    Settings map[string]any `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `MovingAverage` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (movingAverage *MovingAverage) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `MovingAverage` objects.

```go
func (movingAverage *MovingAverage) Equals(other MovingAverage) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `MovingAverage` fields for violations and returns them.

```go
func (movingAverage *MovingAverage) Validate() error
```

## See also

 * <span class="badge builder"></span> [MovingAverageBuilder](./builder-MovingAverageBuilder.md)

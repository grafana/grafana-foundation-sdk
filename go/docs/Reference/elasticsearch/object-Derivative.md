---
title: <span class="badge object-type-struct"></span> Derivative
---
# <span class="badge object-type-struct"></span> Derivative

## Definition

```go
type Derivative struct {
    PipelineAgg *string `json:"pipelineAgg,omitempty"`
    Field *string `json:"field,omitempty"`
    Type string `json:"type"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchDerivativeSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Derivative` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (derivative *Derivative) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Derivative` objects.

```go
func (derivative *Derivative) Equals(other Derivative) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Derivative` fields for violations and returns them.

```go
func (derivative *Derivative) Validate() error
```

## See also

 * <span class="badge builder"></span> [DerivativeBuilder](./builder-DerivativeBuilder.md)

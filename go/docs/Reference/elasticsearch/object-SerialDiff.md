---
title: <span class="badge object-type-struct"></span> SerialDiff
---
# <span class="badge object-type-struct"></span> SerialDiff

## Definition

```go
type SerialDiff struct {
    PipelineAgg *string `json:"pipelineAgg,omitempty"`
    Field *string `json:"field,omitempty"`
    Type string `json:"type"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchSerialDiffSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SerialDiff` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (serialDiff *SerialDiff) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SerialDiff` objects.

```go
func (serialDiff *SerialDiff) Equals(other SerialDiff) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SerialDiff` fields for violations and returns them.

```go
func (serialDiff *SerialDiff) Validate() error
```

## See also

 * <span class="badge builder"></span> [SerialDiffBuilder](./builder-SerialDiffBuilder.md)

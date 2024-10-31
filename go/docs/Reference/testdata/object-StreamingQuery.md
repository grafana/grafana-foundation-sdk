---
title: <span class="badge object-type-struct"></span> StreamingQuery
---
# <span class="badge object-type-struct"></span> StreamingQuery

## Definition

```go
type StreamingQuery struct {
    Bands *int64 `json:"bands,omitempty"`
    Noise float64 `json:"noise"`
    Speed float64 `json:"speed"`
    Spread float64 `json:"spread"`
    // Possible enum values:
    //  - `"fetch"` 
    //  - `"logs"` 
    //  - `"signal"` 
    //  - `"traces"` 
    Type testdata.StreamingQueryType `json:"type"`
    Url *string `json:"url,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `StreamingQuery` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (streamingQuery *StreamingQuery) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `StreamingQuery` objects.

```go
func (streamingQuery *StreamingQuery) Equals(other StreamingQuery) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `StreamingQuery` fields for violations and returns them.

```go
func (streamingQuery *StreamingQuery) Validate() error
```

## See also

 * <span class="badge builder"></span> [StreamingQueryBuilder](./builder-StreamingQueryBuilder.md)

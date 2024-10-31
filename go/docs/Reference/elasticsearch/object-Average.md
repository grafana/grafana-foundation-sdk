---
title: <span class="badge object-type-struct"></span> Average
---
# <span class="badge object-type-struct"></span> Average

## Definition

```go
type Average struct {
    Type string `json:"type"`
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchAverageSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Average` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (average *Average) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Average` objects.

```go
func (average *Average) Equals(other Average) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Average` fields for violations and returns them.

```go
func (average *Average) Validate() error
```

## See also

 * <span class="badge builder"></span> [AverageBuilder](./builder-AverageBuilder.md)

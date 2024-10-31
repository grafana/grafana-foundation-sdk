---
title: <span class="badge object-type-struct"></span> Max
---
# <span class="badge object-type-struct"></span> Max

## Definition

```go
type Max struct {
    Type string `json:"type"`
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchMaxSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Max` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (max *Max) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Max` objects.

```go
func (max *Max) Equals(other Max) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Max` fields for violations and returns them.

```go
func (max *Max) Validate() error
```

## See also

 * <span class="badge builder"></span> [MaxBuilder](./builder-MaxBuilder.md)

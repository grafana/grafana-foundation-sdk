---
title: <span class="badge object-type-struct"></span> Min
---
# <span class="badge object-type-struct"></span> Min

## Definition

```go
type Min struct {
    Type string `json:"type"`
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchMinSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Min` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (min *Min) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Min` objects.

```go
func (min *Min) Equals(other Min) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Min` fields for violations and returns them.

```go
func (min *Min) Validate() error
```

## See also

 * <span class="badge builder"></span> [MinBuilder](./builder-MinBuilder.md)

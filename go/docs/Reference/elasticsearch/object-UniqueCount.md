---
title: <span class="badge object-type-struct"></span> UniqueCount
---
# <span class="badge object-type-struct"></span> UniqueCount

## Definition

```go
type UniqueCount struct {
    Type string `json:"type"`
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Settings *elasticsearch.ElasticsearchUniqueCountSettings `json:"settings,omitempty"`
    Hide *bool `json:"hide,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `UniqueCount` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (uniqueCount *UniqueCount) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `UniqueCount` objects.

```go
func (uniqueCount *UniqueCount) Equals(other UniqueCount) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `UniqueCount` fields for violations and returns them.

```go
func (uniqueCount *UniqueCount) Validate() error
```

## See also

 * <span class="badge builder"></span> [UniqueCountBuilder](./builder-UniqueCountBuilder.md)

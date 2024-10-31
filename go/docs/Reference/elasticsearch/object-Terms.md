---
title: <span class="badge object-type-struct"></span> Terms
---
# <span class="badge object-type-struct"></span> Terms

## Definition

```go
type Terms struct {
    Field *string `json:"field,omitempty"`
    Id string `json:"id"`
    Type string `json:"type"`
    Settings *elasticsearch.ElasticsearchTermsSettings `json:"settings,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `Terms` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (terms *Terms) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `Terms` objects.

```go
func (terms *Terms) Equals(other Terms) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `Terms` fields for violations and returns them.

```go
func (terms *Terms) Validate() error
```

## See also

 * <span class="badge builder"></span> [TermsBuilder](./builder-TermsBuilder.md)

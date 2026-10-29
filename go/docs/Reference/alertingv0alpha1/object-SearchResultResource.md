---
title: <span class="badge object-type-struct"></span> SearchResultResource
---
# <span class="badge object-type-struct"></span> SearchResultResource

#SearchResultResource is the full identity of a hit.

## Definition

```go
type SearchResultResource struct {
    Group string `json:"group"`
    Resource string `json:"resource"`
    Kind string `json:"kind"`
    Name string `json:"name"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SearchResultResource` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (searchResultResource *SearchResultResource) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SearchResultResource` objects.

```go
func (searchResultResource *SearchResultResource) Equals(other SearchResultResource) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SearchResultResource` fields for violations and returns them.

```go
func (searchResultResource *SearchResultResource) Validate() error
```

## See also

 * <span class="badge builder"></span> [SearchResultResourceBuilder](./builder-SearchResultResourceBuilder.md)

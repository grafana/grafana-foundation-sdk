---
title: <span class="badge object-type-struct"></span> SearchResults
---
# <span class="badge object-type-struct"></span> SearchResults

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Definition

```go
type SearchResults struct {
    Metadata alertingv0alpha1.SearchResultsMetadata `json:"metadata"`
    Items []alertingv0alpha1.SearchResultHit `json:"items"`
    Facets map[string][]alertingv0alpha1.FacetValue `json:"facets,omitempty"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SearchResults` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (searchResults *SearchResults) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SearchResults` objects.

```go
func (searchResults *SearchResults) Equals(other SearchResults) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SearchResults` fields for violations and returns them.

```go
func (searchResults *SearchResults) Validate() error
```

## See also

 * <span class="badge builder"></span> [SearchResultsBuilder](./builder-SearchResultsBuilder.md)

---
title: <span class="badge object-type-struct"></span> SearchResultHit
---
# <span class="badge object-type-struct"></span> SearchResultHit

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Definition

```go
type SearchResultHit struct {
    Resource alertingv0alpha1.SearchResultResource `json:"resource"`
    Score *float64 `json:"score,omitempty"`
    Fields alertingv0alpha1.RuleSearchHitFields `json:"fields"`
}
```
## Methods

### <span class="badge object-method"></span> UnmarshalJSONStrict

UnmarshalJSONStrict implements a custom JSON unmarshalling logic to decode `SearchResultHit` from JSON.

Note: the unmarshalling done by this function is strict. It will fail over required fields being absent from the input, fields having an incorrect type, unexpected fields being present, …

```go
func (searchResultHit *SearchResultHit) UnmarshalJSONStrict(raw []byte) error
```

### <span class="badge object-method"></span> Equals

Equals tests the equality of two `SearchResultHit` objects.

```go
func (searchResultHit *SearchResultHit) Equals(other SearchResultHit) bool
```

### <span class="badge object-method"></span> Validate

Validate checks all the validation constraints that may be defined on `SearchResultHit` fields for violations and returns them.

```go
func (searchResultHit *SearchResultHit) Validate() error
```

## See also

 * <span class="badge builder"></span> [SearchResultHitBuilder](./builder-SearchResultHitBuilder.md)

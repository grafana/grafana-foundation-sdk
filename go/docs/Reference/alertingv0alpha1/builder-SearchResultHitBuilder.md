---
title: <span class="badge builder"></span> SearchResultHitBuilder
---
# <span class="badge builder"></span> SearchResultHitBuilder

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Constructor

```go
func NewSearchResultHitBuilder() *SearchResultHitBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SearchResultHitBuilder) Build() (SearchResultHit, error)
```

### <span class="badge object-method"></span> Fields

```go
func (builder *SearchResultHitBuilder) Fields(fields cog.Builder[alertingv0alpha1.RuleSearchHitFields]) *SearchResultHitBuilder
```

### <span class="badge object-method"></span> Resource

```go
func (builder *SearchResultHitBuilder) Resource(resource cog.Builder[alertingv0alpha1.SearchResultResource]) *SearchResultHitBuilder
```

### <span class="badge object-method"></span> Score

```go
func (builder *SearchResultHitBuilder) Score(score float64) *SearchResultHitBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SearchResultHit](./object-SearchResultHit.md)

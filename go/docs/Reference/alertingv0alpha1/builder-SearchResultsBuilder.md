---
title: <span class="badge builder"></span> SearchResultsBuilder
---
# <span class="badge builder"></span> SearchResultsBuilder

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Constructor

```go
func NewSearchResultsBuilder() *SearchResultsBuilder
```
## Methods

### <span class="badge object-method"></span> Build

Builds the object.

```go
func (builder *SearchResultsBuilder) Build() (SearchResults, error)
```

### <span class="badge object-method"></span> Facets

```go
func (builder *SearchResultsBuilder) Facets(facets map[string][]cog.Builder[alertingv0alpha1.FacetValue]) *SearchResultsBuilder
```

### <span class="badge object-method"></span> Items

```go
func (builder *SearchResultsBuilder) Items(items []cog.Builder[alertingv0alpha1.SearchResultHit]) *SearchResultsBuilder
```

### <span class="badge object-method"></span> Metadata

```go
func (builder *SearchResultsBuilder) Metadata(metadata cog.Builder[alertingv0alpha1.SearchResultsMetadata]) *SearchResultsBuilder
```

## See also

 * <span class="badge object-type-struct"></span> [SearchResults](./object-SearchResults.md)

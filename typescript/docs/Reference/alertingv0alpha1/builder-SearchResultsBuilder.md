---
title: <span class="badge builder"></span> SearchResultsBuilder
---
# <span class="badge builder"></span> SearchResultsBuilder

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Constructor

```typescript
new SearchResultsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> facets

```typescript
facets(facets: Record<string, cog.Builder<alertingv0alpha1.FacetValue>[]>)
```

### <span class="badge object-method"></span> items

```typescript
items(items: cog.Builder<alertingv0alpha1.SearchResultHit>[])
```

### <span class="badge object-method"></span> metadata

```typescript
metadata(metadata: cog.Builder<alertingv0alpha1.SearchResultsMetadata>)
```

## See also

 * <span class="badge object-type-interface"></span> [SearchResults](./object-SearchResults.md)

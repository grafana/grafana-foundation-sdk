---
title: <span class="badge object-type-interface"></span> SearchResults
---
# <span class="badge object-type-interface"></span> SearchResults

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Definition

```typescript
export interface SearchResults {
	metadata: alertingv0alpha1.SearchResultsMetadata;
	items: alertingv0alpha1.SearchResultHit[];
	facets?: Record<string, alertingv0alpha1.FacetValue[]>;
}

```
## See also

 * <span class="badge builder"></span> [SearchResultsBuilder](./builder-SearchResultsBuilder.md)

---
title: <span class="badge object-type-interface"></span> SearchResultsMetadata
---
# <span class="badge object-type-interface"></span> SearchResultsMetadata

#SearchResultsMetadata carries the paging token and total authorised match

count.

## Definition

```typescript
export interface SearchResultsMetadata {
	continue?: string;
	totalHits?: number;
	// Always read totalHits together with totalHitsRelation.
	totalHitsRelation?: alertingv0alpha1.TotalHitsRelation;
}

```
## See also

 * <span class="badge builder"></span> [SearchResultsMetadataBuilder](./builder-SearchResultsMetadataBuilder.md)

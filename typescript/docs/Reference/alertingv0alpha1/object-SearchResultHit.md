---
title: <span class="badge object-type-interface"></span> SearchResultHit
---
# <span class="badge object-type-interface"></span> SearchResultHit

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Definition

```typescript
export interface SearchResultHit {
	resource: alertingv0alpha1.SearchResultResource;
	score?: number;
	fields: alertingv0alpha1.RuleSearchHitFields;
}

```
## See also

 * <span class="badge builder"></span> [SearchResultHitBuilder](./builder-SearchResultHitBuilder.md)

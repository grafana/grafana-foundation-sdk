---
title: <span class="badge object-type-enum"></span> TotalHitsRelation
---
# <span class="badge object-type-enum"></span> TotalHitsRelation

#TotalHitsRelation says how totalHits relates to the real number of matching

rules the caller may see: "eq" when it is exact, "lte" when it is an upper

bound because authorisation was applied after the search ranked its results.

## Definition

```typescript
export enum TotalHitsRelation {
	Eq = "eq",
	Lte = "lte",
}

```

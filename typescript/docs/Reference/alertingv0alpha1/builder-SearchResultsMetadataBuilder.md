---
title: <span class="badge builder"></span> SearchResultsMetadataBuilder
---
# <span class="badge builder"></span> SearchResultsMetadataBuilder

#SearchResultsMetadata carries the paging token and total authorised match

count.

## Constructor

```typescript
new SearchResultsMetadataBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```typescript
build()
```

### <span class="badge object-method"></span> continueVal

```typescript
continueVal(continueVal: string)
```

### <span class="badge object-method"></span> totalHits

```typescript
totalHits(totalHits: number)
```

### <span class="badge object-method"></span> totalHitsRelation

Always read totalHits together with totalHitsRelation.

```typescript
totalHitsRelation(totalHitsRelation: alertingv0alpha1.TotalHitsRelation)
```

## See also

 * <span class="badge object-type-interface"></span> [SearchResultsMetadata](./object-SearchResultsMetadata.md)

---
title: <span class="badge builder"></span> SearchResultsMetadataBuilder
---
# <span class="badge builder"></span> SearchResultsMetadataBuilder

#SearchResultsMetadata carries the paging token and total authorised match

count.

## Constructor

```java
new SearchResultsMetadataBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```java
public SearchResultsMetadata build()
```

### <span class="badge object-method"></span> continue

```java
public SearchResultsMetadataBuilder continue(String continueArg)
```

### <span class="badge object-method"></span> totalHits

```java
public SearchResultsMetadataBuilder totalHits(Long totalHits)
```

### <span class="badge object-method"></span> totalHitsRelation

Always read totalHits together with totalHitsRelation.

```java
public SearchResultsMetadataBuilder totalHitsRelation(TotalHitsRelation totalHitsRelation)
```

## See also

 * <span class="badge object-type-class"></span> [SearchResultsMetadata](./object-SearchResultsMetadata.md)

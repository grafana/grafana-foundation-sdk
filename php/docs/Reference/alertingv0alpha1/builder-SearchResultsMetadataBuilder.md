---
title: <span class="badge builder"></span> SearchResultsMetadataBuilder
---
# <span class="badge builder"></span> SearchResultsMetadataBuilder

#SearchResultsMetadata carries the paging token and total authorised match

count.

## Constructor

```php
new SearchResultsMetadataBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> continue

```php
continue(string $continue)
```

### <span class="badge object-method"></span> totalHits

```php
totalHits(int $totalHits)
```

### <span class="badge object-method"></span> totalHitsRelation

Always read totalHits together with totalHitsRelation.

```php
totalHitsRelation(\Grafana\Foundation\Alertingv0alpha1\TotalHitsRelation $totalHitsRelation)
```

## See also

 * <span class="badge object-type-class"></span> [SearchResultsMetadata](./object-SearchResultsMetadata.md)

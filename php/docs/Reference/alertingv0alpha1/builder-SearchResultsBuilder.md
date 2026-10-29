---
title: <span class="badge builder"></span> SearchResultsBuilder
---
# <span class="badge builder"></span> SearchResultsBuilder

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Constructor

```php
new SearchResultsBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> facets

@param array<string, array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\FacetValue>>> $facets

```php
facets(array $facets)
```

### <span class="badge object-method"></span> items

@param array<\Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultHit>> $items

```php
items(array $items)
```

### <span class="badge object-method"></span> metadata

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata> $metadata

```php
metadata(\Grafana\Foundation\Cog\Builder $metadata)
```

## See also

 * <span class="badge object-type-class"></span> [SearchResults](./object-SearchResults.md)

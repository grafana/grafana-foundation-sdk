---
title: <span class="badge builder"></span> SearchResultHitBuilder
---
# <span class="badge builder"></span> SearchResultHitBuilder

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Constructor

```php
new SearchResultHitBuilder()
```
## Methods

### <span class="badge object-method"></span> build

Builds the object.

```php
build()
```

### <span class="badge object-method"></span> fields

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields> $fields

```php
fields(\Grafana\Foundation\Cog\Builder $fields)
```

### <span class="badge object-method"></span> resource

@param \Grafana\Foundation\Cog\Builder<\Grafana\Foundation\Alertingv0alpha1\SearchResultResource> $resource

```php
resource(\Grafana\Foundation\Cog\Builder $resource)
```

### <span class="badge object-method"></span> score

```php
score(float $score)
```

## See also

 * <span class="badge object-type-class"></span> [SearchResultHit](./object-SearchResultHit.md)

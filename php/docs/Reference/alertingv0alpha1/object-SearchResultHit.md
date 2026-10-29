---
title: <span class="badge object-type-class"></span> SearchResultHit
---
# <span class="badge object-type-class"></span> SearchResultHit

#SearchResultHit is a single match: its identity, an optional relevance

score (present only when the query included free text), and the requested

fields.

## Definition

```php
class SearchResultHit implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv0alpha1\SearchResultResource $resource;

    public ?float $score;

    public \Grafana\Foundation\Alertingv0alpha1\RuleSearchHitFields $fields;

}
```
## Methods

### <span class="badge object-method"></span> fromArray

Builds this object from an array.

This function is meant to be used with the return value of `json_decode($json, true)`.

```php
static fromArray(array $inputData)
```

### <span class="badge object-method"></span> jsonSerialize

Returns the data representing this object, preparing it for JSON serialization with `json_encode()`.

```php
jsonSerialize()
```

## See also

 * <span class="badge builder"></span> [SearchResultHitBuilder](./builder-SearchResultHitBuilder.md)

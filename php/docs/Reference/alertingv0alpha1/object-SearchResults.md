---
title: <span class="badge object-type-class"></span> SearchResults
---
# <span class="badge object-type-class"></span> SearchResults

#SearchResults is the response envelope, mirroring

search.grafana.app SearchResults.

## Definition

```php
class SearchResults implements \JsonSerializable
{
    public \Grafana\Foundation\Alertingv0alpha1\SearchResultsMetadata $metadata;

    /**
     * @var array<\Grafana\Foundation\Alertingv0alpha1\SearchResultHit>
     */
    public array $items;

    /**
     * @var array<string, array<\Grafana\Foundation\Alertingv0alpha1\FacetValue>>|null
     */
    public ?array $facets;

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

 * <span class="badge builder"></span> [SearchResultsBuilder](./builder-SearchResultsBuilder.md)

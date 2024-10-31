---
title: <span class="badge object-type-class"></span> ElasticsearchTermsSettings
---
# <span class="badge object-type-class"></span> ElasticsearchTermsSettings

## Definition

```php
class ElasticsearchTermsSettings implements \JsonSerializable
{
    public ?\Grafana\Foundation\Elasticsearch\TermsOrder $order;

    public ?string $size;

    public ?string $minDocCount;

    public ?string $orderBy;

    public ?string $missing;

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

 * <span class="badge builder"></span> [ElasticsearchTermsSettingsBuilder](./builder-ElasticsearchTermsSettingsBuilder.md)

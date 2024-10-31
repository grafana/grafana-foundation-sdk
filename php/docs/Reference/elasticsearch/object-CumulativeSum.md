---
title: <span class="badge object-type-class"></span> CumulativeSum
---
# <span class="badge object-type-class"></span> CumulativeSum

## Definition

```php
class CumulativeSum implements \JsonSerializable
{
    public ?string $pipelineAgg;

    public ?string $field;

    public string $type;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchCumulativeSumSettings $settings;

    public ?bool $hide;

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

 * <span class="badge builder"></span> [CumulativeSumBuilder](./builder-CumulativeSumBuilder.md)

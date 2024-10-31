---
title: <span class="badge object-type-class"></span> ExtendedStats
---
# <span class="badge object-type-class"></span> ExtendedStats

## Definition

```php
class ExtendedStats implements \JsonSerializable
{
    public string $type;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchExtendedStatsSettings $settings;

    public ?string $field;

    public string $id;

    /**
     * @var mixed|null
     */
    public $meta;

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

 * <span class="badge builder"></span> [ExtendedStatsBuilder](./builder-ExtendedStatsBuilder.md)

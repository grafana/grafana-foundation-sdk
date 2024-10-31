---
title: <span class="badge object-type-class"></span> Logs
---
# <span class="badge object-type-class"></span> Logs

## Definition

```php
class Logs implements \JsonSerializable
{
    public string $type;

    public string $id;

    public ?\Grafana\Foundation\Elasticsearch\ElasticsearchLogsSettings $settings;

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

 * <span class="badge builder"></span> [LogsBuilder](./builder-LogsBuilder.md)

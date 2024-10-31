---
title: <span class="badge object-type-class"></span> MetricNamesQuery
---
# <span class="badge object-type-class"></span> MetricNamesQuery

## Definition

```php
class MetricNamesQuery implements \JsonSerializable
{
    public ?string $rawQuery;

    public string $kind;

    public string $subscription;

    public string $resourceGroup;

    public string $resourceName;

    public string $metricNamespace;

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

 * <span class="badge builder"></span> [MetricNamesQueryBuilder](./builder-MetricNamesQueryBuilder.md)

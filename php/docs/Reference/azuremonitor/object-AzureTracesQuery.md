---
title: <span class="badge object-type-class"></span> AzureTracesQuery
---
# <span class="badge object-type-class"></span> AzureTracesQuery

Application Insights Traces sub-query properties

## Definition

```php
class AzureTracesQuery implements \JsonSerializable
{
    /**
     * Specifies the format results should be returned as.
     */
    public ?\Grafana\Foundation\Azuremonitor\ResultFormat $resultFormat;

    /**
     * Array of resource URIs to be queried.
     * @var array<string>|null
     */
    public ?array $resources;

    /**
     * Operation ID. Used only for Traces queries.
     */
    public ?string $operationId;

    /**
     * Types of events to filter by.
     * @var array<string>|null
     */
    public ?array $traceTypes;

    /**
     * Filters for property values.
     * @var array<\Grafana\Foundation\Azuremonitor\AzureTracesFilter>|null
     */
    public ?array $filters;

    /**
     * KQL query to be executed.
     */
    public ?string $query;

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

 * <span class="badge builder"></span> [AzureTracesQueryBuilder](./builder-AzureTracesQueryBuilder.md)

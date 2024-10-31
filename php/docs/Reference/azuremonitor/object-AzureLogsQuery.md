---
title: <span class="badge object-type-class"></span> AzureLogsQuery
---
# <span class="badge object-type-class"></span> AzureLogsQuery

Azure Monitor Logs sub-query properties

## Definition

```php
class AzureLogsQuery implements \JsonSerializable
{
    /**
     * KQL query to be executed.
     */
    public ?string $query;

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
     * If set to true the dashboard time range will be used as a filter for the query. Otherwise the query time ranges will be used. Defaults to false.
     */
    public ?bool $dashboardTime;

    /**
     * If dashboardTime is set to true this value dictates which column the time filter will be applied to. Defaults to the first tables timeSpan column, the first datetime column found, or TimeGenerated
     */
    public ?string $timeColumn;

    /**
     * If set to true the query will be run as a basic logs query
     */
    public ?bool $basicLogsQuery;

    /**
     * Workspace ID. This was removed in Grafana 8, but remains for backwards compat.
     */
    public ?string $workspace;

    /**
     * @deprecated Use resources instead
     */
    public ?string $resource;

    /**
     * @deprecated Use dashboardTime instead
     */
    public ?bool $intersectTime;

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

 * <span class="badge builder"></span> [AzureLogsQueryBuilder](./builder-AzureLogsQueryBuilder.md)

---
title: <span class="badge object-type-class"></span> TempoQuery
---
# <span class="badge object-type-class"></span> TempoQuery

## Definition

```php
class TempoQuery implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * A unique identifier for the query within the list of targets.
     * In server side expressions, the refId is used as a variable name to identify results.
     * By default, the UI will assign A->Z; however setting meaningful names may be useful.
     */
    public string $refId;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * Note this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Specify the query flavor
     * TODO make this required and give it a default
     */
    public ?string $queryType;

    /**
     * TraceQL query or trace ID
     */
    public string $query;

    /**
     * @deprecated Logfmt query to filter traces by their tags. Example: http.status_code=200 error=true
     */
    public ?string $search;

    /**
     * @deprecated Query traces by service name
     */
    public ?string $serviceName;

    /**
     * @deprecated Query traces by span name
     */
    public ?string $spanName;

    /**
     * @deprecated Define the minimum duration to select traces. Use duration format, for example: 1.2s, 100ms
     */
    public ?string $minDuration;

    /**
     * @deprecated Define the maximum duration to select traces. Use duration format, for example: 1.2s, 100ms
     */
    public ?string $maxDuration;

    /**
     * Filters to be included in a PromQL query to select data for the service graph. Example: {client="app",service="app"}
     */
    public ?string $serviceMapQuery;

    /**
     * Use service.namespace in addition to service.name to uniquely identify a service.
     */
    public ?bool $serviceMapIncludeNamespace;

    /**
     * Defines the maximum number of traces that are returned from Tempo
     */
    public ?int $limit;

    /**
     * For mixed data sources the selected datasource is on the query level.
     * For non mixed scenarios this is undefined.
     * TODO find a better way to do this ^ that's friendly to schema
     * TODO this shouldn't be unknown but DataSourceRef | null
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * @var array<\Grafana\Foundation\Tempo\TraceqlFilter>
     */
    public array $filters;

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

### <span class="badge object-method"></span> dataqueryType

Returns the type of this dataquery object.

```php
dataqueryType()
```

## See also

 * <span class="badge builder"></span> [TempoQueryBuilder](./builder-TempoQueryBuilder.md)

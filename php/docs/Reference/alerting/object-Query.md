---
title: <span class="badge object-type-class"></span> Query
---
# <span class="badge object-type-class"></span> Query

## Definition

```php
class Query implements \JsonSerializable
{
    /**
     * Grafana data source unique identifier; it should be '__expr__' for a Server Side Expression operation.
     */
    public ?string $datasourceUid;

    /**
     * JSON is the raw JSON query and includes the above properties as well as custom properties.
     * @var \Grafana\Foundation\Cog\Dataquery|null
     */
    public ?\Grafana\Foundation\Cog\Dataquery $model;

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public ?string $queryType;

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public ?string $refId;

    /**
     * RelativeTimeRange is the per query start and end time
     * for requests.
     */
    public ?\Grafana\Foundation\Alerting\RelativeTimeRange $relativeTimeRange;

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

 * <span class="badge builder"></span> [QueryBuilder](./builder-QueryBuilder.md)

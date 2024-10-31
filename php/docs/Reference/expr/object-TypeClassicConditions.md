---
title: <span class="badge object-type-class"></span> TypeClassicConditions
---
# <span class="badge object-type-class"></span> TypeClassicConditions

## Definition

```php
class TypeClassicConditions implements \JsonSerializable, \Grafana\Foundation\Cog\Dataquery
{
    /**
     * @var array<\Grafana\Foundation\Expr\ExprTypeClassicConditionsConditions>
     */
    public array $conditions;

    /**
     * The datasource
     */
    public ?\Grafana\Foundation\Dashboard\DataSourceRef $datasource;

    /**
     * true if query is disabled (ie should not be returned to the dashboard)
     * NOTE: this does not always imply that the query should not be executed since
     * the results from a hidden query may be used as the input to other queries (SSE etc)
     */
    public ?bool $hide;

    /**
     * Interval is the suggested duration between time points in a time series query.
     * NOTE: the values for intervalMs is not saved in the query model.  It is typically calculated
     * from the interval required to fill a pixels in the visualization
     */
    public ?float $intervalMs;

    /**
     * MaxDataPoints is the maximum number of data points that should be returned from a time series query.
     * NOTE: the values for maxDataPoints is not saved in the query model.  It is typically calculated
     * from the number of pixels visible in a visualization
     */
    public ?int $maxDataPoints;

    /**
     * QueryType is an optional identifier for the type of query.
     * It can be used to distinguish different types of queries.
     */
    public ?string $queryType;

    /**
     * RefID is the unique identifier of the query, set by the frontend call.
     */
    public string $refId;

    /**
     * Optionally define expected query result behavior
     */
    public ?\Grafana\Foundation\Expr\ExprTypeClassicConditionsResultAssertions $resultAssertions;

    /**
     * TimeRange represents the query range
     * NOTE: unlike generic /ds/query, we can now send explicit time values in each query
     * NOTE: the values for timeRange are not saved in a dashboard, they are constructed on the fly
     */
    public ?\Grafana\Foundation\Expr\ExprTypeClassicConditionsTimeRange $timeRange;

    public string $type;

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

 * <span class="badge builder"></span> [TypeClassicConditionsBuilder](./builder-TypeClassicConditionsBuilder.md)
